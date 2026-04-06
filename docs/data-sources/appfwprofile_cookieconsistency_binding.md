---
subcategory: "Application Firewall"
---

# Data Source: appfwprofile_cookieconsistency_binding

The appfwprofile_cookieconsistency_binding data source allows you to retrieve information about the bindings between appfwprofile and cookieconsistency relaxation rules.

## Example Usage

```terraform
data "citrixadc_appfwprofile_cookieconsistency_binding" "demo_binding" {
  name              = "demo_appfwprofile"
  cookieconsistency = "^logon_[0-9A-Za-z]{2,15}$"
}

output "cookieconsistency_state" {
  value = data.citrixadc_appfwprofile_cookieconsistency_binding.demo_binding.state
}

output "alertonly" {
  value = data.citrixadc_appfwprofile_cookieconsistency_binding.demo_binding.alertonly
}
```

## Argument Reference

* `name` - (Required) Name of the profile to which the binding belongs.
* `cookieconsistency` - (Required) The name of the cookie to be checked.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwprofile_cookieconsistency_binding. It is the concatenation of the `name` and `cookieconsistency` attributes separated by a comma.
* `alertonly` - Send SNMP alert?
* `comment` - Any comments about the purpose of profile, or other useful information about the profile.
* `isautodeployed` - Is the rule auto deployed by dynamic profile ?
* `isregex` - Is the cookie name a regular expression?
* `resourceid` - A "id" that identifies the rule.
* `state` - Enabled.
