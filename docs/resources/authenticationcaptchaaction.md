---
subcategory: "Authentication"
---

# Resource: authenticationcaptchaaction

The authenticationcaptchaaction resource is used to create authentication captchaaction resource.


## Example usage

### Using secretkey and sitekey (sensitive attributes - persisted in state)

```hcl
variable "captchaaction_secretkey" {
  type      = string
  sensitive = true
}

variable "captchaaction_sitekey" {
  type      = string
  sensitive = true
}

resource "citrixadc_authenticationcaptchaaction" "tf_captchaaction" {
  name                       = "tf_captchaaction"
  secretkey                  = var.captchaaction_secretkey
  sitekey                    = var.captchaaction_sitekey
  serverurl                  = "http://www.example.com/"
  defaultauthenticationgroup = "new_group"
}
```

### Using secretkey_wo and sitekey_wo (write-only/ephemeral - NOT persisted in state)

The `secretkey_wo` and `sitekey_wo` attributes provide an ephemeral path for the captcha secret key and site key. The values are sent to the ADC but are **not stored in Terraform state**, reducing the risk of secret exposure. To trigger an update when the value changes, increment `secretkey_wo_version` or `sitekey_wo_version`.

```hcl
variable "captchaaction_secretkey" {
  type      = string
  sensitive = true
}

variable "captchaaction_sitekey" {
  type      = string
  sensitive = true
}

resource "citrixadc_authenticationcaptchaaction" "tf_captchaaction" {
  name                       = "tf_captchaaction"
  secretkey_wo               = var.captchaaction_secretkey
  secretkey_wo_version       = 1
  sitekey_wo                 = var.captchaaction_sitekey
  sitekey_wo_version         = 1
  serverurl                  = "http://www.example.com/"
  defaultauthenticationgroup = "new_group"
}
```

To rotate the secrets, update the variable values and bump the versions:

```hcl
resource "citrixadc_authenticationcaptchaaction" "tf_captchaaction" {
  name                       = "tf_captchaaction"
  secretkey_wo               = var.captchaaction_secretkey
  secretkey_wo_version       = 2  # Bumped to trigger update
  sitekey_wo                 = var.captchaaction_sitekey
  sitekey_wo_version         = 2  # Bumped to trigger update
  serverurl                  = "http://www.example.com/"
  defaultauthenticationgroup = "new_group"
}
```


## Argument Reference

* `name` - (Required) Name for the new captcha action. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after an action is created.  The following requirement applies only to the NetScaler CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my action" or 'my action').
* `secretkey` - (Optional, Sensitive) Secret of gateway as established at the captcha source. The value is persisted in Terraform state (encrypted). See also `secretkey_wo` for an ephemeral alternative. At least one of `secretkey` or `secretkey_wo` must be set.
* `secretkey_wo` - (Optional, Sensitive, WriteOnly) Same as `secretkey`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `secretkey_wo_version`. If both `secretkey` and `secretkey_wo` are set, `secretkey_wo` takes precedence. At least one of `secretkey` or `secretkey_wo` must be set.
* `secretkey_wo_version` - (Optional) An integer version tracker for `secretkey_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the value has changed and trigger an update. Defaults to `1`.
* `sitekey` - (Optional, Sensitive) Sitekey to identify gateway fqdn while loading captcha. The value is persisted in Terraform state (encrypted). See also `sitekey_wo` for an ephemeral alternative.
* `sitekey_wo` - (Optional, Sensitive, WriteOnly) Same as `sitekey`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `sitekey_wo_version`. If both `sitekey` and `sitekey_wo` are set, `sitekey_wo` takes precedence.
* `sitekey_wo_version` - (Optional) An integer version tracker for `sitekey_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the value has changed and trigger an update. Defaults to `1`.
* `serverurl` - (Optional) This is the endpoint at which captcha response is validated.
* `defaultauthenticationgroup` - (Optional) This is the group that is added to user sessions that match current policy.
* `scorethreshold` - (Optional) This is the score threshold value for recaptcha v3.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authenticationcaptchaaction. It has the same value as the `name` attribute.


## Import

A authenticationcaptchaaction can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationcaptchaaction.tf_captchaaction tf_captchaaction
```
