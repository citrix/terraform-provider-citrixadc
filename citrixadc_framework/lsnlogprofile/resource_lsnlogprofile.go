package lsnlogprofile

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
var _ resource.Resource = &LsnlogprofileResource{}
var _ resource.ResourceWithConfigure = (*LsnlogprofileResource)(nil)
var _ resource.ResourceWithImportState = (*LsnlogprofileResource)(nil)

func NewLsnlogprofileResource() resource.Resource {
	return &LsnlogprofileResource{}
}

// LsnlogprofileResource defines the resource implementation.
type LsnlogprofileResource struct {
	client *service.NitroClient
}

func (r *LsnlogprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LsnlogprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lsnlogprofile"
}

func (r *LsnlogprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LsnlogprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LsnlogprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lsnlogprofile resource")

	// Create API request body from the model
	lsnlogprofile := lsnlogprofileGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource
	lsnlogprofileName := data.Logprofilename.ValueString()
	_, err := r.client.AddResource(service.Lsnlogprofile.Type(), lsnlogprofileName, &lsnlogprofile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lsnlogprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created lsnlogprofile resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(lsnlogprofileName)

	// Read the updated state back
	if !r.readLsnlogprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "lsnlogprofile not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnlogprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LsnlogprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lsnlogprofile resource")

	found := r.readLsnlogprofileFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *LsnlogprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state LsnlogprofileResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating lsnlogprofile resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Analyticsprofile.Equal(state.Analyticsprofile) {
		tflog.Debug(ctx, "analyticsprofile has changed for lsnlogprofile")
		hasChange = true
	}
	if !data.Logcompact.Equal(state.Logcompact) {
		tflog.Debug(ctx, "logcompact has changed for lsnlogprofile")
		hasChange = true
	}
	if !data.Logipfix.Equal(state.Logipfix) {
		tflog.Debug(ctx, "logipfix has changed for lsnlogprofile")
		hasChange = true
	}
	if !data.Logsessdeletion.Equal(state.Logsessdeletion) {
		tflog.Debug(ctx, "logsessdeletion has changed for lsnlogprofile")
		hasChange = true
	}
	if !data.Logsubscrinfo.Equal(state.Logsubscrinfo) {
		tflog.Debug(ctx, "logsubscrinfo has changed for lsnlogprofile")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		lsnlogprofile := lsnlogprofileGetThePayloadFromthePlan(ctx, &data)
		// Named resource - use UpdateResource
		lsnlogprofileName := data.Logprofilename.ValueString()
		_, err := r.client.UpdateResource(service.Lsnlogprofile.Type(), lsnlogprofileName, &lsnlogprofile)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lsnlogprofile, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated lsnlogprofile resource")
	} else {
		tflog.Debug(ctx, "No changes detected for lsnlogprofile resource, skipping update")
	}

	// Read the updated state back
	if !r.readLsnlogprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "lsnlogprofile not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnlogprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LsnlogprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lsnlogprofile resource")
	// Named resource - delete using DeleteResource
	lsnlogprofileName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Lsnlogprofile.Type(), lsnlogprofileName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete lsnlogprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted lsnlogprofile resource")
}

// Helper function to read lsnlogprofile data from API
func (r *LsnlogprofileResource) readLsnlogprofileFromApi(ctx context.Context, data *LsnlogprofileResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	lsnlogprofileName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Lsnlogprofile.Type(), lsnlogprofileName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lsnlogprofile, got error: %s", err))
		return false
	}

	lsnlogprofileSetAttrFromGet(ctx, data, getResponseData)

	return true
}
