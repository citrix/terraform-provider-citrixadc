---
subcategory: "Load Balancing"
---

# Resource: lbwlm_lbvserver_binding

Associates a load balancing virtual server with a Work Load Manager (WLM) on the Citrix ADC so that the WLM can collect and report load metrics for the traffic handled by that virtual server. Creating this binding tells the named WLM which virtual server it should monitor.

~> **Note** The Work Load Manager (WLM) feature is deprecated on Citrix ADC. New deployments should not rely on it; this resource is provided for managing existing configurations.


## Example usage

```hcl
resource "citrixadc_lbwlm" "tf_lbwlm" {
  wlmname           = "mywlm"
  ipaddress         = "10.222.74.177"
  port              = 3010
  lbuid             = "lb-uid-001"
  katimeout         = 2
}

resource "citrixadc_lbvserver" "tf_lbvserver" {
  name        = "tf_lbvserver"
  servicetype = "HTTP"
  ipv46       = "10.10.10.33"
  port        = 80
}

resource "citrixadc_lbwlm_lbvserver_binding" "tf_lbwlm_lbvserver_binding" {
  wlmname     = citrixadc_lbwlm.tf_lbwlm.wlmname
  vservername = citrixadc_lbvserver.tf_lbvserver.name
}
```


## Argument Reference

* `wlmname` - (Required) The name of the Work Load Manager.
* `vservername` - (Required) Name of the virtual server which is to be bound to the WLM.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbwlm_lbvserver_binding. It is a composite identifier built from the `vservername` and `wlmname` attributes in the form `vservername:<vservername>,wlmname:<wlmname>` (the values are URL-encoded).


## Import

A lbwlm_lbvserver_binding can be imported using its id, in the format `vservername:<vservername>,wlmname:<wlmname>`, e.g.

```shell
terraform import citrixadc_lbwlm_lbvserver_binding.tf_lbwlm_lbvserver_binding vservername:tf_lbvserver,wlmname:mywlm
```
