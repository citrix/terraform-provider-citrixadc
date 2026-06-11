package fis_channel_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func FisChannelBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"ifnum": schema.StringAttribute{
				Required:    true,
				Description: "Interface to be bound to the FIS, specified in slot/port notation (for example, 1/3)",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the FIS to which you want to bind interfaces.",
			},
			"ownernode": schema.Int64Attribute{
				Optional:    true,
				Description: "ID of the cluster node for which you are creating the FIS. Can be configured only through the cluster IP address.",
			},
		},
	}
}
