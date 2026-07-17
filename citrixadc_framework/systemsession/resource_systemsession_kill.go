package systemsession

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SystemsessionKillResource{}
var _ resource.ResourceWithConfigure = (*SystemsessionKillResource)(nil)
var _ resource.ResourceWithImportState = (*SystemsessionKillResource)(nil)

// ValidateConfig enforces the CLI's mandatory mutually-exclusive `(<sid> | -all)` choice
// for the kill action (NitroValidator Pattern 8/17). Keeping this assertion is required
// so the framework invokes ValidateConfig.
var _ resource.ResourceWithValidateConfig = (*SystemsessionKillResource)(nil)

func NewSystemsessionKillResource() resource.Resource {
	return &SystemsessionKillResource{}
}

// SystemsessionKillResource models the NITRO systemsession `?action=kill` (POST) action.
//
// kill is a DESTRUCTIVE one-shot action. NITRO exposes no add/update/delete API for it —
// Create invokes the kill action; Read/Update/Delete are no-ops (there is no un-kill).
//
// WARNING: `kill systemsession -all` terminates ALL admin sessions, INCLUDING the
// provider's own NITRO session. Killing a specific `sid` terminates only that session.
type SystemsessionKillResource struct {
	client *service.NitroClient
}

// SystemsessionKillResourceModel describes the resource data model.
//
// The kill action accepts only `sid` and `all` (Pattern 15: all GET-response-only fields
// such as username/logintime/clientipaddress are intentionally excluded).
type SystemsessionKillResourceModel struct {
	Id  types.String `tfsdk:"id"`
	All types.Bool   `tfsdk:"all"`
	Sid types.Int64  `tfsdk:"sid"`
}

func (r *SystemsessionKillResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SystemsessionKillResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_systemsession_kill"
}

func (r *SystemsessionKillResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

// ValidateConfig enforces the CLI's mutually-exclusive `(<sid> | -all)` choice for the
// kill action: exactly one of `sid` or `all` must be supplied.
func (r *SystemsessionKillResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data SystemsessionKillResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	sidSet := !data.Sid.IsNull() && !data.Sid.IsUnknown()
	allSet := !data.All.IsNull() && !data.All.IsUnknown() && data.All.ValueBool()

	if !sidSet && !allSet {
		resp.Diagnostics.AddError(
			"Invalid Attribute Combination",
			"Exactly one of \"sid\" or \"all\" must be specified to kill a system session.",
		)
		return
	}
	if sidSet && allSet {
		resp.Diagnostics.AddAttributeError(
			path.Root("sid"),
			"Invalid Attribute Combination",
			"Only one of \"sid\" or \"all\" may be specified; they are mutually exclusive.",
		)
	}
}

func (r *SystemsessionKillResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		// DESTRUCTIVE ONE-SHOT ACTION RESOURCE.
		// `kill systemsession -all` terminates ALL admin sessions, INCLUDING the
		// provider's own NITRO session. Killing a specific `sid` terminates that one
		// session. This resource performs the kill action on Create and has no GET-based
		// drift detection (sessions are transient); Read/Update/Delete are no-ops.
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Description: "The ID of the systemsession_kill resource.",
			},
			"all": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					// Action input — re-running with a different value is a new kill action.
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Terminate all the system sessions except the current session.",
			},
			"sid": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "ID of the system session to kill.",
			},
		},
	}
}

func (r *SystemsessionKillResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SystemsessionKillResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Killing systemsession (action-only resource)")
	payload := systemsession_killGetThePayloadFromthePlan(ctx, &data)

	// Action-only resource: kill via ?action=kill (POST). No add/update/delete API.
	err := r.client.ActOnResource(service.Systemsession.Type(), payload, "kill")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to kill systemsession, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Killed systemsession")

	// Synthetic ID for the action-only resource; keeps Read/Delete no-ops addressable.
	data.Id = types.StringValue("systemsession_kill")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Read is a no-op. systemsessions are transient runtime objects; the kill action has
// no persistent footprint to reconcile, so prior state is preserved unchanged (do not
// fail or clear state if the session is gone).
func (r *SystemsessionKillResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SystemsessionKillResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for systemsession_kill; kill is a one-shot action with no GET-based drift detection")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Update is a no-op. All attributes are RequiresReplace and there is no NITRO update
// endpoint for systemsession, so this branch is never meaningfully reached.
func (r *SystemsessionKillResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SystemsessionKillResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for systemsession_kill; all attributes are RequiresReplace and there is no update endpoint")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Delete removes the resource from Terraform state only. There is no inverse of the
// kill action on NITRO, so no API call is made.
func (r *SystemsessionKillResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete is a state-only removal for systemsession_kill; kill has no inverse NITRO API")
}

// systemsession_killGetThePayloadFromthePlan builds the kill action payload. Only `sid`
// and `all` are valid write parameters for ?action=kill (Pattern 15: all read-only GET
// fields are excluded). Returns map[string]interface{} to match the original impl.
func systemsession_killGetThePayloadFromthePlan(ctx context.Context, data *SystemsessionKillResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In systemsession_killGetThePayloadFromthePlan Function")

	payload := map[string]interface{}{}
	if !data.All.IsNull() && !data.All.IsUnknown() {
		payload["all"] = data.All.ValueBool()
	}
	if !data.Sid.IsNull() && !data.Sid.IsUnknown() {
		payload["sid"] = utils.IntPtr(int(data.Sid.ValueInt64()))
	}

	return payload
}
