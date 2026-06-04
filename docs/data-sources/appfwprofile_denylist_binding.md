---
subcategory: "Application Firewall"
---

# Data Source: appfwprofile_denylist_binding

The appfwprofile_denylist_binding data source allows you to retrieve information about a deny-list binding on an application firewall profile.


## Example usage

```terraform
data "citrixadc_appfwprofile_denylist_binding" "tf_binding" {
  name                    = "tf_appfwprofile"
  as_deny_list            = "X-Forwarded-For"
  as_deny_list_location   = "HEADER"
  as_deny_list_value_type = "Keyword"
}

output "state" {
  value = data.citrixadc_appfwprofile_denylist_binding.tf_binding.state
}

output "resourceid" {
  value = data.citrixadc_appfwprofile_denylist_binding.tf_binding.resourceid
}
```


## Argument Reference

* `name` - (Required) Name of the application firewall profile to which the deny-list rule is bound.
* `as_deny_list` - (Required) The deny-list value bound to the profile.
* `as_deny_list_location` - (Required) The scan location to which the deny-list rule applies.
* `as_deny_list_value_type` - (Required) The value type of the deny-list entry.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwprofile_denylist_binding. It is a composite identifier composed of comma-separated `key:value` pairs, in the format `as_deny_list:<as_deny_list>,as_deny_list_location:<as_deny_list_location>,as_deny_list_value_type:<as_deny_list_value_type>,name:<name>`.
* `as_deny_list_action` - The action(s) to take when the deny-list rule matches, as a list of strings. Possible values: [ none, log, RESET, REDIRECT ]
* `comment` - Any comments about the purpose of the profile, or other useful information about the profile.
* `state` - Whether the deny-list rule is enabled. Possible values: [ ENABLED, DISABLED ]
* `isautodeployed` - Indicates whether the rule was auto-deployed by a dynamic profile.
* `resourceid` - A system-generated identifier that identifies the rule.
* `alertonly` - Indicates whether an SNMP alert is sent.
```
