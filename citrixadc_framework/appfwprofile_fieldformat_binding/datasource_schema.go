package appfwprofile_fieldformat_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AppfwprofileFieldformatBindingDataSourceSchema() schema.Schema {
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
			"fieldformat": schema.StringAttribute{
				Required:    true,
				Description: "Name of the form field to which a field format will be assigned.",
			},
			"fieldformatmaxlength": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The maximum allowed length for data in this form field.",
			},
			"fieldformatminlength": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The minimum allowed length for data in this form field.",
			},
			"fieldtype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The field type you are assigning to this form field.",
			},
			"formactionurl_ff": schema.StringAttribute{
				Required:    true,
				Description: "Action URL of the form field to which a field format will be assigned.",
			},
			"isautodeployed": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Is the rule auto deployed by dynamic profile ?",
			},
			"isregexff": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Is the form field name a regular expression?",
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
