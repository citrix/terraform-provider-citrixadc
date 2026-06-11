---
subcategory: "Network"
---

# Data Source: linkset\_interface\_binding

The linkset\_interface\_binding data source allows you to retrieve information about an interface bound to a linkset on the Citrix ADC.


## Example usage

```hcl
data "citrixadc_linkset_interface_binding" "tf_bind" {
  linksetid = "LS/1"
  ifnum     = "1/3"
}

output "bound_interface" {
  value = data.citrixadc_linkset_interface_binding.tf_bind.ifnum
}
```


## Argument Reference

* `linksetid` - (Required) ID of the linkset whose interface binding you want to look up, specified in `LS/x` notation (for example, `LS/1`).
* `ifnum` - (Required) The interface bound to the linkset, specified in `C/U` notation (for example, `1/3`).


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the linkset\_interface\_binding resource. It is a composite key of the form `id:<linksetid>,ifnum:<ifnum>`, with each value URL-encoded.
