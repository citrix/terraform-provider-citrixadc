---
subcategory: "NS"
---

# Resource: nsvpxparam

The nsvpxparam resource is used to set ns vpx parameters.


## Example usage

```hcl
resource "citrixadc_nsvpxparam" "tf_vpxparam" {
    cpuyield = "DEFAULT"
    masterclockcpu1 = "NO"
}
```


## Argument Reference

* `masterclockcpu1` - (Optional) This setting applicable in virtual appliances, to move master clock source cpu from management cpu cpu0 to cpu1 ie PE0. * There are 2 options for the behavior: 1. YES - Allow the Virtual Appliance to move clock source to cpu1. 2. NO - Virtual Appliance will use management cpu ie cpu0 for clock source default option is NO. Possible values: [ YES, NO ]
* `cpuyield` - (Optional) This setting applicable in virtual appliances, is to affect the cpu yield(relinquishing the cpu resources) in any hypervised environment. * There are 3 options for the behavior: 1. YES - Allow the Virtual Appliance to yield its vCPUs periodically, if there is no data traffic. 2. NO - Virtual Appliance will not yield the vCPU. 3. DEFAULT - Restores the default behaviour, according to the license. * Its behavior in different scenarios: 1. As this setting is node specific only, it will not be propagated to other nodes, when executed on Cluster(CLIP) and HA(Primary). 2. In cluster setup, use '-ownerNode' to specify ID of the cluster node. 3. This setting is a system wide implementation and not granular to vCPUs. 4. No effect on the management PE. Possible values: [ DEFAULT, YES, NO ]
* `ownernode` - (Optional) ID of the cluster node for which you are setting the cpuyield. It can be configured only through the cluster IP address.
* `kvmvirtiomultiqueue` - (Optional) This setting applicable on KVM VPX with virtio NICs, is to configure multiple queues for all virtio interfaces.  * There are 2 options for this behavior: 1. YES - Allows VPX to use multiple queues for each virtio interface as configured through the KVM Hypervisor. 2. NO - Each virtio interface within VPX will use a single queue for transmit and receive.  * Its behavior in different scenarios: 1. As this setting is node specific only, it will not be propagated to other nodes, when executed on Cluster(CLIP) and HA(Primary). 2. In cluster setup, use '-ownerNode' to specify ID of the cluster node.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nsvpxparam. It is a unique string prefixed with "tf-nsvpxparam-".
