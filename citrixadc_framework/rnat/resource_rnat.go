package rnat

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/network"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &RnatResource{}
var _ resource.ResourceWithConfigure = (*RnatResource)(nil)
var _ resource.ResourceWithImportState = (*RnatResource)(nil)

func NewRnatResource() resource.Resource {
	return &RnatResource{}
}

// RnatResource defines the resource implementation.
type RnatResource struct {
	client *service.NitroClient
}

func (r *RnatResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *RnatResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rnat"
}

func (r *RnatResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *RnatResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data RnatResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating rnat resource")

	rnatName := data.Name.ValueString()

	// Build the add payload from the plan.
	rnat := rnatGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource.
	_, err := r.client.AddResource(service.Rnat.Type(), rnatName, &rnat)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create rnat, got error: %s", err))
		return
	}

	// ID is the rnat name (matches SDK v2 d.SetId(name)).
	data.Id = types.StringValue(rnatName)

	tflog.Trace(ctx, "Created rnat resource")

	// Read the updated state back.
	if !r.readRnatFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "rnat not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RnatResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data RnatResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading rnat resource")

	found := r.readRnatFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if !found {
		// Resource no longer exists on the appliance - clear it from state.
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RnatResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state RnatResourceModel

	// Read prior state to preserve the ID / live name.
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from config (-> unset).
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID (the current live name) from prior state.
	data.Id = state.Id

	tflog.Debug(ctx, "Updating rnat resource")

	// The current live name of the object is tracked by the ID, not by the
	// (possibly stale) configured name attribute.
	liveName := state.Id.ValueString()

	// 1) Regular in-place update of the SDK v2 updateable fields.
	hasChange := false
	attributesToUnset := []string{}
	if !data.Connfailover.Equal(state.Connfailover) {
		tflog.Debug(ctx, "connfailover has changed for rnat")
		if config.Connfailover.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "connfailover")
		} else {
			hasChange = true
		}
	}
	if !data.Ownergroup.Equal(state.Ownergroup) {
		tflog.Debug(ctx, "ownergroup has changed for rnat")
		hasChange = true
	}
	if !data.Redirectport.Equal(state.Redirectport) {
		tflog.Debug(ctx, "redirectport has changed for rnat")
		hasChange = true
	}
	if !data.Srcippersistency.Equal(state.Srcippersistency) {
		tflog.Debug(ctx, "srcippersistency has changed for rnat")
		if config.Srcippersistency.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "srcippersistency")
		} else {
			hasChange = true
		}
	}
	if !data.Td.Equal(state.Td) {
		tflog.Debug(ctx, "td has changed for rnat")
		hasChange = true
	}
	if !data.Useproxyport.Equal(state.Useproxyport) {
		tflog.Debug(ctx, "useproxyport has changed for rnat")
		if config.Useproxyport.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "useproxyport")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		rnat := rnatGetTheUpdatablePayloadFromThePlan(ctx, &data)
		// Ensure the payload keys off the current live name.
		rnat.Name = liveName
		if _, err := r.client.UpdateResource(service.Rnat.Type(), liveName, &rnat); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update rnat, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated rnat resource fields")
	} else {
		tflog.Debug(ctx, "No updateable field changes detected for rnat resource")
	}

	// Unset attributes removed from config so the appliance reverts them to
	// their NITRO defaults. Keyed off the current live name.
	unsetIdPayload := map[string]interface{}{
		"name": liveName,
	}
	if err := utils.ExecuteUnset(r.client, service.Rnat.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset rnat attributes, got error: %s", err))
		return
	}

	// 2) In-place rename via NITRO ?action=rename when newname is set/changed.
	if !data.Newname.IsNull() && !data.Newname.IsUnknown() && data.Newname.ValueString() != "" &&
		data.Newname.ValueString() != liveName {
		newName := data.Newname.ValueString()
		renamePayload := network.Rnat{
			Name:    liveName,
			Newname: newName,
		}
		if err := r.client.ActOnResource(service.Rnat.Type(), &renamePayload, "rename"); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to rename rnat, got error: %s", err))
			return
		}
		// The live object is now named newName; track it via the ID for the
		// read-back and all future reads/deletes.
		data.Id = types.StringValue(newName)
		tflog.Trace(ctx, "Renamed rnat resource")
	}

	// Read the updated state back (keyed off data.Id).
	if !r.readRnatFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "rnat not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RnatResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data RnatResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting rnat resource")

	// Named resource - delete by the live name (the ID), not the possibly-stale
	// configured name attribute (which is stale after a rename).
	err := r.client.DeleteResource(service.Rnat.Type(), data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete rnat, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted rnat resource")
}

// readRnatFromApi reads rnat data from the appliance keyed off the live name
// (data.Id). Returns false (without error) when the resource no longer exists.
func (r *RnatResource) readRnatFromApi(ctx context.Context, data *RnatResourceModel, diags *diag.Diagnostics) bool {
	rnatName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Rnat.Type(), rnatName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read rnat, got error: %s", err))
		return false
	}

	rnatSetAttrFromGet(ctx, data, getResponseData)

	return true
}
