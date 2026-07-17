package sslfips

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SslfipsResetResource{}
var _ resource.ResourceWithConfigure = (*SslfipsResetResource)(nil)
var _ resource.ResourceWithImportState = (*SslfipsResetResource)(nil)

func NewSslfipsResetResource() resource.Resource {
	return &SslfipsResetResource{}
}

// SslfipsResetResource defines the resource implementation.
type SslfipsResetResource struct {
	client *service.NitroClient
}

// SslfipsResetResourceModel describes the resource data model.
//
// This resource models the NITRO sslfips `?action=reset` action. reset is a
// one-shot side-effect action with no GET endpoint and no inverse API, so
// Read/Update/Delete are no-ops. The reset payload carries no input attributes.
type SslfipsResetResourceModel struct {
	Id types.String `tfsdk:"id"`
}

func (r *SslfipsResetResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslfipsResetResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslfips_reset"
}

func (r *SslfipsResetResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslfipsResetResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslfips_reset resource.",
			},
		},
	}
}

func (r *SslfipsResetResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslfipsResetResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslfips_reset resource")
	payload := sslfips_resetGetThePayloadFromthePlan(ctx, &data)

	// NITRO exposes reset as POST ?action=reset. Use ActOnResource with the
	// case-sensitive "reset" verb (lower-case per the NITRO URL).
	err := r.client.ActOnResource(service.Sslfips.Type(), &payload, "reset")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to reset sslfips, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Reset sslfips resource")

	// reset takes no input attributes; use a static synthetic ID.
	data.Id = types.StringValue("sslfips_reset")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslfipsResetResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// reset is a one-shot action. NITRO has no GET endpoint that reports
	// reset-state, so Read is a pure preserve-state no-op.
	var data SslfipsResetResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for sslfips_reset; NITRO has no query endpoint for reset state")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslfipsResetResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO has no update endpoint for reset; the resource has no writable
	// attributes, so Terraform never invokes Update for a real change.
	var data, state SslfipsResetResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for sslfips_reset; NITRO has no update endpoint")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslfipsResetResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// reset is a one-shot side-effect action. There is no inverse NITRO API
	// (no "un-reset"). Delete simply removes the resource from Terraform state.
	tflog.Debug(ctx, "Delete is a no-op for sslfips_reset; NITRO has no inverse of the reset action")
}

func sslfips_resetGetThePayloadFromthePlan(ctx context.Context, data *SslfipsResetResourceModel) ssl.Sslfips {
	tflog.Debug(ctx, "In sslfips_resetGetThePayloadFromthePlan Function")

	// NITRO `?action=reset` accepts no input attributes; empty payload.
	sslfips := ssl.Sslfips{}

	return sslfips
}
