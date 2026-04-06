---
subcategory: "Application Firewall"
---

# Data Source: appfwprofile_safeobject_binding

The `citrixadc_appfwprofile_safeobject_binding` data source allows you to retrieve information about a specific safe object binding for an Application Firewall profile.

## Example Usage

```terraform
data "citrixadc_appfwprofile_safeobject_binding" "example" {
  name       = "tf_appfwprofile"
  safeobject = "tf_safeobject"
}

output "as_expression" {
  value = data.citrixadc_appfwprofile_safeobject_binding.example.as_expression
}

output "state" {
  value = data.citrixadc_appfwprofile_safeobject_binding.example.state
}
```

## Argument Reference

The following arguments are required:

* `name` - (Required) Name of the profile to which to bind an exemption or rule.
* `safeobject` - (Required) Name of the Safe Object.

## Attribute Reference

In addition to the arguments, the following attributes are exported:

* `id` - The ID of the binding.
* `action` - Safe Object action types. (BLOCK | LOG | STATS | NONE)
* `alertonly` - Send SNMP alert?
* `as_expression` - A regular expression that defines the Safe Object.
* `comment` - Any comments about the purpose of profile, or other useful information about the profile.
* `isautodeployed` - Is the rule auto deployed by dynamic profile?
* `maxmatchlength` - Maximum match length for a Safe Object expression.
* `resourceid` - A "id" that identifies the rule.
* `state` - Enabled.
