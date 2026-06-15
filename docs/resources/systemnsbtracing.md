---
subcategory: "System"
---

# Resource: systemnsbtracing

Controls NSB (NetScaler Buffer) tracing on the Citrix ADC. This is a lifecycle-driven toggle: **creating this resource ENABLES NSB tracing, and destroying it DISABLES it.** There is no in-place on/off attribute to set — the desired state is expressed entirely by the presence or absence of the resource in your configuration. Updates are a no-op (the only configurable attribute, `nodeid`, forces replacement).

Under the hood, the ADC exposes only `?action=enable` and `?action=disable` (each with an empty payload) for this object; the provider maps Create to `enable` and Delete to `disable`.

This is a global singleton, so a single tracing toggle exists per appliance (or per cluster node when `nodeid` is specified).


## Example usage

```hcl
# Enabling NSB tracing on the appliance
resource "citrixadc_systemnsbtracing" "tf_systemnsbtracing" {}
```

```hcl
# Targeting a specific cluster node when reading state back
resource "citrixadc_systemnsbtracing" "tf_systemnsbtracing" {
  nodeid = 0
}
```

To disable NSB tracing, remove this resource from your configuration (or run `terraform destroy`); the provider issues the `disable` action on deletion.


## Argument Reference

* `nodeid` - (Optional) Unique number that identifies the cluster node. This is used only as a GET query filter to target a specific node when reading state; it is not part of the enable/disable action payload. Changing this value forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the systemnsbtracing resource. Because this is a global singleton, the ID is a synthetic constant string `systemnsbtracing-config`.
