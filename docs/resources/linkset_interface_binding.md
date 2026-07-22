---
subcategory: "Network"
---

# Resource: linkset\_interface\_binding

Binds an interface to an existing linkset on the Citrix ADC. A linkset groups several interfaces so the appliance treats them as a single logical entity for bridging and Layer 2 forwarding, which avoids forwarding loops and duplicate broadcast traffic in topologies where the ADC bridges between multiple physical links.


## Example usage

```hcl
resource "citrixadc_linkset" "tf_linkset" {
  id = "LS/1"
}

resource "citrixadc_linkset_interface_binding" "tf_bind" {
  linksetid = citrixadc_linkset.tf_linkset.id
  ifnum     = "1/3"
}
```


## Argument Reference

* `linksetid` - (Required) ID of the linkset to which to bind the interface, specified in `LS/x` notation (for example, `LS/1`). Changing this value forces a new resource to be created.
* `ifnum` - (Required) The interface to be bound to the linkset, specified in `C/U` notation (for example, `1/3`). Changing this value forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the linkset\_interface\_binding resource. It is a composite key of the form `id:<linksetid>,ifnum:<ifnum>`, where each value is URL-encoded (because linkset and interface identifiers contain `/`). For example, the linkset `LS/1` bound to interface `1/3` yields `id:LS%2F1,ifnum:1%2F3`.


## Import

A linkset\_interface\_binding can be imported using its `id`, e.g.

```shell
terraform import citrixadc_linkset_interface_binding.tf_bind "id:LS%2F1,ifnum:1%2F3"
```
