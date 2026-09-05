---
subcategory: "Tunnel"
---

# Data Source: tunneltrafficpolicy

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
* `id` - The id of the tunneltrafficpolicy. It has the same value as the `name` attribute.

### Read-only tunneltrafficpolicy metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_tunneltrafficpolicy` resource). They are Computed/GET-only, and any attribute the appliance does not return is `null`.

* `expressiontype` - Type of policy (Classic/Advanced).
* `hits` - Number of hits.
* `undefhits` - Number of policy UNDEF hits.
* `txbytes` - Number of bytes transmitted.
* `rxbytes` - Number of bytes received.
* `clientttlb` - Total client TTLB value.
* `clienttransactions` - Number of client transactions.
* `serverttlb` - Total server TTLB value.
* `servertransactions` - Number of server transactions.
* `isdefault` - A value of true is returned if it is a default tunnelpolicy.
* `builtin` - Indicates that a variable is a built-in (SYSTEM INTERNAL) type. A list of strings.
* `feature` - The feature to be checked while applying this config.
