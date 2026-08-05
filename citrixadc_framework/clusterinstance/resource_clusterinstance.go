package clusterinstance

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ClusterinstanceResource{}
var _ resource.ResourceWithConfigure = (*ClusterinstanceResource)(nil)
var _ resource.ResourceWithImportState = (*ClusterinstanceResource)(nil)

func NewClusterinstanceResource() resource.Resource {
	return &ClusterinstanceResource{}
}

// ClusterinstanceResource defines the resource implementation.
type ClusterinstanceResource struct {
	client *service.NitroClient
}

func (r *ClusterinstanceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ClusterinstanceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_clusterinstance"
}

func (r *ClusterinstanceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ClusterinstanceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ClusterinstanceResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating clusterinstance resource")

	// Create API request body from the model
	clusterinstance := clusterinstanceGetThePayloadFromthePlan(ctx, &data)

	// Named resource keyed on clid - use AddResource (POST)
	clidValue := fmt.Sprintf("%d", data.Clid.ValueInt64())
	_, err := r.client.AddResource(service.Clusterinstance.Type(), clidValue, &clusterinstance)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create clusterinstance, got error: %s", err))
		return
	}

	// The ID is the clid value (backward-compatible with SDK v2 d.SetId(strconv.Itoa(clid)))
	data.Id = types.StringValue(clidValue)

	tflog.Trace(ctx, "Created clusterinstance resource")

	// Read the updated state back
	if !r.readClusterinstanceFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "clusterinstance not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusterinstanceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ClusterinstanceResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading clusterinstance resource")

	found := r.readClusterinstanceFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if !found {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusterinstanceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state ClusterinstanceResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating clusterinstance resource")

	// Check if there are any changes in updateable attributes (clid is RequiresReplace)
	hasChange := false
	if !data.Backplanebasedview.Equal(state.Backplanebasedview) {
		tflog.Debug(ctx, "backplanebasedview has changed for clusterinstance")
		hasChange = true
	}
	if !data.Clusterproxyarp.Equal(state.Clusterproxyarp) {
		tflog.Debug(ctx, "clusterproxyarp has changed for clusterinstance")
		hasChange = true
	}
	if !data.Deadinterval.Equal(state.Deadinterval) {
		tflog.Debug(ctx, "deadinterval has changed for clusterinstance")
		hasChange = true
	}
	if !data.Dfdretainl2params.Equal(state.Dfdretainl2params) {
		tflog.Debug(ctx, "dfdretainl2params has changed for clusterinstance")
		hasChange = true
	}
	if !data.Hellointerval.Equal(state.Hellointerval) {
		tflog.Debug(ctx, "hellointerval has changed for clusterinstance")
		hasChange = true
	}
	if !data.Inc.Equal(state.Inc) {
		tflog.Debug(ctx, "inc has changed for clusterinstance")
		hasChange = true
	}
	if !data.Nodegroup.Equal(state.Nodegroup) {
		tflog.Debug(ctx, "nodegroup has changed for clusterinstance")
		hasChange = true
	}
	if !data.Preemption.Equal(state.Preemption) {
		tflog.Debug(ctx, "preemption has changed for clusterinstance")
		hasChange = true
	}
	if !data.Processlocal.Equal(state.Processlocal) {
		tflog.Debug(ctx, "processlocal has changed for clusterinstance")
		hasChange = true
	}
	if !data.Quorumtype.Equal(state.Quorumtype) {
		tflog.Debug(ctx, "quorumtype has changed for clusterinstance")
		hasChange = true
	}
	if !data.Retainconnectionsoncluster.Equal(state.Retainconnectionsoncluster) {
		tflog.Debug(ctx, "retainconnectionsoncluster has changed for clusterinstance")
		hasChange = true
	}
	if !data.Secureheartbeats.Equal(state.Secureheartbeats) {
		tflog.Debug(ctx, "secureheartbeats has changed for clusterinstance")
		hasChange = true
	}
	if !data.Syncstatusstrictmode.Equal(state.Syncstatusstrictmode) {
		tflog.Debug(ctx, "syncstatusstrictmode has changed for clusterinstance")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		clusterinstance := clusterinstanceGetThePayloadFromthePlan(ctx, &data)

		// Update is PUT /clusterinstance (clid in body) - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Clusterinstance.Type(), &clusterinstance)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update clusterinstance, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated clusterinstance resource")
	} else {
		tflog.Debug(ctx, "No changes detected for clusterinstance resource, skipping update")
	}

	// Read the updated state back
	if !r.readClusterinstanceFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "clusterinstance not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusterinstanceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ClusterinstanceResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting clusterinstance resource")

	// Named resource - delete via DELETE /clusterinstance/{clid}
	err := r.client.DeleteResource(service.Clusterinstance.Type(), data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete clusterinstance, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted clusterinstance resource")
}

// Helper function to read clusterinstance data from API
func (r *ClusterinstanceResource) readClusterinstanceFromApi(ctx context.Context, data *ClusterinstanceResourceModel, diags *diag.Diagnostics) bool {
	// Named resource keyed on clid - the ID is the plain clid value
	clidName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Clusterinstance.Type(), clidName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read clusterinstance, got error: %s", err))
		return false
	}

	clusterinstanceSetAttrFromGet(ctx, data, getResponseData)

	return true
}
