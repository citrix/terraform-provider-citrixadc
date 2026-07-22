package appfwprofile_jsonblockkeyword_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AppfwprofileJsonblockkeywordBindingDataSourceSchema() schema.Schema {
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
			"iskeyregex_json_blockkeyword": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Is JSON blockkeyword key a regular expression?",
			},
			"jsonblockkeyword": schema.StringAttribute{
				Required:    true,
				Description: "Field name of json block keyword binding",
			},
			"jsonblockkeywordtype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "JSON block keyword type",
			},
			"jsonblockkeywordurl": schema.StringAttribute{
				Required:    true,
				Description: "The json blockkeyword URL.",
			},
			"keyname_json_blockkeyword": schema.StringAttribute{
				Required:    true,
				Description: "JSON block keyword keyname",
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
