---
subcategory: "Network"
---

# Resource: fis_interface_binding

This resource is used to bind an interface to a Failover Interface Set (FIS).


## Example usage

```hcl
resource "citrixadc_fis" "tf_fis" {
  name = "fis1"
}

resource "citrixadc_fis_interface_binding" "tf_fis_interface_binding" {
  name  = citrixadc_fis.tf_fis.name
  ifnum = "1/3"
}
```


## Argument Reference

* `name` - (Required) The name of the FIS to which you want to bind the interface. Changing this forces a new resource to be created.
* `ifnum` - (Required) Interface to be bound to the FIS, specified in slot/port notation (for example, `1/3`). Changing this forces a new resource to be created.
* `ownernode` - (Optional) ID of the cluster node for which you are creating the FIS. Can be configured only through the cluster IP address. Changing this forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the fis_interface_binding. It is a composite key of the form `name:<name>,ifnum:<ifnum>`, where each value is URL-encoded (for example, `name:fis1,ifnum:1%2F3`).
