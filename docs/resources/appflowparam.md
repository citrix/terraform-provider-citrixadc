---
subcategory: "AppFlow"
---

# Resource: appflowparam

The appflowparam resource is used to create appflowparam.


## Example usage

```hcl
resource "citrixadc_appflowparam" "tf_appflowparam" {
  templaterefresh     = 200
  flowrecordinterval  = 100
  httpcookie          = "ENABLED"
  httplocation        = "ENABLED"
}

```


## Argument Reference

* `aaausername` - (Optional) Enable AppFlow AAA Username logging.
* `analyticsauthtoken` - (Optional) Authentication token to be set by the agent.
* `appnamerefresh` - (Optional) Interval, in seconds, at which to send Appnames to the configured collectors. Appname refers to the name of an entity (virtual server, service, or service group) in the Citrix ADC.
* `auditlogs` - (Optional) Enable Auditlogs to be sent to the Telemetry Agent
* `cacheinsight` - (Optional) Flag to determine whether cache records need to be exported or not. If this flag is true and IC is enabled, cache records are exported instead of L7 HTTP records
* `clienttrafficonly` - (Optional) Generate AppFlow records for only the traffic from the client.
* `connectionchaining` - (Optional) Enable connection chaining so that the client server flows of a connection are linked. Also the connection chain ID is propagated across Citrix ADCs, so that in a multi-hop environment the flows belonging to the same logical connection are linked. This id is also logged as part of appflow record
* `cqareporting` - (Optional) TCP CQA reporting enable/disable knob.
* `distributedtracing` - (Optional) Enable generation of the distributed tracing templates in the Appflow records
* `disttracingsamplingrate` - (Optional) Sampling rate for Distributed Tracing
* `emailaddress` - (Optional) Enable AppFlow user email-id logging.
* `events` - (Optional) Enable Events to be sent to the Telemetry Agent
* `flowrecordinterval` - (Optional) Interval, in seconds, at which to send flow records to the configured collectors.
* `gxsessionreporting` - (Optional) Enable this option for Gx session reporting
* `httpauthorization` - (Optional) Include the HTTP Authorization header information.
* `httpcontenttype` - (Optional) Include the HTTP Content-Type header sent from the server to the client to determine the type of the content sent.
* `httpcookie` - (Optional) Include the cookie that was in the HTTP request the appliance received from the client.
* `httpdomain` - (Optional) Include the http domain request to be exported.
* `httphost` - (Optional) Include the host identified in the HTTP request that the appliance received from the client.
* `httplocation` - (Optional) Include the HTTP location headers returned from the HTTP responses.
* `httpmethod` - (Optional) Include the method that was specified in the HTTP request that the appliance received from the client.
* `httpquerywithurl` - (Optional) Include the HTTP query segment along with the URL that the Citrix ADC received from the client.
* `httpreferer` - (Optional) Include the web page that was last visited by the client.
* `httpsetcookie` - (Optional) Include the Set-cookie header sent from the server to the client in response to a HTTP request.
* `httpsetcookie2` - (Optional) Include the Set-cookie header sent from the server to the client in response to a HTTP request.
* `httpurl` - (Optional) Include the http URL that the Citrix ADC received from the client.
* `httpuseragent` - (Optional) Include the client application through which the HTTP request was received by the Citrix ADC.
* `httpvia` - (Optional) Include the httpVia header which contains the IP address of proxy server through which the client accessed the server.
* `httpxforwardedfor` - (Optional) Include the httpXForwardedFor header, which contains the original IP Address of the client using a proxy server to access the server.
* `identifiername` - (Optional) Include the stream identifier name to be exported.
* `identifiersessionname` - (Optional) Include the stream identifier session name to be exported.
* `logstreamovernsip` - (Optional) To use the Citrix ADC IP to send Logstream records instead of the SNIP
* `lsnlogging` - (Optional) On enabling this option, the Citrix ADC will send the Large Scale Nat(LSN) records to the configured collectors.
* `metrics` - (Optional) Enable Citrix ADC Stats to be sent to the Telemetry Agent
* `observationdomainid` - (Optional) An observation domain groups a set of Citrix ADCs based on deployment: cluster, HA etc. A unique Observation Domain ID is required to be assigned to each such group.
* `observationdomainname` - (Optional) Name of the Observation Domain defined by the observation domain ID.
* `observationpointid` - (Optional) An observation point ID is identifier for the NetScaler from which appflow records are being exported. By default, the NetScaler IP is the observation point ID.
* `securityinsightrecordinterval` - (Optional) Interval, in seconds, at which to send security insight flow records to the configured collectors.
* `securityinsighttraffic` - (Optional) Enable/disable the feature individually on appflow action.
* `skipcacheredirectionhttptransaction` - (Optional) Skip Cache http transaction. This HTTP transaction is specific to Cache Redirection module. In Case of Cache Miss there will be another HTTP transaction initiated by the cache server.
* `subscriberawareness` - (Optional) Enable this option for logging end user MSISDN in L4/L7 appflow records
* `subscriberidobfuscation` - (Optional) Enable this option for obfuscating MSISDN in L4/L7 appflow records
* `subscriberidobfuscationalgo` - (Optional) Algorithm(MD5 or SHA256) to be used for obfuscating MSISDN
* `tcpattackcounterinterval` - (Optional) Interval, in seconds, at which to send tcp attack counters to the configured collectors. If 0 is configured, the record is not sent.
* `templaterefresh` - (Optional) Refresh interval, in seconds, at which to export the template data. Because data transmission is in UDP, the templates must be resent at regular intervals.
* `timeseriesovernsip` - (Optional) To use the Citrix ADC IP to send Time series data such as metrics and events, instead of the SNIP
* `udppmtu` - (Optional) MTU, in bytes, for IPFIX UDP packets.
* `urlcategory` - (Optional) Include the URL category record.
* `usagerecordinterval` - (Optional) On enabling this option, the NGS will send bandwidth usage record to configured collectors.
* `videoinsight` - (Optional) Enable/disable the feature individually on appflow action.
* `websaasappusagereporting` - (Optional) On enabling this option, NGS will send data used by Web/saas app at the end of every HTTP transaction to configured collectors.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appflowparam is a unique string prefixed with "tf-appflowparam-"