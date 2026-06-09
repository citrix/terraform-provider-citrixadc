---
subcategory: "Load Balancing"
---

# Data Source: lbwlm_lbvserver_binding

The lbwlm_lbvserver_binding data source allows you to retrieve information about a binding between a load balancing virtual server and a Work Load Manager (WLM) on the Citrix ADC.

~> **Note** The Work Load Manager (WLM) feature is deprecated on Citrix ADC.


## Example usage

```terraform
data "citrixadc_lbwlm_lbvserver_binding" "example" {
  wlmname     = "mywlm"
  vservername = "tf_lbvserver"
}

output "binding_id" {
  value = data.citrixadc_lbwlm_lbvserver_binding.example.id
}
```


## Argument Reference

* `wlmname` - (Required) The name of the Work Load Manager.
* `vservername` - (Required) Name of the virtual server which is bound to the WLM.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbwlm_lbvserver_binding. It is a composite identifier in the form `vservername:<vservername>,wlmname:<wlmname>` (the values are URL-encoded).
