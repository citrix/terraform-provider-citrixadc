---
subcategory: "NS"
---

# Resource: nsconfig_convert

The nsconfig_convert resource converts a classic Citrix ADC configuration file into the equivalent NITRO (declarative) representation. Use it to migrate an existing CLI-style config file into a NITRO graph that can be consumed by automation or stored for later application.

~> **One-shot action.** This is an action resource: applying it performs the configuration conversion; it does not manage a persistent object, so re-applying re-runs the conversion. Each `terraform apply` that creates or replaces this resource runs the conversion once, and changing any argument forces the conversion to run again (replacement). Bump `timestamp` to re-run the conversion when the other arguments are unchanged.


## Example usage

```hcl
resource "citrixadc_nsconfig_convert" "tf_nsconfig_convert" {
  configfile   = "/nsconfig/ns.conf"
  responsefile = "/nsconfig/ns_nitro.json"
  async        = false
  outtype      = "cli"
  timestamp    = "2026-07-15T10:00:00Z"
}
```


## Argument Reference

* `configfile` - (Required) Full path of config file to be converted to nitro. Changing this value forces the resource to be recreated (re-running the convert action).
* `responsefile` - (Optional) Full path of file to store the nitro graph. If not specified, the nitro graph is returned as part of the API response. Changing this value forces the resource to be recreated (re-running the convert action).
* `async` - (Optional) Using this option will run the operation in async mode and return the job id. The job ID can be used later to track the conversion progress via the `show ns job <id>` command. This option is mostly useful for the API to avoid timeouts for large input configurations. Changing this value forces the resource to be recreated (re-running the convert action).
* `outtype` - (Optional) Format to display the difference in configurations. Changing this value forces the resource to be recreated (re-running the convert action).
* `timestamp` - (Required) Timestamp marker used as the resource ID. Can be any string. Change it to re-run the convert action when all other arguments have remained the same. Changing this value forces the resource to be recreated.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the nsconfig_convert resource. It has the same value as the `timestamp` attribute.
