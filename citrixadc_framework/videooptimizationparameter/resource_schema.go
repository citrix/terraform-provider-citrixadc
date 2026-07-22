package videooptimizationparameter

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/videooptimization"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// VideooptimizationparameterResourceModel describes the resource data model.
type VideooptimizationparameterResourceModel struct {
	Id                       types.String  `tfsdk:"id"`
	Quicpacingrate           types.Int64   `tfsdk:"quicpacingrate"`
	Randomsamplingpercentage types.Float64 `tfsdk:"randomsamplingpercentage"`
}

func (r *VideooptimizationparameterResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the videooptimizationparameter resource.",
			},
			"quicpacingrate": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "QUIC Video Pacing Rate (Kbps).",
			},
			"randomsamplingpercentage": schema.Float64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Random Sampling Percentage.",
			},
		},
	}
}

func videooptimizationparameterGetThePayloadFromthePlan(ctx context.Context, data *VideooptimizationparameterResourceModel) videooptimization.Videooptimizationparameter {
	tflog.Debug(ctx, "In videooptimizationparameterGetThePayloadFromthePlan Function")

	// Create API request body from the model
	videooptimizationparameter := videooptimization.Videooptimizationparameter{}
	if !data.Quicpacingrate.IsNull() && !data.Quicpacingrate.IsUnknown() {
		videooptimizationparameter.Quicpacingrate = utils.IntPtr(int(data.Quicpacingrate.ValueInt64()))
	}
	if !data.Randomsamplingpercentage.IsNull() && !data.Randomsamplingpercentage.IsUnknown() {
		videooptimizationparameter.Randomsamplingpercentage = data.Randomsamplingpercentage.ValueFloat64()
	}

	return videooptimizationparameter
}

func videooptimizationparameterSetAttrFromGet(ctx context.Context, data *VideooptimizationparameterResourceModel, getResponseData map[string]interface{}) *VideooptimizationparameterResourceModel {
	tflog.Debug(ctx, "In videooptimizationparameterSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["quicpacingrate"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Quicpacingrate = types.Int64Value(intVal)
		}
	} else {
		data.Quicpacingrate = types.Int64Null()
	}
	if val, ok := getResponseData["randomsamplingpercentage"]; ok && val != nil {
		data.Randomsamplingpercentage = types.Float64Value(val.(float64))
	} else {
		data.Randomsamplingpercentage = types.Float64Null()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("videooptimizationparameter-config")

	return data
}
