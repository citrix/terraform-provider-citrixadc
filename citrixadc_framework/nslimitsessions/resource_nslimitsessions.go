package nslimitsessions

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
var _ resource.Resource = &NslimitsessionsResource{}
var _ resource.ResourceWithConfigure = (*NslimitsessionsResource)(nil)
var _ resource.ResourceWithImportState = (*NslimitsessionsResource)(nil)

func NewNslimitsessionsResource() resource.Resource {
	return &NslimitsessionsResource{}
}

// NslimitsessionsResource defines the resource implementation.
type NslimitsessionsResource struct {
	client *service.NitroClient
}

func (r *NslimitsessionsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NslimitsessionsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nslimitsessions"
}

func (r *NslimitsessionsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NslimitsessionsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NslimitsessionsResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Clearing nslimitsessions resource (action=clear)")
	nslimitsessions := nslimitsessionsGetThePayloadFromthePlan(ctx, &data)

	// nslimitsessions is an action-only resource: NITRO exposes only the
	// clear action (POST ?action=clear). The clear payload carries only
	// limitidentifier.
	err := r.client.ActOnResource(service.Nslimitsessions.Type(), nslimitsessions, "clear")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to clear nslimitsessions, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Cleared nslimitsessions sessions")

	// Synthetic ID = the limitidentifier value (set once here).
	data.Id = types.StringValue(data.Limitidentifier.ValueString())

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NslimitsessionsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NslimitsessionsResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// nslimitsessions sessions are transient and cleared by the action; there
	// is nothing stable to read back. Read is a no-op that preserves state.
	tflog.Debug(ctx, "Read is a no-op for nslimitsessions; clear is an action and sessions are transient")

	// Save (unchanged) data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NslimitsessionsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state NslimitsessionsResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state. nslimitsessions has no update endpoint and
	// limitidentifier is RequiresReplace, so Update is a no-op.
	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for nslimitsessions; all attributes are RequiresReplace")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NslimitsessionsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NslimitsessionsResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// nslimitsessions is action-only (clear); there is no delete endpoint.
	// Delete simply removes the resource from Terraform state.
	tflog.Debug(ctx, "Delete is a state-only removal for nslimitsessions; NITRO exposes no delete endpoint")
}
