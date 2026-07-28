package lbpolicylabel

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func LbpolicylabelDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments to preserve information about this LB policy label.",
			},
			"labelname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the LB policy label. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my lb policy label\" or 'my lb policy label').",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New name for the LB policy label. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.",
			},
			"policylabeltype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Protocols supported by the policylabel. Available Types are :\n* HTTP - HTTP requests.\n* DNS - DNS request.\n* OTHERTCP - OTHERTCP request.\n* SIP_UDP - SIP_UDP request.\n* SIP_TCP - SIP_TCP request.\n* MYSQL - MYSQL request.\n* MSSQL - MSSQL request.\n* ORACLE - ORACLE request.\n* NAT - NAT request.\n* DIAMETER - DIAMETER request.\n* RADIUS - RADIUS request.\n* MQTT - MQTT request.\n* QUIC_BRIDGE - QUIC_BRIDGE request.\n* HTTP_QUIC - HTTP_QUIC request.",
			},
		},
	}
}
