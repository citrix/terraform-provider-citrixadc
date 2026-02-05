---
subcategory: "IPSec"
---

# Data Source `ipsecprofile`

The ipsecprofile data source allows you to retrieve information about IPSec profiles.


## Example usage

```terraform
data "citrixadc_ipsecprofile" "tf_ipsecprofile" {
  name = "my_ipsecprofile"
}

output "ikeversion" {
  value = data.citrixadc_ipsecprofile.tf_ipsecprofile.ikeversion
}

output "livenesscheckinterval" {
  value = data.citrixadc_ipsecprofile.tf_ipsecprofile.livenesscheckinterval
}
```


## Argument Reference

* `name` - (Required) The name of the ipsec profile.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `encalgo` - Type of encryption algorithm (Note: Selection of AES enables AES128).
* `hashalgo` - Type of hashing algorithm.
* `ikeretryinterval` - IKE retry interval for bringing up the connection.
* `ikeversion` - IKE Protocol Version.
* `lifetime` - Lifetime of IKE SA in seconds. Lifetime of IPSec SA will be (lifetime of IKE SA/8).
* `livenesscheckinterval` - Number of seconds after which a notify payload is sent to check the liveliness of the peer. Additional retries are done as per retransmit interval setting. Zero value disables liveliness checks.
* `peerpublickey` - Peer public key file path.
* `perfectforwardsecrecy` - Enable/Disable PFS.
* `privatekey` - Private key file path.
* `psk` - Pre shared key value.
* `publickey` - Public key file path.
* `replaywindowsize` - IPSec Replay window size for the data traffic.
* `retransmissiontime` - The interval in seconds to retry sending the IKE messages to peer, three consecutive attempts are done with doubled interval after every failure.

## Attribute Reference

* `id` - The id of the ipsecprofile. It has the same value as the `name` attribute.


## Import

An ipsecprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_ipsecprofile.tf_ipsecprofile my_ipsecprofile
```
