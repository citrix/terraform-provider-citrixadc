package lsnrtspalgsession

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// LsnrtspalgsessionResourceModel describes the resource data model.
// lsnrtspalgsession is an action-only resource (flush via ?action=flush).
// nodeid is a GET-only cluster filter (Pattern 15) and is therefore excluded
// from the flush payload (see the payload builder).
type LsnrtspalgsessionResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Nodeid    types.Int64  `tfsdk:"nodeid"`
	Sessionid types.String `tfsdk:"sessionid"`
}

func (r *LsnrtspalgsessionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lsnrtspalgsession resource.",
			},
			// nodeid is a GET-only cluster filter, not a flush payload argument
			// (Pattern 15). Read is a no-op (flush is an action on a transient
			// session), so it must NOT be Computed or Terraform reports an unknown
			// value after apply (Pattern 13 schema-flag implication).
			"nodeid": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Unique number that identifies the cluster node.",
			},
			// sessionid is mandatory for the flush action (NITRO marks it red/bold).
			// Required, not Optional+Computed.
			"sessionid": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Session ID for the RTSP call.",
			},
		},
	}
}

// lsnrtspalgsessionGetThePayloadFromthePlan builds the body for the flush action.
// Only sessionid is a valid flush argument. nodeid is a GET-only cluster filter
// and is intentionally excluded from the flush payload (Pattern 15).
func lsnrtspalgsessionGetThePayloadFromthePlan(ctx context.Context, data *LsnrtspalgsessionResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In lsnrtspalgsessionGetThePayloadFromthePlan Function")

	lsnrtspalgsession := map[string]interface{}{}
	if !data.Sessionid.IsNull() && !data.Sessionid.IsUnknown() {
		lsnrtspalgsession["sessionid"] = data.Sessionid.ValueString()
	}

	return lsnrtspalgsession
}

// lsnrtspalgsessionSetAttrFromGetForDatasource faithfully copies the GET response
// into the model for the read-only datasource. The resource itself never calls
// this (its Read is a no-op).
func lsnrtspalgsessionSetAttrFromGetForDatasource(ctx context.Context, data *LsnrtspalgsessionResourceModel, getResponseData map[string]interface{}) *LsnrtspalgsessionResourceModel {
	tflog.Debug(ctx, "In lsnrtspalgsessionSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["nodeid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Nodeid = types.Int64Value(intVal)
		}
	}
	if val, ok := getResponseData["sessionid"]; ok && val != nil {
		data.Sessionid = types.StringValue(val.(string))
	}

	return data
}
