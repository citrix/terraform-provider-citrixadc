package nsaigwprofile

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func NsaigwprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the AIGW Profile.",
			},
			"endpointtype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The type of AI GW endpoint type. Possible values = azureopenai",
			},
			"profiletype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The binding entity for the aigw profile. Possible values = frontend, backend",
			},
			"tokenquota": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Token capacity of the backend server.",
			},
			"quotarefreshfrequency": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Quota refresh rate, in minutes.",
			},
			"authtoken": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
				Description: "Authentication token/API Key for the AI GW Endpoint.",
			},
			"authtoken_wo": schema.StringAttribute{
				Optional:    true,
				Description: "Authentication token/API Key for the AI GW Endpoint.",
			},
			"authtoken_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Increment this version to signal an authtoken_wo update.",
			},
		},
	}
}
