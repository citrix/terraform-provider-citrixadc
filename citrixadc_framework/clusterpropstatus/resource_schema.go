package clusterpropstatus

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ClusterpropstatusResourceModel describes the resource data model.
// clusterpropstatus is an action-only resource (clear via ?action=clear).
// nodeid is the only writable filter.
type ClusterpropstatusResourceModel struct {
	Id     types.String `tfsdk:"id"`
	Nodeid types.Int64  `tfsdk:"nodeid"`
}

func (r *ClusterpropstatusResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the clusterpropstatus resource.",
			},
			// nodeid is the only clear argument. Read is a no-op (clear is an
			// action, clusterpropstatus is not a persistent managed object), so
			// nodeid must NOT be Computed or Terraform reports an unknown value
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

// clusterpropstatusGetThePayloadFromthePlan builds the body for the clear action.
// Build a map containing ONLY the clear arguments that are set so the action body
// never includes invalid arguments. The read-only output fields
// (numpropcmdfailed, cmdstrs) and internal/meta fields (__count,
// _nextgenapiresource) are intentionally excluded.
func clusterpropstatusGetThePayloadFromthePlan(ctx context.Context, data *ClusterpropstatusResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In clusterpropstatusGetThePayloadFromthePlan Function")

	clusterpropstatus := map[string]interface{}{}
	if !data.Nodeid.IsNull() && !data.Nodeid.IsUnknown() {
		clusterpropstatus["nodeid"] = int(data.Nodeid.ValueInt64())
	}

	return clusterpropstatus
}
