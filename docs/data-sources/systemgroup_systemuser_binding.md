---
subcategory: "System"
---

# Data Source: systemgroup_systemuser_binding

The systemgroup_systemuser_binding data source allows you to retrieve information about a binding between a system group and a system user.

## Example Usage

```terraform
data "citrixadc_systemgroup_systemuser_binding" "tf_bind" {
  groupname = "tf_systemgroup"
  username  = "tf_user"
}

output "groupname" {
  value = data.citrixadc_systemgroup_systemuser_binding.tf_bind.groupname
}

output "username" {
  value = data.citrixadc_systemgroup_systemuser_binding.tf_bind.username
}
```

## Argument Reference

* `groupname` - (Required) Name of the system group.
* `username` - (Required) The system user.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the systemgroup_systemuser_binding. It is a system-generated identifier.
