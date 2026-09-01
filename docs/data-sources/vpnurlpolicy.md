---
subcategory: "VPN"
---

# Data Source: vpnurlpolicy

The vpnurlpolicy data source allows you to retrieve information about a VPN URL policy.

## Example usage

```terraform
data "citrixadc_vpnurlpolicy" "example" {
  name = "new_policy"
}

output "action" {
  value = data.citrixadc_vpnurlpolicy.example.action
}

output "rule" {
  value = data.citrixadc_vpnurlpolicy.example.rule
}
```

## Argument Reference

* `name` - (Required) Name for the new urlPolicy.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `action` - Action to be applied by the new urlPolicy if the rule criteria are met.
* `comment` - Any comments to preserve information about this policy.
* `logaction` - Name of messagelog action to use when a request matches this policy.
* `newname` - New name for the vpn urlPolicy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.
* `rule` - Expression, or name of a named expression, specifying the traffic that matches the policy.

### Read-only vpnurlpolicy metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_vpnurlpolicy` resource). They are Computed/GET-only, and any attribute the appliance does not return is `null`.

* `builtin` - Indicates that a variable is a built-in (SYSTEM INTERNAL) type. A list of strings.
* `undefhits` - The number of times the policy evaluation resulted in undefined processing.
