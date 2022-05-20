---
subcategory: "Basic"
---

# Resource: extendedmemoryparam

The extendedmemoryparam resource is used to create Parameter for extended memory used by LSN and Subscriber Store resource.


## Example usage

```hcl
# mamlimit should be less than Maximum Memory Usage Limit
resource "citrixadc_extendedmemoryparam" "tf_extendedmemoryparam" {
  memlimit = 512
}
```


## Argument Reference

* `memlimit` - (Required) Amount of Citrix ADC memory to reserve for the memory used by LSN and Subscriber Session Store feature, in multiples of 2MB. Note: If you later reduce the value of this parameter, the amount of active memory is not reduced. Changing the configured memory limit can only increase the amount of active memory.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the extendedmemoryparam. It is a unique string prefixed with "tf-extendedmemoryparam-"

