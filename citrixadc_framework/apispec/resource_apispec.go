package apispec

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
var _ resource.Resource = &ApispecResource{}
var _ resource.ResourceWithConfigure = (*ApispecResource)(nil)
var _ resource.ResourceWithImportState = (*ApispecResource)(nil)

func NewApispecResource() resource.Resource {
	return &ApispecResource{}
}

// ApispecResource defines the resource implementation.
type ApispecResource struct {
	client *service.NitroClient
}

func (r *ApispecResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ApispecResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_apispec"
}

func (r *ApispecResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ApispecResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ApispecResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating apispec resource")
	apispec := apispecGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Apispec.Type(), name_value, &apispec)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create apispec, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created apispec resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	if !r.readApispecFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "apispec not found immediately after create/update")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ApispecResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ApispecResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading apispec resource")

	found := r.readApispecFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *ApispecResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state ApispecResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating apispec resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.File.Equal(state.File) {
		tflog.Debug(ctx, fmt.Sprintf("file has changed for apispec"))
		hasChange = true
	}
	if !data.Skipvalidation.Equal(state.Skipvalidation) {
		tflog.Debug(ctx, fmt.Sprintf("skipvalidation has changed for apispec"))
		hasChange = true
	}
	if !data.Type.Equal(state.Type) {
		tflog.Debug(ctx, fmt.Sprintf("type has changed for apispec"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model.
		// NITRO `change` action does NOT accept `encrypted`, so use a
		// separate payload builder that omits it.
		apispec := apispecGetTheUpdatePayloadFromthePlan(ctx, &data)
		// Make API call
		// NITRO exposes update via POST /apispec?action=update (the
		// "change" action), not standard PUT. Use ActOnResource.
		err := r.client.ActOnResource(service.Apispec.Type(), &apispec, "update")
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update apispec, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated apispec resource")
	} else {
		tflog.Debug(ctx, "No changes detected for apispec resource, skipping update")
	}

	// Read the updated state back
	if !r.readApispecFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "apispec not found immediately after create/update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ApispecResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ApispecResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting apispec resource")
	// Named resource - delete using DeleteResource
	name_value := data.Name.ValueString()
	err := r.client.DeleteResource(service.Apispec.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete apispec, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted apispec resource")
}

// Helper function to read apispec data from API
func (r *ApispecResource) readApispecFromApi(ctx context.Context, data *ApispecResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Apispec.Type(), name_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read apispec, got error: %s", err))
		return false
	}

	apispecSetAttrFromGet(ctx, data, getResponseData)

	return true
}
