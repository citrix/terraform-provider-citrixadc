---
subcategory: "NS"
---

# Resource: nstrafficdomain_bridgegroup_binding

The nstrafficdomain_bridgegroup_binding resource is used to bind bridgegroup to nstrafficdomain resource.


## Example usage

```hcl
resource "citrixadc_nstrafficdomain" "tf_trafficdomain" {
  td        = 2
  aliasname = "tf_trafficdomain"
  vmac      = "DISABLED"
}
resource "citrixadc_bridgegroup" "tf_bridgegroup" {
  bridgegroup_id     = 2
  dynamicrouting     = "DISABLED"
  ipv6dynamicrouting = "DISABLED"
}
resource "citrixadc_nstrafficdomain_bridgegroup_binding" "tf_binding" {
  td          = citrixadc_nstrafficdomain.tf_trafficdomain.td
  bridgegroup = citrixadc_bridgegroup.tf_bridgegroup.bridgegroup_id
}
```


## Argument Reference

* `td` - (Required) Integer value that uniquely identifies a traffic domain. Minimum value =  1 Maximum value =  4094
* `bridgegroup` - (Required) ID of the configured bridge to bind to this traffic domain. More than one bridge group can be bound to a traffic domain, but the same bridge group cannot be a part of multiple traffic domains. Minimum value =  1 Maximum value =  1000


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nstrafficdomain_bridgegroup_binding. It is the concatenation of `td` and `bridgegroup` attributes separated by comma.


## Import

A nstrafficdomain_bridgegroup_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_nstrafficdomain_bridgegroup_binding.tf_binding 2,2
```
