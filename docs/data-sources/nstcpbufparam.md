---
subcategory: "NS"
---

# Data Source `nstcpbufparam`

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
