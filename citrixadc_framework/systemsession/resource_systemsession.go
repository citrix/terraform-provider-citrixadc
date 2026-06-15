package systemsession

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SystemsessionResource{}
var _ resource.ResourceWithConfigure = (*SystemsessionResource)(nil)
var _ resource.ResourceWithImportState = (*SystemsessionResource)(nil)
var _ resource.ResourceWithValidateConfig = (*SystemsessionResource)(nil)

func NewSystemsessionResource() resource.Resource {
	return &SystemsessionResource{}
}

// SystemsessionResource defines the resource implementation.
//
// systemsession is a DESTRUCTIVE one-shot ACTION resource. NITRO exposes only
// get/get(all)/count and the `kill` action (?action=kill, POST) — there is NO
// add/update/delete API. Create invokes the kill action; Read/Update are no-ops;
// Delete simply drops the resource from Terraform state (there is no un-kill).
//
// WARNING: `kill systemsession -all` terminates ALL admin sessions, INCLUDING the
// provider's own NITRO session. Killing a specific `sid` terminates only that session.
type SystemsessionResource struct {
	client *service.NitroClient
}

func (r *SystemsessionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SystemsessionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_systemsession"
}

func (r *SystemsessionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

// ValidateConfig enforces the CLI's mutually-exclusive `(<sid> | -all)` choice for the
// kill action: exactly one of `sid` or `all` must be supplied.
func (r *SystemsessionResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data SystemsessionResourceModel
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

func (r *SystemsessionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SystemsessionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Killing systemsession (action-only resource)")
	payload := systemsessionGetThePayloadFromthePlan(ctx, &data)

	// Action-only resource: kill via ?action=kill (POST). No add/update/delete API.
	err := r.client.ActOnResource(service.Systemsession.Type(), payload, "kill")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to kill systemsession, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Killed systemsession")

	// Synthetic ID set exactly once (Pattern 6): the sid value, or "all" when all=true.
	if !data.All.IsNull() && !data.All.IsUnknown() && data.All.ValueBool() {
		data.Id = types.StringValue("all")
	} else {
		data.Id = types.StringValue(fmt.Sprintf("%v", data.Sid.ValueInt64()))
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Read is a no-op. systemsessions are transient runtime objects; the kill action has
// no persistent footprint to reconcile, so prior state is preserved unchanged (do not
// fail or clear state if the session is gone).
func (r *SystemsessionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SystemsessionResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for systemsession; kill is a one-shot action with no GET-based drift detection")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Update is a no-op. All attributes are RequiresReplace and there is no NITRO update
// endpoint for systemsession, so this branch is never meaningfully reached.
func (r *SystemsessionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SystemsessionResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for systemsession; all attributes are RequiresReplace and there is no update endpoint")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Delete removes the resource from Terraform state only. There is no inverse of the
// kill action on NITRO, so no API call is made.
func (r *SystemsessionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete is a state-only removal for systemsession; kill has no inverse NITRO API")
}
