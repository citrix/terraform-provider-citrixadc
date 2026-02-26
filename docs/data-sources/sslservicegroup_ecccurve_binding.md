---
subcategory: "SSL"
---

# Data Source: sslservicegroup_ecccurve_binding

The sslservicegroup_ecccurve_binding data source allows you to retrieve information about the ECC curve bindings on an SSL service group.

## Example Usage

```terraform
data "citrixadc_sslservicegroup_ecccurve_binding" "tf_binding" {
  servicegroupname = "tf_servicegroup"
  ecccurvename     = "P_256"
}

output "servicegroupname" {
  value = data.citrixadc_sslservicegroup_ecccurve_binding.tf_binding.servicegroupname
}

output "ecccurvename" {
  value = data.citrixadc_sslservicegroup_ecccurve_binding.tf_binding.ecccurvename
}
```

## Argument Reference

* `servicegroupname` - (Required) The name of the SSL service group to which the SSL policy needs to be bound.
* `ecccurvename` - (Required) Named ECC curve bound to servicegroup.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslservicegroup_ecccurve_binding. It is a system-generated identifier.
