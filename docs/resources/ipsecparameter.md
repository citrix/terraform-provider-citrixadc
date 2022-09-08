---
subcategory: "ipsec"
---

# Resource: ipsecparameter

The ipsecparameter resource is used to update ipsecparameter.


## Example usage

```hcl
resource "citrixadc_ipsecparameter" "tf_ipsecparameter" {
  ikeversion            = "V2"
  encalgo               = ["AES", "3DES"]
  hashalgo              = ["HMAC_SHA1", "HMAC_SHA256"]
  livenesscheckinterval = 50
}
```


## Argument Reference

* `ikeversion` - (Optional) IKE Protocol Version. Possible values: [ V1, V2 ]
* `encalgo` - (Optional) Type of encryption algorithm (Note: Selection of AES enables AES128). Possible values: [ AES, 3DES, AES192, AES256 ]
* `hashalgo` - (Optional) Type of hashing algorithm. Possible values: [ HMAC_SHA1, HMAC_SHA256, HMAC_SHA384, HMAC_SHA512, HMAC_MD5 ]
* `lifetime` - (Optional) Lifetime of IKE SA in seconds. Lifetime of IPSec SA will be (lifetime of IKE SA/8). Minimum value =  480 Maximum value =  31536000
* `livenesscheckinterval` - (Optional) Number of seconds after which a notify payload is sent to check the liveliness of the peer. Additional retries are done as per retransmit interval setting. Zero value disables liveliness checks. Minimum value =  0 Maximum value =  64999
* `replaywindowsize` - (Optional) IPSec Replay window size for the data traffic. Minimum value =  0 Maximum value =  16384
* `ikeretryinterval` - (Optional) IKE retry interval for bringing up the connection. Minimum value =  60 Maximum value =  3600
* `perfectforwardsecrecy` - (Optional) Enable/Disable PFS. Possible values: [ enable, disable ]
* `retransmissiontime` - (Optional) The interval in seconds to retry sending the IKE messages to peer, three consecutive attempts are done with doubled interval after every failure, increases for every retransmit till 6 retransmits. Minimum value =  1 Maximum value =  99


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the ipsecparameter. It is a unique string prefixed with `tf-ipsecparameter-` attribute.
