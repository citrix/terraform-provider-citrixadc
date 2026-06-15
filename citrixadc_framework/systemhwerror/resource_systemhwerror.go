package systemhwerror

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
var _ resource.Resource = &SystemhwerrorResource{}
var _ resource.ResourceWithConfigure = (*SystemhwerrorResource)(nil)
var _ resource.ResourceWithImportState = (*SystemhwerrorResource)(nil)

func NewSystemhwerrorResource() resource.Resource {
	return &SystemhwerrorResource{}
}

// SystemhwerrorResource defines the resource implementation.
type SystemhwerrorResource struct {
	client *service.NitroClient
}

func (r *SystemhwerrorResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SystemhwerrorResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_systemhwerror"
}

func (r *SystemhwerrorResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SystemhwerrorResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SystemhwerrorResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating systemhwerror resource")
	systemhwerror := systemhwerrorGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Action-only resource - NITRO exposes only `check` (?action=check, POST)
	err := r.client.ActOnResource(service.Systemhwerror.Type(), &systemhwerror, "check")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create systemhwerror, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created systemhwerror resource")

	// Set ID for the resource
	data.Id = types.StringValue("systemhwerror-config")

	// No GET endpoint on NITRO side - state is taken directly from the plan

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemhwerrorResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SystemhwerrorResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read is a no-op for systemhwerror: it is an action-only resource
	// (?action=check) with no GET endpoint. Preserve prior state.
	tflog.Debug(ctx, "Read is a no-op for systemhwerror; NITRO has no GET endpoint")

	// Save prior state back unchanged
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemhwerrorResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SystemhwerrorResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Update is a no-op for systemhwerror: it is an action-only resource and
	// all attributes use RequiresReplace, so this branch is never reached with
	// an actual attribute change.
	tflog.Debug(ctx, "Update is a no-op for systemhwerror; all attributes are RequiresReplace")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemhwerrorResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SystemhwerrorResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting systemhwerror resource")
	// Action-only resource - no delete operation on ADC, just remove from state
	tflog.Trace(ctx, "Removed systemhwerror from Terraform state")
}
