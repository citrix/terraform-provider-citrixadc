---
subcategory: "GSLB"
---

# Data Source `gslbparameter`

The gslbparameter data source allows you to retrieve information about GSLB parameter configuration.


## Example usage

```terraform
data "citrixadc_gslbparameter" "tf_gslbparameter" {
}

output "ldnsentrytimeout" {
  value = data.citrixadc_gslbparameter.tf_gslbparameter.ldnsentrytimeout
}

output "rtttolerance" {
  value = data.citrixadc_gslbparameter.tf_gslbparameter.rtttolerance
}
```


## Argument Reference

This datasource does not require any arguments.

## Attribute Reference

The following attributes are available:

* `automaticconfigsync` - GSLB configuration will be synced automatically to remote gslb sites if enabled. Possible values: `ENABLED`, `DISABLED`.
* `dropldnsreq` - Drop LDNS requests if round-trip time (RTT) information is not available. Possible values: `ENABLED`, `DISABLED`.
* `gslbconfigsyncmonitor` - If enabled, remote gslb site's rsync port will be monitored and site is considered for configuration sync only when the monitor is successful. Possible values: `ENABLED`, `DISABLED`.
* `gslbsvcstatedelaytime` - Amount of delay in updating the state of GSLB service to DOWN when MEP goes down. This parameter is applicable only if monitors are not bound to GSLB services.
* `gslbsyncinterval` - Time duartion (in seconds) for which the gslb sync process will wait before checking for config changes.
* `gslbsynclocfiles` - If disabled, Location files will not be synced to the remote sites as part of manual sync and automatic sync. Possible values: `ENABLED`, `DISABLED`.
* `gslbsyncmode` - Mode in which configuration will be synced from master site to remote sites. Possible values: `IncrementalSync`, `FullSync`.
* `gslbsyncsaveconfigcommand` - If enabled, 'save ns config' command will be treated as other GSLB commands and synced to GSLB nodes when auto gslb sync option is enabled. Possible values: `ENABLED`, `DISABLED`.
* `ldnsentrytimeout` - Time, in seconds, after which an inactive LDNS entry is removed.
* `ldnsmask` - The IPv4 network mask with which to create LDNS entries.
* `ldnsprobeorder` - Order in which monitors should be initiated to calculate RTT.
* `mepkeepalivetimeout` - Time duartion (in seconds) during which if no new packets received by Local gslb site from Remote gslb site then mark the MEP connection DOWN.
* `rtttolerance` - Tolerance, in milliseconds, for newly learned round-trip time (RTT) values. If the difference between the old RTT value and the newly computed RTT value is less than or equal to the specified tolerance value, the LDNS entry in the network metric table is not updated with the new RTT value. Prevents the exchange of metrics when variations in RTT values are negligible.
* `svcstatelearningtime` - Time (in seconds) within which local or child site services remain in learning phase. GSLB site will enter the learning phase after reboot, HA failover, Cluster GSLB owner node changes or MEP being enabled on local node. Backup parent (if configured) will selectively move the adopted children's GSLB services to learning phase when primary parent goes down. While a service is in learning period, remote site will not honour the state and stats got through MEP for that service. State can be learnt from health monitor if bound explicitly.
* `undefaction` - Action to perform when policy evaluation creates an UNDEF condition. Possible values: `NOLBACTION`, `RESET`, `DROP`.
* `v6ldnsmasklen` - Mask for creating LDNS entries for IPv6 source addresses. The mask is defined as the number of leading bits to consider, in the source IP address, when creating an LDNS entry.

## Attribute Reference

* `id` - The id of the gslbparameter. It is a system-generated identifier.
