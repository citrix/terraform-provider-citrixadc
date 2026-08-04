package appflowpolicylabel

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
var _ resource.Resource = &AppflowpolicylabelResource{}
var _ resource.ResourceWithConfigure = (*AppflowpolicylabelResource)(nil)
var _ resource.ResourceWithImportState = (*AppflowpolicylabelResource)(nil)

func NewAppflowpolicylabelResource() resource.Resource {
	return &AppflowpolicylabelResource{}
}

// AppflowpolicylabelResource defines the resource implementation.
type AppflowpolicylabelResource struct {
	client *service.NitroClient
}

func (r *AppflowpolicylabelResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppflowpolicylabelResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appflowpolicylabel"
}

func (r *AppflowpolicylabelResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppflowpolicylabelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppflowpolicylabelResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appflowpolicylabel resource")
	// Get payload from plan
	appflowpolicylabel := appflowpolicylabelGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	labelname_value := data.Labelname.ValueString()
	_, err := r.client.AddResource(service.Appflowpolicylabel.Type(), labelname_value, &appflowpolicylabel)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appflowpolicylabel, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created appflowpolicylabel resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Labelname.ValueString()))

	// Read the updated state back
	if !r.readAppflowpolicylabelFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "appflowpolicylabel not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppflowpolicylabelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppflowpolicylabelResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appflowpolicylabel resource")

	found := r.readAppflowpolicylabelFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AppflowpolicylabelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AppflowpolicylabelResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// appflowpolicylabel has no NITRO-updatable attributes (labelname and
	// policylabeltype are both ForceNew/RequiresReplace, and NITRO exposes no
	// update operation for this resource). Any configuration change triggers a
	// replace instead of an update, so here we simply re-read current state.
	tflog.Debug(ctx, "Updating appflowpolicylabel resource - no updatable attributes, re-reading state")

	// Read the updated state back
	if !r.readAppflowpolicylabelFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "appflowpolicylabel not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppflowpolicylabelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppflowpolicylabelResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appflowpolicylabel resource")
	// Named resource - delete using DeleteResource
	labelname_value := data.Labelname.ValueString()
	err := r.client.DeleteResource(service.Appflowpolicylabel.Type(), labelname_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete appflowpolicylabel, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted appflowpolicylabel resource")
}

// Helper function to read appflowpolicylabel data from API
func (r *AppflowpolicylabelResource) readAppflowpolicylabelFromApi(ctx context.Context, data *AppflowpolicylabelResourceModel, diags *diag.Diagnostics) bool {

	// Find with single ID attribute - ID is the plain value
	labelname_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Appflowpolicylabel.Type(), labelname_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appflowpolicylabel, got error: %s", err))
		return false
	}

	appflowpolicylabelSetAttrFromGet(ctx, data, getResponseData)

	return true
}
