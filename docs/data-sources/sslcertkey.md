---
subcategory: "SSL"
---

# Data Source `sslcertkey`

The sslcertkey data source allows you to retrieve information about the TLS certificate keys.


## Example usage

```terraform
data "citrixadc_sslcertkey" "tf_sslcertkey" {
  certkey = "servercert1"
}

output "cert" {
  value = data.citrixadc_sslcertkey.tf_sslcertkey.cert
}

output "key" {
  value = data.citrixadc_sslcertkey.tf_sslcertkey.key
}
```


## Argument Reference

* `certkey` - (Required) Name for the certificate and private-key pair.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `cert` -  Name of and, optionally, path to the X509 certificate file that is used to form the certificate-key pair. The certificate file should be present on the appliance's hard-disk drive or solid-state drive. Storing a certificate in any location other than the default might cause inconsistency in a high availability setup. /nsconfig/ssl/ is the default path.
* `key` -  Name of and, optionally, path to the private-key file that is used to form the certificate-key pair. The certificate file should be present on the appliance's hard-disk drive or solid-state drive. Storing a certificate in any location other than the default might cause inconsistency in a high availability setup. /nsconfig/ssl/ is the default path.
* `password` -  Passphrase that was used to encrypt the private-key. Use this option to load encrypted private-keys in PEM format.
* `fipskey` -  Name of the FIPS key that was created inside the Hardware Security Module (HSM) of a FIPS appliance, or a key that was imported into the HSM.
* `hsmkey` -  Name of the HSM key that was created in the External Hardware Security Module (HSM) of a FIPS appliance.
* `inform` -  Input format of the certificate and the private-key files. The three formats supported by the appliance are: PEM - Privacy Enhanced Mail DER - Distinguished Encoding Rule PFX - Personal Information Exchange. Possible values: [ DER, PEM, PFX ]
* `passplain` -  Pass phrase used to encrypt the private-key. Required when adding an encrypted private-key in PEM format.
* `expirymonitor` -  Issue an alert when the certificate is about to expire. Possible values: [ ENABLED, DISABLED ]
* `notificationperiod` -  Time, in number of days, before certificate expiration, at which to generate an alert that the certificate is about to expire.
* `bundle` -  Parse the certificate chain as a single file after linking the server certificate to its issuer's certificate within the file. Possible values: [ YES, NO ]
* `linkcertkeyname` -  Name of the Certificate Authority certificate-key pair to which to link a certificate-key pair.
* `nodomaincheck` -  Override the check for matching domain names during a certificate update operation.
* `ocspstaplingcache` -  Clear cached ocspStapling response in certkey.
* `deletecertkeyfilesonremoval` -  This option is used to automatically delete certificate/key files from physical device when the added certkey is removed. When deleteCertKeyFilesOnRemoval option is used at rm certkey command, it overwrites the deleteCertKeyFilesOnRemoval setting used at add/set certkey command
* `deletefromdevice` -  Delete cert/key file from file system.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslcertkey. It has the same value as the `certkey` attribute.


## Import

A sslcertkey can be imported using its certkey, e.g.

```shell
terraform import citrixadc_sslcertkey.tf_sslcertkey tf_sslcertkey
```
