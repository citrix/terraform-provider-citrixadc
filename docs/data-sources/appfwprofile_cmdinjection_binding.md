---
subcategory: "Application Firewall"
---

# Data Source: appfwprofile_cmdinjection_binding

The appfwprofile_cmdinjection_binding data source allows you to retrieve information about a cmdinjection binding to an application firewall profile.

## Example Usage

```terraform
data "citrixadc_appfwprofile_cmdinjection_binding" "tf_binding" {
  name                 = "tf_appfwprofile"
  cmdinjection         = "tf_cmdinjection"
  formactionurl_cmd    = "^https://sd2\\-zgw\\.test\\.ctxns\\.com/api/document/content$"
  as_scan_location_cmd = "HEADER"
  as_value_type_cmd    = "Keyword"
  as_value_expr_cmd    = "[a-z]+grep"
}

output "alertonly" {
  value = data.citrixadc_appfwprofile_cmdinjection_binding.tf_binding.alertonly
}

output "comment" {
  value = data.citrixadc_appfwprofile_cmdinjection_binding.tf_binding.comment
}

output "isvalueregex_cmd" {
  value = data.citrixadc_appfwprofile_cmdinjection_binding.tf_binding.isvalueregex_cmd
}
```

## Argument Reference

* `name` - (Required) Name of the profile to which to bind an exemption or rule.
* `cmdinjection` - (Required) Name of the relaxed web form field/header/cookie.
* `formactionurl_cmd` - (Required) The web form action URL.
* `as_scan_location_cmd` - (Required) Location of command injection exception - form field, header or cookie.
* `as_value_type_cmd` - (Optional) Type of the relaxed web form value.
* `as_value_expr_cmd` - (Optional) The web form/header/cookie value expression.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `alertonly` - Send SNMP alert? Possible values: [ ON, OFF ]
* `comment` - Any comments about the purpose of profile, or other useful information about the profile.
* `id` - The id of the appfwprofile_cmdinjection_binding. It is a system-generated identifier.
* `isautodeployed` - Is the rule auto deployed by dynamic profile? Possible values: [ AUTODEPLOYED, NOTAUTODEPLOYED ]
* `isregex_cmd` - Is the relaxed web form field name/header/cookie a regular expression? Possible values: [ REGEX, NOTREGEX ]
* `isvalueregex_cmd` - Is the web form field/header/cookie value a regular expression? Possible values: [ REGEX, NOTREGEX ]
* `resourceid` - A "id" that identifies the rule.
* `state` - Enabled. Possible values: [ ENABLED, DISABLED ]
