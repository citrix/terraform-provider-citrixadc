package nsmemrecovery

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// nsmemrecovery_start is an ACTION-ONLY resource.
//
//   - NITRO exposes only the "start" action:
//     POST /nitro/v1/config/nsmemrecovery?action=start
//     which recovers a configurable percentage of memory from the freepools.
//   - There is NO add/set/get/delete/unset endpoint, so:
//     Create performs the start action, Read is a no-op (nothing to reconcile),
//     Update re-runs the start action when "percentage" changes, and Delete is a
//     state-only removal.
//   - Because there is no GET endpoint, there is NO datasource for
//     nsmemrecovery_start and the resource cannot be verified by reading it back.
var _ resource.Resource = &NsmemrecoveryStartResource{}
var _ resource.ResourceWithConfigure = (*NsmemrecoveryStartResource)(nil)
var _ resource.ResourceWithImportState = (*NsmemrecoveryStartResource)(nil)

func NewNsmemrecoveryStartResource() resource.Resource {
	return &NsmemrecoveryStartResource{}
}

// NsmemrecoveryStartResource defines the resource implementation.
type NsmemrecoveryStartResource struct {
	client *service.NitroClient
}

func (r *NsmemrecoveryStartResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NsmemrecoveryStartResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsmemrecovery_start"
}

func (r *NsmemrecoveryStartResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NsmemrecoveryStartResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NsmemrecoveryStartResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Starting nsmemrecovery (start action)")

	// percentage is Optional+Computed with no schema Default. When the user does
	// not configure it, resolve it to the NITRO default (10) so the Computed
	// value is known and Terraform does not error on a null/unknown result.
	if data.Percentage.IsNull() || data.Percentage.IsUnknown() {
		data.Percentage = types.Int64Value(10)
	}

	nsmemrecovery := nsmemrecoveryStartGetThePayloadFromthePlan(ctx, &data)

	// Action-only resource - the only NITRO operation is the "start" action.
	err := r.client.ActOnResource(service.Nsmemrecovery.Type(), &nsmemrecovery, "start")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to start nsmemrecovery, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Triggered nsmemrecovery start action")

	// Synthetic ID - there is no GET endpoint to read back.
	data.Id = types.StringValue("nsmemrecovery-config")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Read is a no-op. nsmemrecovery has no GET endpoint; there is nothing to
// reconcile, so prior state is preserved unchanged.
func (r *NsmemrecoveryStartResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NsmemrecoveryStartResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for nsmemrecovery_start; NITRO exposes no GET endpoint")

	// Preserve prior state unchanged.
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Update re-runs the "start" action when percentage changes.
func (r *NsmemrecoveryStartResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state NsmemrecoveryStartResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Resolve unknown/null percentage to the NITRO default (10).
	if data.Percentage.IsNull() || data.Percentage.IsUnknown() {
		data.Percentage = types.Int64Value(10)
	}

	tflog.Debug(ctx, "Updating nsmemrecovery_start resource")

	if !data.Percentage.Equal(state.Percentage) {
		tflog.Debug(ctx, "percentage has changed for nsmemrecovery_start; re-running start action")
		nsmemrecovery := nsmemrecoveryStartGetThePayloadFromthePlan(ctx, &data)
		err := r.client.ActOnResource(service.Nsmemrecovery.Type(), &nsmemrecovery, "start")
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nsmemrecovery_start, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Re-triggered nsmemrecovery start action")
	} else {
		tflog.Debug(ctx, "No changes detected for nsmemrecovery_start resource, skipping start action")
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Delete is a no-op. nsmemrecovery has no delete endpoint; the action is not
// reversible and there is no persistent object to remove.
func (r *NsmemrecoveryStartResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete is a no-op for nsmemrecovery_start; NITRO has no delete endpoint")
}
