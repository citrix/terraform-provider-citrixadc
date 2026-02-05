---
subcategory: "Bot"
---

# Data Source `botsettings`

The botsettings data source allows you to retrieve information about Bot Management settings configuration.


## Example usage

```terraform
data "citrixadc_botsettings" "tf_botsettings" {
}

output "sessiontimeout" {
  value = data.citrixadc_botsettings.tf_botsettings.sessiontimeout
}

output "sessioncookiename" {
  value = data.citrixadc_botsettings.tf_botsettings.sessioncookiename
}
```


## Argument Reference

This datasource does not require any arguments.

## Attribute Reference

The following attributes are available:

* `defaultnonintrusiveprofile` - Profile to use when the feature is not enabled but feature is licensed. NonIntrusive checks will be disabled and IPRep cronjob(24 Hours) will be removed if this is set to BOT_BYPASS.
* `defaultprofile` - Profile to use when a connection does not match any policy. Default setting is " ", which sends unmatched connections back to the Citrix ADC without attempting to filter them further.
* `dfprequestlimit` - Number of requests to allow without bot session cookie if device fingerprint is enabled
* `javascriptname` - Name of the JavaScript that the Bot Management feature uses in response. Must begin with a letter or number, and can consist of from 1 to 31 letters, numbers, and the hyphen (-) and underscore (_) symbols.
* `proxypassword` - Password with which user logs on.
* `proxyport` - Proxy Server Port to get updated signatures from AWS.
* `proxyserver` - Proxy Server IP to get updated signatures from AWS.
* `proxyusername` - Proxy Username
* `sessioncookiename` - Name of the SessionCookie that the Bot Management feature uses for tracking. Must begin with a letter or number, and can consist of from 1 to 31 letters, numbers, and the hyphen (-) and underscore (_) symbols.
* `sessiontimeout` - Timeout, in seconds, after which a user session is terminated.
* `signatureautoupdate` - Flag used to enable/disable bot auto update signatures
* `signatureurl` - URL to download the bot signature mapping file from server
* `trapurlautogenerate` - Enable/disable trap URL auto generation. When enabled, trap URL is updated within the configured interval.
* `trapurlinterval` - Time in seconds after which trap URL is updated.
* `trapurllength` - Length of the auto-generated trap URL.

## Attribute Reference

* `id` - The id of the botsettings. It is a system-generated identifier.
