---
subcategory: "AAA"
---

# Data Source `aaagroup_aaauser_binding`

The aaagroup_aaauser_binding data source allows you to retrieve information about bindings between AAA groups and AAA users.


## Example usage

```terraform
data "citrixadc_aaagroup_aaauser_binding" "tf_binding" {
  groupname = "test_group"
  username  = "test_user"
}

output "gotopriorityexpression" {
  value = data.citrixadc_aaagroup_aaauser_binding.tf_binding.gotopriorityexpression
}
```


## Argument Reference

* `groupname` - (Required) Name of the group.
* `username` - (Required) Name of the user to bind to the group.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `gotopriorityexpression` - Expression or other value specifying the priority of the next policy to be evaluated.

## Attribute Reference

* `id` - The id of the aaagroup_aaauser_binding. It is a combination of the `groupname` and `username` attributes.


## Import

A aaagroup_aaauser_binding can be imported using its groupname and username, e.g.

```shell
terraform import citrixadc_aaagroup_aaauser_binding.tf_binding test_group,test_user
```
