---
subcategory: "Traffic Management"
---

# Resource:  tmsamlssoprofile

The tmsamlssoprofile resource is used to create tmsamlssoprofile.


## Example usage

```hcl
resource "citrixadc_tmsamlssoprofile" "tf_tmsamlssoprofile" {
  name                        = "my_tmsamlssoprofile"
  assertionconsumerserviceurl = "https://service.example.com"
  sendpassword                = "ON"
  relaystaterule              = "true"
}

```


## Argument Reference

* `name` - (Required) Name for the new saml single sign-on profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after an SSO action is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my action" or 'my action'). Minimum length =  1
* `samlsigningcertname` - (Optional) Name of the SSL certificate that is used to Sign Assertion. Minimum length =  1
* `assertionconsumerserviceurl` - (Optional) URL to which the assertion is to be sent. Minimum length =  1
* `relaystaterule` - (Optional) Expression to extract relaystate to be sent along with assertion. Evaluation of this expression should return TEXT content. This is typically a targ et url to which user is redirected after the recipient validates SAML token.
* `sendpassword` - (Optional) Option to send password in assertion. Possible values: [ on, off ]
* `samlissuername` - (Optional) The name to be used in requests sent from	Citrix ADC to IdP to uniquely identify Citrix ADC. Minimum length =  1
* `signaturealg` - (Optional) Algorithm to be used to sign/verify SAML transactions. Possible values: [ RSA-SHA1, RSA-SHA256 ]
* `digestmethod` - (Optional) Algorithm to be used to compute/verify digest for SAML transactions. Possible values: [ SHA1, SHA256 ]
* `audience` - (Optional) Audience for which assertion sent by IdP is applicable. This is typically entity name or url that represents ServiceProvider.
* `nameidformat` - (Optional) Format of Name Identifier sent in Assertion. Possible values: [ Unspecified, emailAddress, X509SubjectName, WindowsDomainQualifiedName, kerberos, entity, persistent, transient ]
* `nameidexpr` - (Optional) Expression that will be evaluated to obtain NameIdentifier to be sent in assertion. Maximum length =  128
* `attribute1` - (Optional) Name of attribute1 that needs to be sent in SAML Assertion.
* `attribute1expr` - (Optional) Expression that will be evaluated to obtain attribute1's value to be sent in Assertion. Maximum length =  128
* `attribute1friendlyname` - (Optional) User-Friendly Name of attribute1 that needs to be sent in SAML Assertion.
* `attribute1format` - (Optional) Format of Attribute1 to be sent in Assertion. Possible values: [ URI, Basic ]
* `attribute2` - (Optional) Name of attribute2 that needs to be sent in SAML Assertion.
* `attribute2expr` - (Optional) Expression that will be evaluated to obtain attribute2's value to be sent in Assertion. Maximum length =  128
* `attribute2friendlyname` - (Optional) User-Friendly Name of attribute2 that needs to be sent in SAML Assertion.
* `attribute2format` - (Optional) Format of Attribute2 to be sent in Assertion. Possible values: [ URI, Basic ]
* `attribute3` - (Optional) Name of attribute3 that needs to be sent in SAML Assertion.
* `attribute3expr` - (Optional) Expression that will be evaluated to obtain attribute3's value to be sent in Assertion. Maximum length =  128
* `attribute3friendlyname` - (Optional) User-Friendly Name of attribute3 that needs to be sent in SAML Assertion.
* `attribute3format` - (Optional) Format of Attribute3 to be sent in Assertion. Possible values: [ URI, Basic ]
* `attribute4` - (Optional) Name of attribute4 that needs to be sent in SAML Assertion.
* `attribute4expr` - (Optional) Expression that will be evaluated to obtain attribute4's value to be sent in Assertion. Maximum length =  128
* `attribute4friendlyname` - (Optional) User-Friendly Name of attribute4 that needs to be sent in SAML Assertion.
* `attribute4format` - (Optional) Format of Attribute4 to be sent in Assertion. Possible values: [ URI, Basic ]
* `attribute5` - (Optional) Name of attribute5 that needs to be sent in SAML Assertion.
* `attribute5expr` - (Optional) Expression that will be evaluated to obtain attribute5's value to be sent in Assertion. Maximum length =  128
* `attribute5friendlyname` - (Optional) User-Friendly Name of attribute5 that needs to be sent in SAML Assertion.
* `attribute5format` - (Optional) Format of Attribute5 to be sent in Assertion. Possible values: [ URI, Basic ]
* `attribute6` - (Optional) Name of attribute6 that needs to be sent in SAML Assertion.
* `attribute6expr` - (Optional) Expression that will be evaluated to obtain attribute6's value to be sent in Assertion. Maximum length =  128
* `attribute6friendlyname` - (Optional) User-Friendly Name of attribute6 that needs to be sent in SAML Assertion.
* `attribute6format` - (Optional) Format of Attribute6 to be sent in Assertion. Possible values: [ URI, Basic ]
* `attribute7` - (Optional) Name of attribute7 that needs to be sent in SAML Assertion.
* `attribute7expr` - (Optional) Expression that will be evaluated to obtain attribute7's value to be sent in Assertion. Maximum length =  128
* `attribute7friendlyname` - (Optional) User-Friendly Name of attribute7 that needs to be sent in SAML Assertion.
* `attribute7format` - (Optional) Format of Attribute7 to be sent in Assertion. Possible values: [ URI, Basic ]
* `attribute8` - (Optional) Name of attribute8 that needs to be sent in SAML Assertion.
* `attribute8expr` - (Optional) Expression that will be evaluated to obtain attribute8's value to be sent in Assertion. Maximum length =  128
* `attribute8friendlyname` - (Optional) User-Friendly Name of attribute8 that needs to be sent in SAML Assertion.
* `attribute8format` - (Optional) Format of Attribute8 to be sent in Assertion. Possible values: [ URI, Basic ]
* `attribute9` - (Optional) Name of attribute9 that needs to be sent in SAML Assertion.
* `attribute9expr` - (Optional) Expression that will be evaluated to obtain attribute9's value to be sent in Assertion. Maximum length =  128
* `attribute9friendlyname` - (Optional) User-Friendly Name of attribute9 that needs to be sent in SAML Assertion.
* `attribute9format` - (Optional) Format of Attribute9 to be sent in Assertion. Possible values: [ URI, Basic ]
* `attribute10` - (Optional) Name of attribute10 that needs to be sent in SAML Assertion.
* `attribute10expr` - (Optional) Expression that will be evaluated to obtain attribute10's value to be sent in Assertion. Maximum length =  128
* `attribute10friendlyname` - (Optional) User-Friendly Name of attribute10 that needs to be sent in SAML Assertion.
* `attribute10format` - (Optional) Format of Attribute10 to be sent in Assertion. Possible values: [ URI, Basic ]
* `attribute11` - (Optional) Name of attribute11 that needs to be sent in SAML Assertion.
* `attribute11expr` - (Optional) Expression that will be evaluated to obtain attribute11's value to be sent in Assertion. Maximum length =  128
* `attribute11friendlyname` - (Optional) User-Friendly Name of attribute11 that needs to be sent in SAML Assertion.
* `attribute11format` - (Optional) Format of Attribute11 to be sent in Assertion. Possible values: [ URI, Basic ]
* `attribute12` - (Optional) Name of attribute12 that needs to be sent in SAML Assertion.
* `attribute12expr` - (Optional) Expression that will be evaluated to obtain attribute12's value to be sent in Assertion. Maximum length =  128
* `attribute12friendlyname` - (Optional) User-Friendly Name of attribute12 that needs to be sent in SAML Assertion.
* `attribute12format` - (Optional) Format of Attribute12 to be sent in Assertion. Possible values: [ URI, Basic ]
* `attribute13` - (Optional) Name of attribute13 that needs to be sent in SAML Assertion.
* `attribute13expr` - (Optional) Expression that will be evaluated to obtain attribute13's value to be sent in Assertion. Maximum length =  128
* `attribute13friendlyname` - (Optional) User-Friendly Name of attribute13 that needs to be sent in SAML Assertion.
* `attribute13format` - (Optional) Format of Attribute13 to be sent in Assertion. Possible values: [ URI, Basic ]
* `attribute14` - (Optional) Name of attribute14 that needs to be sent in SAML Assertion.
* `attribute14expr` - (Optional) Expression that will be evaluated to obtain attribute14's value to be sent in Assertion. Maximum length =  128
* `attribute14friendlyname` - (Optional) User-Friendly Name of attribute14 that needs to be sent in SAML Assertion.
* `attribute14format` - (Optional) Format of Attribute14 to be sent in Assertion. Possible values: [ URI, Basic ]
* `attribute15` - (Optional) Name of attribute15 that needs to be sent in SAML Assertion.
* `attribute15expr` - (Optional) Expression that will be evaluated to obtain attribute15's value to be sent in Assertion. Maximum length =  128
* `attribute15friendlyname` - (Optional) User-Friendly Name of attribute15 that needs to be sent in SAML Assertion.
* `attribute15format` - (Optional) Format of Attribute15 to be sent in Assertion. Possible values: [ URI, Basic ]
* `attribute16` - (Optional) Name of attribute16 that needs to be sent in SAML Assertion.
* `attribute16expr` - (Optional) Expression that will be evaluated to obtain attribute16's value to be sent in Assertion. Maximum length =  128
* `attribute16friendlyname` - (Optional) User-Friendly Name of attribute16 that needs to be sent in SAML Assertion.
* `attribute16format` - (Optional) Format of Attribute16 to be sent in Assertion. Possible values: [ URI, Basic ]
* `encryptassertion` - (Optional) Option to encrypt assertion when Citrix ADC sends one. Possible values: [ on, off ]
* `samlspcertname` - (Optional) Name of the SSL certificate of peer/receving party using which Assertion is encrypted. Minimum length =  1
* `encryptionalgorithm` - (Optional) Algorithm to be used to encrypt SAML assertion. Possible values: [ DES3, AES128, AES192, AES256 ]
* `skewtime` - (Optional) This option specifies the number of minutes on either side of current time that the assertion would be valid. For example, if skewTime is 10, then assertion would be valid from (current time - 10) min to (current time + 10) min, ie 20min in all.
* `signassertion` - (Optional) Option to sign portions of assertion when Citrix ADC IDP sends one. Based on the user selection, either Assertion or Response or Both or none can be signed. Possible values: [ NONE, ASSERTION, RESPONSE, BOTH ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the tmsamlssoprofile. It has the same value as the `name` attribute.


## Import

A tmsamlssoprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_tmsamlssoprofile.tf_tmsamlssoprofile my_tmsamlssoprofile
```
