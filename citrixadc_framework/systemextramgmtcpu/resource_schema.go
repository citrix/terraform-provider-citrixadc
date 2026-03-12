package systemextramgmtcpu

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/system"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SystemextramgmtcpuResourceModel describes the resource data model.
type SystemextramgmtcpuResourceModel struct {
	Id      types.String `tfsdk:"id"`
	Enabled types.Bool   `tfsdk:"enabled"`
}

func (r *SystemextramgmtcpuResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the systemextramgmtcpu resource.",
			},
			"enabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Boolean value indicating the effective state of the extra management CPU.",
			},
		},
	}
}

func systemextramgmtcpuGetThePayloadFromtheConfig(ctx context.Context, data *SystemextramgmtcpuResourceModel) system.Systemextramgmtcpu {
	tflog.Debug(ctx, "In systemextramgmtcpuGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	systemextramgmtcpu := system.Systemextramgmtcpu{}
	// if !data.Enabled.IsNull() {
	// 	systemextramgmtcpu.Enabled = utils.BoolPtr(data.Enabled.ValueBool())
	// }

	return systemextramgmtcpu
}

func systemextramgmtcpuSetAttrFromGet(ctx context.Context, data *SystemextramgmtcpuResourceModel, getResponseData map[string]interface{}) *SystemextramgmtcpuResourceModel {
	tflog.Debug(ctx, "In systemextramgmtcpuSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["effectivestate"]; ok && val != nil && val == "ENABLED" {
		data.Enabled = types.BoolValue(true)
	} else {
		data.Enabled = types.BoolValue(false)
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("systemextramgmtcpu-config")

	return data
}
