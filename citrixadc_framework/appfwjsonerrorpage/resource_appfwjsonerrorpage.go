package appfwjsonerrorpage

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
var _ resource.Resource = &AppfwjsonerrorpageResource{}
var _ resource.ResourceWithConfigure = (*AppfwjsonerrorpageResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwjsonerrorpageResource)(nil)

func NewAppfwjsonerrorpageResource() resource.Resource {
	return &AppfwjsonerrorpageResource{}
}

// AppfwjsonerrorpageResource defines the resource implementation.
type AppfwjsonerrorpageResource struct {
	client *service.NitroClient
}

func (r *AppfwjsonerrorpageResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwjsonerrorpageResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwjsonerrorpage"
}

func (r *AppfwjsonerrorpageResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwjsonerrorpageResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwjsonerrorpageResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwjsonerrorpage resource")
	appfwjsonerrorpage := appfwjsonerrorpageGetThePayloadFromtheConfig(ctx, &data)

	// NITRO exposes appfwjsonerrorpage create only via POST ?action=Import (no
	// `add`). Match the SDK v2 resource which used the case-sensitive "Import" verb.
	err := r.client.ActOnResource(service.Appfwjsonerrorpage.Type(), &appfwjsonerrorpage, "Import")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwjsonerrorpage, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created appfwjsonerrorpage resource")

	// Set ID for the resource before reading state (SDK v2: d.SetId(name))
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	if !r.readAppfwjsonerrorpageFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "appfwjsonerrorpage not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwjsonerrorpageResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwjsonerrorpageResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwjsonerrorpage resource")

	found := r.readAppfwjsonerrorpageFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AppfwjsonerrorpageResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO exposes no field-level update for appfwjsonerrorpage (only Import,
	// delete, get and a `change` that takes just the name). Every schema
	// attribute is marked RequiresReplace, so Terraform never invokes Update
	// with field changes — any plan change forces destroy + recreate. This body
	// is therefore a documented no-op that preserves the prior ID and re-reads
	// state for consistency.
	var data, state AppfwjsonerrorpageResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for appfwjsonerrorpage; NITRO has no field update endpoint and all attributes are RequiresReplace")

	if !r.readAppfwjsonerrorpageFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "appfwjsonerrorpage not found immediately after update")
		}
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwjsonerrorpageResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwjsonerrorpageResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwjsonerrorpage resource")
	// SDK v2 deleted via DELETE /appfwjsonerrorpage/{name}.
	name_value := data.Name.ValueString()
	err := r.client.DeleteResource(service.Appfwjsonerrorpage.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete appfwjsonerrorpage, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted appfwjsonerrorpage resource")
}

// Helper function to read appfwjsonerrorpage data from API.
//
// The SDK v2 resource read the object by name (FindResource(name)) and only
// echoed `name` back into state. NITRO's get response carries name/response/src
// but never `comment` or `overwrite`, so appfwjsonerrorpageSetAttrFromGet
// preserves the user-supplied inputs to avoid a perpetual diff / "inconsistent
// result after apply".
func (r *AppfwjsonerrorpageResource) readAppfwjsonerrorpageFromApi(ctx context.Context, data *AppfwjsonerrorpageResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value (name)
	appfwjsonerrorpage_Name := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Appfwjsonerrorpage.Type(), appfwjsonerrorpage_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwjsonerrorpage, got error: %s", err))
		return false
	}

	appfwjsonerrorpageSetAttrFromGet(ctx, data, getResponseData)

	return true
}
