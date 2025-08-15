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

* `adpttimeout` - (Optional) Adapts the configured request timeout based on flow conditions. The timeout is increased or decreased internally and applied on the flow.
* `allowonlywordcharactersandhyphen` - (Optional) Allow only word characters [A-Za-z0-9_] and hyphen [-] in the http request/response header names.
* `altsvc` - (Optional) Choose whether to enable support for Alternative Services.
* `altsvcvalue` - (Optional) Configure a custom Alternative Services header value that should be inserted in the response to advertise a HTTP/SSL/HTTP_QUIC vserver.
* `apdexcltresptimethreshold` - (Optional) This option sets the satisfactory threshold (T) for client response time in milliseconds to be used for APDEX calculations. This means a transaction responding in less than this threshold is considered satisfactory. Transaction responding between T and 4\*T is considered tolerable. Any transaction responding in more than 4\*T time is considered frustrating. Citrix ADC maintains stats for such tolerable and frustrating transcations. And client response time related apdex counters are only updated on a vserver which receives clients traffic.
* `clientiphdrexpr` - (Optional) Name of the header that contains the real client IP address.
* `cmponpush` - (Optional) Start data compression on receiving a TCP packet with PUSH flag set.
* `conmultiplex` - (Optional) Reuse server connections for requests from more than one client connections.
* `dropextracrlf` - (Optional) Drop any extra 'CR' and 'LF' characters present after the header.
* `dropextradata` - (Optional) Drop any extra data when server sends more data than the specified content-length.
* `dropinvalreqs` - (Optional) Drop invalid HTTP requests or responses.
* `grpcholdlimit` - (Optional)
* `grpcholdtimeout` - (Optional)
* `grpclengthdelimitation` - (Optional) Set to DISABLED for gRPC without a length delimitation.
* `http2` - (Optional) Choose whether to enable support for HTTP/2.
* `http2altsvcframe` - (Optional) Choose whether to enable support for sending HTTP/2 ALTSVC frames. When enabled, the ADC sends HTTP/2 ALTSVC frames to HTTP/2 clients, instead of the Alt-Svc response header field. Not applicable to servers.
* `http2direct` - (Optional) Choose whether to enable support for Direct HTTP/2.
* `http2headertablesize` - (Optional)
* `http2initialconnwindowsize` - (Optional) Initial window size for connection level flow control, in bytes.
* `http2initialwindowsize` - (Optional) Initial window size for stream level flow control, in bytes.
* `http2maxconcurrentstreams` - (Optional)
* `http2maxemptyframespermin` - (Optional)
* `http2maxframesize` - (Optional)
* `http2maxheaderlistsize` - (Optional)
* `http2maxpingframespermin` - (Optional)
* `http2maxresetframespermin` - (Optional)
* `http2maxsettingsframespermin` - (Optional)
* `http2minseverconn` - (Optional)
* `http2strictcipher` - (Optional) Choose whether to enable strict HTTP/2 cipher selection
* `http3` - (Optional) Choose whether to enable support for HTTP/3.
* `http3maxheaderblockedstreams` - (Optional)
* `http3maxheaderfieldsectionsize` - (Optional)
* `http3maxheadertablesize` - (Optional)
* `incomphdrdelay` - (Optional)
* `markconnreqinval` - (Optional) Mark CONNECT requests as invalid.
* `markhttp09inval` - (Optional) Mark HTTP/0.9 requests as invalid.
* `markhttpheaderextrawserror` - (Optional) Mark Http header with extra white space as invalid
* `markrfc7230noncompliantinval` - (Optional) Mark RFC7230 non-compliant transaction as invalid
* `marktracereqinval` - (Optional) Mark TRACE requests as invalid.
* `maxheaderlen` - (Optional) Number of bytes to be queued to look for complete header before returning error. If complete header is not obtained after queuing these many bytes, request will be marked as invalid and no L7 processing will be done for that TCP connection.
* `maxreq` - (Optional)
* `maxreusepool` - (Optional)
* `minreusepool` - (Optional)
* `name` - (Required) Name for an HTTP profile. Must begin with a letter, number, or the underscore \(\_\) character. Other characters allowed, after the first character, are the hyphen \(-\), period \(.\), hash \(\#\), space \( \), at \(@\), colon \(:\), and equal \(=\) characters. The name of a HTTP profile cannot be changed after it is created.  CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks \(for example, "my http profile" or 'my http profile'\).
* `persistentetag` - (Optional) Generate the persistent Citrix ADC specific ETag for the HTTP response with ETag header.
* `reqtimeout` - (Optional) Time, in seconds, within which the HTTP request must complete. If the request does not complete within this time, the specified request timeout action is executed. Zero disables the timeout.
* `reqtimeoutaction` - (Optional) Action to take when the HTTP request does not complete within the specified request timeout duration. You can configure the following actions: * RESET - Send RST (reset) to client when timeout occurs. * DROP - Drop silently when timeout occurs. * Custom responder action - Name of the responder action to trigger when timeout occurs, used to send custom message.
* `reusepooltimeout` - (Optional) Idle timeout (in seconds) for server connections in re-use pool. Connections in the re-use pool are flushed, if they remain idle for the configured timeout.
* `rtsptunnel` - (Optional) Allow RTSP tunnel in HTTP. Once application/x-rtsp-tunnelled is seen in Accept or Content-Type header, Citrix ADC does not process Layer 7 traffic on this connection.
* `spdy` - (Optional) Enable SPDYv2 or SPDYv3 or both over SSL vserver. SSL will advertise SPDY support either during NPN Handshake or when client will advertises SPDY support during ALPN handshake. Both SPDY versions are enabled when this parameter is set to ENABLED.
* `weblog` - (Optional) Enable or disable web logging.
* `websocket` - (Optional) HTTP connection to be upgraded to a web socket connection. Once upgraded, Citrix ADC does not process Layer 7 traffic on this connection.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nshttpprofile. It has the same value as the `name` attribute.


## Import

A nshttpprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_nshttpprofile.tf_httpprofile tf_httpprofile
```
