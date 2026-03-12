---
subcategory: "Application Firewall"
---

# Data Source: appfwprofile_starturl_binding

The appfwprofile_starturl_binding data source allows you to retrieve information about appfwprofile starturl bindings.

## Example Usage

```terraform
data "citrixadc_appfwprofile_starturl_binding" "tf_binding" {
  name     = "tfAcc_appfwprofile"
  starturl = "^[^?]+[.](html?|shtml|js|gif|jpg|jpeg|png|swf|pif|pdf|css|csv)$"
}

output "state" {
  value = data.citrixadc_appfwprofile_starturl_binding.tf_binding.state
}

output "alertonly" {
  value = data.citrixadc_appfwprofile_starturl_binding.tf_binding.alertonly
}
```

## Argument Reference

* `name` - (Required) Name of the profile to which to bind an exemption or rule.
* `starturl` - (Required) A regular expression that designates a URL on the Start URL list.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwprofile_starturl_binding. It is a system-generated identifier.
* `alertonly` - Send SNMP alert?
* `comment` - Any comments about the purpose of profile, or other useful information about the profile.
* `isautodeployed` - Is the rule auto deployed by dynamic profile?
* `resourceid` - A "id" that identifies the rule.
* `state` - Enabled.
