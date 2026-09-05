package clusternode

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
var _ resource.Resource = &ClusternodeResource{}
var _ resource.ResourceWithConfigure = (*ClusternodeResource)(nil)
var _ resource.ResourceWithImportState = (*ClusternodeResource)(nil)

func NewClusternodeResource() resource.Resource {
	return &ClusternodeResource{}
}

// ClusternodeResource defines the resource implementation.
type ClusternodeResource struct {
	client *service.NitroClient
}

func (r *ClusternodeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ClusternodeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_clusternode"
}

func (r *ClusternodeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ClusternodeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ClusternodeResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating clusternode resource")

	// Create API request body from the model
	clusternode := clusternodeGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource keyed on nodeid
	clusternodeId := fmt.Sprintf("%d", data.Nodeid.ValueInt64())
	_, err := r.client.AddResource(service.Clusternode.Type(), clusternodeId, &clusternode)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create clusternode, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created clusternode resource")

	// Set ID for the resource before reading state back (matches SDK v2 id = nodeid)
	data.Id = types.StringValue(clusternodeId)

	// Read the updated state back
	if !r.readClusternodeFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "clusternode not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusternodeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ClusternodeResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading clusternode resource")

	found := r.readClusternodeFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *ClusternodeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state ClusternodeResourceModel

	// Read Terraform prior state to preserve ID and detect changes
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read Terraform config to detect attributes removed from config (for unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating clusternode resource")

	// Check if there are any changes in updateable attributes (matches SDK v2)
	hasChange := false
	attributesToUnset := []string{}
	if !data.Backplane.Equal(state.Backplane) {
		tflog.Debug(ctx, "backplane has changed for clusternode")
		hasChange = true
	}
	if !data.Delay.Equal(state.Delay) {
		tflog.Debug(ctx, "delay has changed for clusternode")
		if config.Delay.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "delay")
		} else {
			hasChange = true
		}
	}
	if !data.Nodegroup.Equal(state.Nodegroup) {
		tflog.Debug(ctx, "nodegroup has changed for clusternode")
		hasChange = true
	}
	if !data.Priority.Equal(state.Priority) {
		tflog.Debug(ctx, "priority has changed for clusternode")
		if config.Priority.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "priority")
		} else {
			hasChange = true
		}
	}
	if !data.State.Equal(state.State) {
		tflog.Debug(ctx, "state has changed for clusternode")
		if config.State.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "state")
		} else {
			hasChange = true
		}
	}
	if !data.Tunnelmode.Equal(state.Tunnelmode) {
		tflog.Debug(ctx, "tunnelmode has changed for clusternode")
		if config.Tunnelmode.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "tunnelmode")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		// Named resource - use UpdateResource keyed on nodeid
		clusternode := clusternodeGetTheUpdatablePayloadFromThePlan(ctx, &data)
		clusternodeId := data.Id.ValueString()
		_, err := r.client.UpdateResource(service.Clusternode.Type(), clusternodeId, &clusternode)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update clusternode, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated clusternode resource")
	} else {
		tflog.Debug(ctx, "No changes detected for clusternode resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their NITRO defaults. nodeid is the resource key and is required
	// in the unset payload.
	unsetIdPayload := map[string]interface{}{
		"nodeid": data.Nodeid.ValueInt64(),
	}
	if err := utils.ExecuteUnset(r.client, service.Clusternode.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset clusternode attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readClusternodeFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "clusternode not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusternodeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ClusternodeResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting clusternode resource")

	// Named resource - delete keyed on nodeid, passing delete-time args
	// clearnodegroupconfig (default YES) and force (only when true), matching SDK v2.
	args := make([]string, 0)
	if !data.Clearnodegroupconfig.IsNull() && data.Clearnodegroupconfig.ValueString() != "" {
		args = append(args, fmt.Sprintf("clearnodegroupconfig:%s", data.Clearnodegroupconfig.ValueString()))
	} else {
		args = append(args, "clearnodegroupconfig:YES")
	}
	if !data.Force.IsNull() && data.Force.ValueBool() {
		args = append(args, fmt.Sprintf("force:%t", data.Force.ValueBool()))
	}

	err := r.client.DeleteResourceWithArgs(service.Clusternode.Type(), data.Id.ValueString(), args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete clusternode, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted clusternode resource")
}

// Helper function to read clusternode data from API. Returns false when the resource no
// longer exists on the ADC so the caller can remove it from state.
func (r *ClusternodeResource) readClusternodeFromApi(ctx context.Context, data *ClusternodeResourceModel, diags *diag.Diagnostics) bool {
	// Case 2: Find with single ID attribute (nodeid) - ID is the plain value
	clusternodeId := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Clusternode.Type(), clusternodeId)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read clusternode, got error: %s", err))
		return false
	}

	clusternodeSetAttrFromGet(ctx, data, getResponseData)

	return true
}
