---
subcategory: "Authentication"
---

# Resource: authenticationemailaction

The authenticationemailaction resource is used to create authentication emailaction resource.


## Example usage

```hcl
resource "citrixadc_authenticationemailaction" "tf_emailaction" {
  name      = "tf_emailaction"
  username  = "username"
  password  = "secret"
  serverurl = "www/sdsd.com"
  timeout   = 100
  type      = "SMTP"
}
```


## Argument Reference

* `name` - (Required) Name for the new email action. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after an action is created.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my action" or 'my action').
* `username` - (Required) Username/Clientid/EmailID to be used to authenticate to the server.
* `password` - (Required) Password/Clientsecret to use when authenticating to the server.
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
