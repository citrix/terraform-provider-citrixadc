---
subcategory: "Metrics"
---

# Resource: metricsprofile_uservserver_binding

This resource is used to bind a user vserver to a metrics profile.


## Example usage

```hcl
resource "citrixadc_metricsprofile_uservserver_binding" "tf_metricsprofile_uservserver_binding" {
  name       = "http_metrics_profile"
  entityname = "lb_user_vserver1"
  entitytype = "uservserver"
}
```


## Argument Reference

* `name` - (Required) Name for the metrics profile to which the entity is bound. Must begin with an ASCII alphabetic or underscore (`_`) character, and must contain only ASCII alphanumeric, underscore, hash (`#`), period (`.`), space, colon (`:`), at (`@`), equals (`=`), and hyphen (`-`) characters. Changing this value forces a new resource to be created.
* `entityname` - (Required) Name of the entity (user virtual server) bound to the metrics profile. Changing this value forces a new resource to be created.
* `entitytype` - (Required) Type of the entity bound to the metrics profile. For this binding the value is `uservserver`. Changing this value forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the metricsprofile_uservserver_binding. It is a string of comma-separated `key:value` pairs built from the unique attributes, in the form `entityname:<entityname>,entitytype:<entitytype>,name:<name>` (each value is URL-encoded). For example: `entityname:lb_user_vserver1,entitytype:uservserver,name:http_metrics_profile`.


## Import

A metricsprofile_uservserver_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_metricsprofile_uservserver_binding.tf_metricsprofile_uservserver_binding entityname:lb_user_vserver1,entitytype:uservserver,name:http_metrics_profile
```
