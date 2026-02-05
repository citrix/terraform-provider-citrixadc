---
subcategory: "Authentication"
---

# Data Source `authenticationoauthaction`

The authenticationoauthaction data source allows you to retrieve information about OAuth authentication actions configured on Citrix ADC.


## Example usage

```terraform
data "citrixadc_authenticationoauthaction" "tf_authenticationoauthaction" {
  name = "my_oauth_action"
}

output "tokenendpoint" {
  value = data.citrixadc_authenticationoauthaction.tf_authenticationoauthaction.tokenendpoint
}

output "clientid" {
  value = data.citrixadc_authenticationoauthaction.tf_authenticationoauthaction.clientid
}
```


## Argument Reference

* `name` - (Required) Name for the OAuth Authentication action.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `allowedalgorithms` - Multivalued option to specify allowed token verification algorithms.
* `attribute1` - Name of the attribute to be extracted from OAuth Token and to be stored in the attribute1.
* `attribute2` - Name of the attribute to be extracted from OAuth Token and to be stored in the attribute2.
* `attribute3` - Name of the attribute to be extracted from OAuth Token and to be stored in the attribute3.
* `attribute4` - Name of the attribute to be extracted from OAuth Token and to be stored in the attribute4.
* `attribute5` - Name of the attribute to be extracted from OAuth Token and to be stored in the attribute5.
* `attribute6` - Name of the attribute to be extracted from OAuth Token and to be stored in the attribute6.
* `attribute7` - Name of the attribute to be extracted from OAuth Token and to be stored in the attribute7.
* `attribute8` - Name of the attribute to be extracted from OAuth Token and to be stored in the attribute8.
* `attribute9` - Name of the attribute to be extracted from OAuth Token and to be stored in the attribute9.
* `attribute10` - Name of the attribute to be extracted from OAuth Token and to be stored in the attribute10.
* `attribute11` - Name of the attribute to be extracted from OAuth Token and to be stored in the attribute11.
* `attribute12` - Name of the attribute to be extracted from OAuth Token and to be stored in the attribute12.
* `attribute13` - Name of the attribute to be extracted from OAuth Token and to be stored in the attribute13.
* `attribute14` - Name of the attribute to be extracted from OAuth Token and to be stored in the attribute14.
* `attribute15` - Name of the attribute to be extracted from OAuth Token and to be stored in the attribute15.
* `attribute16` - Name of the attribute to be extracted from OAuth Token and to be stored in the attribute16.
* `attributes` - List of attribute names separated by ',' which needs to be extracted. Note that preceding and trailing spaces will be removed. Attribute name can be 127 bytes and total length of this string should not cross 1023 bytes. These attributes have multi-value support separated by ',' and stored as key-value pair in AAA session.
* `audience` - Audience for which token sent by Authorization server is applicable. This is typically entity name or url that represents the recipient.
* `authentication` - If authentication is disabled, password is not sent in the request.
* `authorizationendpoint` - Authorization endpoint/url to which unauthenticated user will be redirected. Citrix ADC redirects user to this endpoint by adding query parameters including clientid.
* `certendpoint` - URL of the endpoint that contains JWKs (Json Web Key) for JWT (Json Web Token) verification.
* `certfilepath` - Path to the file that contains JWKs (Json Web Key) for JWT (Json Web Token) verification.
* `clientid` - Unique identity of the client/user who is getting authenticated. Authorization server infers client configuration using this ID.
* `clientsecret` - Secret string established by user and authorization server.
* `defaultauthenticationgroup` - This is the default group that is chosen when the authentication succeeds in addition to extracted groups.
* `granttype` - Grant type support. Value can be code or password.
* `graphendpoint` - URL of the Graph API service to learn Enterprise Mobility Services (EMS) endpoints.
* `idtokendecryptendpoint` - URL to which obtained idtoken will be posted to get a decrypted user identity.
* `introspecturl` - URL to which access token would be posted for validation.
* `intunedeviceidexpression` - The expression that will be evaluated to obtain IntuneDeviceId for compliance check against IntuneNAC device compliance endpoint.
* `issuer` - Identity of the server whose tokens are to be accepted.
* `metadataurl` - Well-known configuration endpoint of the Authorization Server. Citrix ADC fetches server details from this endpoint.
* `oauthmiscflags` - Option to set/unset miscellaneous feature flags.
* `oauthtype` - Type of the OAuth implementation. Default value is generic implementation that is applicable for most deployments.
* `pkce` - Option to enable/disable PKCE flow during authentication.
* `refreshinterval` - Interval at which services are monitored for necessary configuration.
* `requestattribute` - Name-Value pairs of attributes to be inserted in request parameter. Configuration format is name=value_expr@@@name2=value2_expr@@@.
* `resourceuri` - Resource URL for OAuth configuration.
* `skewtime` - This option specifies the allowed clock skew in number of minutes that Citrix ADC allows on an incoming token.
* `tenantid` - TenantID of the application. This is usually specific to providers such as Microsoft and usually refers to the deployment identifier.
* `tokenendpoint` - URL to which OAuth token will be posted to verify its authenticity.
* `tokenendpointauthmethod` - Option to select the variant of token authentication method. This method is used while exchanging code with IdP.
* `userinfourl` - URL to which OAuth access token will be posted to obtain user information.
* `usernamefield` - Attribute in the token from which username should be extracted.

## Attribute Reference

* `id` - The id of the authenticationoauthaction. It has the same value as the `name` attribute.


## Import

An authenticationoauthaction can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationoauthaction.tf_authenticationoauthaction my_oauth_action
```
