package nsdiameter

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func NsdiameterDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"identity": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "DiameterIdentity to be used by NS. DiameterIdentity is used to identify a Diameter node uniquely. Before setting up diameter configuration, Citrix ADC (as a Diameter node) MUST be assigned a unique DiameterIdentity.\nexample =>\nset ns diameter -identity netscaler.com\nNow whenever Citrix ADC needs to use identity in diameter messages. It will use 'netscaler.com' as Origin-Host AVP as defined in RFC3588",
			},
			"ownernode": schema.Int64Attribute{
				Required:    true,
				Description: "ID of the cluster node for which the diameter id is set, can be configured only through CLIP",
			},
			"realm": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Diameter Realm to be used by NS.\nexample =>\nset ns diameter -realm com\nNow whenever Citrix ADC system needs to use realm in diameter messages. It will use 'com' as Origin-Realm AVP as defined in RFC3588",
			},
			"serverclosepropagation": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "when a Server connection goes down, whether to close the corresponding client connection if there were requests pending on the server.",
			},
		},
	}
}
