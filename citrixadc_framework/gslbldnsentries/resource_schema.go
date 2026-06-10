package gslbldnsentries

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// GslbldnsentriesResourceModel describes the resource data model.
// gslbldnsentries is an action-only resource (clear via ?action=clear). nodeid is
// a GET-only cluster filter (Pattern 15) and is therefore excluded from the clear
// payload (see the payload builder); clear itself takes no arguments.
type GslbldnsentriesResourceModel struct {
	Id     types.String `tfsdk:"id"`
	Nodeid types.Int64  `tfsdk:"nodeid"`
}

func (r *GslbldnsentriesResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the gslbldnsentries resource.",
			},
			// nodeid is a GET-only cluster filter. Read is a no-op (clear is an
			// action, the cleared LDNS entries are not a persistent managed object),
			// so nodeid must NOT be Computed or Terraform reports an unknown value
			// after apply (Pattern 13 schema-flag implication).
			"nodeid": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Unique number that identifies the cluster node.",
			},
		},
	}
}

// gslbldnsentriesGetThePayloadFromthePlan builds the body for the clear action.
// clear takes no arguments, so the body is always empty. nodeid is a GET-only
// cluster filter and is intentionally excluded from the clear payload (Pattern
// 15).
func gslbldnsentriesGetThePayloadFromthePlan(ctx context.Context, data *GslbldnsentriesResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In gslbldnsentriesGetThePayloadFromthePlan Function")

	gslbldnsentries := map[string]interface{}{}

	return gslbldnsentries
}

// gslbldnsentriesSetAttrFromGetForDatasource faithfully copies the GET (get all)
// response into the model for the read-only datasource. It exposes the filter arg
// present in tfdata (nodeid). The rich read-only output fields are not in tfdata
// and are not modelled. The resource itself never calls this (its Read is a
// no-op).
func gslbldnsentriesSetAttrFromGetForDatasource(ctx context.Context, data *GslbldnsentriesResourceModel, getResponseData map[string]interface{}) *GslbldnsentriesResourceModel {
	tflog.Debug(ctx, "In gslbldnsentriesSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["nodeid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Nodeid = types.Int64Value(intVal)
		}
	}

	return data
}
