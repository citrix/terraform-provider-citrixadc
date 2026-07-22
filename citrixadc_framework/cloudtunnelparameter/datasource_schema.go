package cloudtunnelparameter

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func CloudtunnelparameterDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"controllerfqdn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "0",
			},
			"fqdn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "0",
			},
			"resourcelocation": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "0",
			},
			"subnetresourcelocationmappings": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "0",
			},
		},
	}
}
