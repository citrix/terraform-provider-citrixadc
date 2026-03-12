package sslservice_ecccurve_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SslserviceEcccurveBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"ecccurvename": schema.StringAttribute{
				Required:    true,
				Description: "Named ECC curve bound to service/vserver.",
			},
			"servicename": schema.StringAttribute{
				Required:    true,
				Description: "Name of the SSL service for which to set advanced configuration.",
			},
		},
	}
}
