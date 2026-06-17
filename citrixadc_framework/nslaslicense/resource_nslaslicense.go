package nslaslicense

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
var _ resource.Resource = &NslaslicenseResource{}
var _ resource.ResourceWithConfigure = (*NslaslicenseResource)(nil)
var _ resource.ResourceWithImportState = (*NslaslicenseResource)(nil)

func NewNslaslicenseResource() resource.Resource {
	return &NslaslicenseResource{}
}

// NslaslicenseResource defines the resource implementation.
type NslaslicenseResource struct {
	client *service.NitroClient
}

func (r *NslaslicenseResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NslaslicenseResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nslaslicense"
}

func (r *NslaslicenseResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NslaslicenseResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NslaslicenseResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nslaslicense resource")
	nslaslicense := nslaslicenseGetThePayloadFromthePlan(ctx, &data)

	// Action-only resource: NITRO exposes only the `apply` verb
	// (POST ?action=apply). NOTE: applying a LAS license is DISRUPTIVE /
	// non-idempotent on the appliance.
	err := r.client.ActOnResource(service.Nslaslicense.Type(), &nslaslicense, "apply")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nslaslicense, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created nslaslicense resource")

	// Synthetic ID derived from the filename (Pattern 6); there is no GET
	// endpoint to read a server identifier back.
	data.Id = types.StringValue(data.Filename.ValueString())

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NslaslicenseResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NslaslicenseResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read is a no-op: NITRO exposes no get endpoint for nslaslicense, so we
	// preserve the prior state unchanged (drift detection is impossible).
	tflog.Debug(ctx, "Read is a no-op for nslaslicense (no GET endpoint on NITRO)")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NslaslicenseResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state NslaslicenseResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Update is a no-op for nslaslicense: all attributes are RequiresReplace and
	// NITRO has no update endpoint.
	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for nslaslicense; all attributes are RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NslaslicenseResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NslaslicenseResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Action-only resource: NITRO exposes no delete endpoint. Just remove from
	// Terraform state.
	tflog.Trace(ctx, "Removed nslaslicense from Terraform state (no delete endpoint on NITRO)")
}
