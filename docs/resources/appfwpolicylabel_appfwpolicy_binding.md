---
subcategory: "Application Firewall"
---

# Resource: appfwpolicylabel_appfwpolicy_binding

The appfwpolicylabel_appfwpolicy_binding resource is used to bind appfw policy to appfw policylabel resource.


## Example usage

```hclresource "citrixadc_appfwpolicylabel" "tf_appfwpolicylabel" {
  labelname       = "tf_appfwpolicylabel"
  policylabeltype = "http_req"
}
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
resource "citrixadc_appfwpolicylabel_appfwpolicy_binding" "tf_binding" {
  labelname  = citrixadc_appfwpolicylabel.tf_appfwpolicylabel.labelname
  policyname = citrixadc_appfwpolicy.tf_appfwpolicy.name
  priority   = 90
}
```


## Argument Reference

* `labelname` - (Required) Name of the application firewall policy label.
* `policyname` - (Required) Name of the application firewall policy to bind to the policy label.
* `priority` - (Required) Positive integer specifying the priority of the policy. A lower number specifies a higher priority. Must be unique within a group of policies that are bound to the same bind point or label. Policies are evaluated in the order of their priority numbers.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - (Optional) If the current policy evaluates to TRUE, terminate evaluation of policies bound to the current policy label, and then forward the request to the specified virtual server or evaluate the specified policy label.
* `invoke_labelname` - (Optional) Name of the policy label to invoke if the current policy evaluates to TRUE, the invoke parameter is set, and Label Type is set to Policy Label.
* `labeltype` - (Optional) Type of policy label to invoke if the current policy evaluates to TRUE and the invoke parameter is set. Available settings function as follows: * reqvserver. Invoke the unnamed policy label associated with the specified request virtual server. * policylabel. Invoke the specified user-defined policy label.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwpolicylabel_appfwpolicy_binding. It is the concatenation of `labelname`and `policyname` attributes separated by comma.


## Import

A appfwpolicylabel_appfwpolicy_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_appfwpolicylabel_appfwpolicy_binding.tf_binding tf_appfwpolicylabel,tf_appfwpolicy
```
