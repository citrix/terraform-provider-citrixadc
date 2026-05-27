package appfwarchive_export

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AppfwarchiveExportResource{}
var _ resource.ResourceWithConfigure = (*AppfwarchiveExportResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwarchiveExportResource)(nil)

func NewAppfwarchiveExportResource() resource.Resource {
	return &AppfwarchiveExportResource{}
}

// AppfwarchiveExportResource defines the resource implementation.
type AppfwarchiveExportResource struct {
	client *service.NitroClient
}

func (r *AppfwarchiveExportResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwarchiveExportResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwarchive_export"
}

func (r *AppfwarchiveExportResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwarchiveExportResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwarchiveExportResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwarchive_export resource")
	payload := appfwarchiveExportGetThePayloadFromthePlan(ctx, &data)

	// NITRO exposes export as POST ?action=export. Use ActOnResource with
	// the case-sensitive "export" verb (lower-case e per the NITRO URL).
	err := r.client.ActOnResource(service.Appfwarchive.Type(), &payload, "export")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to export appfwarchive, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Exported appfwarchive resource")

	// ID = name (the archive that was exported); also keeps Read/Delete no-ops
	// addressable by Terraform.
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwarchiveExportResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// NITRO has no GET endpoint that reports export-state, and the appfwarchive
	// `get (all)` response carries no identifying fields anyway. Read is a
	// pure preserve-state no-op.
	var data AppfwarchiveExportResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for appfwarchive_export; NITRO has no query endpoint for export state")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwarchiveExportResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO has no update endpoint for export; every schema attribute is
	// RequiresReplace, so Terraform never invokes Update for a real change.
	var data, state AppfwarchiveExportResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for appfwarchive_export; NITRO has no update endpoint and all attributes are RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwarchiveExportResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// export is a one-shot side-effect action. There is no inverse NITRO API
	// (no "un-export" / no delete-by-export-target). Delete simply removes the
	// resource from Terraform state.
	tflog.Debug(ctx, "Delete is a no-op for appfwarchive_export; NITRO has no inverse of the export action")
}
