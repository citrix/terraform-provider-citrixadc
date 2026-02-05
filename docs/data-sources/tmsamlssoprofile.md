---
subcategory: "Traffic Management"
---

# Data Source: tmsamlssoprofile

The tmsamlssoprofile data source allows you to retrieve information about a SAML SSO profile.

## Example usage

```terraform
data "citrixadc_tmsamlssoprofile" "tf_tmsamlssoprofile" {
  name = "my_tmsamlssoprofile"
}

output "assertionconsumerserviceurl" {
  value = data.citrixadc_tmsamlssoprofile.tf_tmsamlssoprofile.assertionconsumerserviceurl
}

output "relaystaterule" {
  value = data.citrixadc_tmsamlssoprofile.tf_tmsamlssoprofile.relaystaterule
}
```

## Argument Reference

* `name` - (Required) Name for the new saml single sign-on profile.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `assertionconsumerserviceurl` - URL to which the assertion is to be sent.
* `attribute1` - Name of attribute1 that needs to be sent in SAML Assertion.
* `attribute1expr` - Expression that will be evaluated to obtain attribute1's value to be sent in Assertion.
* `attribute1format` - Format of Attribute1 to be sent in Assertion.
* `attribute1friendlyname` - User-Friendly Name of attribute1 that needs to be sent in SAML Assertion.
* `attribute2` - Name of attribute2 that needs to be sent in SAML Assertion.
* `attribute2expr` - Expression that will be evaluated to obtain attribute2's value to be sent in Assertion.
* `attribute2format` - Format of Attribute2 to be sent in Assertion.
* `attribute2friendlyname` - User-Friendly Name of attribute2 that needs to be sent in SAML Assertion.
* `attribute3` - Name of attribute3 that needs to be sent in SAML Assertion.
* `attribute3expr` - Expression that will be evaluated to obtain attribute3's value to be sent in Assertion.
* `attribute3format` - Format of Attribute3 to be sent in Assertion.
* `attribute3friendlyname` - User-Friendly Name of attribute3 that needs to be sent in SAML Assertion.
* `attribute4` - Name of attribute4 that needs to be sent in SAML Assertion.
* `attribute4expr` - Expression that will be evaluated to obtain attribute4's value to be sent in Assertion.
* `attribute4format` - Format of Attribute4 to be sent in Assertion.
* `attribute4friendlyname` - User-Friendly Name of attribute4 that needs to be sent in SAML Assertion.
* `attribute5` - Name of attribute5 that needs to be sent in SAML Assertion.
* `attribute5expr` - Expression that will be evaluated to obtain attribute5's value to be sent in Assertion.
* `attribute5format` - Format of Attribute5 to be sent in Assertion.
* `attribute5friendlyname` - User-Friendly Name of attribute5 that needs to be sent in SAML Assertion.
* `attribute6` - Name of attribute6 that needs to be sent in SAML Assertion.
* `attribute6expr` - Expression that will be evaluated to obtain attribute6's value to be sent in Assertion.
* `attribute6format` - Format of Attribute6 to be sent in Assertion.
* `attribute6friendlyname` - User-Friendly Name of attribute6 that needs to be sent in SAML Assertion.
* `attribute7` - Name of attribute7 that needs to be sent in SAML Assertion.
* `attribute7expr` - Expression that will be evaluated to obtain attribute7's value to be sent in Assertion.
* `attribute7format` - Format of Attribute7 to be sent in Assertion.
* `attribute7friendlyname` - User-Friendly Name of attribute7 that needs to be sent in SAML Assertion.
* `attribute8` - Name of attribute8 that needs to be sent in SAML Assertion.
* `attribute8expr` - Expression that will be evaluated to obtain attribute8's value to be sent in Assertion.
* `attribute8format` - Format of Attribute8 to be sent in Assertion.
* `attribute8friendlyname` - User-Friendly Name of attribute8 that needs to be sent in SAML Assertion.
* `attribute9` - Name of attribute9 that needs to be sent in SAML Assertion.
* `attribute9expr` - Expression that will be evaluated to obtain attribute9's value to be sent in Assertion.
* `attribute9format` - Format of Attribute9 to be sent in Assertion.
* `attribute9friendlyname` - User-Friendly Name of attribute9 that needs to be sent in SAML Assertion.
* `attribute10` - Name of attribute10 that needs to be sent in SAML Assertion.
* `attribute10expr` - Expression that will be evaluated to obtain attribute10's value to be sent in Assertion.
* `attribute10format` - Format of Attribute10 to be sent in Assertion.
* `attribute10friendlyname` - User-Friendly Name of attribute10 that needs to be sent in SAML Assertion.
* `attribute11` - Name of attribute11 that needs to be sent in SAML Assertion.
* `attribute11expr` - Expression that will be evaluated to obtain attribute11's value to be sent in Assertion.
* `attribute11format` - Format of Attribute11 to be sent in Assertion.
* `attribute11friendlyname` - User-Friendly Name of attribute11 that needs to be sent in SAML Assertion.
* `attribute12` - Name of attribute12 that needs to be sent in SAML Assertion.
* `attribute12expr` - Expression that will be evaluated to obtain attribute12's value to be sent in Assertion.
* `attribute12format` - Format of Attribute12 to be sent in Assertion.
* `attribute12friendlyname` - User-Friendly Name of attribute12 that needs to be sent in SAML Assertion.
* `attribute13` - Name of attribute13 that needs to be sent in SAML Assertion.
* `attribute13expr` - Expression that will be evaluated to obtain attribute13's value to be sent in Assertion.
* `attribute13format` - Format of Attribute13 to be sent in Assertion.
* `attribute13friendlyname` - User-Friendly Name of attribute13 that needs to be sent in SAML Assertion.
* `attribute14` - Name of attribute14 that needs to be sent in SAML Assertion.
* `attribute14expr` - Expression that will be evaluated to obtain attribute14's value to be sent in Assertion.
* `attribute14format` - Format of Attribute14 to be sent in Assertion.
* `attribute14friendlyname` - User-Friendly Name of attribute14 that needs to be sent in SAML Assertion.
* `attribute15` - Name of attribute15 that needs to be sent in SAML Assertion.
* `attribute15expr` - Expression that will be evaluated to obtain attribute15's value to be sent in Assertion.
* `attribute15format` - Format of Attribute15 to be sent in Assertion.
* `attribute15friendlyname` - User-Friendly Name of attribute15 that needs to be sent in SAML Assertion.
* `attribute16` - Name of attribute16 that needs to be sent in SAML Assertion.
* `attribute16expr` - Expression that will be evaluated to obtain attribute16's value to be sent in Assertion.
* `attribute16format` - Format of Attribute16 to be sent in Assertion.
* `attribute16friendlyname` - User-Friendly Name of attribute16 that needs to be sent in SAML Assertion.
* `audience` - Audience for which assertion sent by IdP is applicable. This is typically entity name or url that represents ServiceProvider.
* `digestmethod` - Algorithm to be used to compute/verify digest for SAML transactions.
* `encryptassertion` - Option to encrypt assertion when Citrix ADC sends one.
* `encryptionalgorithm` - Algorithm to be used to encrypt SAML assertion.
* `nameidexpr` - Expression that will be evaluated to obtain NameIdentifier to be sent in assertion.
* `nameidformat` - Format of Name Identifier sent in Assertion.
* `relaystaterule` - Expression to extract relaystate to be sent along with assertion. Evaluation of this expression should return TEXT content. This is typically a target url to which user is redirected after the recipient validates SAML token.
* `samlissuername` - The name to be used in requests sent from Citrix ADC to IdP to uniquely identify Citrix ADC.
* `samlsigningcertname` - Name of the SSL certificate that is used to Sign Assertion.
* `samlspcertname` - Name of the SSL certificate of peer/receving party using which Assertion is encrypted.
* `sendpassword` - Option to send password in assertion.
* `signassertion` - Option to sign portions of assertion when Citrix ADC IDP sends one. Based on the user selection, either Assertion or Response or Both or none can be signed.
* `signaturealg` - Algorithm to be used to sign/verify SAML transactions.
* `skewtime` - This option specifies the number of minutes on either side of current time that the assertion would be valid. For example, if skewTime is 10, then assertion would be valid from (current time - 10) min to (current time + 10) min, ie 20min in all.

## Attribute Reference

* `id` - The id of the tmsamlssoprofile. It has the same value as the `name` attribute.

## Import

A tmsamlssoprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_tmsamlssoprofile.tf_tmsamlssoprofile my_tmsamlssoprofile
```
