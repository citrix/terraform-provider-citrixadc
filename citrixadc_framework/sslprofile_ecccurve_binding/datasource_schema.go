package sslprofile_ecccurve_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SslprofileEcccurveBindingDataSourceModel describes the datasource data model.
//
// The datasource looks up a SINGLE ecccurve binding on a profile (name + a single
// ecccurvename), unlike the resource which manages a list of curve names. It
// therefore has its own model struct.
type SslprofileEcccurveBindingDataSourceModel struct {
	Id             types.String `tfsdk:"id"`
	Cipherpriority types.Int64  `tfsdk:"cipherpriority"`
	Ecccurvename   types.String `tfsdk:"ecccurvename"`
	Name           types.String `tfsdk:"name"`
}

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
