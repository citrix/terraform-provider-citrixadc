package systemkek

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/system"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// systemkek_change is an ACTION-ONLY resource. NITRO exposes only the `change`
// verb, which is POST /nitro/v1/config/systemkek?action=update. There is NO add,
// get, delete, or count endpoint.
//
// WARNING: applying this resource ROTATES the appliance Key Encryption Key (KEK).
// This action is IRREVERSIBLE and NON-IDEMPOTENT: each apply backs up the old
// keys and generates brand-new keys. Because there is no GET endpoint, drift
// cannot be detected; the only way to "re-run" the rotation is to recreate the
// resource (every attribute is RequiresReplace).

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SystemkekChangeResource{}
var _ resource.ResourceWithConfigure = (*SystemkekChangeResource)(nil)
var _ resource.ResourceWithImportState = (*SystemkekChangeResource)(nil)

func NewSystemkekChangeResource() resource.Resource {
	return &SystemkekChangeResource{}
}

// SystemkekChangeResource defines the resource implementation.
type SystemkekChangeResource struct {
	client *service.NitroClient
}

// SystemkekChangeResourceModel describes the resource data model.
type SystemkekChangeResourceModel struct {
	Id    types.String `tfsdk:"id"`
	Level types.String `tfsdk:"level"`
}

func (r *SystemkekChangeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SystemkekChangeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_systemkek_change"
}

func (r *SystemkekChangeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SystemkekChangeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the systemkek_change resource.",
			},
			"level": schema.StringAttribute{
				// CLI + NITRO mandatory (tfdata wrongly had is_required:false) -> Required (Pattern 8).
				// RequiresReplace: re-applying forces a fresh KEK rotation; there is no
				// update/GET endpoint so this is the only way to re-run the action.
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Type of update KEK to be performed.\n*basic : The level basic will backup old keys and create new keys and respond back.\n*extended : The level extended will backup old keys and create new keys, update\nns.conf, nscfg.db, all ns.conf for same release, in all partitions. While doing so\n will block all config changes and once done shall respond back.",
			},
		},
	}
}

func (r *SystemkekChangeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SystemkekChangeResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating systemkek_change resource (rotating appliance KEK)")
	systemkek := systemkek_changeGetThePayloadFromthePlan(ctx, &data)

	// Action-only resource: the `change` verb is POST ?action=update.
	// This ROTATES the appliance KEK (irreversible, non-idempotent).
	err := r.client.ActOnResource(service.Systemkek.Type(), &systemkek, "update")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create systemkek, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created systemkek_change resource")

	// Set synthetic constant ID (no GET endpoint, so nothing to read back).
	data.Id = types.StringValue("systemkek_change")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemkekChangeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SystemkekChangeResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read is a no-op: NITRO exposes no GET endpoint for systemkek, so we
	// preserve the prior state unchanged. Drift detection is impossible.
	tflog.Debug(ctx, "Read is a no-op for systemkek_change; no GET endpoint on NITRO side")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemkekChangeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SystemkekChangeResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Update is a no-op for systemkek_change; the only attribute (level) is
	// RequiresReplace, so any change forces recreation (a fresh KEK rotation)
	// and this method is never reached for an attribute change.
	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for systemkek_change; all attributes are RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemkekChangeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SystemkekChangeResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// No delete endpoint on NITRO side for systemkek (KEK rotation cannot be
	// undone); just remove the resource from Terraform state.
	tflog.Debug(ctx, "Deleting systemkek_change resource (state removal only; no NITRO delete endpoint)")
	tflog.Trace(ctx, "Removed systemkek_change from Terraform state")
}

func systemkek_changeGetThePayloadFromthePlan(ctx context.Context, data *SystemkekChangeResourceModel) system.Systemkek {
	tflog.Debug(ctx, "In systemkek_changeGetThePayloadFromthePlan Function")

	// Create API request body from the model
	systemkek := system.Systemkek{}
	if !data.Level.IsNull() && !data.Level.IsUnknown() {
		systemkek.Level = data.Level.ValueString()
	}

	return systemkek
}
