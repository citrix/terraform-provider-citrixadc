package rewritepolicylabel

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func RewritepolicylabelDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments to preserve information about this rewrite policy label.",
			},
			"labelname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the rewrite policy label. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the rewrite policy label is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my rewrite policy label\" or 'my rewrite policy label').",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New name for the rewrite policy label. \nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my policy label\" or 'my policy label').",
			},
			"transform": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Types of transformations allowed by the policies bound to the label. For Rewrite, the following types are supported:\n* http_req - HTTP requests\n* http_res - HTTP responses\n* othertcp_req - Non-HTTP TCP requests\n* othertcp_res - Non-HTTP TCP responses\n* url - URLs\n* text - Text strings\n* clientless_vpn_req - Citrix ADC clientless VPN requests\n* clientless_vpn_res - Citrix ADC clientless VPN responses\n* sipudp_req - SIP requests\n* sipudp_res - SIP responses\n* diameter_req - DIAMETER requests\n* diameter_res - DIAMETER responses\n* radius_req - RADIUS requests\n* radius_res - RADIUS responses\n* dns_req - DNS requests\n* dns_res - DNS responses\n* mqtt_req - MQTT requests\n* mqtt_res - MQTT responses",
			},
		},
	}
}
