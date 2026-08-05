package lsntransportprofile

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
var _ resource.Resource = &LsntransportprofileResource{}
var _ resource.ResourceWithConfigure = (*LsntransportprofileResource)(nil)
var _ resource.ResourceWithImportState = (*LsntransportprofileResource)(nil)

func NewLsntransportprofileResource() resource.Resource {
	return &LsntransportprofileResource{}
}

// LsntransportprofileResource defines the resource implementation.
type LsntransportprofileResource struct {
	client *service.NitroClient
}

func (r *LsntransportprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LsntransportprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lsntransportprofile"
}

func (r *LsntransportprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LsntransportprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LsntransportprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lsntransportprofile resource")

	// Create API request body from the model
	lsntransportprofile := lsntransportprofileGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource
	transportprofilename := data.Transportprofilename.ValueString()
	_, err := r.client.AddResource(service.Lsntransportprofile.Type(), transportprofilename, &lsntransportprofile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lsntransportprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created lsntransportprofile resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(transportprofilename)

	// Read the updated state back
	if !r.readLsntransportprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "lsntransportprofile not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsntransportprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LsntransportprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lsntransportprofile resource")

	found := r.readLsntransportprofileFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *LsntransportprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state LsntransportprofileResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating lsntransportprofile resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Finrsttimeout.Equal(state.Finrsttimeout) {
		tflog.Debug(ctx, "finrsttimeout has changed for lsntransportprofile")
		hasChange = true
	}
	if !data.Groupsessionlimit.Equal(state.Groupsessionlimit) {
		tflog.Debug(ctx, "groupsessionlimit has changed for lsntransportprofile")
		hasChange = true
	}
	if !data.Portpreserveparity.Equal(state.Portpreserveparity) {
		tflog.Debug(ctx, "portpreserveparity has changed for lsntransportprofile")
		hasChange = true
	}
	if !data.Portpreserverange.Equal(state.Portpreserverange) {
		tflog.Debug(ctx, "portpreserverange has changed for lsntransportprofile")
		hasChange = true
	}
	if !data.Portquota.Equal(state.Portquota) {
		tflog.Debug(ctx, "portquota has changed for lsntransportprofile")
		hasChange = true
	}
	if !data.Sessionquota.Equal(state.Sessionquota) {
		tflog.Debug(ctx, "sessionquota has changed for lsntransportprofile")
		hasChange = true
	}
	if !data.Sessiontimeout.Equal(state.Sessiontimeout) {
		tflog.Debug(ctx, "sessiontimeout has changed for lsntransportprofile")
		hasChange = true
	}
	if !data.Stuntimeout.Equal(state.Stuntimeout) {
		tflog.Debug(ctx, "stuntimeout has changed for lsntransportprofile")
		hasChange = true
	}
	if !data.Syncheck.Equal(state.Syncheck) {
		tflog.Debug(ctx, "syncheck has changed for lsntransportprofile")
		hasChange = true
	}
	if !data.Synidletimeout.Equal(state.Synidletimeout) {
		tflog.Debug(ctx, "synidletimeout has changed for lsntransportprofile")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model, restricted to updatable fields
		lsntransportprofile := lsntransportprofileGetTheUpdatablePayloadFromThePlan(ctx, &data)
		// Named resource - use UpdateResource
		transportprofilename := data.Transportprofilename.ValueString()
		_, err := r.client.UpdateResource(service.Lsntransportprofile.Type(), transportprofilename, &lsntransportprofile)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lsntransportprofile, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated lsntransportprofile resource")
	} else {
		tflog.Debug(ctx, "No changes detected for lsntransportprofile resource, skipping update")
	}

	// Read the updated state back
	if !r.readLsntransportprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "lsntransportprofile not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsntransportprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LsntransportprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lsntransportprofile resource")

	// Named resource - delete using DeleteResource
	transportprofilename := data.Id.ValueString()
	err := r.client.DeleteResource(service.Lsntransportprofile.Type(), transportprofilename)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete lsntransportprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted lsntransportprofile resource")
}

// Helper function to read lsntransportprofile data from API
func (r *LsntransportprofileResource) readLsntransportprofileFromApi(ctx context.Context, data *LsntransportprofileResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	transportprofilename := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Lsntransportprofile.Type(), transportprofilename)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lsntransportprofile, got error: %s", err))
		return false
	}

	lsntransportprofileSetAttrFromGet(ctx, data, getResponseData)

	return true
}
