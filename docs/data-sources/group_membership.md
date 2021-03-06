---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "gitlab_group_membership Data Source - terraform-provider-gitlab"
subcategory: ""
description: |-
  Provide details about a list of group members in the gitlab provider. The results include id, username, name and more about the requested members.
  Note: exactly one of groupid or fullpath must be provided.
---

# gitlab_group_membership (Data Source)

Provide details about a list of group members in the gitlab provider. The results include id, username, name and more about the requested members.

> **Note**: exactly one of group_id or full_path must be provided.

## Example Usage

```terraform
# By group's ID
data "gitlab_group_membership" "example" {
  group_id = 123
}

# By group's full path
data "gitlab_group_membership" "example" {
  full_path = "foo/bar"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **access_level** (String) Only return members with the desired access level. Acceptable values are: `guest`, `reporter`, `developer`, `maintainer`, `owner`.
- **full_path** (String) The full path of the group.
- **group_id** (Number) The ID of the group.
- **id** (String) The ID of this resource.

### Read-Only

- **members** (List of Object) The list of group members. (see [below for nested schema](#nestedatt--members))

<a id="nestedatt--members"></a>
### Nested Schema for `members`

Read-Only:

- **access_level** (String)
- **avatar_url** (String)
- **expires_at** (String)
- **id** (Number)
- **name** (String)
- **state** (String)
- **username** (String)
- **web_url** (String)


