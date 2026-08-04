package appfwxmlerrorpage

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
var _ resource.Resource = &AppfwxmlerrorpageResource{}
var _ resource.ResourceWithConfigure = (*AppfwxmlerrorpageResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwxmlerrorpageResource)(nil)

func NewAppfwxmlerrorpageResource() resource.Resource {
	return &AppfwxmlerrorpageResource{}
}

// AppfwxmlerrorpageResource defines the resource implementation.
type AppfwxmlerrorpageResource struct {
	client *service.NitroClient
}

func (r *AppfwxmlerrorpageResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwxmlerrorpageResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwxmlerrorpage"
}

func (r *AppfwxmlerrorpageResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwxmlerrorpageResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwxmlerrorpageResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwxmlerrorpage resource")

	// Get payload from plan
	appfwxmlerrorpage := appfwxmlerrorpageGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Named resource imported via the "Import" NITRO action (matches SDK v2)
	err := r.client.ActOnResource(service.Appfwxmlerrorpage.Type(), &appfwxmlerrorpage, "Import")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwxmlerrorpage, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created appfwxmlerrorpage resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(data.Name.ValueString())

	// Read the updated state back
	if !r.readAppfwxmlerrorpageFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "appfwxmlerrorpage not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwxmlerrorpageResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwxmlerrorpageResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwxmlerrorpage resource")

	found := r.readAppfwxmlerrorpageFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AppfwxmlerrorpageResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AppfwxmlerrorpageResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating appfwxmlerrorpage resource")

	// All configurable attributes are ForceNew (RequiresReplace), matching the
	// SDK v2 resource which has no update operation, so no in-place NITRO call is
	// made here; the state is simply refreshed.

	// Read the updated state back
	if !r.readAppfwxmlerrorpageFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "appfwxmlerrorpage not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwxmlerrorpageResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwxmlerrorpageResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwxmlerrorpage resource")
	// Named resource - delete using DeleteResource
	name := data.Id.ValueString()
	err := r.client.DeleteResource(service.Appfwxmlerrorpage.Type(), name)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete appfwxmlerrorpage, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted appfwxmlerrorpage resource")
}

// Helper function to read appfwxmlerrorpage data from API
func (r *AppfwxmlerrorpageResource) readAppfwxmlerrorpageFromApi(ctx context.Context, data *AppfwxmlerrorpageResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value (name)
	name := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Appfwxmlerrorpage.Type(), name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwxmlerrorpage, got error: %s", err))
		return false
	}

	appfwxmlerrorpageSetAttrFromGet(ctx, data, getResponseData)

	return true
}
