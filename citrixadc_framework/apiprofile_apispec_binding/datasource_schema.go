package apiprofile_apispec_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func ApiprofileApispecBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"apispec": schema.StringAttribute{
				Required:    true,
				Description: "Name for the API spec which will be binded to the profile.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the API profile in which to bind the API apispec(s).",
			},
		},
	}
}