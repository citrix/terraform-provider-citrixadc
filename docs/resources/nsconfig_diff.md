---
subcategory: "NS"
---

# Resource: nsconfig_diff

The nsconfig_diff resource compares two Citrix ADC configuration sets and reports the differences between them. Use it to review what has changed between a saved configuration and the running configuration (or between two saved configuration locations) before promoting or auditing a change.

~> **One-shot action.** This is an action resource: applying it performs the configuration diff; it does not manage a persistent object, so re-applying re-runs the diff. Each `terraform apply` that creates or replaces this resource runs the diff once, and changing any argument forces the diff to run again (replacement). Bump `timestamp` to re-run the diff when the other arguments are unchanged.


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
