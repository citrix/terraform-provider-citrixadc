---
subcategory: "Network"
---

# Resource: fis_interface_binding

Binds a physical interface to a Failover Interface Set (FIS) on the Citrix ADC. A FIS groups one or more interfaces (or channels) so that the failover state of the set as a whole tracks the state of its members. Use this resource to add an interface (for example, `1/3`) as a member of an existing FIS created with the `citrixadc_fis` resource.

~> **Note** This binding has no readable NITRO GET endpoint. The Citrix ADC exposes only add and delete operations for `fis_interface_binding`, and the aggregate `fis_binding/<name>` endpoint does not surface interface members (verified live on NS14.1). Terraform therefore manages the binding (create and delete) but cannot refresh it from the appliance: the plan/state values are authoritative and drift is not detected. The bound interfaces are visible on the CLI with `show fis <name>`.


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
