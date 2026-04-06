---
subcategory: "Application Firewall"
---

# Data Source: appfwprofile_jsonxssurl_binding

The appfwprofile_jsonxssurl_binding data source allows you to retrieve information about appfwprofile jsonxssurl bindings.

## Example Usage

```terraform
data "citrixadc_appfwprofile_jsonxssurl_binding" "tf_binding" {
  name       = "tf_appfwprofile"
  jsonxssurl = "^https://sd2\\-zgw\\.test\\.ctxns\\.com/api/document/content$"
}

output "state" {
  value = data.citrixadc_appfwprofile_jsonxssurl_binding.tf_binding.state
}

output "comment" {
  value = data.citrixadc_appfwprofile_jsonxssurl_binding.tf_binding.comment
}
```

## Argument Reference

* `name` - (Required) Name of the profile to which to bind an exemption or rule.
* `jsonxssurl` - (Required) A regular expression that designates a URL on the Json XSS URL list for which XSS violations are relaxed. Enclose URLs in double quotes to ensure preservation of any embedded spaces or non-alphanumeric characters.
* `keyname_json_xss` - (Optional) An expression that designates a keyname on the JSON XSS URL for which XSS injection violations are relaxed.
* `as_value_type_json_xss` - (Optional) Type of the relaxed JSON XSS key value.
* `as_value_expr_json_xss` - (Optional) The JSON XSS key value expression.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwprofile_jsonxssurl_binding. It is a system-generated identifier.
* `alertonly` - Send SNMP alert?
* `comment` - Any comments about the purpose of profile, or other useful information about the profile.
* `isautodeployed` - Is the rule auto deployed by dynamic profile?
* `iskeyregex_json_xss` - Is the key name a regular expression?
* `isvalueregex_json_xss` - Is the JSON XSS key value a regular expression?
* `resourceid` - A "id" that identifies the rule.
* `state` - Enabled.
