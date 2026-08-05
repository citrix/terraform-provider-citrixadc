package lsnappsprofile

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
var _ resource.Resource = &LsnappsprofileResource{}
var _ resource.ResourceWithConfigure = (*LsnappsprofileResource)(nil)
var _ resource.ResourceWithImportState = (*LsnappsprofileResource)(nil)

func NewLsnappsprofileResource() resource.Resource {
	return &LsnappsprofileResource{}
}

// LsnappsprofileResource defines the resource implementation.
type LsnappsprofileResource struct {
	client *service.NitroClient
}

func (r *LsnappsprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LsnappsprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lsnappsprofile"
}

func (r *LsnappsprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LsnappsprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LsnappsprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lsnappsprofile resource")

	lsnappsprofile := lsnappsprofileGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource (POST)
	appsprofilename := data.Appsprofilename.ValueString()
	_, err := r.client.AddResource(service.Lsnappsprofile.Type(), appsprofilename, &lsnappsprofile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lsnappsprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created lsnappsprofile resource")

	// Generate the resource ID (single unique attribute - plain value)
	data.Id = types.StringValue(appsprofilename)

	// Read the updated state back
	if !r.readLsnappsprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "lsnappsprofile not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnappsprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LsnappsprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lsnappsprofile resource")

	found := r.readLsnappsprofileFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *LsnappsprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state LsnappsprofileResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating lsnappsprofile resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Filtering.Equal(state.Filtering) {
		tflog.Debug(ctx, "filtering has changed for lsnappsprofile")
		hasChange = true
	}
	if !data.Ippooling.Equal(state.Ippooling) {
		tflog.Debug(ctx, "ippooling has changed for lsnappsprofile")
		hasChange = true
	}
	if !data.L2info.Equal(state.L2info) {
		tflog.Debug(ctx, "l2info has changed for lsnappsprofile")
		hasChange = true
	}
	if !data.Mapping.Equal(state.Mapping) {
		tflog.Debug(ctx, "mapping has changed for lsnappsprofile")
		hasChange = true
	}
	if !data.Tcpproxy.Equal(state.Tcpproxy) {
		tflog.Debug(ctx, "tcpproxy has changed for lsnappsprofile")
		hasChange = true
	}
	if !data.Td.Equal(state.Td) {
		tflog.Debug(ctx, "td has changed for lsnappsprofile")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model (excludes ForceNew transportprotocol)
		lsnappsprofile := lsnappsprofileGetTheUpdatablePayloadFromThePlan(ctx, &data)
		// SDK v2 parity - update is via PUT with the name carried in the body
		err := r.client.UpdateUnnamedResource(service.Lsnappsprofile.Type(), &lsnappsprofile)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lsnappsprofile, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated lsnappsprofile resource")
	} else {
		tflog.Debug(ctx, "No changes detected for lsnappsprofile resource, skipping update")
	}

	// Read the updated state back
	if !r.readLsnappsprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "lsnappsprofile not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnappsprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LsnappsprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lsnappsprofile resource")

	// Named resource - delete using DeleteResource
	appsprofilename := data.Id.ValueString()
	err := r.client.DeleteResource(service.Lsnappsprofile.Type(), appsprofilename)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete lsnappsprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted lsnappsprofile resource")
}

// Helper function to read lsnappsprofile data from API.
// Returns false (without an error diagnostic) when the resource no longer exists.
func (r *LsnappsprofileResource) readLsnappsprofileFromApi(ctx context.Context, data *LsnappsprofileResourceModel, diags *diag.Diagnostics) bool {
	// Case 2: Find with single ID attribute - ID is the plain value
	appsprofilename := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Lsnappsprofile.Type(), appsprofilename)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lsnappsprofile, got error: %s", err))
		return false
	}

	lsnappsprofileSetAttrFromGet(ctx, data, getResponseData)

	return true
}
