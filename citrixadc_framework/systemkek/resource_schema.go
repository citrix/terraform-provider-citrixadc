package systemkek

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/system"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SystemkekResourceModel describes the resource data model.
type SystemkekResourceModel struct {
	Id    types.String `tfsdk:"id"`
	Level types.String `tfsdk:"level"`
}

func (r *SystemkekResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the systemkek resource.",
			},
			"level": schema.StringAttribute{
				// CLI + NITRO mandatory (tfdata wrongly had is_required:false) -> Required (Pattern 8).
				// RequiresReplace: re-applying forces a fresh KEK rotation; there is no
				// update/GET endpoint so this is the only way to re-run the action.
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Type of update KEK to be performed.\n*basic : The level basic will backup old keys and create new keys and respond back.\n*extended : The level extended will backup old keys and create new keys, update\nns.conf, nscfg.db, all ns.conf for same release, in all partitions. While doing so\n will block all config changes and once done shall respond back.",
			},
		},
	}
}

func systemkekGetThePayloadFromthePlan(ctx context.Context, data *SystemkekResourceModel) system.Systemkek {
	tflog.Debug(ctx, "In systemkekGetThePayloadFromthePlan Function")

	// Create API request body from the model
	systemkek := system.Systemkek{}
	if !data.Level.IsNull() && !data.Level.IsUnknown() {
		systemkek.Level = data.Level.ValueString()
	}

	return systemkek
}
