---
subcategory: "VPN"
---

# Data Source: vpntrafficpolicy

The vpntrafficpolicy data source allows you to retrieve information about a VPN traffic policy.

## Example usage

```terraform
data "citrixadc_vpntrafficpolicy" "tf_vpntrafficpolicy" {
  name = "tf_vpntrafficpolicy"
}

output "rule" {
  value = data.citrixadc_vpntrafficpolicy.tf_vpntrafficpolicy.rule
}

output "action" {
  value = data.citrixadc_vpntrafficpolicy.tf_vpntrafficpolicy.action
}
```

## Argument Reference

* `name` - (Required) Name for the traffic policy. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpntrafficpolicy. It is the same as the `name` attribute.
* `action` - Action to apply to traffic that matches the policy.
* `rule` - Expression, or name of a named expression, against which traffic is evaluated.
