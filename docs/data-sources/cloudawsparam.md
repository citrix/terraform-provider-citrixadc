---
subcategory: "Cloud"
---

# Data Source: cloudawsparam

The cloudawsparam data source allows you to retrieve information about the AWS integration parameters configured on the Citrix ADC.


## Example usage

```terraform
data "citrixadc_cloudawsparam" "tf_cloudawsparam" {
}

output "rolearn" {
  value = data.citrixadc_cloudawsparam.tf_cloudawsparam.rolearn
}
```


## Argument Reference

This datasource does not require any arguments.

## Attribute Reference

The following attributes are available:

* `rolearn` - IAM Role ARN that the Citrix ADC assumes when accessing AWS services.
* `id` - The id of the cloudawsparam.
