---
subcategory: "Authentication"
---

# Resource: authenticationemailaction

The authenticationemailaction resource is used to create authentication emailaction resource.


## Example usage

### Using password (sensitive attribute - persisted in state)

```hcl
variable "authenticationemailaction_password" {
  type      = string
  sensitive = true
}

resource "citrixadc_authenticationemailaction" "tf_emailaction" {
  name      = "tf_emailaction"
  username  = "username"
  password  = var.authenticationemailaction_password
  serverurl = "www/sdsd.com"
  timeout   = 100
  type      = "SMTP"
}
```

### Using password_wo (write-only/ephemeral - NOT persisted in state)

The `password_wo` attribute provides an ephemeral path for the password. The value is sent to the ADC but is **not stored in Terraform state**, reducing the risk of secret exposure. To trigger an update when the value changes, increment `password_wo_version`.

```hcl
variable "authenticationemailaction_password" {
  type      = string
  sensitive = true
}

resource "citrixadc_authenticationemailaction" "tf_emailaction" {
  name                = "tf_emailaction"
  username            = "username"
  password_wo         = var.authenticationemailaction_password
  password_wo_version = 1
  serverurl           = "www/sdsd.com"
  timeout             = 100
  type                = "SMTP"
}
```

To rotate the secret, update the variable value and bump the version:

```hcl
resource "citrixadc_authenticationemailaction" "tf_emailaction" {
  name                = "tf_emailaction"
  username            = "username"
  password_wo         = var.authenticationemailaction_password
  password_wo_version = 2  # Bumped to trigger update
  serverurl           = "www/sdsd.com"
  timeout             = 100
  type                = "SMTP"
}
```


## Argument Reference

* `name` - (Required) Name for the new email action. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after an action is created.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my action" or 'my action').
* `username` - (Required) Username/Clientid/EmailID to be used to authenticate to the server.
* `password` - (Optional, Sensitive) Password/Clientsecret to use when authenticating to the server. The value is persisted in Terraform state (encrypted). See also `password_wo` for an ephemeral alternative. At least one of `password` or `password_wo` must be set.
* `password_wo` - (Optional, Sensitive, WriteOnly) Same as `password`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `password_wo_version`. If both `password` and `password_wo` are set, `password_wo` takes precedence. At least one of `password` or `password_wo` must be set.
* `password_wo_version` - (Optional) An integer version tracker for `password_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the value has changed and trigger an update. Defaults to `1`.
* `serverurl` - (Required) Address of the server that delivers the message. It is fully qualified fqdn such as http(s):// or smtp(s):// for http and smtp protocols respectively. For SMTP, the port number is mandatory like smtps://smtp.example.com:25.
* `content` - (Optional) Content to be delivered to the user. "$code" string within the content will be replaced with the actual one-time-code to be sent.
* `defaultauthenticationgroup` - (Optional) This is the group that is added to user sessions that match current IdP policy. It can be used in policies to identify relying party trust.
* `emailaddress` - (Optional) An optional expression that yields user's email. When not configured, user's default mail address would be used. When configured, result of this expression is used as destination email address.
* `timeout` - (Optional) Time after which the code expires.
* `type` - (Optional) Type of the email action. Default type is SMTP.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authenticationemailaction. It has the same value as the `name` attribute.


## Import

A authenticationemailaction can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationemailaction.tf_emailaction tf_emailaction
```
