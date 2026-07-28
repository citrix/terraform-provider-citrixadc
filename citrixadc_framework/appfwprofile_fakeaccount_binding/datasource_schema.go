package appfwprofile_fakeaccount_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AppfwprofileFakeaccountBindingDataSourceSchema() schema.Schema {
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
			"fakeaccount": schema.StringAttribute{
				Required:    true,
				Description: "Field name of the fake account rule.",
			},
			"formexpression": schema.StringAttribute{
				// At-most-one arm (mutually exclusive with formurl_fad). A binding
				// only ever has ONE arm populated, so this is an Optional lookup
				// filter, not a mandatory key. Mirrors the resource's at-most-one
				// composite-ID model.
				Optional:    true,
				Computed:    true,
				Description: "A regular expression that defines the Fake Account.",
			},
			"formurl_fad": schema.StringAttribute{
				// At-most-one arm (mutually exclusive with formexpression). Optional
				// lookup filter, not a mandatory key.
				Optional:    true,
				Computed:    true,
				Description: "The fake account detection URL.",
			},
			"isautodeployed": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Is the rule auto deployed by dynamic profile ?",
			},
			"isfieldnameregex": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Is Fake Account Detection field name regex?",
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
			"tag": schema.StringAttribute{
				Required:    true,
				Description: "A tag expression that defines the Fake Account.",
			},
		},
	}
}
