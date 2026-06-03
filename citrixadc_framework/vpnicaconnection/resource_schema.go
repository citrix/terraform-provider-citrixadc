package vpnicaconnection

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

// VpnicaconnectionResourceModel describes the resource data model.
// vpnicaconnection is an action-only resource (kill via ?action=kill).
// nodeid is intentionally absent: per the NITRO doc it is a GET-only filter
// argument, not a kill-payload property (Pattern 15).
type VpnicaconnectionResourceModel struct {
	Id         types.String `tfsdk:"id"`
	All        types.Bool   `tfsdk:"all"`
	Transproto types.String `tfsdk:"transproto"`
	Username   types.String `tfsdk:"username"`
}

func (r *VpnicaconnectionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnicaconnection resource.",
			},
			// User-facing kill arguments. Read is a no-op (kill is an action,
			// the killed connection is not a persistent managed object), so
			// these must NOT be Computed or Terraform reports an unknown value
			// after apply (Pattern 13 schema-flag implication).
			"all": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Terminate all active icaconnections.",
			},
			"transproto": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Transport type for the existing Existing ICA conenction.",
			},
			"username": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "User name for which ica connections needs to be terminated.",
			},
		},
	}
}

// vpnicaconnectionGetThePayloadFromthePlan builds the body for the kill action.
// The vendored vpn.Vpnicaconnection struct carries read-only fields and nodeid
// (a GET-only filter); build a map containing ONLY the kill arguments that are
// set so the action body never includes invalid arguments (Pattern 3 + 15).
func vpnicaconnectionGetThePayloadFromthePlan(ctx context.Context, data *VpnicaconnectionResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In vpnicaconnectionGetThePayloadFromthePlan Function")

	vpnicaconnection := map[string]interface{}{}
	if !data.All.IsNull() && !data.All.IsUnknown() {
		vpnicaconnection["all"] = data.All.ValueBool()
	}
	if !data.Transproto.IsNull() && !data.Transproto.IsUnknown() {
		vpnicaconnection["transproto"] = data.Transproto.ValueString()
	}
	if !data.Username.IsNull() && !data.Username.IsUnknown() {
		vpnicaconnection["username"] = data.Username.ValueString()
	}

	return vpnicaconnection
}
