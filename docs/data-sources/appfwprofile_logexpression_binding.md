---
subcategory: "Application Firewall"
---

# Data Source: appfwprofile_logexpression_binding

The appfwprofile_logexpression_binding data source allows you to retrieve information about a specific logexpression binding to an appfwprofile resource.

## Example Usage

```terraform
data "citrixadc_appfwprofile_logexpression_binding" "tf_binding" {
  name          = "tf_appfwprofile"
  logexpression = "tf_logexp"
}

output "as_logexpression" {
  value = data.citrixadc_appfwprofile_logexpression_binding.tf_binding.as_logexpression
}

output "state" {
  value = data.citrixadc_appfwprofile_logexpression_binding.tf_binding.state
}
```

## Argument Reference

* `name` - (Required) Name of the profile to which to bind an exemption or rule.
* `logexpression` - (Required) Name of LogExpression object.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `alertonly` - Send SNMP alert?
* `as_logexpression` - LogExpression to log when violation happened on appfw profile.
* `comment` - Any comments about the purpose of profile, or other useful information about the profile.
* `id` - The id of the appfwprofile_logexpression_binding. It is a system-generated identifier.
* `isautodeployed` - Is the rule auto deployed by dynamic profile?
* `resourceid` - A "id" that identifies the rule.
* `state` - Enabled.
