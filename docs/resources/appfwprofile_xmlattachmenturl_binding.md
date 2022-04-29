---
subcategory: "Application Firewall"
---

# Resource: appfwprofile_xmlattachmenturl_binding

The appfwprofile_xmlattachmenturl_binding resource is used to bind xmlattachmenturl to appfwprofile.


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
resource "citrixadc_appfwprofile_xmlattachmenturl_binding" "tf_binding" {
  name                          = citrixadc_appfwprofile.tf_appfwprofile.name
  xmlattachmenturl              = ".*"
  xmlattachmentcontenttype      = "abc*"
  alertonly                     = "ON"
  state                         = "ENABLED"
  isautodeployed                = "AUTODEPLOYED"
  comment                       = "Testing"
  xmlattachmentcontenttypecheck = "ON"
  xmlmaxattachmentsize          = "1000"
  xmlmaxattachmentsizecheck     = "ON"
}
```


## Argument Reference

* `name` - (Required) Name of the profile to which to bind an exemption or rule.
* `xmlattachmenturl` - (Required) XML attachment URL regular expression length.
* `alertonly` - (Optional) Send SNMP alert?
* `comment` - (Optional) Any comments about the purpose of profile, or other useful information about the profile.
* `isautodeployed` - (Optional) Is the rule auto deployed by dynamic profile ?
* `resourceid` - (Optional) A "id" that identifies the rule.
* `ruletype` - (Optional) Specifies rule type of binding
* `state` - (Optional) Enabled.
* `xmlattachmentcontenttype` - (Optional) Specify content-type regular expression.
* `xmlattachmentcontenttypecheck` - (Optional) State if XML attachment content-type check is ON or OFF. Protects against XML requests with illegal attachments.
* `xmlmaxattachmentsize` - (Optional) Specify maximum attachment size.
* `xmlmaxattachmentsizecheck` - (Optional) State if XML Max attachment size Check is ON or OFF. Protects against XML requests with large attachment data.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwprofile_xmlattachmenturl_binding. It is the concatenation of `name`  and `xmlattachmenturl` attributes separated by comma.


## Import

A appfwprofile_xmlattachmenturl_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_appfwprofile_xmlattachmenturl_binding.tf_binding tf_appfwprofile,.*
```
