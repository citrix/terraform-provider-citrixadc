---
subcategory: "Application Firewall"
---

# Data Source: appfwprofile_fieldconsistency_binding

The appfwprofile_fieldconsistency_binding data source allows you to retrieve information about Application Firewall Profile to Field Consistency binding.

## Example Usage

```terraform
data "citrixadc_appfwprofile_fieldconsistency_binding" "tf_binding" {
  name              = "tf_appfwprofile"
  fieldconsistency  = "tf_field"
  formactionurl_ffc = "^https://sd2\\-zgw\\.test\\.ctxns\\.com/api/document/content$"
}

output "state" {
  value = data.citrixadc_appfwprofile_fieldconsistency_binding.tf_binding.state
}

output "alertonly" {
  value = data.citrixadc_appfwprofile_fieldconsistency_binding.tf_binding.alertonly
}
```

## Argument Reference

* `name` - (Required) Name of the profile to which to bind an exemption or rule.
* `fieldconsistency` - (Required) The web form field name.
* `formactionurl_ffc` - (Required) The web form action URL.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwprofile_fieldconsistency_binding. It is a system-generated identifier.
* `alertonly` - Send SNMP alert?
* `comment` - Any comments about the purpose of profile, or other useful information about the profile.
* `isautodeployed` - Is the rule auto deployed by dynamic profile ?
* `isregex_ffc` - Is the web form field name a regular expression?
* `resourceid` - A "id" that identifies the rule.
* `state` - Enabled.
