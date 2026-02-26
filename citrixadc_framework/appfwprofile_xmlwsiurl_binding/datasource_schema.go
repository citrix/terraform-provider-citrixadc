package appfwprofile_xmlwsiurl_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AppfwprofileXmlwsiurlBindingDataSourceSchema() schema.Schema {
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
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enabled.",
			},
			"xmlwsichecks": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify a comma separated list of relevant WS-I rule IDs. (R1140, R1141)",
			},
			"xmlwsiurl": schema.StringAttribute{
				Required:    true,
				Description: "XML WS-I URL regular expression length.",
			},
		},
	}
}
