---
subcategory: "Authentication"
---

# Resource: authenticationpolicylabel_authenticationpolicy_binding

The authenticationpolicylabel_authenticationpolicy_binding resource is used to bind authenticationpolicylabel to authenticationpolicy.


## Example usage

```hcl
resource "citrixadc_authenticationpolicylabel" "tf_authenticationpolicylabel" {
  labelname = "tf_authenticationpolicylabel"
  type      = "AAATM_REQ"
  comment   = "Testingresource"
}
resource "citrixadc_authenticationldapaction" "tf_authenticationldapaction" {
  name          = "ldapaction"
  serverip      = "1.2.3.4"
  serverport    = 8080
  authtimeout   = 1
  ldaploginname = "username"
}
resource "citrixadc_authenticationpolicy" "tf_authenticationpolicy" {
  name   = "tf_authenticationpolicy"
  rule   = "true"
  action = citrixadc_authenticationldapaction.tf_authenticationldapaction.name
}
resource "citrixadc_authenticationpolicylabel_authenticationpolicy_binding" "tf_bind" {
  labelname  = citrixadc_authenticationpolicylabel.tf_authenticationpolicylabel.labelname
  policyname = citrixadc_authenticationpolicy.tf_authenticationpolicy.name
  priority   = 20
}
```


## Argument Reference

* `labelname` - (Required) Name of the authentication policy label to which to bind the policy.
* `policyname` - (Required) Name of the authentication policy to bind to the policy label.
* `priority` - (Required) Specifies the priority of the policy.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `nextfactor` - (Optional) On success invoke label.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authenticationpolicylabel_authenticationpolicy_binding. It is the concatenation of the `labelname` and `policyname` attributes separated by a comma.


## Import

A authenticationpolicylabel_authenticationpolicy_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationpolicylabel_authenticationpolicy_binding.tf_bind tf_authenticationpolicylabel,tf_authenticationpolicy
```
