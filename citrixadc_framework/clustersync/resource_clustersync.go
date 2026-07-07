package clustersync

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// clustersync is an ACTION-ONLY, ZERO-ATTRIBUTE resource.
//
//   - NITRO exposes only the Force action:
//     POST /nitro/v1/config/clustersync?action=Force, which forces a
//     synchronization of the cluster configuration across nodes.
//   - There is NO add/set/get/delete endpoint, so:
//     Create performs the Force action, Read/Update are no-ops (there is nothing
//     to reconcile), and Delete is a state-only removal.
//   - Because there is no GET endpoint, there is NO datasource for clustersync.
var _ resource.Resource = &ClustersyncResource{}
var _ resource.ResourceWithConfigure = (*ClustersyncResource)(nil)
var _ resource.ResourceWithImportState = (*ClustersyncResource)(nil)

func NewClustersyncResource() resource.Resource {
	return &ClustersyncResource{}
}

// ClustersyncResource defines the resource implementation.
type ClustersyncResource struct {
	client *service.NitroClient
}

func (r *ClustersyncResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ClustersyncResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_clustersync"
}

func (r *ClustersyncResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ClustersyncResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ClustersyncResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating clustersync resource (Force action)")
	clustersync := clustersyncGetThePayloadFromthePlan(ctx, &data)

	// Action-only resource - NITRO exposes only POST ?action=Force
	err := r.client.ActOnResource(service.Clustersync.Type(), &clustersync, "Force")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to force clustersync, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Forced clustersync")

	// Synthetic ID - no GET endpoint exists to derive it from
	data.Id = types.StringValue("clustersync-config")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Read is a no-op. clustersync has no GET endpoint; there is nothing to reconcile.
func (r *ClustersyncResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ClustersyncResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for clustersync; NITRO exposes no GET endpoint (action=Force only)")

	// Preserve prior state unchanged - no GET endpoint to reconcile against
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Update is a no-op. clustersync has no attributes and no set endpoint.
func (r *ClustersyncResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state ClustersyncResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Update is a no-op for clustersync; it has no attributes and no set endpoint")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Delete is a no-op. clustersync has no delete endpoint; the action is not
// reversible and there is no persistent object to remove.
func (r *ClustersyncResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete is a no-op for clustersync; NITRO has no delete endpoint")
}
