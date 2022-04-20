---
subcategory: "Bot"
---

# Resource: botprofile_captcha_binding

The botprofile_captcha_binding resource is used to bind captcha to botprofile resource.


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
resource "citrixadc_botprofile_captcha_binding" "tf_binding" {
  name                = citrixadc_botprofile.tf_botprofile.name
  captcharesource     = "true"
  bot_captcha_url     = "www.example.com"
  bot_captcha_action  = ["NONE"]
  bot_captcha_enabled = "OFF"
  logmessage          = "Testing"
  retryattempts       = 4
}
```


## Argument Reference

* `name` - (Required) Name for the profile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.), pound (#), space ( ), at (@), equals (=), colon (:), and underscore (_) characters. Cannot be changed after the profile is added.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my profile" or 'my profile').
* `bot_captcha_url` - (Required) URL for which the Captcha action, if configured under IP reputation, TPS or device fingerprint, need to be applied.
* `bot_bind_comment` - (Optional) Any comments about this binding.
* `bot_captcha_action` - (Optional) One or more actions to be taken when client fails captcha challenge. Only, log action can be configured with DROP, REDIRECT or RESET action.
* `bot_captcha_enabled` - (Optional) Enable or disable the captcha binding.
* `captcharesource` - (Optional) Captcha action binding. For each URL, only one binding is allowed. To update the values of an existing URL binding, user has to first unbind that binding, and then needs to bind the URL again with new values. Maximum 30 bindings can be configured per profile.
* `graceperiod` - (Optional) Time (in seconds) duration for which no new captcha challenge is sent after current captcha challenge has been answered successfully.
* `logmessage` - (Optional) Message to be logged for this binding.
* `muteperiod` - (Optional) Time (in seconds) duration for which client which failed captcha need to wait until allowed to try again. The requests from this client are silently dropped during the mute period.
* `requestsizelimit` - (Optional) Length of body request (in Bytes) up to (equal or less than) which captcha challenge will be provided to client. Above this length threshold the request will be dropped. This is to avoid DOS and DDOS attacks.
* `retryattempts` - (Optional) Number of times client can retry solving the captcha.
* `waittime` - (Optional) Wait time in seconds for which ADC needs to wait for the Captcha response. This is to avoid DOS attacks.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the botprofile_captcha_binding. It is the concatenation of `name` and `bot_captcha_url` attributes seperated by comma.


## Import

A botprofile_captcha_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_botprofile_captcha_binding.tf_binding tf_botprofile,www.example.com
```
