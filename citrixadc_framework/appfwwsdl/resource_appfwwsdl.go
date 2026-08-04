package appfwwsdl

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
var _ resource.Resource = &AppfwwsdlResource{}
var _ resource.ResourceWithConfigure = (*AppfwwsdlResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwwsdlResource)(nil)

func NewAppfwwsdlResource() resource.Resource {
	return &AppfwwsdlResource{}
}

// AppfwwsdlResource defines the resource implementation.
type AppfwwsdlResource struct {
	client *service.NitroClient
}

func (r *AppfwwsdlResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwwsdlResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwwsdl"
}

func (r *AppfwwsdlResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwwsdlResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwwsdlResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwwsdl resource")
	appfwwsdl := appfwwsdlGetThePayloadFromthePlan(ctx, &data)

	// NITRO exposes appfwwsdl create only via POST ?action=Import (no `add`).
	// Use ActOnResource with the case-sensitive "Import" verb (matches SDK v2).
	err := r.client.ActOnResource(service.Appfwwsdl.Type(), &appfwwsdl, "Import")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwwsdl, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created appfwwsdl resource")

	// Set ID for the resource before reading state (matches SDK v2 d.SetId(name))
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	if !r.readAppfwwsdlFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "appfwwsdl not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwwsdlResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwwsdlResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwwsdl resource")

	found := r.readAppfwwsdlFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AppfwwsdlResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO exposes no update endpoint for appfwwsdl that accepts the Import
	// inputs (src/overwrite/comment). Every schema attribute is marked
	// RequiresReplace (SDK v2 had every attribute ForceNew and no update func),
	// so Terraform will never actually invoke Update with field changes — any
	// plan change forces destroy + recreate. This body is therefore a documented
	// no-op that preserves the prior ID and re-reads state for consistency.
	var data, state AppfwwsdlResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for appfwwsdl; NITRO has no compatible update endpoint and all attributes are RequiresReplace")

	if !r.readAppfwwsdlFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "appfwwsdl not found immediately after update")
		}
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwwsdlResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwwsdlResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwwsdl resource")
	// Named resource - delete using DeleteResource (matches SDK v2)
	name_value := data.Id.ValueString()
	err := r.client.DeleteResource(service.Appfwwsdl.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete appfwwsdl, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted appfwwsdl resource")
}

// Helper function to read appfwwsdl data from API.
//
// NITRO supports `get` by name (URL: /appfwwsdl/<name>) but the response payload
// only echoes back `name`, `response`, and `_nextgenapiresource`. The
// user-supplied Import inputs `comment`, `overwrite`, and `src` are NEVER
// returned, so touching them would null them on every Read and cause a perpetual
// diff. appfwwsdlSetAttrFromGet preserves the existing plan/state values for
// those inputs (mirrors the SDK v2 read, which only set `name`).
func (r *AppfwwsdlResource) readAppfwwsdlFromApi(ctx context.Context, data *AppfwwsdlResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Appfwwsdl.Type(), name_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwwsdl, got error: %s", err))
		return false
	}

	appfwwsdlSetAttrFromGet(ctx, data, getResponseData)

	return true
}
