package nsspparams

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// NsspparamsResourceModel describes the resource data model.
type NsspparamsResourceModel struct {
	Id            types.String `tfsdk:"id"`
	Basethreshold types.Int64  `tfsdk:"basethreshold"`
	Throttle      types.String `tfsdk:"throttle"`
}

func (r *NsspparamsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nsspparams resource.",
			},
			"basethreshold": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(200),
				Description: "Maximum number of server connections that can be opened before surge protection is activated.",
			},
			"throttle": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("Normal"),
				Description: "Rate at which the system opens connections to the server.",
			},
		},
	}
}

func nsspparamsGetThePayloadFromtheConfig(ctx context.Context, data *NsspparamsResourceModel) ns.Nsspparams {
	tflog.Debug(ctx, "In nsspparamsGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	nsspparams := ns.Nsspparams{}
	if !data.Basethreshold.IsNull() {
		nsspparams.Basethreshold = utils.IntPtr(int(data.Basethreshold.ValueInt64()))
	}
	if !data.Throttle.IsNull() {
		nsspparams.Throttle = data.Throttle.ValueString()
	}

	return nsspparams
}

func nsspparamsSetAttrFromGet(ctx context.Context, data *NsspparamsResourceModel, getResponseData map[string]interface{}) *NsspparamsResourceModel {
	tflog.Debug(ctx, "In nsspparamsSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["basethreshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Basethreshold = types.Int64Value(intVal)
		}
	} else {
		data.Basethreshold = types.Int64Null()
	}
	if val, ok := getResponseData["throttle"]; ok && val != nil {
		data.Throttle = types.StringValue(val.(string))
	} else {
		data.Throttle = types.StringNull()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("nsspparams-config")

	return data
}
