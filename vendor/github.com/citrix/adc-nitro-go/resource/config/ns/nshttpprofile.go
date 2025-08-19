/*
* Copyright (c) 2021 Citrix Systems, Inc.
*
*   Licensed under the Apache License, Version 2.0 (the "License");
*   you may not use this file except in compliance with the License.
*   You may obtain a copy of the License at
*
*       http://www.apache.org/licenses/LICENSE-2.0
*
*  Unless required by applicable law or agreed to in writing, software
*   distributed under the License is distributed on an "AS IS" BASIS,
*   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
*   See the License for the specific language governing permissions and
*   limitations under the License.
*/

package ns

/**
* Configuration for HTTP profile resource.
*/
type Nshttpprofile struct {
	/**
	* Name for an HTTP profile. Must begin with a letter, number, or the underscore \(_\) character. Other characters allowed, after the first character, are the hyphen \(-\), period \(.\), hash \(\#\), space \( \), at \(@\), colon \(:\), and equal \(=\) characters. The name of a HTTP profile cannot be changed after it is created.
		CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks \(for example, "my http profile" or 'my http profile'\).
	*/
	Name string `json:"name,omitempty"`
	/**
	* Drop invalid HTTP requests or responses.
	*/
	Dropinvalreqs string `json:"dropinvalreqs,omitempty"`
	/**
	* Mark HTTP/0.9 requests as invalid.
	*/
	Markhttp09inval string `json:"markhttp09inval,omitempty"`
	/**
	* Mark CONNECT requests as invalid.
	*/
	Markconnreqinval string `json:"markconnreqinval,omitempty"`
	/**
	* Mark TRACE requests as invalid.
	*/
	Marktracereqinval string `json:"marktracereqinval,omitempty"`
	/**
	* Mark RFC7230 non-compliant transaction as invalid
	*/
	Markrfc7230noncompliantinval string `json:"markrfc7230noncompliantinval,omitempty"`
	/**
	* Mark Http header with extra white space as invalid
	*/
	Markhttpheaderextrawserror string `json:"markhttpheaderextrawserror,omitempty"`
	/**
	* Start data compression on receiving a TCP packet with PUSH flag set.
	*/
	Cmponpush string `json:"cmponpush,omitempty"`
	/**
	* Reuse server connections for requests from more than one client connections.
	*/
	Conmultiplex string `json:"conmultiplex,omitempty"`
	/**
	* Maximum limit on the number of connections, from the Citrix ADC to a particular server that are kept in the reuse pool. This setting is helpful for optimal memory utilization and for reducing the idle connections to the server just after the peak time. Zero implies no limit on reuse pool size. If non-zero value is given, it has to be greater than or equal to the number of running Packet Engines.
	*/
	Maxreusepool int `json:"maxreusepool,omitempty"`
	/**
	* Drop any extra 'CR' and 'LF' characters present after the header.
	*/
	Dropextracrlf string `json:"dropextracrlf,omitempty"`
	/**
	* Maximum time to wait, in milliseconds, between incomplete header packets. If the header packets take longer to arrive at Citrix ADC, the connection is silently dropped.
	*/
	Incomphdrdelay int `json:"incomphdrdelay,omitempty"`
	/**
	* HTTP connection to be upgraded to a web socket connection. Once upgraded, Citrix ADC does not process Layer 7 traffic on this connection.
	*/
	Websocket string `json:"websocket,omitempty"`
	/**
	* Allow RTSP tunnel in HTTP. Once application/x-rtsp-tunnelled is seen in Accept or Content-Type header, Citrix ADC does not process Layer 7 traffic on this connection.
	*/
	Rtsptunnel string `json:"rtsptunnel,omitempty"`
	/**
	* Time, in seconds, within which the HTTP request must complete. If the request does not complete within this time, the specified request timeout action is executed. Zero disables the timeout.
	*/
	Reqtimeout int `json:"reqtimeout,omitempty"`
	/**
	* Adapts the configured request timeout based on flow conditions. The timeout is increased or decreased internally and applied on the flow.
	*/
	Adpttimeout string `json:"adpttimeout,omitempty"`
	/**
	* Action to take when the HTTP request does not complete within the specified request timeout duration. You can configure the following actions:
		* RESET - Send RST (reset) to client when timeout occurs.
		* DROP - Drop silently when timeout occurs.
		* Custom responder action - Name of the responder action to trigger when timeout occurs, used to send custom message.
	*/
	Reqtimeoutaction string `json:"reqtimeoutaction,omitempty"`
	/**
	* Drop any extra data when server sends more data than the specified content-length.
	*/
	Dropextradata string `json:"dropextradata,omitempty"`
	/**
	* Enable or disable web logging.
	*/
	Weblog string `json:"weblog,omitempty"`
	/**
	* Name of the header that contains the real client IP address.
	*/
	Clientiphdrexpr string `json:"clientiphdrexpr,omitempty"`
	/**
	* Maximum number of requests allowed on a single connection. Zero implies no limit on the number of requests.
	*/
	Maxreq int `json:"maxreq,omitempty"`
	/**
	* Generate the persistent Citrix ADC specific ETag for the HTTP response with ETag header.
	*/
	Persistentetag string `json:"persistentetag,omitempty"`
	/**
	* Enable SPDYv2 or SPDYv3 or both over SSL vserver. SSL will advertise SPDY support either during NPN Handshake or when client will advertises SPDY support during ALPN handshake. Both SPDY versions are enabled when this parameter is set to ENABLED.
	*/
	Spdy string `json:"spdy,omitempty"`
	/**
	* Choose whether to enable support for HTTP/2.
	*/
	Http2 string `json:"http2,omitempty"`
	/**
	* Choose whether to enable support for Direct HTTP/2.
	*/
	Http2direct string `json:"http2direct,omitempty"`
	/**
	* Choose whether to enable strict HTTP/2 cipher selection
	*/
	Http2strictcipher string `json:"http2strictcipher,omitempty"`
	/**
	* Choose whether to enable support for sending HTTP/2 ALTSVC frames. When enabled, the ADC sends HTTP/2 ALTSVC frames to HTTP/2 clients, instead of the Alt-Svc response header field. Not applicable to servers.
	*/
	Http2altsvcframe string `json:"http2altsvcframe,omitempty"`
	/**
	* Choose whether to enable support for Alternative Services.
	*/
	Altsvc string `json:"altsvc,omitempty"`
	/**
	* Configure a custom Alternative Services header value that should be inserted in the response to advertise a HTTP/SSL/HTTP_QUIC vserver.
	*/
	Altsvcvalue string `json:"altsvcvalue,omitempty"`
	/**
	* Idle timeout (in seconds) for server connections in re-use pool. Connections in the re-use pool are flushed, if they remain idle for the configured timeout.
	*/
	Reusepooltimeout int `json:"reusepooltimeout,omitempty"`
	/**
	* Number of bytes to be queued to look for complete header before returning error. If complete header is not obtained after queuing these many bytes, request will be marked as invalid and no L7 processing will be done for that TCP connection.
	*/
	Maxheaderlen int `json:"maxheaderlen,omitempty"`
	/**
	* Minimum limit on the number of connections, from the Citrix ADC to a particular server that are kept in the reuse pool. This setting is helpful for optimal memory utilization and for reducing the idle connections to the server just after the peak time. Zero implies no limit on reuse pool size.
	*/
	Minreusepool int `json:"minreusepool,omitempty"`
	/**
	* Maximum size of header list that the Citrix ADC is prepared to accept, in bytes. NOTE: The actual plain text header size that the Citrix ADC accepts is limited by maxHeaderLen. Please change maxHeaderLen parameter as well when modifying http2MaxHeaderListSize.
	*/
	Http2maxheaderlistsize int `json:"http2maxheaderlistsize,omitempty"`
	/**
	* Maximum size of the frame payload that the Citrix ADC is willing to receive, in bytes.
	*/
	Http2maxframesize int `json:"http2maxframesize,omitempty"`
	/**
	* Maximum number of concurrent streams that is allowed per connection.
	*/
	Http2maxconcurrentstreams int `json:"http2maxconcurrentstreams,omitempty"`
	/**
	* Initial window size for connection level flow control, in bytes.
	*/
	Http2initialconnwindowsize int `json:"http2initialconnwindowsize,omitempty"`
	/**
	* Initial window size for stream level flow control, in bytes.
	*/
	Http2initialwindowsize int `json:"http2initialwindowsize,omitempty"`
	/**
	* Maximum size of the header compression table used to decode header blocks, in bytes.
	*/
	Http2headertablesize int `json:"http2headertablesize,omitempty"`
	/**
	* Minimum number of HTTP2 connections established to backend server, on receiving HTTP requests from client before multiplexing the streams into the available HTTP/2 connections.
	*/
	Http2minseverconn int `json:"http2minseverconn,omitempty"`
	/**
	* Maximum number of ping frames allowed in HTTP2 connection per minute
	*/
	Http2maxpingframespermin int `json:"http2maxpingframespermin,omitempty"`
	/**
	* Maximum number of settings frames allowed in HTTP2 connection per minute
	*/
	Http2maxsettingsframespermin int `json:"http2maxsettingsframespermin,omitempty"`
	/**
	* Maximum number of reset frames allowed in HTTP/2 connection per minute
	*/
	Http2maxresetframespermin int `json:"http2maxresetframespermin,omitempty"`
	/**
	* Maximum number of empty  frames allowed in HTTP2 connection per minute
	*/
	Http2maxemptyframespermin int `json:"http2maxemptyframespermin,omitempty"`
	/**
	* Maximum size in bytes allowed to buffer gRPC packets till trailer is received
	*/
	Grpcholdlimit int `json:"grpcholdlimit,omitempty"`
	/**
	* Maximum time in milliseconds allowed to buffer gRPC packets till trailer is received. The value should be in multiples of 100
	*/
	Grpcholdtimeout int `json:"grpcholdtimeout,omitempty"`
	/**
	* Set to DISABLED for gRPC without a length delimitation.
	*/
	Grpclengthdelimitation string `json:"grpclengthdelimitation,omitempty"`
	/**
	* This option sets the satisfactory threshold (T) for client response time in milliseconds to be used for APDEX calculations. This means a transaction responding in less than this threshold is considered satisfactory. Transaction responding between T and 4*T is considered tolerable. Any transaction responding in more than 4*T time is considered frustrating. Citrix ADC maintains stats for such tolerable and frustrating transcations. And client response time related apdex counters are only updated on a vserver which receives clients traffic.
	*/
	Apdexcltresptimethreshold int `json:"apdexcltresptimethreshold,omitempty"`
	/**
	* Choose whether to enable support for HTTP/3.
	*/
	Http3 string `json:"http3,omitempty"`
	/**
	* Maximum size of the HTTP/3 header field section, in bytes.
	*/
	Http3maxheaderfieldsectionsize int `json:"http3maxheaderfieldsectionsize,omitempty"`
	/**
	* Maximum size of the HTTP/3 QPACK dynamic header table, in bytes.
	*/
	Http3maxheadertablesize int `json:"http3maxheadertablesize,omitempty"`
	/**
	* Maximum number of HTTP/3 streams that can be blocked while HTTP/3 headers are being decoded.
	*/
	Http3maxheaderblockedstreams int `json:"http3maxheaderblockedstreams,omitempty"`
	/**
	* Number of bytes allowed for header field for HTTP header. If number of bytes exceeds beyond configured value, then request will be marked invalid
	*/
	Maxheaderfieldlen int `json:"maxheaderfieldlen,omitempty"`
	/**
	* Maximum number of incoming RST_STREAM frames allowed in HTTP/2 connection per minute
	*/
	Http2maxrxresetframespermin int `json:"http2maxrxresetframespermin,omitempty"`
	/**
	* Choose whether to enable support for WebTransport over HTTP/3.
	*/
	Http3webtransport string `json:"http3webtransport,omitempty"`
	/**
	* Minimum number of HTTP/3 connections established to backend server, on receiving HTTP requests from client before multiplexing the streams into the available HTTP/3 connections.
	*/
	Http3minseverconn int `json:"http3minseverconn,omitempty"`
	/**
	* Application pipeline request buffering size, in bytes.
	*/
	Httppipelinebuffsize int `json:"httppipelinebuffsize,omitempty"`
	/**
	* When enabled allows only the word characters [A-Za-z0-9_] and hyphen [-] in the request/response header names and the connection will be reset for the other characters. When disabled allows any visible (printing) characters (%21-%7E) except delimiters (double quotes and "(),/:;<=>?@[]{}").
	*/
	Allowonlywordcharactersandhyphen string `json:"allowonlywordcharactersandhyphen,omitempty"`
	/**
	* Validates the length of the Host header and its syntax. Also includes validation of the port number if specified
	*/
	Hostheadervalidation string `json:"hostheadervalidation,omitempty"`
	/**
	* Maximum number of allowed occurrences of header fields that share the same field name. This threshold is enforced only for well-known header field names recognized by NetScaler. If the value is set to 0, then it will be similar to previous behavior, Where we store only 15 duplicate headers and rest are parsed and send to the server.
	*/
	Maxduplicateheaderfields int `json:"maxduplicateheaderfields,omitempty"`
	/**
	* Pass protocol upgrade request to the server.
	*/
	Passprotocolupgrade string `json:"passprotocolupgrade,omitempty"`
	/**
	* Choose whether to enable HTTP/2 Extended CONNECT mechanism.
	*/
	Http2extendedconnect string `json:"http2extendedconnect,omitempty"`
	//------- Read only Parameter ---------;

	Refcnt string `json:"refcnt,omitempty"`
	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`
	Apdexsvrresptimethreshold string `json:"apdexsvrresptimethreshold,omitempty"`
	Dropinvalreqswarning string `json:"dropinvalreqswarning,omitempty"`

}
