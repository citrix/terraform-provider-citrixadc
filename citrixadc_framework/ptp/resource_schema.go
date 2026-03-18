package ptp

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// PtpResourceModel describes the resource data model.
type PtpResourceModel struct {
	Id    types.String `tfsdk:"id"`
	State types.String `tfsdk:"state"`
}

func (r *PtpResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the ptp resource.",
			},
			"state": schema.StringAttribute{
				Required:    true,
				Default:     stringdefault.StaticString("ENABLE"),
				Description: "Enables or disables Precision Time Protocol (PTP) on the appliance. If you disable PTP, make sure you enable Network Time Protocol (NTP) on the cluster.",
			},
		},
	}
}

func ptpGetThePayloadFromtheConfig(ctx context.Context, data *PtpResourceModel) network.Ptp {
	tflog.Debug(ctx, "In ptpGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	ptp := network.Ptp{}
	if !data.State.IsNull() {
		ptp.State = data.State.ValueString()
	}

	return ptp
}

func ptpSetAttrFromGet(ctx context.Context, data *PtpResourceModel, getResponseData map[string]interface{}) *PtpResourceModel {
	tflog.Debug(ctx, "In ptpSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	} else {
		data.State = types.StringNull()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("ptp-config")

	return data
}
