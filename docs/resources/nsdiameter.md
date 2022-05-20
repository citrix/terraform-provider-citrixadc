---
subcategory: "NS"
---

# Resource: nsdiameter

The nsdiameter resource is used to create Diameter Parameters resource.


## Example usage

```hcl
resource "citrixadc_nsdiameter" "tf_nsdiameter" {
  identity               = "citrixadc.com"
  realm                  = "com"
  serverclosepropagation = "OFF"
}
```


## Argument Reference

* `identity` - (Optional) DiameterIdentity to be used by NS. DiameterIdentity is used to identify a Diameter node uniquely. Before setting up diameter configuration, Citrix ADC (as a Diameter node) MUST be assigned a unique DiameterIdentity. example => set ns diameter -identity netscaler.com Now whenever Citrix ADC needs to use identity in diameter messages. It will use 'netscaler.com' as Origin-Host AVP as defined in RFC3588 . Minimum length =  1
* `realm` - (Optional) Diameter Realm to be used by NS. example => set ns diameter -realm com Now whenever Citrix ADC system needs to use realm in diameter messages. It will use 'com' as Origin-Realm AVP as defined in RFC3588 . Minimum length =  1
* `serverclosepropagation` - (Optional) when a Server connection goes down, whether to close the corresponding client connection if there were requests pending on the server. Possible values: [ YES, NO ]
* `ownernode` - (Optional) ID of the cluster node for which the diameter id is set, can be configured only through CLIP. Minimum value =  0 Maximum value =  31


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nsdiameter. It is a unique string prefixed with "tf-nsdiameter-"

