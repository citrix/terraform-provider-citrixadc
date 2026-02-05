---
subcategory: "Authentication"
---

# Data Source `authenticationsamlaction`

The authenticationsamlaction data source allows you to retrieve information about SAML authentication actions.


## Example usage

```terraform
data "citrixadc_authenticationsamlaction" "tf_samlaction" {
  name = "my_samlaction"
}

output "metadataurl" {
  value = data.citrixadc_authenticationsamlaction.tf_samlaction.metadataurl
}

output "samltwofactor" {
  value = data.citrixadc_authenticationsamlaction.tf_samlaction.samltwofactor
}
```


## Argument Reference

* `name` - (Required) Name for the SAML server profile (action).

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `artifactresolutionserviceurl` - URL of the Artifact Resolution Service on IdP to which Citrix ADC will post artifact to get actual SAML token.
* `attribute1` - Name of the attribute in SAML Assertion whose value needs to be extracted and stored as attribute1. Maximum length of the extracted attribute is 239 bytes.
* `attribute2` - Name of the attribute in SAML Assertion whose value needs to be extracted and stored as attribute2. Maximum length of the extracted attribute is 239 bytes.
* `attribute3` - Name of the attribute in SAML Assertion whose value needs to be extracted and stored as attribute3. Maximum length of the extracted attribute is 239 bytes.
* `attribute4` - Name of the attribute in SAML Assertion whose value needs to be extracted and stored as attribute4. Maximum length of the extracted attribute is 239 bytes.
* `attribute5` - Name of the attribute in SAML Assertion whose value needs to be extracted and stored as attribute5. Maximum length of the extracted attribute is 239 bytes.
* `attribute6` - Name of the attribute in SAML Assertion whose value needs to be extracted and stored as attribute6. Maximum length of the extracted attribute is 239 bytes.
* `attribute7` - Name of the attribute in SAML Assertion whose value needs to be extracted and stored as attribute7. Maximum length of the extracted attribute is 239 bytes.
* `attribute8` - Name of the attribute in SAML Assertion whose value needs to be extracted and stored as attribute8. Maximum length of the extracted attribute is 239 bytes.
* `attribute9` - Name of the attribute in SAML Assertion whose value needs to be extracted and stored as attribute9. Maximum length of the extracted attribute is 239 bytes.
* `attribute10` - Name of the attribute in SAML Assertion whose value needs to be extracted and stored as attribute10. Maximum length of the extracted attribute is 239 bytes.
* `attribute11` - Name of the attribute in SAML Assertion whose value needs to be extracted and stored as attribute11. Maximum length of the extracted attribute is 239 bytes.
* `attribute12` - Name of the attribute in SAML Assertion whose value needs to be extracted and stored as attribute12. Maximum length of the extracted attribute is 239 bytes.
* `attribute13` - Name of the attribute in SAML Assertion whose value needs to be extracted and stored as attribute13. Maximum length of the extracted attribute is 239 bytes.
* `attribute14` - Name of the attribute in SAML Assertion whose value needs to be extracted and stored as attribute14. Maximum length of the extracted attribute is 239 bytes.
* `attribute15` - Name of the attribute in SAML Assertion whose value needs to be extracted and stored as attribute15. Maximum length of the extracted attribute is 239 bytes.
* `attribute16` - Name of the attribute in SAML Assertion whose value needs to be extracted and stored as attribute16. Maximum length of the extracted attribute is 239 bytes.
* `attributeconsumingserviceindex` - Index/ID of the attribute specification at Identity Provider (IdP). IdP will locate attributes requested by SP using this index and send those attributes in Assertion.
* `attributes` - List of attribute names separated by ',' which needs to be extracted. Note that preceeding and trailing spaces will be removed. Attribute name can be 127 bytes and total length of this string should not cross 2047 bytes. These attributes have multi-value support separated by ',' and stored as key-value pair in AAA session.
* `audience` - Audience for which assertion sent by IdP is applicable. This is typically entity name or url that represents ServiceProvider.
* `authnctxclassref` - This element specifies the authentication class types that are requested from IdP (IdentityProvider).
* `customauthnctxclassref` - This element specifies the custom authentication class reference to be sent as a part of the Authentication Request that is sent by the SP to SAML IDP. The input string must be the body of the authentication class being requested.
* `defaultauthenticationgroup` - This is the default group that is chosen when the authentication succeeds in addition to extracted groups.
* `digestmethod` - Algorithm to be used to compute/verify digest for SAML transactions.
* `enforceusername` - Option to choose whether the username that is extracted from SAML assertion can be edited in login page while doing second factor.
* `forceauthn` - Option that forces authentication at the Identity Provider (IdP) that receives Citrix ADC's request.
* `groupnamefield` - Name of the tag in assertion that contains user groups.
* `logoutbinding` - This element specifies the transport mechanism of saml logout messages.
* `logouturl` - SingleLogout URL on IdP to which logoutRequest will be sent on Citrix ADC session cleanup.
* `metadatarefreshinterval` - Interval in minutes for fetching metadata from specified metadata URL.
* `metadataurl` - This URL is used for obtaining saml metadata. Note that it fills samlIdPCertName and samlredirectUrl fields so those fields should not be updated when metadataUrl present.
* `preferredbindtype` - This element specifies the preferred binding types for sso and logout for metadata configuration.
* `relaystaterule` - Boolean expression that will be evaluated to validate the SAML Response.
* `requestedauthncontext` - This element specifies the authentication context requirements of authentication statements returned in the response.
* `samlacsindex` - Index/ID of the metadata entry corresponding to this configuration.
* `samlbinding` - This element specifies the transport mechanism of saml messages.
* `samlidpcertname` - Name of the SSL certificate used to verify responses from SAML Identity Provider (IdP). Note that if metadateURL is present then this filed should be empty.
* `samlissuername` - The name to be used in requests sent from Citrix ADC to IdP to uniquely identify Citrix ADC.
* `samlredirecturl` - URL to which users are redirected for authentication. Note that if metadateURL is present then this filed should be empty.
* `samlrejectunsignedassertion` - Reject unsigned SAML assertions. ON option results in rejection of Assertion that is received without signature. STRICT option ensures that both Response and Assertion are signed.
* `samlsigningcertname` - Name of the SSL certificate to sign requests from ServiceProvider (SP) to Identity Provider (IdP).
* `samltwofactor` - Option to enable second factor after SAML.
* `samluserfield` - SAML user ID, as given in the SAML assertion.
* `sendthumbprint` - Option to send thumbprint instead of x509 certificate in SAML request.
* `signaturealg` - Algorithm to be used to sign/verify SAML transactions.
* `skewtime` - This option specifies the allowed clock skew in number of minutes that Citrix ADC ServiceProvider allows on an incoming assertion. For example, if skewTime is 10, then assertion would be valid from (current time - 10) min to (current time + 10) min, ie 20min in all.
* `statechecks` - Boolean expression that will be evaluated to validate HTTP requests on SAML endpoints.
* `storesamlresponse` - Option to store entire SAML Response through the life of user session.

## Attribute Reference

* `id` - The id of the authenticationsamlaction. It has the same value as the `name` attribute.


## Import

A authenticationsamlaction can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationsamlaction.tf_samlaction my_samlaction
```
