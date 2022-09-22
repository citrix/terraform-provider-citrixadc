---
subcategory: "Application Firewall"
---

# Resource: appfwprofile_creditcardnumber_binding

The appfwprofile_creditcardnumber_binding resource is used to bind creditcard number to appfw profile.


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
resource "citrixadc_appfwprofile_creditcardnumber_binding" "tf_binding" {
  name                = citrixadc_appfwprofile.tf_appfwprofile.name
  creditcardnumberurl = "www.example.com"
  creditcardnumber    = "123456789"
  isautodeployed      = "AUTODEPLOYED"
  alertonly           = "ON"
  state               = "ENABLED"
  comment             = "Testing"
}
```


## Argument Reference

* `name` - (Required) Name of the profile to which to bind an exemption or rule.
* `creditcardnumber` - (Required) The object expression that is to be excluded from safe commerce check
* `creditcardnumberurl` - (Required) The url for which the list of credit card numbers are needed to be bypassed from inspection
* `alertonly` - (Optional) Send SNMP alert?
* `comment` - (Optional) Any comments about the purpose of profile, or other useful information about the profile.
* `isautodeployed` - (Optional) Is the rule auto deployed by dynamic profile ?
* `resourceid` - (Optional) A "id" that identifies the rule.
* `ruletype` - (Optional) Specifies rule type of binding
* `state` - (Optional) Enabled.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwprofile_creditcardnumber_binding. It is the concatenation of `name` , `creditcardnumber` and `creditcardnumberurl` attributes separated by comma.


## Import

A appfwprofile_creditcardnumber_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_appfwprofile_creditcardnumber_binding.tf_binding tf_csaction
```
