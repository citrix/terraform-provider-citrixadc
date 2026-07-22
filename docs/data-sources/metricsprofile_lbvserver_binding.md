---
subcategory: "Metrics"
---

# Data Source: metricsprofile_lbvserver_binding

The metricsprofile_lbvserver_binding data source allows you to retrieve information about an existing binding between a load balancing virtual server and a metrics profile.


## Example usage

```terraform
data "citrixadc_metricsprofile_lbvserver_binding" "example" {
  name       = "http_metrics"
  entityname = "lb_vserver1"
  entitytype = "lbvserver"
}

output "bound_entity" {
  value = data.citrixadc_metricsprofile_lbvserver_binding.example.entityname
}
```


## Argument Reference

* `name` - (Required) Name of the metrics profile to which the entity is bound.
* `entityname` - (Required) Name of the entity (load balancing virtual server) bound to the metrics profile.
* `entitytype` - (Required) Type of the entity bound to the metrics profile. For this binding the value is always `lbvserver`.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the `metricsprofile_lbvserver_binding` resource. It is a composite of the unique attributes expressed as comma-separated `key:value` pairs, in the form `entityname:<entityname>,entitytype:<entitytype>,name:<name>`.
* `name` - Name of the metrics profile to which the entity is bound.
* `entityname` - Name of the entity bound to the metrics profile.
* `entitytype` - Type of the entity bound to the metrics profile.
