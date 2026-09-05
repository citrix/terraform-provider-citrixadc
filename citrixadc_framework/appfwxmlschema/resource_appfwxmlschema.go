package appfwxmlschema

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
var _ resource.Resource = &AppfwxmlschemaResource{}
var _ resource.ResourceWithConfigure = (*AppfwxmlschemaResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwxmlschemaResource)(nil)

func NewAppfwxmlschemaResource() resource.Resource {
	return &AppfwxmlschemaResource{}
}

// AppfwxmlschemaResource defines the resource implementation.
type AppfwxmlschemaResource struct {
	client *service.NitroClient
}

func (r *AppfwxmlschemaResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwxmlschemaResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwxmlschema"
}

func (r *AppfwxmlschemaResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwxmlschemaResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwxmlschemaResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwxmlschema resource")

	appfwxmlschema := appfwxmlschemaGetThePayloadFromtheConfig(ctx, &data)

	// Named resource - appfwxmlschema is created via the NITRO "Import" action
	// (matches the SDK v2 client.ActOnResource(..., "Import") behaviour).
	name := data.Name.ValueString()
	err := r.client.ActOnResource(service.Appfwxmlschema.Type(), &appfwxmlschema, "Import")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwxmlschema, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created appfwxmlschema resource")

	// Set ID for the resource before reading state (matches SDK v2 d.SetId(name))
	data.Id = types.StringValue(name)

	// Read the created resource back
	if !r.readAppfwxmlschemaFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "appfwxmlschema not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwxmlschemaResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwxmlschemaResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwxmlschema resource")

	found := r.readAppfwxmlschemaFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AppfwxmlschemaResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AppfwxmlschemaResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating appfwxmlschema resource")

	// All appfwxmlschema attributes are ForceNew (RequiresReplace) and NITRO exposes no
	// in-place update operation, so there is nothing to update here - just refresh state.
	if !r.readAppfwxmlschemaFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "appfwxmlschema not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwxmlschemaResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwxmlschemaResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwxmlschema resource")
	// Named resource - delete using DeleteResource (matches SDK v2 behaviour)
	name := data.Id.ValueString()
	err := r.client.DeleteResource(service.Appfwxmlschema.Type(), name)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete appfwxmlschema, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted appfwxmlschema resource")
}

// Helper function to read appfwxmlschema data from API
func (r *AppfwxmlschemaResource) readAppfwxmlschemaFromApi(ctx context.Context, data *AppfwxmlschemaResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value (the name)
	appfwxmlschema_Name := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Appfwxmlschema.Type(), appfwxmlschema_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwxmlschema, got error: %s", err))
		return false
	}

	appfwxmlschemaSetAttrFromGet(ctx, data, getResponseData)

	return true
}
