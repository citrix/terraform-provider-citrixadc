---
subcategory: "NS"
---

# Data Source `nspartition`

The nspartition data source allows you to retrieve information about NetScaler partitions.


## Example usage

```terraform
resource "citrixadc_nspartition" "my_partition" {
  partitionname = "my_partition"
  maxbandwidth  = 10240
  minbandwidth  = 10240
  maxconn       = 1024
  maxmemlimit   = 10
}

data "citrixadc_nspartition" "my_partition_data" {
  partitionname = citrixadc_nspartition.my_partition.partitionname
}

output "partition_mac" {
  value = data.citrixadc_nspartition.my_partition_data.partitionmac
}

output "maxbandwidth" {
  value = data.citrixadc_nspartition.my_partition_data.maxbandwidth
}
```


## Argument Reference

* `partitionname` - (Required) Name of the Partition. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `force` - Switches to new admin partition without prompt for saving configuration. Configuration will not be saved.
* `maxbandwidth` - Maximum bandwidth, in Kbps, that the partition can consume. A zero value indicates the bandwidth is unrestricted on the partition and it can consume up to the system limits.
* `maxconn` - Maximum number of concurrent connections that can be open in the partition. A zero value indicates no limit on number of open connections.
* `maxmemlimit` - Maximum memory, in megabytes, allocated to the partition. A zero value indicates the memory is unlimited on the partition and it can consume up to the system limits.
* `minbandwidth` - Minimum bandwidth, in Kbps, that the partition can consume. A zero value indicates the bandwidth is unrestricted on the partition and it can consume up to the system limits.
* `partitionmac` - Special MAC address for the partition which is used for communication over shared vlans in this partition. If not specified, the MAC address is auto-generated.
* `save` - Switches to new admin partition without prompt for saving configuration. Configuration will be saved.
* `id` - The id of the nspartition. It has the same value as the `partitionname` attribute.
