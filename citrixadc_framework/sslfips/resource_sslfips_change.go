package sslfips

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"
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
var _ resource.Resource = &SslfipsChangeResource{}
var _ resource.ResourceWithConfigure = (*SslfipsChangeResource)(nil)
var _ resource.ResourceWithImportState = (*SslfipsChangeResource)(nil)

func NewSslfipsChangeResource() resource.Resource {
	return &SslfipsChangeResource{}
}

// SslfipsChangeResource defines the resource implementation.
type SslfipsChangeResource struct {
	client *service.NitroClient
}

// SslfipsChangeResourceModel describes the resource data model.
//
// This resource models the NITRO sslfips `change` action. NOTE: the NITRO
// operation is named `change` but its URL query parameter is literally
// `?action=update` (and the CLI verb is `update ssl fips`). The action is a
// one-shot side-effect with no GET endpoint and no inverse API, so
// Read/Update/Delete are no-ops. The change payload carries only `fipsfw`.
type SslfipsChangeResourceModel struct {
	Id     types.String `tfsdk:"id"`
	Fipsfw types.String `tfsdk:"fipsfw"`
}

func (r *SslfipsChangeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslfipsChangeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslfips_change"
}

func (r *SslfipsChangeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslfipsChangeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslfips_change resource.",
			},
			// CLI marks -fipsFW mandatory and the NITRO change payload marks
			// fipsfw red/bold (required); tfdata's is_required:false is wrong.
			"fipsfw": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Path to the FIPS firmware file. Maximum length: 63.",
			},
		},
	}
}

func (r *SslfipsChangeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslfipsChangeResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslfips_change resource")
	payload := sslfips_changeGetThePayloadFromthePlan(ctx, &data)

	// NITRO's `change` operation is exposed at POST ?action=update (NOT
	// ?action=change, which does not exist on the appliance). Dispatch the
	// case-sensitive lower-case "update" verb.
	err := r.client.ActOnResource(service.Sslfips.Type(), &payload, "update")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to change sslfips, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Changed sslfips resource")

	// ID = the firmware path that was applied; keeps Read/Delete no-ops
	// addressable by Terraform.
	data.Id = types.StringValue(fmt.Sprintf("sslfips_change-%v", data.Fipsfw.ValueString()))

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslfipsChangeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// change is a one-shot action. NITRO has no GET endpoint that reports
	// change-state, so Read is a pure preserve-state no-op.
	var data SslfipsChangeResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for sslfips_change; NITRO has no query endpoint for change state")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslfipsChangeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO has no update endpoint for this wrapper; every schema attribute is
	// RequiresReplace, so Terraform never invokes Update for a real change.
	var data, state SslfipsChangeResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for sslfips_change; all attributes are RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslfipsChangeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// change is a one-shot side-effect action. There is no inverse NITRO API.
	// Delete simply removes the resource from Terraform state.
	tflog.Debug(ctx, "Delete is a no-op for sslfips_change; NITRO has no inverse of the change action")
}

func sslfips_changeGetThePayloadFromthePlan(ctx context.Context, data *SslfipsChangeResourceModel) ssl.Sslfips {
	tflog.Debug(ctx, "In sslfips_changeGetThePayloadFromthePlan Function")

	// NITRO `change` (?action=update) accepts only `fipsfw`.
	sslfips := ssl.Sslfips{}
	if !data.Fipsfw.IsNull() && !data.Fipsfw.IsUnknown() {
		sslfips.Fipsfw = data.Fipsfw.ValueString()
	}

	return sslfips
}
