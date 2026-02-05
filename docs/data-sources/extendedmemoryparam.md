---
subcategory: "System"
---

# Data Source `extendedmemoryparam`

The extendedmemoryparam data source allows you to retrieve information about the extended memory parameters configuration.


## Example usage

```terraform
data "citrixadc_extendedmemoryparam" "tf_extendedmemoryparam" {
}

output "memlimit" {
  value = data.citrixadc_extendedmemoryparam.tf_extendedmemoryparam.memlimit
}
```


## Argument Reference

No required arguments. This data source retrieves the global extended memory parameter configuration.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `memlimit` - Amount of NetScaler memory to reserve for the memory used by LSN and Subscriber Session Store feature, in multiples of 2MB. Note: If you later reduce the value of this parameter, the amount of active memory is not reduced. Changing the configured memory limit can only increase the amount of active memory.

## Attribute Reference

* `id` - The id of the extendedmemoryparam.
