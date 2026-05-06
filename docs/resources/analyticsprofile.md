---
subcategory: "Analytics"
---

# Resource: analyticsprofile

The analyticsprofile resource is used to create an analytics profile on a Citrix ADC appliance.


## Example usage

### Using analyticsauthtoken (sensitive attribute - persisted in state)

```hcl
variable "analyticsprofile_analyticsauthtoken" {
  type      = string
  sensitive = true
}

resource "citrixadc_analyticsprofile" "tf_analyticsprofile" {
  name               = "my_analyticsprofile"
  type               = "webinsight"
  analyticsauthtoken = var.analyticsprofile_analyticsauthtoken
  httppagetracking   = "DISABLED"
  httpurl            = "ENABLED"
}
```

### Using analyticsauthtoken_wo (write-only/ephemeral - NOT persisted in state)

The `analyticsauthtoken_wo` attribute provides an ephemeral path for the authentication token. The value is sent to the ADC but is **not stored in Terraform state**, reducing the risk of secret exposure. To trigger an update when the value changes, increment `analyticsauthtoken_wo_version`.

```hcl
variable "analyticsprofile_analyticsauthtoken" {
  type      = string
  sensitive = true
}

resource "citrixadc_analyticsprofile" "tf_analyticsprofile" {
  name                          = "my_analyticsprofile"
  type                          = "webinsight"
  analyticsauthtoken_wo         = var.analyticsprofile_analyticsauthtoken
  analyticsauthtoken_wo_version = 1
  httppagetracking              = "DISABLED"
  httpurl                       = "ENABLED"
}
```

To rotate the token, update the variable value and increment the version:

```hcl
resource "citrixadc_analyticsprofile" "tf_analyticsprofile" {
  name                          = "my_analyticsprofile"
  type                          = "webinsight"
  analyticsauthtoken_wo         = var.analyticsprofile_analyticsauthtoken
  analyticsauthtoken_wo_version = 2  # Bumped to trigger update
  httppagetracking              = "DISABLED"
  httpurl                       = "ENABLED"
}
```


## Argument Reference

* `name` - (Required) Name for the analytics profile. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my appflow profile" or 'my appflow profile').
* `type` - (Required) This option indicates what information needs to be collected and exported. Possible values: [ global, webinsight, tcpinsight, securityinsight, videoinsight, hdxinsight, gatewayinsight, timeseries, lsninsight, botinsight, CIinsight ]
* `allhttpheaders` - (Optional) On enabling this option, the Citrix ADC will log all the request and response headers. Possible values: [ ENABLED, DISABLED ]. Defaults to `"DISABLED"`.
* `analyticsauthtoken` - (Optional, Sensitive) Token for authenticating with the endpoint. If the endpoint requires the Authorization header in a particular format, specify the complete format as the value to this parameter. For eg., in case of splunk, the Authorization header is required to be of the form - Splunk <auth-token>. The value is persisted in Terraform state (encrypted). See also `analyticsauthtoken_wo` for an ephemeral alternative.
* `analyticsauthtoken_wo` - (Optional, Sensitive, WriteOnly) Same as `analyticsauthtoken`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `analyticsauthtoken_wo_version`. If both `analyticsauthtoken` and `analyticsauthtoken_wo` are set, `analyticsauthtoken_wo` takes precedence.
* `analyticsauthtoken_wo_version` - (Optional) An integer version tracker for `analyticsauthtoken_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the value has changed and trigger an update. Defaults to `1`.
* `analyticsendpointcontenttype` - (Optional) By default, application/json content-type is used. If this needs to be overridden, specify the value.
* `analyticsendpointmetadata` - (Optional) If the endpoint requires some metadata to be present before the actual json data, specify the same.
* `analyticsendpointurl` - (Optional) The URL at which to upload the analytics data on the endpoint.
* `auditlogs` - (Optional) This option indicates whether auditlog should be sent to the REST collector. Possible values: [ ENABLED, DISABLED ]. Defaults to `"DISABLED"`.
* `collectors` - (Optional) The collector can be an IP, an appflow collector name, a service or a vserver. If IP is specified, the transport is considered as logstream and default port of 5557 is taken. If collector name is specified, the collector properties are taken from the configured collector. If service is specified, the configured service is assumed as the collector. If vserver is specified, the services bound to it are considered as collectors and the records are load balanced.
* `cqareporting` - (Optional) On enabling this option, the Citrix ADC will log TCP CQA parameters. Possible values: [ ENABLED, DISABLED ]. Defaults to `"DISABLED"`.
* `dataformatfile` - (Optional) This option is for configuring the file containing the data format and metadata required by the analytics endpoint.
* `events` - (Optional) This option indicates whether events should be sent to the REST collector. Possible values: [ ENABLED, DISABLED ]. Defaults to `"DISABLED"`.
* `grpcstatus` - (Optional) On enabling this option, the Citrix ADC will log the gRPC status headers. Possible values: [ ENABLED, DISABLED ]. Defaults to `"DISABLED"`.
* `httpauthentication` - (Optional) On enabling this option, the Citrix ADC will log Authentication header. Possible values: [ ENABLED, DISABLED ]. Defaults to `"DISABLED"`.
* `httpclientsidemeasurements` - (Optional) On enabling this option, the Citrix ADC will insert a javascript into the HTTP response to collect the client side page-timings and will send the same to the configured collectors. Possible values: [ ENABLED, DISABLED ]. Defaults to `"DISABLED"`.
* `httpcontenttype` - (Optional) On enabling this option, the Citrix ADC will log content-length header. Possible values: [ ENABLED, DISABLED ]. Defaults to `"DISABLED"`.
* `httpcookie` - (Optional) On enabling this option, the Citrix ADC will log cookie header. Possible values: [ ENABLED, DISABLED ]. Defaults to `"DISABLED"`.
* `httpcustomheaders` - (Optional) Specify the list of custom headers to be exported in web transaction records.
* `httpdomainname` - (Optional) On enabling this option, the Citrix ADC will log domain name. Possible values: [ ENABLED, DISABLED ]. Defaults to `"DISABLED"`.
* `httphost` - (Optional) On enabling this option, the Citrix ADC will log the Host header in appflow records. Possible values: [ ENABLED, DISABLED ]. Defaults to `"DISABLED"`.
* `httplocation` - (Optional) On enabling this option, the Citrix ADC will log location header. Possible values: [ ENABLED, DISABLED ]. Defaults to `"DISABLED"`.
* `httpmethod` - (Optional) On enabling this option, the Citrix ADC will log the method header in appflow records. Possible values: [ ENABLED, DISABLED ]. Defaults to `"DISABLED"`.
* `httppagetracking` - (Optional) On enabling this option, the Citrix ADC will link the embedded objects of a page together. Possible values: [ ENABLED, DISABLED ]. Defaults to `"DISABLED"`.
* `httpreferer` - (Optional) On enabling this option, the Citrix ADC will log the referer header in appflow records. Possible values: [ ENABLED, DISABLED ]. Defaults to `"DISABLED"`.
* `httpsetcookie` - (Optional) On enabling this option, the Citrix ADC will log set-cookie header. Possible values: [ ENABLED, DISABLED ]. Defaults to `"DISABLED"`.
* `httpsetcookie2` - (Optional) On enabling this option, the Citrix ADC will log set-cookie2 header. Possible values: [ ENABLED, DISABLED ]. Defaults to `"DISABLED"`.
* `httpurl` - (Optional) On enabling this option, the Citrix ADC will log the URL in appflow records. Possible values: [ ENABLED, DISABLED ]. Defaults to `"DISABLED"`.
* `httpurlquery` - (Optional) On enabling this option, the Citrix ADC will log URL Query. Possible values: [ ENABLED, DISABLED ]. Defaults to `"DISABLED"`.
* `httpuseragent` - (Optional) On enabling this option, the Citrix ADC will log User-Agent header. Possible values: [ ENABLED, DISABLED ]. Defaults to `"DISABLED"`.
* `httpvia` - (Optional) On enabling this option, the Citrix ADC will log the Via header. Possible values: [ ENABLED, DISABLED ]. Defaults to `"DISABLED"`.
* `httpxforwardedforheader` - (Optional) On enabling this option, the Citrix ADC will log X-Forwarded-For header. Possible values: [ ENABLED, DISABLED ]. Defaults to `"DISABLED"`.
* `integratedcache` - (Optional) On enabling this option, the Citrix ADC will log the Integrated Caching appflow records. Possible values: [ ENABLED, DISABLED ]. Defaults to `"DISABLED"`.
* `managementlog` - (Optional) This option indicates whether managementlog should be sent to the REST collector.
* `metrics` - (Optional) This option indicates whether metrics should be sent to the REST collector. Possible values: [ ENABLED, DISABLED ]. Defaults to `"DISABLED"`.
* `metricsexportfrequency` - (Optional) This option is for configuring the metrics export frequency in seconds, frequency value must be in [30,300] seconds range. Defaults to `30`.
* `outputmode` - (Optional) This option indicates the format of REST API POST body. It depends on the consumer of the analytics data. Possible values: [ avro, prometheus, influx ]. Defaults to `"avro"`.
* `schemafile` - (Optional) This option is for configuring json schema file containing a list of counters to be exported by metricscollector.
* `servemode` - (Optional) This option is for setting the mode of how data is provided. Possible values: [ Push, Pull ]. Defaults to `"Push"`.
* `tcpburstreporting` - (Optional) On enabling this option, the Citrix ADC will log TCP burst parameters. Possible values: [ ENABLED, DISABLED ]. Defaults to `"ENABLED"`.
* `topn` - (Optional) On enabling this topn support, the topn information of the stream identifier this profile is bound to will be exported to the analytics endpoint. Possible values: [ ENABLED, DISABLED ]. Defaults to `"DISABLED"`.
* `urlcategory` - (Optional) On enabling this option, the Citrix ADC will send the URL category record. Possible values: [ ENABLED, DISABLED ]. Defaults to `"DISABLED"`.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the analyticsprofile. It has the same value as the `name` attribute.


## Import

An analyticsprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_analyticsprofile.tf_analyticsprofile my_analyticsprofile
```
