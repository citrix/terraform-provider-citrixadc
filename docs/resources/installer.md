---
subcategory: "Utility"
---

# Resource: installer

The installer resource is used to install updated version build onto a target ADC.


## Example usage

```hcl
resource "citrixadc_installer" "tf_installer" {
	url =  "file:///var/tmp/build_mana_47_24_nc_64.tgz"
    y = true
    l = false
    wait_until_reachable = true
}
```


## Argument Reference

* `url` - (Optional) Url for the build file. Must be in the following formats: http://[user]:[password]@host/path/to/file https://[user]:[password]@host/path/to/file sftp://[user]:[password]@host/path/to/file scp://[user]:[password]@host/path/to/file ftp://[user]:[password]@host/path/to/file file://path/to/file.
* `y` - (Optional) Do not prompt for yes/no before rebooting.
* `l` - (Optional) Use this flag to enable callhome.
* `enhancedupgrade` - (Optional) Use this flag for upgrading from/to enhancement mode.
* `resizeswapvar` - (Optional) Use this flag to change swap size on ONLY 64bit nCore/MCNS/VMPE systems NON-VPX systems.
* `wait_until_reachable` - (Optional) Boolean value to determine if the resource should wait for the ADC to become reachable after the build is installed.
* `reachable_timeout` - (Optional) Time period to wait untill the target ADC becomes reachable. Default value "10m"
* `reachable_poll_delay` - (Optional) Time delay before the first poll. Must be sufficiently large to allow the ADC to start reboot. Default value "60s".
* `reachable_poll_interval` - (Optional) Time interval between polls for reachability. Default value "60s".
* `reachable_poll_timeout` - (Optional) Time period to wait before the http poll request times out. Default value "20s".


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the installer. It is a random string prefixed with "tf-installer-".
