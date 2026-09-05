---
subcategory: "AAA"
---

# Resource: aaaotpparameter

The aaaotpparameter> resource is used to update aaaotpparameter.


## Example usage

```hcl
resource "citrixadc_aaaotpparameter" "tf_aaaotpparameter" {
  encryption = "ON"
  maxotpdevices = 5
}
```


## Argument Reference

* `encryption` - (Optional) To encrypt otp secret in AD or not. Default value is OFF. Possible values: [ on, off ]
* `maxotpdevices` - (Optional) Maximum number of otp devices user can register. Default value is 4. Max value is 255. Minimum value =  0 Maximum value =  255
* `otptype` - (Optional) Input flag to generate OTP for the given type. Possible values = gwtest


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the aaaotpparameter. It is a unique string prefixed with `tf-aaaotpparameter-`.


## Import

A aaaotpparameter can be imported using its id, e.g.

```shell
terraform import citrixadc_aaaotpparameter.tf_aaaotpparameter aaaotpparameter-config
```
