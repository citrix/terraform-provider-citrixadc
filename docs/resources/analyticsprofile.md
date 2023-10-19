---
subcategory: "Analytics"
---

# Resource: analyticsprofile

The analyticsprofile resource is used to create analyticsprofile.


## Example usage

```hcl
resource "citrixadc_analyticsprofile" "tf_analyticsprofile" {
  name             = "my_analyticsprofile"
  type             = "webinsight"
  httppagetracking = "DISABLED"
  httpurl          = "ENABLED"
}
```


## Argument Reference

* `name` - (Required) Name for the analytics profile. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my appflow profile" or 'my appflow profile'). Minimum length =  1 Maximum length =  127
* `type` - (Required) This option indicates what information needs to be collected and exported. Possible values: [ global, webinsight, tcpinsight, securityinsight, videoinsight, hdxinsight, gatewayinsight, timeseries, lsninsight, botinsight, CIinsight ]
* `collectors` - (Optional) The collector can be an IP, an appflow collector name, a service or a vserver. If IP is specified, the transport is considered as logstream and default port of 5557 is taken. If collector name is specified, the collector properties are taken from the configured collector. If service is specified, the configured service is assumed as the collector. If vserver is specified, the services bound to it are considered as collectors and the records are load balanced. Minimum length =  1
* `httpclientsidemeasurements` - (Optional) On enabling this option, the Citrix ADC will insert a javascript into the HTTP response to collect the client side page-timings and will send the same to the configured collectors. Possible values: [ ENABLED, DISABLED ]
* `httppagetracking` - (Optional) On enabling this option, the Citrix ADC will link the embedded objects of a page together. Possible values: [ ENABLED, DISABLED ]
* `httpurl` - (Optional) On enabling this option, the Citrix ADC will log the URL in appflow records. Possible values: [ ENABLED, DISABLED ]
* `httphost` - (Optional) On enabling this option, the Citrix ADC will log the Host header in appflow records. Possible values: [ ENABLED, DISABLED ]
* `httpmethod` - (Optional) On enabling this option, the Citrix ADC will log the method header in appflow records. Possible values: [ ENABLED, DISABLED ]
* `httpreferer` - (Optional) On enabling this option, the Citrix ADC will log the referer header in appflow records. Possible values: [ ENABLED, DISABLED ]
* `httpuseragent` - (Optional) On enabling this option, the Citrix ADC will log User-Agent header. Possible values: [ ENABLED, DISABLED ]
* `httpcookie` - (Optional) On enabling this option, the Citrix ADC will log cookie header. Possible values: [ ENABLED, DISABLED ]
* `httplocation` - (Optional) On enabling this option, the Citrix ADC will log location header. Possible values: [ ENABLED, DISABLED ]
* `urlcategory` - (Optional) On enabling this option, the Citrix ADC will send the URL category record. Possible values: [ ENABLED, DISABLED ]
* `allhttpheaders` - (Optional) On enabling this option, the Citrix ADC will log all the request and response headers. Possible values: [ ENABLED, DISABLED ]
* `httpcontenttype` - (Optional) On enabling this option, the Citrix ADC will log content-length header. Possible values: [ ENABLED, DISABLED ]
* `httpauthentication` - (Optional) On enabling this option, the Citrix ADC will log Authentication header. Possible values: [ ENABLED, DISABLED ]
* `httpvia` - (Optional) On enabling this option, the Citrix ADC will Via header. Possible values: [ ENABLED, DISABLED ]
* `httpxforwardedforheader` - (Optional) On enabling this option, the Citrix ADC will log X-Forwarded-For header. Possible values: [ ENABLED, DISABLED ]
* `httpsetcookie` - (Optional) On enabling this option, the Citrix ADC will log set-cookie header. Possible values: [ ENABLED, DISABLED ]
* `httpsetcookie2` - (Optional) On enabling this option, the Citrix ADC will log set-cookie2 header. Possible values: [ ENABLED, DISABLED ]
* `httpdomainname` - (Optional) On enabling this option, the Citrix ADC will log domain name. Possible values: [ ENABLED, DISABLED ]
* `httpurlquery` - (Optional) On enabling this option, the Citrix ADC will log URL Query. Possible values: [ ENABLED, DISABLED ]
* `tcpburstreporting` - (Optional) On enabling this option, the Citrix ADC will log TCP burst parameters. Possible values: [ ENABLED, DISABLED ]
* `cqareporting` - (Optional) On enabling this option, the Citrix ADC will log TCP CQA parameters. Possible values: [ ENABLED, DISABLED ]
* `integratedcache` - (Optional) On enabling this option, the Citrix ADC will log the Integrated Caching appflow records. Possible values: [ ENABLED, DISABLED ]
* `grpcstatus` - (Optional) On enabling this option, the Citrix ADC will log the gRPC status headers. Possible values: [ ENABLED, DISABLED ]
* `outputmode` - (Optional) This option indicates the format of REST API POST body. It depends on the consumer of the analytics data. Possible values: [ avro, prometheus, influx ]
* `metrics` - (Optional) This option indicates the whether metrics should be sent to the REST collector. Possible values: [ ENABLED, DISABLED ]
* `events` - (Optional) This option indicates the whether events should be sent to the REST collector. Possible values: [ ENABLED, DISABLED ]
* `auditlogs` - (Optional) This option indicates the whether auditlog should be sent to the REST collector. Possible values: [ ENABLED, DISABLED ]
* `servemode` - (Optional) This option is for setting the mode of how data is provided. Default value: Push | Possible values: [ Push, Pull ]
* `auditlogs` - (Optional) This option indicates the whether auditlog should be sent to the REST collector. Possible values: [ ENABLED, DISABLED ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the analyticsprofile. It has the same value as the `name` attribute.


## Import

A analyticsprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_analyticsprofile.tf_analyticsprofile my_analyticsprofile
```
