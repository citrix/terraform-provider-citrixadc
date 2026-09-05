package vrid

import (
	"context"
	"fmt"
	"strconv"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &VridResource{}
var _ resource.ResourceWithConfigure = (*VridResource)(nil)
var _ resource.ResourceWithImportState = (*VridResource)(nil)

func NewVridResource() resource.Resource {
	return &VridResource{}
}

// VridResource defines the resource implementation.
type VridResource struct {
	client *service.NitroClient
}

func (r *VridResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VridResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vrid"
}

func (r *VridResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VridResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VridResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vrid resource")

	vrid := vridGetThePayloadFromtheConfig(ctx, &data)

	// Named resource - use AddResource keyed by vrid_id (matches SDK v2 semantics)
	vridIdStr := strconv.Itoa(int(data.Vrid_id.ValueInt64()))
	_, err := r.client.AddResource(service.Vrid.Type(), vridIdStr, &vrid)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vrid, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created vrid resource")

	// Set ID for the resource before reading state (plain vrid_id value, matches SDK v2)
	data.Id = types.StringValue(vridIdStr)

	// Read the updated state back
	if !r.readVridFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "vrid not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VridResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VridResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vrid resource")

	found := r.readVridFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *VridResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state VridResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from config (for unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating vrid resource")

	// Check if there are any changes in updateable attributes (vrid_id is RequiresReplace)
	hasChange := false
	attributesToUnset := []string{}
	if !data.All.Equal(state.All) {
		tflog.Debug(ctx, "all has changed for vrid")
		hasChange = true
	}
	if !data.Ownernode.Equal(state.Ownernode) {
		tflog.Debug(ctx, "ownernode has changed for vrid")
		hasChange = true
	}
	if !data.Preemption.Equal(state.Preemption) {
		tflog.Debug(ctx, "preemption has changed for vrid")
		if config.Preemption.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "preemption")
		} else {
			hasChange = true
		}
	}
	if !data.Preemptiondelaytimer.Equal(state.Preemptiondelaytimer) {
		tflog.Debug(ctx, "preemptiondelaytimer has changed for vrid")
		if config.Preemptiondelaytimer.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "preemptiondelaytimer")
		} else {
			hasChange = true
		}
	}
	if !data.Priority.Equal(state.Priority) {
		tflog.Debug(ctx, "priority has changed for vrid")
		if config.Priority.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "priority")
		} else {
			hasChange = true
		}
	}
	if !data.Sharing.Equal(state.Sharing) {
		tflog.Debug(ctx, "sharing has changed for vrid")
		if config.Sharing.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "sharing")
		} else {
			hasChange = true
		}
	}
	if !data.Trackifnumpriority.Equal(state.Trackifnumpriority) {
		tflog.Debug(ctx, "trackifnumpriority has changed for vrid")
		if config.Trackifnumpriority.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "trackifnumpriority")
		} else {
			hasChange = true
		}
	}
	if !data.Tracking.Equal(state.Tracking) {
		tflog.Debug(ctx, "tracking has changed for vrid")
		if config.Tracking.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "tracking")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		vrid := vridGetThePayloadFromtheConfig(ctx, &data)
		// Named resource - use UpdateResource keyed by vrid_id
		vridIdStr := strconv.Itoa(int(data.Vrid_id.ValueInt64()))
		_, err := r.client.UpdateResource(service.Vrid.Type(), vridIdStr, &vrid)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vrid, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated vrid resource")
	} else {
		tflog.Debug(ctx, "No changes detected for vrid resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults.
	unsetIdPayload := map[string]interface{}{
		"id": int(data.Vrid_id.ValueInt64()),
	}
	if err := utils.ExecuteUnset(r.client, service.Vrid.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset vrid attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readVridFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "vrid not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VridResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VridResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vrid resource")

	// Named resource - delete using DeleteResource keyed by the ID (vrid_id value)
	err := r.client.DeleteResource(service.Vrid.Type(), data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete vrid, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted vrid resource")
}

// Helper function to read vrid data from API. Returns false when the resource
// no longer exists on the ADC (so callers can remove it from state).
func (r *VridResource) readVridFromApi(ctx context.Context, data *VridResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain vrid_id value
	vridName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Vrid.Type(), vridName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vrid, got error: %s", err))
		return false
	}

	vridSetAttrFromGet(ctx, data, getResponseData)

	return true
}
