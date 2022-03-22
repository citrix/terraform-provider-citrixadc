---
subcategory: "Authentication"
---

# Resource: authenticationcaptchaaction

The authenticationcaptchaaction resource is used to create authentication captchaaction resource.


## Example usage

```hcl
resource "citrixadc_authenticationcaptchaaction" "tf_captchaaction" {
  name                       = "tf_captchaaction"
  secretkey                  = "secret"
  sitekey                    = "key"
  serverurl                  = "http://www.example.com/"
  defaultauthenticationgroup = "new_group"
}
```


## Argument Reference

* `name` - (Required) Name for the new captcha action. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after an action is created.  The following requirement applies only to the NetScaler CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my action" or 'my action').
* `secretkey` - (Required) Secret of gateway as established at the captcha source.
* `sitekey` - (Required) Sitekey to identify gateway fqdn while loading captcha.
* `serverurl` - (Optional) This is the endpoint at which captcha response is validated.
* `defaultauthenticationgroup` - (Optional) This is the group that is added to user sessions that match current policy.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authenticationcaptchaaction. It has the same value as the `name` attribute.


## Import

A authenticationcaptchaaction can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationcaptchaaction.tf_captchaaction tf_captchaaction
```
