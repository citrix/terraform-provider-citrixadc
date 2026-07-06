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

### Using passplain (sensitive attribute - persisted in state)

The `passplain` attribute stores the passphrase in the Terraform state file (encrypted). It is retained for backward compatibility:

```hcl
resource "citrixadc_sslcertkey" "tf_sslcertkey_legacy" {
  certkey = "tf_sslcertkey_legacy"
  cert = "/nsconfig/ssl/certificate1.crt"
  key = "/nsconfig/ssl/encrypted_key1.pem"
  
  # Legacy approach (passphrase stored in state file)
  passplain = "my-secret-passphrase"
}
```

### Using passplain_wo (write-only/ephemeral - NOT persisted in state)

The `passplain_wo` attribute provides an ephemeral path for the private-key pass phrase. The value is sent to the ADC but is **not stored in Terraform state**, reducing the risk of secret exposure. To trigger an update when the pass phrase changes, increment `passplain_wo_version`.

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

To rotate the pass phrase, update the variable value and bump the version:

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

### Automatic Certificate Renewal Detection

This example demonstrates automatic certificate renewal detection using file content hashing. When certificate files are updated, the sslcertkey resource will automatically detect changes and update accordingly:

```hcl
# Create certificate file
resource "citrixadc_systemfile" "cert_file" {
  filename     = "servercert.cert"
  filelocation = "/nsconfig/ssl"
  filecontent  = file("path/to/your/certificate.cert")
}

# Create private key file
resource "citrixadc_systemfile" "key_file" {
  filename     = "servercert.key"
  filelocation = "/nsconfig/ssl"
  filecontent  = file("path/to/your/private.key")
}

variable "sslcertkey_passplain_wo" {
  type      = string
  sensitive = true
  description = "Passphrase for encrypted private key (not stored in state)"
}

# SSL certificate with automatic renewal detection
resource "citrixadc_sslcertkey" "tf_sslcertkey_auto_renewal" {
  certkey   = "tf_sslcertkey_auto_renewal"
  cert      = "/nsconfig/ssl/servercert.cert"
  key       = "/nsconfig/ssl/servercert.key"
  
  # Hash-based change detection for automatic renewal
  cert_hash = sha256(citrixadc_systemfile.cert_file.filecontent)
  key_hash  = sha256(citrixadc_systemfile.key_file.filecontent)
  
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
* `passplain` - (Optional, Sensitive) Pass phrase used to encrypt the private-key. Required when adding an encrypted private-key in PEM format. The value is persisted in Terraform state (encrypted). See also `passplain_wo` for an ephemeral alternative.
* `passplain_wo` - (Optional, Sensitive, WriteOnly) Same as `passplain`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `passplain_wo_version`. If both `passplain` and `passplain_wo` are set, `passplain_wo` takes precedence.
* `passplain_wo_version` - (Optional) An integer version tracker for `passplain_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the value has changed and trigger an update. Defaults to `1`.
* `expirymonitor` - (Optional) Issue an alert when the certificate is about to expire. Possible values: [ ENABLED, DISABLED ]
* `notificationperiod` - (Optional) Time, in number of days, before certificate expiration, at which to generate an alert that the certificate is about to expire.
* `bundle` - (Optional) Parse the certificate chain as a single file after linking the server certificate to its issuer's certificate within the file. Possible values: [ YES, NO ]
* `linkcertkeyname` - (Optional) Name of the Certificate Authority certificate-key pair to which to link a certificate-key pair.
* `nodomaincheck` - (Optional) Override the check for matching domain names during a certificate update operation.
* `ocspstaplingcache` - (Optional) Clear cached ocspStapling response in certkey.
* `deletecertkeyfilesonremoval` - (Optional) This option is used to automatically delete certificate/key files from physical device when the added certkey is removed. When deleteCertKeyFilesOnRemoval option is used at rm certkey command, it overwrites the deleteCertKeyFilesOnRemoval setting used at add/set certkey command
* `deletefromdevice` - (Optional) Delete cert/key file from file system.
* `cert_hash` - (Optional) Hash of the certificate file content. Used internally to detect certificate file changes for automatic renewal. Typically set using `sha256(systemfile.filecontent)`.
* `key_hash` - (Optional) Hash of the private key file content. Used internally to detect key file changes for automatic renewal. Typically set using `sha256(systemfile.filecontent)`.

## Automatic Certificate Renewal Detection

The `cert_hash` and `key_hash` attributes provide automatic certificate renewal detection capability:

### How It Works

1. **Hash Calculation**: Use Terraform's built-in `sha256()` function to calculate hashes of certificate and key file contents
2. **Change Detection**: During `terraform plan/apply`, the system compares current file content hashes with stored state hashes
3. **Automatic Updates**: If hashes differ (indicating file content changed), the sslcertkey resource automatically triggers an update using the Citrix ADC Change API
4. **State Sync**: Updated hashes are stored in Terraform state for future comparisons

### When Hash Attributes Are Needed

**Important Note**: The `cert_hash` and `key_hash` attributes are only needed when the certificate file content changes but the file paths (`cert` and `key` attributes) remain the same. 

- **Use hashes when**: Certificate content is updated but file names/paths stay the same
- **Hashes not needed when**: Certificate renewal involves changing file names or paths (e.g., `certificate-2024.crt` to `certificate-2025.crt`)

In cases where the `cert` and/or `key` file paths change during renewal, Terraform will automatically detect the attribute changes and update the certificate without requiring hash-based detection.



## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslcertkey. It has the same value as the `certkey` attribute.


## Import

A sslcertkey can be imported using its certkey, e.g.

```shell
terraform import citrixadc_sslcertkey.tf_sslcertkey tf_sslcertkey
```
