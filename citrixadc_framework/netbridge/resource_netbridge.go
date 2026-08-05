package netbridge

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
var _ resource.Resource = &NetbridgeResource{}
var _ resource.ResourceWithConfigure = (*NetbridgeResource)(nil)
var _ resource.ResourceWithImportState = (*NetbridgeResource)(nil)

func NewNetbridgeResource() resource.Resource {
	return &NetbridgeResource{}
}

// NetbridgeResource defines the resource implementation.
type NetbridgeResource struct {
	client *service.NitroClient
}

func (r *NetbridgeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NetbridgeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_netbridge"
}

func (r *NetbridgeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NetbridgeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NetbridgeResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating netbridge resource")

	// Create API request body from the model
	netbridge := netbridgeGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	netbridgeName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Netbridge.Type(), netbridgeName, &netbridge)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create netbridge, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created netbridge resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	if !r.readNetbridgeFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "netbridge not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NetbridgeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetbridgeResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading netbridge resource")

	found := r.readNetbridgeFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *NetbridgeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state NetbridgeResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating netbridge resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Vxlanvlanmap.Equal(state.Vxlanvlanmap) {
		tflog.Debug(ctx, "vxlanvlanmap has changed for netbridge")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		netbridge := netbridgeGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Named resource - use UpdateResource
		netbridgeName := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Netbridge.Type(), netbridgeName, &netbridge)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update netbridge, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated netbridge resource")
	} else {
		tflog.Debug(ctx, "No changes detected for netbridge resource, skipping update")
	}

	// Read the updated state back
	if !r.readNetbridgeFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "netbridge not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NetbridgeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NetbridgeResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting netbridge resource")

	// Named resource - delete using DeleteResource
	netbridgeName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Netbridge.Type(), netbridgeName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete netbridge, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted netbridge resource")
}

// Helper function to read netbridge data from API
func (r *NetbridgeResource) readNetbridgeFromApi(ctx context.Context, data *NetbridgeResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	netbridgeName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Netbridge.Type(), netbridgeName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read netbridge, got error: %s", err))
		return false
	}

	netbridgeSetAttrFromGet(ctx, data, getResponseData)

	return true
}
