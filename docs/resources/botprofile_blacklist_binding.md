---
subcategory: "Bot"
---

# Resource: botprofile_blacklist_binding

The botprofile_blacklist_binding resource is used to bind blacklist to botprofile.


## Example usage

```hcl
resource "citrixadc_botprofile" "tf_botprofile" {
  name                     = "tf_botprofile"
  errorurl                 = "http://www.citrix.com"
  trapurl                  = "/http://www.citrix.com"
  comment                  = "tf_botprofile comment"
  bot_enable_white_list    = "ON"
  bot_enable_black_list    = "ON"
  bot_enable_rate_limit    = "ON"
  devicefingerprint        = "ON"
  devicefingerprintaction  = ["LOG", "RESET"]
  bot_enable_ip_reputation = "ON"
  trap                     = "ON"
  trapaction               = ["LOG", "RESET"]
  bot_enable_tps           = "ON"
}
resource "citrixadc_botprofile_blacklist_binding" "tf_binding" {
  name                  = citrixadc_botprofile.tf_botprofile.name
  bot_blacklist         = "true"
  bot_blacklist_type    = "IPv4"
  bot_blacklist_value   = "1.3.5.7"
  bot_bind_comment      = "TestingBlacklist"
  bot_blacklist_enabled = "ON"
  bot_blacklist_action  = ["LOG", "RESET"]
  logmessage            = "HelloTesting"
}
```


## Argument Reference

* `name` - (Required) Name for the profile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.), pound (#), space ( ), at (@), equals (=), colon (:), and underscore (_) characters. Cannot be changed after the profile is added.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my profile" or 'my profile').
* `bot_blacklist_value` - (Required) Value of the bot black-list entry.
* `bot_bind_comment` - (Optional) Any comments about this binding.
* `bot_blacklist` - (Optional) Blacklist binding. Maximum 32 bindings can be configured per profile for Blacklist detection.
* `bot_blacklist_action` - (Optional) One or more actions to be taken if  bot is detected based on this Blacklist binding. Only LOG action can be combined with DROP or RESET action.
* `bot_blacklist_enabled` - (Optional) Enabled or disbaled black-list binding.
* `bot_blacklist_type` - (Optional) Type of the black-list entry.
* `logmessage` - (Optional) Message to be logged for this binding.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the botprofile_blacklist_binding. It is the concatenation of `name` and `bot_blacklist_value` attributes separated by comma.


## Import

A botprofile_blacklist_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_botprofile_blacklist_binding.tf_binding tf_botprofile,1.3.5.7
