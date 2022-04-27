---
subcategory: "Application Firewall"
---

# Resource: appfwprofile_cmdinjection_binding

The appfwprofile_cmdinjection_binding resource is used to bind cmdinjection to appfw profile.


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
resource "citrixadc_appfwprofile_cmdinjection_binding" "tf_binding" {
  name                 = citrixadc_appfwprofile.tf_appfwprofile.name
  cmdinjection         = "tf_cmdinjection"
  formactionurl_cmd    = "http://10.10.10.10/"
  as_scan_location_cmd = "HEADER"
  as_value_type_cmd    = "Keyword"
  as_value_expr_cmd    = "[a-z]+grep"
  alertonly            = "OFF"
  isvalueregex_cmd     = "REGEX"
  isautodeployed       = "NOTAUTODEPLOYED"
  comment              = "Testing"
}
```


## Argument Reference

* `name` - (Required) Name of the profile to which to bind an exemption or rule.
* `cmdinjection` - (Required) Name of the relaxed web form field/header/cookie
* `formactionurl_cmd` - (Required) The web form action URL.
* `alertonly` - (Optional) Send SNMP alert?
* `as_scan_location_cmd` - (Optional) Location of command injection exception - form field, header or cookie.
* `as_value_expr_cmd` - (Optional) The web form/header/cookie value expression.
* `as_value_type_cmd` - (Optional) Type of the relaxed web form value
* `comment` - (Optional) Any comments about the purpose of profile, or other useful information about the profile.
* `isautodeployed` - (Optional) Is the rule auto deployed by dynamic profile ?
* `isregex_cmd` - (Optional) Is the relaxed web form field name/header/cookie a regular expression?
* `isvalueregex_cmd` - (Optional) Is the web form field/header/cookie value a regular expression?
* `resourceid` - (Optional) A "id" that identifies the rule.
* `ruletype` - (Optional) Specifies rule type of binding
* `state` - (Optional) Enabled.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwprofile_cmdinjection_binding. It is the concatenation of `name` ,`cmdinjection` and `formactionurl_cmd` attributes separated by comma.


## Import

A appfwprofile_cmdinjection_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_appfwprofile_cmdinjection_binding.tf_binding tf_appfwprofile,tf_cmdinjection,http://10.10.10.10/
```
