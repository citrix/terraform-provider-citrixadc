package dbdbprofile

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
var _ resource.Resource = &DbdbprofileResource{}
var _ resource.ResourceWithConfigure = (*DbdbprofileResource)(nil)
var _ resource.ResourceWithImportState = (*DbdbprofileResource)(nil)

func NewDbdbprofileResource() resource.Resource {
	return &DbdbprofileResource{}
}

// DbdbprofileResource defines the resource implementation.
type DbdbprofileResource struct {
	client *service.NitroClient
}

func (r *DbdbprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *DbdbprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dbdbprofile"
}

func (r *DbdbprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *DbdbprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data DbdbprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating dbdbprofile resource")

	// Create API request body from the model
	dbdbprofile := dbdbprofileGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	dbdbprofileName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Dbdbprofile.Type(), dbdbprofileName, &dbdbprofile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create dbdbprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created dbdbprofile resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	if !r.readDbdbprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "dbdbprofile not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DbdbprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DbdbprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading dbdbprofile resource")

	found := r.readDbdbprofileFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *DbdbprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state DbdbprofileResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from configuration (for unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating dbdbprofile resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	attributesToUnset := []string{}
	if !data.Conmultiplex.Equal(state.Conmultiplex) {
		tflog.Debug(ctx, "conmultiplex has changed for dbdbprofile")
		if config.Conmultiplex.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "conmultiplex")
		} else {
			hasChange = true
		}
	}
	if !data.Enablecachingconmuxoff.Equal(state.Enablecachingconmuxoff) {
		tflog.Debug(ctx, "enablecachingconmuxoff has changed for dbdbprofile")
		if config.Enablecachingconmuxoff.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "enablecachingconmuxoff")
		} else {
			hasChange = true
		}
	}
	if !data.Interpretquery.Equal(state.Interpretquery) {
		tflog.Debug(ctx, "interpretquery has changed for dbdbprofile")
		if config.Interpretquery.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "interpretquery")
		} else {
			hasChange = true
		}
	}
	if !data.Kcdaccount.Equal(state.Kcdaccount) {
		tflog.Debug(ctx, "kcdaccount has changed for dbdbprofile")
		hasChange = true
	}
	if !data.Stickiness.Equal(state.Stickiness) {
		tflog.Debug(ctx, "stickiness has changed for dbdbprofile")
		if config.Stickiness.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "stickiness")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		// Create API request body from the model
		dbdbprofile := dbdbprofileGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Update URL carries the name in the payload (no name path segment)
		err := r.client.UpdateUnnamedResource(service.Dbdbprofile.Type(), &dbdbprofile)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update dbdbprofile, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated dbdbprofile resource")
	} else {
		tflog.Debug(ctx, "No changes detected for dbdbprofile resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults.
	unsetIdPayload := map[string]interface{}{
		"name": data.Name.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Dbdbprofile.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset dbdbprofile attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readDbdbprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "dbdbprofile not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DbdbprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data DbdbprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting dbdbprofile resource")
	// Named resource - delete using DeleteResource
	dbdbprofileName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Dbdbprofile.Type(), dbdbprofileName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete dbdbprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted dbdbprofile resource")
}

// Helper function to read dbdbprofile data from API
func (r *DbdbprofileResource) readDbdbprofileFromApi(ctx context.Context, data *DbdbprofileResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	dbdbprofileName := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Dbdbprofile.Type(), dbdbprofileName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read dbdbprofile, got error: %s", err))
		return false
	}

	dbdbprofileSetAttrFromGet(ctx, data, getResponseData)

	return true
}
