package radiusnode

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/basic"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// RadiusnodeResourceModel describes the resource data model.
type RadiusnodeResourceModel struct {
	Id              types.String `tfsdk:"id"`
	Nodeprefix      types.String `tfsdk:"nodeprefix"`
	Radkey          types.String `tfsdk:"radkey"`
	RadkeyWo        types.String `tfsdk:"radkey_wo"`
	RadkeyWoVersion types.Int64  `tfsdk:"radkey_wo_version"`
}

func (r *RadiusnodeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the radiusnode resource.",
			},
			"nodeprefix": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "IP address/IP prefix of radius node in CIDR format",
			},
			"radkey": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "The key shared between the RADIUS server and clients.\n      Required for NetScaler to communicate with the RADIUS nodes.",
			},
			"radkey_wo": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				WriteOnly:   true,
				Description: "The key shared between the RADIUS server and clients.\n      Required for NetScaler to communicate with the RADIUS nodes.",
			},
			"radkey_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(1),
				Description: "Increment this version to signal a radkey_wo update.",
			},
		},
	}
}

func radiusnodeGetThePayloadFromthePlan(ctx context.Context, data *RadiusnodeResourceModel) basic.Radiusnode {
	tflog.Debug(ctx, "In radiusnodeGetThePayloadFromthePlan Function")

	// Create API request body from the model
	radiusnode := basic.Radiusnode{}
	if !data.Nodeprefix.IsNull() && !data.Nodeprefix.IsUnknown() {
		radiusnode.Nodeprefix = data.Nodeprefix.ValueString()
	}
	if !data.Radkey.IsNull() && !data.Radkey.IsUnknown() {
		radiusnode.Radkey = data.Radkey.ValueString()
	}
	// Skip write-only attribute: radkey_wo
	// Skip version tracker attribute: radkey_wo_version

	return radiusnode
}

func radiusnodeGetThePayloadFromtheConfig(ctx context.Context, data *RadiusnodeResourceModel, payload *basic.Radiusnode) {
	tflog.Debug(ctx, "In radiusnodeGetThePayloadFromtheConfig Function")

	// Add write-only attributes from config to the provided payload
	// Handle write-only secret attribute: radkey_wo -> radkey
	if !data.RadkeyWo.IsNull() {
		radkeyWo := data.RadkeyWo.ValueString()
		if radkeyWo != "" {
			payload.Radkey = radkeyWo
		}
	}
}

func radiusnodeSetAttrFromGet(ctx context.Context, data *RadiusnodeResourceModel, getResponseData map[string]interface{}) *RadiusnodeResourceModel {
	tflog.Debug(ctx, "In radiusnodeSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["nodeprefix"]; ok && val != nil {
		data.Nodeprefix = types.StringValue(val.(string))
	} else {
		data.Nodeprefix = types.StringNull()
	}
	// radkey is not returned by NITRO API (secret/ephemeral) - retain from config
	// radkey_wo is not returned by NITRO API (secret/ephemeral) - retain from config
	// radkey_wo_version is not returned by NITRO API (secret/ephemeral) - retain from config

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Nodeprefix.ValueString()))

	return data
}
