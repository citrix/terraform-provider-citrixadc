package bridgegroup

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func BridgegroupDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"dynamicrouting": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable dynamic routing for this bridgegroup.",
			},
			"bridgegroup_id": schema.Int64Attribute{
				Required:    true,
				Description: "An integer that uniquely identifies the bridge group.",
			},
			"ipv6dynamicrouting": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable all IPv6 dynamic routing protocols on all VLANs bound to this bridgegroup. Note: For the ENABLED setting to work, you must configure IPv6 dynamic routing protocols from the VTYSH command line.",
			},
		},
	}
}
