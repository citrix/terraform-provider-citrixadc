---
subcategory: "Metrics"
---

# Data Source: metricsprofile_service_binding

The metricsprofile_service_binding data source allows you to retrieve information about a binding between a service entity and a metrics profile on the Citrix ADC.


## Example usage

```terraform
data "citrixadc_metricsprofile_service_binding" "example" {
  name       = "http_metrics_profile"
  entityname = "svc_web1"
  entitytype = "service"
}

output "binding_id" {
  value = data.citrixadc_metricsprofile_service_binding.example.id
}
```


## Argument Reference

* `name` - (Required) Name of the metrics profile to which the entity is bound.
* `entityname` - (Required) Name of the entity (service) bound to the metrics profile.
* `entitytype` - (Required) Type of the entity bound to the metrics profile. For this binding the value is the constant `service`.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the metricsprofile_service_binding. It is a composite of the unique attributes formatted as comma-separated `key:value` pairs: `entityname:<entityname>,entitytype:<entitytype>,name:<name>`.
