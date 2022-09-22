---
subcategory: "IPsec"
---

# Resource: ipsecprofile

The ipsecprofile resource is used to create ipsecprofile.


## Example usage

```hcl
resource "citrixadc_ipsecprofile" "tf_ipsecprofile" {
  name                  = "my_ipsecprofile"
  ikeversion            = "V2"
  encalgo               = ["AES", "3DES"]
  hashalgo              = ["HMAC_SHA1", "HMAC_SHA256"]
  livenesscheckinterval = 50
  psk                   = "GCC5VcY0TQ+0TfjGwCrR+cQthm5UnBPB"
}
```


## Argument Reference

* `name` - (Required) The name of the ipsec profile. Minimum length =  1 Maximum length =  32
* `ikeversion` - (Optional) IKE Protocol Version. Possible values: [ V1, V2 ]
* `encalgo` - (Optional) Type of encryption algorithm (Note: Selection of AES enables AES128). Possible values: [ AES, 3DES, AES192, AES256 ]
* `hashalgo` - (Optional) Type of hashing algorithm. Possible values: [ HMAC_SHA1, HMAC_SHA256, HMAC_SHA384, HMAC_SHA512, HMAC_MD5 ]
* `lifetime` - (Optional) Lifetime of IKE SA in seconds. Lifetime of IPSec SA will be (lifetime of IKE SA/8). Minimum value =  480 Maximum value =  31536000
* `psk` - (Optional) Pre shared key value.
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
