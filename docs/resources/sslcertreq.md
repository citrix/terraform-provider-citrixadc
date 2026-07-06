---
subcategory: "SSL"
---

# Resource: sslcertreq

The sslcertreq resource is used to create ssl certreq file.


## Example usage

### Basic usage

```hcl
resource "citrixadc_sslcertreq" "tf_sslcertreq" {
  reqfile          = "/nsconfig/ssl/test-ca.csr"
  keyfile          = "/nsconfig/ssl/key1.pem"
  countryname      = "in"
  statename        = "kar"
  organizationname = "xyz"
}
```

### Using pempassphrase (sensitive attribute - persisted in state)

```hcl
variable "sslcertreq_pempassphrase" {
  type      = string
  sensitive = true
}

resource "citrixadc_sslcertreq" "tf_sslcertreq" {
  reqfile          = "/nsconfig/ssl/test-ca.csr"
  keyfile          = "/nsconfig/ssl/key1.pem"
  countryname      = "in"
  statename        = "kar"
  organizationname = "xyz"
  pempassphrase    = var.sslcertreq_pempassphrase
}
```

### Using pempassphrase_wo (write-only/ephemeral - NOT persisted in state)

The `pempassphrase_wo` attribute provides an ephemeral path for the PEM pass phrase. The value is sent to the ADC but is **not stored in Terraform state**, reducing the risk of secret exposure. To change the value, increment `pempassphrase_wo_version`; because the secret is immutable on the ADC, this **destroys and recreates** the resource.

```hcl
variable "sslcertreq_pempassphrase" {
  type      = string
  sensitive = true
}

resource "citrixadc_sslcertreq" "tf_sslcertreq" {
  reqfile                  = "/nsconfig/ssl/test-ca.csr"
  keyfile                  = "/nsconfig/ssl/key1.pem"
  countryname              = "in"
  statename                = "kar"
  organizationname         = "xyz"
  pempassphrase_wo         = var.sslcertreq_pempassphrase
  pempassphrase_wo_version = 1
}
```

### Using challengepassword (sensitive attribute - persisted in state)

```hcl
variable "sslcertreq_challengepassword" {
  type      = string
  sensitive = true
}

resource "citrixadc_sslcertreq" "tf_sslcertreq" {
  reqfile             = "/nsconfig/ssl/test-ca.csr"
  keyfile             = "/nsconfig/ssl/key1.pem"
  countryname         = "in"
  statename           = "kar"
  organizationname    = "xyz"
  challengepassword   = var.sslcertreq_challengepassword
}
```

### Using challengepassword_wo (write-only/ephemeral - NOT persisted in state)

The `challengepassword_wo` attribute provides an ephemeral path for the challenge password. The value is sent to the ADC but is **not stored in Terraform state**, reducing the risk of secret exposure. To change the value, increment `challengepassword_wo_version`; because the secret is immutable on the ADC, this **destroys and recreates** the resource.

```hcl
variable "sslcertreq_challengepassword" {
  type      = string
  sensitive = true
}

resource "citrixadc_sslcertreq" "tf_sslcertreq" {
  reqfile                      = "/nsconfig/ssl/test-ca.csr"
  keyfile                      = "/nsconfig/ssl/key1.pem"
  countryname                  = "in"
  statename                    = "kar"
  organizationname             = "xyz"
  challengepassword_wo         = var.sslcertreq_challengepassword
  challengepassword_wo_version = 1
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
* `pempassphrase` - (Optional, Sensitive) Pass phrase used to encrypt the private key. The value is persisted in Terraform state (encrypted). See also `pempassphrase_wo` for an ephemeral alternative. Minimum length =  1 Maximum length =  31
* `pempassphrase_wo` - (Optional, Sensitive, WriteOnly) Same as `pempassphrase`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `pempassphrase_wo_version`. If both `pempassphrase` and `pempassphrase_wo` are set, `pempassphrase_wo` takes precedence.
* `pempassphrase_wo_version` - (Optional) An integer version tracker for `pempassphrase_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the value has changed. Note: this secret is immutable on the ADC, so changing `pempassphrase_wo_version` (or `pempassphrase`/`pempassphrase_wo`) forces the resource to be **destroyed and recreated** rather than updated in place. Defaults to `1`.
* `organizationunitname` - (Optional) Name of the division or section in the organization that will use the certificate. Minimum length =  1
* `localityname` - (Optional) Name of the city or town in which your organization's head office is located. Minimum length =  1
* `commonname` - (Optional) Fully qualified domain name for the company or web site. The common name must match the name used by DNS servers to do a DNS lookup of your server. Most browsers use this information for authenticating the server's certificate during the SSL handshake. If the server name in the URL does not match the common name as given in the server certificate, the browser terminates the SSL handshake or prompts the user with a warning message. Do not use wildcard characters, such as asterisk (*) or question mark (?), and do not use an IP address as the common name. The common name must not contain the protocol specifier <http://> or <https://>. Minimum length =  1
* `emailaddress` - (Optional) Contact person's e-mail address. This address is publically displayed as part of the certificate. Provide an e-mail address that is monitored by an administrator who can be contacted about the certificate. Minimum length =  1
* `challengepassword` - (Optional, Sensitive) Pass phrase, embedded in the certificate signing request that is shared only between the client or server requesting the certificate and the SSL certificate issuer (typically the certificate authority). This pass phrase can be used to authenticate a client or server that is requesting a certificate from the certificate authority. The value is persisted in Terraform state (encrypted). See also `challengepassword_wo` for an ephemeral alternative. Minimum length =  4
* `challengepassword_wo` - (Optional, Sensitive, WriteOnly) Same as `challengepassword`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `challengepassword_wo_version`. If both `challengepassword` and `challengepassword_wo` are set, `challengepassword_wo` takes precedence.
* `challengepassword_wo_version` - (Optional) An integer version tracker for `challengepassword_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the value has changed. Note: this secret is immutable on the ADC, so changing `challengepassword_wo_version` (or `challengepassword`/`challengepassword_wo`) forces the resource to be **destroyed and recreated** rather than updated in place. Defaults to `1`.
* `companyname` - (Optional) Additional name for the company or web site. Minimum length =  1
* `digestmethod` - (Optional) Digest algorithm used in creating CSR. Possible values: [ SHA1, SHA256 ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslcertreq. It is a unique string prefixed with "tf-sslcertreq-"

