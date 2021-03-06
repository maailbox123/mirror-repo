---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "gitlab_service_github Resource - terraform-provider-gitlab"
subcategory: ""
description: |-
  NOTE: requires either EE (self-hosted) or Silver and above (GitLab.com).
  This resource manages a GitHub integration https://docs.gitlab.com/ee/user/project/integrations/github.html that updates pipeline statuses on a GitHub repo's pull requests.
---

# gitlab_service_github (Resource)

**NOTE**: requires either EE (self-hosted) or Silver and above (GitLab.com).

This resource manages a [GitHub integration](https://docs.gitlab.com/ee/user/project/integrations/github.html) that updates pipeline statuses on a GitHub repo's pull requests.

## Example Usage

```terraform
resource "gitlab_project" "awesome_project" {
  name             = "awesome_project"
  description      = "My awesome project."
  visibility_level = "public"
}

resource "gitlab_service_github" "github" {
  project        = gitlab_project.awesome_project.id
  token          = "REDACTED"
  repository_url = "https://github.com/gitlabhq/terraform-provider-gitlab"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **project** (String) ID of the project you want to activate integration on.
- **repository_url** (String) The URL of the GitHub repo to integrate with, e,g, https://github.com/gitlabhq/terraform-provider-gitlab.
- **token** (String, Sensitive) A GitHub personal access token with at least `repo:status` scope.

### Optional

- **id** (String) The ID of this resource.
- **static_context** (Boolean) Append instance name instead of branch to the status. Must enable to set a GitLab status check as _required_ in GitHub. See [Static / dynamic status check names] to learn more.

### Read-Only

- **active** (Boolean) Whether the integration is active.
- **created_at** (String) Create time.
- **title** (String) Title.
- **updated_at** (String) Update time.

## Import

Import is supported using the following syntax:

```shell
# You can import a service_github state using `terraform import <resource> <project_id>`:
terraform import gitlab_service_github.github 1
```
