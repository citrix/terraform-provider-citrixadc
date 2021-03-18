---
subcategory: "QuickBridge"
---

# Resource: quicbridgeprofile

The `quicbridgeprofile` resource is used to create Citrix ADC Quick Bridge Profiles.

## Example usage

``` hcl
resource citrixadc_quicbridgeprofile demo_quicbridge {
  name             = "demo_quicbridge"
  routingalgorithm = "PLAINTEXT" # OPTIONAL
  serveridlength   = 4           # OPTIONAL
}
resource "citrixadc_lbvserver" "demo_quicbridge_lbvserver" {
  name                  = "demo_quicbridge_vserver"
  ipv46                 = "10.202.11.11"
  lbmethod              = "TOKEN"
  persistencetype       = "CUSTOMSERVERID"
  port                  = 8080
  servicetype           = "QUIC_BRIDGE"
  quicbridgeprofilename = citrixadc_quicbridgeprofile.demo_quicbridge.name
}
```

## Argument Reference

* `name` - Name for the QUIC profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals sign (=), and hyphen (-) characters. Cannot be changed after the profile is created.
* `routingalgorithm` - (Optional) Routing algorithm to generate routable connection IDs.
* `serveridlength` - (Optional) Length of serverid to encode/decode server information.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the quicbridgeprofile. It has the same value as the `name` attribute.

## Import


