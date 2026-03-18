---
subcategory: "Bot"
---

# Data Source: botprofile_ratelimit_binding

The botprofile_ratelimit_binding data source allows you to retrieve information about the bindings between botprofile and ratelimit entries.

## Example Usage

```terraform
data "citrixadc_botprofile_ratelimit_binding" "tf_binding" {
  name                = "tf_botprofile"
  bot_rate_limit_type = "SESSION"
  cookiename          = "name"
}

output "bot_rate_limit_enabled" {
  value = data.citrixadc_botprofile_ratelimit_binding.tf_binding.bot_rate_limit_enabled
}

output "rate" {
  value = data.citrixadc_botprofile_ratelimit_binding.tf_binding.rate
}

output "timeslice" {
  value = data.citrixadc_botprofile_ratelimit_binding.tf_binding.timeslice
}
```

## Argument Reference

* `name` - (Required) Name for the profile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.), pound (#), space ( ), at (@), equals (=), colon (:), and underscore (_) characters. Cannot be changed after the profile is added. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my profile" or 'my profile').
* `bot_rate_limit_type` - (Required) Rate-limiting type Following rate-limiting types are allowed: *SOURCE_IP - Rate-limiting based on the client IP. *SESSION - Rate-limiting based on the configured cookie name. *URL - Rate-limiting based on the configured URL. *GEOLOCATION - Rate-limiting based on the configured country name. *JA3_FINGERPRINT - Rate-limiting based on client SSL JA3 fingerprint.
* `bot_rate_limit_url` - (Optional) URL for the resource based rate-limiting.
* `condition` - (Optional) Expression to be used in a rate-limiting condition. This expression result must be a boolean value.
* `cookiename` - (Optional) Cookie name which is used to identify the session for session rate-limiting.
* `countrycode` - (Optional) Country name which is used for geolocation rate-limiting.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the botprofile_ratelimit_binding. It is the concatenation of the `name` and `bot_rate_limit_type` attributes separated by a comma.
* `bot_ratelimit` - Rate-limit binding. Maximum 30 bindings can be configured per profile for rate-limit detection. For SOURCE_IP type, only one binding can be configured, and for URL type, only one binding is allowed per URL, and for SESSION type, only one binding is allowed for a cookie name. To update the values of an existing binding, user has to first unbind that binding, and then needs to bind again with new values.
* `bot_bind_comment` - Any comments about this binding.
* `bot_rate_limit_action` - One or more actions to be taken when the current rate becomes more than the configured rate. Only LOG action can be combined with DROP, REDIRECT, RESPOND_STATUS_TOO_MANY_REQUESTS or RESET action.
* `bot_rate_limit_enabled` - Enable or disable rate-limit binding.
* `limittype` - Rate-Limiting traffic Type
* `logmessage` - Message to be logged for this binding.
* `rate` - Maximum number of requests that are allowed in this session in the given period time.
* `timeslice` - Time interval during which requests are tracked to check if they cross the given rate.
