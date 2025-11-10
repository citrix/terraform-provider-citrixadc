---
subcategory: "DNS"
---

# Resource: dnskey

The dnskey resource is used to create DNS key.


## Example usage

```hcl
resource "citrixadc_dnskey" "dnskey" {
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


## Argument Reference

* `keyname` - (Required) Name of the public-private key pair to publish in the zone.
* `privatekey` - (Required) File name of the private key.
* `publickey` - (Required) File name of the public key.
* `algorithm` - (Optional) Algorithm to generate for zone signing.
* `expires` - (Optional) Time period for which to consider the key valid, after the key is used to sign a zone.
* `filenameprefix` - (Optional) Common prefix for the names of the generated public and private key files and the Delegation Signer (DS) resource record. During key generation, the .key, .private, and .ds suffixes are appended automatically to the file name prefix to produce the names of the public key, the private key, and the DS record, respectively.
* `keysize` - (Optional) Size of the key, in bits.
* `keytype` - (Optional) Type of key to create.
* `notificationperiod` - (Optional) Time at which to generate notification of key expiration, specified as number of days, hours, or minutes before expiry. Must be less than the expiry period. The notification is an SNMP trap sent to an SNMP manager. To enable the appliance to send the trap, enable the DNSKEY-EXPIRY SNMP alarm.
* `password` - (Optional) Passphrase for reading the encrypted public/private DNS keys
* `src` - (Optional) URL (protocol, host, path, and file name) from where the DNS key file will be imported. NOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access. This is a mandatory argument
* `ttl` - (Optional) Time to Live (TTL), in seconds, for the DNSKEY resource record created in the zone. TTL is the time for which the record must be cached by the DNS proxies. If the TTL is not specified, either the DNS zone's minimum TTL or the default value of 3600 is used.
* `units1` - (Optional) Units for the expiry period.
* `units2` - (Optional) Units for the notification period.
* `zonename` - (Optional) Name of the zone for which to create a key.
* `autorollover` - (Optional) Flag to enable/disable key rollover automatically. Note: * Key name will be appended with _AR1 for successor key. For e.g. current key=k1, successor key=k1_AR1. * Key name can be truncated if current name length is more than 58 bytes to accomodate the suffix.
* `revoke` - (Optional) Revoke the key. Note: This operation is non-reversible.
* `rollovermethod` - (Optional) Method used for automatic rollover. * Key type: ZSK, Method: PrePublication or DoubleSignature. * Key type: KSK, Method: DoubleRRSet.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the dnskey. It has the same value as the `keyname` attribute.


## Import

A dnskey can be imported using its name, e.g.

```shell
terraform import citrixadc_dnskey.dnskey adckey_1
```
