---
subcategory: "NS"
---

# Resource: nspartition_bridgegroup_binding

The nspartition_bridgegroup_binding resource is used to bind bridgegroup to nspartition resource.


## Example usage

```hcl
resource "citrixadc_nspartition" "tf_nspartition" {
  partitionname = "tf_nspartition"
  maxbandwidth  = 10240
  minbandwidth  = 512
  maxconn       = 512
  maxmemlimit   = 11
}
resource "citrixadc_bridgegroup" "tf_bridgegroup" {
  bridgegroup_id     = 2
  dynamicrouting     = "DISABLED"
  ipv6dynamicrouting = "DISABLED"
}
resource "citrixadc_nspartition_bridgegroup_binding" "tf_binding" {
  partitionname = citrixadc_nspartition.tf_nspartition.partitionname
  bridgegroup   = citrixadc_bridgegroup.tf_bridgegroup.bridgegroup_id
}
```


## Argument Reference

* `partitionname` - (Required) Name of the Partition. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Minimum length =  1
* `bridgegroup` - (Required) Identifier of the bridge group that is assigned to this partition. Minimum value =  1 Maximum value =  1000


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nspartition_bridgegroup_binding. It is the concatenation of `partitionname` and `bridgeroup` attributes separated by comma.


## Import

A nspartition_bridgegroup_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_nspartition_bridgegroup_binding.tf_binding tf_nspartition,2
```
