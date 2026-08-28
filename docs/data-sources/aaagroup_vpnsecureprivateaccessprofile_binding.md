---
subcategory: "AAA"
---

# Data Source: aaagroup_vpnsecureprivateaccessprofile_binding

The aaagroup_vpnsecureprivateaccessprofile_binding data source allows you to retrieve information about a Secure Private Access Profile binding to an AAA group.

## Example Usage

```terraform
data "citrixadc_aaagroup_vpnsecureprivateaccessprofile_binding" "tf_binding" {
  groupname                  = "my_group"
  secureprivateaccessprofile = "tf_spa_profile"
}

output "gotopriorityexpression" {
  value = data.citrixadc_aaagroup_vpnsecureprivateaccessprofile_binding.tf_binding.gotopriorityexpression
}
```

## Argument Reference

* `groupname` - (Required) Name of the group that you are binding.
* `secureprivateaccessprofile` - (Required) Name of the Secure Private Access Profile bound to the group.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `gotopriorityexpression` - Expression or other value specifying the next policy to evaluate if the current policy evaluates to TRUE.
* `id` - The id of the aaagroup_vpnsecureprivateaccessprofile_binding. It is the concatenation of `groupname` and `secureprivateaccessprofile` attributes separated by a comma.
