---
subcategory: "VPN"
---

# Data Source: vpnsessionpolicy

The vpnsessionpolicy data source allows you to retrieve information about a VPN session policy.

## Example usage

```terraform
data "citrixadc_vpnsessionpolicy" "foo" {
  name = "tf_vpnsessionpolicy"
}

output "rule" {
  value = data.citrixadc_vpnsessionpolicy.foo.rule
}

output "action" {
  value = data.citrixadc_vpnsessionpolicy.foo.action
}
```

## Argument Reference

* `name` - (Required) Name for the new session policy that is applied after the user logs on to Citrix Gateway.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnsessionpolicy. It is the same as the `name` attribute.
* `action` - Action to be applied by the new session policy if the rule criteria are met.
* `rule` - Expression, or name of a named expression, specifying the traffic that matches the policy.
