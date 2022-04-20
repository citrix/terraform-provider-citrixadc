---
subcategory: "Bot"
---

# Resource: botprofile_trapinsertionurl_binding

The botprofile_trapinsertionurl_binding resource is used to bind trapinsertionurl to botprofile.


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
resource "citrixadc_botprofile_trapinsertionurl_binding" "tf_binding" {
  name                           = citrixadc_botprofile.tf_botprofile.name
  trapinsertionurl               = "true"
  bot_trap_url                   = "www.example.com"
  bot_bind_comment               = "testing"
  bot_trap_url_insertion_enabled = "OFF"
}
```


## Argument Reference

* `name` - (Required) Name for the profile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.), pound (#), space ( ), at (@), equals (=), colon (:), and underscore (_) characters. Cannot be changed after the profile is added.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my profile" or 'my profile').
* `bot_trap_url` - (Required) Request URL regex pattern for which Trap URL is inserted.
* `bot_bind_comment` - (Optional) Any comments about this binding.
* `bot_trap_url_insertion_enabled` - (Optional) Enable or disable the request URL pattern.
* `logmessage` - (Optional) Message to be logged for this binding.
* `trapinsertionurl` - (Optional) Bind the trap URL for the configured request URLs. Maximum 30 bindings can be configured per profile.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the botprofile_trapinsertionurl_binding. It is the concatenation of `name` and `bot_trap_url` attributes seperated by comma.


## Import

A botprofile_trapinsertionurl_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_botprofile_trapinsertionurl_binding.tf_binding tf_botprofile,www.example.com
```
