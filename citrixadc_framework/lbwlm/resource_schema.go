package lbwlm

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/lb"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// LbwlmResourceModel describes the resource data model.
type LbwlmResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Ipaddress types.String `tfsdk:"ipaddress"`
	Katimeout types.Int64  `tfsdk:"katimeout"`
	Lbuid     types.String `tfsdk:"lbuid"`
	Port      types.Int64  `tfsdk:"port"`
	Wlmname   types.String `tfsdk:"wlmname"`
}

func (r *LbwlmResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lbwlm resource.",
			},
			"ipaddress": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The IP address of the WLM.",
			},
			"katimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The idle time period after which Citrix ADC would probe the WLM. The value ranges from 1 to 1440 minutes.",
			},
			"lbuid": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The LBUID for the Load Balancer to communicate to the Work Load Manager.",
			},
			"port": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "The port of the WLM.",
			},
			"wlmname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The name of the Work Load Manager.",
			},
		},
	}
}

func lbwlmGetThePayloadFromthePlan(ctx context.Context, data *LbwlmResourceModel) lb.Lbwlm {
	tflog.Debug(ctx, "In lbwlmGetThePayloadFromthePlan Function")

	// Create API request body from the model
	lbwlm := lb.Lbwlm{}
	if !data.Ipaddress.IsNull() && !data.Ipaddress.IsUnknown() {
		lbwlm.Ipaddress = data.Ipaddress.ValueString()
	}
	if !data.Katimeout.IsNull() && !data.Katimeout.IsUnknown() {
		lbwlm.Katimeout = utils.IntPtr(int(data.Katimeout.ValueInt64()))
	}
	if !data.Lbuid.IsNull() && !data.Lbuid.IsUnknown() {
		lbwlm.Lbuid = data.Lbuid.ValueString()
	}
	if !data.Port.IsNull() && !data.Port.IsUnknown() {
		lbwlm.Port = utils.IntPtr(int(data.Port.ValueInt64()))
	}
	if !data.Wlmname.IsNull() && !data.Wlmname.IsUnknown() {
		lbwlm.Wlmname = data.Wlmname.ValueString()
	}

	return lbwlm
}

// lbwlmGetTheUpdatePayloadFromthePlan builds the payload for the NITRO PUT (update) call.
// Per the NITRO doc, the update payload accepts ONLY wlmname (key) + katimeout;
// ipaddress/port/lbuid are create-only and must not be sent on update.
func lbwlmGetTheUpdatePayloadFromthePlan(ctx context.Context, data *LbwlmResourceModel) lb.Lbwlm {
	tflog.Debug(ctx, "In lbwlmGetTheUpdatePayloadFromthePlan Function")

	lbwlm := lb.Lbwlm{}
	if !data.Wlmname.IsNull() && !data.Wlmname.IsUnknown() {
		lbwlm.Wlmname = data.Wlmname.ValueString()
	}
	if !data.Katimeout.IsNull() && !data.Katimeout.IsUnknown() {
		lbwlm.Katimeout = utils.IntPtr(int(data.Katimeout.ValueInt64()))
	}

	return lbwlm
}

func lbwlmSetAttrFromGet(ctx context.Context, data *LbwlmResourceModel, getResponseData map[string]interface{}) *LbwlmResourceModel {
	tflog.Debug(ctx, "In lbwlmSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["ipaddress"]; ok && val != nil {
		data.Ipaddress = types.StringValue(val.(string))
	} else {
		data.Ipaddress = types.StringNull()
	}
	if val, ok := getResponseData["katimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Katimeout = types.Int64Value(intVal)
		}
	} else {
		data.Katimeout = types.Int64Null()
	}
	if val, ok := getResponseData["lbuid"]; ok && val != nil {
		data.Lbuid = types.StringValue(val.(string))
	} else {
		data.Lbuid = types.StringNull()
	}
	if val, ok := getResponseData["port"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Port = types.Int64Value(intVal)
		}
	} else {
		data.Port = types.Int64Null()
	}
	if val, ok := getResponseData["wlmname"]; ok && val != nil {
		data.Wlmname = types.StringValue(val.(string))
	} else {
		data.Wlmname = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Wlmname.ValueString()))

	return data
}
