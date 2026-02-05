package ipsecalgprofile

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func IpsecalgprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"connfailover": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Mode in which the connection failover feature must operate for the IPSec Alg. After a failover, established UDP connections and ESP packet flows are kept active and resumed on the secondary appliance. Recomended setting is ENABLED.",
			},
			"espgatetimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Timeout ESP in seconds as no ESP packets are seen after IKE negotiation",
			},
			"espsessiontimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "ESP session timeout in minutes.",
			},
			"ikesessiontimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "IKE session timeout in minutes",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the ipsec alg profile",
			},
		},
	}
}
