package lsnrtspalgprofile

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
var _ resource.Resource = &LsnrtspalgprofileResource{}
var _ resource.ResourceWithConfigure = (*LsnrtspalgprofileResource)(nil)
var _ resource.ResourceWithImportState = (*LsnrtspalgprofileResource)(nil)

func NewLsnrtspalgprofileResource() resource.Resource {
	return &LsnrtspalgprofileResource{}
}

// LsnrtspalgprofileResource defines the resource implementation.
type LsnrtspalgprofileResource struct {
	client *service.NitroClient
}

func (r *LsnrtspalgprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LsnrtspalgprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lsnrtspalgprofile"
}

func (r *LsnrtspalgprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LsnrtspalgprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LsnrtspalgprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lsnrtspalgprofile resource")

	lsnrtspalgprofile := lsnrtspalgprofileGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	rtspalgprofilename_value := data.Rtspalgprofilename.ValueString()
	_, err := r.client.AddResource(service.Lsnrtspalgprofile.Type(), rtspalgprofilename_value, &lsnrtspalgprofile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lsnrtspalgprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created lsnrtspalgprofile resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Rtspalgprofilename.ValueString()))

	// Read the updated state back
	if !r.readLsnrtspalgprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "lsnrtspalgprofile not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnrtspalgprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LsnrtspalgprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lsnrtspalgprofile resource")

	found := r.readLsnrtspalgprofileFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *LsnrtspalgprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state LsnrtspalgprofileResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating lsnrtspalgprofile resource")

	// Check if there are any changes in updateable attributes
	// (rtspalgprofilename is ForceNew/RequiresReplace and never reaches Update)
	hasChange := false
	if !data.Rtspidletimeout.Equal(state.Rtspidletimeout) {
		tflog.Debug(ctx, "rtspidletimeout has changed for lsnrtspalgprofile")
		hasChange = true
	}
	if !data.Rtspportrange.Equal(state.Rtspportrange) {
		tflog.Debug(ctx, "rtspportrange has changed for lsnrtspalgprofile")
		hasChange = true
	}
	if !data.Rtsptransportprotocol.Equal(state.Rtsptransportprotocol) {
		tflog.Debug(ctx, "rtsptransportprotocol has changed for lsnrtspalgprofile")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		lsnrtspalgprofile := lsnrtspalgprofileGetThePayloadFromtheConfig(ctx, &data)
		// Make API call
		// Named resource - use UpdateResource
		rtspalgprofilename_value := data.Rtspalgprofilename.ValueString()
		_, err := r.client.UpdateResource(service.Lsnrtspalgprofile.Type(), rtspalgprofilename_value, &lsnrtspalgprofile)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lsnrtspalgprofile, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated lsnrtspalgprofile resource")
	} else {
		tflog.Debug(ctx, "No changes detected for lsnrtspalgprofile resource, skipping update")
	}

	// Read the updated state back
	if !r.readLsnrtspalgprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "lsnrtspalgprofile not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnrtspalgprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LsnrtspalgprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lsnrtspalgprofile resource")

	// Named resource - delete using DeleteResource
	rtspalgprofilename_value := data.Rtspalgprofilename.ValueString()
	err := r.client.DeleteResource(service.Lsnrtspalgprofile.Type(), rtspalgprofilename_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete lsnrtspalgprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted lsnrtspalgprofile resource")
}

// Helper function to read lsnrtspalgprofile data from API
func (r *LsnrtspalgprofileResource) readLsnrtspalgprofileFromApi(ctx context.Context, data *LsnrtspalgprofileResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	rtspalgprofilename_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Lsnrtspalgprofile.Type(), rtspalgprofilename_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lsnrtspalgprofile, got error: %s", err))
		return false
	}

	lsnrtspalgprofileSetAttrFromGet(ctx, data, getResponseData)

	return true
}
