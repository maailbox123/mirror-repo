---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "gitlab_project_freeze_period Resource - terraform-provider-gitlab"
subcategory: ""
description: |-
  This resource allows you to create and manage freeze periods. For further information on freeze periods, consult the gitlab documentation https://docs.gitlab.com/ee/api/freeze_periods.html#create-a-freeze-period.
---

# gitlab_project_freeze_period (Resource)

This resource allows you to create and manage freeze periods. For further information on freeze periods, consult the [gitlab documentation](https://docs.gitlab.com/ee/api/freeze_periods.html#create-a-freeze-period).

## Example Usage

```terraform
resource "gitlab_project_freeze_period" "schedule" {
  project_id    = gitlab_project.foo.id
  freeze_start  = "0 23 * * 5"
  freeze_end    = "0 7 * * 1"
  cron_timezone = "UTC"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **freeze_end** (String) End of the Freeze Period in cron format (e.g. `0 2 * * *`).
- **freeze_start** (String) Start of the Freeze Period in cron format (e.g. `0 1 * * *`).
- **project_id** (String) The id of the project to add the schedule to.

### Optional

- **cron_timezone** (String) The timezone.
- **id** (String) The ID of this resource.

## Import

Import is supported using the following syntax:

```shell
# GitLab project freeze periods can be imported using an id made up of `project_id:freeze_period_id`, e.g.
terraform import gitlab_project_freeze_period.schedule "12345:1337"
```
