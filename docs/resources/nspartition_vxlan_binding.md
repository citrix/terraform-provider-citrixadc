---
subcategory: "NS"
---

# Resource: nspartition_vxlan_binding

The nspartition_vxlan_binding resource is used to bind vxlan to nspartition resource.


## Example usage

```hcl
resource "citrixadc_nspartition" "tf_nspartition" {
  partitionname = "tf_nspartition"
  maxbandwidth  = 10240
  minbandwidth  = 512
  maxconn       = 512
  maxmemlimit   = 11
}
resource "citrixadc_vxlan" "tf_vxlan" {
  vxlanid            = 123
  port               = 33
  dynamicrouting     = "DISABLED"
  ipv6dynamicrouting = "DISABLED"
  innervlantagging   = "ENABLED"
}
resource "citrixadc_nspartition_vxlan_binding" "tf_binding" {
  partitionname = citrixadc_nspartition.tf_nspartition.partitionname
  vxlan         = citrixadc_vxlan.tf_vxlan.vxlanid
}
```


## Argument Reference

* `vxlan` - (Required) Identifier of the vxlan that is assigned to this partition. Minimum value =  1 Maximum value =  16777215
* `partitionname` - (Required) Name of the Partition. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Minimum length =  1


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nspartition_vxlan_binding. It ids concatenation of `partitionname` and `vxlan` attributes separated by comma.


## Import

A nspartition_vxlan_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_nspartition_vxlan_binding.tf_binding tf_nspartition,123
```
