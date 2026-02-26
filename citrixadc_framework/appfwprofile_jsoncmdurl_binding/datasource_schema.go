package appfwprofile_jsoncmdurl_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AppfwprofileJsoncmdurlBindingDataSourceSchema() schema.Schema {
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
			"as_value_expr_json_cmd": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The JSON CMD key value expression.",
			},
			"as_value_type_json_cmd": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of the relaxed JSON CMD key value",
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
			"iskeyregex_json_cmd": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Is the key name a regular expression?",
			},
			"isvalueregex_json_cmd": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Is the JSON CMD key value a regular expression?",
			},
			"jsoncmdurl": schema.StringAttribute{
				Required:    true,
				Description: "A regular expression that designates a URL on the Json CMD URL list for which Command injection violations are relaxed.\nEnclose URLs in double quotes to ensure preservation of any embedded spaces or non-alphanumeric characters.",
			},
			"keyname_json_cmd": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "An expression that designates a keyname on the JSON CMD URL for which Command injection violations are relaxed.",
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
