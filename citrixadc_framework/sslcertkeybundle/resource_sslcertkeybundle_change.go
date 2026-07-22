package sslcertkeybundle

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SslcertkeybundleChangeResource{}
var _ resource.ResourceWithConfigure = (*SslcertkeybundleChangeResource)(nil)

func NewSslcertkeybundleChangeResource() resource.Resource {
	return &SslcertkeybundleChangeResource{}
}

// SslcertkeybundleChangeResource defines the resource implementation.
type SslcertkeybundleChangeResource struct {
	client *service.NitroClient
}

// SslcertkeybundleChangeResourceModel describes the resource data model.
//
// This resource models the NITRO sslcertkeybundle `change` action. The NITRO doc
// labels the operation `change`, but the on-wire verb is `?action=update` (the
// backing CLI verb is `update`), so Create calls ActOnResource(..., "update").
// change is a one-shot side-effect action with no GET endpoint and no inverse
// API, so Read/Update/Delete are no-ops. The change payload carries
// certkeybundlename, bundlefile, and passplain.
type SslcertkeybundleChangeResourceModel struct {
	Id                types.String `tfsdk:"id"`
	Certkeybundlename types.String `tfsdk:"certkeybundlename"`
	Bundlefile        types.String `tfsdk:"bundlefile"`
	Passplain         types.String `tfsdk:"passplain"`
}

func (r *SslcertkeybundleChangeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslcertkeybundle_change"
}

func (r *SslcertkeybundleChangeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslcertkeybundleChangeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslcertkeybundle_change resource.",
			},
			// NITRO change payload marks certkeybundlename as mandatory.
			"certkeybundlename": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name given to the certKeyBundle. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Maximum length: 127.",
			},
			// Optional for the change action (mandatory only for add).
			"bundlefile": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of and, optionally, path to the X509 certificate bundle file that is used to form the certificate-key bundle. /nsconfig/ssl/ is the default path. Maximum length: 255.",
			},
			// Secret: pass phrase to decrypt an encrypted private-key in the bundle.
			"passplain": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Pass phrase used to encrypt the private-key. Required when the certificate bundle file contains an encrypted private-key in PEM format. Maximum length: 31.",
			},
		},
	}
}

func (r *SslcertkeybundleChangeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslcertkeybundleChangeResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslcertkeybundle_change resource")
	payload := sslcertkeybundle_changeGetThePayloadFromthePlan(ctx, &data)

	// NITRO doc labels the op `change`, but the on-wire verb is ?action=update.
	// Pass the case-sensitive "update" verb (NOT "change").
	err := r.client.ActOnResource(service.Sslcertkeybundle.Type(), &payload, "update")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to change sslcertkeybundle, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Changed sslcertkeybundle resource")

	data.Id = types.StringValue(fmt.Sprintf("sslcertkeybundle_change-%v", data.Certkeybundlename.ValueString()))

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslcertkeybundleChangeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// change is a one-shot action. NITRO has no GET endpoint that reports
	// change-state, so Read is a pure preserve-state no-op.
	var data SslcertkeybundleChangeResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for sslcertkeybundle_change; NITRO has no query endpoint for change state")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslcertkeybundleChangeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO has no update endpoint for this wrapper; every schema attribute is
	// RequiresReplace, so Terraform never invokes Update for a real change.
	var data, state SslcertkeybundleChangeResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for sslcertkeybundle_change; all attributes are RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslcertkeybundleChangeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// change is a one-shot side-effect action. There is no inverse NITRO API.
	// Delete simply removes the resource from Terraform state.
	tflog.Debug(ctx, "Delete is a no-op for sslcertkeybundle_change; NITRO has no inverse of the change action")
}

func sslcertkeybundle_changeGetThePayloadFromthePlan(ctx context.Context, data *SslcertkeybundleChangeResourceModel) ssl.Sslcertkeybundle {
	tflog.Debug(ctx, "In sslcertkeybundle_changeGetThePayloadFromthePlan Function")

	// NITRO `?action=update` accepts certkeybundlename, bundlefile, passplain.
	sslcertkeybundle := ssl.Sslcertkeybundle{}
	if !data.Certkeybundlename.IsNull() && !data.Certkeybundlename.IsUnknown() {
		sslcertkeybundle.Certkeybundlename = data.Certkeybundlename.ValueString()
	}
	if !data.Bundlefile.IsNull() && !data.Bundlefile.IsUnknown() {
		sslcertkeybundle.Bundlefile = data.Bundlefile.ValueString()
	}
	if !data.Passplain.IsNull() && !data.Passplain.IsUnknown() {
		sslcertkeybundle.Passplain = data.Passplain.ValueString()
	}

	return sslcertkeybundle
}
