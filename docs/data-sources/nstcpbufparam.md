---
subcategory: "NS"
---

# Data Source: nstcpbufparam

The nstcpbufparam data source allows you to retrieve information about TCP buffering parameters.


## Example usage

```terraform
data "citrixadc_nstcpbufparam" "my_nstcpbufparam" {
}

output "size" {
  value = data.citrixadc_nstcpbufparam.my_nstcpbufparam.size
}

output "memlimit" {
  value = data.citrixadc_nstcpbufparam.my_nstcpbufparam.memlimit
}
```


## Argument Reference

This data source takes no arguments.

## Attribute Reference

The following attributes are available:

* `size` - TCP buffering size per connection, in kilobytes.
* `memlimit` - Maximum memory, in megabytes, that can be used for buffering.
* `id` - The id of the nstcpbufparam resource.

### Read-only nstcpbufparam metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_nstcpbufparam` resource) and are computed. Any attribute the appliance does not return is `null`.

* `builtin` - Flag to determine if TCP buffering is built-in or not. A list of strings (possible values: MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL).
* `feature` - The feature to be checked while applying this config.
