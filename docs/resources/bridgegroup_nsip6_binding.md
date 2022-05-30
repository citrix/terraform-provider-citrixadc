---
subcategory: "Network"
---

# Resource: bridgegroup_nsip6_binding

The bridgegroup_nsip6_binding resource is used to bind nsip6 to bridgegroup resource.


## Example usage

```hcl
resource "citrixadc_bridgegroup" "tf_bridgegroup" {
  bridgegroup_id     = 2
  dynamicrouting     = "DISABLED"
  ipv6dynamicrouting = "DISABLED"
}
resource "citrixadc_nsip6" "test_nsip" {
  ipv6address = "2001:db8:100::fb/64"
  type        = "VIP"
  icmp        = "DISABLED"
}
resource "citrixadc_bridgegroup_nsip6_binding" "tf_binding" {
  bridgegroup_id = citrixadc_bridgegroup.tf_bridgegroup.bridgegroup_id
  ipaddress      = citrixadc_nsip6.test_nsip.ipv6address
}
```


## Argument Reference

* `bridgegroup_id` - (Required) The integer that uniquely identifies the bridge group. Minimum value =  1 Maximum value =  1000
* `ipaddress` - (Required) The IP address assigned to the  bridge group.
* `td` - (Optional) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0. Minimum value =  0 Maximum value =  4094
* `ownergroup` - (Optional) The owner node group in a Cluster for this vlan. Minimum length =  1
* `netmask` - (Optional) A subnet mask associated with the network address. Minimum length =  1


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the bridgegroup_nsip6_binding. It is the concatenation of both `bridgegroup_id` and `ipaddress` attributes separated by comma.


## Import

A bridgegroup_nsip6_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_bridgegroup_nsip6_binding.tf_binding 2,2001:db8:100::fb/64
```
