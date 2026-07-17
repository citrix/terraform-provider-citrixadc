package policyurlset

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/policy"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &PolicyurlsetExportResource{}
var _ resource.ResourceWithConfigure = (*PolicyurlsetExportResource)(nil)
var _ resource.ResourceWithImportState = (*PolicyurlsetExportResource)(nil)

func NewPolicyurlsetExportResource() resource.Resource {
	return &PolicyurlsetExportResource{}
}

// PolicyurlsetExportResource defines the resource implementation.
type PolicyurlsetExportResource struct {
	client *service.NitroClient
}

// PolicyurlsetExportResourceModel describes the resource data model.
//
// This resource models the NITRO policyurlset `?action=export` action. export
// is a one-shot side-effect action with no GET endpoint and no inverse API, so
// Read/Update/Delete are no-ops. The export payload carries `name` and `url`.
type PolicyurlsetExportResourceModel struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
	Url  types.String `tfsdk:"url"`
}

func (r *PolicyurlsetExportResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *PolicyurlsetExportResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_policyurlset_export"
}

func (r *PolicyurlsetExportResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *PolicyurlsetExportResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the policyurlset_export resource.",
			},
			// NITRO export payload marks `name` as mandatory.
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Unique name of the url set. Maximum length: 127.",
			},
			// CLI marks `-url` mandatory for export (Pattern 8: tfdata/NITRO
			// under-constrained; CLI is authoritative).
			"url": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "URL (protocol, host, path and file name) to which the CSV file will be exported. HTTP, HTTPS and FTP protocols are supported. Maximum length: 2047.",
			},
		},
	}
}

func (r *PolicyurlsetExportResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data PolicyurlsetExportResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating policyurlset_export resource")
	payload := policyurlset_exportGetThePayloadFromthePlan(ctx, &data)

	// NITRO exposes export as POST ?action=export.
	err := r.client.ActOnResource(service.Policyurlset.Type(), &payload, "export")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to export policyurlset, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Exported policyurlset resource")

	data.Id = types.StringValue(fmt.Sprintf("policyurlset_export-%v", data.Name.ValueString()))

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PolicyurlsetExportResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// export is a one-shot action. NITRO has no GET endpoint that reports
	// export-state, so Read is a pure preserve-state no-op.
	var data PolicyurlsetExportResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for policyurlset_export; NITRO has no query endpoint for export state")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PolicyurlsetExportResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO has no update endpoint for export; every schema attribute is
	// RequiresReplace, so Terraform never invokes Update for a real change.
	var data, state PolicyurlsetExportResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for policyurlset_export; all attributes are RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PolicyurlsetExportResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// export is a one-shot side-effect action. There is no inverse NITRO API.
	// Delete simply removes the resource from Terraform state.
	tflog.Debug(ctx, "Delete is a no-op for policyurlset_export; NITRO has no inverse of the export action")
}

func policyurlset_exportGetThePayloadFromthePlan(ctx context.Context, data *PolicyurlsetExportResourceModel) policy.Policyurlset {
	tflog.Debug(ctx, "In policyurlset_exportGetThePayloadFromthePlan Function")

	// NITRO `?action=export` accepts only `name` and `url`.
	policyurlset := policy.Policyurlset{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		policyurlset.Name = data.Name.ValueString()
	}
	if !data.Url.IsNull() && !data.Url.IsUnknown() {
		policyurlset.Url = data.Url.ValueString()
	}

	return policyurlset
}
