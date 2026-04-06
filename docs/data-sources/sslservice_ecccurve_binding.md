---
subcategory: "SSL"
---

# Data Source: sslservice_ecccurve_binding

The sslservice_ecccurve_binding data source allows you to retrieve information about the ECC curve bindings on an SSL service.

## Example Usage

```terraform
data "citrixadc_sslservice_ecccurve_binding" "tf_binding" {
  servicename  = "tf_service"
  ecccurvename = "P_256"
}

output "servicename" {
  value = data.citrixadc_sslservice_ecccurve_binding.tf_binding.servicename
}

output "ecccurvename" {
  value = data.citrixadc_sslservice_ecccurve_binding.tf_binding.ecccurvename
}
```

## Argument Reference

* `servicename` - (Required) Name of the SSL service for which to set advanced configuration.
* `ecccurvename` - (Required) Named ECC curve bound to service/vserver.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslservice_ecccurve_binding. It is a system-generated identifier.
