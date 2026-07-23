---
subcategory: "System"
---

# Resource: systemnsbtracing

This resource is used to manage NSB (NetScaler Buffer) tracing on the Citrix ADC.


## Example usage

```hcl
# Enabling NSB tracing on the appliance
resource "citrixadc_systemnsbtracing" "tf_systemnsbtracing" {}
```

```hcl
# Targeting a specific cluster node
resource "citrixadc_systemnsbtracing" "tf_systemnsbtracing" {
  nodeid = 0
}
```

To disable NSB tracing, remove this resource from your configuration (or run `terraform destroy`); the provider issues the `disable` action on deletion.


## Argument Reference

* `nodeid` - (Optional) Unique number that identifies the cluster node. Changing this value forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the systemnsbtracing. It is set to `systemnsbtracing-config`.


## Import

A systemnsbtracing can be imported using its id, e.g.

```shell
terraform import citrixadc_systemnsbtracing.tf_systemnsbtracing systemnsbtracing-config
```
