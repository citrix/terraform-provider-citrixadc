---
subcategory: "Network"
---

# Resource: rnatparam

The rnatparam resource is used to create RNAT parameter resource.


## Example usage

```hcl
resource "citrixadc_rnatparam" "tf_rnatparam" {
  tcpproxy         = "ENABLED"
  srcippersistency = "DISABLED"
}
```


## Argument Reference

* `tcpproxy` - (Optional) Enable TCP proxy, which enables the Citrix ADC to optimize the RNAT TCP traffic by using Layer 4 features. Possible values: [ ENABLED, DISABLED ]
* `srcippersistency` - (Optional) Enable source ip persistency, which enables the Citrix ADC to use the RNAT ips using source ip. Possible values: [ ENABLED, DISABLED ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the rnatparam. It is a unique string prefixed with "tf-rnatparam-"

