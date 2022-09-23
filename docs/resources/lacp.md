---
subcategory: "Network"
---

# Resource: lacp

The lacp resource is used to configure aggregation control protocol lacp.


## Example usage

```hcl
resource "citrixadc_lacp" "tf_lacp" {
  syspriority = 30
  ownernode   = 0
}
```


## Argument Reference

* `syspriority` - (Required) Priority number that determines which peer device of an LACP LA channel can have control over the LA channel. This parameter is globally applied to all LACP channels on the Citrix ADC. The lower the number, the higher the priority. Minimum value =  1 Maximum value =  65535
* `ownernode` - (Required) The owner node in a cluster for which we want to set the lacp priority. Owner node can vary from 0 to 31. Ownernode value of 254 is used for Cluster.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lacp. It has the same value as the `ownernode` attribute.


## Import

A lacp can be imported using its name, e.g.

```shell
terraform import citrixadc_lacp.tf_lacp 0
```
