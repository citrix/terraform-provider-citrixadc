---
subcategory: "IPSec"
---

# Data Source `ipsecparameter`

The ipsecparameter data source allows you to retrieve information about IPSec parameter configuration.


## Example usage

```terraform
data "citrixadc_ipsecparameter" "tf_ipsecparameter" {
}

output "ikeversion" {
  value = data.citrixadc_ipsecparameter.tf_ipsecparameter.ikeversion
}

output "livenesscheckinterval" {
  value = data.citrixadc_ipsecparameter.tf_ipsecparameter.livenesscheckinterval
}
```


## Argument Reference

This datasource does not require any arguments.

## Attribute Reference

The following attributes are available:

* `encalgo` - Type of encryption algorithm (Note: Selection of AES enables AES128).
* `hashalgo` - Type of hashing algorithm.
* `ikeretryinterval` - IKE retry interval for bringing up the connection.
* `ikeversion` - IKE Protocol Version.
* `lifetime` - Lifetime of IKE SA in seconds. Lifetime of IPSec SA will be (lifetime of IKE SA/8).
* `livenesscheckinterval` - Number of seconds after which a notify payload is sent to check the liveliness of the peer. Additional retries are done as per retransmit interval setting. Zero value disables liveliness checks.
* `perfectforwardsecrecy` - Enable/Disable PFS.
* `replaywindowsize` - IPSec Replay window size for the data traffic.
* `retransmissiontime` - The interval in seconds to retry sending the IKE messages to peer, three consecutive attempts are done with doubled interval after every failure, increases for every retransmit till 6 retransmits.

## Attribute Reference

* `id` - The id of the ipsecparameter. It is a system-generated identifier.
