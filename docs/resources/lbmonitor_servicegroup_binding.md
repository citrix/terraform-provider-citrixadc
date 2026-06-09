---
subcategory: "Load Balancing"
---

# Resource: lbmonitor_servicegroup_binding

Binds an existing load balancing monitor to a service group so that the Citrix ADC continuously probes the members of that group and tracks their health. Once bound, the monitor's health checks govern whether the service group members are considered UP or DOWN for load balancing decisions, and the configured weight influences how traffic is distributed among healthy members.

~> **No drift detection.** The NITRO API does not expose a get/get(all)/count endpoint for this binding. Applying the configuration creates the binding on the appliance, but the provider cannot read it back. The `Read` operation is a no-op that preserves the prior Terraform state unchanged, so out-of-band changes to the binding on the ADC are **not** detected. Because of this, all attributes are also marked as forcing replacement (`RequiresReplace`): any change recreates the binding rather than updating it in place.


## Example usage

```hcl
resource "citrixadc_lbmonitor" "tf_http_monitor" {
  monitorname = "http_monitor"
  type        = "HTTP"
  interval    = 5
  resptimeout = 3
}

resource "citrixadc_servicegroup" "tf_servicegroup" {
  servicegroupname = "web_servicegroup1"
  servicetype      = "HTTP"
}

resource "citrixadc_lbmonitor_servicegroup_binding" "tf_binding" {
  monitorname      = citrixadc_lbmonitor.tf_http_monitor.monitorname
  servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
  weight           = 50
}
```


## Argument Reference

The following arguments are supported:

* `monitorname` - (Required) Name of the monitor to bind. Changing this forces a new resource to be created.
* `servicegroupname` - (Required) Name of the service group to bind the monitor to. Changing this forces a new resource to be created.
* `servicename` - (Optional) Name of the service or service group within the binding. Changing this forces a new resource to be created.
* `weight` - (Optional) Weight to assign to the binding between the monitor and the service group. Changing this forces a new resource to be created.
* `state` - (Optional) State of the monitor. The state setting for a monitor of a given type affects all monitors of that type. For example, if an HTTP monitor is enabled, all HTTP monitors on the appliance are (or remain) enabled; if disabled, all HTTP monitors on the appliance are disabled. Defaults to `"ENABLED"`. Changing this forces a new resource to be created.
* `dup_state` - (Optional) NITRO duplicate of `state`; mirrors the `state` attribute. Defaults to `"ENABLED"`. Changing this forces a new resource to be created.
* `dup_weight` - (Optional) NITRO duplicate of `weight`; mirrors the `weight` attribute. Defaults to `1`. Changing this forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the `lbmonitor_servicegroup_binding` resource. It is a comma-separated, URL-encoded composite key of the form `monitorname:<monitorname>,servicegroupname:<servicegroupname>` (and additionally `,servicename:<servicename>` when a service name is specified), for example `monitorname:http_monitor,servicegroupname:web_servicegroup1`.


## Import

A `lbmonitor_servicegroup_binding` resource can be imported using its composite id (`monitorname:<monitorname>,servicegroupname:<servicegroupname>[,servicename:<servicename>]`), e.g.

```shell
terraform import citrixadc_lbmonitor_servicegroup_binding.tf_binding monitorname:http_monitor,servicegroupname:web_servicegroup1
```
