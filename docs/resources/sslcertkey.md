---
subcategory: "SSL"
---

# Resource: sslcertkey

The sslcertkey resource is used to create TLS certificate keys.


## Example Usage

### Basic SSL Certificate without Passphrase

```hcl
resource "citrixadc_sslcertkey" "tf_sslcertkey" {
  certkey = "tf_sslcertkey"
  cert = "/nsconfig/ssl/certificate1.crt"
  key = "/nsconfig/ssl/key1.pem"
  notificationperiod = 40
  expirymonitor = "ENABLED"
}
```

### Using Legacy `passplain` (Backward Compatibility)

The `passplain` attribute is maintained for backward compatibility but stores the passphrase in the state file:

```hcl
resource "citrixadc_sslcertkey" "tf_sslcertkey_legacy" {
  certkey = "tf_sslcertkey_legacy"
  cert = "/nsconfig/ssl/certificate1.crt"
  key = "/nsconfig/ssl/encrypted_key1.pem"
  
  # Legacy approach (passphrase stored in state file)
  passplain = "my-secret-passphrase"
}
```

### SSL Certificate with Ephemeral Passphrase Support

This example demonstrates using `passplain_wo` (write-only) for enhanced security. The passphrase is not stored in the Terraform state file.

```hcl
variable "sslcertkey_passplain_wo" {
  type      = string
  sensitive = true
  description = "Passphrase for encrypted private key (not stored in state)"
}

resource "citrixadc_sslcertkey" "tf_sslcertkey_encrypted" {
  certkey = "tf_sslcertkey_encrypted"
  cert = "/nsconfig/ssl/certificate1.crt"
  key = "/nsconfig/ssl/encrypted_key1.pem"
  notificationperiod = 40
  expirymonitor = "ENABLED"
  
  # Ephemeral passphrase (write-only, not stored in state)
  passplain_wo = var.sslcertkey_passplain_wo
  passplain_wo_version = 1
}
```

### Updating the Passphrase

To update the passphrase for an encrypted private key, increment the `passplain_wo_version` value:

```hcl
variable "sslcertkey_passplain_wo" {
  type      = string
  sensitive = true
  description = "New passphrase for encrypted private key"
}

resource "citrixadc_sslcertkey" "tf_sslcertkey_encrypted" {
  certkey = "tf_sslcertkey_encrypted"
  cert = "/nsconfig/ssl/certificate1.crt"
  key = "/nsconfig/ssl/encrypted_key1.pem"
  notificationperiod = 40
  expirymonitor = "ENABLED"
  
  # Update passphrase by incrementing version
  passplain_wo = var.sslcertkey_passplain_wo
  passplain_wo_version = 2  # Incremented from 1 to trigger update
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
* `passplain` - (Optional) Pass phrase used to encrypt the private-key. Required when adding an encrypted private-key in PEM format. **Note:** This value is stored in the Terraform state file. For enhanced security, use `passplain_wo` instead.
* `passplain_wo` - (Optional, Write-Only) Pass phrase used to encrypt the private-key. Required when adding an encrypted private-key in PEM format. **Recommended over `passplain`** for security reasons as this value is not stored in the Terraform state file. Must be used together with `passplain_wo_version`.
* `passplain_wo_version` - (Optional) Version counter used to trigger updates when `passplain_wo` changes. Increment this value whenever you need to update the passphrase. Default: 1
* `expirymonitor` - (Optional) Issue an alert when the certificate is about to expire. Possible values: [ ENABLED, DISABLED ]
* `notificationperiod` - (Optional) Time, in number of days, before certificate expiration, at which to generate an alert that the certificate is about to expire.
* `bundle` - (Optional) Parse the certificate chain as a single file after linking the server certificate to its issuer's certificate within the file. Possible values: [ YES, NO ]
* `linkcertkeyname` - (Optional) Name of the Certificate Authority certificate-key pair to which to link a certificate-key pair.
* `nodomaincheck` - (Optional) Override the check for matching domain names during a certificate update operation.
* `ocspstaplingcache` - (Optional) Clear cached ocspStapling response in certkey.
* `deletecertkeyfilesonremoval` - (Optional) This option is used to automatically delete certificate/key files from physical device when the added certkey is removed. When deleteCertKeyFilesOnRemoval option is used at rm certkey command, it overwrites the deleteCertKeyFilesOnRemoval setting used at add/set certkey command
* `deletefromdevice` - (Optional) Delete cert/key file from file system.

## Ephemeral Passphrase Support

The `passplain_wo` (write-only) and `passplain_wo_version` attributes provide ephemeral passphrase support for enhanced security:

### Why Use Ephemeral Passphrases?

- **Security**: The passphrase is never stored in the Terraform state file
- **Compliance**: Meets security requirements that prohibit storing secrets in state
- **Best Practice**: Follows Terraform's recommended pattern for sensitive values that should not persist

### How It Works

1. **Initial Creation**: Provide the passphrase via `passplain_wo` and set `passplain_wo_version` to 1
2. **Updating Passphrase**: Change the passphrase value and increment `passplain_wo_version` (e.g., to 2)
3. **Version Tracking**: Terraform uses the version change to detect when the passphrase needs updating

### Important Notes

- `passplain_wo` and `passplain_wo_version` must be used together
- The passphrase is sent to the Citrix ADC but never stored in Terraform state
- For backward compatibility, `passplain` is still supported but stores the value in state
- Do not use both `passplain` and `passplain_wo` simultaneously; choose one approach


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslcertkey. It has the same value as the `certkey` attribute.


## Import

A sslcertkey can be imported using its certkey, e.g.

```shell
terraform import citrixadc_sslcertkey.tf_sslcertkey tf_sslcertkey
```
