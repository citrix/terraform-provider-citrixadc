package sslfipssimsource

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
var _ resource.Resource = &SslfipssimsourceResource{}
var _ resource.ResourceWithConfigure = (*SslfipssimsourceResource)(nil)
var _ resource.ResourceWithImportState = (*SslfipssimsourceResource)(nil)

func NewSslfipssimsourceResource() resource.Resource {
	return &SslfipssimsourceResource{}
}

// SslfipssimsourceResource defines the resource implementation.
type SslfipssimsourceResource struct {
	client *service.NitroClient
}

func (r *SslfipssimsourceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslfipssimsourceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslfipssimsource"
}

func (r *SslfipssimsourceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslfipssimsourceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslfipssimsourceResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslfipssimsource resource")
	sslfipssimsource := sslfipssimsourceGetThePayloadFromthePlan(ctx, &data)

	// Action-only resource: the NITRO doc exposes only `enable` and `init` actions
	// (no add/get/delete). `enable` is the primary action used for Create.
	// WARNING: DISRUPTIVE and FIPS-only - requires dedicated FIPS hardware.
	err := r.client.ActOnResource(service.Sslfipssimsource.Type(), &sslfipssimsource, "enable")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslfipssimsource, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created sslfipssimsource resource")

	// Set synthetic ID for the resource (no GET endpoint to derive it from).
	data.Id = types.StringValue("sslfipssimsource-config")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslfipssimsourceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslfipssimsourceResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// No GET endpoint on the NITRO side (action-only resource) - Read is a no-op
	// that preserves prior state. Drift detection is impossible by definition.
	tflog.Debug(ctx, "Read is a no-op for sslfipssimsource; no GET endpoint on NITRO side")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslfipssimsourceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SslfipssimsourceResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Update is a no-op for sslfipssimsource; all attributes are RequiresReplace and
	// there is no NITRO update endpoint (action-only resource).
	tflog.Debug(ctx, "Update is a no-op for sslfipssimsource")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslfipssimsourceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslfipssimsourceResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslfipssimsource resource")
	// Action-only resource - no delete operation on ADC, just remove from state
	tflog.Trace(ctx, "Removed sslfipssimsource from Terraform state")
}
