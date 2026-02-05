---
subcategory: "Tunnel"
---

# Data Source `tunneltrafficpolicy`

The tunneltrafficpolicy data source allows you to retrieve information about a tunnel traffic policy.


## Example usage

```terraform
data "citrixadc_tunneltrafficpolicy" "tf_tunneltrafficpolicy" {
  name = "my_tunneltrafficpolicy"
}

output "action" {
  value = data.citrixadc_tunneltrafficpolicy.tf_tunneltrafficpolicy.action
}

output "rule" {
  value = data.citrixadc_tunneltrafficpolicy.tf_tunneltrafficpolicy.rule
}
```


## Argument Reference

* `name` - (Required) Name for the tunnel traffic policy. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `action` - Name of the built-in compression action to associate with the policy.
* `comment` - Any comments to preserve information about this policy.
* `logaction` - Name of the messagelog action to use for requests that match this policy.
* `newname` - New name for the tunnel traffic policy. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.
* `rule` - Expression, against which traffic is evaluated.

## Attribute Reference

* `id` - The id of the tunneltrafficpolicy. It has the same value as the `name` attribute.


## Import

A tunneltrafficpolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_tunneltrafficpolicy.tf_tunneltrafficpolicy my_tunneltrafficpolicy
```
