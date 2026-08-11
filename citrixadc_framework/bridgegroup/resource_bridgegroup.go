package bridgegroup

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
var _ resource.Resource = &BridgegroupResource{}
var _ resource.ResourceWithConfigure = (*BridgegroupResource)(nil)
var _ resource.ResourceWithImportState = (*BridgegroupResource)(nil)

func NewBridgegroupResource() resource.Resource {
	return &BridgegroupResource{}
}

// BridgegroupResource defines the resource implementation.
type BridgegroupResource struct {
	client *service.NitroClient
}

func (r *BridgegroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *BridgegroupResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_bridgegroup"
}

func (r *BridgegroupResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *BridgegroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data BridgegroupResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating bridgegroup resource")

	bridgegroup := bridgegroupGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	bridgegroupIdStr := fmt.Sprintf("%d", data.Bridgegroupid.ValueInt64())
	_, err := r.client.AddResource(service.Bridgegroup.Type(), bridgegroupIdStr, &bridgegroup)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create bridgegroup, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created bridgegroup resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(bridgegroupIdStr)

	// Read the updated state back
	if !r.readBridgegroupFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "bridgegroup not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BridgegroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data BridgegroupResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading bridgegroup resource")

	found := r.readBridgegroupFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *BridgegroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state BridgegroupResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from config (for unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state (bridgegroup_id is RequiresReplace, so it never changes here)
	data.Id = state.Id

	tflog.Debug(ctx, "Updating bridgegroup resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	attributesToUnset := []string{}
	if !data.Dynamicrouting.Equal(state.Dynamicrouting) {
		tflog.Debug(ctx, "dynamicrouting has changed for bridgegroup")
		if config.Dynamicrouting.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "dynamicrouting")
		} else {
			hasChange = true
		}
	}
	if !data.Ipv6dynamicrouting.Equal(state.Ipv6dynamicrouting) {
		tflog.Debug(ctx, "ipv6dynamicrouting has changed for bridgegroup")
		if config.Ipv6dynamicrouting.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "ipv6dynamicrouting")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		// Create API request body from the model
		bridgegroup := bridgegroupGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Named resource - use UpdateResource
		_, err := r.client.UpdateResource(service.Bridgegroup.Type(), data.Id.ValueString(), &bridgegroup)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update bridgegroup, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated bridgegroup resource")
	} else {
		tflog.Debug(ctx, "No changes detected for bridgegroup resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults.
	unsetIdPayload := map[string]interface{}{
		"id": data.Bridgegroupid.ValueInt64(),
	}
	if err := utils.ExecuteUnset(r.client, service.Bridgegroup.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset bridgegroup attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readBridgegroupFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "bridgegroup not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BridgegroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data BridgegroupResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting bridgegroup resource")

	// Named resource - delete using DeleteResource
	err := r.client.DeleteResource(service.Bridgegroup.Type(), data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete bridgegroup, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted bridgegroup resource")
}

// Helper function to read bridgegroup data from API
func (r *BridgegroupResource) readBridgegroupFromApi(ctx context.Context, data *BridgegroupResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	bridgegroupName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Bridgegroup.Type(), bridgegroupName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read bridgegroup, got error: %s", err))
		return false
	}

	bridgegroupSetAttrFromGet(ctx, data, getResponseData)

	return true
}
