---
subcategory: "Application Firewall"
---

# Resource: appfwprofile_csrftag_binding

The appfwprofile_csrftag_binding resource is used to bind csrftag to appfwprofile.


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
resource "citrixadc_appfwprofile_csrftag_binding" "tf_binding" {
  name              = citrixadc_appfwprofile.tf_appfwprofile.name
  csrftag           = "www.source.com"
  csrfformactionurl = "www.action.com"
  isautodeployed    = "NOTAUTODEPLOYED"
  comment           = "Testing"
  state             = "ENABLED"
  alertonly         = "OFF"
}
```


## Argument Reference

* `name` - (Required) Name of the profile to which to bind an exemption or rule.
* `csrftag` - (Required) The web form originating URL.
* `csrfformactionurl` - (Required) The web form action URL.
* `alertonly` - (Optional) Send SNMP alert?. Possible values: [ ON, OFF ]
* `comment` - (Optional) Any comments about the purpose of profile, or other useful information about the profile.
* `isautodeployed` - (Optional) Is the rule auto deployed by dynamic profile ?. Possible values: [ AUTODEPLOYED, NOTAUTODEPLOYED ]
* `resourceid` - (Optional) A unique id that identifies the rule.
* `ruletype` - (Optional) Specifies rule type of binding.
* `state` - (Optional) Enabled. Possible values: [ ENABLED, DISABLED ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwprofile_csrftag_binding. It is the concatenation of the `name`, `csrftag` and `csrfformactionurl` attributes separated by a comma.


## Import

A appfwprofile_csrftag_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_appfwprofile_csrftag_binding.tf_binding tf_appfwprofile,www.source.com,www.action.com
```
