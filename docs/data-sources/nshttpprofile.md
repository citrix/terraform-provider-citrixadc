---
subcategory: "NS"
---

# Data Source: citrixadc_nshttpprofile

The `citrixadc_nshttpprofile` data source allows you to retrieve information about an HTTP profile configured on the Citrix ADC.

## Example Usage

```terraform
data "citrixadc_nshttpprofile" "example" {
  name = "my_http_profile"
}

output "http2_enabled" {
  value = data.citrixadc_nshttpprofile.example.http2
}

output "max_reuse_pool" {
  value = data.citrixadc_nshttpprofile.example.maxreusepool
}
```

## Argument Reference

* `name` - (Required) Name for the HTTP profile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.), pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore characters. Cannot be changed after the profile is created.

## Attribute Reference

In addition to the argument, the following attributes are available:

* `id` - The ID of the HTTP profile. It has the same value as the `name` attribute.
* `adpttimeout` - Adapts the configured request timeout based on flow conditions. The timeout is increased or decreased internally and applied on the flow.
* `allowonlywordcharactersandhyphen` - When enabled, allows only word characters [A-Za-z0-9_] and hyphen [-] in the request/response header names.
* `altsvc` - Enable or disable support for Alternative Services.
* `altsvcvalue` - Custom Alternative Services header value to advertise a HTTP/SSL/HTTP_QUIC vserver.
* `apdexcltresptimethreshold` - Satisfactory threshold (T) for client response time in milliseconds for APDEX calculations.
* `clientiphdrexpr` - Name of the header that contains the real client IP address.
* `cmponpush` - Start data compression on receiving a TCP packet with PUSH flag set.
* `conmultiplex` - Reuse server connections for requests from more than one client connection.
* `dropextracrlf` - Drop any extra 'CR' and 'LF' characters present after the header.
* `dropextradata` - Drop any extra data when server sends more data than the specified content-length.
* `dropinvalreqs` - Drop invalid HTTP requests or responses.
* `grpcholdlimit` - Maximum size in bytes allowed to buffer gRPC packets till trailer is received.
* `grpcholdtimeout` - Maximum time in milliseconds allowed to buffer gRPC packets till trailer is received.
* `grpclengthdelimitation` - Set to DISABLED for gRPC without a length delimitation.
* `hostheadervalidation` - Validates the length of the Host header and its syntax.
* `http2` - Enable or disable support for HTTP/2.
* `http2altsvcframe` - Enable or disable support for sending HTTP/2 ALTSVC frames.
* `http2direct` - Enable or disable support for Direct HTTP/2.
* `http2extendedconnect` - Enable or disable HTTP/2 Extended CONNECT mechanism.
* `http2headertablesize` - Maximum size of the header compression table used to decode header blocks, in bytes.
* `http2initialconnwindowsize` - Initial window size for connection level flow control, in bytes.
* `http2initialwindowsize` - Initial window size for stream level flow control, in bytes.
* `http2maxconcurrentstreams` - Maximum number of concurrent streams allowed per connection.
* `http2maxemptyframespermin` - Maximum number of empty frames allowed in HTTP2 connection per minute.
* `http2maxframesize` - Maximum size of the frame payload that the Citrix ADC is willing to receive, in bytes.
* `http2maxheaderlistsize` - Maximum size of header list that the Citrix ADC is prepared to accept, in bytes.
* `http2maxpingframespermin` - Maximum number of PING frames allowed in HTTP2 connection per minute.
* `http2maxresetframespermin` - Maximum number of outgoing RST_STREAM frames allowed in HTTP/2 connection per minute.
* `http2maxrxresetframespermin` - Maximum number of incoming RST_STREAM frames allowed in HTTP/2 connection per minute.
* `http2maxsettingsframespermin` - Maximum number of SETTINGS frames allowed in HTTP2 connection per minute.
* `http2minseverconn` - Minimum number of HTTP2 connections established to backend server.
* `http2strictcipher` - Enable or disable strict HTTP/2 cipher selection.
* `http3` - Enable or disable support for HTTP/3.
* `http3maxheaderblockedstreams` - Maximum number of HTTP/3 streams that can be blocked while HTTP/3 headers are being decoded.
* `http3maxheaderfieldsectionsize` - Maximum size of the HTTP/3 header field section, in bytes.
* `http3maxheadertablesize` - Maximum size of the HTTP/3 QPACK dynamic header table, in bytes.
* `http3minseverconn` - Minimum number of HTTP/3 connections established to backend server.
* `http3webtransport` - Enable or disable HTTP/3 WebTransport support.
* `httppipelinebuffsize` - Maximum size in bytes for HTTP pipeline buffer.
* `incomphdrdelay` - Maximum time in milliseconds to wait for a complete HTTP header to arrive.
* `markconnreqinval` - Mark CONNECT requests as invalid.
* `markhttp09inval` - Mark HTTP/0.9 requests as invalid.
* `markhttpheaderextrawserror` - Mark HTTP requests with extra whitespace as errors.
* `markrfc7230noncompliantinval` - Mark requests that do not comply with RFC 7230 as invalid.
* `marktracereqinval` - Mark TRACE requests as invalid.
* `maxduplicateheaderfields` - Maximum number of duplicate header fields allowed.
* `maxheaderfieldlen` - Maximum length of an HTTP header field, in bytes.
* `maxheaderlen` - Maximum length of all HTTP headers combined, in bytes.
* `maxreq` - Maximum number of HTTP requests allowed on a persistent connection.
* `maxreusepool` - Maximum size of connection pool for reusing server connections.
* `minreusepool` - Minimum size of connection pool for reusing server connections.
* `passprotocolupgrade` - Pass through protocol upgrade requests.
* `persistentetag` - Enable or disable generation of persistent entity tags.
* `reqtimeout` - Request timeout in milliseconds.
* `reqtimeoutaction` - Action to take when request times out.
* `reusepooltimeout` - Timeout in seconds for idle connections in the reuse pool.
* `rtsptunnel` - Enable or disable RTSP tunneling.
* `weblog` - Enable or disable web logging.
* `websocket` - Enable or disable WebSocket support.
