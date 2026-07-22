package appfwprofile_bypasslist_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AppfwprofileBypasslistBindingDataSourceSchema() schema.Schema {
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
			"as_bypass_list": schema.StringAttribute{
				Required:    true,
				Description: "Bypass List Value",
			},
			"as_bypass_list_action": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Bypass List Action",
			},
			"as_bypass_list_location": schema.StringAttribute{
				Required:    true,
				Description: "Bypass List scan location",
			},
			"as_bypass_list_value_type": schema.StringAttribute{
				Required:    true,
				Description: "Bypass List value type",
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
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enabled.",
			},
		},
	}
}
