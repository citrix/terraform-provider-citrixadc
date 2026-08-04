package appfwprofile

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
var _ resource.Resource = &AppfwprofileResource{}
var _ resource.ResourceWithConfigure = (*AppfwprofileResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwprofileResource)(nil)

func NewAppfwprofileResource() resource.Resource {
	return &AppfwprofileResource{}
}

// AppfwprofileResource defines the resource implementation.
type AppfwprofileResource struct {
	client *service.NitroClient
}

func (r *AppfwprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile"
}

func (r *AppfwprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwprofile resource")

	// Get payload from plan
	appfwprofile := appfwprofileGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	appfwprofileName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Appfwprofile.Type(), appfwprofileName, &appfwprofile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created appfwprofile resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(appfwprofileName)

	// Read the updated state back
	if !r.readAppfwprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "appfwprofile not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwprofile resource")

	found := r.readAppfwprofileFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AppfwprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AppfwprofileResourceModel

	// Read Terraform prior state to preserve ID and detect changes
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating appfwprofile resource")

	// Build payload restricted to updatable fields that actually changed
	appfwprofile, hasChange := appfwprofileGetTheUpdatablePayloadFromThePlan(ctx, &data, &state)

	if hasChange {
		// Named resource - use UpdateResource
		appfwprofileName := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Appfwprofile.Type(), appfwprofileName, &appfwprofile)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appfwprofile, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated appfwprofile resource")
	} else {
		tflog.Debug(ctx, "No changes detected for appfwprofile resource, skipping update")
	}

	// Read the updated state back
	if !r.readAppfwprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "appfwprofile not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwprofile resource")
	// Named resource - delete using DeleteResource
	appfwprofileName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Appfwprofile.Type(), appfwprofileName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete appfwprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted appfwprofile resource")
}

// Helper function to read appfwprofile data from API
func (r *AppfwprofileResource) readAppfwprofileFromApi(ctx context.Context, data *AppfwprofileResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain profile name
	appfwprofileName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Appfwprofile.Type(), appfwprofileName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile, got error: %s", err))
		return false
	}

	appfwprofileSetAttrFromGet(ctx, data, getResponseData)

	return true
}
