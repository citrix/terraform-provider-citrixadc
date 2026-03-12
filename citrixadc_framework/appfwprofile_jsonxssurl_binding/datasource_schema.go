package appfwprofile_jsonxssurl_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AppfwprofileJsonxssurlBindingDataSourceSchema() schema.Schema {
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
			"as_value_expr_json_xss": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The JSON XSS key value expression.",
			},
			"as_value_type_json_xss": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of the relaxed JSON XSS key value",
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
			"iskeyregex_json_xss": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Is the key name a regular expression?",
			},
			"isvalueregex_json_xss": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Is the JSON XSS key value a regular expression?",
			},
			"jsonxssurl": schema.StringAttribute{
				Required:    true,
				Description: "A regular expression that designates a URL on the Json XSS URL list for which XSS violations are relaxed.\nEnclose URLs in double quotes to ensure preservation of any embedded spaces or non-alphanumeric characters.",
			},
			"keyname_json_xss": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "An expression that designates a keyname on the JSON XSS URL for which XSS injection violations are relaxed.",
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
