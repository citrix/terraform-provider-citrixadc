---
subcategory: "IPSec"
---

# Resource: ipsecprofile

The ipsecprofile resource is used to create IPSec profiles on Citrix ADC.


## Example usage

### Basic usage

```hcl
resource "citrixadc_ipsecprofile" "tf_ipsecprofile" {
  name                  = "my_ipsecprofile"
  ikeversion            = "V2"
  encalgo               = ["AES", "3DES"]
  hashalgo              = ["HMAC_SHA1", "HMAC_SHA256"]
  livenesscheckinterval = 50
}
```

### Using psk (sensitive attribute - persisted in state)

```hcl
variable "ipsecprofile_psk" {
  type      = string
  sensitive = true
}

resource "citrixadc_ipsecprofile" "tf_ipsecprofile" {
  name       = "my_ipsecprofile"
  ikeversion = "V2"
  psk        = var.ipsecprofile_psk
}
```

### Using psk_wo (write-only/ephemeral - NOT persisted in state)

The `psk_wo` attribute provides an ephemeral path for the pre-shared key. The value is sent to the ADC but is **not stored in Terraform state**, reducing the risk of secret exposure. To trigger an update when the key changes, increment `psk_wo_version`.

```hcl
variable "ipsecprofile_psk" {
  type      = string
  sensitive = true
}

resource "citrixadc_ipsecprofile" "tf_ipsecprofile" {
  name           = "my_ipsecprofile"
  ikeversion     = "V2"
  psk_wo         = var.ipsecprofile_psk
  psk_wo_version = 1
}
```

To rotate the key, update the variable value and bump the version:

```hcl
resource "citrixadc_ipsecprofile" "tf_ipsecprofile" {
  name           = "my_ipsecprofile"
  ikeversion     = "V2"
  psk_wo         = var.ipsecprofile_psk
  psk_wo_version = 2  # Bumped to trigger update
}
```


## Argument Reference

* `name` - (Required) The name of the ipsec profile. Minimum length =  1 Maximum length =  32
* `ikeversion` - (Optional) IKE Protocol Version. Possible values: [ V1, V2 ]
* `encalgo` - (Optional) Type of encryption algorithm (Note: Selection of AES enables AES128). Possible values: [ AES, 3DES, AES192, AES256 ]
* `hashalgo` - (Optional) Type of hashing algorithm. Possible values: [ HMAC_SHA1, HMAC_SHA256, HMAC_SHA384, HMAC_SHA512, HMAC_MD5 ]
* `lifetime` - (Optional) Lifetime of IKE SA in seconds. Lifetime of IPSec SA will be (lifetime of IKE SA/8). Minimum value =  480 Maximum value =  31536000
* `psk` - (Optional, Sensitive) Pre shared key value. The value is persisted in Terraform state (encrypted). See also `psk_wo` for an ephemeral alternative.
* `psk_wo` - (Optional, Sensitive, WriteOnly) Same as `psk`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `psk_wo_version`. If both `psk` and `psk_wo` are set, `psk_wo` takes precedence.
* `psk_wo_version` - (Optional) A user-managed integer version tracker for `psk_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the value has changed and re-send the write-only secret, which forces the resource to be replaced. This value is entirely user-controlled and has no default.
* `publickey` - (Optional) Public key file path.
* `privatekey` - (Optional) Private key file path.
* `peerpublickey` - (Optional) Peer public key file path.
* `livenesscheckinterval` - (Optional) Number of seconds after which a notify payload is sent to check the liveliness of the peer. Additional retries are done as per retransmit interval setting. Zero value disables liveliness checks. Minimum value =  0 Maximum value =  64999
* `replaywindowsize` - (Optional) IPSec Replay window size for the data traffic. Minimum value =  0 Maximum value =  16384
* `ikeretryinterval` - (Optional) IKE retry interval for bringing up the connection. Minimum value =  60 Maximum value =  3600
* `retransmissiontime` - (Optional) The interval in seconds to retry sending the IKE messages to peer, three consecutive attempts are done with doubled interval after every failure. Minimum value =  1 Maximum value =  99
* `perfectforwardsecrecy` - (Optional) Enable/Disable PFS. Possible values: [ enable, disable ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the ipsecprofile. It has the same value as the `name` attribute.


## Import

A ipsecprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_ipsecprofile.tf_ipsecprofile my_ipsecprofile
```
