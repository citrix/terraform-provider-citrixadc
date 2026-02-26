---
subcategory: "Application Firewall"
---

# Data Source: appfwprofile_excluderescontenttype_binding

The appfwprofile_excluderescontenttype_binding data source allows you to retrieve information about Application Firewall Profile to Exclude Response Content-Type binding.

## Example Usage

```terraform
data "citrixadc_appfwprofile_excluderescontenttype_binding" "tf_binding" {
  name                  = "tf_appfwprofile"
  excluderescontenttype = "expressionexample"
}

output "alertonly" {
  value = data.citrixadc_appfwprofile_excluderescontenttype_binding.tf_binding.alertonly
}

output "state" {
  value = data.citrixadc_appfwprofile_excluderescontenttype_binding.tf_binding.state
}
```

## Argument Reference

* `name` - (Required) Name of the profile to which to bind an exemption or rule.
* `excluderescontenttype` - (Required) A regular expression that represents the content type of the response that are to be excluded from inspection.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwprofile_excluderescontenttype_binding. It is a system-generated identifier.
* `alertonly` - Send SNMP alert?
* `comment` - Any comments about the purpose of profile, or other useful information about the profile.
* `isautodeployed` - Is the rule auto deployed by dynamic profile ?
* `resourceid` - A "id" that identifies the rule.
* `state` - Enabled.
