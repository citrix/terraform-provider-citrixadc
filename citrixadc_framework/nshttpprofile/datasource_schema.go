package nshttpprofile

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func NshttpprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"adpttimeout": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Adapts the configured request timeout based on flow conditions. The timeout is increased or decreased internally and applied on the flow.",
			},
			"allowonlywordcharactersandhyphen": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "When enabled allows only the word characters [A-Za-z0-9_] and hyphen [-] in the request/response header names and the connection will be reset for the other characters. When disabled allows any visible (printing) characters (%21-%7E) except delimiters (double quotes and \"(),/:;<=>?@[]{}\").",
			},
			"altsvc": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Choose whether to enable support for Alternative Services.",
			},
			"altsvcvalue": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Configure a custom Alternative Services header value that should be inserted in the response to advertise a HTTP/SSL/HTTP_QUIC vserver.",
			},
			"apdexcltresptimethreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "This option sets the satisfactory threshold (T) for client response time in milliseconds to be used for APDEX calculations. This means a transaction responding in less than this threshold is considered satisfactory. Transaction responding between T and 4*T is considered tolerable. Any transaction responding in more than 4*T time is considered frustrating. Citrix ADC maintains stats for such tolerable and frustrating transcations. And client response time related apdex counters are only updated on a vserver which receives clients traffic.",
			},
			"clientiphdrexpr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the header that contains the real client IP address.",
			},
			"cmponpush": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Start data compression on receiving a TCP packet with PUSH flag set.",
			},
			"conmultiplex": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Reuse server connections for requests from more than one client connections.",
			},
			"dropextracrlf": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Drop any extra 'CR' and 'LF' characters present after the header.",
			},
			"dropextradata": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Drop any extra data when server sends more data than the specified content-length.",
			},
			"dropinvalreqs": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Drop invalid HTTP requests or responses.",
			},
			"grpcholdlimit": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum size in bytes allowed to buffer gRPC packets till trailer is received",
			},
			"grpcholdtimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum time in milliseconds allowed to buffer gRPC packets till trailer is received. The value should be in multiples of 100",
			},
			"grpclengthdelimitation": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Set to DISABLED for gRPC without a length delimitation.",
			},
			"hostheadervalidation": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Validates the length of the Host header and its syntax. Also includes validation of the port number if specified",
			},
			"http2": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Choose whether to enable support for HTTP/2.",
			},
			"http2altsvcframe": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Choose whether to enable support for sending HTTP/2 ALTSVC frames. When enabled, the ADC sends HTTP/2 ALTSVC frames to HTTP/2 clients, instead of the Alt-Svc response header field. Not applicable to servers.",
			},
			"http2direct": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Choose whether to enable support for Direct HTTP/2.",
			},
			"http2extendedconnect": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Choose whether to enable HTTP/2 Extended CONNECT mechanism.",
			},
			"http2headertablesize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum size of the header compression table used to decode header blocks, in bytes.",
			},
			"http2initialconnwindowsize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Initial window size for connection level flow control, in bytes.",
			},
			"http2initialwindowsize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Initial window size for stream level flow control, in bytes.",
			},
			"http2maxconcurrentstreams": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of concurrent streams that is allowed per connection.",
			},
			"http2maxemptyframespermin": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of empty frames allowed in HTTP2 connection per minute",
			},
			"http2maxframesize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum size of the frame payload that the Citrix ADC is willing to receive, in bytes.",
			},
			"http2maxheaderlistsize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum size of header list that the Citrix ADC is prepared to accept, in bytes. NOTE: The actual plain text header size that the Citrix ADC accepts is limited by maxHeaderLen. Please change maxHeaderLen parameter as well when modifying http2MaxHeaderListSize.",
			},
			"http2maxpingframespermin": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of PING frames allowed in HTTP2 connection per minute",
			},
			"http2maxresetframespermin": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of outgoing RST_STREAM frames allowed in HTTP/2 connection per minute",
			},
			"http2maxrxresetframespermin": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of incoming RST_STREAM frames allowed in HTTP/2 connection per minute",
			},
			"http2maxsettingsframespermin": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of SETTINGS frames allowed in HTTP2 connection per minute",
			},
			"http2minseverconn": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum number of HTTP2 connections established to backend server, on receiving HTTP requests from client before multiplexing the streams into the available HTTP/2 connections.",
			},
			"http2strictcipher": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Choose whether to enable strict HTTP/2 cipher selection",
			},
			"http3": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Choose whether to enable support for HTTP/3.",
			},
			"http3maxheaderblockedstreams": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of HTTP/3 streams that can be blocked while HTTP/3 headers are being decoded.",
			},
			"http3maxheaderfieldsectionsize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum size of the HTTP/3 header field section, in bytes.",
			},
			"http3maxheadertablesize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum size of the HTTP/3 QPACK dynamic header table, in bytes.",
			},
			"http3minseverconn": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum number of HTTP/3 connections established to backend server, on receiving HTTP requests from client before multiplexing the streams into the available HTTP/3 connections.",
			},
			"http3webtransport": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Choose whether to enable support for WebTransport over HTTP/3.",
			},
			"httppipelinebuffsize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Application pipeline request buffering size, in bytes.",
			},
			"incomphdrdelay": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum time to wait, in milliseconds, between incomplete header packets. If the header packets take longer to arrive at Citrix ADC, the connection is silently dropped.",
			},
			"markconnreqinval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Mark CONNECT requests as invalid.",
			},
			"markhttp09inval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Mark HTTP/0.9 requests as invalid.",
			},
			"markhttpheaderextrawserror": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Mark Http header with extra white space as invalid",
			},
			"markrfc7230noncompliantinval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Mark RFC7230 non-compliant transaction as invalid",
			},
			"marktracereqinval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Mark TRACE requests as invalid.",
			},
			"maxduplicateheaderfields": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of allowed occurrences of header fields that share the same field name. This threshold is enforced only for well-known header field names recognized by NetScaler. If the value is set to 0, then it will be similar to previous behavior, Where we store only 15 duplicate headers and rest are parsed and send to the server.",
			},
			"maxheaderfieldlen": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of bytes allowed for header field for HTTP header. If number of bytes exceeds beyond configured value, then request will be marked invalid",
			},
			"maxheaderlen": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of bytes to be queued to look for complete header before returning error. If complete header is not obtained after queuing these many bytes, request will be marked as invalid and no L7 processing will be done for that TCP connection.",
			},
			"maxreq": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of requests allowed on a single connection. Zero implies no limit on the number of requests.",
			},
			"maxreusepool": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum limit on the number of connections, from the Citrix ADC to a particular server that are kept in the reuse pool. This setting is helpful for optimal memory utilization and for reducing the idle connections to the server just after the peak time. Zero implies no limit on reuse pool size. If non-zero value is given, it has to be greater than or equal to the number of running Packet Engines.",
			},
			"minreusepool": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum limit on the number of connections, from the Citrix ADC to a particular server that are kept in the reuse pool. This setting is helpful for optimal memory utilization and for reducing the idle connections to the server just after the peak time. Zero implies no limit on reuse pool size.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for an HTTP profile. Must begin with a letter, number, or the underscore \\(_\\) character. Other characters allowed, after the first character, are the hyphen \\(-\\), period \\(.\\), hash \\(\\#\\), space \\( \\), at \\(@\\), colon \\(:\\), and equal \\(=\\) characters. The name of a HTTP profile cannot be changed after it is created.\n\nCLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks \\(for example, \"my http profile\" or 'my http profile'\\).",
			},
			"passprotocolupgrade": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Pass protocol upgrade request to the server.",
			},
			"persistentetag": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Generate the persistent Citrix ADC specific ETag for the HTTP response with ETag header.",
			},
			"reqtimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in seconds, within which the HTTP request must complete. If the request does not complete within this time, the specified request timeout action is executed. Zero disables the timeout.",
			},
			"reqtimeoutaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Action to take when the HTTP request does not complete within the specified request timeout duration. You can configure the following actions:\n* RESET - Send RST (reset) to client when timeout occurs.\n* DROP - Drop silently when timeout occurs.\n* Custom responder action - Name of the responder action to trigger when timeout occurs, used to send custom message.",
			},
			"reusepooltimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Idle timeout (in seconds) for server connections in re-use pool. Connections in the re-use pool are flushed, if they remain idle for the configured timeout.",
			},
			"rtsptunnel": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow RTSP tunnel in HTTP. Once application/x-rtsp-tunnelled is seen in Accept or Content-Type header, Citrix ADC does not process Layer 7 traffic on this connection.",
			},
			"weblog": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable web logging.",
			},
			"websocket": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "HTTP connection to be upgraded to a web socket connection. Once upgraded, Citrix ADC does not process Layer 7 traffic on this connection.",
			},
		},
	}
}
