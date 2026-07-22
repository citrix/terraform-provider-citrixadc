---
subcategory: "Metrics"
---

# Data Source: metricsprofile_csvserver_binding

The `metricsprofile_csvserver_binding` data source allows you to retrieve information about a specific binding between a metrics profile and a content switching virtual server (csvserver) on the Citrix ADC.

## Example usage

```terraform
data "citrixadc_metricsprofile_csvserver_binding" "tf_metricsprofile_csvserver_binding" {
  name       = "http_metrics_profile"
  entityname = "cs_vserver1"
  entitytype = "csvserver"
}

output "entityname" {
  value = data.citrixadc_metricsprofile_csvserver_binding.tf_metricsprofile_csvserver_binding.entityname
}

output "entitytype" {
  value = data.citrixadc_metricsprofile_csvserver_binding.tf_metricsprofile_csvserver_binding.entitytype
}
```

## Argument Reference

* `name` - (Required) Name of the metrics profile to which the entity is bound.
* `entityname` - (Required) Name of the content switching virtual server bound to the metrics profile.
* `entitytype` - (Required) Type of the entity bound to the metrics profile. For this binding the value is the constant `csvserver`.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the metricsprofile_csvserver_binding. It is a composite identifier composed of comma-separated `key:value` pairs (each value URL-encoded), in the format `entityname:<entityname>,entitytype:<entitytype>,name:<name>`.
