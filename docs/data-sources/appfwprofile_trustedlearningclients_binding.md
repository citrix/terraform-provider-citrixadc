---
subcategory: "Application Firewall"
---

# Data Source: appfwprofile_trustedlearningclients_binding

The appfwprofile_trustedlearningclients_binding data source allows you to retrieve information about a specific trusted learning client binding to an application firewall profile.

## Example Usage

```terraform
data "citrixadc_appfwprofile_trustedlearningclients_binding" "tf_binding" {
  name                   = "tf_appfwprofile"
  trustedlearningclients = "1.2.31.1/32"
}

output "state" {
  value = data.citrixadc_appfwprofile_trustedlearningclients_binding.tf_binding.state
}

output "comment" {
  value = data.citrixadc_appfwprofile_trustedlearningclients_binding.tf_binding.comment
}
```

## Argument Reference

The following arguments are required:

* `name` - (Required) Name of the profile to which to bind an exemption or rule.
* `trustedlearningclients` - (Required) Specify trusted host/network IP.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwprofile_trustedlearningclients_binding. It is a system-generated identifier.
* `alertonly` - Send SNMP alert?
* `comment` - Any comments about the purpose of profile, or other useful information about the profile.
* `isautodeployed` - Is the rule auto deployed by dynamic profile?
* `resourceid` - A "id" that identifies the rule.
* `state` - Enabled.
