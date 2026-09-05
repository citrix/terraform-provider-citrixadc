package ssldynamicclientcertcache

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SsldynamicclientcertcacheFlushResource models the NITRO
// ssldynamicclientcertcache `?action=flush` action.
//
//   - NITRO exposes flush as POST /config/ssldynamicclientcertcache?action=flush.
//     There is no add/set/delete endpoint and no GET endpoint. The backing
//     service/config/ssl.Ssldynamicclientcertcache{} struct is empty.
//   - Create performs the flush action. Read/Update/Delete are no-ops: flushing
//     the dynamic client-certificate cache has no persistent object to reconcile
//     or remove.
//   - flush accepts NO attributes (empty payload); the model carries only the
//     synthetic id.
var _ resource.Resource = &SsldynamicclientcertcacheFlushResource{}
var _ resource.ResourceWithConfigure = (*SsldynamicclientcertcacheFlushResource)(nil)
var _ resource.ResourceWithImportState = (*SsldynamicclientcertcacheFlushResource)(nil)

func NewSsldynamicclientcertcacheFlushResource() resource.Resource {
	return &SsldynamicclientcertcacheFlushResource{}
}

// SsldynamicclientcertcacheFlushResource defines the resource implementation.
type SsldynamicclientcertcacheFlushResource struct {
	client *service.NitroClient
}

func (r *SsldynamicclientcertcacheFlushResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SsldynamicclientcertcacheFlushResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ssldynamicclientcertcache_flush"
}

func (r *SsldynamicclientcertcacheFlushResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SsldynamicclientcertcacheFlushResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SsldynamicclientcertcacheFlushResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Flushing ssldynamicclientcertcache (action-only resource)")
	payload := ssldynamicclientcertcacheFlushGetThePayloadFromthePlan(ctx, &data)

	// NITRO exposes flush as POST ?action=flush. Use ActOnResource with the
	// case-sensitive "flush" verb (lower-case per the NITRO URL).
	err := r.client.ActOnResource(service.Ssldynamicclientcertcache.Type(), &payload, "flush")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to flush ssldynamicclientcertcache, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Flushed ssldynamicclientcertcache")

	// Synthetic ID for the action-only resource; keeps Read/Delete no-ops
	// addressable by Terraform.
	data.Id = types.StringValue("ssldynamicclientcertcache")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SsldynamicclientcertcacheFlushResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// flush is a one-shot action. NITRO has no GET endpoint that reports
	// flush-state, so Read is a pure preserve-state no-op.
	var data SsldynamicclientcertcacheFlushResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for ssldynamicclientcertcache_flush; NITRO has no query endpoint for flush state")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SsldynamicclientcertcacheFlushResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO has no update endpoint for flush; the resource has no read/write
	// attributes, so Terraform never invokes Update for a real change.
	var data, state SsldynamicclientcertcacheFlushResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for ssldynamicclientcertcache_flush; it has no read/write attributes")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SsldynamicclientcertcacheFlushResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// flush is a one-shot side-effect action. There is no inverse NITRO API
	// (no "un-flush"). Delete simply removes the resource from Terraform state.
	tflog.Debug(ctx, "Delete is a no-op for ssldynamicclientcertcache_flush; NITRO has no inverse of the flush action")
}
