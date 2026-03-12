package appfwprofile_cmdinjection_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AppfwprofileCmdinjectionBindingDataSourceSchema() schema.Schema {
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
			"as_scan_location_cmd": schema.StringAttribute{
				Optional:    true,
				Description: "Location of command injection exception - form field, header or cookie.",
			},
			"as_value_expr_cmd": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The web form/header/cookie value expression.",
			},
			"as_value_type_cmd": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of the relaxed web form value",
			},
			"cmdinjection": schema.StringAttribute{
				Required:    true,
				Description: "Name of the relaxed web form field/header/cookie",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments about the purpose of profile, or other useful information about the profile.",
			},
			"formactionurl_cmd": schema.StringAttribute{
				Required:    true,
				Description: "The web form action URL.",
			},
			"isautodeployed": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Is the rule auto deployed by dynamic profile ?",
			},
			"isregex_cmd": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Is the relaxed web form field name/header/cookie a regular expression?",
			},
			"isvalueregex_cmd": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Is the web form field/header/cookie value a regular expression?",
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
