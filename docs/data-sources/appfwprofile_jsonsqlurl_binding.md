---
subcategory: "Application Firewall"
---

# Data Source: appfwprofile_jsonsqlurl_binding

The appfwprofile_jsonsqlurl_binding data source allows you to retrieve information about appfwprofile jsonsqlurl bindings.

## Example Usage

```terraform
data "citrixadc_appfwprofile_jsonsqlurl_binding" "tf_binding" {
  name                   = "tf_appfwprofile"
  jsonsqlurl             = "[abc][a-z]a*"
  keyname_json_sql       = "id"
  as_value_type_json_sql = "SpecialString"
  as_value_expr_json_sql = "p"
}

output "state" {
  value = data.citrixadc_appfwprofile_jsonsqlurl_binding.tf_binding.state
}

output "comment" {
  value = data.citrixadc_appfwprofile_jsonsqlurl_binding.tf_binding.comment
}
```

## Argument Reference

* `name` - (Required) Name of the profile to which to bind an exemption or rule.
* `jsonsqlurl` - (Required) A regular expression that designates a URL on the Json SQL URL list for which SQL violations are relaxed.
* `keyname_json_sql` - (Required) An expression that designates a keyname on the JSON SQL URL for which SQL injection violations are relaxed.
* `as_value_type_json_sql` - (Required) Type of the relaxed JSON SQL key value.
* `as_value_expr_json_sql` - (Required) The JSON SQL key value expression.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwprofile_jsonsqlurl_binding. It is the concatenation of the `name`, `jsonsqlurl`, `keyname_json_sql`, `as_value_type_json_sql` and `as_value_expr_json_sql` attributes separated by a comma.
* `alertonly` - Send SNMP alert?
* `comment` - Any comments about the purpose of profile, or other useful information about the profile.
* `isautodeployed` - Is the rule auto deployed by dynamic profile?
* `iskeyregex_json_sql` - Is the key name a regular expression?
* `isvalueregex_json_sql` - Is the JSON SQL key value a regular expression?
* `resourceid` - A "id" that identifies the rule.
* `ruletype` - Specifies rule type of binding.
* `state` - Enabled.
