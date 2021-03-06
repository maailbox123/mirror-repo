---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "gitlab_project_custom_attribute Resource - terraform-provider-gitlab"
subcategory: ""
description: |-
  This resource allows you to set custom attributes for a project.
---

# gitlab_project_custom_attribute (Resource)

This resource allows you to set custom attributes for a project.

## Example Usage

```terraform
resource "gitlab_project_custom_attribute" "attr" {
  project = "42"
  key     = "location"
  value   = "Greenland"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **key** (String) Key for the Custom Attribute.
- **project** (Number) The id of the project.
- **value** (String) Value for the Custom Attribute.

### Optional

- **id** (String) The ID of this resource.

## Import

Import is supported using the following syntax:

```shell
# You can import a project custom attribute using an id made up of `{project-id}:{key}`, e.g.
terraform import gitlab_project_custom_attribute.attr 42:location
```
