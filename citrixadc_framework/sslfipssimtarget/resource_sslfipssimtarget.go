package sslfipssimtarget

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
var _ resource.Resource = &SslfipssimtargetResource{}
var _ resource.ResourceWithConfigure = (*SslfipssimtargetResource)(nil)
var _ resource.ResourceWithImportState = (*SslfipssimtargetResource)(nil)

func NewSslfipssimtargetResource() resource.Resource {
	return &SslfipssimtargetResource{}
}

// SslfipssimtargetResource defines the resource implementation.
type SslfipssimtargetResource struct {
	client *service.NitroClient
}

func (r *SslfipssimtargetResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslfipssimtargetResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslfipssimtarget"
}

func (r *SslfipssimtargetResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslfipssimtargetResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslfipssimtargetResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslfipssimtarget resource")
	sslfipssimtarget := sslfipssimtargetGetThePayloadFromthePlan(ctx, &data)

	// Action-only resource: the NITRO doc exposes only `enable` and `init` actions
	// (no add/get/delete). `enable` is the primary action used for Create.
	// WARNING: DISRUPTIVE and FIPS-only - requires dedicated FIPS hardware.
	err := r.client.ActOnResource(service.Sslfipssimtarget.Type(), &sslfipssimtarget, "enable")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslfipssimtarget, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created sslfipssimtarget resource")

	// Set synthetic ID for the resource (no GET endpoint to derive it from).
	data.Id = types.StringValue("sslfipssimtarget-config")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslfipssimtargetResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslfipssimtargetResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// No GET endpoint on the NITRO side (action-only resource) - Read is a no-op
	// that preserves prior state. Drift detection is impossible by definition.
	tflog.Debug(ctx, "Read is a no-op for sslfipssimtarget; no GET endpoint on NITRO side")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslfipssimtargetResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SslfipssimtargetResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Update is a no-op for sslfipssimtarget; all attributes are RequiresReplace and
	// there is no NITRO update endpoint (action-only resource).
	tflog.Debug(ctx, "Update is a no-op for sslfipssimtarget")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslfipssimtargetResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslfipssimtargetResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslfipssimtarget resource")
	// Action-only resource - no delete operation on ADC, just remove from state
	tflog.Trace(ctx, "Removed sslfipssimtarget from Terraform state")
}
