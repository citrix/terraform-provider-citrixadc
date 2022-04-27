---
subcategory: "Application Firewall"
---

# Resource: appfwprofile_excluderescontenttype_binding

The appfwprofile_excluderescontenttype_binding resource is used to bind excluderescontenttype to appfw profile.


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
resource "citrixadc_appfwprofile_excluderescontenttype_binding" "tf_binding" {
  name                  = citrixadc_appfwprofile.tf_appfwprofile.name
  excluderescontenttype = "expressionexample"
  state                 = "DISABLED"
  isautodeployed        = "AUTODEPLOYED"
  alertonly             = "ON"
  comment               = "Testing"
}
```


## Argument Reference

* `name` - (Required) Name of the profile to which to bind an exemption or rule.
* `excluderescontenttype` - (Required) A regular expression that represents the content type of the response that are to be excluded from inspection.
* `alertonly` - (Optional) Send SNMP alert?
* `comment` - (Optional) Any comments about the purpose of profile, or other useful information about the profile.
* `isautodeployed` - (Optional) Is the rule auto deployed by dynamic profile ?
* `resourceid` - (Optional) A "id" that identifies the rule.
* `ruletype` - (Optional) Specifies rule type of binding
* `state` - (Optional) Enabled.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwprofile_excluderescontenttype_binding. It is the concatenation of `name` and `excluderescontenttype` attributes separated by comma.


## Import

A appfwprofile_excluderescontenttype_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_appfwprofile_excluderescontenttype_binding.tf_binding tf_appfwprofile,expressionexample
```
