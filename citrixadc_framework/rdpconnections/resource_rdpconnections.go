package rdpconnections

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
var _ resource.Resource = &RdpconnectionsResource{}
var _ resource.ResourceWithConfigure = (*RdpconnectionsResource)(nil)
var _ resource.ResourceWithImportState = (*RdpconnectionsResource)(nil)

func NewRdpconnectionsResource() resource.Resource {
	return &RdpconnectionsResource{}
}

// RdpconnectionsResource defines the resource implementation.
//
// rdpconnections is an action-only resource. NITRO exposes only get(all),
// count, and ?action=kill (POST). There is NO add, NO set/update, NO delete
// endpoint. The resource therefore models the "kill" action (Pattern 13):
// Create fires ?action=kill; Read/Update/Delete are no-ops. The read-only
// telemetry (endpointip, endpointport, targetip, targetport, peid) is exposed
// by the datasource, not the resource.
type RdpconnectionsResource struct {
	client *service.NitroClient
}

func (r *RdpconnectionsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *RdpconnectionsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rdpconnections"
}

func (r *RdpconnectionsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *RdpconnectionsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data RdpconnectionsResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating rdpconnections resource (firing ?action=kill)")
	payload := rdpconnectionsGetThePayloadFromthePlan(ctx, &data)

	// rdpconnections has no add/set endpoint. Fire the kill action.
	// NITRO verb is case-sensitive: "kill".
	err := r.client.ActOnResource(service.Rdpconnections.Type(), payload, "kill")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to kill rdpconnections, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Killed rdpconnections")

	// Synthetic ID: there is no persistent object to key off. Derive from the
	// kill selector so distinct kills are distinguishable, but the resource is
	// otherwise transient.
	if !data.Username.IsNull() && !data.Username.IsUnknown() && data.Username.ValueString() != "" {
		data.Id = types.StringValue(fmt.Sprintf("rdpconnections-kill-%s", data.Username.ValueString()))
	} else {
		data.Id = types.StringValue("rdpconnections-kill-all")
	}

	// Save data into Terraform state (no Read: transient table, no GET-by-key).
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Read is a no-op. rdpconnections is a transient diagnostics table with no
// GET-by-key endpoint; the killed connections are not a persistent object, so
// there is nothing to reconcile. Preserve prior state unchanged.
func (r *RdpconnectionsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data RdpconnectionsResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for rdpconnections (action-only kill, transient table)")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Update is a no-op. All attributes (username, all) are RequiresReplace, so
// Terraform never routes a change through Update. There is no set endpoint.
func (r *RdpconnectionsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state RdpconnectionsResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for rdpconnections; all attributes are RequiresReplace and there is no set endpoint")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Delete is a no-op. rdpconnections has no delete endpoint; the kill action is
// fire-and-forget. Removing the resource simply drops it from state.
func (r *RdpconnectionsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete is a no-op for rdpconnections (no delete endpoint on NITRO side)")
}
