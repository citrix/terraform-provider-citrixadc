package nsicapprofile

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func NsicapprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"allow204": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or Disable sending Allow: 204 header in ICAP request.",
			},
			"connectionkeepalive": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If enabled, Citrix ADC keeps the ICAP connection alive after a transaction to reuse it to send next ICAP request.",
			},
			"hostheader": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "ICAP Host Header",
			},
			"inserthttprequest": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Exact HTTP request, in the form of an expression, which the Citrix ADC encapsulates and sends to the ICAP server. If you set this parameter, the ICAP request is sent using only this header. This can be used when the HTTP header is not available to send or ICAP server only needs part of the incoming HTTP request. The request expression is constrained by the feature for which it is used.\nThe Citrix ADC does not check the validity of this request. You must manually validate the request.",
			},
			"inserticapheaders": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Insert custom ICAP headers in the ICAP request to send to ICAP server. The headers can be static or can be dynamically constructed using PI Policy Expression. For example, to send static user agent and Client's IP address, the expression can be specified as \"User-Agent: NS-ICAP-Client/V1.0\\r\\nX-Client-IP: \"+CLIENT.IP.SRC+\"\\r\\n\".\nThe Citrix ADC does not check the validity of the specified header name-value. You must manually validate the specified header syntax.",
			},
			"logaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the audit message action which would be evaluated on receiving the ICAP response to emit the logs.",
			},
			"mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "ICAP Mode of operation. It is a mandatory argument while creating an icapprofile.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for an ICAP profile. Must begin with a letter, number, or the underscore \\(_\\) character. Other characters allowed, after the first character, are the hyphen \\(-\\), period \\(.\\), hash \\(\\#\\), space \\( \\), at \\(@\\), colon \\(:\\), and equal \\(=\\) characters. The name of a ICAP profile cannot be changed after it is created.\n\nCLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks \\(for example, \"my icap profile\" or 'my icap profile'\\).",
			},
			"preview": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or Disable preview header with ICAP request. This feature allows an ICAP server to see the beginning of a transaction, then decide if it wants to opt-out of the transaction early instead of receiving the remainder of the request message.",
			},
			"previewlength": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Value of Preview Header field. Citrix ADC uses the minimum of this set value and the preview size received on OPTIONS response.",
			},
			"queryparams": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Query parameters to be included with ICAP request URI. Entered values should be in arg=value format. For more than one parameters, add & separated values. e.g.: arg1=val1&arg2=val2.",
			},
			"reqtimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in seconds, within which the remote server should respond to the ICAP-request. If the Netscaler does not receive full response with this time, the specified request timeout action is performed. Zero value disables this timeout functionality.",
			},
			"reqtimeoutaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the action to perform if the Vserver/Server representing the remote service does not respond with any response within the timeout value configured. The Supported actions are\n* BYPASS - This Ignores the remote server response and sends the request/response to Client/Server.\n           * If the ICAP response with Encapsulated headers is not received within the request-timeout value configured, this Ignores the remote ICAP server response and sends the Full request/response to Server/Client.\n* RESET - Reset the client connection by closing it. The client program, such as a browser, will handle this and may inform the user. The client may then resend the request if desired.\n* DROP - Drop the request without sending a response to the user.",
			},
			"uri": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URI representing icap service. It is a mandatory argument while creating an icapprofile.",
			},
			"useragent": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "ICAP User Agent Header String",
			},
		},
	}
}
