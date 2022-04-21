---
subcategory: "Bot"
---

# Resource: botprofile_ratelimit_binding

The botprofile_ratelimit_binding resource is used to bind ratelimit to botprofile resource.


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
resource "citrixadc_botprofile_ratelimit_binding" "tf_binding" {
  name                   = citrixadc_botprofile.tf_botprofile.name
  bot_ratelimit          = "true"
  bot_rate_limit_type    = "SESSION"
  bot_rate_limit_enabled = "ON"
  cookiename             = "name"
  rate                   = 3
  timeslice              = 20
  bot_rate_limit_action  = ["LOG", "DROP"]
}
```


## Argument Reference

* `name` - (Required) Name for the profile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.), pound (#), space ( ), at (@), equals (=), colon (:), and underscore (_) characters. Cannot be changed after the profile is added.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my profile" or 'my profile').
* `bot_rate_limit_type` - (Required) Rate-limiting type Following rate-limiting types are allowed: *SOURCE_IP - Rate-limiting based on the client IP. *SESSION - Rate-limiting based on the configured cookie name. *URL - Rate-limiting based on the configured URL.
* `bot_bind_comment` - (Optional) Any comments about this binding.
* `bot_rate_limit_action` - (Optional) One or more actions to be taken when the current rate becomes more than the configured rate. Only LOG action can be combined with DROP, REDIRECT or RESET action.
* `bot_rate_limit_enabled` - (Optional) Enable or disable rate-limit binding.
* `bot_rate_limit_url` - (Optional) URL for the resource based rate-limiting.
* `bot_ratelimit` - (Optional) Rate-limit binding. Maximum 30 bindings can be configured per profile for rate-limit detection. For SOURCE_IP type, only one binding can be configured, and for URL type, only one binding is allowed per URL, and for SESSION type, only one binding is allowed for a cookie name. To update the values of an existing binding, user has to first unbind that binding, and then needs to bind again with new values.
* `cookiename` - (Optional) Cookie name which is used to identify the session for session rate-limiting.
* `logmessage` - (Optional) Message to be logged for this binding.
* `rate` - (Optional) Maximum number of requests that are allowed in this session in the given period time.
* `timeslice` - (Optional) Time interval during which requests are tracked to check if they cross the given rate.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the botprofile_ratelimit_binding. It is the concatenation of `name` and `bot_rate_limit_type` attributes separated by comma,


## Import

A botprofile_ratelimit_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_botprofile_ratelimit_binding.tf_binding tf_botprofile,SESSION
```
