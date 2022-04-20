---
subcategory: "Bot"
---

# Resource: botprofile_logexpression_binding

The botprofile_logexpression_binding resource is used to bind logexpression to botprofile resource.


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
resource "citrixadc_botprofile_logexpression_binding" "tf_binding" {
  name                     = citrixadc_botprofile.tf_botprofile.name
  logexpression            = "true"
  bot_log_expression_name  = "tf_logname"
  bot_log_expression_value = "HTTP.REQ.BODY.CONTAINS(\"ANDROID\")"
  bot_bind_comment         = "LogTesting"
}
```


## Argument Reference

* `name` - (Required) Name for the profile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.), pound (#), space ( ), at (@), equals (=), colon (:), and underscore (_) characters. Cannot be changed after the profile is added.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my profile" or 'my profile').
* `bot_log_expression_name` - (Required) Name of the log expression object.
* `bot_bind_comment` - (Optional) Any comments about this binding.
* `bot_log_expression_enabled` - (Optional) Enable or disable the log expression binding.
* `bot_log_expression_value` - (Optional) Expression whose result to be logged when violation happened on the bot profile.
* `logexpression` - (Optional) Log expression binding.
* `logmessage` - (Optional) Message to be logged for this binding.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the botprofile_logexpression_binding. It is the concatenation of `name` and `bot_log_expression_name` attributes seperated by comma.


## Import

A botprofile_logexpression_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_botprofile_logexpression_binding.tf_binding tf_botprofile,tf_logname
```
