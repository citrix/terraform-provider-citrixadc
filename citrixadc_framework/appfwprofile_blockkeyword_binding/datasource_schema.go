package appfwprofile_blockkeyword_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AppfwprofileBlockkeywordBindingDataSourceSchema() schema.Schema {
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
			"as_blockkeyword_formurl": schema.StringAttribute{
				Required:    true,
				Description: "The blockkeyword form action URL.",
			},
			"as_fieldname_isregex_blockkeyword": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Is block keyword field name regular expression?",
			},
			"blockkeyword": schema.StringAttribute{
				Required:    true,
				Description: "Field name of the block keyword binding.",
			},
			"blockkeywordtype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "block keyword type(literal|PCRE)",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments about the purpose of profile, or other useful information about the profile.",
			},
			"fieldname": schema.StringAttribute{
				Required:    true,
				Description: "A block keyword field name",
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
