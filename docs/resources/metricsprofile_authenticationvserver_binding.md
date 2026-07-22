---
subcategory: "Metrics"
---

# Resource: metricsprofile_authenticationvserver_binding

The metricsprofile_authenticationvserver_binding resource binds an authentication virtual server to a metrics profile on the Citrix ADC. Creating this binding makes the ADC collect and export the metrics defined by the referenced metrics profile for the specified authentication vserver, allowing you to monitor that vserver through your metrics pipeline.


## Example usage

```hcl
resource "citrixadc_metricsprofile_authenticationvserver_binding" "tf_binding" {
  name       = "http_metrics_profile"
  entityname = "auth_vserver1"
  entitytype = "authvserver"
}
```


## Argument Reference

* `name` - (Required) Name for the metrics profile to which the entity is bound. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Changing this attribute forces a new resource to be created.
* `entityname` - (Required) Name of the authentication virtual server to bind to the metrics profile. Changing this attribute forces a new resource to be created.
* `entitytype` - (Required) Type of the entity bound to the metrics profile. For this binding the value is always `authvserver`. Changing this attribute forces a new resource to be created.

This binding is immutable. Any change results in the binding being destroyed and recreated.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the metricsprofile_authenticationvserver_binding. It is a composite identifier built from the `entityname`, `entitytype`, and `name` attributes as comma-separated `key:value` pairs (for example, `entityname:auth_vserver1,entitytype:authvserver,name:http_metrics_profile`).


## Import

A metricsprofile_authenticationvserver_binding can be imported using its id, which is the composite of `name`, `entityname`, and `entitytype`. Provide the value in the same comma-separated `key:value` form used for the `id` attribute, e.g.

```shell
terraform import citrixadc_metricsprofile_authenticationvserver_binding.tf_binding "entityname:auth_vserver1,entitytype:authvserver,name:http_metrics_profile"
```
