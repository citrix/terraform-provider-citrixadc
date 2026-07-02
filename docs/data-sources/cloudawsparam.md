---
subcategory: "Cloud"
---

# Data Source: cloudawsparam

The cloudawsparam data source allows you to retrieve the AWS integration parameters configured on the Citrix ADC, such as the IAM Role ARN that the appliance assumes when accessing AWS services. This is a singleton configuration, so no lookup argument is required.


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
* `id` - The id of the cloudawsparam. It is a system-generated identifier.
