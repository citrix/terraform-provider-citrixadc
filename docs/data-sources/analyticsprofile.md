---
subcategory: "Analytics"
---

# Data Source `analyticsprofile`

The analyticsprofile data source allows you to retrieve information about analytics profiles.


## Example usage

```terraform
data "citrixadc_analyticsprofile" "tf_analyticsprofile" {
  name = "my_analyticsprofile"
}

output "type" {
  value = data.citrixadc_analyticsprofile.tf_analyticsprofile.type
}

output "httppagetracking" {
  value = data.citrixadc_analyticsprofile.tf_analyticsprofile.httppagetracking
}
```


## Argument Reference

* `name` - (Required) Name for the analytics profile. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `allhttpheaders` - On enabling this option, the Citrix ADC will log all the request and response headers.
* `analyticsauthtoken` - Token for authenticating with the endpoint. If the endpoint requires the Authorization header in a particular format, specify the complete format as the value to this parameter. For eg., in case of splunk, the Authorizaiton header is required to be of the form - Splunk <auth-token>.
* `analyticsendpointcontenttype` - By default, application/json content-type is used. If this needs to be overridden, specify the value.
* `analyticsendpointmetadata` - If the endpoint requires some metadata to be present before the actual json data, specify the same.
* `analyticsendpointurl` - The URL at which to upload the analytics data on the endpoint
* `auditlogs` - This option indicates the whether auditlog should be sent to the REST collector.
* `collectors` - The collector can be an IP, an appflow collector name, a service or a vserver. If IP is specified, the transport is considered as logstream and default port of 5557 is taken. If collector name is specified, the collector properties are taken from the configured collector. If service is specified, the configured service is assumed as the collector. If vserver is specified, the services bound to it are considered as collectors and the records are load balanced.
* `cqareporting` - On enabling this option, the Citrix ADC will log TCP CQA parameters.
* `dataformatfile` - This option is for configuring the file containing the data format and metadata required by the analytics endpoint.
* `events` - This option indicates the whether events should be sent to the REST collector.
* `grpcstatus` - On enabling this option, the Citrix ADC will log the gRPC status headers
* `httpauthentication` - On enabling this option, the Citrix ADC will log Authentication header.
* `httpclientsidemeasurements` - On enabling this option, the Citrix ADC will insert a javascript into the HTTP response to collect the client side page-timings and will send the same to the configured collectors.
* `httpcontenttype` - On enabling this option, the Citrix ADC will log content-length header.
* `httpcookie` - On enabling this option, the Citrix ADC will log cookie header.
* `httpcustomheaders` - Specify the list of custom headers to be exported in web transaction records.
* `httpdomainname` - On enabling this option, the Citrix ADC will log domain name.
* `httphost` - On enabling this option, the Citrix ADC will log the Host header in appflow records
* `httplocation` - On enabling this option, the Citrix ADC will log location header.
* `httpmethod` - On enabling this option, the Citrix ADC will log the method header in appflow records
* `httppagetracking` - On enabling this option, the Citrix ADC will link the embedded objects of a page together.
* `httpreferer` - On enabling this option, the Citrix ADC will log the referer header in appflow records
* `httpsetcookie` - On enabling this option, the Citrix ADC will log set-cookie header.
* `httpsetcookie2` - On enabling this option, the Citrix ADC will log set-cookie2 header.
* `httpurl` - On enabling this option, the Citrix ADC will log the URL in appflow records
* `httpurlquery` - On enabling this option, the Citrix ADC will log URL Query.
* `httpuseragent` - On enabling this option, the Citrix ADC will log User-Agent header.
* `httpvia` - On enabling this option, the Citrix ADC will Via header.
* `httpxforwardedforheader` - On enabling this option, the Citrix ADC will log X-Forwarded-For header.
* `integratedcache` - On enabling this option, the Citrix ADC will log the Integrated Caching appflow records
* `managementlog` - This option indicates the whether managementlog should be sent to the REST collector.
* `metrics` - This option indicates the whether metrics should be sent to the REST collector.
* `metricsexportfrequency` - This option is for configuring the metrics export frequency in seconds, frequency value must be in [30,300] seconds range
* `outputmode` - This option indicates the format of REST API POST body. It depends on the consumer of the analytics data.
* `schemafile` - This option is for configuring json schema file containing a list of counters to be exported by metricscollector
* `servemode` - This option is for setting the mode of how data is provided
* `tcpburstreporting` - On enabling this option, the Citrix ADC will log TCP burst parameters.
* `topn` - On enabling this topn support, the topn information of the stream identifier this profile is bound to will be exported to the analytics endpoint.
* `type` - This option indicates what information needs to be collected and exported.
* `urlcategory` - On enabling this option, the Citrix ADC will send the URL category record.

## Attribute Reference

* `id` - The id of the analyticsprofile. It has the same value as the `name` attribute.


## Import

A analyticsprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_analyticsprofile.tf_analyticsprofile my_analyticsprofile
```
