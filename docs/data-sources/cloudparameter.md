---
subcategory: "Cloud"
---

# Data Source: cloudparameter

The cloudparameter data source allows you to retrieve information about the global cloud parameters configured on the Citrix ADC.


## Example usage

```hcl
data "citrixadc_cloudparameter" "tf_cloudparameter" {
}

output "cloud_controllerfqdn" {
  value = data.citrixadc_cloudparameter.tf_cloudparameter.controllerfqdn
}

output "cloud_customerid" {
  value = data.citrixadc_cloudparameter.tf_cloudparameter.customerid
}
```


## Argument Reference

This data source takes no arguments; it always reads the singleton `cloudparameter` configuration.


## Attribute Reference

The following attributes are available:

* `id` - The id of the cloudparameter. It is set to `cloudparameter-config`.
* `controllerfqdn` - FQDN of the controller to which the Citrix ADC SDProxy connects.
* `controllerport` - Port number of the controller to which the Citrix ADC SDProxy connects.
* `customerid` - Customer ID of the Citrix Cloud customer.
* `instanceid` - Instance ID of the customer provided by Trust.
* `resourcelocation` - Resource Location of the customer provided by Trust.
* `deployment` - Describes whether the customer is a Staging, Production, or Dev Citrix Cloud customer. Possible values: [ Production, Staging, Dev ]
* `connectorresidence` - Identifies where the connector is located. Possible values: [ None, Onprem, Aws, Azure, Cpx ]

~> **Note:** `activationcode` is a write-only field. The NITRO GET/show operation never returns it, so this data source does not expose it as a readable value (it is always null). Use the `citrixadc_cloudparameter` resource to configure the activation code.
