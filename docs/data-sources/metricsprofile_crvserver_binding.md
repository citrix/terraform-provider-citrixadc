---
subcategory: "Metrics"
---

# Data Source: metricsprofile_crvserver_binding

The metricsprofile_crvserver_binding data source allows you to retrieve information about an existing binding between a Cache Redirection (CR) virtual server and a metrics profile.


## Example usage

```terraform
data "citrixadc_metricsprofile_crvserver_binding" "example" {
  name       = "http_metrics_profile"
  entityname = "cr_vsrv1"
  entitytype = "crvserver"
}

output "bound_entity" {
  value = data.citrixadc_metricsprofile_crvserver_binding.example.entityname
}
```


## Argument Reference

* `name` - (Required) Name for the metrics profile to which the entity is bound.
* `entityname` - (Required) Name of the entity (Cache Redirection virtual server) bound to the metrics profile.
* `entitytype` - (Required) Type of the entity bound to the metrics profile. For this binding the value is the constant `crvserver`.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the metricsprofile_crvserver_binding. It is a comma-separated list of `key:value` pairs built from the unique attributes, with each value URL-encoded: `entityname:<entityname>,entitytype:<entitytype>,name:<name>`.
