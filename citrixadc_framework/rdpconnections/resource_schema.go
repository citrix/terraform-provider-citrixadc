package rdpconnections

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

// RdpconnectionsResourceModel describes the resource data model. The resource
// is an action-only "kill": only the two kill selectors (username, all) plus a
// synthetic id are modelled here. The read-only telemetry lives on the
// datasource model, not here.
type RdpconnectionsResourceModel struct {
	Id       types.String `tfsdk:"id"`
	All      types.Bool   `tfsdk:"all"`
	Username types.String `tfsdk:"username"`
}

func (r *RdpconnectionsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the rdpconnections resource.",
			},
			// kill selectors. Both Optional (one-of semantics, not enforced as a
			// group). Not Computed: Read is a no-op, so no server value ever
			// resolves them. RequiresReplace: any change re-fires the kill.
			"all": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Terminate all active rdpconnections.",
			},
			"username": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "User name for which to display connections.",
			},
		},
	}
}

// rdpconnectionsGetThePayloadFromthePlan builds the ?action=kill payload. Null
// selectors are omitted so a bare kill (kill everything the appliance decides)
// and all=true are both expressible. Returns map[string]interface{} because the
// kill action needs only the two selectors, not the full vendored struct.
func rdpconnectionsGetThePayloadFromthePlan(ctx context.Context, data *RdpconnectionsResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In rdpconnectionsGetThePayloadFromthePlan Function")

	payload := map[string]interface{}{}
	if !data.Username.IsNull() && !data.Username.IsUnknown() {
		payload["username"] = data.Username.ValueString()
	}
	if !data.All.IsNull() && !data.All.IsUnknown() {
		payload["all"] = data.All.ValueBool()
	}

	return payload
}
