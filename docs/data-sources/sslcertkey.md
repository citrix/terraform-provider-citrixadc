---
subcategory: "SSL"
---

# Data Source: sslcertkey

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

# Read-only certificate metadata returned by the appliance
output "days_to_expiration" {
  value = data.citrixadc_sslcertkey.tf_sslcertkey.daystoexpiration
}

output "subject" {
  value = data.citrixadc_sslcertkey.tf_sslcertkey.subject
}

output "status" {
  value = data.citrixadc_sslcertkey.tf_sslcertkey.status
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
* `id` - The id of the sslcertkey. It has the same value as the `certkey` attribute.

### Read-only certificate metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_sslcertkey` resource). Any attribute the appliance does not return is `null`.

* `signaturealg` - Algorithm used to sign the certificate.
* `certificatetype` - Certificate type(s) (for example `ROOT_CERT`, `INTM_CERT`, `SRVR_CERT`, `CLNT_CERT`). A list of strings.
* `serial` - Serial number of the certificate.
* `issuer` - Distinguished name of the certificate issuer.
* `clientcertnotbefore` - Date and time from which the certificate is valid.
* `clientcertnotafter` - Date and time after which the certificate expires.
* `daystoexpiration` - Number of days remaining before the certificate expires.
* `subject` - Distinguished name of the certificate subject.
* `publickey` - Public key algorithm of the certificate.
* `publickeysize` - Public key size, in bits.
* `version` - Version number of the certificate.
* `priority` - Priority of the certificate.
* `status` - Validity status of the certificate (for example `Valid`, `Expired`).
* `passcrypt` - (Sensitive) Encrypted passphrase of the private-key as stored on the appliance.
* `data` - Number of references to the certificate-key pair.
* `servicename` - Service name to which the certificate is bound.
* `sandns` - Subject Alternative Name (DNS) entries of the certificate.
* `sanipadd` - Subject Alternative Name (IP address) entries of the certificate.
* `ocspresponsestatus` - OCSP response status for the certificate.
* `builtin` - Whether the certificate-key pair is built-in. A list of strings.
* `feature` - The feature to be checked while applying this configuration.
* `certkeydigest` - Digest (fingerprint) of the certificate.
* `certificatesource` - Source of the certificate.
* `certkeystatus` - Status of the certificate-key pair.


## Import

A sslcertkey can be imported using its certkey, e.g.

```shell
terraform import citrixadc_sslcertkey.tf_sslcertkey tf_sslcertkey
```
