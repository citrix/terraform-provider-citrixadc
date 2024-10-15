---
subcategory: "System"
---

# Resource: systemgroup_systemcmdpolicy_binding

The systemgroup_systemcmdpolicy_binding resource is used to bind systemcmdpolicy to systemgroup.

~>  If you are using this resource to bind `systemcmdpolicy` to a `systemgroup`, do not define the `cmdpolicybinding` attribute in the systemgroup resource.


## Example usage

```hcl
resource "citrixadc_systemgroup" "tf_systemgroup" {
  groupname    = "tf_systemgroup"
  timeout      = 999
  promptstring = "bye>"
}

resource "citrixadc_systemcmdpolicy" "tf_policy" {
  policyname = "tf_policy"
  action     = "DENY"
  cmdspec    = "add.*"
}

resource "citrixadc_systemgroup_systemcmdpolicy_binding" "tf_bind" {
  groupname  = citrixadc_systemgroup.tf_systemgroup.groupname
  policyname = citrixadc_systemcmdpolicy.tf_policy.policyname
  priority   = 100
}
```


## Argument Reference

* `policyname` - (Required) The name of command policy.
* `priority` - (Required) The priority of the command policy.
* `groupname` - (Required) Name of the system group. Minimum length =  1


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the systemgroup_systemcmdpolicy_binding. It is the concatenation of the `groupname` and `policyname` attributes separated by a comma.


## Import

A systemgroup_systemcmdpolicy_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_systemgroup_systemcmdpolicy_binding.tf_bind tf_systemgroup,tf_policy
```
