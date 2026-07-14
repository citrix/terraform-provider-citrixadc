package vpnpcoipconnection

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &VpnpcoipconnectionResource{}
var _ resource.ResourceWithConfigure = (*VpnpcoipconnectionResource)(nil)
var _ resource.ResourceWithImportState = (*VpnpcoipconnectionResource)(nil)

func NewVpnpcoipconnectionResource() resource.Resource {
	return &VpnpcoipconnectionResource{}
}

// VpnpcoipconnectionResource defines the resource implementation.
// vpnpcoipconnection is an ACTION-ONLY resource: NITRO exposes get(all), count and
// kill (POST ?action=kill). There is no add, update/set, or delete endpoint.
type VpnpcoipconnectionResource struct {
	client *service.NitroClient
}

func (r *VpnpcoipconnectionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnpcoipconnectionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnpcoipconnection"
}

func (r *VpnpcoipconnectionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnpcoipconnectionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnpcoipconnectionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating (kill action) vpnpcoipconnection resource")
	payload := vpnpcoipconnectionGetThePayloadFromthePlan(ctx, &data)

	// kill is a POST ?action=kill action (Pattern 1). There is no add endpoint;
	// AddResource/UpdateUnnamedResource would target a nonexistent verb.
	err := r.client.ActOnResource(service.Vpnpcoipconnection.Type(), payload, "kill")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to kill vpnpcoipconnection, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Killed vpnpcoipconnection")

	// Set synthetic ID once, here in Create. The killed connection is not a
	// queryable managed object, so this ID is purely a Terraform state handle,
	// not a NITRO lookup key. Compose it from the supplied kill scope so distinct
	// kills produce distinct IDs; a bare "kill all" uses a static ID. username is
	// optional, so the ID must never be left empty (otherwise Terraform reports
	// "resource ID not set" after a successful apply).
	id := "kill-all"
	if !data.Username.IsNull() && !data.Username.IsUnknown() && data.Username.ValueString() != "" {
		id = fmt.Sprintf("kill:username:%s", data.Username.ValueString())
	}
	data.Id = types.StringValue(id)

	// Read is a no-op (no stable GET-backed managed object); state is what the
	// plan supplied plus the synthetic ID.
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnpcoipconnectionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnpcoipconnectionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read is a documented no-op. vpnpcoipconnection.kill terminates active PCoIP
	// connections; the killed connection is not a persistent object keyed by
	// these inputs, so a GET cannot stably re-resolve "this" record. Re-fetching
	// would cause perpetual drift / spurious resource removal. Preserve state.
	tflog.Debug(ctx, "Read is a no-op for vpnpcoipconnection (kill action has no stable GET-backed object); preserving state")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnpcoipconnectionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state VpnpcoipconnectionResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Update is a no-op: NITRO exposes no update/set endpoint for vpnpcoipconnection,
	// and every attribute is RequiresReplace, so Terraform never reaches here with
	// real diffs (Pattern 5).
	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for vpnpcoipconnection; no NITRO update endpoint and all attributes are RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnpcoipconnectionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnpcoipconnectionResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete is a state-only removal. "kill" is a one-shot action with no inverse
	// API; there is no NITRO delete endpoint for vpnpcoipconnection. Removing the
	// resource from Terraform state is the only meaningful operation.
	tflog.Debug(ctx, "Delete is a state-only removal for vpnpcoipconnection; kill has no inverse NITRO endpoint")
}
