---
subcategory: "NS"
---

# Resource: nsconfig_clear

The nsconfig_clear resource is used to apply the clear operation for ns config.


## Example usage

```hcl
resource "citrixadc_nsconfig_clear" "foo" {
    force = false
    level = "full"
    rbaconfig = "YES"
}
```


## Argument Reference

* `force` - (Required) Configurations will be cleared without prompting for confirmation.
* `level` - (Required) Types of configurations to be cleared. * basic: Clears all configurations except the following: - NSIP, default route (gateway), static routes, MIPs, and SNIPs - Network settings (DG, VLAN, RHI and DNS settings) - Cluster settings - HA node definitions - Feature and mode settings - nsroot password * extended: Clears the same configurations as the 'basic' option. In addition, it clears the feature and mode settings. * full: Clears all configurations except NSIP, default route, and interface settings. Note: When you clear the configurations through the cluster IP address, by specifying the level as 'full', the cluster is deleted and all cluster nodes become standalone appliances. The 'basic' and 'extended' levels are propagated to the cluster nodes. Possible values: [ basic, extended, full ]
* `rbaconfig` - (Required) RBA configurations and TACACS policies bound to system global will not be cleared if RBA is set to NO.This option is applicable only for BASIC level of clear configuration.Default is YES, which will clear rba configurations. Possible values: [ YES, NO ]
* `timestamp` - (Required) the timestamp of the operation. Can be any string. Used to force the operation again if all other attributes have remained the same.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nsconfig_clear. It has the same value as the `timestamp` attribute.
