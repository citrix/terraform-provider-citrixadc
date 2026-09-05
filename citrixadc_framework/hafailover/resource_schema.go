package hafailover

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ha"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// HafailoverResourceModel describes the resource data model.
//
// hafailover is a synthetic action resource (NITRO exposes only the "Force"
// action, no GET). To preserve backward compatibility with the SDK v2
// implementation, the model carries the same three user-facing attributes:
//   - force     : whether to force the failover without confirmation (payload)
//   - ipaddress : IP address of the HA node whose state drives the action
//   - state     : the desired/observed HA state of that node
type HafailoverResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Force     types.Bool   `tfsdk:"force"`
	Ipaddress types.String `tfsdk:"ipaddress"`
	State     types.String `tfsdk:"state"`
}

func (r *HafailoverResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the hafailover resource.",
			},
			"force": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					// GH #1436: avoid spurious destroy+recreate on upgrade
					boolplanmodifier.UseStateForUnknown(),
					boolplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Force a failover without prompting for confirmation.",
			},
			"ipaddress": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The IP address of the HA node whose state is inspected to decide whether a failover must be forced.",
			},
			"state": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The desired HA state of the node identified by ipaddress. A failover is forced when the current state differs from this value.",
			},
		},
	}
}

// hafailoverGetThePayloadFromtheConfig builds the NITRO payload for the Force action.
func hafailoverGetThePayloadFromtheConfig(ctx context.Context, data *HafailoverResourceModel) ha.Hafailover {
	tflog.Debug(ctx, "In hafailoverGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	hafailover := ha.Hafailover{}
	if !data.Force.IsNull() && !data.Force.IsUnknown() {
		hafailover.Force = data.Force.ValueBool()
	}

	return hafailover
}
