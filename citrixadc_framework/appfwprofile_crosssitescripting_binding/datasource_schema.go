package appfwprofile_crosssitescripting_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AppfwprofileCrosssitescriptingBindingDataSourceSchema() schema.Schema {
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
			"as_scan_location_xss": schema.StringAttribute{
				Required:    true,
				Description: "Location of cross-site scripting exception - form field, header, cookie or URL.",
			},
			"as_value_expr_xss": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The web form value expression.",
			},
			"as_value_type_xss": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The web form value type.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments about the purpose of profile, or other useful information about the profile.",
			},
			"crosssitescripting": schema.StringAttribute{
				Required:    true,
				Description: "The web form field name.",
			},
			"formactionurl_xss": schema.StringAttribute{
				Required:    true,
				Description: "The web form action URL.",
			},
			"isautodeployed": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Is the rule auto deployed by dynamic profile ?",
			},
			"isregex_xss": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Is the web form field name a regular expression?",
			},
			"isvalueregex_xss": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Is the web form field value a regular expression?",
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
