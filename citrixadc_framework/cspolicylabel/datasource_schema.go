package cspolicylabel

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func CspolicylabelDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"cspolicylabeltype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Protocol supported by the policy label. All policies bound to the policy label must either match the specified protocol or be a subtype of that protocol. Available settings function as follows:\n* HTTP - Supports policies that process HTTP traffic. Used to access unencrypted Web sites. (The default.)\n* SSL - Supports policies that process HTTPS/SSL encrypted traffic. Used to access encrypted Web sites.\n* TCP - Supports policies that process any type of TCP traffic, including HTTP.\n* SSL_TCP - Supports policies that process SSL-encrypted TCP traffic, including SSL.\n* UDP - Supports policies that process any type of UDP-based traffic, including DNS.\n* DNS - Supports policies that process DNS traffic.\n* ANY - Supports all types of policies except HTTP, SSL, and TCP.\n* SIP_UDP - Supports policies that process UDP based Session Initiation Protocol (SIP) traffic. SIP initiates, manages, and terminates multimedia communications sessions, and has emerged as the standard for Internet telephony (VoIP).\n* RTSP - Supports policies that process Real Time Streaming Protocol (RTSP) traffic. RTSP provides delivery of multimedia and other streaming data, such as audio, video, and other types of streamed media.\n* RADIUS - Supports policies that process Remote Authentication Dial In User Service (RADIUS) traffic. RADIUS supports combined authentication, authorization, and auditing services for network management.\n* MYSQL - Supports policies that process MYSQL traffic.\n* MSSQL - Supports policies that process Microsoft SQL traffic.",
			},
			"labelname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the policy label. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters.\nThe label name must be unique within the list of policy labels for content switching.\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my policylabel\" or 'my policylabel').",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The new name of the content switching policylabel.",
			},
		},
	}
}
