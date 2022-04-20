---
subcategory: "Bot"
---

# Resource: botprofile_tps_binding

The botprofile_tps_binding resource is used to bind tps to botprofile resource.


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
resource "citrixadc_botprofile_tps_binding" "tf_binding" {
  name         = citrixadc_botprofile.tf_botprofile.name
  bot_tps_type = "SOURCE_IP"
  bot_tps      = "true"
  logmessage   = "Hellobinding"
  threshold    = 3
  percentage   = 20
}
```


## Argument Reference

* `name` - (Required) Name for the profile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.), pound (#), space ( ), at (@), equals (=), colon (:), and underscore (_) characters. Cannot be changed after the profile is added.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my profile" or 'my profile').
* `bot_tps_type` - (Required) Type of TPS binding.
* `bot_bind_comment` - (Optional) Any comments about this binding.
* `bot_tps` - (Optional) TPS binding. For each type only binding can be configured. To  update the values of an existing binding, user has to first unbind that binding, and then needs to bind again with new values.
* `bot_tps_action` - (Optional) One to more actions to be taken if bot is detected based on this TPS binding. Only LOG action can be combined with DROP, RESET, REDIRECT, or MITIGIATION action.
* `logmessage` - (Optional) Message to be logged for this binding.
* `percentage` - (Optional) Maximum percentage increase in the requests from (or to) a IP, Geolocation, URL or Host in 30 minutes interval.
* `threshold` - (Optional) Maximum number of requests that are allowed from (or to) a IP, Geolocation, URL or Host in 1 second time interval.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the botprofile_tps_binding. It is the concatenation of `name` and `bot_tps_type` attributes seperated by comma.


## Import

A botprofile_tps_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_botprofile_tps_binding.tf_binding tf_botprofile,SOURCE_IP
```
