---
subcategory: "Authentication"
---

# Resource: authenticationsamlaction

The authenticationsamlaction resource is used to create authenticationsamlaction resource.


## Example usage

```hcl
resource "citrixadc_authenticationsamlaction" "tf_samlaction" {
  name                    = "tf_samlaction"
  metadataurl             = "http://www.example.com"
  samltwofactor           = "OFF"
  requestedauthncontext   = "minimum"
  digestmethod            = "SHA1"
  signaturealg            = "RSA-SHA256"
  metadatarefreshinterval = 1
}
```


## Argument Reference

* `name` - (Required) Name for the SAML server profile (action).  Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after SAML profile is created.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my authentication action" or 'my authentication action').
* `artifactresolutionserviceurl` - (Optional) URL of the Artifact Resolution Service on IdP to which Citrix ADC will post artifact to get actual SAML token.
* `attribute1` - (Optional) Name of the attribute in SAML Assertion whose value needs to be extracted and stored as attribute1. Maximum length of the extracted attribute is 239 bytes.
* `attribute10` - (Optional) Name of the attribute in SAML Assertion whose value needs to be extracted and stored as attribute10. Maximum length of the extracted attribute is 239 bytes.
* `attribute11` - (Optional) Name of the attribute in SAML Assertion whose value needs to be extracted and stored as attribute11. Maximum length of the extracted attribute is 239 bytes.
* `attribute12` - (Optional) Name of the attribute in SAML Assertion whose value needs to be extracted and stored as attribute12. Maximum length of the extracted attribute is 239 bytes.
* `attribute13` - (Optional) Name of the attribute in SAML Assertion whose value needs to be extracted and stored as attribute13. Maximum length of the extracted attribute is 239 bytes.
* `attribute14` - (Optional) Name of the attribute in SAML Assertion whose value needs to be extracted and stored as attribute14. Maximum length of the extracted attribute is 239 bytes.
* `attribute15` - (Optional) Name of the attribute in SAML Assertion whose value needs to be extracted and stored as attribute15. Maximum length of the extracted attribute is 239 bytes.
* `attribute16` - (Optional) Name of the attribute in SAML Assertion whose value needs to be extracted and stored as attribute16. Maximum length of the extracted attribute is 239 bytes.
* `attribute2` - (Optional) Name of the attribute in SAML Assertion whose value needs to be extracted and stored as attribute2. Maximum length of the extracted attribute is 239 bytes.
* `attribute3` - (Optional) Name of the attribute in SAML Assertion whose value needs to be extracted and stored as attribute3. Maximum length of the extracted attribute is 239 bytes.
* `attribute4` - (Optional) Name of the attribute in SAML Assertion whose value needs to be extracted and stored as attribute4. Maximum length of the extracted attribute is 239 bytes.
* `attribute5` - (Optional) Name of the attribute in SAML Assertion whose value needs to be extracted and stored as attribute5. Maximum length of the extracted attribute is 239 bytes.
* `attribute6` - (Optional) Name of the attribute in SAML Assertion whose value needs to be extracted and stored as attribute6. Maximum length of the extracted attribute is 239 bytes.
* `attribute7` - (Optional) Name of the attribute in SAML Assertion whose value needs to be extracted and stored as attribute7. Maximum length of the extracted attribute is 239 bytes.
* `attribute8` - (Optional) Name of the attribute in SAML Assertion whose value needs to be extracted and stored as attribute8. Maximum length of the extracted attribute is 239 bytes.
* `attribute9` - (Optional) Name of the attribute in SAML Assertion whose value needs to be extracted and stored as attribute9. Maximum length of the extracted attribute is 239 bytes.
* `attributeconsumingserviceindex` - (Optional) Index/ID of the attribute specification at Identity Provider (IdP). IdP will locate attributes requested by SP using this index and send those attributes in Assertion
* `attributes` - (Optional) List of attribute names separated by ',' which needs to be extracted.  Note that preceeding and trailing spaces will be removed.  Attribute name can be 127 bytes and total length of this string should not cross 2047 bytes. These attributes have multi-value support separated by ',' and stored as key-value pair in AAA session
* `audience` - (Optional) Audience for which assertion sent by IdP is applicable. This is typically entity name or url that represents ServiceProvider
* `authnctxclassref` - (Optional) This element specifies the authentication class types that are requested from IdP (IdentityProvider). InternetProtocol: This is applicable when a principal is authenticated through the use of a provided IP address. InternetProtocolPassword: This is applicable when a principal is authenticated through the use of a provided IP address, in addition to a username/password. Kerberos: This is applicable when the principal has authenticated using a password to a local authentication authority, in order to acquire a Kerberos ticket. MobileOneFactorUnregistered: This indicates authentication of the mobile device without requiring explicit end-user interaction. MobileTwoFactorUnregistered: This indicates two-factor based authentication during mobile customer registration process, such as secure device and user PIN. MobileOneFactorContract: Reflects mobile contract customer registration procedures and a single factor authentication. MobileTwoFactorContract: Reflects mobile contract customer registration procedures and a two-factor based authentication. Password: This class is applicable when a principal authenticates using password over unprotected http session. PasswordProtectedTransport: This class is applicable when a principal authenticates to an authentication authority through the presentation of a password over a protected session. PreviousSession: This class is applicable when a principal had authenticated to an authentication authority at some point in the past using any authentication context. X509: This indicates that the principal authenticated by means of a digital signature where the key was validated as part of an X.509 Public Key Infrastructure. PGP: This indicates that the principal authenticated by means of a digital signature where the key was validated as part of a PGP Public Key Infrastructure. SPKI: This indicates that the principal authenticated by means of a digital signature where the key was validated via an SPKI Infrastructure. XMLDSig: This indicates that the principal authenticated by means of a digital signature according to the processing rules specified in the XML Digital Signature specification. Smartcard: This indicates that the principal has authenticated using smartcard. SmartcardPKI: This class is applicable when a principal authenticates to an authentication authority through a two-factor authentication mechanism using a smartcard with enclosed private key and a PIN. SoftwarePKI: This class is applicable when a principal uses an X.509 certificate stored in software to authenticate to the authentication authority. Telephony: This class is used to indicate that the principal authenticated via the provision of a fixed-line telephone number, transported via a telephony protocol such as ADSL. NomadTelephony: Indicates that the principal is "roaming" and authenticates via the means of the line number, a user suffix, and a password element. PersonalTelephony: This class is used to indicate that the principal authenticated via the provision of a fixed-line telephone. AuthenticatedTelephony: Indicates that the principal authenticated via the means of the line number, a user suffix, and a password element. SecureRemotePassword: This class is applicable when the authentication was performed by means of Secure Remote Password. TLSClient: This class indicates that the principal authenticated by means of a client certificate, secured with the SSL/TLS transport. TimeSyncToken: This is applicable when a principal authenticates through a time synchronization token. Unspecified: This indicates that the authentication was performed by unspecified means. Windows: This indicates that Windows integrated authentication is utilized for authentication.
* `customauthnctxclassref` - (Optional) This element specifies the custom authentication class reference to be sent as a part of the Authentication Request that is sent by the SP to SAML IDP. The input string must be the body of the authentication class being requested. Input format: Alphanumeric string or URL specifying the body of the Request.If more than one string has to be provided, then the same can be done by specifying the classes as a string of comma separated values. Example input: set authentication samlaction samlact1 -customAuthnCtxClassRef http://www.class1.com/LoA1,http://www.class2.com/LoA2
* `defaultauthenticationgroup` - (Optional) This is the default group that is chosen when the authentication succeeds in addition to extracted groups.
* `digestmethod` - (Optional) Algorithm to be used to compute/verify digest for SAML transactions
* `enforceusername` - (Optional) Option to choose whether the username that is extracted from SAML assertion can be edited in login page while doing second factor
* `forceauthn` - (Optional) Option that forces authentication at the Identity Provider (IdP) that receives Citrix ADC's request
* `groupnamefield` - (Optional) Name of the tag in assertion that contains user groups.
* `logoutbinding` - (Optional) This element specifies the transport mechanism of saml logout messages.
* `logouturl` - (Optional) SingleLogout URL on IdP to which logoutRequest will be sent on Citrix ADC session cleanup.
* `metadatarefreshinterval` - (Optional) Interval in minutes for fetching metadata from specified metadata URL
* `metadataurl` - (Optional) This URL is used for obtaining saml metadata. Note that it fills samlIdPCertName and samlredirectUrl fields so those fields should not be updated when metadataUrl present
* `relaystaterule` - (Optional) Boolean expression that will be evaluated to validate the SAML Response. Examples:  set authentication samlaction <actionname> -relaystateRule 'AAA.LOGIN.RELAYSTATE.EQ("https://fqdn.com/")' set authentication samlaction <actionname> -relaystateRule 'AAA.LOGIN.RELAYSTATE.CONTAINS("https://fqdn.com/")' set authentication samlaction <actionname> -relaystateRule 'AAA.LOGIN.RELAYSTATE.CONTAINS_ANY("patset_name")' set authentication samlAction samlsp -relaystateRule 'AAA.LOGIN.RELAYSTATE.REGEX_MATCH(re#http://<regex>.com/#)'.
* `requestedauthncontext` - (Optional) This element specifies the authentication context requirements of authentication statements returned in the response.
* `samlacsindex` - (Optional) Index/ID of the metadata entry corresponding to this configuration.
* `samlbinding` - (Optional) This element specifies the transport mechanism of saml messages.
* `samlidpcertname` - (Optional) Name of the SSL certificate used to verify responses from SAML Identity Provider (IdP). Note that if metadateURL is present then this filed should be empty.
* `samlissuername` - (Optional) The name to be used in requests sent from	Citrix ADC to IdP to uniquely identify Citrix ADC.
* `samlredirecturl` - (Optional) URL to which users are redirected for authentication. Note that if metadateURL is present then this filed should be empty
* `samlrejectunsignedassertion` - (Optional) Reject unsigned SAML assertions. ON option results in rejection of Assertion that is received without signature. STRICT option ensures that both Response and Assertion are signed. OFF allows unsigned Assertions.
* `samlsigningcertname` - (Optional) Name of the SSL certificate to sign requests from ServiceProvider (SP) to Identity Provider (IdP).
* `samltwofactor` - (Optional) Option to enable second factor after SAML
* `samluserfield` - (Optional) SAML user ID, as given in the SAML assertion.
* `sendthumbprint` - (Optional) Option to send thumbprint instead of x509 certificate in SAML request
* `signaturealg` - (Optional) Algorithm to be used to sign/verify SAML transactions
* `skewtime` - (Optional) This option specifies the allowed clock skew in number of minutes that Citrix ADC ServiceProvider allows on an incoming assertion. For example, if skewTime is 10, then assertion would be valid from (current time - 10) min to (current time + 10) min, ie 20min in all.
* `storesamlresponse` - (Optional) Option to store entire SAML Response through the life of user session.
* `preferredbindtype` - (Optional) This element specifies the preferred binding types for sso and logout for metadata configuration.
* `statechecks` - (Optional) Boolean expression that will be evaluated to validate HTTP requests on SAML endpoints. Examples: set authentication samlaction <actionname> -stateChecks 'HTTP.REQ.HOSTNAME.EQ("https://fqdn.com/")'


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authenticationsamlaction. It has the same value as the `name` attribute.


## Import

A authenticationsamlaction can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationsamlaction.tf_samlaction tf_samlaction
```
