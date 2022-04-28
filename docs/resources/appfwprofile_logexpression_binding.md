---
subcategory: "Application Firewall"
---

# Resource: appfwprofile_logexpression_binding

The appfwprofile_logexpression_binding resource is used to bind logexpression to appfwprofile resource.


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
resource "citrixadc_appfwprofile_logexpression_binding" "tf_binding" {
  name             = citrixadc_appfwprofile.tf_appfwprofile.name
  logexpression    = "tf_logexp"
  as_logexpression = "HTTP.REQ.IS_VALID"
  alertonly        = "ON"
  isautodeployed   = "AUTODEPLOYED"
  comment          = "Testing"
  state            = "ENABLED"
}
```


## Argument Reference

* `name` - (Required) Name of the profile to which to bind an exemption or rule.
* `logexpression` - (Required) Name of LogExpression object.
* `alertonly` - (Optional) Send SNMP alert?
* `as_logexpression` - (Optional) LogExpression to log when violation happened on appfw profile
* `comment` - (Optional) Any comments about the purpose of profile, or other useful information about the profile.
* `isautodeployed` - (Optional) Is the rule auto deployed by dynamic profile ?
* `resourceid` - (Optional) A "id" that identifies the rule.
* `ruletype` - (Optional) Specifies rule type of binding
* `state` - (Optional) Enabled.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwprofile_logexpression_binding. It is the concatenation of `name` and `logexpression` attributes separated by comma.


## Import

A appfwprofile_logexpression_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_appfwprofile_logexpression_binding.tf_binding tf_appfwprofile,tf_logexp
```
