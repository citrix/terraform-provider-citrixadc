package radiusnode

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func RadiusnodeDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"nodeprefix": schema.StringAttribute{
				Required:    true,
				Description: "IP address/IP prefix of radius node in CIDR format",
			},
			"radkey": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The key shared between the RADIUS server and clients.\n      Required for NetScaler to communicate with the RADIUS nodes.",
			},
		},
	}
}
