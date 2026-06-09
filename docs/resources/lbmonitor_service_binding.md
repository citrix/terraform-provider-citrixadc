---
subcategory: "Load Balancing"
---

# Resource: lbmonitor_service_binding

Binds an existing load balancing monitor to a service (or service group) so that the Citrix ADC continuously probes the bound member and tracks its health. Once bound, the monitor's health checks govern whether the service is considered UP or DOWN for load balancing decisions, and the configured weight influences how traffic is distributed among healthy members.

~> **No drift detection.** The NITRO API does not expose a get/get(all)/count endpoint for this binding. Applying the configuration creates the binding on the appliance, but the provider cannot read it back. The `Read` operation is a no-op that preserves the prior Terraform state unchanged, so out-of-band changes to the binding on the ADC are **not** detected. Because of this, all attributes are also marked as forcing replacement (`RequiresReplace`): any change recreates the binding rather than updating it in place.


## Example usage

```hcl
resource "citrixadc_lbmonitor" "tf_http_monitor" {
  monitorname = "http_monitor"
  type        = "HTTP"
  interval    = 5
  resptimeout = 3
}

resource "citrixadc_service" "tf_service" {
  name        = "web_service1"
  ip          = "10.10.10.20"
  servicetype = "HTTP"
  port        = 80
}

resource "citrixadc_lbmonitor_service_binding" "tf_binding" {
  monitorname = citrixadc_lbmonitor.tf_http_monitor.monitorname
  servicename = citrixadc_service.tf_service.name
  weight      = 50
}
```


## Argument Reference

The following arguments are supported:

* `monitorname` - (Required) Name of the monitor to bind. Changing this forces a new resource to be created.
* `servicename` - (Optional) Name of the service to bind the monitor to. Either `servicename` or `servicegroupname` must be specified. Changing this forces a new resource to be created.
* `servicegroupname` - (Optional) Name of the service group to bind the monitor to. Either `servicename` or `servicegroupname` must be specified. Changing this forces a new resource to be created.
* `weight` - (Optional) Weight to assign to the binding between the monitor and the service. Changing this forces a new resource to be created.
* `state` - (Optional) State of the monitor. The state setting for a monitor of a given type affects all monitors of that type. For example, if an HTTP monitor is enabled, all HTTP monitors on the appliance are (or remain) enabled; if disabled, all HTTP monitors on the appliance are disabled. Defaults to `"ENABLED"`. Changing this forces a new resource to be created.
* `dup_state` - (Optional) NITRO duplicate of `state`; mirrors the `state` attribute. Defaults to `"ENABLED"`. Changing this forces a new resource to be created.
* `dup_weight` - (Optional) NITRO duplicate of `weight`; mirrors the `weight` attribute. Defaults to `1`. Changing this forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the `lbmonitor_service_binding` resource. It is a comma-separated, URL-encoded composite key of the form `monitorname:<monitorname>,servicename:<servicename>` (and additionally `,servicegroupname:<servicegroupname>` when a service group is bound), for example `monitorname:http_monitor,servicename:web_service1`.


## Import

A `lbmonitor_service_binding` resource can be imported using its composite id (`monitorname:<monitorname>,servicename:<servicename>[,servicegroupname:<servicegroupname>]`), e.g.

```shell
terraform import citrixadc_lbmonitor_service_binding.tf_binding monitorname:http_monitor,servicename:web_service1
```
