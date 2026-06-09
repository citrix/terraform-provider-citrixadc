package hasync

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// HasyncResourceModel describes the resource data model.
type HasyncResourceModel struct {
	Id    types.String `tfsdk:"id"`
	Force types.Bool   `tfsdk:"force"`
	Save  types.String `tfsdk:"save"`
}

func (r *HasyncResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the hasync resource.",
			},
			"force": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Force synchronization regardless of the state of HA propagation and HA synchronization on either node.",
			},
			"save": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "After synchronization, automatically save the configuration in the secondary node configuration file (ns.conf) without prompting for confirmation.",
			},
		},
	}
}

func hasyncGetThePayloadFromthePlan(ctx context.Context, data *HasyncResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In hasyncGetThePayloadFromthePlan Function")

	// Build the Force action payload. force/save are included only when set.
	hasync := make(map[string]interface{})
	if !data.Force.IsNull() && !data.Force.IsUnknown() {
		hasync["force"] = data.Force.ValueBool()
	}
	if !data.Save.IsNull() && !data.Save.IsUnknown() {
		hasync["save"] = data.Save.ValueString()
	}

	return hasync
}
