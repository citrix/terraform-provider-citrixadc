package aaapreauthenticationpolicy

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
var _ resource.Resource = &AaapreauthenticationpolicyResource{}
var _ resource.ResourceWithConfigure = (*AaapreauthenticationpolicyResource)(nil)
var _ resource.ResourceWithImportState = (*AaapreauthenticationpolicyResource)(nil)

func NewAaapreauthenticationpolicyResource() resource.Resource {
	return &AaapreauthenticationpolicyResource{}
}

// AaapreauthenticationpolicyResource defines the resource implementation.
type AaapreauthenticationpolicyResource struct {
	client *service.NitroClient
}

func (r *AaapreauthenticationpolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AaapreauthenticationpolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_aaapreauthenticationpolicy"
}

func (r *AaapreauthenticationpolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AaapreauthenticationpolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AaapreauthenticationpolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating aaapreauthenticationpolicy resource")
	// Get payload from plan
	aaapreauthenticationpolicy := aaapreauthenticationpolicyGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Aaapreauthenticationpolicy.Type(), name_value, &aaapreauthenticationpolicy)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create aaapreauthenticationpolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created aaapreauthenticationpolicy resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	if !r.readAaapreauthenticationpolicyFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "aaapreauthenticationpolicy not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaapreauthenticationpolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AaapreauthenticationpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading aaapreauthenticationpolicy resource")

	found := r.readAaapreauthenticationpolicyFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AaapreauthenticationpolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AaapreauthenticationpolicyResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating aaapreauthenticationpolicy resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Reqaction.Equal(state.Reqaction) {
		tflog.Debug(ctx, "reqaction has changed for aaapreauthenticationpolicy")
		hasChange = true
	}
	if !data.Rule.Equal(state.Rule) {
		tflog.Debug(ctx, "rule has changed for aaapreauthenticationpolicy")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model, restricted to NITRO-updatable fields
		aaapreauthenticationpolicy := aaapreauthenticationpolicyGetTheUpdatablePayloadFromThePlan(ctx, &data)
		// Make API call
		// NITRO update is an unnamed PUT (name is carried in the payload)
		err := r.client.UpdateUnnamedResource(service.Aaapreauthenticationpolicy.Type(), &aaapreauthenticationpolicy)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update aaapreauthenticationpolicy, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated aaapreauthenticationpolicy resource")
	} else {
		tflog.Debug(ctx, "No changes detected for aaapreauthenticationpolicy resource, skipping update")
	}

	// Read the updated state back
	if !r.readAaapreauthenticationpolicyFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "aaapreauthenticationpolicy not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaapreauthenticationpolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AaapreauthenticationpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting aaapreauthenticationpolicy resource")
	// Named resource - delete using DeleteResource
	name_value := data.Name.ValueString()
	err := r.client.DeleteResource(service.Aaapreauthenticationpolicy.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete aaapreauthenticationpolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted aaapreauthenticationpolicy resource")
}

// Helper function to read aaapreauthenticationpolicy data from API
func (r *AaapreauthenticationpolicyResource) readAaapreauthenticationpolicyFromApi(ctx context.Context, data *AaapreauthenticationpolicyResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Aaapreauthenticationpolicy.Type(), name_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read aaapreauthenticationpolicy, got error: %s", err))
		return false
	}

	aaapreauthenticationpolicySetAttrFromGet(ctx, data, getResponseData)

	return true
}
