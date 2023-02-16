---
subcategory: "SSL"
---

# Resource: sslcertreq

The sslcertreq resource is used to create ssl certreq file.


## Example usage

```hcl
resource "citrixadc_sslcertreq" "tf_sslcertreq" {
  reqfile          = "/nsconfig/ssl/test-ca.csr"
  keyfile          = "/nsconfig/ssl/key1.pem"
  countryname      = "in"
  statename        = "kar"
  organizationname = "xyz"
}
```


## Argument Reference

* `reqfile` - (Required) Name for and, optionally, path to the certificate signing request (CSR). /nsconfig/ssl/ is the default path. Maximum length =  63
* `countryname` - (Required) Two letter ISO code for your country. For example, US for United States. Minimum length =  2 Maximum length =  2
* `organizationname` - (Required) Name of the organization that will use this certificate. The organization name (corporation, limited partnership, university, or government agency) must be registered with some authority at the national, state, or city level. Use the legal name under which the organization is registered. Do not abbreviate the organization name and do not use the following characters in the name: Angle brackets (< >) tilde (~), exclamation mark, at (@), pound (#), zero (0), caret (^), asterisk (*), forward slash (/), square brackets ([ ]), question mark (?). Minimum length =  1
* `statename` - (Required) Full name of the state or province where your organization is located. Do not abbreviate. Minimum length =  1
* `keyfile` - (Optional) Name of and, optionally, path to the private key used to create the certificate signing request, which then becomes part of the certificate-key pair. The private key can be either an RSA or a DSA key. The key must be present in the appliance's local storage. /nsconfig/ssl is the default path. Maximum length =  63
* `subjectaltname` - (Optional) Subject Alternative Name (SAN) is an extension to X.509 that allows various values to be associated with a security certificate using a subjectAltName field. These values are called "Subject Alternative Names" (SAN). Names include: 1. Email addresses 2. IP addresses 3. URIs 4. DNS names (this is usually also provided as the Common Name RDN within the Subject field of the main certificate.) 5. Directory names (alternative Distinguished Names to that given in the Subject). Minimum length =  1
* `fipskeyname` - (Optional) Name of the FIPS key used to create the certificate signing request. FIPS keys are created inside the Hardware Security Module of the FIPS card. Minimum length =  1 Maximum length =  31
* `keyform` - (Optional) Format in which the key is stored on the appliance. Possible values: [ DER, PEM ]
* `pempassphrase` - (Optional) . Minimum length =  1 Maximum length =  31
* `organizationunitname` - (Optional) Name of the division or section in the organization that will use the certificate. Minimum length =  1
* `localityname` - (Optional) Name of the city or town in which your organization's head office is located. Minimum length =  1
* `commonname` - (Optional) Fully qualified domain name for the company or web site. The common name must match the name used by DNS servers to do a DNS lookup of your server. Most browsers use this information for authenticating the server's certificate during the SSL handshake. If the server name in the URL does not match the common name as given in the server certificate, the browser terminates the SSL handshake or prompts the user with a warning message. Do not use wildcard characters, such as asterisk (*) or question mark (?), and do not use an IP address as the common name. The common name must not contain the protocol specifier <http://> or <https://>. Minimum length =  1
* `emailaddress` - (Optional) Contact person's e-mail address. This address is publically displayed as part of the certificate. Provide an e-mail address that is monitored by an administrator who can be contacted about the certificate. Minimum length =  1
* `challengepassword` - (Optional) Pass phrase, embedded in the certificate signing request that is shared only between the client or server requesting the certificate and the SSL certificate issuer (typically the certificate authority). This pass phrase can be used to authenticate a client or server that is requesting a certificate from the certificate authority. Minimum length =  4
* `companyname` - (Optional) Additional name for the company or web site. Minimum length =  1
* `digestmethod` - (Optional) Digest algorithm used in creating CSR. Possible values: [ SHA1, SHA256 ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslcertreq. It is a unique string prefixed with "tf-sslcertreq-"

