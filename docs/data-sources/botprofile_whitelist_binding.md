---
subcategory: "Bot"
---

# Data Source: botprofile_whitelist_binding

The botprofile_whitelist_binding data source allows you to retrieve information about a specific botprofile whitelist binding.

## Example Usage

```terraform
data "citrixadc_botprofile_whitelist_binding" "tf_binding" {
  name                = "tf_botprofile"
  bot_whitelist_value = "1.2.1.2"
}

output "bot_whitelist_type" {
  value = data.citrixadc_botprofile_whitelist_binding.tf_binding.bot_whitelist_type
}

output "bot_bind_comment" {
  value = data.citrixadc_botprofile_whitelist_binding.tf_binding.bot_bind_comment
}
```

## Argument Reference

* `name` - (Required) Name for the profile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.), pound (#), space ( ), at (@), equals (=), colon (:), and underscore (_) characters. Cannot be changed after the profile is added.
* `bot_whitelist_value` - (Required) Value of bot white-list entry.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the botprofile_whitelist_binding. It is a system-generated identifier.
* `bot_whitelist` - Whitelist binding. Maximum 32 bindings can be configured per profile for Whitelist detection.
* `bot_bind_comment` - Any comments about this binding.
* `bot_whitelist_enabled` - Enabled or disabled white-list binding.
* `bot_whitelist_type` - Type of the white-list entry.
* `log` - Enable logging for Whitelist binding.
* `logmessage` - Message to be logged for this binding.
