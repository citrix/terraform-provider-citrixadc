---
subcategory: "Application Firewall"
---

# Data Source: appfwprofile_bypasslist_binding

The appfwprofile_bypasslist_binding data source allows you to retrieve information about a bypass-list (relaxation) binding on an application firewall profile.


## Example usage

```terraform
data "citrixadc_appfwprofile_bypasslist_binding" "tf_binding" {
  name                      = "tf_appfwprofile"
  as_bypass_list            = "X-Forwarded-For"
  as_bypass_list_location   = "HEADER"
  as_bypass_list_value_type = "Keyword"
}

output "state" {
  value = data.citrixadc_appfwprofile_bypasslist_binding.tf_binding.state
}

output "resourceid" {
  value = data.citrixadc_appfwprofile_bypasslist_binding.tf_binding.resourceid
}
```


## Argument Reference

* `name` - (Required) Name of the application firewall profile to which the bypass-list rule is bound.
* `as_bypass_list` - (Required) The bypass-list value bound to the profile.
* `as_bypass_list_location` - (Required) The scan location to which the bypass-list rule applies.
* `as_bypass_list_value_type` - (Required) The value type of the bypass-list entry.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwprofile_bypasslist_binding. It is a composite identifier composed of comma-separated `key:value` pairs, in the format `as_bypass_list:<as_bypass_list>,as_bypass_list_location:<as_bypass_list_location>,as_bypass_list_value_type:<as_bypass_list_value_type>,name:<name>`.
* `as_bypass_list_action` - The action to take when the bypass-list rule matches. Possible values: [ none, log ]
* `comment` - Any comments about the purpose of the profile, or other useful information about the profile.
* `state` - Whether the bypass-list rule is enabled. Possible values: [ ENABLED, DISABLED ]
* `isautodeployed` - Indicates whether the rule was auto-deployed by a dynamic profile.
* `resourceid` - A system-generated identifier that identifies the rule.
* `alertonly` - Indicates whether an SNMP alert is sent.
```
