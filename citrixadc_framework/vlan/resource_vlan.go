package vlan

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
var _ resource.Resource = &VlanResource{}
var _ resource.ResourceWithConfigure = (*VlanResource)(nil)
var _ resource.ResourceWithImportState = (*VlanResource)(nil)

func NewVlanResource() resource.Resource {
	return &VlanResource{}
}

// VlanResource defines the resource implementation.
type VlanResource struct {
	client *service.NitroClient
}

func (r *VlanResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VlanResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vlan"
}

func (r *VlanResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VlanResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VlanResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vlan resource")

	// Create API request body from the model
	vlan := vlanGetThePayloadFromtheConfig(ctx, &data)

	// Named resource keyed on vlanid - use AddResource
	vlanIdStr := fmt.Sprintf("%d", data.Vlanid.ValueInt64())
	_, err := r.client.AddResource(service.Vlan.Type(), vlanIdStr, &vlan)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vlan, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created vlan resource")

	// Set ID for the resource before reading state (SDK v2 used the vlanid string as the ID)
	data.Id = types.StringValue(vlanIdStr)

	// Read the updated state back
	if !r.readVlanFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "vlan not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VlanResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VlanResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vlan resource")

	found := r.readVlanFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *VlanResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state VlanResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from config (for unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state (vlanid is RequiresReplace, so it never changes here)
	data.Id = state.Id

	tflog.Debug(ctx, "Updating vlan resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	attributesToUnset := []string{}
	if !data.Aliasname.Equal(state.Aliasname) {
		tflog.Debug(ctx, "aliasname has changed for vlan")
		hasChange = true
	}
	if !data.Dynamicrouting.Equal(state.Dynamicrouting) {
		tflog.Debug(ctx, "dynamicrouting has changed for vlan")
		if config.Dynamicrouting.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "dynamicrouting")
		} else {
			hasChange = true
		}
	}
	if !data.Ipv6dynamicrouting.Equal(state.Ipv6dynamicrouting) {
		tflog.Debug(ctx, "ipv6dynamicrouting has changed for vlan")
		if config.Ipv6dynamicrouting.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "ipv6dynamicrouting")
		} else {
			hasChange = true
		}
	}
	if !data.Mtu.Equal(state.Mtu) {
		tflog.Debug(ctx, "mtu has changed for vlan")
		hasChange = true
	}
	if !data.Sharing.Equal(state.Sharing) {
		tflog.Debug(ctx, "sharing has changed for vlan")
		if config.Sharing.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "sharing")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		// Create API request body from the model
		vlan := vlanGetThePayloadFromtheConfig(ctx, &data)

		// Named resource - use UpdateResource
		vlanIdStr := data.Id.ValueString()
		_, err := r.client.UpdateResource(service.Vlan.Type(), vlanIdStr, &vlan)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vlan, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated vlan resource")
	} else {
		tflog.Debug(ctx, "No changes detected for vlan resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults.
	unsetIdPayload := map[string]interface{}{
		"id": data.Vlanid.ValueInt64(),
	}
	if err := utils.ExecuteUnset(r.client, service.Vlan.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset vlan attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readVlanFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "vlan not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VlanResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VlanResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vlan resource")

	// Named resource - delete using DeleteResource
	err := r.client.DeleteResource(service.Vlan.Type(), data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete vlan, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted vlan resource")
}

// Helper function to read vlan data from API
func (r *VlanResource) readVlanFromApi(ctx context.Context, data *VlanResourceModel, diags *diag.Diagnostics) bool {
	// Case 2: Find with single ID attribute - ID is the plain vlanid value
	vlanIdStr := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Vlan.Type(), vlanIdStr)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vlan, got error: %s", err))
		return false
	}

	vlanSetAttrFromGet(ctx, data, getResponseData)

	return true
}
