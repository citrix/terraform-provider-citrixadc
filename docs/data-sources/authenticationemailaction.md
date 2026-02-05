---
subcategory: "Authentication"
---

# Data Source `authenticationemailaction`

The authenticationemailaction data source allows you to retrieve information about authentication email actions.


## Example usage

```terraform
data "citrixadc_authenticationemailaction" "tf_emailaction" {
  name = "tf_emailaction_ds"
}

output "username" {
  value = data.citrixadc_authenticationemailaction.tf_emailaction.username
}

output "serverurl" {
  value = data.citrixadc_authenticationemailaction.tf_emailaction.serverurl
}

output "type" {
  value = data.citrixadc_authenticationemailaction.tf_emailaction.type
}
```


## Argument Reference

* `name` - (Required) Name for the new email action. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after an action is created.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authenticationemailaction. It has the same value as the `name` attribute.
* `content` - Content to be delivered to the user. "$code" string within the content will be replaced with the actual one-time-code to be sent.
* `defaultauthenticationgroup` - This is the group that is added to user sessions that match current IdP policy. It can be used in policies to identify relying party trust.
* `emailaddress` - An optional expression that yields user's email. When not configured, user's default mail address would be used. When configured, result of this expression is used as destination email address.
* `password` - Password/Clientsecret to use when authenticating to the server.
* `serverurl` - Address of the server that delivers the message. It is fully qualified fqdn such as http(s):// or smtp(s):// for http and smtp protocols respectively. For SMTP, the port number is mandatory like smtps://smtp.example.com:25.
* `timeout` - Time after which the code expires.
* `type` - Type of the email action. Default type is SMTP.
* `username` - Username/Clientid/EmailID to be used to authenticate to the server.
