---
subcategory: "Application Firewall"
---

# Data Source: citrixadc_appfwsettings

The appfwsettings data source allows you to retrieve information about Application Firewall global settings configuration.


## Example usage

```terraform
data "citrixadc_appfwsettings" "tf_appfwsettings" {
}

output "defaultprofile" {
  value = data.citrixadc_appfwsettings.tf_appfwsettings.defaultprofile
}

output "sessiontimeout" {
  value = data.citrixadc_appfwsettings.tf_appfwsettings.sessiontimeout
}
```


## Argument Reference

This datasource does not require any arguments.

## Attribute Reference

The following attributes are available:

* `ceflogging` - Enable CEF format logs. Possible values: [ ON, OFF ]
* `centralizedlearning` - Flag used to enable/disable ADM centralized learning. Possible values: [ ON, OFF ]
* `clientiploggingheader` - Name of an HTTP header that contains the IP address that the client used to connect to the protected web site or service.
* `cookieflags` - Add the specified flags to AppFW cookies. Available setttings function as follows: None - Do not add flags to AppFW cookies. HTTP Only - Add the HTTP Only flag to AppFW cookies, which prevent scripts from accessing them. Secure - Add Secure flag to AppFW cookies. All - Add both HTTPOnly and Secure flag to AppFW cookies.
* `cookiepostencryptprefix` - String that is prepended to all encrypted cookie values.
* `defaultprofile` - Profile to use when a connection does not match any policy. Default setting is APPFW_BYPASS, which sends unmatched connections back to the Citrix ADC without attempting to filter them further.
* `entitydecoding` - Transform multibyte (double- or half-width) characters to single width characters. Possible values: [ ON, OFF ]
* `geolocationlogging` - Enable Geo-Location Logging in CEF format logs. Possible values: [ ON, OFF ]
* `importsizelimit` - Maximum cumulative size in bytes of all objects imported to Netscaler. The user is not allowed to import an object if the operation exceeds the currently configured limit.
* `learnratelimit` - Maximum number of connections per second that the application firewall learning engine examines to generate new relaxations for learning-enabled security checks. The application firewall drops any connections above this limit from the list of connections used by the learning engine.
* `logmalformedreq` - Log requests that are so malformed that application firewall parsing doesn't occur. Possible values: [ ON, OFF ]
* `malformedreqaction` - Flag to define action on malformed requests that application firewall cannot parse.
* `proxypassword` - Password with which proxy user logs on.
* `proxyport` - Proxy Server Port to get updated signatures from AWS.
* `proxyserver` - Proxy Server IP to get updated signatures from AWS.
* `proxyusername` - Proxy Username.
* `sessioncookiename` - Name of the session cookie that the application firewall uses to track user sessions. Must begin with a letter or number, and can consist of from 1 to 31 letters, numbers, and the hyphen (-) and underscore (_) symbols.
* `sessionlifetime` - Maximum amount of time (in seconds) that the application firewall allows a user session to remain active, regardless of user activity. After this time, the user session is terminated. Before continuing to use the protected web site, the user must establish a new session by opening a designated start URL. A value of 0 represents infinite time.
* `sessionlimit` - Maximum number of sessions that the application firewall allows to be active, regardless of user activity. After the max_limit reaches, No more user session will be created.
* `sessiontimeout` - Timeout, in seconds, after which a user session is terminated. Before continuing to use the protected web site, the user must establish a new session by opening a designated start URL.
* `signatureautoupdate` - Flag used to enable/disable auto update signatures. Possible values: [ ON, OFF ]
* `signatureurl` - URL to download the mapping file from server.
* `undefaction` - Profile to use when an application firewall policy evaluates to undefined (UNDEF). An UNDEF event indicates an internal error condition. The APPFW_BLOCK built-in profile is the default setting. You can specify a different built-in or user-created profile as the UNDEF profile.
* `useconfigurablesecretkey` - Use configurable secret key in AppFw operations. Possible values: [ ON, OFF ]

## Attribute Reference

* `id` - The id of the appfwsettings. It is a system-generated identifier.
