---
subcategory: "Application Firewall"
---

# Data Source: appfwprofile_sqlinjection_binding

The `appfwprofile_sqlinjection_binding` data source allows you to retrieve information about a specific HTML SQL injection binding for an Application Firewall profile.


## Example Usage

```terraform
data "citrixadc_appfwprofile_sqlinjection_binding" "example" {
  name                 = "demo_appfwprofile"
  sqlinjection         = "data"
  formactionurl_sql    = "^https://citrix.csg.com/analytics/saw.dll$"
  as_scan_location_sql = "FORMFIELD"
  as_value_type_sql    = "Keyword"
  as_value_expr_sql    = ".*"
}

output "ruletype" {
  value = data.citrixadc_appfwprofile_sqlinjection_binding.example.ruletype
}

output "state" {
  value = data.citrixadc_appfwprofile_sqlinjection_binding.example.state
}
```


## Argument Reference

* `name` - (Required) Name of the profile to which to bind an exemption or rule.
* `sqlinjection` - (Required) The web form field name.
* `formactionurl_sql` - (Required) The web form action URL.
* `as_scan_location_sql` - (Required) Location of SQL injection exception - form field, header or cookie.
* `as_value_type_sql` - (Required) The web form value type.
* `as_value_expr_sql` - (Required) The web form value expression.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the `appfwprofile_sqlinjection_binding`. It is the concatenation of the `name`, `sqlinjection`, `formactionurl_sql`, `as_scan_location_sql`, `as_value_type_sql` and `as_value_expr_sql` attributes separated by a comma.
* `alertonly` - Send SNMP alert?
* `comment` - Any comments about the purpose of profile, or other useful information about the profile.
* `isautodeployed` - Is the rule auto deployed by dynamic profile?
* `isregex_sql` - Is the web form field name a regular expression?
* `isvalueregex_sql` - Is the web form field value a regular expression?
* `resourceid` - A "id" that identifies the rule.
* `ruletype` - Specifies rule type of binding.
* `state` - Enabled.
