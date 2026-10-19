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

### Using passplain_wo (write-only/ephemeral - NOT persisted in state)

When the updated private-key is an encrypted PEM file, the pass phrase can be supplied through the `passplain_wo` attribute. The value is sent to the ADC but is **not stored in Terraform state**, reducing the risk of secret exposure. To re-apply the update when the pass phrase changes, increment `passplain_wo_version`.

```hcl
variable "sslcertkey_update_passplain_wo" {
  type        = string
  sensitive   = true
  description = "Passphrase for the encrypted private key (not stored in state)"
}

resource "citrixadc_sslcertkey_update" "tf_sslcertkey_update_encrypted" {
  certkey = "tf_sslcertkey_update_encrypted"
  cert    = "/nsconfig/ssl/certificate3.crt"
  key     = "/nsconfig/ssl/encrypted_key3.pem"

  # Ephemeral passphrase (write-only, not stored in state)
  passplain_wo         = var.sslcertkey_update_passplain_wo
  passplain_wo_version = 1
}
```

To rotate the pass phrase, update the variable value and bump the version:

```hcl
resource "citrixadc_sslcertkey_update" "tf_sslcertkey_update_encrypted" {
  certkey = "tf_sslcertkey_update_encrypted"
  cert    = "/nsconfig/ssl/certificate3.crt"
  key     = "/nsconfig/ssl/encrypted_key3.pem"

  # Re-apply the update by incrementing the version
  passplain_wo         = var.sslcertkey_update_passplain_wo
  passplain_wo_version = 2  # Incremented from 1 to trigger the update
}
```


## Argument Reference

* `certkey` - (Required) Name for the certificate and private-key pair.
* `cert` - (Optional) Name of and, optionally, path to the X509 certificate file that is used to form the certificate-key pair. The certificate file should be present on the appliance's hard-disk drive or solid-state drive. Storing a certificate in any location other than the default might cause inconsistency in a high availability setup. /nsconfig/ssl/ is the default path.
* `key` - (Optional) Name of and, optionally, path to the private-key file that is used to form the certificate-key pair. The certificate file should be present on the appliance's hard-disk drive or solid-state drive. Storing a certificate in any location other than the default might cause inconsistency in a high availability setup. /nsconfig/ssl/ is the default path.
* `password` - (Optional) Passphrase that was used to encrypt the private-key. Use this option to load encrypted private-keys in PEM format.
* `fipskey` - (Optional) Name of the FIPS key that was created inside the Hardware Security Module (HSM) of a FIPS appliance, or a key that was imported into the HSM.
* `inform` - (Optional) Input format of the certificate and the private-key files. The three formats supported by the appliance are: PEM - Privacy Enhanced Mail DER - Distinguished Encoding Rule PFX - Personal Information Exchange. Possible values: [ DER, PEM, PFX ]
* `passplain` - (Optional, Sensitive) Pass phrase used to encrypt the private-key. Required when updating to an encrypted private-key in PEM format. The value is persisted in Terraform state (encrypted). See also `passplain_wo` for an ephemeral alternative.
* `passplain_wo` - (Optional, Sensitive, WriteOnly) Same as `passplain`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `passplain_wo_version`. If both `passplain` and `passplain_wo` are set, `passplain` takes precedence.
* `passplain_wo_version` - (Optional) An integer version tracker for `passplain_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the value has changed and re-apply the update. Defaults to `1`.
* `nodomaincheck` - (Optional) Override the check for matching domain names during a certificate update operation.
* `linkcertkeyname` - (Optional) Name of the Certificate Authority certificate-key pair to which to link a certificate-key pair.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslcertkey_update. It has the same value as the `certkey` attribute.