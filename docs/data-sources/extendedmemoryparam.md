---
subcategory: "System"
---

# Data Source: extendedmemoryparam

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
* `id` - The id of the extendedmemoryparam.

### Read-only extendedmemoryparam metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_extendedmemoryparam` resource). They are Computed / GET-only, and any attribute the appliance does not return is `null`.

* `memlimitactive` - The active memory limit for extendedmemory on the system. Active memory limit could be different from configured memory limit. This could happen when memory limit could not be increased due to unavailability, or could not be decreased as it is already in use. This active memory limit configures the current memory limit for LSN and Subscriber Session Store.
* `maxmemlimit` - The maximum value of memory limit for extendedmemory on the system. Actual available memory may be less. This is maximum memory that can be utilized by LSN and Subscriber Session Store modules.
* `minrequiredmemory` - The minimum memory requirement for extendedmemory. This is minimum memory required for LSN and Subscriber Session Store Modules.
