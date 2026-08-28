---
subcategory: "AAA"
---

# Resource: aaagroup_vpnsecureprivateaccessprofile_binding

The aaagroup_vpnsecureprivateaccessprofile_binding resource is used to bind a Secure Private Access Profile to an AAA group.


## Example usage

```hcl
resource "citrixadc_aaagroup" "tf_aaagroup" {
  groupname = "my_group"
  weight    = 100
}

resource "citrixadc_aaagroup_vpnsecureprivateaccessprofile_binding" "tf_binding" {
  groupname                  = citrixadc_aaagroup.tf_aaagroup.groupname
  secureprivateaccessprofile = "tf_spa_profile"
  gotopriorityexpression     = "END"
}
```


## Argument Reference

* `groupname` - (Required) Name of the group that you are binding. Minimum length =  1
* `secureprivateaccessprofile` - (Required) Name of the Secure Private Access Profile bound to the group.
* `gotopriorityexpression` - (Optional) Expression or other value specifying the next policy to evaluate if the current policy evaluates to TRUE. Specify one of the following values: NEXT, END, USE_INVOCATION_RESULT, or an expression that evaluates to a number.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the aaagroup_vpnsecureprivateaccessprofile_binding. It is the concatenation of `groupname` and `secureprivateaccessprofile` attributes separated by a comma.


## Import

A aaagroup_vpnsecureprivateaccessprofile_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_aaagroup_vpnsecureprivateaccessprofile_binding.tf_binding my_group,tf_spa_profile
```
