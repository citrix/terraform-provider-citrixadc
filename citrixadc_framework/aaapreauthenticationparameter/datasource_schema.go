package aaapreauthenticationparameter

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AaapreauthenticationparameterDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"deletefiles": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "String specifying the path(s) to and name(s) of the files to be deleted by the EPA tool, as a string of between 1 and 1023 characters.",
			},
			"killprocess": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "String specifying the name of a process to be terminated by the EPA tool.",
			},
			"preauthenticationaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Deny or allow login on the basis of end point analysis results.",
			},
			"rule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the Citrix ADC named rule, or an expression, to be evaluated by the EPA tool.",
			},
		},
	}
}
