package lbpersistentsessions

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

// LbpersistentsessionsResourceModel describes the resource data model.
// lbpersistentsessions is an action-only resource (clear via ?action=clear).
// vserver and persistenceparameter are optional clear filters. nodeid is a
// GET-only cluster filter (Pattern 15) and is therefore excluded from the clear
// payload (see the payload builder).
type LbpersistentsessionsResourceModel struct {
	Id                   types.String `tfsdk:"id"`
	Nodeid               types.Int64  `tfsdk:"nodeid"`
	Persistenceparameter types.String `tfsdk:"persistenceparameter"`
	Vserver              types.String `tfsdk:"vserver"`
}

func (r *LbpersistentsessionsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lbpersistentsessions resource.",
			},
			// All clear arguments are optional filters. Read is a no-op (clear is an
			// action, the cleared sessions are not a persistent managed object), so
			// these must NOT be Computed or Terraform reports an unknown value after
			// apply (Pattern 13 schema-flag implication).
			"nodeid": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Unique number that identifies the cluster node.",
			},
			"persistenceparameter": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The persistence parameter whose persistence sessions are to be flushed.",
			},
			"vserver": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The name of the virtual server.",
			},
		},
	}
}

// lbpersistentsessionsGetThePayloadFromthePlan builds the body for the clear
// action. Build a map containing ONLY the clear arguments that are set so the
// action body never includes invalid arguments. nodeid is a GET-only cluster
// filter and is intentionally excluded from the clear payload (Pattern 15).
func lbpersistentsessionsGetThePayloadFromthePlan(ctx context.Context, data *LbpersistentsessionsResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In lbpersistentsessionsGetThePayloadFromthePlan Function")

	lbpersistentsessions := map[string]interface{}{}
	if !data.Persistenceparameter.IsNull() && !data.Persistenceparameter.IsUnknown() {
		lbpersistentsessions["persistenceparameter"] = data.Persistenceparameter.ValueString()
	}
	if !data.Vserver.IsNull() && !data.Vserver.IsUnknown() {
		lbpersistentsessions["vserver"] = data.Vserver.ValueString()
	}

	return lbpersistentsessions
}

// lbpersistentsessionsSetAttrFromGetForDatasource faithfully copies the GET (get
// all) response into the model for the read-only datasource. It exposes the
// filter args present in tfdata (vserver, nodeid). The rich read-only output
// fields are not in tfdata and are not modelled. The resource itself never calls
// this (its Read is a no-op).
func lbpersistentsessionsSetAttrFromGetForDatasource(ctx context.Context, data *LbpersistentsessionsResourceModel, getResponseData map[string]interface{}) *LbpersistentsessionsResourceModel {
	tflog.Debug(ctx, "In lbpersistentsessionsSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["nodeid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Nodeid = types.Int64Value(intVal)
		}
	}
	if val, ok := getResponseData["persistenceparameter"]; ok && val != nil {
		data.Persistenceparameter = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["vserver"]; ok && val != nil {
		data.Vserver = types.StringValue(val.(string))
	}

	return data
}
