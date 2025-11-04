---
subcategory: "Authentication"
---

# Resource: authenticationoauthaction

The authenticationoauthaction resource is used to create OAuth authentication action resource.


## Example usage

```hcl
resource "citrixadc_authenticationoauthaction" "tf_authenticationoauthaction" {
  name                  = "tf_authenticationoauthaction"
  authorizationendpoint = "https://example.com/"
  tokenendpoint         = "https://ssexample.com/"
  clientid              = "cliId"
  clientsecret          = "secret"
  resourceuri           = "http://www.sd.com"
}
```


## Argument Reference

* `name` - (Required) Name for the OAuth Authentication action.  Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the profile is created.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my authentication action" or 'my authentication action').
* `tokenendpoint` - (Required) URL to which OAuth token will be posted to verify its authenticity. User obtains this token from Authorization server upon successful authentication. Citrix ADC will validate presented token by posting it to the URL configured.
* `authorizationendpoint` - (Required) Authorization endpoint/url to which unauthenticated user will be redirected. Citrix ADC redirects user to this endpoint by adding query parameters including clientid. If this parameter not specified then as default value we take Token Endpoint/URL value. Please note that Authorization Endpoint or Token Endpoint is mandatory for oauthAction
* `clientid` - (Required) Unique identity of the client/user who is getting authenticated. Authorization server infers client configuration using this ID
* `clientsecret` - (Required) Secret string established by user and authorization server
* `allowedalgorithms` - (Optional) Multivalued option to specify allowed token verification algorithms.
* `attribute1` - (Optional) Name of the attribute to be extracted from OAuth Token and to be stored in the attribute1
* `attribute10` - (Optional) Name of the attribute to be extracted from OAuth Token and to be stored in the attribute10
* `attribute11` - (Optional) Name of the attribute to be extracted from OAuth Token and to be stored in the attribute11
* `attribute12` - (Optional) Name of the attribute to be extracted from OAuth Token and to be stored in the attribute12
* `attribute13` - (Optional) Name of the attribute to be extracted from OAuth Token and to be stored in the attribute13
* `attribute14` - (Optional) Name of the attribute to be extracted from OAuth Token and to be stored in the attribute14
* `attribute15` - (Optional) Name of the attribute to be extracted from OAuth Token and to be stored in the attribute15
* `attribute16` - (Optional) Name of the attribute to be extracted from OAuth Token and to be stored in the attribute16
* `attribute2` - (Optional) Name of the attribute to be extracted from OAuth Token and to be stored in the attribute2
* `attribute3` - (Optional) Name of the attribute to be extracted from OAuth Token and to be stored in the attribute3
* `attribute4` - (Optional) Name of the attribute to be extracted from OAuth Token and to be stored in the attribute4
* `attribute5` - (Optional) Name of the attribute to be extracted from OAuth Token and to be stored in the attribute5
* `attribute6` - (Optional) Name of the attribute to be extracted from OAuth Token and to be stored in the attribute6
* `attribute7` - (Optional) Name of the attribute to be extracted from OAuth Token and to be stored in the attribute7
* `attribute8` - (Optional) Name of the attribute to be extracted from OAuth Token and to be stored in the attribute8
* `attribute9` - (Optional) Name of the attribute to be extracted from OAuth Token and to be stored in the attribute9
* `attributes` - (Optional) List of attribute names separated by ',' which needs to be extracted. Note that preceding and trailing spaces will be removed. Attribute name can be 127 bytes and total length of this string should not cross 1023 bytes. These attributes have multi-value support separated by ',' and stored as key-value pair in AAA session
* `audience` - (Optional) Audience for which token sent by Authorization server is applicable. This is typically entity name or url that represents the recipient
* `authentication` - (Optional) If authentication is disabled, password is not sent in the request.
* `certendpoint` - (Optional) URL of the endpoint that contains JWKs (Json Web Key) for JWT (Json Web Token) verification.
* `certfilepath` - (Optional) Path to the file that contains JWKs (Json Web Key) for JWT (Json Web Token) verification.
* `defaultauthenticationgroup` - (Optional) This is the default group that is chosen when the authentication succeeds in addition to extracted groups.
* `granttype` - (Optional) Grant type support. value can be code or password
* `graphendpoint` - (Optional) URL of the Graph API service to learn Enterprise Mobility Services (EMS) endpoints.
* `idtokendecryptendpoint` - (Optional) URL to which obtained idtoken will be posted to get a decrypted user identity. Encrypted idtoken will be obtained by posting OAuth token to token endpoint. In order to decrypt idtoken, Citrix ADC posts request to the URL configured
* `introspecturl` - (Optional) URL to which access token would be posted for validation
* `issuer` - (Optional) Identity of the server whose tokens are to be accepted.
* `metadataurl` - (Optional) Well-known configuration endpoint of the Authorization Server. Citrix ADC fetches server details from this endpoint.
* `oauthtype` - (Optional) Type of the OAuth implementation. Default value is generic implementation that is applicable for most deployments.
* `pkce` - (Optional) Option to enable/disable PKCE flow during authentication.
* `refreshinterval` - (Optional) Interval at which services are monitored for necessary configuration.
* `resourceuri` - (Optional) Resource URL for Oauth configuration.
* `skewtime` - (Optional) This option specifies the allowed clock skew in number of minutes that Citrix ADC allows on an incoming token. For example, if skewTime is 10, then token would be valid from (current time - 10) min to (current time + 10) min, ie 20min in all.
* `tenantid` - (Optional) TenantID of the application. This is usually specific to providers such as Microsoft and usually refers to the deployment identifier.
* `tokenendpointauthmethod` - (Optional) Option to select the variant of token authentication method. This method is used while exchanging code with IdP.
* `userinfourl` - (Optional) URL to which OAuth access token will be posted to obtain user information.
* `usernamefield` - (Optional) Attribute in the token from which username should be extracted.
* `intunedeviceidexpression` - (Optional) The expression that will be evaluated to obtain IntuneDeviceId for compliance check against IntuneNAC device compliance endpoint. The expression is applicable when the OAuthType is INTUNE. The maximum length allowed to be used as IntuneDeviceId for the device compliance check from the computed response after the expression evaluation is 41. Examples: add authentication oauthAction <actionName> -intuneDeviceIdExpression 'AAA.LOGIN.INTUNEURI.AFTER_STR("IntuneDeviceId://")'
* `oauthmiscflags` - (Optional) Option to set/unset miscellaneous feature flags. Available values function as follows: * Base64Encode_Authorization_With_Padding - On setting this value, for endpoints (token and introspect), basic authorization header will be base64 encoded with padding. * EnableJWTRequest - By enabling this field, Authorisation request to IDP will have jwt signed 'request' parameter
* `requestattribute` - (Optional) Name-Value pairs of attributes to be inserted in request parameter. Configuration format is name=value_expr@@@name2=value2_expr@@@. '@@@' is used as delimiter between Name-Value pairs. name is a literal string whose value is 127 characters and does not contain '=' character. Value is advanced policy expression terminated by @@@ delimiter. Last value need not contain the delimiter.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authenticationoauthaction. It has the same value as the `name` attribute.


## Import

A authenticationoauthaction can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationoauthaction.tf_authenticationoauthaction tf_authenticationoauthaction
```
