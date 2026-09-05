package lbroute

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

// LbrouteResourceModel describes the resource data model.
type LbrouteResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Gatewayname types.String `tfsdk:"gatewayname"`
	Netmask     types.String `tfsdk:"netmask"`
	Network     types.String `tfsdk:"network"`
	Td          types.Int64  `tfsdk:"td"`
}

func (r *LbrouteResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lbroute resource.",
			},
			"gatewayname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The name of the route.",
			},
			"netmask": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The netmask to which the route belongs.",
			},
			"network": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The IP address of the network to which the route belongs.",
			},
			"td": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					// td was Optional + ForceNew in SDK v2. It is now
					// Optional+Computed (the ADC returns 0 when unset), so keep
					// the computed value stable and only force replacement when
					// the user actually configures and changes it.
					int64planmodifier.UseStateForUnknown(),
					int64planmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
		},
	}
}

func lbrouteGetThePayloadFromtheConfig(ctx context.Context, data *LbrouteResourceModel) lb.Lbroute {
	tflog.Debug(ctx, "In lbrouteGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	lbroute := lb.Lbroute{}
	if !data.Gatewayname.IsNull() && !data.Gatewayname.IsUnknown() {
		lbroute.Gatewayname = data.Gatewayname.ValueString()
	}
	if !data.Netmask.IsNull() && !data.Netmask.IsUnknown() {
		lbroute.Netmask = data.Netmask.ValueString()
	}
	if !data.Network.IsNull() && !data.Network.IsUnknown() {
		lbroute.Network = data.Network.ValueString()
	}
	// td is only sent when explicitly configured (matches SDK v2, which set td
	// only when GetRawConfig had a non-null value). When Optional+Computed and
	// unset, td is Unknown here, so guarding on IsUnknown avoids sending 0.
	if !data.Td.IsNull() && !data.Td.IsUnknown() {
		lbroute.Td = utils.IntPtr(int(data.Td.ValueInt64()))
	}

	return lbroute
}

func lbrouteSetAttrFromGet(ctx context.Context, data *LbrouteResourceModel, getResponseData map[string]interface{}) *LbrouteResourceModel {
	tflog.Debug(ctx, "In lbrouteSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["gatewayname"]; ok && val != nil {
		data.Gatewayname = types.StringValue(val.(string))
	} else {
		data.Gatewayname = types.StringNull()
	}
	if val, ok := getResponseData["netmask"]; ok && val != nil {
		data.Netmask = types.StringValue(val.(string))
	} else {
		data.Netmask = types.StringNull()
	}
	if val, ok := getResponseData["network"]; ok && val != nil {
		data.Network = types.StringValue(val.(string))
	} else {
		data.Network = types.StringNull()
	}
	if val, ok := getResponseData["td"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Td = types.Int64Value(intVal)
		}
	} else {
		// Omit-on-default guard: NITRO may omit td when it is 0. Only null it
		// when the value is unknown; never clobber a known configured value.
		if data.Td.IsUnknown() {
			data.Td = types.Int64Null()
		}
	}

	// Set ID for the resource.
	// SDK v2 used "network,netmask,gatewayname" (see resource_id_mapping.json).
	data.Id = types.StringValue(fmt.Sprintf("%s,%s,%s", data.Network.ValueString(), data.Netmask.ValueString(), data.Gatewayname.ValueString()))

	return data
}
