package appfwprotofile

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AppfwprotofileResource{}
var _ resource.ResourceWithConfigure = (*AppfwprotofileResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwprotofileResource)(nil)

func NewAppfwprotofileResource() resource.Resource {
	return &AppfwprotofileResource{}
}

// AppfwprotofileResource defines the resource implementation.
type AppfwprotofileResource struct {
	client *service.NitroClient
}

func (r *AppfwprotofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwprotofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprotofile"
}

func (r *AppfwprotofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwprotofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwprotofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwprotofile resource")
	appfwprotofile := appfwprotofileGetThePayloadFromthePlan(ctx, &data)

	// NITRO exposes appfwprotofile create only via POST ?action=Import (no `add`).
	// Use ActOnResource with the case-sensitive "Import" verb.
	err := r.client.ActOnResource(service.Appfwprotofile.Type(), &appfwprotofile, "Import")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwprotofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created appfwprotofile resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	r.readAppfwprotofileFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprotofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwprotofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwprotofile resource")

	r.readAppfwprotofileFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprotofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO exposes no update endpoint for appfwprotofile that accepts the
	// Import inputs (src/overwrite/comment). The `change` action only accepts
	// `name` and is not a true update. Every schema attribute is marked
	// RequiresReplace, so Terraform will never actually invoke Update with
	// field changes — any plan change forces destroy + recreate. This body is
	// therefore a documented no-op that preserves the prior ID and re-reads
	// state for consistency.
	var data, state AppfwprotofileResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for appfwprotofile; NITRO has no compatible update endpoint and all attributes are RequiresReplace")

	r.readAppfwprotofileFromApi(ctx, &data, &resp.Diagnostics)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprotofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwprotofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwprotofile resource")
	// Named resource - delete using DeleteResource
	name_value := data.Name.ValueString()
	err := r.client.DeleteResource(service.Appfwprotofile.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete appfwprotofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted appfwprotofile resource")
}

// Helper function to read appfwprotofile data from API.
//
// NITRO supports `get` by name (URL: /appfwprotofile/<name>) but the response
// payload only echoes back `name`, `response`, `src`, and `_nextgenapiresource`.
// The user-supplied write-only inputs `comment` and `overwrite` are NEVER
// returned, so touching them here would null them on every Read and cause a
// perpetual diff. Preserve the existing plan/state values for those inputs.
func (r *AppfwprotofileResource) readAppfwprotofileFromApi(ctx context.Context, data *AppfwprotofileResourceModel, diags *diag.Diagnostics) {

	// Case 2: Find with single ID attribute - ID is the plain value
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Appfwprotofile.Type(), name_Name)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwprotofile, got error: %s", err))
		return
	}

	appfwprotofileSetAttrFromGet(ctx, data, getResponseData)

}
