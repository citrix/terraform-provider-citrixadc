---
subcategory: "SSL"
---

# Resource: sslcertkey_update

The sslcertkey_update resource is used to update TLS certificate and key.


## Example usage

```hcl
resource "citrixadc_sslcertkey_update" "tf_sslcertkey_update" {
  certkey = "tf_sslcertkey_update"
  cert    = "/nsconfig/ssl/certificate3.crt"
  key     = "/nsconfig/ssl/key3.pem"
}
```


## Argument Reference

* `certkey` - (Required) Name for the certificate and private-key pair.
* `cert` - (Optional) Name of and, optionally, path to the X509 certificate file that is used to form the certificate-key pair. The certificate file should be present on the appliance's hard-disk drive or solid-state drive. Storing a certificate in any location other than the default might cause inconsistency in a high availability setup. /nsconfig/ssl/ is the default path.
* `key` - (Optional) Name of and, optionally, path to the private-key file that is used to form the certificate-key pair. The certificate file should be present on the appliance's hard-disk drive or solid-state drive. Storing a certificate in any location other than the default might cause inconsistency in a high availability setup. /nsconfig/ssl/ is the default path.
* `password` - (Optional) Passphrase that was used to encrypt the private-key. Use this option to load encrypted private-keys in PEM format.
* `fipskey` - (Optional) Name of the FIPS key that was created inside the Hardware Security Module (HSM) of a FIPS appliance, or a key that was imported into the HSM.
* `inform` - (Optional) Input format of the certificate and the private-key files. The three formats supported by the appliance are: PEM - Privacy Enhanced Mail DER - Distinguished Encoding Rule PFX - Personal Information Exchange. Possible values: [ DER, PEM, PFX ]
* `passplain` - (Optional) Pass phrase used to encrypt the private-key. Required when adding an encrypted private-key in PEM format.
* `nodomaincheck` - (Optional) Override the check for matching domain names during a certificate update operation.
* `linkcertkeyname` - (Optional) Name of the Certificate Authority certificate-key pair to which to link a certificate-key pair.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslcertkey_update. It has the same value as the `certkey` attribute.


