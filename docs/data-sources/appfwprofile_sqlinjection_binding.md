---
subcategory: "Application Firewall"
---

# Data Source: appfwprofile_sqlinjection_binding

The `citrixadc_appfwprofile_sqlinjection_binding` data source allows you to retrieve information about a specific SQL injection binding for an Application Firewall profile.

## Example Usage

```terraform
data "citrixadc_appfwprofile_sqlinjection_binding" "example" {
  name                 = "demo_appfwprofile"
  sqlinjection         = "data"
  as_scan_location_sql = "FORMFIELD"
  formactionurl_sql    = "^https://citrix.csg.com/analytics/saw.dll$"
  as_value_type_sql    = "Keyword"
  as_value_expr_sql    = ".*"
}

output "alertonly" {
  value = data.citrixadc_appfwprofile_sqlinjection_binding.example.alertonly
}

output "state" {
  value = data.citrixadc_appfwprofile_sqlinjection_binding.example.state
}
```

## Argument Reference

The following arguments are required:

* `name` - (Required) Name of the profile to which to bind an exemption or rule.
* `sqlinjection` - (Required) The web form field name.
* `as_scan_location_sql` - (Required) Location of SQL injection exception - form field, header or cookie.
* `formactionurl_sql` - (Required) The web form action URL.
* `as_value_type_sql` - (optional) The web form value type.
* `as_value_expr_sql` - (optional) The web form value expression.
* `ruletype` - (optional) Specifies rule type of binding.

## Attribute Reference

In addition to the arguments, the following attributes are exported:

* `id` - The ID of the binding.
* `alertonly` - Send SNMP alert?
* `comment` - Any comments about the purpose of profile, or other useful information about the profile.
* `isautodeployed` - Is the rule auto deployed by dynamic profile?
* `isregex_sql` - Is the web form field name a regular expression?
* `isvalueregex_sql` - Is the web form field value a regular expression?
* `resourceid` - A "id" that identifies the rule.
* `state` - Enabled.
