---
subcategory: "AppFlow"
---

# Data Source `appflowparam`

The appflowparam data source allows you to retrieve information about AppFlow parameters configuration.


## Example usage

```terraform
data "citrixadc_appflowparam" "tf_appflowparam" {
}

output "templaterefresh" {
  value = data.citrixadc_appflowparam.tf_appflowparam.templaterefresh
}

output "flowrecordinterval" {
  value = data.citrixadc_appflowparam.tf_appflowparam.flowrecordinterval
}
```


## Argument Reference

This datasource does not require any arguments.

## Attribute Reference

The following attributes are available:

* `aaausername` - Enable AppFlow AAA Username logging.
* `analyticsauthtoken` - Authentication token to be set by the agent.
* `appnamerefresh` - Interval, in seconds, at which to send Appnames to the configured collectors. Appname refers to the name of an entity (virtual server, service, or service group) in the Citrix ADC.
* `auditlogs` - Enable Auditlogs to be sent to the Telemetry Agent
* `cacheinsight` - Flag to determine whether cache records need to be exported or not. If this flag is true and IC is enabled, cache records are exported instead of L7 HTTP records
* `clienttrafficonly` - Generate AppFlow records for only the traffic from the client.
* `connectionchaining` - Enable connection chaining so that the client server flows of a connection are linked. Also the connection chain ID is propagated across Citrix ADCs, so that in a multi-hop environment the flows belonging to the same logical connection are linked. This id is also logged as part of appflow record
* `cqareporting` - TCP CQA reporting enable/disable knob.
* `distributedtracing` - Enable generation of the distributed tracing templates in the Appflow records
* `disttracingsamplingrate` - Sampling rate for Distributed Tracing
* `emailaddress` - Enable AppFlow user email-id logging.
* `events` - Enable Events to be sent to the Telemetry Agent
* `flowrecordinterval` - Interval, in seconds, at which to send flow records to the configured collectors.
* `gxsessionreporting` - Enable this option for Gx session reporting
* `httpauthorization` - Include the HTTP Authorization header information.
* `httpcontenttype` - Include the HTTP Content-Type header sent from the server to the client to determine the type of the content sent.
* `httpcookie` - Include the cookie that was in the HTTP request the appliance received from the client.
* `httpdomain` - Include the http domain request to be exported.
* `httphost` - Include the host identified in the HTTP request that the appliance received from the client.
* `httplocation` - Include the HTTP location headers returned from the HTTP responses.
* `httpmethod` - Include the method that was specified in the HTTP request that the appliance received from the client.
* `httpquerywithurl` - Include the HTTP query segment along with the URL that the Citrix ADC received from the client.
* `httpreferer` - Include the web page that was last visited by the client.
* `httpsetcookie` - Include the Set-cookie header sent from the server to the client in response to a HTTP request.
* `httpsetcookie2` - Include the Set-cookie header sent from the server to the client in response to a HTTP request.
* `httpurl` - Include the http URL that the Citrix ADC received from the client.
* `httpuseragent` - Include the client application through which the HTTP request was received by the Citrix ADC.
* `httpvia` - Include the httpVia header which contains the IP address of proxy server through which the client accessed the server.
* `httpxforwardedfor` - Include the httpXForwardedFor header, which contains the original IP Address of the client using a proxy server to access the server.
* `identifiername` - Include the stream identifier name to be exported.
* `identifiersessionname` - Include the stream identifier session name to be exported.
* `logstreamovernsip` - To use the Citrix ADC IP to send Logstream records instead of the SNIP
* `lsnlogging` - On enabling this option, the Citrix ADC will send the Large Scale Nat(LSN) records to the configured collectors.
* `metrics` - Enable Citrix ADC Stats to be sent to the Telemetry Agent
* `observationdomainid` - An observation domain groups a set of Citrix ADCs based on deployment: cluster, HA etc. A unique Observation Domain ID is required to be assigned to each such group.
* `observationdomainname` - Name of the Observation Domain defined by the observation domain ID.
* `observationpointid` - An observation point ID is identifier for the NetScaler from which appflow records are being exported. By default, the NetScaler IP is the observation point ID.
* `securityinsightrecordinterval` - Interval, in seconds, at which to send security insight flow records to the configured collectors.
* `securityinsighttraffic` - Enable/disable the feature individually on appflow action.
* `skipcacheredirectionhttptransaction` - Skip Cache http transaction. This HTTP transaction is specific to Cache Redirection module. In Case of Cache Miss there will be another HTTP transaction initiated by the cache server.
* `subscriberawareness` - Enable this option for logging end user MSISDN in L4/L7 appflow records
* `subscriberidobfuscation` - Enable this option for obfuscating MSISDN in L4/L7 appflow records
* `subscriberidobfuscationalgo` - Algorithm(MD5 or SHA256) to be used for obfuscating MSISDN
* `tcpattackcounterinterval` - Interval, in seconds, at which to send tcp attack counters to the configured collectors. If 0 is configured, the record is not sent.
* `templaterefresh` - Refresh interval, in seconds, at which to export the template data. Because data transmission is in UDP, the templates must be resent at regular intervals.
* `timeseriesovernsip` - To use the Citrix ADC IP to send Time series data such as metrics and events, instead of the SNIP
* `udppmtu` - MTU, in bytes, for IPFIX UDP packets.
* `urlcategory` - Include the URL category record.
* `usagerecordinterval` - On enabling this option, the NGS will send bandwidth usage record to configured collectors.
* `videoinsight` - Enable/disable the feature individually on appflow action.
* `websaasappusagereporting` - On enabling this option, NGS will send data used by Web/saas app at the end of every HTTP transaction to configured collectors.

## Attribute Reference

* `id` - The id of the appflowparam. It is a system-generated identifier.
