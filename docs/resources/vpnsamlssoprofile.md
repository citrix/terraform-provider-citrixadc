---
subcategory: "Vpn"
---

# Resource: vpnsamlssoprofile

The vpnsamlssoprofile resource is used to create of vpn SAML sso profile resource.


## Example usage

```hcl
resource "citrixadc_vpnsamlssoprofile" "tf_vpnsamlssoprofile" {
  name                        = "tf_vpnsamlssoprofile"
  assertionconsumerserviceurl = "http://www.example.com"
  sendpassword                = "ON"
  relaystaterule              = "true"
  samlissuername              = "asdf"
  signaturealg                = "RSA-SHA1"
  digestmethod                = "SHA256"
  nameidformat                = "Unspecified"
}
```


## Argument Reference

* `name` - (Required) Name for the new saml single sign-on profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after an SSO action is created.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my action" or 'my action').
* `assertionconsumerserviceurl` - (Required) URL to which the assertion is to be sent.
* `attribute1` - (Optional) Name of attribute1 that needs to be sent in SAML Assertion
* `attribute10` - (Optional) Name of attribute10 that needs to be sent in SAML Assertion
* `attribute10expr` - (Optional) Expression that will be evaluated to obtain attribute10's value to be sent in Assertion
* `attribute10format` - (Optional) Format of Attribute10 to be sent in Assertion.
* `attribute10friendlyname` - (Optional) User-Friendly Name of attribute10 that needs to be sent in SAML Assertion
* `attribute11` - (Optional) Name of attribute11 that needs to be sent in SAML Assertion
* `attribute11expr` - (Optional) Expression that will be evaluated to obtain attribute11's value to be sent in Assertion
* `attribute11format` - (Optional) Format of Attribute11 to be sent in Assertion.
* `attribute11friendlyname` - (Optional) User-Friendly Name of attribute11 that needs to be sent in SAML Assertion
* `attribute12` - (Optional) Name of attribute12 that needs to be sent in SAML Assertion
* `attribute12expr` - (Optional) Expression that will be evaluated to obtain attribute12's value to be sent in Assertion
* `attribute12format` - (Optional) Format of Attribute12 to be sent in Assertion.
* `attribute12friendlyname` - (Optional) User-Friendly Name of attribute12 that needs to be sent in SAML Assertion
* `attribute13` - (Optional) Name of attribute13 that needs to be sent in SAML Assertion
* `attribute13expr` - (Optional) Expression that will be evaluated to obtain attribute13's value to be sent in Assertion
* `attribute13format` - (Optional) Format of Attribute13 to be sent in Assertion.
* `attribute13friendlyname` - (Optional) User-Friendly Name of attribute13 that needs to be sent in SAML Assertion
* `attribute14` - (Optional) Name of attribute14 that needs to be sent in SAML Assertion
* `attribute14expr` - (Optional) Expression that will be evaluated to obtain attribute14's value to be sent in Assertion
* `attribute14format` - (Optional) Format of Attribute14 to be sent in Assertion.
* `attribute14friendlyname` - (Optional) User-Friendly Name of attribute14 that needs to be sent in SAML Assertion
* `attribute15` - (Optional) Name of attribute15 that needs to be sent in SAML Assertion
* `attribute15expr` - (Optional) Expression that will be evaluated to obtain attribute15's value to be sent in Assertion
* `attribute15format` - (Optional) Format of Attribute15 to be sent in Assertion.
* `attribute15friendlyname` - (Optional) User-Friendly Name of attribute15 that needs to be sent in SAML Assertion
* `attribute16` - (Optional) Name of attribute16 that needs to be sent in SAML Assertion
* `attribute16expr` - (Optional) Expression that will be evaluated to obtain attribute16's value to be sent in Assertion
* `attribute16format` - (Optional) Format of Attribute16 to be sent in Assertion.
* `attribute16friendlyname` - (Optional) User-Friendly Name of attribute16 that needs to be sent in SAML Assertion
* `attribute1expr` - (Optional) Expression that will be evaluated to obtain attribute1's value to be sent in Assertion
* `attribute1format` - (Optional) Format of Attribute1 to be sent in Assertion.
* `attribute1friendlyname` - (Optional) User-Friendly Name of attribute1 that needs to be sent in SAML Assertion
* `attribute2` - (Optional) Name of attribute2 that needs to be sent in SAML Assertion
* `attribute2expr` - (Optional) Expression that will be evaluated to obtain attribute2's value to be sent in Assertion
* `attribute2format` - (Optional) Format of Attribute2 to be sent in Assertion.
* `attribute2friendlyname` - (Optional) User-Friendly Name of attribute2 that needs to be sent in SAML Assertion
* `attribute3` - (Optional) Name of attribute3 that needs to be sent in SAML Assertion
* `attribute3expr` - (Optional) Expression that will be evaluated to obtain attribute3's value to be sent in Assertion
* `attribute3format` - (Optional) Format of Attribute3 to be sent in Assertion.
* `attribute3friendlyname` - (Optional) User-Friendly Name of attribute3 that needs to be sent in SAML Assertion
* `attribute4` - (Optional) Name of attribute4 that needs to be sent in SAML Assertion
* `attribute4expr` - (Optional) Expression that will be evaluated to obtain attribute4's value to be sent in Assertion
* `attribute4format` - (Optional) Format of Attribute4 to be sent in Assertion.
* `attribute4friendlyname` - (Optional) User-Friendly Name of attribute4 that needs to be sent in SAML Assertion
* `attribute5` - (Optional) Name of attribute5 that needs to be sent in SAML Assertion
* `attribute5expr` - (Optional) Expression that will be evaluated to obtain attribute5's value to be sent in Assertion
* `attribute5format` - (Optional) Format of Attribute5 to be sent in Assertion.
* `attribute5friendlyname` - (Optional) User-Friendly Name of attribute5 that needs to be sent in SAML Assertion
* `attribute6` - (Optional) Name of attribute6 that needs to be sent in SAML Assertion
* `attribute6expr` - (Optional) Expression that will be evaluated to obtain attribute6's value to be sent in Assertion
* `attribute6format` - (Optional) Format of Attribute6 to be sent in Assertion.
* `attribute6friendlyname` - (Optional) User-Friendly Name of attribute6 that needs to be sent in SAML Assertion
* `attribute7` - (Optional) Name of attribute7 that needs to be sent in SAML Assertion
* `attribute7expr` - (Optional) Expression that will be evaluated to obtain attribute7's value to be sent in Assertion
* `attribute7format` - (Optional) Format of Attribute7 to be sent in Assertion.
* `attribute7friendlyname` - (Optional) User-Friendly Name of attribute7 that needs to be sent in SAML Assertion
* `attribute8` - (Optional) Name of attribute8 that needs to be sent in SAML Assertion
* `attribute8expr` - (Optional) Expression that will be evaluated to obtain attribute8's value to be sent in Assertion
* `attribute8format` - (Optional) Format of Attribute8 to be sent in Assertion.
* `attribute8friendlyname` - (Optional) User-Friendly Name of attribute8 that needs to be sent in SAML Assertion
* `attribute9` - (Optional) Name of attribute9 that needs to be sent in SAML Assertion
* `attribute9expr` - (Optional) Expression that will be evaluated to obtain attribute9's value to be sent in Assertion
* `attribute9format` - (Optional) Format of Attribute9 to be sent in Assertion.
* `attribute9friendlyname` - (Optional) User-Friendly Name of attribute9 that needs to be sent in SAML Assertion
* `audience` - (Optional) Audience for which assertion sent by IdP is applicable. This is typically entity name or url that represents ServiceProvider
* `digestmethod` - (Optional) Algorithm to be used to compute/verify digest for SAML transactions
* `encryptassertion` - (Optional) Option to encrypt assertion when Citrix ADC sends one.
* `encryptionalgorithm` - (Optional) Algorithm to be used to encrypt SAML assertion
* `nameidexpr` - (Optional) Expression that will be evaluated to obtain NameIdentifier to be sent in assertion
* `nameidformat` - (Optional) Format of Name Identifier sent in Assertion.
* `relaystaterule` - (Optional) Expression to extract relaystate to be sent along with assertion. Evaluation of this expression should return TEXT content. This is typically a target url to which user is redirected after the recipient validates SAML token
* `samlissuername` - (Optional) The name to be used in requests sent from Citrix ADC to IdP to uniquely identify Citrix ADC.
* `samlsigningcertname` - (Optional) Name of the signing authority as given in the SAML server's SSL certificate.
* `samlspcertname` - (Optional) Name of the SSL certificate of peer/receving party using which Assertion is encrypted.
* `sendpassword` - (Optional) Option to send password in assertion.
* `signassertion` - (Optional) Option to sign portions of assertion when Citrix ADC IDP sends one. Based on the user selection, either Assertion or Response or Both or none can be signed
* `signaturealg` - (Optional) Algorithm to be used to sign/verify SAML transactions
* `signatureservice` - (Optional) Name of the service in cloud used to sign the data
* `skewtime` - (Optional) This option specifies the number of minutes on either side of current time that the assertion would be valid. For example, if skewTime is 10, then assertion would be valid from (current time - 10) min to (current time + 10) min, ie 20min in all.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnsamlssoprofile. It has the same value as the `name` attribute.


## Import

A vpnsamlssoprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_vpnsamlssoprofile.tf_vpnsamlssoprofile tf_vpnsamlssoprofile
```
