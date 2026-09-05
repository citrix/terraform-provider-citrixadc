package quicbridgeprofile

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
var _ resource.Resource = &QuicbridgeprofileResource{}
var _ resource.ResourceWithConfigure = (*QuicbridgeprofileResource)(nil)
var _ resource.ResourceWithImportState = (*QuicbridgeprofileResource)(nil)

func NewQuicbridgeprofileResource() resource.Resource {
	return &QuicbridgeprofileResource{}
}

// QuicbridgeprofileResource defines the resource implementation.
type QuicbridgeprofileResource struct {
	client *service.NitroClient
}

func (r *QuicbridgeprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *QuicbridgeprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_quicbridgeprofile"
}

func (r *QuicbridgeprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *QuicbridgeprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data QuicbridgeprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating quicbridgeprofile resource")

	// Create API request body from the model
	quicbridgeprofile := quicbridgeprofileGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	quicbridgeprofileName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Quicbridgeprofile.Type(), quicbridgeprofileName, &quicbridgeprofile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create quicbridgeprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created quicbridgeprofile resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(quicbridgeprofileName)

	// Read the updated state back
	if !r.readQuicbridgeprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "quicbridgeprofile not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *QuicbridgeprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data QuicbridgeprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading quicbridgeprofile resource")

	found := r.readQuicbridgeprofileFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *QuicbridgeprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state QuicbridgeprofileResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from config (unset targets)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating quicbridgeprofile resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	attributesToUnset := []string{}
	if !data.Routingalgorithm.Equal(state.Routingalgorithm) {
		tflog.Debug(ctx, "routingalgorithm has changed for quicbridgeprofile")
		if config.Routingalgorithm.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "routingalgorithm")
		} else {
			hasChange = true
		}
	}
	if !data.Serveridlength.Equal(state.Serveridlength) {
		tflog.Debug(ctx, "serveridlength has changed for quicbridgeprofile")
		if config.Serveridlength.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "serveridlength")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		// Create API request body from the model
		quicbridgeprofile := quicbridgeprofileGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Named resource - use UpdateResource
		quicbridgeprofileName := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Quicbridgeprofile.Type(), quicbridgeprofileName, &quicbridgeprofile)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update quicbridgeprofile, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated quicbridgeprofile resource")
	} else {
		tflog.Debug(ctx, "No changes detected for quicbridgeprofile resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults.
	unsetIdPayload := map[string]interface{}{
		"name": data.Name.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Quicbridgeprofile.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset quicbridgeprofile attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readQuicbridgeprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "quicbridgeprofile not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *QuicbridgeprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data QuicbridgeprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting quicbridgeprofile resource")
	// Named resource - delete using DeleteResource
	quicbridgeprofileName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Quicbridgeprofile.Type(), quicbridgeprofileName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete quicbridgeprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted quicbridgeprofile resource")
}

// Helper function to read quicbridgeprofile data from API
func (r *QuicbridgeprofileResource) readQuicbridgeprofileFromApi(ctx context.Context, data *QuicbridgeprofileResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	quicbridgeprofileName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Quicbridgeprofile.Type(), quicbridgeprofileName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read quicbridgeprofile, got error: %s", err))
		return false
	}

	quicbridgeprofileSetAttrFromGet(ctx, data, getResponseData)

	return true
}
