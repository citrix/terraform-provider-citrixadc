package appfwprofile_restvalidation_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AppfwprofileRestvalidationBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"alertonly": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Send SNMP alert?",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments about the purpose of profile, or other useful information about the profile.",
			},
			"isautodeployed": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Is the rule auto deployed by dynamic profile ?",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the profile to which to bind an exemption or rule.",
			},
			"resourceid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "A \"id\" that identifies the rule.",
			},
			"rest_validation_action": schema.StringAttribute{
				Required:    true,
				Description: "Action to be taken for traffic matching the configured relaxation rule.",
			},
			"restvalidation": schema.StringAttribute{
				Required:    true,
				Description: "Exempt REST endpoints or any other URLs matching the given pattern from the API schema validation check. Example: GET:/v1/bookstore/viewbooks",
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enabled.",
			},
		},
	}
}