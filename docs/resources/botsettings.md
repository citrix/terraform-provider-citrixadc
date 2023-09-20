---
subcategory: "Bot"
---

# Resource: botsettings

The botsettings resource is used to update the ADC BOT settings.


## Example usage

```hcl
resource "citrixadc_botsettings" "default" {
  sessiontimeout      = "900"
  proxyport           = "8080"
  sessioncookiename   = "citrix_bot_id"
  trapurlinterval     = "3600"
  trapurllength       = "32"
}


```


## Argument Reference

* `defaultprofile` - (Optional) Profile to use when a connection does not match any policy. Default setting is " ", which sends unmatched connections back to the Citrix ADC without attempting to filter them further.
* `defaultnonintrusiveprofile` - (Optional) Profile to use when the feature is not enabled but feature is licensed. Default value: BOT_STATS, Possible values = BOT_BYPASS, BOT_STATS, BOT_LOG
* `javascriptname` - (Optional) Name of the JavaScript that the Bot Management feature  uses in response. Must begin with a letter or number, and can consist of from 1 to 31 letters, numbers, and the hyphen (-) and underscore (_) symbols. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my cookie name" or 'my cookie name').
* `sessiontimeout` - (Optional) Timeout, in seconds, after which a user session is terminated.
* `sessioncookiename` - (Optional) Name of the SessionCookie that the Bot Management feature uses for tracking. Must begin with a letter or number, and can consist of from 1 to 31 letters, numbers, and the hyphen (-) and underscore (_) symbols. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my cookie name" or 'my cookie name').
* `dfprequestlimit` - (Optional) Number of requests to allow without bot session cookie if device fingerprint is enabled.
* `signatureautoupdate` - (Optional) Flag used to enable/disable bot auto update signatures. Possible values: [ on, off ]
* `signatureurl` - (Optional) URL to download the bot signature mapping file from server.
* `proxyserver` - (Optional) Proxy Server IP to get updated signatures from AWS.
* `proxyport` - (Optional) Proxy Server Port to get updated signatures from AWS. Range 1-65535 * in CLI is represented as 65535 in NITRO API
* `trapurlautogenerate` - (Optional) Enable/disable trap URL auto generation. When enabled, trap URL is updated within the configured interval. Default value: OFF Possible values = ON, OFF
* `trapurlinterval` - (Optional)Time in seconds after which trap URL is updated. Default value: 3600 Minimum value = 300 Maximum value = 86400
* `trapurllength` - (Optional) Length of the auto-generated trap URL. Default value: 32 Minimum value = 10 Maximum value = 255 

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the botsettings.It is a unique string prefixed with "tf-botsettings".

## Import

A appfwsettings can be imported using its id, e.g.

```shell
terraform import citrixadc_botsettings.default tf-appfwsettings-1234567890
```
