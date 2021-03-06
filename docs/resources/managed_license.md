---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "gitlab_managed_license Resource - terraform-provider-gitlab"
subcategory: ""
description: |-
  This resource allows you to add rules for managing licenses on a project.
  For additional information, please see the gitlab documentation https://docs.gitlab.com/ee/user/compliance/license_compliance/.
  ~> Using this resource requires an active gitlab ultimate https://about.gitlab.com/pricing/subscription.
---

# gitlab_managed_license (Resource)

This resource allows you to add rules for managing licenses on a project.
For additional information, please see the [gitlab documentation](https://docs.gitlab.com/ee/user/compliance/license_compliance/).

~> Using this resource requires an active [gitlab ultimate](https://about.gitlab.com/pricing/)subscription.

## Example Usage

```terraform
resource "gitlab_project" "foo" {
  name             = "example project"
  description      = "Lorem Ipsum"
  visibility_level = "public"
}

resource "gitlab_managed_license" "mit" {
  project         = gitlab_project.foo.id
  name            = "MIT license"
  approval_status = "approved"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **approval_status** (String) Whether the license is approved or not. Only 'approved' or 'blacklisted' allowed.
- **name** (String) The name of the managed license (I.e., 'Apache License 2.0' or 'MIT license')
- **project** (String) The ID of the project under which the managed license will be created.

### Optional

- **id** (String) The ID of this resource.

## Import

Import is supported using the following syntax:

```shell
# You can import this resource with an id made up of `{project-id}:{license-id}`, e.g.
terraform import gitlab_managed_license.foo 1:2
```
