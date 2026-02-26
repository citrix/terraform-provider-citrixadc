package sslvserver_ecccurve_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SslvserverEcccurveBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"ecccurvename": schema.StringAttribute{
				Required:    true,
				Description: "Named ECC curve bound to vserver/service.",
			},
			"vservername": schema.StringAttribute{
				Required:    true,
				Description: "Name of the SSL virtual server.",
			},
		},
	}
}
