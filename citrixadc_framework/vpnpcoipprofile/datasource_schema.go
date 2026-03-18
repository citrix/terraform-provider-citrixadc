package vpnpcoipprofile

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func VpnpcoipprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"conserverurl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Connection server URL",
			},
			"icvverification": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "ICV verification for PCOIP transport packets.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "name of PCoIP profile",
			},
			"sessionidletimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "PCOIP Idle Session timeout",
			},
		},
	}
}
