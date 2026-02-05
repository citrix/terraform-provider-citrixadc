---
subcategory: "Authentication"
---

# Data Source: citrixadc_authenticationoauthidpprofile

This data source is used to retrieve information about a specific `authenticationoauthidpprofile` resource.

## Example Usage

```hcl
data "citrixadc_authenticationoauthidpprofile" "example" {
  name = "my_oauth_idp_profile"
}
```

## Argument Reference

* `name` - (Required) Name for the new OAuth Identity Provider (IdP) single sign-on profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after an action is created.

## Attribute Reference

In addition to the argument, the following attributes are exported:

* `id` - The ID of the authenticationoauthidpprofile.
* `attributes` - Name-Value pairs of attributes to be inserted in idtoken. Configuration format is name=value_expr@@@name2=value2_expr@@@. '@@@' is used as delimiter between Name-Value pairs. name is a literal string whose value is 127 characters and does not contain '=' character. Value is advanced policy expression terminated by @@@ delimiter. Last value need not contain the delimiter.
* `audience` - Audience for which token is being sent by Citrix ADC IdP. This is typically entity name or url that represents the recipient.
* `clientid` - Unique identity of the relying party requesting for authentication.
* `clientsecret` - Unique secret string to authorize relying party at authorization server.
* `configservice` - Name of the entity that is used to obtain configuration for the current authentication request. It is used only in Citrix Cloud.
* `defaultauthenticationgroup` - This group will be part of AAA session's internal group list. This will be helpful to admin in Nfactor flow to decide right AAA configuration for Relaying Party. In authentication policy AAA.USER.IS_MEMBER_OF("<default_auth_group>") is way to use this feature.
* `encrypttoken` - Option to encrypt token when Citrix ADC IDP sends one.
* `issuer` - The name to be used in requests sent from Citrix ADC to IdP to uniquely identify Citrix ADC.
* `redirecturl` - URL endpoint on relying party to which the OAuth token is to be sent.
* `refreshinterval` - Interval at which Relying Party metadata is refreshed.
* `relyingpartymetadataurl` - This is the endpoint at which Citrix ADC IdP can get details about Relying Party (RP) being configured. Metadata response should include endpoints for jwks_uri for RP public key(s).
* `sendpassword` - Option to send encrypted password in idtoken.
* `signaturealg` - Algorithm to be used to sign OpenID tokens.
* `signatureservice` - Name of the service in cloud used to sign the data. This is applicable only if signature if offloaded to cloud.
* `skewtime` - This option specifies the duration for which the token sent by Citrix ADC IdP is valid. For example, if skewTime is 10, then token would be valid from (current time - 10) min to (current time + 10) min, ie 20min in all.
