---
subcategory: "AAA"
---

# Data Source `aaaotpparameter`

The aaaotpparameter data source allows you to retrieve information about AAA OTP (One-Time Password) parameters configuration.


## Example usage

```terraform
data "citrixadc_aaaotpparameter" "tf_aaaotpparameter" {
}

output "encryption" {
  value = data.citrixadc_aaaotpparameter.tf_aaaotpparameter.encryption
}

output "maxotpdevices" {
  value = data.citrixadc_aaaotpparameter.tf_aaaotpparameter.maxotpdevices
}
```


## Argument Reference

This datasource does not require any arguments.

## Attribute Reference

The following attributes are available:

* `encryption` - To encrypt otp secret in AD or not. Default value is OFF.
* `maxotpdevices` - Maximum number of otp devices user can register. Default value is 4. Max value is 255.

## Attribute Reference

* `id` - The id of the aaaotpparameter. It is a system-generated identifier.
