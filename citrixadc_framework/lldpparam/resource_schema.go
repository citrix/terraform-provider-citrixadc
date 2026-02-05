package lldpparam

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/lldp"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// LldpparamResourceModel describes the resource data model.
type LldpparamResourceModel struct {
	Id             types.String `tfsdk:"id"`
	Holdtimetxmult types.Int64  `tfsdk:"holdtimetxmult"`
	Mode           types.String `tfsdk:"mode"`
	Timer          types.Int64  `tfsdk:"timer"`
}

func (r *LldpparamResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lldpparam resource.",
			},
			"holdtimetxmult": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(4),
				Description: "A multiplier for calculating the duration for which the receiving device stores the LLDP information in its database before discarding or removing it. The duration is calculated as the holdtimeTxMult (Holdtime Multiplier) parameter value multiplied by the timer (Timer) parameter value.",
			},
			"mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Global mode of Link Layer Discovery Protocol (LLDP) on the Citrix ADC. The resultant LLDP mode of an interface depends on the LLDP mode configured at the global and the interface levels.",
			},
			"timer": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(30),
				Description: "Interval, in seconds, between LLDP packet data units (LLDPDUs).  that the Citrix ADC sends to a directly connected device.",
			},
		},
	}
}

func lldpparamGetThePayloadFromtheConfig(ctx context.Context, data *LldpparamResourceModel) lldp.Lldpparam {
	tflog.Debug(ctx, "In lldpparamGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	lldpparam := lldp.Lldpparam{}
	if !data.Holdtimetxmult.IsNull() {
		lldpparam.Holdtimetxmult = utils.IntPtr(int(data.Holdtimetxmult.ValueInt64()))
	}
	if !data.Mode.IsNull() {
		lldpparam.Mode = data.Mode.ValueString()
	}
	if !data.Timer.IsNull() {
		lldpparam.Timer = utils.IntPtr(int(data.Timer.ValueInt64()))
	}

	return lldpparam
}

func lldpparamSetAttrFromGet(ctx context.Context, data *LldpparamResourceModel, getResponseData map[string]interface{}) *LldpparamResourceModel {
	tflog.Debug(ctx, "In lldpparamSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["holdtimetxmult"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Holdtimetxmult = types.Int64Value(intVal)
		}
	} else {
		data.Holdtimetxmult = types.Int64Null()
	}
	if val, ok := getResponseData["mode"]; ok && val != nil {
		data.Mode = types.StringValue(val.(string))
	} else {
		data.Mode = types.StringNull()
	}
	if val, ok := getResponseData["timer"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Timer = types.Int64Value(intVal)
		}
	} else {
		data.Timer = types.Int64Null()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("lldpparam-config")

	return data
}
