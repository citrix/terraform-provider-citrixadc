---
subcategory: "NS"
---

# Resource: nstcpbufparam

The nstcpbufparam resource is used to create tcp buffer parameter resource.


## Example usage

```hcl
resource "citrixadc_nstcpbufparam" "tf_nstcpbufparam" {
  size     = 64
  memlimit = 16
}
```


## Argument Reference

* `size` - (Optional) TCP buffering size per connection, in kilobytes. Minimum value =  4 Maximum value =  20480
* `memlimit` - (Optional) Maximum memory, in megabytes, that can be used for buffering.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nstcpbufparam. It is a unique string prefixed with "tf-nstcpbufparam-"

