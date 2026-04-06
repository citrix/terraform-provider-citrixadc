---
subcategory: "Bot"
---

# Data Source: botprofile_trapinsertionurl_binding

The botprofile_trapinsertionurl_binding data source allows you to retrieve information about trapinsertionurl bindings to a botprofile.


## Example usage

```terraform
data "citrixadc_botprofile_trapinsertionurl_binding" "tf_binding" {
  name             = "tf_botprofile"
  bot_trap_url     = "www.example.com"
}

output "bot_bind_comment" {
  value = data.citrixadc_botprofile_trapinsertionurl_binding.tf_binding.bot_bind_comment
}

output "bot_trap_url_insertion_enabled" {
  value = data.citrixadc_botprofile_trapinsertionurl_binding.tf_binding.bot_trap_url_insertion_enabled
}
```


## Argument Reference

* `name` - (Required) Name for the profile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.), pound (#), space ( ), at (@), equals (=), colon (:), and underscore (_) characters. Cannot be changed after the profile is added.
* `bot_trap_url` - (Required) Request URL regex pattern for which Trap URL is inserted.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `trapinsertionurl` - Bind the trap URL for the configured request URLs. Maximum 30 bindings can be configured per profile.
* `bot_bind_comment` - Any comments about this binding.
* `bot_trap_url_insertion_enabled` - Enable or disable the request URL pattern.
* `id` - The id of the botprofile_trapinsertionurl_binding. It is a system-generated identifier.
* `logmessage` - Message to be logged for this binding.
