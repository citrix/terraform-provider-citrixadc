---
subcategory: "Network"
---

# Resource: bridgegroup

The bridgegroup resource is used to create bridge group resource.


## Example usage

```hcl
resource "citrixadc_bridgegroup" "tf_bridgegroup" {
  bridgegroup_id     = 2
  dynamicrouting     = "DISABLED"
  ipv6dynamicrouting = "DISABLED"
}
```


## Argument Reference

* `bridgegroup_id` - (Required) An integer that uniquely identifies the bridge group. Minimum value =  1 Maximum value =  1000
* `dynamicrouting` - (Optional) Enable dynamic routing for this bridgegroup. Possible values: [ ENABLED, DISABLED ]
* `ipv6dynamicrouting` - (Optional) Enable all IPv6 dynamic routing protocols on all VLANs bound to this bridgegroup. Note: For the ENABLED setting to work, you must configure IPv6 dynamic routing protocols from the VTYSH command line. Possible values: [ ENABLED, DISABLED ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the bridgegroup. It has the same value as the `bridgegroup_id` attribute.


## Import

A bridgegroup can be imported using its bridgegroup_id, e.g.

```shell
terraform import citrixadc_bridgegroup.tf_bridgegroup 2
```
