---
subcategory: "Metrics"
---

# Resource: metricsprofile_service_binding

This resource is used to bind a service to a metrics profile.


## Example usage

```hcl
resource "citrixadc_metricsprofile_service_binding" "tf_metricsprofile_service_binding" {
  name       = "http_metrics_profile"
  entityname = "svc_web1"
  entitytype = "service"
}
```


## Argument Reference

* `name` - (Required) Name of the metrics profile to which the entity is bound. Must begin with an ASCII alphabetic or underscore (`_`) character, and must contain only ASCII alphanumeric, underscore, hash (`#`), period (`.`), space, colon (`:`), at (`@`), equals (`=`), and hyphen (`-`) characters. Changing this value forces a new resource to be created.
* `entityname` - (Required) Name of the entity (service) bound to the metrics profile. Changing this value forces a new resource to be created.
* `entitytype` - (Required) Type of the entity bound to the metrics profile. For this binding the value is the constant `service`. Changing this value forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the metricsprofile_service_binding. It is a composite of the unique attributes formatted as comma-separated `key:value` pairs, where each value is URL-encoded: `entityname:<entityname>,entitytype:<entitytype>,name:<name>` (for example, `entityname:svc_web1,entitytype:service,name:http_metrics_profile`).


## Import

A metricsprofile_service_binding can be imported using its id, which is the comma-separated `key:value` composite described above, e.g.

```shell
terraform import citrixadc_metricsprofile_service_binding.tf_metricsprofile_service_binding "entityname:svc_web1,entitytype:service,name:http_metrics_profile"
```
