package aaasession

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
var _ resource.Resource = &AaasessionResource{}
var _ resource.ResourceWithConfigure = (*AaasessionResource)(nil)
var _ resource.ResourceWithImportState = (*AaasessionResource)(nil)

func NewAaasessionResource() resource.Resource {
	return &AaasessionResource{}
}

// AaasessionResource defines the resource implementation.
// aaasession is an ACTION-ONLY resource: NITRO exposes get(all), count and
// kill (POST ?action=kill). There is no add, update/set, or delete endpoint.
type AaasessionResource struct {
	client *service.NitroClient
}

func (r *AaasessionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AaasessionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_aaasession"
}

func (r *AaasessionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AaasessionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AaasessionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating (kill action) aaasession resource")
	payload := aaasessionGetThePayloadFromthePlan(ctx, &data)

	// kill is a POST ?action=kill action (Pattern 1). There is no add endpoint;
	// AddResource/UpdateUnnamedResource would target a nonexistent verb.
	err := r.client.ActOnResource(service.Aaasession.Type(), payload, "kill")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to kill aaasession, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Killed aaasession")

	// Set synthetic ID once, here in Create. The killed session is not a
	// queryable managed object, so this ID is purely a Terraform state handle,
	// not a NITRO lookup key.
	data.Id = types.StringValue("aaasession-kill")

	// Read is a no-op (no stable GET-backed managed object); state is what the
	// plan supplied plus the synthetic ID.
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaasessionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AaasessionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read is a documented no-op (Pattern 13). aaasession.kill terminates active
	// AAA-TM/VPN sessions; the killed session is not a persistent object keyed by
	// these inputs, so a GET cannot stably re-resolve "this" record. Re-fetching
	// would cause perpetual drift / spurious resource removal. Preserve state.
	tflog.Debug(ctx, "Read is a no-op for aaasession (kill action has no stable GET-backed object); preserving state")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaasessionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AaasessionResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Update is a no-op (Pattern 5): NITRO exposes no update/set endpoint for
	// aaasession, and every attribute is RequiresReplace, so Terraform never
	// reaches here with real diffs.
	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for aaasession; no NITRO update endpoint and all attributes are RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaasessionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AaasessionResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete is a state-only removal. "kill" is a one-shot action with no inverse
	// API; there is no NITRO delete endpoint for aaasession. Removing the
	// resource from Terraform state is the only meaningful operation.
	tflog.Debug(ctx, "Delete is a state-only removal for aaasession; kill has no inverse NITRO endpoint")
}
