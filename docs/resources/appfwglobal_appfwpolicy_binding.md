---
subcategory: "Application Firewall"
---

# Resource: appfwglobal_appfwpolicy_binding

The appfwglobal_appfwpolicy_binding resource is used to bind appfwpolicy to appfwglobal configuration.


## Example usage

```hcl
resource "citrixadc_appfwprofile" "tf_appfwprofile" {
  name                     = "tf_appfwprofile"
  bufferoverflowaction     = ["none"]
  contenttypeaction        = ["none"]
  cookieconsistencyaction  = ["none"]
  creditcard               = ["none"]
  creditcardaction         = ["none"]
  crosssitescriptingaction = ["none"]
  csrftagaction            = ["none"]
  denyurlaction            = ["none"]
  dynamiclearning          = ["none"]
  fieldconsistencyaction   = ["none"]
  fieldformataction        = ["none"]
  fileuploadtypesaction    = ["none"]
  inspectcontenttypes      = ["none"]
  jsondosaction            = ["none"]
  jsonsqlinjectionaction   = ["none"]
  jsonxssaction            = ["none"]
  multipleheaderaction     = ["none"]
  sqlinjectionaction       = ["none"]
  starturlaction           = ["none"]
  type                     = ["HTML"]
  xmlattachmentaction      = ["none"]
  xmldosaction             = ["none"]
  xmlformataction          = ["none"]
  xmlsoapfaultaction       = ["none"]
  xmlsqlinjectionaction    = ["none"]
  xmlvalidationaction      = ["none"]
  xmlwsiaction             = ["none"]
  xmlxssaction             = ["none"]
}
resource "citrixadc_appfwpolicy" "tf_appfwpolicy" {
  name        = "tf_appfwpolicy"
  profilename = citrixadc_appfwprofile.tf_appfwprofile.name
  rule        = "true"
}
resource "citrixadc_appfwglobal_appfwpolicy_binding" "tf_binding" {
  policyname = citrixadc_appfwpolicy.tf_appfwpolicy.name
  priority   = 30
  state      = "ENABLED"
  type       = "REQ_DEFAULT"
}
```


## Argument Reference

* `policyname` - (Required) Name of the policy.
* `globalbindtype` - (Optional) 0
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - (Optional) If the current policy evaluates to TRUE, terminate evaluation of policies bound to the current policy label, and then forward the request to the specified virtual server or evaluate the specified policy label.
* `labelname` - (Optional) Name of the policy label to invoke if the current policy evaluates to TRUE, the invoke parameter is set, and Label Type is set to Policy Label.
* `labeltype` - (Optional) Type of policy label invocation.
* `priority` - (Optional) The priority of the policy.
* `state` - (Optional) Enable or disable the binding to activate or deactivate the policy. This is applicable to classic policies only.
* `type` - (Optional) Bind point to which to policy is bound.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwglobal_appfwpolicy_binding. It has the same value as the `policyname` attribute.


## Import

A appfwglobal_appfwpolicy_binding can be imported using its policyname, e.g.

```shell
terraform import citrixadc_appfwglobal_appfwpolicy_binding.tf_binding tf_appfwpolicy
```
