package systemkek

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// systemkek is an ACTION-ONLY resource. NITRO exposes only the `change` verb,
// which is POST /nitro/v1/config/systemkek?action=update. There is NO add, get,
// delete, or count endpoint.
//
// WARNING: applying this resource ROTATES the appliance Key Encryption Key (KEK).
// This action is IRREVERSIBLE and NON-IDEMPOTENT: each apply backs up the old
// keys and generates brand-new keys. Because there is no GET endpoint, drift
// cannot be detected; the only way to "re-run" the rotation is to recreate the
// resource (every attribute is RequiresReplace).

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SystemkekResource{}
var _ resource.ResourceWithConfigure = (*SystemkekResource)(nil)
var _ resource.ResourceWithImportState = (*SystemkekResource)(nil)

func NewSystemkekResource() resource.Resource {
	return &SystemkekResource{}
}

// SystemkekResource defines the resource implementation.
type SystemkekResource struct {
	client *service.NitroClient
}

func (r *SystemkekResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SystemkekResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_systemkek"
}

func (r *SystemkekResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SystemkekResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SystemkekResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating systemkek resource (rotating appliance KEK)")
	systemkek := systemkekGetThePayloadFromthePlan(ctx, &data)

	// Action-only resource: the `change` verb is POST ?action=update.
	// This ROTATES the appliance KEK (irreversible, non-idempotent).
	err := r.client.ActOnResource(service.Systemkek.Type(), &systemkek, "update")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create systemkek, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created systemkek resource")

	// Set synthetic constant ID (no GET endpoint, so nothing to read back).
	data.Id = types.StringValue("systemkek-config")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemkekResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SystemkekResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read is a no-op: NITRO exposes no GET endpoint for systemkek, so we
	// preserve the prior state unchanged. Drift detection is impossible.
	tflog.Debug(ctx, "Read is a no-op for systemkek; no GET endpoint on NITRO side")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemkekResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SystemkekResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Update is a no-op for systemkek; the only attribute (level) is
	// RequiresReplace, so any change forces recreation (a fresh KEK rotation)
	// and this method is never reached for an attribute change.
	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for systemkek; all attributes are RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemkekResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SystemkekResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// No delete endpoint on NITRO side for systemkek (KEK rotation cannot be
	// undone); just remove the resource from Terraform state.
	tflog.Debug(ctx, "Deleting systemkek resource (state removal only; no NITRO delete endpoint)")
	tflog.Trace(ctx, "Removed systemkek from Terraform state")
}
