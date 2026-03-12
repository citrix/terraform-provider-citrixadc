package nshostname

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func NshostnameDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"hostname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Host name for the Citrix ADC.",
			},
			"ownernode": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "ID of the cluster node for which you are setting the hostname. Can be configured only through the cluster IP address.",
			},
		},
	}
}
