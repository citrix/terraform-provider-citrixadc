package sslfipssimsource

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
var _ resource.Resource = &SslfipssimsourceEnableResource{}
var _ resource.ResourceWithConfigure = (*SslfipssimsourceEnableResource)(nil)
var _ resource.ResourceWithImportState = (*SslfipssimsourceEnableResource)(nil)

func NewSslfipssimsourceEnableResource() resource.Resource {
	return &SslfipssimsourceEnableResource{}
}

// SslfipssimsourceEnableResource defines the resource implementation.
type SslfipssimsourceEnableResource struct {
	client *service.NitroClient
}

// SslfipssimsourceEnableResourceModel describes the resource data model.
//
// This resource models the NITRO sslfipssimsource `?action=enable` action. enable
// is a one-shot side-effect action with no GET endpoint and no inverse API, so
// Read/Update/Delete are no-ops. The enable payload carries the mandatory
// targetsecret and sourcesecret attributes (certfile is init-only).
type SslfipssimsourceEnableResourceModel struct {
	Id           types.String `tfsdk:"id"`
	Sourcesecret types.String `tfsdk:"sourcesecret"`
	Targetsecret types.String `tfsdk:"targetsecret"`
}

func (r *SslfipssimsourceEnableResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslfipssimsourceEnableResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslfipssimsource_enable"
}

func (r *SslfipssimsourceEnableResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslfipssimsourceEnableResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		// WARNING: DISRUPTIVE and FIPS-only. sslfipssimsource exposes only the
		// `enable` and `init` NITRO actions (no get/add/delete). This resource
		// models the `enable` action, whose payload is {targetsecret, sourcesecret}.
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslfipssimsource_enable resource.",
			},
			"sourcesecret": schema.StringAttribute{
				// Required for the enable action (Pattern 8).
				Required:  true,
				Sensitive: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for and, optionally, path to the source FIPS appliance's secret data. /nsconfig/ssl/ is the default path.",
			},
			"targetsecret": schema.StringAttribute{
				// Required for the enable action (Pattern 8).
				Required:  true,
				Sensitive: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of and, optionally, path to the target FIPS appliance's secret data. /nsconfig/ssl/ is the default path.",
			},
		},
	}
}

func (r *SslfipssimsourceEnableResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslfipssimsourceEnableResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Enabling sslfipssimsource (action-only resource)")
	payload := sslfipssimsource_enableGetThePayloadFromthePlan(ctx, &data)

	// NITRO exposes enable as POST ?action=enable. Use ActOnResource with the
	// case-sensitive "enable" verb (lower-case per the NITRO URL).
	// WARNING: DISRUPTIVE and FIPS-only - requires dedicated FIPS hardware.
	err := r.client.ActOnResource(service.Sslfipssimsource.Type(), &payload, "enable")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to enable sslfipssimsource, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Enabled sslfipssimsource")

	// Synthetic ID for the action-only resource; keeps Read/Delete no-ops
	// addressable by Terraform (no GET endpoint to derive it from).
	data.Id = types.StringValue("sslfipssimsource_enable")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslfipssimsourceEnableResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// enable is a one-shot action. NITRO has no GET endpoint that reports
	// enable-state, so Read is a pure preserve-state no-op.
	var data SslfipssimsourceEnableResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for sslfipssimsource_enable; NITRO has no query endpoint for enable state")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslfipssimsourceEnableResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO has no update endpoint for enable; every schema attribute is
	// RequiresReplace, so Terraform never invokes Update for a real change.
	var data, state SslfipssimsourceEnableResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for sslfipssimsource_enable; NITRO has no update endpoint and all attributes are RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslfipssimsourceEnableResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// enable is a one-shot side-effect action. There is no inverse NITRO API.
	// Delete simply removes the resource from Terraform state.
	tflog.Debug(ctx, "Delete is a no-op for sslfipssimsource_enable; NITRO has no inverse of the enable action")
}

func sslfipssimsource_enableGetThePayloadFromthePlan(ctx context.Context, data *SslfipssimsourceEnableResourceModel) ssl.Sslfipssimsource {
	tflog.Debug(ctx, "In sslfipssimsource_enableGetThePayloadFromthePlan Function")

	// Create API request body from the model. Only the enable action's fields
	// (targetsecret, sourcesecret) are set; certfile is init-only and omitted.
	sslfipssimsource := ssl.Sslfipssimsource{}
	if !data.Sourcesecret.IsNull() && !data.Sourcesecret.IsUnknown() {
		sslfipssimsource.Sourcesecret = data.Sourcesecret.ValueString()
	}
	if !data.Targetsecret.IsNull() && !data.Targetsecret.IsUnknown() {
		sslfipssimsource.Targetsecret = data.Targetsecret.ValueString()
	}

	return sslfipssimsource
}
