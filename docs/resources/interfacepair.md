---
subcategory: "Network"
---

# Resource: interfacepair

The interfacepair resource is used to create interfacepair.


## Example usage

```hcl
resource "citrixadc_interfacepair" "tf_interfacepair" {
  interface_id = 2
  ifnum        = ["LA/2", "LA/3"]
}

```


## Argument Reference

* `interface_id` - (Required) The Interface pair id. Minimum value =  1 Maximum value =  255
* `ifnum` - (Required) The constituent interfaces in the interface pair. Minimum length =  1


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the interfacepair. It has the same value as the `interface_id` attribute.
