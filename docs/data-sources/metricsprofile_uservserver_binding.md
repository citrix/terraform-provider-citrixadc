---
subcategory: "Metrics"
---

# Data Source: metricsprofile_uservserver_binding

The metricsprofile_uservserver_binding data source allows you to retrieve information about an existing binding between a user-defined virtual server and a metrics profile on the Citrix ADC.


## Example usage

```terraform
data "citrixadc_metricsprofile_uservserver_binding" "example" {
  name       = "http_metrics_profile"
  entityname = "lb_user_vserver1"
  entitytype = "uservserver"
}

output "metricsprofile_uservserver_binding_id" {
  value = data.citrixadc_metricsprofile_uservserver_binding.example.id
}
```


## Argument Reference

* `name` - (Required) Name for the metrics profile to which the entity is bound.
* `entityname` - (Required) Name of the entity (user virtual server) bound to the metrics profile.
* `entitytype` - (Required) Type of the entity bound to the metrics profile. For this binding the value is `uservserver`.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the metricsprofile_uservserver_binding. It is a string of comma-separated `key:value` pairs in the form `entityname:<entityname>,entitytype:<entitytype>,name:<name>`.
