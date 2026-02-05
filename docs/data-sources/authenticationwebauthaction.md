---
subcategory: "Authentication"
---

# Data Source `authenticationwebauthaction`

The authenticationwebauthaction data source allows you to retrieve information about Web Authentication actions.


## Example usage

```terraform
data "citrixadc_authenticationwebauthaction" "tf_webauthaction" {
  name = "my_webauthaction"
}

output "serverip" {
  value = data.citrixadc_authenticationwebauthaction.tf_webauthaction.serverip
}

output "scheme" {
  value = data.citrixadc_authenticationwebauthaction.tf_webauthaction.scheme
}
```


## Argument Reference

* `name` - (Required) Name for the Web Authentication action.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `attribute1` - Expression that would be evaluated to extract attribute1 from the webauth response.
* `attribute2` - Expression that would be evaluated to extract attribute2 from the webauth response.
* `attribute3` - Expression that would be evaluated to extract attribute3 from the webauth response.
* `attribute4` - Expression that would be evaluated to extract attribute4 from the webauth response.
* `attribute5` - Expression that would be evaluated to extract attribute5 from the webauth response.
* `attribute6` - Expression that would be evaluated to extract attribute6 from the webauth response.
* `attribute7` - Expression that would be evaluated to extract attribute7 from the webauth response.
* `attribute8` - Expression that would be evaluated to extract attribute8 from the webauth response.
* `attribute9` - Expression that would be evaluated to extract attribute9 from the webauth response.
* `attribute10` - Expression that would be evaluated to extract attribute10 from the webauth response.
* `attribute11` - Expression that would be evaluated to extract attribute11 from the webauth response.
* `attribute12` - Expression that would be evaluated to extract attribute12 from the webauth response.
* `attribute13` - Expression that would be evaluated to extract attribute13 from the webauth response.
* `attribute14` - Expression that would be evaluated to extract attribute14 from the webauth response.
* `attribute15` - Expression that would be evaluated to extract attribute15 from the webauth response.
* `attribute16` - Expression that would be evaluated to extract attribute16 from the webauth response.
* `defaultauthenticationgroup` - This is the default group that is chosen when the authentication succeeds in addition to extracted groups.
* `fullreqexpr` - Exact HTTP request, in the form of an expression, which the Citrix ADC sends to the authentication server. The Citrix ADC does not check the validity of this request. One must manually validate the request.
* `scheme` - Type of scheme for the web server.
* `serverip` - IP address of the web server to be used for authentication.
* `serverport` - Port on which the web server accepts connections.
* `successrule` - Expression, that checks to see if authentication is successful.
