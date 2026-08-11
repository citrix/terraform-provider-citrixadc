package vrid6

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
var _ resource.Resource = &Vrid6Resource{}
var _ resource.ResourceWithConfigure = (*Vrid6Resource)(nil)
var _ resource.ResourceWithImportState = (*Vrid6Resource)(nil)

func NewVrid6Resource() resource.Resource {
	return &Vrid6Resource{}
}

// Vrid6Resource defines the resource implementation.
type Vrid6Resource struct {
	client *service.NitroClient
}

func (r *Vrid6Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *Vrid6Resource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vrid6"
}

func (r *Vrid6Resource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *Vrid6Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data Vrid6ResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vrid6 resource")

	vrid6 := vrid6GetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource keyed by vrid6_id (SDK v2 parity)
	vrid6IdStr := strconv.Itoa(int(data.Vrid6_id.ValueInt64()))
	_, err := r.client.AddResource(service.Vrid6.Type(), vrid6IdStr, &vrid6)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vrid6, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created vrid6 resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(vrid6IdStr)

	// Read the updated state back
	if !r.readVrid6FromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "vrid6 not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Vrid6Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data Vrid6ResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vrid6 resource")

	found := r.readVrid6FromApi(ctx, &data, &resp.Diagnostics)
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

func (r *Vrid6Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state Vrid6ResourceModel

	// Read Terraform prior state to preserve ID and detect changes
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from config (candidates for unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating vrid6 resource")

	// Check if there are any changes in updateable attributes (SDK v2 parity).
	// vrid6_id is ForceNew (RequiresReplace) and never reaches Update.
	hasChange := false
	attributesToUnset := []string{}
	if !data.All.Equal(state.All) {
		tflog.Debug(ctx, "all has changed for vrid6")
		hasChange = true
	}
	if !data.Ownernode.Equal(state.Ownernode) {
		tflog.Debug(ctx, "ownernode has changed for vrid6")
		hasChange = true
	}
	if !data.Preemption.Equal(state.Preemption) {
		tflog.Debug(ctx, "preemption has changed for vrid6")
		if config.Preemption.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "preemption")
		} else {
			hasChange = true
		}
	}
	if !data.Preemptiondelaytimer.Equal(state.Preemptiondelaytimer) {
		tflog.Debug(ctx, "preemptiondelaytimer has changed for vrid6")
		if config.Preemptiondelaytimer.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "preemptiondelaytimer")
		} else {
			hasChange = true
		}
	}
	if !data.Priority.Equal(state.Priority) {
		tflog.Debug(ctx, "priority has changed for vrid6")
		if config.Priority.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "priority")
		} else {
			hasChange = true
		}
	}
	if !data.Sharing.Equal(state.Sharing) {
		tflog.Debug(ctx, "sharing has changed for vrid6")
		if config.Sharing.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "sharing")
		} else {
			hasChange = true
		}
	}
	if !data.Trackifnumpriority.Equal(state.Trackifnumpriority) {
		tflog.Debug(ctx, "trackifnumpriority has changed for vrid6")
		if config.Trackifnumpriority.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "trackifnumpriority")
		} else {
			hasChange = true
		}
	}
	if !data.Tracking.Equal(state.Tracking) {
		tflog.Debug(ctx, "tracking has changed for vrid6")
		if config.Tracking.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "tracking")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		vrid6 := vrid6GetThePayloadFromthePlan(ctx, &data)
		// Named resource - use UpdateResource keyed by vrid6_id (SDK v2 parity)
		vrid6IdStr := strconv.Itoa(int(data.Vrid6_id.ValueInt64()))
		_, err := r.client.UpdateResource(service.Vrid6.Type(), vrid6IdStr, &vrid6)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vrid6, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated vrid6 resource")
	} else {
		tflog.Debug(ctx, "No changes detected for vrid6 resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their NITRO defaults. The vrid6 resource is keyed by its numeric
	// id, so the unset payload carries "id".
	unsetIdPayload := map[string]interface{}{
		"id": strconv.Itoa(int(data.Vrid6_id.ValueInt64())),
	}
	if err := utils.ExecuteUnset(r.client, service.Vrid6.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset vrid6 attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readVrid6FromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "vrid6 not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Vrid6Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data Vrid6ResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vrid6 resource")

	// Named resource - delete using DeleteResource keyed by the ID (vrid6_id)
	err := r.client.DeleteResource(service.Vrid6.Type(), data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete vrid6, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted vrid6 resource")
}

// Helper function to read vrid6 data from API.
// Returns false (without adding an error) when the resource no longer exists.
func (r *Vrid6Resource) readVrid6FromApi(ctx context.Context, data *Vrid6ResourceModel, diags *diag.Diagnostics) bool {
	// Case 2: Find with single ID attribute - ID is the plain value
	vrid6Name := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Vrid6.Type(), vrid6Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vrid6, got error: %s", err))
		return false
	}

	vrid6SetAttrFromGet(ctx, data, getResponseData)

	return true
}
