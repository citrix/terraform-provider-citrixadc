---
subcategory: "Application Firewall"
---

# Resource: appfwprofile_jsonsqlurl_binding

The appfwprofile_jsonsqlurl_binding resource is used to bind jsonsqlurl to appfwprofile.


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
resource "citrixadc_appfwprofile_jsonsqlurl_binding" "tf_binding" {
  name           = citrixadc_appfwprofile.tf_appfwprofile.name
  jsonsqlurl     = "[abc][a-z]a*"
  keyname_json_sql = "id"
  as_value_type_json_sql = "SpecialString"
  as_value_expr_json_sql = "p"
  isautodeployed = "AUTODEPLOYED"
  state          = "ENABLED"
  alertonly      = "ON"
  comment        = "Testing"
}
resource "citrixadc_appfwprofile_jsonsqlurl_binding" "tf_binding2" {
  name           = citrixadc_appfwprofile.tf_appfwprofile.name
  jsonsqlurl     = "[abc][a-z]a*"
  isautodeployed = "AUTODEPLOYED"
  state          = "ENABLED"
  alertonly      = "ON"
  comment        = "Testing"
}
```


## Argument Reference

* `name` - (Required) Name of the profile to which to bind an exemption or rule.
* `jsonsqlurl` - (Required) A regular expression that designates a URL on the Json SQL URL list for which SQL violations are relaxed. Enclose URLs in double quotes to ensure preservation of any embedded spaces or non-alphanumeric characters.
* `alertonly` - (Optional) Send SNMP alert?
* `comment` - (Optional) Any comments about the purpose of profile, or other useful information about the profile.
* `isautodeployed` - (Optional) Is the rule auto deployed by dynamic profile ?
* `resourceid` - (Optional) A "id" that identifies the rule.
* `ruletype` - (Optional) Specifies rule type of binding
* `state` - (Optional) Enabled.
* `isvalueregex_json_sql` - (Optional) Is the JSON SQL key value a regular expression? Possible values = REGEX, NOTREGEX
* `iskeyregex_json_sql` - (Optional) Is the key name a regular expression? Possible values = REGEX, NOTREGEX
* `as_value_type_json_sql` - (Optional) Type of the relaxed JSON SQL key value. Possible values = Keyword, SpecialString, Wildchar.
* `as_value_expr_json_sql` - (Optional) The JSON SQL key value expression.
* `keyname_json_sql` - (Optional) An expression that designates a keyname on the JSON SQL URL for which SQL injection violations are relaxed.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwprofile_jsonsqlurl_binding. It is concatenation of `name`,`jsonsqlurl`,`keyname_json_sql`,`as_value_type_json_sql` and `as_value_expr_json_sql`attributes separated by comma.


## Import

An appfwprofile_jsonsqlurl_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_appfwprofile_jsonsqlurl_binding.tf_binding tf_appfwprofile,[abc][a-z]a*,id,SpecialString,p
```

An appfwprofile_jsonsqlurl_binding which does not have values set for keyname_json_sql, as_value_type_json_sql and as_value_expr_json_sql can be imported using its id, e.g.

```shell
terraform import citrixadc_appfwprofile_jsonsqlurl_binding.tf_binding2 tf_appfwprofile,[abc][a-z]a*
```
