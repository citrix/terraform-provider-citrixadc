---
subcategory: "Authentication"
---

# Resource: authenticationoauthidpprofile

The authenticationoauthidpprofile resource is used to create authenticationOAuthIdpProfile resource.


## Example usage

```hcl
resource "citrixadc_authenticationoauthidpprofile" "tf_idpprofile" {
  name         = "tf_idpprofile"
  clientid     = "cliId"
  clientsecret = "secret"
  redirecturl  = "http://www.example.com/1/"
}
```


## Argument Reference

* `name` - (Required) Name for the new OAuth Identity Provider (IdP) single sign-on profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after an action is created.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my action" or 'my action').
* `clientid` - (Required) Unique identity of the relying party requesting for authentication.
* `clientsecret` - (Required) Unique secret string to authorize relying party at authorization server.
* `redirecturl` - (Required) URL endpoint on relying party to which the OAuth token is to be sent.
* `attributes` - (Optional) Name-Value pairs of attributes to be inserted in idtoken. Configuration format is name=value_expr@@@name2=value2_expr@@@. '@@@' is used as delimiter between Name-Value pairs. name is a literal string whose value is 127 characters and does not contain '=' character. Value is advanced policy expression terminated by @@@ delimiter. Last value need not contain the delimiter.
* `audience` - (Optional) Audience for which token is being sent by Citrix ADC IdP. This is typically entity name or url that represents the recipient
* `configservice` - (Optional) Name of the entity that is used to obtain configuration for the current authentication request. It is used only in Citrix Cloud.
* `defaultauthenticationgroup` - (Optional) This group will be part of AAA session's internal group list. This will be helpful to admin in Nfactor flow to decide right AAA configuration for Relaying Party. In authentication policy AAA.USER.IS_MEMBER_OF("<default_auth_group>")  is way to use this feature.
* `encrypttoken` - (Optional) Option to encrypt token when Citrix ADC IDP sends one.
* `issuer` - (Optional) The name to be used in requests sent from	Citrix ADC to IdP to uniquely identify Citrix ADC.
* `refreshinterval` - (Optional) Interval at which Relying Party metadata is refreshed.
* `relyingpartymetadataurl` - (Optional) This is the endpoint at which Citrix ADC IdP can get details about Relying Party (RP) being configured. Metadata response should include endpoints for jwks_uri for RP public key(s).
* `sendpassword` - (Optional) Option to send encrypted password in idtoken.
* `signaturealg` - (Optional) Algorithm to be used to sign OpenID tokens.
* `signatureservice` - (Optional) Name of the service in cloud used to sign the data. This is applicable only if signature if offloaded to cloud.
* `skewtime` - (Optional) This option specifies the duration for which the token sent by Citrix ADC IdP is valid. For example, if skewTime is 10, then token would be valid from (current time - 10) min to (current time + 10) min, ie 20min in all.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authenticationoauthidpprofile. It has the same value as the `name` attribute.


## Import

A authenticationoauthidpprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationoauthidpprofile.tf_idpprofile tf_idpprofile
```
