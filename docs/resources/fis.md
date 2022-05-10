---
subcategory: "Network"
---

# Resource: fis

The fis resource is used to create fis resource.


## Example usage

```hcl
resource "citrixadc_fis" "tf_fis" {
  name = "tf_fis"  
}
```


## Argument Reference

* `name` - (Required) Name for the FIS to be created. Leading character must be a number or letter. Other characters allowed, after the first character, are @ _ - . (period) : (colon) # and space ( ). Note: In a cluster setup, the FIS name on each node must be unique.
* `ownernode` - (Optional) ID of the cluster node for which you are creating the FIS. Can be configured only through the cluster IP address.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the fis. It has the same value as the `name` attribute.


## Import

A fis can be imported using its name, e.g.

```shell
terraform import citrixadc_fis.tf_fis tf_fis
```
