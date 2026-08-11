package ipsecalgprofile

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
var _ resource.Resource = &IpsecalgprofileResource{}
var _ resource.ResourceWithConfigure = (*IpsecalgprofileResource)(nil)
var _ resource.ResourceWithImportState = (*IpsecalgprofileResource)(nil)

func NewIpsecalgprofileResource() resource.Resource {
	return &IpsecalgprofileResource{}
}

// IpsecalgprofileResource defines the resource implementation.
type IpsecalgprofileResource struct {
	client *service.NitroClient
}

func (r *IpsecalgprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *IpsecalgprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ipsecalgprofile"
}

func (r *IpsecalgprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *IpsecalgprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data IpsecalgprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating ipsecalgprofile resource")

	// Create API request body from the model
	ipsecalgprofile := ipsecalgprofileGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	ipsecalgprofileName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Ipsecalgprofile.Type(), ipsecalgprofileName, &ipsecalgprofile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create ipsecalgprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created ipsecalgprofile resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(ipsecalgprofileName)

	// Read the updated state back
	if !r.readIpsecalgprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "ipsecalgprofile not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IpsecalgprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data IpsecalgprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading ipsecalgprofile resource")

	found := r.readIpsecalgprofileFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *IpsecalgprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state IpsecalgprofileResourceModel

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

	tflog.Debug(ctx, "Updating ipsecalgprofile resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	attributesToUnset := []string{}
	if !data.Connfailover.Equal(state.Connfailover) {
		tflog.Debug(ctx, "connfailover has changed for ipsecalgprofile")
		if config.Connfailover.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "connfailover")
		} else {
			hasChange = true
		}
	}
	if !data.Espgatetimeout.Equal(state.Espgatetimeout) {
		tflog.Debug(ctx, "espgatetimeout has changed for ipsecalgprofile")
		if config.Espgatetimeout.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "espgatetimeout")
		} else {
			hasChange = true
		}
	}
	if !data.Espsessiontimeout.Equal(state.Espsessiontimeout) {
		tflog.Debug(ctx, "espsessiontimeout has changed for ipsecalgprofile")
		if config.Espsessiontimeout.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "espsessiontimeout")
		} else {
			hasChange = true
		}
	}
	if !data.Ikesessiontimeout.Equal(state.Ikesessiontimeout) {
		tflog.Debug(ctx, "ikesessiontimeout has changed for ipsecalgprofile")
		if config.Ikesessiontimeout.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "ikesessiontimeout")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		// Create API request body from the model
		ipsecalgprofile := ipsecalgprofileGetThePayloadFromtheConfig(ctx, &data)

		// Make API call
		// Matches SDK v2: PUT to /config/ipsecalgprofile with name in the payload
		err := r.client.UpdateUnnamedResource(service.Ipsecalgprofile.Type(), &ipsecalgprofile)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update ipsecalgprofile, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated ipsecalgprofile resource")
	} else {
		tflog.Debug(ctx, "No changes detected for ipsecalgprofile resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults.
	unsetIdPayload := map[string]interface{}{
		"name": data.Name.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Ipsecalgprofile.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset ipsecalgprofile attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readIpsecalgprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "ipsecalgprofile not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IpsecalgprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data IpsecalgprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting ipsecalgprofile resource")

	// Named resource - delete using DeleteResource
	ipsecalgprofileName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Ipsecalgprofile.Type(), ipsecalgprofileName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete ipsecalgprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted ipsecalgprofile resource")
}

// Helper function to read ipsecalgprofile data from API
func (r *IpsecalgprofileResource) readIpsecalgprofileFromApi(ctx context.Context, data *IpsecalgprofileResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	ipsecalgprofileName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Ipsecalgprofile.Type(), ipsecalgprofileName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read ipsecalgprofile, got error: %s", err))
		return false
	}

	ipsecalgprofileSetAttrFromGet(ctx, data, getResponseData)

	return true
}
