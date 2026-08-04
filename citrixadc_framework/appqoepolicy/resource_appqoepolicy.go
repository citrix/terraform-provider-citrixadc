package appqoepolicy

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
var _ resource.Resource = &AppqoepolicyResource{}
var _ resource.ResourceWithConfigure = (*AppqoepolicyResource)(nil)
var _ resource.ResourceWithImportState = (*AppqoepolicyResource)(nil)

func NewAppqoepolicyResource() resource.Resource {
	return &AppqoepolicyResource{}
}

// AppqoepolicyResource defines the resource implementation.
type AppqoepolicyResource struct {
	client *service.NitroClient
}

func (r *AppqoepolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppqoepolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appqoepolicy"
}

func (r *AppqoepolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppqoepolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppqoepolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appqoepolicy resource")
	// Get payload from plan
	appqoepolicy := appqoepolicyGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	appqoepolicyName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Appqoepolicy.Type(), appqoepolicyName, &appqoepolicy)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appqoepolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created appqoepolicy resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	if !r.readAppqoepolicyFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "appqoepolicy not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppqoepolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppqoepolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appqoepolicy resource")

	found := r.readAppqoepolicyFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AppqoepolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AppqoepolicyResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating appqoepolicy resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Action.Equal(state.Action) {
		tflog.Debug(ctx, "action has changed for appqoepolicy")
		hasChange = true
	}
	if !data.Rule.Equal(state.Rule) {
		tflog.Debug(ctx, "rule has changed for appqoepolicy")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model, restricted to updatable fields
		appqoepolicy := appqoepolicyGetTheUpdatablePayloadFromThePlan(ctx, &data)
		// Make API call
		// appqoepolicy is updated via an unnamed PUT (name is carried in the payload)
		err := r.client.UpdateUnnamedResource(service.Appqoepolicy.Type(), &appqoepolicy)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appqoepolicy, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated appqoepolicy resource")
	} else {
		tflog.Debug(ctx, "No changes detected for appqoepolicy resource, skipping update")
	}

	// Read the updated state back
	if !r.readAppqoepolicyFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "appqoepolicy not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppqoepolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppqoepolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appqoepolicy resource")
	// Named resource - delete using DeleteResource
	appqoepolicyName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Appqoepolicy.Type(), appqoepolicyName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete appqoepolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted appqoepolicy resource")
}

// Helper function to read appqoepolicy data from API
func (r *AppqoepolicyResource) readAppqoepolicyFromApi(ctx context.Context, data *AppqoepolicyResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	appqoepolicyName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Appqoepolicy.Type(), appqoepolicyName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appqoepolicy, got error: %s", err))
		return false
	}

	appqoepolicySetAttrFromGet(ctx, data, getResponseData)

	return true
}
