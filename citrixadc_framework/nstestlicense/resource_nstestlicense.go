package nstestlicense

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// nstestlicense is an ACTION-ONLY resource.
//
//   - NITRO exposes the apply action (POST /config/nstestlicense?action=apply)
//     plus a keyless get(all). There is no add/set/delete endpoint.
//   - Create performs the apply action (applies a test/eval license, which is
//     potentially disruptive). Read/Update/Delete are no-ops.
//   - There are no read/write attributes. Use the datasource to read back the
//     (read-only) get(all) attributes.
var _ resource.Resource = &NstestlicenseResource{}
var _ resource.ResourceWithConfigure = (*NstestlicenseResource)(nil)
var _ resource.ResourceWithImportState = (*NstestlicenseResource)(nil)

func NewNstestlicenseResource() resource.Resource {
	return &NstestlicenseResource{}
}

// NstestlicenseResource defines the resource implementation.
type NstestlicenseResource struct {
	client *service.NitroClient
}

func (r *NstestlicenseResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NstestlicenseResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nstestlicense"
}

func (r *NstestlicenseResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NstestlicenseResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NstestlicenseResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating (apply action) nstestlicense resource")
	nstestlicense := nstestlicenseGetThePayloadFromthePlan(ctx, &data)

	// apply is a POST ?action=apply action. There is no add endpoint.
	err := r.client.ActOnResource(nstestlicenseResourceType, &nstestlicense, "apply")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to apply nstestlicense, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Applied nstestlicense")

	// Synthetic ID - there is no addressable object to read back.
	data.Id = types.StringValue("nstestlicense-config")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Read is a no-op. Use the datasource for get(all).
func (r *NstestlicenseResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NstestlicenseResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for nstestlicense; use the datasource for get(all)")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Update is a no-op. nstestlicense has no read/write attributes and no set endpoint.
func (r *NstestlicenseResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state NstestlicenseResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id

	tflog.Debug(ctx, "Update is a no-op for nstestlicense; it has no read/write attributes")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Delete is a no-op. nstestlicense has no delete endpoint; the apply action has
// no persistent object to remove.
func (r *NstestlicenseResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete is a no-op for nstestlicense; NITRO has no delete endpoint")
}
