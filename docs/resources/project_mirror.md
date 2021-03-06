---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "gitlab_project_mirror Resource - terraform-provider-gitlab"
subcategory: ""
description: |-
  This resource allows you to add a mirror target for the repository, all changes will be synced to the remote target.
  -> This is for pushing changes to a remote repository. Pull Mirroring can be configured using a combination of the
  import_url, mirror, and mirror_trigger_builds properties on the gitlab_project resource.
  For further information on mirroring, consult the
  gitlab documentation https://docs.gitlab.com/ee/user/project/repository/repository_mirroring.html#repository-mirroring.
---

# gitlab_project_mirror (Resource)

This resource allows you to add a mirror target for the repository, all changes will be synced to the remote target.

-> This is for *pushing* changes to a remote repository. *Pull Mirroring* can be configured using a combination of the
`import_url`, `mirror`, and `mirror_trigger_builds` properties on the `gitlab_project` resource.

For further information on mirroring, consult the
[gitlab documentation](https://docs.gitlab.com/ee/user/project/repository/repository_mirroring.html#repository-mirroring).

## Example Usage

```terraform
resource "gitlab_project_mirror" "foo" {
  project = "1"
  url     = "https://username:password@github.com/org/repository.git"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **project** (String) The id of the project.
- **url** (String, Sensitive) The URL of the remote repository to be mirrored.

### Optional

- **enabled** (Boolean) Determines if the mirror is enabled.
- **id** (String) The ID of this resource.
- **keep_divergent_refs** (Boolean) Determines if divergent refs are skipped.
- **only_protected_branches** (Boolean) Determines if only protected branches are mirrored.

### Read-Only

- **mirror_id** (Number) Mirror ID.

## Import

Import is supported using the following syntax:

```shell
# GitLab project mirror can be imported using an id made up of `project_id:mirror_id`, e.g.
terraform import gitlab_project_mirror.foo "12345:1337"
```
