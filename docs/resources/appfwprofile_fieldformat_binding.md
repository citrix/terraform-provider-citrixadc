---
subcategory: "Application Firewall"
---

# Resource: appfwprofile_fieldformat_binding

The appfwprofile_fieldformat_binding resource is used to bind fieldformat to appfwprofile.


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
resource "citrixadc_appfwprofile_fieldformat_binding" "tf_binding" {
  name                 = citrixadc_appfwprofile.tf_appfwprofile.name
  fieldformat          = "tf_field"
  formactionurl_ff     = "http://www.example.com"
  comment              = "Testing"
  state                = "ENABLED"
  fieldformatmaxlength = 20
  isregexff            = "NOTREGEX"
  fieldtype            = "alpha"
  alertonly            = "OFF"
}
```


## Argument Reference

* `name` - (Required) Name of the profile to which to bind an exemption or rule.
* `fieldformat` - (Required) Name of the form field to which a field format will be assigned.
* `formactionurl_ff` - (Required) Action URL of the form field to which a field format will be assigned.
* `alertonly` - (Optional) Send SNMP alert?
* `comment` - (Optional) Any comments about the purpose of profile, or other useful information about the profile.
* `fieldformatmaxlength` - (Optional) The maximum allowed length for data in this form field.
* `fieldformatminlength` - (Optional) The minimum allowed length for data in this form field.
* `fieldtype` - (Optional) The field type you are assigning to this form field.
* `isautodeployed` - (Optional) Is the rule auto deployed by dynamic profile ?
* `isregex_ff` - (Optional) Is the form field name a regular expression?
* `resourceid` - (Optional) A "id" that identifies the rule.
* `ruletype` - (Optional) Specifies rule type of binding
* `state` - (Optional) Enabled.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwprofile_fieldformat_binding. It is the concatenation of `name`,`fieldformat` and `formactionurl_ff` attributes separated by comma.


## Import

A appfwprofile_fieldformat_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_appfwprofile_fieldformat_binding.tf_binding tf_appfwprofile,tf_field,http://www.example.com
```
