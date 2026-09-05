---
subcategory: "Network"
---

# Data Source: interfacepair

The interfacepair data source allows you to retrieve information about interface pairs configured on the Citrix ADC.


## Example usage

```terraform
data "citrixadc_interfacepair" "tf_interfacepair" {
  interface_id = 1
}

output "ifnum" {
  value = data.citrixadc_interfacepair.tf_interfacepair.ifnum
}
```


## Argument Reference

* `interface_id` - (Required) The Interface pair id.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `ifnum` - The constituent interfaces in the interface pair.
* `id` - The id of the interfacepair. It has the same value as the `interface_id` attribute.

### Read-only interfacepair metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_interfacepair` resource). They are Computed / GET-only. Any attribute the appliance does not return is `null`.

* `ifaces` - Names of all member interfaces of this Interface Pair.
