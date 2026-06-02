---
subcategory: "Metrics"
---

# Resource: metricsprofile_csvserver_binding

Binds a content switching virtual server (csvserver) to a metrics profile on the Citrix ADC. Creating this binding tells the ADC to collect the metrics defined in the named metrics profile for the specified content switching virtual server, so its traffic statistics can be exported through the configured metrics collector.


## Example usage

```hcl
resource "citrixadc_metricsprofile_csvserver_binding" "tf_metricsprofile_csvserver_binding" {
  name       = "http_metrics_profile"
  entityname = "cs_vserver1"
  entitytype = "csvserver"
}
```


## Argument Reference

* `name` - (Required) Name of the metrics profile to which the entity is bound. Must begin with an ASCII alphabetic or underscore (`_`) character, and must contain only ASCII alphanumeric, underscore, hash (`#`), period (`.`), space, colon (`:`), at (`@`), equals (`=`), and hyphen (`-`) characters. Changing this value forces a new resource to be created.
* `entityname` - (Required) Name of the content switching virtual server bound to the metrics profile. Changing this value forces a new resource to be created.
* `entitytype` - (Required) Type of the entity bound to the metrics profile. For this binding the value is the constant `csvserver`. Changing this value forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the metricsprofile_csvserver_binding. It is a composite identifier composed of comma-separated `key:value` pairs (each value URL-encoded), in the format `entityname:<entityname>,entitytype:<entitytype>,name:<name>`.


## Import

A metricsprofile_csvserver_binding can be imported using its id, which is the composite `key:value` string described above, e.g.

```shell
terraform import citrixadc_metricsprofile_csvserver_binding.tf_metricsprofile_csvserver_binding entityname:cs_vserver1,entitytype:csvserver,name:http_metrics_profile
```
