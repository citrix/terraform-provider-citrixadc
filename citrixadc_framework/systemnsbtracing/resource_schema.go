package systemnsbtracing

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SystemnsbtracingResourceModel describes the resource data model.
type SystemnsbtracingResourceModel struct {
	Id     types.String `tfsdk:"id"`
	Nodeid types.Int64  `tfsdk:"nodeid"`
}

func (r *SystemnsbtracingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the systemnsbtracing resource.",
			},
			// nodeid is a GET-only query filter (args=nodeid) used to target a
			// specific cluster node; it is NOT part of the empty enable/disable
			// action payload (Pattern 15). Kept Optional (no Computed) so it can be
			// supplied as the GET filter; changing it re-targets a different node, so
			// RequiresReplace.
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

// systemnsbtracingSetAttrFromGet populates the resource model from the GET response.
// nodeid is a user-supplied GET filter, not echoed back as a settable property, so
// it is preserved from plan/state and not overwritten here. The synthetic ID is set
// exactly once in Create (Pattern 6), so it is NOT recomputed here.
func systemnsbtracingSetAttrFromGet(ctx context.Context, data *SystemnsbtracingResourceModel, getResponseData map[string]interface{}) *SystemnsbtracingResourceModel {
	tflog.Debug(ctx, "In systemnsbtracingSetAttrFromGet Function")

	// configuredstate / effectivestate are NITRO read-only properties not present in
	// the model; nodeid is a write-time GET filter preserved from plan/state. Nothing
	// to copy from the GET response.

	return data
}
