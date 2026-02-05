---
subcategory: "Network"
---

# Data Source `fis`

The fis data source allows you to retrieve information about FIS (Forward Information Base) configurations.


## Example usage

```terraform
data "citrixadc_fis" "tf_fis" {
  name = "tf_fis"
}

output "ownernode" {
  value = data.citrixadc_fis.tf_fis.ownernode
}
```


## Argument Reference

* `name` - (Required) Name for the FIS to be created. Leading character must be a number or letter. Other characters allowed, after the first character, are @ _ - . (period) : (colon) # and space ( ). Note: In a cluster setup, the FIS name on each node must be unique.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `ownernode` - ID of the cluster node for which you are creating the FIS. Can be configured only through the cluster IP address.
* `id` - The id of the fis. It has the same value as the `name` attribute.
