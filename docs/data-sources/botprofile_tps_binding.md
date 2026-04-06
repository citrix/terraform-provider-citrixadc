---
subcategory: "Bot"
---

# Data Source: botprofile_tps_binding

The botprofile_tps_binding data source allows you to retrieve information about the TPS binding configuration for a bot profile.

## Example Usage

```terraform
data "citrixadc_botprofile_tps_binding" "tf_binding" {
  name         = "tf_botprofile"
  bot_tps_type = "REQUEST_URL"
}

output "logmessage" {
  value = data.citrixadc_botprofile_tps_binding.tf_binding.logmessage
}

output "threshold" {
  value = data.citrixadc_botprofile_tps_binding.tf_binding.threshold
}

output "percentage" {
  value = data.citrixadc_botprofile_tps_binding.tf_binding.percentage
}
```

## Argument Reference

* `name` - (Required) Name for the profile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.), pound (#), space ( ), at (@), equals (=), colon (:), and underscore (_) characters. Cannot be changed after the profile is added.
* `bot_tps_type` - (Required) Type of TPS binding.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the botprofile_tps_binding. It is a system-generated identifier.
* `bot_tps` - TPS binding. For each type only binding can be configured. To update the values of an existing binding, user has to first unbind that binding, and then needs to bind again with new values.
* `bot_bind_comment` - Any comments about this binding.
* `bot_tps_action` - One to more actions to be taken if bot is detected based on this TPS binding. Only LOG action can be combined with DROP, RESET, REDIRECT, or MITIGIATION action.
* `bot_tps_enabled` - Enabled or disabled TPS binding.
* `logmessage` - Message to be logged for this binding.
* `percentage` - Maximum percentage increase in the requests from (or to) a IP, Geolocation, URL or Host in 30 minutes interval.
* `threshold` - Maximum number of requests that are allowed from (or to) a IP, Geolocation, URL or Host in 1 second time interval.
