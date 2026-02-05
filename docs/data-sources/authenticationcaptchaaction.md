---
subcategory: "Authentication"
---

# Data Source `authenticationcaptchaaction`

The authenticationcaptchaaction data source allows you to retrieve information about authentication captcha actions.


## Example usage

```terraform
data "citrixadc_authenticationcaptchaaction" "tf_captchaaction" {
  name = "my_captchaaction"
}

output "serverurl" {
  value = data.citrixadc_authenticationcaptchaaction.tf_captchaaction.serverurl
}

output "scorethreshold" {
  value = data.citrixadc_authenticationcaptchaaction.tf_captchaaction.scorethreshold
}
```


## Argument Reference

* `name` - (Required) Name for the new captcha action. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `defaultauthenticationgroup` - This is the group that is added to user sessions that match current policy.
* `scorethreshold` - This is the score threshold value for recaptcha v3.
* `secretkey` - Secret of gateway as established at the captcha source. (Note: This value is encrypted when retrieved from the API)
* `serverurl` - This is the endpoint at which captcha response is validated.
* `sitekey` - Sitekey to identify gateway fqdn while loading captcha. (Note: This value is encrypted when retrieved from the API)

## Attribute Reference

* `id` - The id of the authenticationcaptchaaction. It has the same value as the `name` attribute.


## Import

A authenticationcaptchaaction can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationcaptchaaction.tf_captchaaction my_captchaaction
```
