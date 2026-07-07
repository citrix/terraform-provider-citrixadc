package cluster

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// cluster is an ACTION-ONLY resource.
//
//   - NITRO exposes only the "join" action for the cluster object:
//     POST /nitro/v1/config/cluster?action=join, which joins the current node
//     to the cluster identified by the cluster IP (clip) using the CCO nsroot
//     password.
//   - There is NO add/set/get/delete endpoint for this object, so:
//     Create performs the join action, Read/Update are no-ops (there is nothing
//     to reconcile / no GET endpoint), and Delete is a state-only removal.
//   - Because there is no GET endpoint, there is NO datasource for cluster and
//     the resource cannot be verified by reading it back.
//   - WARNING: applying this resource JOINS the appliance to a cluster, which is
//     a disruptive operation. It is intended for deliberate, operator-initiated
//     use only.
var _ resource.Resource = &ClusterResource{}
var _ resource.ResourceWithConfigure = (*ClusterResource)(nil)
var _ resource.ResourceWithImportState = (*ClusterResource)(nil)

func NewClusterResource() resource.Resource {
	return &ClusterResource{}
}

// ClusterResource defines the resource implementation.
type ClusterResource struct {
	client *service.NitroClient
}

func (r *ClusterResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ClusterResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cluster"
}

func (r *ClusterResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ClusterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ClusterResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Joining the node to the cluster (cluster join action)")
	cluster := clusterGetThePayloadFromtheConfig(ctx, &data)

	// Action-only resource - the only NITRO operation is the "join" action
	// (POST /nitro/v1/config/cluster?action=join).
	err := r.client.ActOnResource(service.Cluster.Type(), &cluster, "join")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to join the cluster, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Triggered cluster join")

	// Synthetic ID - there is no GET endpoint to read back. Use the cluster IP.
	data.Id = types.StringValue(data.Clip.ValueString())

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Read is a no-op. cluster has no GET endpoint; there is nothing to reconcile.
func (r *ClusterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ClusterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for cluster; NITRO exposes no GET endpoint")

	// Preserve prior state unchanged.
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Update is a no-op. cluster exposes only the join action and all attributes
// use RequiresReplace, so Terraform will never reach this with an in-place
// change.
func (r *ClusterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state ClusterResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Update is a no-op for cluster; it exposes only the join action and all attributes are RequiresReplace")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Delete is a no-op. cluster has no delete endpoint; the join action is not
// reversible through this object and there is no persistent object to remove.
func (r *ClusterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete is a no-op for cluster; NITRO has no delete endpoint")
}
