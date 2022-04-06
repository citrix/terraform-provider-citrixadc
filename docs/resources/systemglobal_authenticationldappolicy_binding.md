---
subcategory: "System"
---

# Resource: systemglobal_authenticationldappolicy_binding

The systemglobal_authenticationldappolicy_binding resource is used to bind authenticationldappolicy to systemglobal.


## Example usage

```hcl
resource "citrixadc_authenticationldapaction" "tf_authenticationldapaction" {
  name          = "tf_ldapaction"
  serverip      = "1.2.3.4"
  serverport    = 8080
  authtimeout   = 1
  ldaploginname = "username"
}
resource "citrixadc_authenticationldappolicy" "tf_authenticationldappolicy" {
  name      = "tf_authenticationldappolicy"
  rule      = "NS_TRUE"
  reqaction = citrixadc_authenticationldapaction.tf_authenticationldapaction.name
}
resource "citrixadc_systemglobal_authenticationldappolicy_binding" "tf_bind" {
  policyname     = citrixadc_authenticationldappolicy.tf_authenticationldappolicy.name
  globalbindtype = "RNAT_GLOBAL"
  priority       = 88
  feature        = "SYSTEM"
}
```


## Argument Reference

* `policyname` - (Required) The name of the  command policy.
* `builtin` - (Optional) Indicates that a variable is a built-in (SYSTEM INTERNAL) type.
* `feature` - (Optional) The feature to be checked while applying this config
* `globalbindtype` - (Optional) 0
* `gotopriorityexpression` - (Optional) Applicable only to advance authentication policy. Expression or other value specifying the next policy to be evaluated if the current policy evaluates to TRUE.  Specify one of the following values: * NEXT - Evaluate the policy with the next higher priority number. * END - End policy evaluation.
* `nextfactor` - (Optional) On success invoke label. Applicable for advanced authentication policy binding
* `priority` - (Optional) The priority of the command policy.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the systemglobal_authenticationldappolicy_binding. It has the same value as the `policyname` attribute.


## Import

A systemglobal_authenticationldappolicy_binding can be imported using its policyname, e.g.

```shell
terraform import citrixadc_systemglobal_authenticationldappolicy_binding.tf_bind tf_authenticationldappolicy
```
