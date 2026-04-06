---
subcategory: "Bot"
---

# Data Source: botprofile_captcha_binding

The botprofile_captcha_binding data source allows you to retrieve information about the bindings between botprofile and captcha URLs.

## Example Usage

```terraform
data "citrixadc_botprofile_captcha_binding" "tf_binding" {
  name            = "tf_botprofile"
  bot_captcha_url = "www.example.com"
}

output "bot_captcha_enabled" {
  value = data.citrixadc_botprofile_captcha_binding.tf_binding.bot_captcha_enabled
}

output "retryattempts" {
  value = data.citrixadc_botprofile_captcha_binding.tf_binding.retryattempts
}
```

## Argument Reference

* `name` - (Required) Name for the profile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.), pound (#), space ( ), at (@), equals (=), colon (:), and underscore (_) characters. Cannot be changed after the profile is added. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my profile" or 'my profile').
* `bot_captcha_url` - (Required) URL for which the Captcha action, if configured under IP reputation, TPS or device fingerprint, need to be applied.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the botprofile_captcha_binding. It is the concatenation of the `name` and `bot_captcha_url` attributes separated by a comma.
* `captcharesource` - Captcha action binding. For each URL, only one binding is allowed. To update the values of an existing URL binding, user has to first unbind that binding, and then needs to bind the URL again with new values. Maximum 30 bindings can be configured per profile.
* `bot_bind_comment` - Any comments about this binding.
* `bot_captcha_action` - One or more actions to be taken when client fails captcha challenge. Only, log action can be configured with DROP, REDIRECT or RESET action.
* `bot_captcha_enabled` - Enable or disable the captcha binding.
* `graceperiod` - Time (in seconds) duration for which no new captcha challenge is sent after current captcha challenge has been answered successfully.
* `logmessage` - Message to be logged for this binding.
* `muteperiod` - Time (in seconds) duration for which client which failed captcha need to wait until allowed to try again. The requests from this client are silently dropped during the mute period.
* `requestsizelimit` - Length of body request (in Bytes) up to (equal or less than) which captcha challenge will be provided to client. Above this length threshold the request will be dropped. This is to avoid DOS and DDOS attacks.
* `retryattempts` - Number of times client can retry solving the captcha.
* `waittime` - Wait time in seconds for which ADC needs to wait for the Captcha response. This is to avoid DOS attacks.
