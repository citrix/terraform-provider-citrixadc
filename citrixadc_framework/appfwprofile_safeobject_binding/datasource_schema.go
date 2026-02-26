package appfwprofile_safeobject_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func AppfwprofileSafeobjectBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"action": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Safe Object action types. (BLOCK | LOG | STATS | NONE)",
			},
			"alertonly": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Send SNMP alert?",
			},
			"as_expression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "A regular expression that defines the Safe Object.",
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
			"maxmatchlength": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum match length for a Safe Object expression.",
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
			"safeobject": schema.StringAttribute{
				Required:    true,
				Description: "Name of the Safe Object.",
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enabled.",
			},
		},
	}
}
