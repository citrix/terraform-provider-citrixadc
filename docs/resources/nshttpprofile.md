---
subcategory: "NS"
---

# Resource: nshttpprofile

The nshttpprofile resource is used to create HTTP profiles.


## Example usage

```hcl
resource "citrixadc_nshttpprofile" "tf_httpprofile" {
    name  = "tf_httpprofile"
    http2 = "ENABLED"
}  
```


## Argument Reference

* `name` - (Required) Name for an HTTP profile.The name of a HTTP profile cannot be changed after it is created.
* `dropinvalreqs` - (Optional) Drop invalid HTTP requests or responses. Possible values: [ ENABLED, DISABLED ]
* `markhttp09inval` - (Optional) Mark HTTP/0.9 requests as invalid. Possible values: [ ENABLED, DISABLED ]
* `markconnreqinval` - (Optional) Mark CONNECT requests as invalid. Possible values: [ ENABLED, DISABLED ]
* `marktracereqinval` - (Optional) Mark TRACE requests as invalid. Possible values: [ ENABLED, DISABLED ]
* `cmponpush` - (Optional) Start data compression on receiving a TCP packet with PUSH flag set. Possible values: [ ENABLED, DISABLED ]
* `conmultiplex` - (Optional) Reuse server connections for requests from more than one client connections. Possible values: [ ENABLED, DISABLED ]
* `maxreusepool` - (Optional) 
* `dropextracrlf` - (Optional) Drop any extra 'CR' and 'LF' characters present after the header. Possible values: [ ENABLED, DISABLED ]
* `incomphdrdelay` - (Optional) 
* `websocket` - (Optional) HTTP connection to be upgraded to a web socket connection. Once upgraded, Citrix ADC does not process Layer 7 traffic on this connection. Possible values: [ ENABLED, DISABLED ]
* `rtsptunnel` - (Optional) Allow RTSP tunnel in HTTP. Once application/x-rtsp-tunnelled is seen in Accept or Content-Type header, Citrix ADC does not process Layer 7 traffic on this connection. Possible values: [ ENABLED, DISABLED ]
* `reqtimeout` - (Optional) Time, in seconds, within which the HTTP request must complete. If the request does not complete within this time, the specified request timeout action is executed. Zero disables the timeout.
* `adpttimeout` - (Optional) Adapts the configured request timeout based on flow conditions. The timeout is increased or decreased internally and applied on the flow. Possible values: [ ENABLED, DISABLED ]
* `reqtimeoutaction` - (Optional) Action to take when the HTTP request does not complete within the specified request timeout duration. You can configure the following actions: * RESET - Send RST (reset) to client when timeout occurs. * DROP - Drop silently when timeout occurs. * Custom responder action - Name of the responder action to trigger when timeout occurs, used to send custom message.
* `dropextradata` - (Optional) Drop any extra data when server sends more data than the specified content-length. Possible values: [ ENABLED, DISABLED ]
* `weblog` - (Optional) Enable or disable web logging. Possible values: [ ENABLED, DISABLED ]
* `clientiphdrexpr` - (Optional) Name of the header that contains the real client IP address.
* `maxreq` - (Optional) 
* `persistentetag` - (Optional) Generate the persistent Citrix ADC specific ETag for the HTTP response with ETag header. Possible values: [ ENABLED, DISABLED ]
* `spdy` - (Optional) Enable SPDYv2 or SPDYv3 or both over SSL vserver. SSL will advertise SPDY support either during NPN Handshake or when client will advertises SPDY support during ALPN handshake. Both SPDY versions are enabled when this parameter is set to ENABLED. Possible values: [ DISABLED, ENABLED, V2, V3 ]
* `http2` - (Optional) Choose whether to enable support for HTTP/2. Possible values: [ ENABLED, DISABLED ]
* `http2direct` - (Optional) Choose whether to enable support for Direct HTTP/2. Possible values: [ ENABLED, DISABLED ]
* `http2strictcipher` - (Optional) Choose whether to enable strict HTTP/2 cipher selection. Possible values: [ ENABLED, DISABLED ]
* `altsvc` - (Optional) Choose whether to enable support for Alternative Service. Possible values: [ ENABLED, DISABLED ]
* `reusepooltimeout` - (Optional) Idle timeout (in seconds) for server connections in re-use pool. Connections in the re-use pool are flushed, if they remain idle for the configured timeout.
* `maxheaderlen` - (Optional) Number of bytes to be queued to look for complete header before returning error. If complete header is not obtained after queuing these many bytes, request will be marked as invalid and no L7 processing will be done for that TCP connection.
* `minreusepool` - (Optional) 
* `http2maxheaderlistsize` - (Optional) 
* `http2maxframesize` - (Optional) 
* `http2maxconcurrentstreams` - (Optional) 
* `http2initialwindowsize` - (Optional) Initial window size for stream level flow control, in bytes.
* `http2headertablesize` - (Optional) 
* `http2minseverconn` - (Optional) 
* `apdexcltresptimethreshold` - (Optional) This option sets the satisfactory threshold (T) for client response time in milliseconds to be used for APDEX calculations. This means a transaction responding in less than this threshold is considered satisfactory. Transaction responding between T and 4\*T is considered tolerable. Any transaction responding in more than 4\*T time is considered frustrating. Citrix ADC maintains stats for such tolerable and frustrating transcations. And client response time related apdex counters are only updated on a vserver which receives clients traffic.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nshttpprofile. It has the same value as the `name` attribute.


## Import

A nshttpprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_nshttpprofile.tf_httpprofile tf_httpprofile
```
