package appfwarchive

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
var _ resource.Resource = &AppfwarchiveResource{}
var _ resource.ResourceWithConfigure = (*AppfwarchiveResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwarchiveResource)(nil)

func NewAppfwarchiveResource() resource.Resource {
	return &AppfwarchiveResource{}
}

// AppfwarchiveResource defines the resource implementation.
type AppfwarchiveResource struct {
	client *service.NitroClient
}

func (r *AppfwarchiveResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwarchiveResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwarchive"
}

func (r *AppfwarchiveResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwarchiveResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwarchiveResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwarchive resource")
	appfwarchive := appfwarchiveGetThePayloadFromthePlan(ctx, &data)

	// NITRO exposes appfwarchive create only via POST ?action=Import (no `add`).
	// Use ActOnResource with the case-sensitive "Import" verb.
	err := r.client.ActOnResource(service.Appfwarchive.Type(), &appfwarchive, "Import")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwarchive, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created appfwarchive resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	r.readAppfwarchiveFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwarchiveResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwarchiveResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwarchive resource")

	r.readAppfwarchiveFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwarchiveResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO exposes no update endpoint for appfwarchive (only Import, export,
	// delete, get). Every schema attribute is marked RequiresReplace, so
	// Terraform will never actually invoke Update with field changes — any plan
	// change forces destroy + recreate. This body is therefore a documented
	// no-op that preserves the prior ID and re-reads state for consistency.
	var data, state AppfwarchiveResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for appfwarchive; NITRO has no update endpoint and all attributes are RequiresReplace")

	r.readAppfwarchiveFromApi(ctx, &data, &resp.Diagnostics)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwarchiveResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwarchiveResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwarchive resource")
	// NITRO supports DELETE /appfwarchive/{name}.
	name_value := data.Name.ValueString()
	err := r.client.DeleteResource(service.Appfwarchive.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete appfwarchive, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted appfwarchive resource")
}

// Helper function to read appfwarchive data from API.
//
// NITRO's appfwarchive `get (all)` response carries only "response" and
// "_nextgenapiresource" — there is NO per-archive identifying field (no
// `name`). Therefore we cannot match a specific archive by name against the
// listing. We treat a successful, non-empty GET as confirmation that the
// resource exists, and leave the user-supplied plan/state values untouched
// (see appfwarchiveSetAttrFromGet for details).
func (r *AppfwarchiveResource) readAppfwarchiveFromApi(ctx context.Context, data *AppfwarchiveResourceModel, diags *diag.Diagnostics) {

	findParams := service.FindParams{
		ResourceType:             service.Appfwarchive.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err := r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwarchive, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "appfwarchive returned empty array")
		return
	}

	// The response carries no identifying fields; pass the first entry so that
	// any future response fields can be honoured, but preserve plan/state values.
	appfwarchiveSetAttrFromGet(ctx, data, dataArr[0])
}
