package gitlab

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	gitlab "github.com/xanzy/go-gitlab"
)

func resourceGitlabProjectAccessToken() *schema.Resource {
	// lintignore: XR002 // TODO: Resolve this tfproviderlint issue
	return &schema.Resource{
		CreateContext: resourceGitlabProjectAccessTokenCreate,
		ReadContext:   resourceGitlabProjectAccessTokenRead,
		DeleteContext: resourceGitlabProjectAccessTokenDelete,

		Schema: map[string]*schema.Schema{
			"project": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"scopes": {
				Type:     schema.TypeSet,
				Required: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.StringInSlice([]string{"api", "read_api", "read_repository", "write_repository"}, false),
				},
			},
			"expires_at": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: func(i interface{}, k string) (warnings []string, errors []error) {
					v := i.(string)

					if _, err := time.Parse("2006-01-02", v); err != nil {
						errors = append(errors, fmt.Errorf("expected %q to be a valid YYYY-MM-DD date, got %q: %+v", k, i, err))
					}

					return warnings, errors
				},
				ForceNew: true,
			},
			"token": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
			"active": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"created_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"revoked": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"user_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func resourceGitlabProjectAccessTokenCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*gitlab.Client)

	project := d.Get("project").(int)
	options := &gitlab.CreateProjectAccessTokenOptions{
		Name:   gitlab.String(d.Get("name").(string)),
		Scopes: *stringSetToStringSlice(d.Get("scopes").(*schema.Set)),
	}

	log.Printf("[DEBUG] create gitlab ProjectAccessToken %s %s for project ID %d", *options.Name, options.Scopes, project)

	if v, ok := d.GetOk("expires_at"); ok {
		parsedExpiresAt, err := time.Parse("2006-01-02", v.(string))
		if err != nil {
			return diag.Errorf("Invalid expires_at date: %v", err)
		}
		parsedExpiresAtISOTime := gitlab.ISOTime(parsedExpiresAt)
		options.ExpiresAt = &parsedExpiresAtISOTime
		log.Printf("[DEBUG] create gitlab ProjectAccessToken %s with expires_at %s for project ID %d", *options.Name, *options.ExpiresAt, project)
	}

	projectAccessToken, _, err := client.ProjectAccessTokens.CreateProjectAccessToken(project, options, gitlab.WithContext(ctx))
	if err != nil {
		return diag.FromErr(err)
	}

	log.Printf("[DEBUG] created gitlab ProjectAccessToken %d - %s for project ID %d", projectAccessToken.ID, *options.Name, project)

	projectString := strconv.Itoa(project)
	PATstring := strconv.Itoa(projectAccessToken.ID)
	d.SetId(buildTwoPartID(&projectString, &PATstring))
	d.Set("token", projectAccessToken.Token)

	return resourceGitlabProjectAccessTokenRead(ctx, d, meta)
}

func resourceGitlabProjectAccessTokenRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	projectString, PATstring, err := parseTwoPartID(d.Id())
	if err != nil {
		return diag.Errorf("Error parsing ID: %s", d.Id())
	}

	client := meta.(*gitlab.Client)

	project, err := strconv.Atoi(projectString)
	if err != nil {
		return diag.Errorf("%s cannot be converted to int", projectString)
	}

	projectAccessTokenID, err := strconv.Atoi(PATstring)
	if err != nil {
		return diag.Errorf("%s cannot be converted to int", PATstring)
	}

	log.Printf("[DEBUG] read gitlab ProjectAccessToken %d, project ID %d", projectAccessTokenID, project)

	//there is a slight possibility to not find an existing item, for example
	// 1. item is #101 (ie, in the 2nd page)
	// 2. I load first page (ie. I don't find my target item)
	// 3. A concurrent operation remove item 99 (ie, my target item shift to 1st page)
	// 4. a concurrent operation add an item
	// 5: I load 2nd page  (ie. I don't find my target item)
	// 6. Total pages and total items properties are unchanged (from the perspective of the reader)

	page := 1
	for page != 0 {
		projectAccessTokens, response, err := client.ProjectAccessTokens.ListProjectAccessTokens(project, &gitlab.ListProjectAccessTokensOptions{Page: page, PerPage: 100}, gitlab.WithContext(ctx))
		if err != nil {
			return diag.FromErr(err)
		}

		for _, projectAccessToken := range projectAccessTokens {
			if projectAccessToken.ID == projectAccessTokenID {

				d.Set("project", project)
				d.Set("name", projectAccessToken.Name)
				if projectAccessToken.ExpiresAt != nil {
					d.Set("expires_at", projectAccessToken.ExpiresAt.String())
				}
				d.Set("active", projectAccessToken.Active)
				d.Set("created_at", projectAccessToken.CreatedAt.String())
				d.Set("revoked", projectAccessToken.Revoked)
				d.Set("user_id", projectAccessToken.UserID)
				d.Set("scopes", projectAccessToken.Scopes) // lintignore: R004,XR004 // TODO: Resolve this tfproviderlint issue

				return nil
			}
		}

		page = response.NextPage
	}

	log.Printf("[DEBUG] failed to read gitlab ProjectAccessToken %d, project ID %d", projectAccessTokenID, project)
	d.SetId("")
	return nil
}

func resourceGitlabProjectAccessTokenDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	projectString, PATstring, err := parseTwoPartID(d.Id())
	if err != nil {
		return diag.Errorf("Error parsing ID: %s", d.Id())
	}

	client := meta.(*gitlab.Client)

	project, err := strconv.Atoi(projectString)
	if err != nil {
		return diag.Errorf("%s cannot be converted to int", projectString)
	}

	projectAccessTokenID, err := strconv.Atoi(PATstring)
	if err != nil {
		return diag.Errorf("%s cannot be converted to int", PATstring)
	}

	log.Printf("[DEBUG] Delete gitlab ProjectAccessToken %s", d.Id())
	_, err = client.ProjectAccessTokens.DeleteProjectAccessToken(project, projectAccessTokenID, gitlab.WithContext(ctx))
	return diag.FromErr(err)
}
