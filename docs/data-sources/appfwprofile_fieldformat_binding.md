---
subcategory: "Application Firewall"
---

# Data Source: appfwprofile_fieldformat_binding

The appfwprofile_fieldformat_binding data source allows you to retrieve information about a fieldformat binding to an application firewall profile.

## Example Usage

```terraform
data "citrixadc_appfwprofile_fieldformat_binding" "tf_binding" {
  name             = "tf_appfwprofile"
  fieldformat      = "tf_field"
  formactionurl_ff = "^https://sd2\\-zgw\\.test\\.ctxns\\.com/api/document/content$"
}

output "comment" {
  value = data.citrixadc_appfwprofile_fieldformat_binding.tf_binding.comment
}

output "state" {
  value = data.citrixadc_appfwprofile_fieldformat_binding.tf_binding.state
}

output "fieldtype" {
  value = data.citrixadc_appfwprofile_fieldformat_binding.tf_binding.fieldtype
}
```

## Argument Reference

* `name` - (Required) Name of the profile to which to bind an exemption or rule.
* `fieldformat` - (Required) Name of the form field to which a field format will be assigned.
* `formactionurl_ff` - (Required) Action URL of the form field to which a field format will be assigned.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `alertonly` - Send SNMP alert?
* `comment` - Any comments about the purpose of profile, or other useful information about the profile.
* `fieldformatmaxlength` - The maximum allowed length for data in this form field.
* `fieldformatminlength` - The minimum allowed length for data in this form field.
* `fieldtype` - The field type you are assigning to this form field.
* `id` - The id of the appfwprofile_fieldformat_binding. It is a system-generated identifier.
* `isautodeployed` - Is the rule auto deployed by dynamic profile?
* `isregexff` - Is the form field name a regular expression?
* `resourceid` - A "id" that identifies the rule.
* `state` - Enabled.
