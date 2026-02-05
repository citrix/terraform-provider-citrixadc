package aaaotpparameter

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/aaa"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// AaaotpparameterResourceModel describes the resource data model.
type AaaotpparameterResourceModel struct {
	Id            types.String `tfsdk:"id"`
	Encryption    types.String `tfsdk:"encryption"`
	Maxotpdevices types.Int64  `tfsdk:"maxotpdevices"`
}

func (r *AaaotpparameterResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the aaaotpparameter resource.",
			},
			"encryption": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "To encrypt otp secret in AD or not. Default value is OFF",
			},
			"maxotpdevices": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(4),
				Description: "Maximum number of otp devices user can register. Default value is 4. Max value is 255",
			},
		},
	}
}

func aaaotpparameterGetThePayloadFromtheConfig(ctx context.Context, data *AaaotpparameterResourceModel) aaa.Aaaotpparameter {
	tflog.Debug(ctx, "In aaaotpparameterGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	aaaotpparameter := aaa.Aaaotpparameter{}
	if !data.Encryption.IsNull() {
		aaaotpparameter.Encryption = data.Encryption.ValueString()
	}
	if !data.Maxotpdevices.IsNull() {
		aaaotpparameter.Maxotpdevices = utils.IntPtr(int(data.Maxotpdevices.ValueInt64()))
	}

	return aaaotpparameter
}

func aaaotpparameterSetAttrFromGet(ctx context.Context, data *AaaotpparameterResourceModel, getResponseData map[string]interface{}) *AaaotpparameterResourceModel {
	tflog.Debug(ctx, "In aaaotpparameterSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["encryption"]; ok && val != nil {
		data.Encryption = types.StringValue(val.(string))
	} else {
		data.Encryption = types.StringNull()
	}
	if val, ok := getResponseData["maxotpdevices"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxotpdevices = types.Int64Value(intVal)
		}
	} else {
		data.Maxotpdevices = types.Int64Null()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("aaaotpparameter-config")

	return data
}
