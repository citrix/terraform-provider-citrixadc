---
subcategory: "Metrics"
---

# Resource: metricsprofile_gslbvserver_binding

Binds a GSLB (Global Server Load Balancing) virtual server to a metrics profile, so that the time-series metrics collected by that profile (such as request rate, latency, and connection counters) are gathered for the bound GSLB vserver. Create one of these bindings for each GSLB vserver whose traffic you want a given metrics profile to track.


## Example usage

```hcl
resource "citrixadc_metricsprofile_gslbvserver_binding" "tf_metricsprofile_gslbvserver_binding" {
  name       = "http_metrics_profile"
  entityname = "gslb_vsrv1"
  entitytype = "gslbvserver"
}
```


## Argument Reference

* `name` - (Required) Name for the metrics profile to which the entity is bound. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Changing this attribute forces a new resource to be created.
* `entityname` - (Required) Name of the entity (GSLB virtual server) bound to the metrics profile. Changing this attribute forces a new resource to be created.
* `entitytype` - (Required) Type of the entity bound to the metrics profile. For this binding the value is the constant `gslbvserver`. Changing this attribute forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the metricsprofile_gslbvserver_binding. It is a comma-separated list of `key:value` pairs built from the unique attributes, with each value URL-encoded: `entityname:<entityname>,entitytype:<entitytype>,name:<name>`. For example: `entityname:gslb_vsrv1,entitytype:gslbvserver,name:http_metrics_profile`.


## Import

A metricsprofile_gslbvserver_binding can be imported using its id (the comma-separated `key:value` form), e.g.

```shell
terraform import citrixadc_metricsprofile_gslbvserver_binding.tf_metricsprofile_gslbvserver_binding "entityname:gslb_vsrv1,entitytype:gslbvserver,name:http_metrics_profile"
```
