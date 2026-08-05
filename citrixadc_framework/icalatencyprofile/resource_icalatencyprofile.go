package icalatencyprofile

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
var _ resource.Resource = &IcalatencyprofileResource{}
var _ resource.ResourceWithConfigure = (*IcalatencyprofileResource)(nil)
var _ resource.ResourceWithImportState = (*IcalatencyprofileResource)(nil)

func NewIcalatencyprofileResource() resource.Resource {
	return &IcalatencyprofileResource{}
}

// IcalatencyprofileResource defines the resource implementation.
type IcalatencyprofileResource struct {
	client *service.NitroClient
}

func (r *IcalatencyprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *IcalatencyprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_icalatencyprofile"
}

func (r *IcalatencyprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *IcalatencyprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data IcalatencyprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating icalatencyprofile resource")

	icalatencyprofile := icalatencyprofileGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Icalatencyprofile.Type(), name_value, &icalatencyprofile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create icalatencyprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created icalatencyprofile resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(name_value)

	// Read the updated state back
	if !r.readIcalatencyprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "icalatencyprofile not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IcalatencyprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data IcalatencyprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading icalatencyprofile resource")

	found := r.readIcalatencyprofileFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *IcalatencyprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state IcalatencyprofileResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating icalatencyprofile resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.L7latencymaxnotifycount.Equal(state.L7latencymaxnotifycount) {
		tflog.Debug(ctx, "l7latencymaxnotifycount has changed for icalatencyprofile")
		hasChange = true
	}
	if !data.L7latencymonitoring.Equal(state.L7latencymonitoring) {
		tflog.Debug(ctx, "l7latencymonitoring has changed for icalatencyprofile")
		hasChange = true
	}
	if !data.L7latencynotifyinterval.Equal(state.L7latencynotifyinterval) {
		tflog.Debug(ctx, "l7latencynotifyinterval has changed for icalatencyprofile")
		hasChange = true
	}
	if !data.L7latencythresholdfactor.Equal(state.L7latencythresholdfactor) {
		tflog.Debug(ctx, "l7latencythresholdfactor has changed for icalatencyprofile")
		hasChange = true
	}
	if !data.L7latencywaittime.Equal(state.L7latencywaittime) {
		tflog.Debug(ctx, "l7latencywaittime has changed for icalatencyprofile")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		icalatencyprofile := icalatencyprofileGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Named resource - use UpdateResource
		name_value := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Icalatencyprofile.Type(), name_value, &icalatencyprofile)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update icalatencyprofile, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated icalatencyprofile resource")
	} else {
		tflog.Debug(ctx, "No changes detected for icalatencyprofile resource, skipping update")
	}

	// Read the updated state back
	if !r.readIcalatencyprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "icalatencyprofile not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IcalatencyprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data IcalatencyprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting icalatencyprofile resource")
	// Named resource - delete using DeleteResource
	name_value := data.Name.ValueString()
	err := r.client.DeleteResource(service.Icalatencyprofile.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete icalatencyprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted icalatencyprofile resource")
}

// Helper function to read icalatencyprofile data from API.
// Returns false when the resource no longer exists on the ADC.
func (r *IcalatencyprofileResource) readIcalatencyprofileFromApi(ctx context.Context, data *IcalatencyprofileResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain name value
	name_Name := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Icalatencyprofile.Type(), name_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read icalatencyprofile, got error: %s", err))
		return false
	}

	icalatencyprofileSetAttrFromGet(ctx, data, getResponseData)

	return true
}
