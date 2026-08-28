package hasecureheartbeats

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ha"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// HasecureheartbeatsResourceModel describes the resource data model.
type HasecureheartbeatsResourceModel struct {
	Id             types.String `tfsdk:"id"`
	State          types.String `tfsdk:"state"`
	Hapsk          types.String `tfsdk:"hapsk"`
	HapskWo        types.String `tfsdk:"hapsk_wo"`
	HapskWoVersion types.Int64  `tfsdk:"hapsk_wo_version"`
}

func (r *HasecureheartbeatsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the hasecureheartbeats resource.",
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "By enabling this option, HA heartbeats are securely exchanged between nodes. Possible values = ENABLED, DISABLED",
			},
			"hapsk": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "Pre shared key to be used for securing HA heartbeats.",
			},
			"hapsk_wo": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				WriteOnly:   true,
				Description: "Pre shared key to be used for securing HA heartbeats.",
			},
			"hapsk_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(1),
				Description: "Increment this version to signal a hapsk_wo update.",
			},
		},
	}
}

func hasecureheartbeatsGetThePayloadFromthePlan(ctx context.Context, data *HasecureheartbeatsResourceModel) ha.Hasecureheartbeats {
	tflog.Debug(ctx, "In hasecureheartbeatsGetThePayloadFromthePlan Function")

	// Create API request body from the model
	hasecureheartbeats := ha.Hasecureheartbeats{}
	if !data.State.IsNull() && !data.State.IsUnknown() {
		hasecureheartbeats.State = data.State.ValueString()
	}
	if !data.Hapsk.IsNull() && !data.Hapsk.IsUnknown() {
		hasecureheartbeats.Hapsk = data.Hapsk.ValueString()
	}
	// Skip write-only attribute: hapsk_wo
	// Skip version tracker attribute: hapsk_wo_version

	return hasecureheartbeats
}

func hasecureheartbeatsGetThePayloadFromtheConfig(ctx context.Context, data *HasecureheartbeatsResourceModel, payload *ha.Hasecureheartbeats) {
	tflog.Debug(ctx, "In hasecureheartbeatsGetThePayloadFromtheConfig Function")

	// Add write-only attributes from config to the provided payload
	// Handle write-only secret attribute: hapsk_wo -> hapsk
	if !data.HapskWo.IsNull() {
		hapskWo := data.HapskWo.ValueString()
		if hapskWo != "" {
			payload.Hapsk = hapskWo
		}
	}
}

func hasecureheartbeatsSetAttrFromGet(ctx context.Context, data *HasecureheartbeatsResourceModel, getResponseData map[string]interface{}) *HasecureheartbeatsResourceModel {
	tflog.Debug(ctx, "In hasecureheartbeatsSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	} else {
		data.State = types.StringNull()
	}
	// hapsk is not returned by NITRO API in usable form (secret/ephemeral) - retain from config
	// hapsk_wo is not returned by NITRO API (secret/ephemeral) - retain from config
	// hapsk_wo_version is not returned by NITRO API (secret/ephemeral) - retain from config

	// Set ID for the resource
	// Case 1: No unique attributes - static ID (singleton)
	data.Id = types.StringValue("hasecureheartbeats-config")

	return data
}
