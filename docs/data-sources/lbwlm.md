---
subcategory: "Load Balancing"
---

# Data Source: lbwlm

The lbwlm data source allows you to retrieve information about a configured Web/Workload Manager (WLM) registration on the Citrix ADC, looked up by its name.

> **Note:** WLM (Web/Workload Manager, also referred to as CAS) is a deprecated NetScaler/Citrix ADC feature.


## Example usage

```terraform
data "citrixadc_lbwlm" "example" {
  wlmname = "wlm1"
}

output "wlm_ipaddress" {
  value = data.citrixadc_lbwlm.example.ipaddress
}
```


## Argument Reference

* `wlmname` - (Required) The name of the Work Load Manager.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the lbwlm resource. It has the same value as the `wlmname` attribute.
* `ipaddress` - The IP address of the WLM.
* `port` - The port of the WLM.
* `lbuid` - The LBUID for the Load Balancer to communicate to the Work Load Manager.
* `katimeout` - The idle time period after which Citrix ADC would probe the WLM. The value ranges from 1 to 1440 minutes.
