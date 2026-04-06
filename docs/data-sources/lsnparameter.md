---
subcategory: "Lsn"
---

# Data Source: lsnparameter

The lsnparameter data source allows you to retrieve information about the LSN (Large Scale NAT) global parameters.

## Example usage

```terraform
data "citrixadc_lsnparameter" "tf_lsnparameter" {
}

output "sessionsync" {
  value = data.citrixadc_lsnparameter.tf_lsnparameter.sessionsync
}

output "subscrsessionremoval" {
  value = data.citrixadc_lsnparameter.tf_lsnparameter.subscrsessionremoval
}
```

## Argument Reference

This datasource does not require any arguments.

## Attribute Reference

The following attributes are available:

* `memlimit` - Amount of Citrix ADC memory to reserve for the LSN feature, in multiples of 2MB. Note: If you later reduce the value of this parameter, the amount of active memory is not reduced. Changing the configured memory limit can only increase the amount of active memory. This command is deprecated, use 'set extendedmemoryparam -memlimit' instead.
* `sessionsync` - Synchronize all LSN sessions with the secondary node in a high availability (HA) deployment (global synchronization). After a failover, established TCP connections and UDP packet flows are kept active and resumed on the secondary node (new primary). The global session synchronization parameter and session synchronization parameters (group level) of all LSN groups are enabled by default. For a group, when both the global level and the group level LSN session synchronization parameters are enabled, the primary node synchronizes information of all LSN sessions related to this LSN group with the secondary node.
* `subscrsessionremoval` - LSN global setting for controlling subscriber aware session removal, when this is enabled, when ever the subscriber info is deleted from subscriber database, sessions corresponding to that subscriber will be removed. if this setting is disabled, subscriber sessions will be timed out as per the idle time out settings.

## Attribute Reference

* `id` - The id of the lsnparameter. It is a system-generated identifier.
