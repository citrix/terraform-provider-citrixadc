---
subcategory: "Lsn"
---

# Resource: lsnparameter

The lsnparameter resource is used to create lsnparameter.


## Example usage

```hcl
resource "citrixadc_lsnparameter" "tf_lsnparameter" {
  sessionsync          = "ENABLED"
  subscrsessionremoval = "ENABLED"
}
```


## Argument Reference

* `sessionsync` - (Optional) Synchronize all LSN sessions with the secondary node in a high availability (HA) deployment (global synchronization). After a failover, established TCP connections and UDP packet flows are kept active and resumed on the secondary node (new primary).  The global session synchronization parameter and session synchronization parameters (group level) of all LSN groups are enabled by default.  For a group, when both the global level and the group level LSN session synchronization parameters are enabled, the primary node synchronizes information of all LSN sessions related to this LSN group with the secondary node.
* `subscrsessionremoval` - (Optional) LSN global setting for controlling subscriber aware session removal, when this is enabled, when ever the subscriber info is deleted from subscriber database, sessions corresponding to that subscriber will be removed. if this setting is disabled, subscriber sessions will be timed out as per the idle time out settings.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lsnparameter. It is a unique string prefixed with  `tf-lsnparameter-`.