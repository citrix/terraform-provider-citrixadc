package sslfipssimtarget

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
var _ resource.Resource = &SslfipssimtargetInitResource{}
var _ resource.ResourceWithConfigure = (*SslfipssimtargetInitResource)(nil)

func NewSslfipssimtargetInitResource() resource.Resource {
	return &SslfipssimtargetInitResource{}
}

// SslfipssimtargetInitResource defines the resource implementation.
//
// This resource models the NITRO sslfipssimtarget `?action=init` action. init is
// a one-shot side-effect action with no GET endpoint and no inverse API, so
// Read/Update/Delete are no-ops. The init payload carries certfile, keyvector and
// targetsecret (all mandatory per the NITRO doc and CLI).
// WARNING: DISRUPTIVE and FIPS-only - requires dedicated FIPS hardware.
type SslfipssimtargetInitResource struct {
	client *service.NitroClient
}

// SslfipssimtargetInitResourceModel describes the resource data model.
type SslfipssimtargetInitResourceModel struct {
	Id           types.String `tfsdk:"id"`
	Certfile     types.String `tfsdk:"certfile"`
	Keyvector    types.String `tfsdk:"keyvector"`
	Targetsecret types.String `tfsdk:"targetsecret"`
}

func (r *SslfipssimtargetInitResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslfipssimtarget_init"
}

func (r *SslfipssimtargetInitResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslfipssimtargetInitResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		// WARNING: DISRUPTIVE and FIPS-only. sslfipssimtarget_init models the
		// NITRO `init` action (no get/add/delete). It requires dedicated FIPS
		// hardware and is unsupported on non-FIPS VPX appliances.
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslfipssimtarget_init resource.",
			},
			"certfile": schema.StringAttribute{
				// Required for the init action (CLI/NITRO mandatory).
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of and, optionally, path to the source FIPS appliance's certificate file. /nsconfig/ssl/ is the default path.",
			},
			"keyvector": schema.StringAttribute{
				// Required for the init action (CLI/NITRO mandatory).
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of and, optionally, path to the target FIPS appliance's key vector. /nsconfig/ssl/ is the default path.",
			},
			"targetsecret": schema.StringAttribute{
				// Required for the init action (CLI/NITRO mandatory).
				Required:  true,
				Sensitive: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for and, optionally, path to the target FIPS appliance's secret data. The default input path for the secret data is /nsconfig/ssl/.",
			},
		},
	}
}

func (r *SslfipssimtargetInitResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslfipssimtargetInitResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Initializing sslfipssimtarget (action-only resource)")
	payload := sslfipssimtarget_initGetThePayloadFromthePlan(ctx, &data)

	// NITRO exposes init as POST ?action=init. Use ActOnResource with the
	// case-sensitive "init" verb (lower-case per the NITRO URL).
	err := r.client.ActOnResource(service.Sslfipssimtarget.Type(), &payload, "init")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to init sslfipssimtarget, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Initialized sslfipssimtarget")

	// Synthetic ID for the action-only resource (no GET endpoint to derive it from).
	data.Id = types.StringValue("sslfipssimtarget_init")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslfipssimtargetInitResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslfipssimtargetInitResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// No GET endpoint on the NITRO side (action-only resource) - Read is a no-op
	// that preserves prior state. Drift detection is impossible by definition.
	tflog.Debug(ctx, "Read is a no-op for sslfipssimtarget_init; no GET endpoint on NITRO side")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslfipssimtargetInitResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SslfipssimtargetInitResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Update is a no-op for sslfipssimtarget_init; all attributes are
	// RequiresReplace and there is no NITRO update endpoint (action-only resource).
	tflog.Debug(ctx, "Update is a no-op for sslfipssimtarget_init")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslfipssimtargetInitResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// init is a one-shot side-effect action. There is no inverse NITRO API.
	// Delete simply removes the resource from Terraform state.
	tflog.Debug(ctx, "Delete is a no-op for sslfipssimtarget_init; NITRO has no inverse of the init action")
}

func sslfipssimtarget_initGetThePayloadFromthePlan(ctx context.Context, data *SslfipssimtargetInitResourceModel) ssl.Sslfipssimtarget {
	tflog.Debug(ctx, "In sslfipssimtarget_initGetThePayloadFromthePlan Function")

	// Create API request body from the model. init accepts only certfile,
	// keyvector and targetsecret.
	sslfipssimtarget := ssl.Sslfipssimtarget{}
	if !data.Certfile.IsNull() && !data.Certfile.IsUnknown() {
		sslfipssimtarget.Certfile = data.Certfile.ValueString()
	}
	if !data.Keyvector.IsNull() && !data.Keyvector.IsUnknown() {
		sslfipssimtarget.Keyvector = data.Keyvector.ValueString()
	}
	if !data.Targetsecret.IsNull() && !data.Targetsecret.IsUnknown() {
		sslfipssimtarget.Targetsecret = data.Targetsecret.ValueString()
	}

	return sslfipssimtarget
}
