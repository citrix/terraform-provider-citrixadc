---
subcategory: "NS"
---

# Resource: nsconfig_save

The nsconfig_save resource is used to apply the save operation for ns config.


## Example usage

```hcl
resource "citrixadc_nsconfig_save" "tf_ns_save" {
    all        = true
    timestamp  = "2020-03-24T12:37:06Z"
}
```


## Argument Reference

* `all` - (Optional) Use this option to do saveconfig for all partitions.
* `timestamp` - (Required) the timestamp of the operation. Can be any string. Used to force the operation again if all other attributes have remained the same.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nsconfig_save. It has the same value as the `timestamp` attribute.
