package aaauser

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
var _ resource.Resource = &AaauserResource{}
var _ resource.ResourceWithConfigure = (*AaauserResource)(nil)
var _ resource.ResourceWithImportState = (*AaauserResource)(nil)

func NewAaauserResource() resource.Resource {
	return &AaauserResource{}
}

// AaauserResource defines the resource implementation.
type AaauserResource struct {
	client *service.NitroClient
}

func (r *AaauserResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AaauserResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_aaauser"
}

func (r *AaauserResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AaauserResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config AaauserResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating aaauser resource")
	// Get payload from plan (regular attributes)
	aaauser := aaauserGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	aaauserGetThePayloadFromtheConfig(ctx, &config, &aaauser)

	// Make API call
	// Named resource - use AddResource
	username_value := data.Username.ValueString()
	_, err := r.client.AddResource(service.Aaauser.Type(), username_value, &aaauser)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create aaauser, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created aaauser resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Username.ValueString()))

	// Read the updated state back
	r.readAaauserFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaauserResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AaauserResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading aaauser resource")

	r.readAaauserFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaauserResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state AaauserResourceModel

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

	tflog.Debug(ctx, "Updating aaauser resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	// Check secret attribute password or its version tracker
	if !data.Password.Equal(state.Password) {
		tflog.Debug(ctx, fmt.Sprintf("password has changed for aaauser"))
		hasChange = true
	} else if !data.PasswordWoVersion.Equal(state.PasswordWoVersion) {
		tflog.Debug(ctx, fmt.Sprintf("password_wo_version has changed for aaauser"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		// Get payload from plan (regular attributes)
		aaauser := aaauserGetThePayloadFromthePlan(ctx, &data)
		// Add write-only attributes from config to the payload
		aaauserGetThePayloadFromtheConfig(ctx, &config, &aaauser)
		// Make API call
		// Named resource - use UpdateResource
		username_value := data.Username.ValueString()
		_, err := r.client.UpdateResource(service.Aaauser.Type(), username_value, &aaauser)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update aaauser, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated aaauser resource")
	} else {
		tflog.Debug(ctx, "No changes detected for aaauser resource, skipping update")
	}

	// Read the updated state back
	r.readAaauserFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaauserResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AaauserResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting aaauser resource")
	// Named resource - delete using DeleteResource
	username_value := data.Username.ValueString()
	err := r.client.DeleteResource(service.Aaauser.Type(), username_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete aaauser, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted aaauser resource")
}

// Helper function to read aaauser data from API
func (r *AaauserResource) readAaauserFromApi(ctx context.Context, data *AaauserResourceModel, diags *diag.Diagnostics) {

	// Case 2: Find with single ID attribute - ID is the plain value
	username_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Aaauser.Type(), username_Name)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read aaauser, got error: %s", err))
		return
	}

	aaauserSetAttrFromGet(ctx, data, getResponseData)

}
