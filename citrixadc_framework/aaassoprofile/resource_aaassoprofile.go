package aaassoprofile

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
var _ resource.Resource = &AaassoprofileResource{}
var _ resource.ResourceWithConfigure = (*AaassoprofileResource)(nil)
var _ resource.ResourceWithImportState = (*AaassoprofileResource)(nil)
var _ resource.ResourceWithValidateConfig = (*AaassoprofileResource)(nil)

func NewAaassoprofileResource() resource.Resource {
	return &AaassoprofileResource{}
}

// AaassoprofileResource defines the resource implementation.
type AaassoprofileResource struct {
	client *service.NitroClient
}

func (r *AaassoprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AaassoprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_aaassoprofile"
}

func (r *AaassoprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AaassoprofileResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data AaassoprofileResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Validate that either password or password_wo is specified
	if data.Password.IsNull() && data.PasswordWo.IsNull() {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Missing Required Attribute",
			"Either \"password\" or \"password_wo\" must be specified.",
		)
	}
}

func (r *AaassoprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config AaassoprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating aaassoprofile resource")
	// Get payload from plan (regular attributes)
	aaassoprofile := aaassoprofileGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	aaassoprofileGetThePayloadFromtheConfig(ctx, &config, &aaassoprofile)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Aaassoprofile.Type(), name_value, &aaassoprofile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create aaassoprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created aaassoprofile resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	r.readAaassoprofileFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaassoprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AaassoprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading aaassoprofile resource")

	r.readAaassoprofileFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaassoprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state AaassoprofileResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating aaassoprofile resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	// Check secret attribute password or its version tracker
	if !data.Password.Equal(state.Password) {
		tflog.Debug(ctx, fmt.Sprintf("password has changed for aaassoprofile"))
		hasChange = true
	} else if !data.PasswordWoVersion.Equal(state.PasswordWoVersion) {
		tflog.Debug(ctx, fmt.Sprintf("password_wo_version has changed for aaassoprofile"))
		hasChange = true
	}
	if !data.Username.Equal(state.Username) {
		tflog.Debug(ctx, fmt.Sprintf("username has changed for aaassoprofile"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		// Get payload from plan (regular attributes)
		aaassoprofile := aaassoprofileGetThePayloadFromthePlan(ctx, &data)
		// Add write-only attributes from config to the payload
		aaassoprofileGetThePayloadFromtheConfig(ctx, &config, &aaassoprofile)
		// Make API call
		// Named resource - use UpdateResource
		name_value := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Aaassoprofile.Type(), name_value, &aaassoprofile)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update aaassoprofile, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated aaassoprofile resource")
	} else {
		tflog.Debug(ctx, "No changes detected for aaassoprofile resource, skipping update")
	}

	// Read the updated state back
	r.readAaassoprofileFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaassoprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AaassoprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting aaassoprofile resource")
	// Named resource - delete using DeleteResource
	name_value := data.Name.ValueString()
	err := r.client.DeleteResource(service.Aaassoprofile.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete aaassoprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted aaassoprofile resource")
}

// Helper function to read aaassoprofile data from API
func (r *AaassoprofileResource) readAaassoprofileFromApi(ctx context.Context, data *AaassoprofileResourceModel, diags *diag.Diagnostics) {

	// Case 2: Find with single ID attribute - ID is the plain value
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Aaassoprofile.Type(), name_Name)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read aaassoprofile, got error: %s", err))
		return
	}

	aaassoprofileSetAttrFromGet(ctx, data, getResponseData)

}
