---
subcategory: "Application Firewall"
---

# Data Source: appfwprofile_jsoncmdurl_binding

The appfwprofile_jsoncmdurl_binding data source allows you to retrieve information about a specific json command URL binding to an application firewall profile.


## Example Usage

```terraform
data "citrixadc_appfwprofile_jsoncmdurl_binding" "tf_binding" {
  name                   = "tf_appfwprofile"
  jsoncmdurl             = "^https://sd2\\-zgw\\.test\\.ctxns\\.com/api/v1/resource/temp$"
  keyname_json_cmd       = "id"
  as_value_type_json_cmd = "SpecialString"
  as_value_expr_json_cmd = "$"
}

output "alertonly" {
  value = data.citrixadc_appfwprofile_jsoncmdurl_binding.tf_binding.alertonly
}

output "comment" {
  value = data.citrixadc_appfwprofile_jsoncmdurl_binding.tf_binding.comment
}

output "state" {
  value = data.citrixadc_appfwprofile_jsoncmdurl_binding.tf_binding.state
}
```


## Argument Reference

* `name` - (Required) Name of the profile to which to bind an exemption or rule.
* `jsoncmdurl` - (Required) A regular expression that designates a URL on the Json CMD URL list for which Command injection violations are relaxed.
* `keyname_json_cmd` - (Optional) An expression that designates a keyname on the JSON CMD URL for which Command injection violations are relaxed.
* `as_value_type_json_cmd` - (Optional) Type of the relaxed JSON CMD key value.
* `as_value_expr_json_cmd` - (Optional) The JSON CMD key value expression.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `alertonly` - Send SNMP alert?
* `comment` - Any comments about the purpose of profile, or other useful information about the profile.
* `id` - The id of the appfwprofile_jsoncmdurl_binding. It is a system-generated identifier.
* `isautodeployed` - Is the rule auto deployed by dynamic profile ?
* `iskeyregex_json_cmd` - Is the key name a regular expression?
* `isvalueregex_json_cmd` - Is the JSON CMD key value a regular expression?
* `resourceid` - A "id" that identifies the rule.
* `state` - Enabled.
