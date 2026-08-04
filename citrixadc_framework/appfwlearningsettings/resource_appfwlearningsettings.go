package appfwlearningsettings

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
var _ resource.Resource = &AppfwlearningsettingsResource{}
var _ resource.ResourceWithConfigure = (*AppfwlearningsettingsResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwlearningsettingsResource)(nil)

func NewAppfwlearningsettingsResource() resource.Resource {
	return &AppfwlearningsettingsResource{}
}

// AppfwlearningsettingsResource defines the resource implementation.
type AppfwlearningsettingsResource struct {
	client *service.NitroClient
}

func (r *AppfwlearningsettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwlearningsettingsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwlearningsettings"
}

func (r *AppfwlearningsettingsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwlearningsettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwlearningsettingsResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwlearningsettings resource")
	// Get payload from plan
	appfwlearningsettings := appfwlearningsettingsGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Unnamed resource - use UpdateUnnamedResource (NITRO only supports update/unset/get)
	err := r.client.UpdateUnnamedResource(service.Appfwlearningsettings.Type(), &appfwlearningsettings)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwlearningsettings, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created appfwlearningsettings resource")

	// Set ID for the resource before reading state (ID is the profilename, matching SDK v2 d.SetId(profilename))
	data.Id = types.StringValue(data.Profilename.ValueString())

	// Read the updated state back
	if !r.readAppfwlearningsettingsFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "appfwlearningsettings not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwlearningsettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwlearningsettingsResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwlearningsettings resource")

	found := r.readAppfwlearningsettingsFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AppfwlearningsettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AppfwlearningsettingsResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating appfwlearningsettings resource")

	// Create API request body from the model
	appfwlearningsettings := appfwlearningsettingsGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Unnamed resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Appfwlearningsettings.Type(), &appfwlearningsettings)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appfwlearningsettings, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Updated appfwlearningsettings resource")

	// Read the updated state back
	if !r.readAppfwlearningsettingsFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "appfwlearningsettings not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwlearningsettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwlearningsettingsResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwlearningsettings resource")

	// appfwlearningsettings has no NITRO delete operation (only update/unset/get).
	// Matching SDK v2 behavior (d.SetId("")), we simply drop it from Terraform state.
	tflog.Trace(ctx, "Deleted appfwlearningsettings resource from state")
}

// Helper function to read appfwlearningsettings data from API
func (r *AppfwlearningsettingsResource) readAppfwlearningsettingsFromApi(ctx context.Context, data *AppfwlearningsettingsResourceModel, diags *diag.Diagnostics) bool {

	// The ID is the profilename (matching SDK v2 d.SetId(profilename)).
	appfwlearningsettingsName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Appfwlearningsettings.Type(), appfwlearningsettingsName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwlearningsettings, got error: %s", err))
		return false
	}

	appfwlearningsettingsSetAttrFromGet(ctx, data, getResponseData)

	return true
}
