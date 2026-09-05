package appalgparam

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
var _ resource.Resource = &AppalgparamResource{}
var _ resource.ResourceWithConfigure = (*AppalgparamResource)(nil)
var _ resource.ResourceWithImportState = (*AppalgparamResource)(nil)

func NewAppalgparamResource() resource.Resource {
	return &AppalgparamResource{}
}

// AppalgparamResource defines the resource implementation.
type AppalgparamResource struct {
	client *service.NitroClient
}

func (r *AppalgparamResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppalgparamResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appalgparam"
}

func (r *AppalgparamResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppalgparamResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppalgparamResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appalgparam resource")

	appalgparam := appalgparamGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Singleton resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Appalgparam.Type(), &appalgparam)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appalgparam, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created appalgparam resource")

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("appalgparam-config")

	// Read the updated state back
	if !r.readAppalgparamFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "appalgparam not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppalgparamResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppalgparamResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appalgparam resource")

	found := r.readAppalgparamFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AppalgparamResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AppalgparamResourceModel

	// Read Terraform prior state to preserve ID and detect changes
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating appalgparam resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Pptpgreidletimeout.Equal(state.Pptpgreidletimeout) {
		tflog.Debug(ctx, "pptpgreidletimeout has changed for appalgparam")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		appalgparam := appalgparamGetThePayloadFromtheConfig(ctx, &data)

		// Make API call
		// Singleton resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Appalgparam.Type(), &appalgparam)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appalgparam, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated appalgparam resource")
	} else {
		tflog.Debug(ctx, "No changes detected for appalgparam resource, skipping update")
	}

	// Read the updated state back
	if !r.readAppalgparamFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "appalgparam not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppalgparamResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppalgparamResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appalgparam resource")

	// For appalgparam, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted appalgparam resource from state")
}

// Helper function to read appalgparam data from API
func (r *AppalgparamResource) readAppalgparamFromApi(ctx context.Context, data *AppalgparamResourceModel, diags *diag.Diagnostics) bool {

	// Case 1: Simple find without ID (singleton resource)
	getResponseData, err := r.client.FindResource(service.Appalgparam.Type(), "")
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appalgparam, got error: %s", err))
		return false
	}

	appalgparamSetAttrFromGet(ctx, data, getResponseData)

	return true
}
