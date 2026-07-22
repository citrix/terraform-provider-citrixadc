---
subcategory: "Cloud"
---

# Resource: cloudawsparam

Configures the AWS integration parameters used by the Citrix ADC to authenticate against AWS services. Set the IAM Role ARN that the ADC assumes when it interacts with AWS (for example, to fetch instance metadata or access other AWS resources) so that cloud integrations work without embedding long-lived credentials on the appliance. This is a singleton resource: a single configuration object always exists on the ADC, so creating this resource updates the existing settings rather than adding a new object.


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
