package nssourceroutecachetable

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// nssourceroutecachetable is an ACTION-ONLY resource.
//
//   - NITRO exposes the flush action (POST /config/nssourceroutecachetable?action=flush)
//     plus a keyless get(all) and count. There is no add/set/delete endpoint.
//   - Create performs the flush action. Read/Update/Delete are no-ops: flushing
//     the cache table has no persistent object to reconcile or remove.
//   - There are no read/write attributes. Use the datasource to read back the
//     (read-only) get(all) attributes.
var _ resource.Resource = &NssourceroutecachetableResource{}
var _ resource.ResourceWithConfigure = (*NssourceroutecachetableResource)(nil)
var _ resource.ResourceWithImportState = (*NssourceroutecachetableResource)(nil)

func NewNssourceroutecachetableResource() resource.Resource {
	return &NssourceroutecachetableResource{}
}

// NssourceroutecachetableResource defines the resource implementation.
type NssourceroutecachetableResource struct {
	client *service.NitroClient
}

func (r *NssourceroutecachetableResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NssourceroutecachetableResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nssourceroutecachetable"
}

func (r *NssourceroutecachetableResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NssourceroutecachetableResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NssourceroutecachetableResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating (flush action) nssourceroutecachetable resource")
	nssourceroutecachetable := nssourceroutecachetableGetThePayloadFromthePlan(ctx, &data)

	// flush is a POST ?action=flush action. There is no add endpoint.
	err := r.client.ActOnResource(nssourceroutecachetableResourceType, &nssourceroutecachetable, "flush")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to flush nssourceroutecachetable, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Flushed nssourceroutecachetable")

	// Synthetic ID - there is no addressable object to read back.
	data.Id = types.StringValue("nssourceroutecachetable-config")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Read is a no-op. Use the datasource for get(all).
func (r *NssourceroutecachetableResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NssourceroutecachetableResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for nssourceroutecachetable; use the datasource for get(all)")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Update is a no-op. nssourceroutecachetable has no read/write attributes and no set endpoint.
func (r *NssourceroutecachetableResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state NssourceroutecachetableResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id

	tflog.Debug(ctx, "Update is a no-op for nssourceroutecachetable; it has no read/write attributes")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Delete is a no-op. nssourceroutecachetable has no delete endpoint; the flush
// action has no persistent object to remove.
func (r *NssourceroutecachetableResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete is a no-op for nssourceroutecachetable; NITRO has no delete endpoint")
}
