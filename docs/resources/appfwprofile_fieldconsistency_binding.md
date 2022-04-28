---
subcategory: "Application Firewall"
---

# Resource: appfwprofile_fieldconsistency_binding

The appfwprofile_fieldconsistency_binding resource is used to bind fieldconsistency to appfwprofile.


## Example usage

```hclresource "citrixadc_appfwprofile" "tf_appfwprofile" {
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
resource "citrixadc_appfwprofile_fieldconsistency_binding" "tf_binding" {
  name              = citrixadc_appfwprofile.tf_appfwprofile.name
  fieldconsistency  = "tf_field"
  formactionurl_ffc = "www.example.com"
  isautodeployed    = "NOTAUTODEPLOYED"
  state             = "DISABLED"
  alertonly         = "OFF"
  isregex_ffc       = "REGEX"
  comment           = "Testing"
}
```


## Argument Reference

* `name` - (Required) Name of the profile to which to bind an exemption or rule.
* `fieldconsistency` - (Required) The web form field name.
* `formactionurl_ffc` - (Required) The web form action URL.
* `alertonly` - (Optional) Send SNMP alert?
* `comment` - (Optional) Any comments about the purpose of profile, or other useful information about the profile.
* `isautodeployed` - (Optional) Is the rule auto deployed by dynamic profile ?
* `isregex_ffc` - (Optional) Is the web form field name a regular expression?
* `resourceid` - (Optional) A "id" that identifies the rule.
* `ruletype` - (Optional) Specifies rule type of binding
* `state` - (Optional) Enabled.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwprofile_fieldconsistency_binding. It is the concatenation of `name` , `fieldconsistency` and `formactionurl_ffc` attributes separated by comma.


## Import

A appfwprofile_fieldconsistency_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_appfwprofile_fieldconsistency_binding.tf_binding tf_appfwprofile,tf_field,www.example.com
```
