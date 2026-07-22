---
subcategory: "Metrics"
---

# Resource: metricsprofile_crvserver_binding

Binds a Cache Redirection (CR) virtual server to a metrics profile, so that the time-series metrics collected by that profile (such as request rate, latency, and connection counters) are gathered for the bound CR vserver. Create one of these bindings for each CR vserver whose traffic you want a given metrics profile to track.


## Example usage

```hcl
resource "citrixadc_metricsprofile_crvserver_binding" "tf_metricsprofile_crvserver_binding" {
  name       = "http_metrics_profile"
  entityname = "cr_vsrv1"
  entitytype = "crvserver"
}
```


## Argument Reference

* `name` - (Required) Name for the metrics profile to which the entity is bound. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Changing this attribute forces a new resource to be created.
* `entityname` - (Required) Name of the entity (Cache Redirection virtual server) bound to the metrics profile. Changing this attribute forces a new resource to be created.
* `entitytype` - (Required) Type of the entity bound to the metrics profile. For this binding the value is the constant `crvserver`. Changing this attribute forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the metricsprofile_crvserver_binding. It is a comma-separated list of `key:value` pairs built from the unique attributes, with each value URL-encoded: `entityname:<entityname>,entitytype:<entitytype>,name:<name>`. For example: `entityname:cr_vsrv1,entitytype:crvserver,name:http_metrics_profile`.


## Import

A metricsprofile_crvserver_binding can be imported using its id (the comma-separated `key:value` form), e.g.

```shell
terraform import citrixadc_metricsprofile_crvserver_binding.tf_metricsprofile_crvserver_binding "entityname:cr_vsrv1,entitytype:crvserver,name:http_metrics_profile"
```
