package vpnpcoipvserverprofile

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func VpnpcoipvserverprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"logindomain": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Login domain for PCoIP users",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "name of PCoIP vserver profile",
			},
			"udpport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "UDP port for PCoIP data traffic",
			},
		},
	}
}
