---
subcategory: "Application Firewall"
---

# Resource: appfwprofile_safeobject_binding

The appfwprofile_safeobject_binding resource is used to bind safeobject to appfw profile. 


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
resource "citrixadc_appfwprofile_safeobject_binding" "tf_binding" {
  name           = citrixadc_appfwprofile.tf_appfwprofile.name
  safeobject     = "tf_safeobject"
  as_expression  = "regularexpression"
  maxmatchlength = 10
  state          = "DISABLED"
  alertonly      = "OFF"
  isautodeployed = "AUTODEPLOYED"
  comment        = "Example"
  action         = ["block", "log"]
}
```


## Argument Reference

* `name` - (Required) Name of the profile to which to bind an exemption or rule.
* `safeobject` - (Required) Name of the Safe Object.
* `action` - (Optional) Safe Object action types. (BLOCK | LOG | STATS | NONE)
* `alertonly` - (Optional) Send SNMP alert?
* `as_expression` - (Optional) A regular expression that defines the Safe Object.
* `comment` - (Optional) Any comments about the purpose of profile, or other useful information about the profile.
* `isautodeployed` - (Optional) Is the rule auto deployed by dynamic profile ?
* `maxmatchlength` - (Optional) Maximum match length for a Safe Object expression.
* `resourceid` - (Optional) A "id" that identifies the rule.
* `ruletype` - (Optional) Specifies rule type of binding
* `state` - (Optional) Enabled.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwprofile_safeobject_binding. It is theconcatenation of `name` and `safeobject` attributes separated by comma.


## Import

A appfwprofile_safeobject_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_appfwprofile_safeobject_binding.tf_binding tf_appfwprofile,tf_safeobject
```
