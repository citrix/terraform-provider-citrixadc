---
subcategory: "Bot"
---

# Data Source: botprofile_blacklist_binding

The botprofile_blacklist_binding data source allows you to retrieve information about the bindings between botprofile and blacklist entries.

## Example Usage

```terraform
data "citrixadc_botprofile_blacklist_binding" "tf_binding" {
  name                = "tf_botprofile"
  bot_blacklist_value = "1.3.5.7"
}

output "bot_blacklist_type" {
  value = data.citrixadc_botprofile_blacklist_binding.tf_binding.bot_blacklist_type
}

output "bot_blacklist_enabled" {
  value = data.citrixadc_botprofile_blacklist_binding.tf_binding.bot_blacklist_enabled
}
```

## Argument Reference

* `name` - (Required) Name for the profile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.), pound (#), space ( ), at (@), equals (=), colon (:), and underscore (_) characters. Cannot be changed after the profile is added. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my profile" or 'my profile').
* `bot_blacklist_value` - (Required) Value of the bot black-list entry.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the botprofile_blacklist_binding. It is the concatenation of the `name` and `bot_blacklist_value` attributes separated by a comma.
* `bot_blacklist` - Blacklist binding. Maximum 32 bindings can be configured per profile for Blacklist detection.
* `bot_bind_comment` - Any comments about this binding.
* `bot_blacklist_action` - One or more actions to be taken if bot is detected based on this Blacklist binding. Only LOG action can be combined with DROP or RESET action.
* `bot_blacklist_enabled` - Enabled or disabled black-list binding.
* `bot_blacklist_type` - Type of the black-list entry.
* `logmessage` - Message to be logged for this binding.
