package sslfipssimtarget

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
var _ resource.Resource = &SslfipssimtargetEnableResource{}
var _ resource.ResourceWithConfigure = (*SslfipssimtargetEnableResource)(nil)
var _ resource.ResourceWithImportState = (*SslfipssimtargetEnableResource)(nil)

func NewSslfipssimtargetEnableResource() resource.Resource {
	return &SslfipssimtargetEnableResource{}
}

// SslfipssimtargetEnableResource defines the resource implementation.
//
// This resource models the NITRO sslfipssimtarget `?action=enable` action. enable
// is a one-shot side-effect action with no GET endpoint and no inverse API, so
// Read/Update/Delete are no-ops. The enable payload carries only keyvector and
// sourcesecret (both mandatory per the NITRO doc and CLI).
// WARNING: DISRUPTIVE and FIPS-only - requires dedicated FIPS hardware.
type SslfipssimtargetEnableResource struct {
	client *service.NitroClient
}

// SslfipssimtargetEnableResourceModel describes the resource data model.
type SslfipssimtargetEnableResourceModel struct {
	Id           types.String `tfsdk:"id"`
	Keyvector    types.String `tfsdk:"keyvector"`
	Sourcesecret types.String `tfsdk:"sourcesecret"`
}

func (r *SslfipssimtargetEnableResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslfipssimtargetEnableResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslfipssimtarget_enable"
}

func (r *SslfipssimtargetEnableResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslfipssimtargetEnableResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		// WARNING: DISRUPTIVE and FIPS-only. sslfipssimtarget_enable models the
		// NITRO `enable` action (no get/add/delete). It requires dedicated FIPS
		// hardware and is unsupported on non-FIPS VPX appliances.
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslfipssimtarget_enable resource.",
			},
			"keyvector": schema.StringAttribute{
				// Required for the enable action (CLI/NITRO mandatory).
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of and, optionally, path to the target FIPS appliance's key vector. /nsconfig/ssl/ is the default path.",
			},
			"sourcesecret": schema.StringAttribute{
				// Required for the enable action (CLI/NITRO mandatory).
				Required:  true,
				Sensitive: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of and, optionally, path to the source FIPS appliance's secret data. /nsconfig/ssl/ is the default path.",
			},
		},
	}
}

func (r *SslfipssimtargetEnableResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslfipssimtargetEnableResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Enabling sslfipssimtarget (action-only resource)")
	payload := sslfipssimtarget_enableGetThePayloadFromthePlan(ctx, &data)

	// NITRO exposes enable as POST ?action=enable. Use ActOnResource with the
	// case-sensitive "enable" verb (lower-case per the NITRO URL).
	err := r.client.ActOnResource(service.Sslfipssimtarget.Type(), &payload, "enable")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to enable sslfipssimtarget, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Enabled sslfipssimtarget")

	// Synthetic ID for the action-only resource (no GET endpoint to derive it from).
	data.Id = types.StringValue("sslfipssimtarget_enable")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslfipssimtargetEnableResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslfipssimtargetEnableResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// No GET endpoint on the NITRO side (action-only resource) - Read is a no-op
	// that preserves prior state. Drift detection is impossible by definition.
	tflog.Debug(ctx, "Read is a no-op for sslfipssimtarget_enable; no GET endpoint on NITRO side")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslfipssimtargetEnableResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SslfipssimtargetEnableResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Update is a no-op for sslfipssimtarget_enable; all attributes are
	// RequiresReplace and there is no NITRO update endpoint (action-only resource).
	tflog.Debug(ctx, "Update is a no-op for sslfipssimtarget_enable")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslfipssimtargetEnableResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// enable is a one-shot side-effect action. There is no inverse NITRO API.
	// Delete simply removes the resource from Terraform state.
	tflog.Debug(ctx, "Delete is a no-op for sslfipssimtarget_enable; NITRO has no inverse of the enable action")
}

func sslfipssimtarget_enableGetThePayloadFromthePlan(ctx context.Context, data *SslfipssimtargetEnableResourceModel) ssl.Sslfipssimtarget {
	tflog.Debug(ctx, "In sslfipssimtarget_enableGetThePayloadFromthePlan Function")

	// Create API request body from the model. enable accepts only keyvector and
	// sourcesecret.
	sslfipssimtarget := ssl.Sslfipssimtarget{}
	if !data.Keyvector.IsNull() && !data.Keyvector.IsUnknown() {
		sslfipssimtarget.Keyvector = data.Keyvector.ValueString()
	}
	if !data.Sourcesecret.IsNull() && !data.Sourcesecret.IsUnknown() {
		sslfipssimtarget.Sourcesecret = data.Sourcesecret.ValueString()
	}

	return sslfipssimtarget
}
