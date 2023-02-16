---
subcategory: "SSL"
---

# Resource: sslcert

The sslcert resource is used to create ssl cert file.


## Example usage

```hcl
resource "citrixadc_sslcertreq" "tf_sslcertreq" {
  reqfile          = "/nsconfig/ssl/test-ca.csr"
  keyfile          = "/nsconfig/ssl/key1.pem"
  countryname      = "in"
  statename        = "kar"
  organizationname = "xyz"
}
resource "citrixadc_sslcert" "tf_sslcert" {
  certfile = "/nsconfig/ssl/certificate1.crt"
  reqfile  = "/nsconfig/ssl/test-ca.csr"
  certtype = "ROOT_CERT"
  keyfile  = "/nsconfig/ssl/key1.pem"
  depends_on = [
    citrixadc_sslcertreq.tf_sslcertreq
  ]
}
```


## Argument Reference

* `certfile` - (Required) Name for and, optionally, path to the generated certificate file. /nsconfig/ssl/ is the default path. Maximum length =  63
* `reqfile` - (Required) Name for and, optionally, path to the certificate-signing request (CSR). /nsconfig/ssl/ is the default path. Maximum length =  63
* `certtype` - (Required) Type of certificate to generate. Specify one of the following: * ROOT_CERT - Self-signed Root-CA certificate. You must specify the key file name. The generated Root-CA certificate can be used for signing end-user client or server certificates or to create Intermediate-CA certificates. * INTM_CERT - Intermediate-CA certificate. * CLNT_CERT - End-user client certificate used for client authentication. * SRVR_CERT - SSL server certificate used on SSL servers for end-to-end encryption. Possible values: [ ROOT_CERT, INTM_CERT, CLNT_CERT, SRVR_CERT ]
* `keyfile` - (Optional) Name for and, optionally, path to the private key. You can either use an existing RSA or DSA key that you own or create a new private key on the Citrix ADC. This file is required only when creating a self-signed Root-CA certificate. The key file is stored in the /nsconfig/ssl directory by default. If the input key specified is an encrypted key, you are prompted to enter the PEM pass phrase that was used for encrypting the key. Maximum length =  63
* `keyform` - (Optional) Format in which the key is stored on the appliance. Possible values: [ DER, PEM ]
* `pempassphrase` - (Optional) . Minimum length =  1 Maximum length =  31
* `days` - (Optional) Number of days for which the certificate will be valid, beginning with the time and day (system time) of creation. Minimum value =  1 Maximum value =  3650
* `subjectaltname` - (Optional) Subject Alternative Name (SAN) is an extension to X.509 that allows various values to be associated with a security certificate using a subjectAltName field. These values are called "Subject Alternative Names" (SAN). Names include: 1. Email addresses 2. IP addresses 3. URIs 4. DNS names (This is usually also provided as the Common Name RDN within the Subject field of the main certificate.) 5. directory names (alternative Distinguished Names to that given in the Subject). Minimum length =  1
* `certform` - (Optional) Format in which the certificate is stored on the appliance. Possible values: [ DER, PEM ]
* `cacert` - (Optional) Name of the CA certificate file that issues and signs the Intermediate-CA certificate or the end-user client and server certificates. Maximum length =  63
* `cacertform` - (Optional) Format of the CA certificate. Possible values: [ DER, PEM ]
* `cakey` - (Optional) Private key, associated with the CA certificate that is used to sign the Intermediate-CA certificate or the end-user client and server certificate. If the CA key file is password protected, the user is prompted to enter the pass phrase that was used to encrypt the key. Maximum length =  63
* `cakeyform` - (Optional) Format for the CA certificate. Possible values: [ DER, PEM ]
* `caserial` - (Optional) Serial number file maintained for the CA certificate. This file contains the serial number of the next certificate to be issued or signed by the CA. If the specified file does not exist, a new file is created, with /nsconfig/ssl/ as the default path. If you do not specify a proper path for the existing serial file, a new serial file is created. This might change the certificate serial numbers assigned by the CA certificate to each of the certificates it signs. Maximum length =  63


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslcert. It is a unique string prefixed with "tf-sslcert-"

