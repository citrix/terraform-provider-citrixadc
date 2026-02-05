package lbroute6

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

// Lbroute6ResourceModel describes the resource data model.
type Lbroute6ResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Gatewayname types.String `tfsdk:"gatewayname"`
	Network     types.String `tfsdk:"network"`
	Td          types.Int64  `tfsdk:"td"`
}

func (r *Lbroute6Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lbroute6 resource.",
			},
			"gatewayname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The name of the route.",
			},
			"network": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The destination network.",
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

func lbroute6GetThePayloadFromtheConfig(ctx context.Context, data *Lbroute6ResourceModel) lb.Lbroute6 {
	tflog.Debug(ctx, "In lbroute6GetThePayloadFromtheConfig Function")

	// Create API request body from the model
	lbroute6 := lb.Lbroute6{}
	if !data.Gatewayname.IsNull() {
		lbroute6.Gatewayname = data.Gatewayname.ValueString()
	}
	if !data.Network.IsNull() {
		lbroute6.Network = data.Network.ValueString()
	}
	if !data.Td.IsNull() {
		lbroute6.Td = utils.IntPtr(int(data.Td.ValueInt64()))
	}

	return lbroute6
}

func lbroute6SetAttrFromGet(ctx context.Context, data *Lbroute6ResourceModel, getResponseData map[string]interface{}) *Lbroute6ResourceModel {
	tflog.Debug(ctx, "In lbroute6SetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["gatewayname"]; ok && val != nil {
		data.Gatewayname = types.StringValue(val.(string))
	} else {
		data.Gatewayname = types.StringNull()
	}
	if val, ok := getResponseData["network"]; ok && val != nil {
		data.Network = types.StringValue(val.(string))
	} else {
		data.Network = types.StringNull()
	}
	if val, ok := getResponseData["td"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Td = types.Int64Value(intVal)
		} else {
			data.Td = types.Int64Value(0)
		}
	} else {
		// If td is not present in response, default to 0
		data.Td = types.Int64Value(0)
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated
	data.Id = types.StringValue(fmt.Sprintf("%s,%d", data.Network.ValueString(), data.Td.ValueInt64()))

	return data
}
