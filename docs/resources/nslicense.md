---
subcategory: "NS"
---

# Resource: nslicense

The nslicense resource is used to upload and apply a license file to the target ADC.


## Example usage

```hcl
resource "citrixadc_nslicense" "tf_license" {

    license_file = "CNS_V10000_SERVER_PLT_Retail.lic"
}
```


## Argument Reference

* `license_file` - (Required) License file to upload.
* `ssh_host` - (Optional) The ssh host ip address that will be used for the sftp transfer of the license file.
* `ssh_username` - (Optional) The user name for the ssh connection.
* `ssh_password` - (Optional) The password for the ssh connection.
* `ssh_port` - (Optional) The port for the ssh connection.
* `ssh_host_pubkey` - (Required) The ADC public ssh host key.
* `reboot` - (Optional) Set this to true to reboot and wait for the ADC to become responsive.
* `poll_delay` - (Optional) Time to wait before the first poll after reboot. Defaults to "60s".
* `poll_interval` - (Optional) Interval between polls. Defaults to "60s".
* `poll_timeout` - (Optional) Timeout for a poll attempt. Default to "10s".

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nslicense. It has the same value as the `license_file` attribute.
