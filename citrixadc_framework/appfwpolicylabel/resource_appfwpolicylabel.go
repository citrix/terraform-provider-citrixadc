package appfwpolicylabel

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
var _ resource.Resource = &AppfwpolicylabelResource{}
var _ resource.ResourceWithConfigure = (*AppfwpolicylabelResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwpolicylabelResource)(nil)

func NewAppfwpolicylabelResource() resource.Resource {
	return &AppfwpolicylabelResource{}
}

// AppfwpolicylabelResource defines the resource implementation.
type AppfwpolicylabelResource struct {
	client *service.NitroClient
}

func (r *AppfwpolicylabelResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwpolicylabelResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwpolicylabel"
}

func (r *AppfwpolicylabelResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwpolicylabelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwpolicylabelResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwpolicylabel resource")
	// Get payload from plan
	appfwpolicylabel := appfwpolicylabelGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	labelname_value := data.Labelname.ValueString()
	_, err := r.client.AddResource(service.Appfwpolicylabel.Type(), labelname_value, &appfwpolicylabel)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwpolicylabel, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created appfwpolicylabel resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Labelname.ValueString()))

	// Read the updated state back
	if !r.readAppfwpolicylabelFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "appfwpolicylabel not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwpolicylabelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwpolicylabelResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwpolicylabel resource")

	found := r.readAppfwpolicylabelFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AppfwpolicylabelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AppfwpolicylabelResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating appfwpolicylabel resource")

	// appfwpolicylabel has no NITRO update operation (add/delete/get/rename only) and all
	// configurable attributes carry RequiresReplace, so no in-place update call is issued.

	// Read the updated state back
	if !r.readAppfwpolicylabelFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "appfwpolicylabel not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwpolicylabelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwpolicylabelResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwpolicylabel resource")
	// Named resource - delete using DeleteResource
	labelname_value := data.Labelname.ValueString()
	err := r.client.DeleteResource(service.Appfwpolicylabel.Type(), labelname_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete appfwpolicylabel, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted appfwpolicylabel resource")
}

// Helper function to read appfwpolicylabel data from API
func (r *AppfwpolicylabelResource) readAppfwpolicylabelFromApi(ctx context.Context, data *AppfwpolicylabelResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	labelname_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Appfwpolicylabel.Type(), labelname_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwpolicylabel, got error: %s", err))
		return false
	}

	appfwpolicylabelSetAttrFromGet(ctx, data, getResponseData)

	return true
}
