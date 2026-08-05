package ssllogprofile

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
var _ resource.Resource = &SsllogprofileResource{}
var _ resource.ResourceWithConfigure = (*SsllogprofileResource)(nil)
var _ resource.ResourceWithImportState = (*SsllogprofileResource)(nil)

func NewSsllogprofileResource() resource.Resource {
	return &SsllogprofileResource{}
}

// SsllogprofileResource defines the resource implementation.
type SsllogprofileResource struct {
	client *service.NitroClient
}

func (r *SsllogprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SsllogprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ssllogprofile"
}

func (r *SsllogprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SsllogprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SsllogprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating ssllogprofile resource")

	ssllogprofile := ssllogprofileGetThePayloadFromtheConfig(ctx, &data)

	// Named resource - use AddResource
	ssllogprofileName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Ssllogprofile.Type(), ssllogprofileName, &ssllogprofile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create ssllogprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created ssllogprofile resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(ssllogprofileName)

	// Read the updated state back
	if !r.readSsllogprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "ssllogprofile not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SsllogprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SsllogprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading ssllogprofile resource")

	found := r.readSsllogprofileFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *SsllogprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SsllogprofileResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating ssllogprofile resource")

	// Check if there are any changes in updateable attributes
	// (name is ForceNew/RequiresReplace and never reaches Update)
	hasChange := false
	if !data.Ssllogclauth.Equal(state.Ssllogclauth) {
		tflog.Debug(ctx, "ssllogclauth has changed for ssllogprofile")
		hasChange = true
	}
	if !data.Ssllogclauthfailures.Equal(state.Ssllogclauthfailures) {
		tflog.Debug(ctx, "ssllogclauthfailures has changed for ssllogprofile")
		hasChange = true
	}
	if !data.Sslloghs.Equal(state.Sslloghs) {
		tflog.Debug(ctx, "sslloghs has changed for ssllogprofile")
		hasChange = true
	}
	if !data.Sslloghsfailures.Equal(state.Sslloghsfailures) {
		tflog.Debug(ctx, "sslloghsfailures has changed for ssllogprofile")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		ssllogprofile := ssllogprofileGetThePayloadFromtheConfig(ctx, &data)
		// Named resource - use UpdateResource
		ssllogprofileName := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Ssllogprofile.Type(), ssllogprofileName, &ssllogprofile)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update ssllogprofile, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated ssllogprofile resource")
	} else {
		tflog.Debug(ctx, "No changes detected for ssllogprofile resource, skipping update")
	}

	// Read the updated state back
	if !r.readSsllogprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "ssllogprofile not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SsllogprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SsllogprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting ssllogprofile resource")

	// Named resource - delete using DeleteResource
	ssllogprofileName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Ssllogprofile.Type(), ssllogprofileName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete ssllogprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted ssllogprofile resource")
}

// Helper function to read ssllogprofile data from API
func (r *SsllogprofileResource) readSsllogprofileFromApi(ctx context.Context, data *SsllogprofileResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value (name)
	ssllogprofileName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Ssllogprofile.Type(), ssllogprofileName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read ssllogprofile, got error: %s", err))
		return false
	}

	ssllogprofileSetAttrFromGet(ctx, data, getResponseData)

	return true
}
