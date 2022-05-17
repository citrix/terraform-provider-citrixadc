---
subcategory: "NS"
---

# Resource: nstrafficdomain

The nstrafficdomain resource is used to create Traffic Domain resource.


## Example usage

```hcl
resource "citrixadc_nstrafficdomain" "tf_trafficdomain" {
  td        = 2
  aliasname = "tf_trafficdomain"
  vmac      = "ENABLED"
}
```


## Argument Reference

* `td` - (Required) Integer value that uniquely identifies a traffic domain. Minimum value =  1 Maximum value =  4094
* `aliasname` - (Optional) Name of traffic domain  being added. Minimum length =  1 Maximum length =  31
* `vmac` - (Optional) Associate the traffic domain with a VMAC address instead of with VLANs. The Citrix ADC then sends the VMAC address of the traffic domain in all responses to ARP queries for network entities in that domain. As a result, the ADC can segregate subsequent incoming traffic for this traffic domain on the basis of the destination MAC address, because the destination MAC address is the VMAC address of the traffic domain. After creating entities on a traffic domain, you can easily manage and monitor them by performing traffic domain level operations. Possible values: [ ENABLED, DISABLED ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nstrafficdomain. It has the same value as the `td` attribute.


## Import

A nstrafficdomain can be imported using its td, e.g.

```shell
terraform import citrixadc_nstrafficdomain.tf_trafficdomain 2
```
