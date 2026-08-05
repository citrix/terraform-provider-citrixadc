package lsnsipalgprofile

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
var _ resource.Resource = &LsnsipalgprofileResource{}
var _ resource.ResourceWithConfigure = (*LsnsipalgprofileResource)(nil)
var _ resource.ResourceWithImportState = (*LsnsipalgprofileResource)(nil)

func NewLsnsipalgprofileResource() resource.Resource {
	return &LsnsipalgprofileResource{}
}

// LsnsipalgprofileResource defines the resource implementation.
type LsnsipalgprofileResource struct {
	client *service.NitroClient
}

func (r *LsnsipalgprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LsnsipalgprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lsnsipalgprofile"
}

func (r *LsnsipalgprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LsnsipalgprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LsnsipalgprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lsnsipalgprofile resource")

	lsnsipalgprofile := lsnsipalgprofileGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Named resource - use AddResource (NITRO add is POST /nitro/v1/config/lsnsipalgprofile)
	lsnsipalgprofileName := data.Sipalgprofilename.ValueString()
	_, err := r.client.AddResource(service.Lsnsipalgprofile.Type(), lsnsipalgprofileName, &lsnsipalgprofile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lsnsipalgprofile, got error: %s", err))
		return
	}

	// Set ID for the resource before reading state
	data.Id = types.StringValue(lsnsipalgprofileName)

	tflog.Trace(ctx, "Created lsnsipalgprofile resource")

	// Read the updated state back
	if !r.readLsnsipalgprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "lsnsipalgprofile not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnsipalgprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LsnsipalgprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lsnsipalgprofile resource")

	found := r.readLsnsipalgprofileFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *LsnsipalgprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state LsnsipalgprofileResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating lsnsipalgprofile resource")

	// Create API request body from the model
	lsnsipalgprofile := lsnsipalgprofileGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// NITRO update is PUT /nitro/v1/config/lsnsipalgprofile (name carried in body) - matches SDK v2
	err := r.client.UpdateUnnamedResource(service.Lsnsipalgprofile.Type(), &lsnsipalgprofile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lsnsipalgprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Updated lsnsipalgprofile resource")

	// Read the updated state back
	if !r.readLsnsipalgprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "lsnsipalgprofile not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnsipalgprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LsnsipalgprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lsnsipalgprofile resource")

	// Named resource - delete using DeleteResource by ID (the live name)
	lsnsipalgprofileName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Lsnsipalgprofile.Type(), lsnsipalgprofileName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete lsnsipalgprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted lsnsipalgprofile resource")
}

// Helper function to read lsnsipalgprofile data from API.
// Returns false (without error) when the resource no longer exists on the ADC.
func (r *LsnsipalgprofileResource) readLsnsipalgprofileFromApi(ctx context.Context, data *LsnsipalgprofileResourceModel, diags *diag.Diagnostics) bool {
	// Case 2: Find with single ID attribute - ID is the plain value (sipalgprofilename)
	lsnsipalgprofileName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Lsnsipalgprofile.Type(), lsnsipalgprofileName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lsnsipalgprofile, got error: %s", err))
		return false
	}

	lsnsipalgprofileSetAttrFromGet(ctx, data, getResponseData)

	return true
}
