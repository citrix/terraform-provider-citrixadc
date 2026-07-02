package lldpneighbors

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// LldpneighborsResourceModel describes the resource data model.
type LldpneighborsResourceModel struct {
	Id     types.String `tfsdk:"id"`
	Ifnum  types.String `tfsdk:"ifnum"`
	Nodeid types.Int64  `tfsdk:"nodeid"`
}

func (r *LldpneighborsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lldpneighbors resource.",
			},
			// lldpneighbors is an action-only (clear) resource with no GET/add
			// endpoint. ifnum/nodeid are GET filters, not clear-action args, so
			// they are plain Optional inputs here (not Computed - Read is a no-op
			// and could never resolve an "unknown" value at apply time).
			"ifnum": schema.StringAttribute{
				Optional:    true,
				Description: "Interface Name",
			},
			"nodeid": schema.Int64Attribute{
				Optional:    true,
				Description: "Unique number that identifies the cluster node.",
			},
		},
	}
}

// lldpneighborsSetAttrFromGetForDatasource faithfully copies the GET response
// fields into the model for the datasource. The datasource has no prior
// plan/state to preserve, and it sets its own synthetic ID.
func lldpneighborsSetAttrFromGetForDatasource(ctx context.Context, data *LldpneighborsResourceModel, getResponseData map[string]interface{}) *LldpneighborsResourceModel {
	tflog.Debug(ctx, "In lldpneighborsSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["ifnum"]; ok && val != nil {
		if s, ok := val.(string); ok {
			data.Ifnum = types.StringValue(s)
		}
	}
	if val, ok := getResponseData["nodeid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Nodeid = types.Int64Value(intVal)
		}
	}

	return data
}
