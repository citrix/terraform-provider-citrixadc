package feoparameter

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/feo"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// FeoparameterResourceModel describes the resource data model.
type FeoparameterResourceModel struct {
	Id                 types.String `tfsdk:"id"`
	Cssinlinethressize types.Int64  `tfsdk:"cssinlinethressize"`
	Imginlinethressize types.Int64  `tfsdk:"imginlinethressize"`
	Jpegqualitypercent types.Int64  `tfsdk:"jpegqualitypercent"`
	Jsinlinethressize  types.Int64  `tfsdk:"jsinlinethressize"`
}

func (r *FeoparameterResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the feoparameter resource.",
			},
			// Optional+Computed (no Default) to match the SDK v2 contract: the
			// ADC always returns these values (NITRO defaults), so they are read
			// back and computed rather than defaulted in the schema. A schema
			// Default without Computed is invalid and panics the framework.
			"cssinlinethressize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Threshold value of the file size (in bytes) for converting external CSS files to inline CSS files.",
			},
			"imginlinethressize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum file size of an image (in bytes), for coverting linked images to inline images.",
			},
			"jpegqualitypercent": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The percentage value of a JPEG image quality to be reduced. Range: 0 - 100",
			},
			"jsinlinethressize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Threshold value of the file size (in bytes), for converting external JavaScript files to inline JavaScript files.",
			},
		},
	}
}

func feoparameterGetThePayloadFromthePlan(ctx context.Context, data *FeoparameterResourceModel) feo.Feoparameter {
	tflog.Debug(ctx, "In feoparameterGetThePayloadFromthePlan Function")

	// Create API request body from the model. Only send attributes that carry a
	// concrete, configured value. Unconfigured Optional+Computed attributes are
	// Unknown (not Null) in the plan, so guard on both to avoid sending a bogus
	// zero value to the ADC.
	feoparameter := feo.Feoparameter{}
	if !data.Cssinlinethressize.IsNull() && !data.Cssinlinethressize.IsUnknown() {
		feoparameter.Cssinlinethressize = utils.IntPtr(int(data.Cssinlinethressize.ValueInt64()))
	}
	if !data.Imginlinethressize.IsNull() && !data.Imginlinethressize.IsUnknown() {
		feoparameter.Imginlinethressize = utils.IntPtr(int(data.Imginlinethressize.ValueInt64()))
	}
	if !data.Jpegqualitypercent.IsNull() && !data.Jpegqualitypercent.IsUnknown() {
		feoparameter.Jpegqualitypercent = utils.IntPtr(int(data.Jpegqualitypercent.ValueInt64()))
	}
	if !data.Jsinlinethressize.IsNull() && !data.Jsinlinethressize.IsUnknown() {
		feoparameter.Jsinlinethressize = utils.IntPtr(int(data.Jsinlinethressize.ValueInt64()))
	}

	return feoparameter
}

func feoparameterSetAttrFromGet(ctx context.Context, data *FeoparameterResourceModel, getResponseData map[string]interface{}) *FeoparameterResourceModel {
	tflog.Debug(ctx, "In feoparameterSetAttrFromGet Function")

	// Convert API response to model. When a value is absent from the GET
	// response, only null it out if it was Unknown (never configured); otherwise
	// preserve the configured value to avoid "inconsistent result after apply".
	if val, ok := getResponseData["cssinlinethressize"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Cssinlinethressize = types.Int64Value(intVal)
		}
	} else if data.Cssinlinethressize.IsUnknown() {
		data.Cssinlinethressize = types.Int64Null()
	}
	if val, ok := getResponseData["imginlinethressize"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Imginlinethressize = types.Int64Value(intVal)
		}
	} else if data.Imginlinethressize.IsUnknown() {
		data.Imginlinethressize = types.Int64Null()
	}
	if val, ok := getResponseData["jpegqualitypercent"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Jpegqualitypercent = types.Int64Value(intVal)
		}
	} else if data.Jpegqualitypercent.IsUnknown() {
		data.Jpegqualitypercent = types.Int64Null()
	}
	if val, ok := getResponseData["jsinlinethressize"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Jsinlinethressize = types.Int64Value(intVal)
		}
	} else if data.Jsinlinethressize.IsUnknown() {
		data.Jsinlinethressize = types.Int64Null()
	}

	// Set ID for the resource.
	// Singleton (unnamed) resource - static ID matching SDK v2 behavior.
	data.Id = types.StringValue("feoparameter-config")

	return data
}
