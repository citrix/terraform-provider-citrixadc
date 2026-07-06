---
subcategory: "DNS"
---

# Resource: dnskey

The dnskey resource is used to create DNS key.


## Example usage

### Basic usage

```hcl
resource "citrixadc_dnskey" "example" {
  keyname            = "adckey_1"
  publickey          = "/nsconfig/dns/demo.key"
  privatekey         = "/nsconfig/dns/demo.private"
  expires            = 120
  units1             = "DAYS"
  notificationperiod = 7
  units2             = "DAYS"
  ttl                = 3600
}
```

### Using password (sensitive attribute - persisted in state)

```hcl
variable "dnskey_password" {
  type      = string
  sensitive = true
}

resource "citrixadc_dnskey" "example" {
  keyname    = "adckey_1"
  publickey  = "/nsconfig/dns/demo.key"
  privatekey = "/nsconfig/dns/demo.private"
  password   = var.dnskey_password
}
```

### Using password_wo (write-only/ephemeral - NOT persisted in state)

The `password_wo` attribute provides an ephemeral path for the passphrase used to read the encrypted public/private DNS keys. The value is sent to the Citrix ADC but is **not stored in Terraform state**, reducing the risk of secret exposure. To change the value, increment `password_wo_version`; because the secret is immutable on the ADC, this **destroys and recreates** the resource.

```hcl
variable "dnskey_password" {
  type      = string
  sensitive = true
}

resource "citrixadc_dnskey" "example" {
  keyname             = "adckey_1"
  publickey           = "/nsconfig/dns/demo.key"
  privatekey          = "/nsconfig/dns/demo.private"
  password_wo         = var.dnskey_password
  password_wo_version = 1
}
```

To rotate the secret, update the variable value and bump the version:

```hcl
resource "citrixadc_dnskey" "example" {
  keyname             = "adckey_1"
  publickey           = "/nsconfig/dns/demo.key"
  privatekey          = "/nsconfig/dns/demo.private"
  password_wo         = var.dnskey_password
  password_wo_version = 2  # Bumped: forces destroy & recreate
}
```


## Argument Reference

* `keyname` - (Required) Name of the public-private key pair to publish in the zone.
* `privatekey` - (Required) File name of the private key.
* `publickey` - (Required) File name of the public key.
* `algorithm` - (Optional) Algorithm to generate the key. Defaults to `"RSASHA1"`.
* `autorollover` - (Optional) Flag to enable/disable key rollover automatically. Note: Key name will be appended with _AR1 for successor key. For e.g. current key=k1, successor key=k1_AR1. Key name can be truncated if current name length is more than 58 bytes to accomodate the suffix. Defaults to `"DISABLED"`.
* `expires` - (Optional) Time period for which to consider the key valid, after the key is used to sign a zone. Defaults to `120`.
* `filenameprefix` - (Optional) Common prefix for the names of the generated public and private key files and the Delegation Signer (DS) resource record. During key generation, the .key, .private, and .ds suffixes are appended automatically to the file name prefix to produce the names of the public key, the private key, and the DS record, respectively.
* `keysize` - (Optional) Size of the key, in bits. Defaults to `512`.
* `keytype` - (Optional) Type of key to create. Defaults to `"ZSK"`.
* `notificationperiod` - (Optional) Time at which to generate notification of key expiration, specified as number of days, hours, or minutes before expiry. Must be less than the expiry period. The notification is an SNMP trap sent to an SNMP manager. To enable the appliance to send the trap, enable the DNSKEY-EXPIRY SNMP alarm. In case autorollover option is enabled, rollover for successor key will be initiated at this time. No notification trap will be sent. Defaults to `7`.
* `password` - (Optional, Sensitive) Passphrase for reading the encrypted public/private DNS keys. The value is persisted in Terraform state (encrypted). See also `password_wo` for an ephemeral alternative.
* `password_wo` - (Optional, Sensitive, WriteOnly) Same as `password`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `password_wo_version`. If both `password` and `password_wo` are set, `password_wo` takes precedence.
* `password_wo_version` - (Optional) An integer version tracker for `password_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the value has changed. Note: this secret is immutable on the ADC, so changing `password_wo_version` (or `password`/`password_wo`) forces the resource to be **destroyed and recreated** rather than updated in place. Defaults to `1`.
* `revoke` - (Optional) Revoke the key. Note: This operation is non-reversible.
* `rollovermethod` - (Optional) Method used for automatic rollover. Key type: ZSK, Method: PrePublication or DoubleSignature. Key type: KSK, Method: DoubleRRSet.
* `src` - (Optional) URL (protocol, host, path, and file name) from where the DNS key file will be imported. NOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access.
* `ttl` - (Optional) Time to Live (TTL), in seconds, for the DNSKEY resource record created in the zone. TTL is the time for which the record must be cached by the DNS proxies. If the TTL is not specified, either the DNS zone's minimum TTL or the default value of 3600 is used. Defaults to `3600`.
* `units1` - (Optional) Units for the expiry period. Defaults to `"DAYS"`.
* `units2` - (Optional) Units for the notification period. Defaults to `"DAYS"`.
* `zonename` - (Optional) Name of the zone for which to create a key.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the dnskey. It has the same value as the `keyname` attribute.


## Import

A dnskey can be imported using its name, e.g.

```shell
terraform import citrixadc_dnskey.example adckey_1
```
