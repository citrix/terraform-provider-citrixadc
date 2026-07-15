---
subcategory: "Network"
---

# Resource: rnatparam

Configures the global Reverse Network Address Translation (RNAT) parameters on the Citrix ADC. These settings control how the appliance handles RNAT traffic, letting you enable TCP proxy optimization and source-IP persistency for connections that are source-NATed by the appliance. This is a singleton resource: a single RNAT parameter configuration always exists on the ADC, so creating this resource updates the existing global settings rather than creating a new object.


## Example usage

```hcl
resource "citrixadc_rnatparam" "tf_rnatparam" {
  tcpproxy         = "ENABLED"
  srcippersistency = "DISABLED"
}
```


## Argument Reference

* `tcpproxy` - (Optional) Enable TCP proxy, which enables the Citrix ADC to optimize the RNAT TCP traffic by using Layer 4 features. Possible values: [ ENABLED, DISABLED ]. Defaults to `"ENABLED"`.
* `srcippersistency` - (Optional) Enable source ip persistency, which enables the Citrix ADC to use the RNAT ips using source ip. Possible values: [ ENABLED, DISABLED ]. Defaults to `"DISABLED"`.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the rnatparam resource. Because this is a singleton resource, the id is a fixed synthetic string with the value `"rnatparam-config"`.


## Import

A rnatparam can be imported using its id, e.g.

```shell
terraform import citrixadc_rnatparam.tf_rnatparam rnatparam-config
```
