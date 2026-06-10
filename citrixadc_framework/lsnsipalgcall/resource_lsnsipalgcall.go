package lsnsipalgcall

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
var _ resource.Resource = &LsnsipalgcallResource{}
var _ resource.ResourceWithConfigure = (*LsnsipalgcallResource)(nil)
var _ resource.ResourceWithImportState = (*LsnsipalgcallResource)(nil)

func NewLsnsipalgcallResource() resource.Resource {
	return &LsnsipalgcallResource{}
}

// LsnsipalgcallResource defines the resource implementation.
type LsnsipalgcallResource struct {
	client *service.NitroClient
}

func (r *LsnsipalgcallResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LsnsipalgcallResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lsnsipalgcall"
}

func (r *LsnsipalgcallResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

// Create performs the flush action. lsnsipalgcall is an ACTION-ONLY runtime
// resource: NITRO exposes only flush (POST ?action=flush), get/get-byname, and
// count. There is no add/update/delete. The flush action terminates the SIP ALG
// call identified by callid.
func (r *LsnsipalgcallResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LsnsipalgcallResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Flushing lsnsipalgcall resource")
	lsnsipalgcall := lsnsipalgcallGetThePayloadFromthePlan(ctx, &data)

	// Action-only resource - perform the flush action
	err := r.client.ActOnResource(service.Lsnsipalgcall.Type(), &lsnsipalgcall, "flush")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to flush lsnsipalgcall, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Flushed lsnsipalgcall resource")

	// Set ID for the resource (synthetic, from callid). Set once here.
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Callid.ValueString()))

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Read is a no-op for this action-only runtime resource. The flushed SIP ALG
// call is transient runtime state; a missing call after flush is expected and
// must not be treated as out-of-band deletion. Preserve the prior state.
func (r *LsnsipalgcallResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LsnsipalgcallResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for lsnsipalgcall; action-only flush resource, state preserved")

	// Save unchanged data back into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Update is a no-op. All attributes are RequiresReplace and NITRO exposes no
// update endpoint for this action-only resource.
func (r *LsnsipalgcallResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state LsnsipalgcallResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Update is a no-op for lsnsipalgcall; action-only flush resource, all attributes are RequiresReplace")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Delete is state-only. NITRO has no delete endpoint for this action-only
// resource; removing it from state simply forgets the prior flush invocation.
func (r *LsnsipalgcallResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LsnsipalgcallResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Delete is a state-only removal for lsnsipalgcall; action-only flush resource has no delete endpoint")
}
