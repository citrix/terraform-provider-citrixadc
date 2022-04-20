---
subcategory: "Bot"
---

# Resource: botprofile_whitelist_binding

The botprofile_whitelist_binding resource is used to bind whitelist to botprofile.


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
resource "citrixadc_botprofile_whitelist_binding" "tf_binding" {
  name                  = citrixadc_botprofile.tf_botprofile.name
  bot_whitelist         = "true"
  bot_whitelist_type    = "IPv4"
  bot_whitelist_value   = "1.2.1.2"
  bot_bind_comment      = "TestingWhiteList"
  bot_whitelist_enabled = "ON"
  log                   = "ON"
  logmessage            = "BotWhiteListAdded"
}
```


## Argument Reference

* `name` - (Required) Name for the profile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.), pound (#), space ( ), at (@), equals (=), colon (:), and underscore (_) characters. Cannot be changed after the profile is added.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my profile" or 'my profile').
* `bot_whitelist_value` - (Required) Value of bot white-list entry.
* `bot_bind_comment` - (Optional) Any comments about this binding.
* `bot_whitelist` - (Optional) Whitelist binding. Maximum 32 bindings can be configured per profile for Whitelist detection.
* `bot_whitelist_enabled` - (Optional) Enabled or disabled white-list binding.
* `bot_whitelist_type` - (Optional) Type of the white-list entry.
* `log` - (Optional) Enable logging for Whitelist binding.
* `logmessage` - (Optional) Message to be logged for this binding.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the botprofile_whitelist_binding. It is the concatenation of `name` and `bot_whitelist_value` attributes seperated by comma.


## Import

A botprofile_whitelist_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_botprofile_whitelist_binding.tf_binding tf_botprofile,1.2.1.2
```
