package hasecureheartbeats

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func HasecureheartbeatsDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "By enabling this option, HA heartbeats are securely exchanged between nodes. Possible values = ENABLED, DISABLED",
			},
			"hapsk": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
				Description: "Pre shared key to be used for securing HA heartbeats.",
			},
		},
	}
}

type HasecureheartbeatsDataSourceModel struct {
	Id    types.String `tfsdk:"id"`
	State types.String `tfsdk:"state"`
	Hapsk types.String `tfsdk:"hapsk"`
}

func hasecureheartbeatsDataSourceSetAttrFromGet(ctx context.Context, data *HasecureheartbeatsDataSourceModel, getResponseData map[string]interface{}) *HasecureheartbeatsDataSourceModel {
	tflog.Debug(ctx, "In hasecureheartbeatsDataSourceSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	} else {
		data.State = types.StringNull()
	}
	// hapsk is not returned by NITRO API in usable form (secret/ephemeral) - retain from config

	// Set ID for the resource
	// Case 1: No unique attributes - static ID (singleton)
	data.Id = types.StringValue("hasecureheartbeats-config")

	return data
}
