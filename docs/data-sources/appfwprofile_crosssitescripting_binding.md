---
subcategory: "Application Firewall"
---

# Data Source: appfwprofile_crosssitescripting_binding

The `citrixadc_appfwprofile_crosssitescripting_binding` data source allows you to retrieve information about a specific cross-site scripting binding for an Application Firewall profile.

## Example Usage

```terraform
data "citrixadc_appfwprofile_crosssitescripting_binding" "example" {
  name                 = "tf_appfwprofile"
  crosssitescripting   = "file"
  formactionurl_xss    = "^https://sd2\\-zgw\\.test\\.ctxns\\.com/api/document/content$"
  as_scan_location_xss = "FORMFIELD"
  as_value_type_xss    = "Tag"
  as_value_expr_xss    = ".*"
}

output "alertonly" {
  value = data.citrixadc_appfwprofile_crosssitescripting_binding.example.alertonly
}

output "state" {
  value = data.citrixadc_appfwprofile_crosssitescripting_binding.example.state
}
```

## Argument Reference

The following arguments are required:

* `name` - (Required) Name of the profile to which to bind an exemption or rule.
* `crosssitescripting` - (Required) The web form field name.
* `formactionurl_xss` - (Required) The web form action URL.
* `as_scan_location_xss` - (Required) Location of cross-site scripting exception - form field, header, cookie or URL.
* `as_value_type_xss` - (Optional) The web form value type.
* `as_value_expr_xss` - (Optional) The web form value expression.

## Attribute Reference

In addition to the arguments, the following attributes are exported:

* `id` - The ID of the binding.
* `alertonly` - Send SNMP alert?
* `comment` - Any comments about the purpose of profile, or other useful information about the profile.
* `isautodeployed` - Is the rule auto deployed by dynamic profile?
* `isregex_xss` - Is the web form field name a regular expression?
* `isvalueregex_xss` - Is the web form field value a regular expression?
* `resourceid` - A "id" that identifies the rule.
* `state` - Enabled.
