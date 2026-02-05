---
subcategory: "Bot"
---

# Data Source: citrixadc_botprofile

Use this data source to retrieve information about an existing Bot Profile.

The `citrixadc_botprofile` data source allows you to retrieve details of a bot profile by its name. This is useful for referencing existing bot profiles in your Terraform configurations without managing them directly.

## Example usage

```hcl
# Retrieve an existing bot profile
data "citrixadc_botprofile" "example" {
  name = "my_bot_profile"
}

# Use the retrieved profile data in a bot policy
resource "citrixadc_botpolicy" "example_policy" {
  name        = "example_bot_policy"
  profilename = data.citrixadc_botprofile.example.name
  rule        = "true"
}

# Reference profile attributes
output "profile_errorurl" {
  value = data.citrixadc_botprofile.example.errorurl
}

output "profile_sessiontimeout" {
  value = data.citrixadc_botprofile.example.sessiontimeout
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name for the profile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.), pound (#), space ( ), at (@), equals (=), colon (:), and underscore (_) characters. Cannot be changed after the profile is added.

## Attribute Reference

In addition to the argument above, the following attributes are exported:

* `id` - The ID of the bot profile (same as name).

* `addcookieflags` - Add the specified flags to bot session cookies. Available settings function as follows:
  * None - Do not add flags to cookies.
  * HTTP Only - Add the HTTP Only flag to cookies, which prevents scripts from accessing cookies.
  * Secure - Add Secure flag to cookies.
  * All - Add both HTTPOnly and Secure flags to cookies.

* `bot_enable_black_list` - Enable black-list bot detection.

* `bot_enable_ip_reputation` - Enable IP-reputation bot detection.

* `bot_enable_rate_limit` - Enable rate-limit bot detection.

* `bot_enable_tps` - Enable TPS.

* `bot_enable_white_list` - Enable white-list bot detection.

* `clientipexpression` - Expression to get the client IP.

* `comment` - Any comments about the purpose of profile, or other useful information about the profile.

* `devicefingerprint` - Enable device-fingerprint bot detection.

* `devicefingerprintaction` - Action to be taken for device-fingerprint based bot detection.

* `devicefingerprintmobile` - Enabling bot device fingerprint protection for mobile clients.

* `dfprequestlimit` - Number of requests to allow without bot session cookie if device fingerprint is enabled.

* `errorurl` - URL that Bot protection uses as the Error URL.

* `headlessbrowserdetection` - Enable Headless Browser detection.

* `kmdetection` - Enable keyboard-mouse based bot detection.

* `kmeventspostbodylimit` - Size of the KM data send by the browser, needs to be processed on ADC.

* `kmjavascriptname` - Name of the JavaScript file that the Bot Management feature will insert in the response for keyboard-mouse based detection. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.

* `sessioncookiename` - Name of the SessionCookie that the Bot Management feature uses for tracking. Must begin with a letter or number, and can consist of from 1 to 31 letters, numbers, and the hyphen (-) and underscore (_) symbols.

* `sessiontimeout` - Timeout, in seconds, after which a user session is terminated.

* `signature` - Name of object containing bot static signature details.

* `signaturemultipleuseragentheaderaction` - Actions to be taken if multiple User-Agent headers are seen in a request (Applicable if Signature check is enabled). Log action should be combined with other actions.

* `signaturenouseragentheaderaction` - Actions to be taken if no User-Agent header in the request (Applicable if Signature check is enabled).

* `spoofedreqaction` - Actions to be taken on a spoofed request (A request spoofing good bot user agent string).

* `trap` - Enable trap bot detection.

* `trapaction` - Action to be taken for bot trap based bot detection.

* `trapurl` - URL that Bot protection uses as the Trap URL.

* `verboseloglevel` - Bot verbose Logging. Based on the log level, ADC will log additional information whenever client is detected as a bot.

## Common Use Cases

### Retrieve Profile for Bot Policy

```hcl
data "citrixadc_botprofile" "bot_protection" {
  name = "bot_protection_profile"
}

resource "citrixadc_botpolicy" "bot_policy" {
  name        = "bot_protection_policy"
  profilename = data.citrixadc_botprofile.bot_protection.name
  rule        = "CLIENT.IP.SRC.IN_SUBNET(10.0.0.0/8)"
}
```

### Reference Profile in Multiple Policies

```hcl
data "citrixadc_botprofile" "common_bot_profile" {
  name = "common_bot_protection"
}

resource "citrixadc_botpolicy" "web_bot_policy" {
  name        = "web_bot_policy"
  profilename = data.citrixadc_botprofile.common_bot_profile.name
  rule        = "HTTP.REQ.URL.STARTSWITH(\"/web\")"
}

resource "citrixadc_botpolicy" "api_bot_policy" {
  name        = "api_bot_policy"
  profilename = data.citrixadc_botprofile.common_bot_profile.name
  rule        = "HTTP.REQ.URL.STARTSWITH(\"/api\")"
}
```

### Check Profile Configuration

```hcl
data "citrixadc_botprofile" "existing_profile" {
  name = "production_bot_profile"
}

output "bot_profile_details" {
  value = {
    name                     = data.citrixadc_botprofile.existing_profile.name
    errorurl                 = data.citrixadc_botprofile.existing_profile.errorurl
    trapurl                  = data.citrixadc_botprofile.existing_profile.trapurl
    sessiontimeout           = data.citrixadc_botprofile.existing_profile.sessiontimeout
    devicefingerprint        = data.citrixadc_botprofile.existing_profile.devicefingerprint
    bot_enable_white_list    = data.citrixadc_botprofile.existing_profile.bot_enable_white_list
    bot_enable_black_list    = data.citrixadc_botprofile.existing_profile.bot_enable_black_list
    headlessbrowserdetection = data.citrixadc_botprofile.existing_profile.headlessbrowserdetection
  }
}
```

## Notes

* Bot profiles are used to configure various bot management features including device fingerprinting, trap URLs, rate limiting, and more.
* Bot profiles must be associated with bot policies to be applied to traffic.
* The profile name is case-sensitive and must match exactly.
* Some attributes may return empty values if they were not configured in the profile.
