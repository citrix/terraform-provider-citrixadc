package systemhwerror

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/system"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SystemhwerrorCheckResource{}
var _ resource.ResourceWithConfigure = (*SystemhwerrorCheckResource)(nil)

func NewSystemhwerrorCheckResource() resource.Resource {
	return &SystemhwerrorCheckResource{}
}

// SystemhwerrorCheckResource defines the resource implementation.
type SystemhwerrorCheckResource struct {
	client *service.NitroClient
}

// SystemhwerrorCheckResourceModel describes the resource data model.
//
// This resource models the NITRO systemhwerror `?action=check` action. check is
// a one-shot side-effect action with no GET endpoint and no inverse API, so
// Read/Update/Delete are no-ops.
type SystemhwerrorCheckResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Diskcheck types.Bool   `tfsdk:"diskcheck"`
}

func (r *SystemhwerrorCheckResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_systemhwerror_check"
}

func (r *SystemhwerrorCheckResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SystemhwerrorCheckResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the systemhwerror_check resource.",
			},
			"diskcheck": schema.BoolAttribute{
				Required: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Perform only disk error checking.",
			},
		},
	}
}

func (r *SystemhwerrorCheckResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SystemhwerrorCheckResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Checking systemhwerror (action-only resource)")
	payload := systemhwerror_checkGetThePayloadFromthePlan(ctx, &data)

	// NITRO exposes check as POST ?action=check. Use ActOnResource with the
	// case-sensitive "check" verb (per the NITRO URL).
	err := r.client.ActOnResource(service.Systemhwerror.Type(), &payload, "check")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create systemhwerror, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Checked systemhwerror")

	// Synthetic ID for the action-only resource; keeps Read/Delete no-ops
	// addressable by Terraform.
	data.Id = types.StringValue("systemhwerror_check")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemhwerrorCheckResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// check is a one-shot action. NITRO has no GET endpoint that reports
	// check-state, so Read is a pure preserve-state no-op.
	var data SystemhwerrorCheckResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for systemhwerror_check; NITRO has no query endpoint for check state")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemhwerrorCheckResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO has no update endpoint for check; every schema attribute is
	// RequiresReplace, so Terraform never invokes Update for a real change.
	var data, state SystemhwerrorCheckResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for systemhwerror_check; NITRO has no update endpoint and all attributes are RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemhwerrorCheckResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// check is a one-shot side-effect action. There is no inverse NITRO API
	// (no "un-check"). Delete simply removes the resource from Terraform state.
	tflog.Debug(ctx, "Delete is a no-op for systemhwerror_check; NITRO has no inverse of the check action")
}

func systemhwerror_checkGetThePayloadFromthePlan(ctx context.Context, data *SystemhwerrorCheckResourceModel) system.Systemhwerror {
	tflog.Debug(ctx, "In systemhwerror_checkGetThePayloadFromthePlan Function")

	// Create API request body from the model
	systemhwerror := system.Systemhwerror{}
	if !data.Diskcheck.IsNull() && !data.Diskcheck.IsUnknown() {
		systemhwerror.Diskcheck = data.Diskcheck.ValueBool()
	}

	return systemhwerror
}
