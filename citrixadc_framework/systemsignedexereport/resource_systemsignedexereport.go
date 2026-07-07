package systemsignedexereport

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// systemsignedexereport is an ACTION-ONLY, ZERO-ATTRIBUTE resource modelling an
// ENABLE/DISABLE toggle.
//
//   - NITRO exposes only the enable/disable actions:
//     POST /nitro/v1/config/systemsignedexereport?action=enable
//     POST /nitro/v1/config/systemsignedexereport?action=disable
//     (empty body). There is NO add/set/get/delete endpoint.
//   - The enable/disable pair is modelled as a clean inverse:
//     Create performs the "enable" action, Delete performs the "disable" action.
//     Read/Update are no-ops (there is nothing to reconcile).
//   - Because there is no GET endpoint, there is NO datasource for this resource
//     and it cannot be verified by reading it back.
var _ resource.Resource = &SystemsignedexereportResource{}
var _ resource.ResourceWithConfigure = (*SystemsignedexereportResource)(nil)
var _ resource.ResourceWithImportState = (*SystemsignedexereportResource)(nil)

func NewSystemsignedexereportResource() resource.Resource {
	return &SystemsignedexereportResource{}
}

// SystemsignedexereportResource defines the resource implementation.
type SystemsignedexereportResource struct {
	client *service.NitroClient
}

func (r *SystemsignedexereportResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SystemsignedexereportResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_systemsignedexereport"
}

func (r *SystemsignedexereportResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SystemsignedexereportResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SystemsignedexereportResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Enabling the systemsignedexereport (enable action)")
	payload := systemsignedexereportGetThePayloadFromthePlan(ctx, &data)

	// Action-only resource - Create maps to the "enable" action.
	err := r.client.ActOnResource(systemsignedexereportResourceType, payload, "enable")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to enable systemsignedexereport, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Enabled systemsignedexereport")

	// Synthetic ID - there is no GET endpoint to read back.
	data.Id = types.StringValue("systemsignedexereport-config")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Read is a no-op. systemsignedexereport has no GET endpoint; there is nothing
// to reconcile.
func (r *SystemsignedexereportResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SystemsignedexereportResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for systemsignedexereport; NITRO exposes no GET endpoint")

	// Preserve prior state unchanged.
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Update is a no-op. systemsignedexereport has no attributes and no set endpoint.
func (r *SystemsignedexereportResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SystemsignedexereportResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Update is a no-op for systemsignedexereport; it has no attributes and no set endpoint")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Delete performs the inverse of Create: the "disable" action.
func (r *SystemsignedexereportResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SystemsignedexereportResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Disabling the systemsignedexereport (disable action)")
	payload := systemsignedexereportGetThePayloadFromthePlan(ctx, &data)

	// Delete maps to the "disable" action (inverse of Create's "enable").
	err := r.client.ActOnResource(systemsignedexereportResourceType, payload, "disable")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to disable systemsignedexereport, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Disabled systemsignedexereport")
}
