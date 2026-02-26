package sslservicegroup_ecccurve_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SslservicegroupEcccurveBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"ecccurvename": schema.StringAttribute{
				Required:    true,
				Description: "Named ECC curve bound to servicegroup.",
			},
			"servicegroupname": schema.StringAttribute{
				Required:    true,
				Description: "The name of the SSL service to which the SSL policy needs to be bound.",
			},
		},
	}
}
