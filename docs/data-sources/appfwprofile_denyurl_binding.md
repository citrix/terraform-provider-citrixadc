---
subcategory: "Application Firewall"
---

# Data Source: appfwprofile_denyurl_binding

The appfwprofile_denyurl_binding data source allows you to retrieve information about Application Firewall Profile to DenyURL binding.

## Example Usage

```terraform
data "citrixadc_appfwprofile_denyurl_binding" "tf_binding" {
  name    = "tf_appfwprofile"
  denyurl = "test[.][^/?]*(|[?].*)$"
}

output "alertonly" {
  value = data.citrixadc_appfwprofile_denyurl_binding.tf_binding.alertonly
}

output "state" {
  value = data.citrixadc_appfwprofile_denyurl_binding.tf_binding.state
}
```

## Argument Reference

* `name` - (Required) Name of the profile to which to bind an exemption or rule.
* `denyurl` - (Required) A regular expression that designates a URL on the Deny URL list.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwprofile_denyurl_binding. It is a system-generated identifier.
* `alertonly` - Send SNMP alert?
* `comment` - Any comments about the purpose of profile, or other useful information about the profile.
* `isautodeployed` - Is the rule auto deployed by dynamic profile ?
* `resourceid` - A "id" that identifies the rule.
* `state` - Enabled.
