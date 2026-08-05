package cmpaction

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
var _ resource.Resource = &CmpactionResource{}
var _ resource.ResourceWithConfigure = (*CmpactionResource)(nil)
var _ resource.ResourceWithImportState = (*CmpactionResource)(nil)

func NewCmpactionResource() resource.Resource {
	return &CmpactionResource{}
}

// CmpactionResource defines the resource implementation.
type CmpactionResource struct {
	client *service.NitroClient
}

func (r *CmpactionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *CmpactionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cmpaction"
}

func (r *CmpactionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *CmpactionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data CmpactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating cmpaction resource")
	// Get payload from plan
	cmpaction := cmpactionGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	cmpactionName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Cmpaction.Type(), cmpactionName, &cmpaction)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create cmpaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created cmpaction resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", cmpactionName))

	// Read the updated state back
	if !r.readCmpactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "cmpaction not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CmpactionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data CmpactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading cmpaction resource")

	found := r.readCmpactionFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *CmpactionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state CmpactionResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating cmpaction resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Addvaryheader.Equal(state.Addvaryheader) {
		tflog.Debug(ctx, "addvaryheader has changed for cmpaction")
		hasChange = true
	}
	if !data.Cmptype.Equal(state.Cmptype) {
		tflog.Debug(ctx, "cmptype has changed for cmpaction")
		hasChange = true
	}
	if !data.Varyheadervalue.Equal(state.Varyheadervalue) {
		tflog.Debug(ctx, "varyheadervalue has changed for cmpaction")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model, restricted to NITRO-updatable fields
		cmpaction := cmpactionGetTheUpdatablePayloadFromThePlan(ctx, &data)
		// Make API call
		// NITRO update for cmpaction is a PUT to /config/cmpaction (name carried in the
		// payload, not the URL) - use UpdateUnnamedResource (matches SDK v2 behavior).
		err := r.client.UpdateUnnamedResource(service.Cmpaction.Type(), &cmpaction)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update cmpaction, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated cmpaction resource")
	} else {
		tflog.Debug(ctx, "No changes detected for cmpaction resource, skipping update")
	}

	// Read the updated state back
	if !r.readCmpactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "cmpaction not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CmpactionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data CmpactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting cmpaction resource")
	// Named resource - delete using DeleteResource
	cmpactionName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Cmpaction.Type(), cmpactionName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete cmpaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted cmpaction resource")
}

// Helper function to read cmpaction data from API
func (r *CmpactionResource) readCmpactionFromApi(ctx context.Context, data *CmpactionResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	cmpactionName := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Cmpaction.Type(), cmpactionName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read cmpaction, got error: %s", err))
		return false
	}

	cmpactionSetAttrFromGet(ctx, data, getResponseData)

	return true
}
