package quicparam

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/quic"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// QuicparamResourceModel describes the resource data model.
type QuicparamResourceModel struct {
	Id                types.String `tfsdk:"id"`
	Quicsecrettimeout types.Int64  `tfsdk:"quicsecrettimeout"`
}

func (r *QuicparamResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the quicparam resource.",
			},
			"quicsecrettimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(3600),
				Description: "Rotation frequency, in seconds, for the secret used to generate address validation tokens that will be issued in QUIC Retry packets and QUIC NEW_TOKEN frames sent by the Citrix ADC. A value of 0 can be configured if secret rotation is not desired.",
			},
		},
	}
}

func quicparamGetThePayloadFromthePlan(ctx context.Context, data *QuicparamResourceModel) quic.Quicparam {
	tflog.Debug(ctx, "In quicparamGetThePayloadFromthePlan Function")

	// Create API request body from the model
	quicparam := quic.Quicparam{}
	if !data.Quicsecrettimeout.IsNull() && !data.Quicsecrettimeout.IsUnknown() {
		quicparam.Quicsecrettimeout = utils.IntPtr(int(data.Quicsecrettimeout.ValueInt64()))
	}

	return quicparam
}

func quicparamSetAttrFromGet(ctx context.Context, data *QuicparamResourceModel, getResponseData map[string]interface{}) *QuicparamResourceModel {
	tflog.Debug(ctx, "In quicparamSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["quicsecrettimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Quicsecrettimeout = types.Int64Value(intVal)
		}
	} else {
		data.Quicsecrettimeout = types.Int64Null()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("quicparam-config")

	return data
}
