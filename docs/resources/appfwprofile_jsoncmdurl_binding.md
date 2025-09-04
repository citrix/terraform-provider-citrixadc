---
subcategory: "Application Firewall"
---

# Resource: appfwprofile_jsoncmdurl_binding

The appfwprofile_jsoncmdurl_binding resource is used to bind jsoncmdurl to appfw profile resource.


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
resource "citrixadc_appfwprofile_jsoncmdurl_binding" "tf_binding" {
  name           = citrixadc_appfwprofile.tf_appfwprofile.name
  jsoncmdurl     = "www.example.com"
  alertonly      = "ON"
  isautodeployed = "AUTODEPLOYED"
  comment        = "Testing"
  state          = "DISABLED"
}
```


## Argument Reference

* `name` - (Required) Name of the profile to which to bind an exemption or rule.
* `jsoncmdurl` - (Required) A regular expression that designates a URL on the Json CMD URL list for which Command injection violations are relaxed. Enclose URLs in double quotes to ensure preservation of any embedded spaces or non-alphanumeric characters.
* `alertonly` - (Optional) Send SNMP alert?
* `comment` - (Optional) Any comments about the purpose of profile, or other useful information about the profile.
* `isautodeployed` - (Optional) Is the rule auto deployed by dynamic profile ?
* `resourceid` - (Optional) A "id" that identifies the rule.
* `ruletype` - (Optional) Specifies rule type of binding
* `state` - (Optional) Enabled.
* `iskeyregex_json_cmd` - (Optional) Is the key name a regular expression?
* `keyname_json_cmd` - (Optional) An expression that designates a keyname on the JSON CMD URL for which Command injection violations are relaxed.
* `as_value_type_json_cmd` - (Optional) Type of the relaxed JSON CMD key value.
* `as_value_expr_json_cmd` - (Optional) The JSON CMD key value expression.
* `isvalueregex_json_cmd` - (Optional) Is the JSON CMD key value a regular expression?


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwprofile_jsoncmdurl_binding. It is the concatenation of `name` and `jsoncmdurl` attributes separated by comma.


## Import

A appfwprofile_jsoncmdurl_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_appfwprofile_jsoncmdurl_binding.tf_binding tf_appfwprofile,www.example.com
```
