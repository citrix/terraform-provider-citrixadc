---
subcategory: "Utility"
---

# Resource: pinger

The ping resource is used to send ping requests from the target ADC.


## Example usage

```hcl
resource "citrixadc_pinger" "tf_pinger" {
    hostname = "localhost"
}
```


## Argument Reference

* `c` - (Optional) Number of packets to send. The default value is infinite. For Nitro API, defalut value is taken as 5.
* `i` - (Optional) Waiting time, in seconds. The default value is 1 second.
* `n` - (Optional) Numeric output only. No name resolution.
* `p` - (Optional) Pattern to fill in packets.  Can be up to 16 bytes, useful for diagnosing data-dependent problems.
* `q` - (Optional) Quiet output. Only the summary is printed. For Nitro API, this flag is set by default.
* `s` - (Optional) Data size, in bytes. The default value is 56.
* `t` - (Optional) Time-out, in seconds, before ping exits.
* `hostname` - (Optional) Address of host to ping.
* `forcenew_id_set` - (Optional) A set of terraform resource ids. Any change in the set will trigger the execution of the pinger.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the pinger. It is a random string prefixed with "tf-pinger-".
