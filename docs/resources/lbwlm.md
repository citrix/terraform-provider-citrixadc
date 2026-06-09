---
subcategory: "Load Balancing"
---

# Resource: lbwlm

Registers a Web/Workload Manager (WLM) with the Citrix ADC so that load-balancing decisions can be delegated to an external Content Application Server (CAS). The WLM communicates with the ADC over a configured IP address and port, and the ADC periodically probes it using a keep-alive timeout. Create an `lbwlm` resource when you need the ADC to consult an external workload manager for server-selection guidance.

> **Note:** WLM (Web/Workload Manager, also referred to as CAS) is a deprecated NetScaler/Citrix ADC feature. Use this resource only when interoperating with an existing WLM/CAS deployment.


## Example usage

```hcl
resource "citrixadc_lbwlm" "example" {
  wlmname   = "wlm1"
  lbuid     = "LB001"
  ipaddress = "192.168.10.50"
  port      = 3013
  katimeout = 2
}
```


## Argument Reference

The following arguments are supported:

* `wlmname` - (Required) The name of the Work Load Manager. This value is also used as the resource `id`. Changing this attribute forces the resource to be recreated.
* `lbuid` - (Required) The LBUID for the Load Balancer to communicate to the Work Load Manager. This attribute is set only at creation time; changing it forces the resource to be recreated.
* `ipaddress` - (Optional) The IP address of the WLM. This attribute is set only at creation time; changing it forces the resource to be recreated.
* `port` - (Optional) The port of the WLM. This attribute is set only at creation time; changing it forces the resource to be recreated.
* `katimeout` - (Optional) The idle time period after which Citrix ADC would probe the WLM. The value ranges from 1 to 1440 minutes. Defaults to `2`. This is the only attribute that can be updated in place; changing it does not force recreation.

### Create-only versus updateable attributes

Only `katimeout` can be modified after creation and is applied via an in-place update. The `wlmname`, `lbuid`, `ipaddress`, and `port` attributes are create-only — changing any of them forces Terraform to destroy and recreate the resource.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the lbwlm resource. It has the same value as the `wlmname` attribute.


## Import

An lbwlm resource can be imported using its `wlmname`, e.g.

```shell
terraform import citrixadc_lbwlm.example wlm1
```
