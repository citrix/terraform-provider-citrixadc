---
subcategory: "DPS"
---

# Data Source: dpsparameter

The dpsparameter data source allows you to retrieve information about the DPS
(Citrix Cloud connectivity) parameters configuration.


## Example usage

```terraform
data "citrixadc_dpsparameter" "tf_dpsparameter" {
}

output "customerid" {
  value = data.citrixadc_dpsparameter.tf_dpsparameter.customerid
}

output "deployment" {
  value = data.citrixadc_dpsparameter.tf_dpsparameter.deployment
}
```


## Argument Reference

This datasource does not require any arguments.

## Attribute Reference

The following attributes are available:

* `customerid` - Customer ID of the Citrix Cloud customer.
* `deployment` - Describes if the customer is connecting to Commerical/JapanCloud/Gov Citrix Cloud customer. Possible values: [ COMMERCIAL, GOV, JAPANCLOUD ]
* `serviceurl` - Service URL of the Citrix Cloud customer.
* `id` - The id of the dpsparameter. It is a system-generated identifier.
