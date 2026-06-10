package lsnsipalgcall

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/lsn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// LsnsipalgcallResourceModel describes the resource data model.
type LsnsipalgcallResourceModel struct {
	Id     types.String `tfsdk:"id"`
	Callid types.String `tfsdk:"callid"`
	Nodeid types.Int64  `tfsdk:"nodeid"`
}

func (r *LsnsipalgcallResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lsnsipalgcall resource.",
			},
			"callid": schema.StringAttribute{
				// NITRO flush marks callid mandatory; it is the only field in the
				// flush payload. Required (not Optional+Computed). Read is a no-op,
				// so it must not be Computed.
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Call ID for the SIP call.",
			},
			"nodeid": schema.Int64Attribute{
				// nodeid is a GET-only filter argument (it appears only in the get
				// args=callid:...,nodeid:... list, NOT in the flush payload).
				// Kept Optional, excluded from the flush payload (Pattern 15).
				// Read is a no-op, so it must not be Computed.
				Optional:    true,
				Description: "Unique number that identifies the cluster node.",
			},
		},
	}
}

func lsnsipalgcallGetThePayloadFromthePlan(ctx context.Context, data *LsnsipalgcallResourceModel) lsn.Lsnsipalgcall {
	tflog.Debug(ctx, "In lsnsipalgcallGetThePayloadFromthePlan Function")

	// Create API request body from the model.
	// Only callid is accepted by the flush action. nodeid is a GET-only filter
	// (Pattern 15) and is intentionally excluded from the flush payload.
	lsnsipalgcall := lsn.Lsnsipalgcall{}
	if !data.Callid.IsNull() && !data.Callid.IsUnknown() {
		lsnsipalgcall.Callid = data.Callid.ValueString()
	}

	return lsnsipalgcall
}

// lsnsipalgcallSetAttrFromGetForDatasource faithfully copies the GET response
// into the model for the datasource read path. It sets the synthetic ID since
// the datasource never calls Create.
func lsnsipalgcallSetAttrFromGetForDatasource(ctx context.Context, data *LsnsipalgcallResourceModel, getResponseData map[string]interface{}) *LsnsipalgcallResourceModel {
	tflog.Debug(ctx, "In lsnsipalgcallSetAttrFromGetForDatasource Function")

	// Convert API response to model
	if val, ok := getResponseData["callid"]; ok && val != nil {
		data.Callid = types.StringValue(val.(string))
	} else {
		data.Callid = types.StringNull()
	}
	if val, ok := getResponseData["nodeid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Nodeid = types.Int64Value(intVal)
		}
	} else {
		data.Nodeid = types.Int64Null()
	}

	// Set ID for the datasource (synthetic, from callid)
	data.Id = types.StringValue(data.Callid.ValueString())

	return data
}
