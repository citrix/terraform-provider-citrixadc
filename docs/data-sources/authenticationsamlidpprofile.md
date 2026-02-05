---
subcategory: "Authentication"
---

# Data Source: citrixadc_authenticationsamlidpprofile

Use this data source to retrieve information about an existing SAML Identity Provider (IdP) profile.

The `citrixadc_authenticationsamlidpprofile` data source allows you to retrieve details of an authentication SAML IdP profile by its name. This is useful for referencing existing profiles in your Terraform configurations without managing them directly.

## Example usage

```hcl
# Retrieve an existing authenticationsamlidpprofile
data "citrixadc_authenticationsamlidpprofile" "example" {
  name = "my_samlidpprofile"
}

# Use the retrieved profile data
output "saml_issuer_name" {
  value = data.citrixadc_authenticationsamlidpprofile.example.samlissuername
}

output "assertion_consumer_service_url" {
  value = data.citrixadc_authenticationsamlidpprofile.example.assertionconsumerserviceurl
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the SAML single sign-on profile to retrieve. Must match an existing profile name.

## Attribute Reference

In addition to the argument above, the following attributes are exported:

* `id` - The ID of the SAML IdP profile (same as name).

### Basic Configuration

* `samlissuername` - The name to be used in requests sent from Citrix ADC to IdP to uniquely identify Citrix ADC.
* `assertionconsumerserviceurl` - URL to which the assertion is to be sent.
* `serviceproviderid` - Unique identifier of the Service Provider that sends SAML Request. Citrix ADC will ensure that the Issuer of the SAML Request matches this URI.
* `audience` - Audience for which assertion sent by IdP is applicable. This is typically entity name or url that represents ServiceProvider.

### Certificate Configuration

* `samlspcertname` - Name of the SSL certificate of SAML Relying Party. This certificate is used to verify signature of the incoming AuthnRequest from a Relying Party or Service Provider.
* `samlspcertversion` - Version of the certificate in signature service used to verify the signature of the incoming AuthnRequest from a Relying Party or Service Provider.
* `samlidpcertname` - Name of the certificate used to sign the SAMLResposne that is sent to Relying Party or Service Provider after successful authentication.
* `samlsigningcertversion` - Version of the certificate in signature service used to sign the SAMLResposne that is sent to Relying Party or Service Provider after successful authentication.

### Security Settings

* `rejectunsignedrequests` - Option to Reject unsigned SAML Requests. ON option denies any authentication requests that arrive without signature.
* `signaturealg` - Algorithm to be used to sign/verify SAML transactions.
* `digestmethod` - Algorithm to be used to compute/verify digest for SAML transactions.
* `signassertion` - Option to sign portions of assertion when Citrix ADC IDP sends one. Based on the user selection, either Assertion or Response or Both or none can be signed.

### Encryption Settings

* `encryptassertion` - Option to encrypt assertion when Citrix ADC IDP sends one.
* `encryptionalgorithm` - Algorithm to be used to encrypt SAML assertion.
* `keytransportalg` - Key transport algorithm to be used in encryption of SAML assertion.

### Binding and Transport

* `samlbinding` - This element specifies the transport mechanism of saml messages.
* `logoutbinding` - This element specifies the transport mechanism of saml logout messages.

### Name Identifier

* `nameidformat` - Format of Name Identifier sent in Assertion.
* `nameidexpr` - Expression that will be evaluated to obtain NameIdentifier to be sent in assertion.

### Attributes (1-16)

The profile supports up to 16 custom attributes that can be sent in SAML Assertions:

* `attribute1` - Name of attribute1 that needs to be sent in SAML Assertion.
* `attribute1expr` - Expression that will be evaluated to obtain attribute1's value to be sent in Assertion.
* `attribute1format` - Format of Attribute1 to be sent in Assertion.
* `attribute1friendlyname` - User-Friendly Name of attribute1 that needs to be sent in SAML Assertion.

Similar patterns apply for `attribute2` through `attribute16` with their respective expr, format, and friendlyname variants.

### ACS URL Rule

* `acsurlrule` - Expression that will be evaluated to allow Assertion Consumer Service URI coming in the SAML Request.

### Logout Configuration

* `splogouturl` - Endpoint on the ServiceProvider (SP) to which logout messages are to be sent.

### Metadata

* `metadataurl` - This URL is used for obtaining samlidp metadata.
* `metadatarefreshinterval` - Interval in minute for fetching metadata from specified metadata URL.

### Additional Settings

* `defaultauthenticationgroup` - This group will be part of AAA session's internal group list. This is helpful to admin in Nfactor flow to decide right AAA configuration for Relaying Party.
* `skewtime` - This option specifies the number of minutes on either side of current time that the assertion would be valid. For example, if skewTime is 10, then assertion would be valid from (current time - 10) min to (current time + 10) min, ie 20min in all.
* `sendpassword` - Option to send password in assertion.
* `signatureservice` - Name of the service in cloud used to sign the data.

## Notes

* The datasource retrieves read-only information about existing SAML IdP profiles.
* All attributes marked as Optional in the schema are computed from the NetScaler ADC configuration.
* The profile must exist on the NetScaler ADC before it can be retrieved using this datasource.
