package cloudallowedngsticketprofile

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &CloudallowedngsticketprofileResource{}
var _ resource.ResourceWithConfigure = (*CloudallowedngsticketprofileResource)(nil)
var _ resource.ResourceWithImportState = (*CloudallowedngsticketprofileResource)(nil)

func NewCloudallowedngsticketprofileResource() resource.Resource {
	return &CloudallowedngsticketprofileResource{}
}

// CloudallowedngsticketprofileResource defines the resource implementation.
type CloudallowedngsticketprofileResource struct {
	client *service.NitroClient
}

func (r *CloudallowedngsticketprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *CloudallowedngsticketprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cloudallowedngsticketprofile"
}

func (r *CloudallowedngsticketprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *CloudallowedngsticketprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data CloudallowedngsticketprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating cloudallowedngsticketprofile resource")
	cloudallowedngsticketprofile := cloudallowedngsticketprofileGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Cloudallowedngsticketprofile.Type(), name_value, &cloudallowedngsticketprofile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create cloudallowedngsticketprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created cloudallowedngsticketprofile resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	r.readCloudallowedngsticketprofileFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CloudallowedngsticketprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data CloudallowedngsticketprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading cloudallowedngsticketprofile resource")

	r.readCloudallowedngsticketprofileFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CloudallowedngsticketprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state CloudallowedngsticketprofileResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating cloudallowedngsticketprofile resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Creator.Equal(state.Creator) {
		tflog.Debug(ctx, fmt.Sprintf("creator has changed for cloudallowedngsticketprofile"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		cloudallowedngsticketprofile := cloudallowedngsticketprofileGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Named resource - use UpdateResource
		name_value := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Cloudallowedngsticketprofile.Type(), name_value, &cloudallowedngsticketprofile)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update cloudallowedngsticketprofile, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated cloudallowedngsticketprofile resource")
	} else {
		tflog.Debug(ctx, "No changes detected for cloudallowedngsticketprofile resource, skipping update")
	}

	// Read the updated state back
	r.readCloudallowedngsticketprofileFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CloudallowedngsticketprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data CloudallowedngsticketprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting cloudallowedngsticketprofile resource")
	// Named resource - delete using DeleteResource
	name_value := data.Name.ValueString()
	err := r.client.DeleteResource(service.Cloudallowedngsticketprofile.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete cloudallowedngsticketprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted cloudallowedngsticketprofile resource")
}

// Helper function to read cloudallowedngsticketprofile data from API
func (r *CloudallowedngsticketprofileResource) readCloudallowedngsticketprofileFromApi(ctx context.Context, data *CloudallowedngsticketprofileResourceModel, diags *diag.Diagnostics) {

	// Case 2: Find with single ID attribute - ID is the plain value
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Cloudallowedngsticketprofile.Type(), name_Name)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read cloudallowedngsticketprofile, got error: %s", err))
		return
	}

	cloudallowedngsticketprofileSetAttrFromGet(ctx, data, getResponseData)

}
