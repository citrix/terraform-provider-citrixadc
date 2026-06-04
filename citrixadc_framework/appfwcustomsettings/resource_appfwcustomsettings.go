package appfwcustomsettings

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
var _ resource.Resource = &AppfwcustomsettingsResource{}
var _ resource.ResourceWithConfigure = (*AppfwcustomsettingsResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwcustomsettingsResource)(nil)

func NewAppfwcustomsettingsResource() resource.Resource {
	return &AppfwcustomsettingsResource{}
}

// AppfwcustomsettingsResource defines the resource implementation.
type AppfwcustomsettingsResource struct {
	client *service.NitroClient
}

func (r *AppfwcustomsettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwcustomsettingsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwcustomsettings"
}

func (r *AppfwcustomsettingsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwcustomsettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwcustomsettingsResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwcustomsettings resource")
	payload := appfwcustomsettingsGetThePayloadFromthePlan(ctx, &data)

	// NITRO exposes only export as POST ?action=export. Use ActOnResource with
	// the case-sensitive "export" verb (lower-case e per the NITRO URL).
	err := r.client.ActOnResource(service.Appfwcustomsettings.Type(), &payload, "export")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to export appfwcustomsettings, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Exported appfwcustomsettings resource")

	// ID = name (the custom-settings config that was exported); keeps the
	// Read/Update/Delete no-ops addressable by Terraform.
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwcustomsettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// NITRO exposes no GET endpoint for appfwcustomsettings (only ?action=export).
	// Read is a pure preserve-state no-op.
	var data AppfwcustomsettingsResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for appfwcustomsettings; NITRO has no query endpoint for export state")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwcustomsettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO has no update endpoint for appfwcustomsettings; every schema attribute
	// is RequiresReplace, so Terraform never invokes Update for a real change.
	var data, state AppfwcustomsettingsResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for appfwcustomsettings; NITRO has no update endpoint and all attributes are RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwcustomsettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// export is a one-shot side-effect action. There is no inverse NITRO API
	// (no delete endpoint). Delete simply removes the resource from Terraform state.
	tflog.Debug(ctx, "Delete is a no-op for appfwcustomsettings; NITRO has no inverse of the export action")
}
