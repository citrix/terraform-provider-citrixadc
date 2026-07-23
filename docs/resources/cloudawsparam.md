---
subcategory: "Cloud"
---

# Resource: cloudawsparam

This resource is used to manage the AWS cloud parameters.


## Example usage

```hcl
resource "citrixadc_cloudawsparam" "tf_cloudawsparam" {
  rolearn = "arn:aws:iam::123456789012:role/citrix-adc-role"
}
```


## Argument Reference

* `rolearn` - (Required) IAM Role ARN that the Citrix ADC assumes when accessing AWS services. Maximum length =  2048


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the cloudawsparam. It is set to `cloudawsparam-config`.


## Import

A cloudawsparam can be imported using its id (a fixed synthetic constant), e.g.

```shell
terraform import citrixadc_cloudawsparam.tf_cloudawsparam cloudawsparam-config
```
