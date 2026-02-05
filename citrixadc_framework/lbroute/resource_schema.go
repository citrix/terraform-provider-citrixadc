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
					int64planmodifier.RequiresReplace(),
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
	if !data.Gatewayname.IsNull() {
		lbroute.Gatewayname = data.Gatewayname.ValueString()
	}
	if !data.Netmask.IsNull() {
		lbroute.Netmask = data.Netmask.ValueString()
	}
	if !data.Network.IsNull() {
		lbroute.Network = data.Network.ValueString()
	}
	if !data.Td.IsNull() {
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
		data.Td = types.Int64Null()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated
	data.Id = types.StringValue(fmt.Sprintf("%s,%s,%d", data.Network.ValueString(), data.Netmask.ValueString(), data.Td.ValueInt64()))

	return data
}
