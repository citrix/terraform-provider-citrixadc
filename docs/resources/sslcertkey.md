---
subcategory: "SSL"
---

# Resource: sslcertkey

The sslcertkey resource is used to create TLS certificate keys.


## Example usage

```hcl
resource "citrixadc_sslcertkey" "tf_sslcertkey" {
  certkey = "tf_sslcertkey"
  cert = "/nsconfig/ssl/certificate1.crt"
  key = "/nsconfig/ssl/key1.pem"
  notificationperiod = 40
  expirymonitor = "ENABLED"
}
```

## Example usage 2

```hcl

variable "sslcertkey_passplain_wo" {
  type      = string
  sensitive = true
}

resource "citrixadc_sslcertkey" "tf_sslcertkey" {
  certkey = "tf_sslcertkey"
  cert = "/nsconfig/ssl/certificate1.crt"
  key = "/nsconfig/ssl/key1.pem"
  notificationperiod = 40
  expirymonitor = "ENABLED"
  passplain_wo = var.sslcertkey_passplain_wo
  passplain_wo_version = 1
}

```


## Argument Reference

* `certkey` - (Required) Name for the certificate and private-key pair.
* `cert` - (Required) Name of and, optionally, path to the X509 certificate file that is used to form the certificate-key pair. The certificate file should be present on the appliance's hard-disk drive or solid-state drive. Storing a certificate in any location other than the default might cause inconsistency in a high availability setup. /nsconfig/ssl/ is the default path.
* `key` - (Optional) Name of and, optionally, path to the private-key file that is used to form the certificate-key pair. The certificate file should be present on the appliance's hard-disk drive or solid-state drive. Storing a certificate in any location other than the default might cause inconsistency in a high availability setup. /nsconfig/ssl/ is the default path.
* `password` - (Optional) Passphrase that was used to encrypt the private-key. Use this option to load encrypted private-keys in PEM format.
* `fipskey` - (Optional) Name of the FIPS key that was created inside the Hardware Security Module (HSM) of a FIPS appliance, or a key that was imported into the HSM.
* `hsmkey` - (Optional) Name of the HSM key that was created in the External Hardware Security Module (HSM) of a FIPS appliance.
* `inform` - (Optional) Input format of the certificate and the private-key files. The three formats supported by the appliance are: PEM - Privacy Enhanced Mail DER - Distinguished Encoding Rule PFX - Personal Information Exchange. Possible values: [ DER, PEM, PFX ]
* `passplain` - (Optional) Pass phrase used to encrypt the private-key. Required when adding an encrypted private-key in PEM format.
* `expirymonitor` - (Optional) Issue an alert when the certificate is about to expire. Possible values: [ ENABLED, DISABLED ]
* `notificationperiod` - (Optional) Time, in number of days, before certificate expiration, at which to generate an alert that the certificate is about to expire.
* `bundle` - (Optional) Parse the certificate chain as a single file after linking the server certificate to its issuer's certificate within the file. Possible values: [ YES, NO ]
* `linkcertkeyname` - (Optional) Name of the Certificate Authority certificate-key pair to which to link a certificate-key pair.
* `nodomaincheck` - (Optional) Override the check for matching domain names during a certificate update operation.
* `ocspstaplingcache` - (Optional) Clear cached ocspStapling response in certkey.
* `deletecertkeyfilesonremoval` - (Optional) This option is used to automatically delete certificate/key files from physical device when the added certkey is removed. When deleteCertKeyFilesOnRemoval option is used at rm certkey command, it overwrites the deleteCertKeyFilesOnRemoval setting used at add/set certkey command
* `deletefromdevice` - (Optional) Delete cert/key file from file system.
* `passplain_wo` - (Optional, Write-Only required unless passplain is provided) Pass phrase used to encrypt the private-key. Required when adding an encrypted private-key in PEM format. Note that this may show up in logs, and it will not be stored in the state file.
* `passplain_wo_version` - (Optional) Used together with passplain_wo to trigger an update. Increment this value when an update to passplain_wo is required.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslcertkey. It has the same value as the `certkey` attribute.


## Import

A sslcertkey can be imported using its certkey, e.g.

```shell
terraform import citrixadc_sslcertkey.tf_sslcertkey tf_sslcertkey
```
