package sslprofile_ecccurve_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SslprofileEcccurveBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"cipherpriority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Priority of the cipher binding",
			},
			"ecccurvename": schema.StringAttribute{
				Required:    true,
				Description: "Named ECC curve bound to vserver/service.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the SSL profile.",
			},
		},
	}
}
