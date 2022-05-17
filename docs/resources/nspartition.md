---
subcategory: "NS"
---

# Resource: nspartition

The nspartition resource is used to create ns partition resource.


## Example usage

```hcl
resource "citrixadc_nspartition" "tf_nspartition" {
  partitionname = "tf_nspartition"
  maxbandwidth  = 10240
  minbandwidth  = 512
  maxconn       = 512
  maxmemlimit   = 11
}
```


## Argument Reference

* `partitionname` - (Required) Name of the Partition. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Minimum length =  1
* `maxbandwidth` - (Optional) Maximum bandwidth, in Kbps, that the partition can consume. A zero value indicates the bandwidth is unrestricted on the partition and it can consume up to the system limits.
* `minbandwidth` - (Optional) Minimum bandwidth, in Kbps, that the partition can consume. A zero value indicates the bandwidth is unrestricted on the partition and it can consume up to the system limits.
* `maxconn` - (Optional) Maximum number of concurrent connections that can be open in the partition. A zero value indicates no limit on number of open connections.
* `maxmemlimit` - (Optional) Maximum memory, in megabytes, allocated to the partition.  A zero value indicates the memory is unlimited on the partition and it can consume up to the system limits. Minimum value =  0 Maximum value =  1048576
* `partitionmac` - (Optional) Special MAC address for the partition which is used for communication over shared vlans in this partition. If not specified, the MAC address is auto-generated.
* `force` - (Optional) Switches to new admin partition without prompt for saving configuration. Configuration will not be saved.
* `save` - (Optional) Switches to new admin partition without prompt for saving configuration. Configuration will be saved.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nspartition. It has the same value as the `partitionname` attribute.


## Import

A nspartition can be imported using its partitionname, e.g.

```shell
terraform import citrixadc_nspartition.tf_nspartition tf_nspartition
```
