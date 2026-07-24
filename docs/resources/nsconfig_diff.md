---
subcategory: "NS"
---

# Resource: nsconfig_diff

This resource is used to compare two Citrix ADC configuration sets and report their differences.


## Example usage

```hcl
resource "citrixadc_nsconfig_diff" "tf_nsconfig_diff" {
  config1              = "running"
  config2              = "saved"
  outtype              = "cli"
  ignoredevicespecific = true
  timestamp            = "2026-07-15T10:00:00Z"
}
```


## Argument Reference

* `config1` - (Optional) Location of the configurations. Changing this value forces the resource to be recreated (re-running the diff action).
* `config2` - (Optional) Location of the configurations. Changing this value forces the resource to be recreated (re-running the diff action).
* `outtype` - (Optional) Format to display the difference in configurations. Possible values: [ cli, xml ]. Changing this value forces the resource to be recreated (re-running the diff action).
* `template` - (Optional) File that contains the commands to be compared. Changing this value forces the resource to be recreated (re-running the diff action).
* `ignoredevicespecific` - (Optional) Suppress device specific differences. Changing this value forces the resource to be recreated (re-running the diff action).
* `timestamp` - (Required) Timestamp marker used as the resource ID. Can be any string. Change it to re-run the diff action when all other arguments have remained the same. Changing this value forces the resource to be recreated.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the nsconfig_diff resource. It has the same value as the `timestamp` attribute.
