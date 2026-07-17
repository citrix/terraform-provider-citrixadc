package nstestlicense

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// nstestlicenseResourceType is the NITRO resource-type string. nstestlicense is
// not present in the adc-nitro-go service enum, so the type is declared here.
//
// It is relocated here from the (now deleted) resource_schema.go because the
// preserved datasource (datasource_nstestlicense.go) references it in its
// FindResource call.
const nstestlicenseResourceType = "nstestlicense"

// nstestlicense_apply is an ACTION-ONLY resource.
//
//   - NITRO exposes the apply action (POST /config/nstestlicense?action=apply)
//     with an empty payload ({"nstestlicense":{}}). There is no add/set/delete
//     endpoint.
//   - Create performs the apply action (applies a test/eval license, which is
//     potentially disruptive). Read/Update/Delete are no-ops.
//   - The apply action takes no arguments (Usage: apply ns testlicense); the
//     model carries only the synthetic id. Use the citrixadc_nstestlicense
//     datasource to read back the (read-only) get(all) attributes.
var _ resource.Resource = &NstestlicenseApplyResource{}
var _ resource.ResourceWithConfigure = (*NstestlicenseApplyResource)(nil)
var _ resource.ResourceWithImportState = (*NstestlicenseApplyResource)(nil)

func NewNstestlicenseApplyResource() resource.Resource {
	return &NstestlicenseApplyResource{}
}

// NstestlicenseApplyResource defines the resource implementation.
type NstestlicenseApplyResource struct {
	client *service.NitroClient
}

// NstestlicenseApplyResourceModel describes the resource data model.
//
// The apply action is a ZERO-ATTRIBUTE action; the model carries only the
// synthetic id.
type NstestlicenseApplyResourceModel struct {
	Id types.String `tfsdk:"id"`
}

func (r *NstestlicenseApplyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NstestlicenseApplyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nstestlicense_apply"
}

func (r *NstestlicenseApplyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NstestlicenseApplyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nstestlicense_apply resource. It is a synthetic value (nstestlicense_apply).",
			},
		},
	}
}

func (r *NstestlicenseApplyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NstestlicenseApplyResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Applying nstestlicense (action-only resource)")
	payload := nstestlicense_applyGetThePayloadFromthePlan(ctx, &data)

	// NITRO exposes apply as POST ?action=apply. Use ActOnResource with the
	// case-sensitive "apply" verb (lower-case per the NITRO URL).
	err := r.client.ActOnResource(nstestlicenseResourceType, &payload, "apply")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to apply nstestlicense, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Applied nstestlicense")

	// Synthetic ID for the action-only resource; keeps Read/Delete no-ops
	// addressable by Terraform.
	data.Id = types.StringValue("nstestlicense_apply")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NstestlicenseApplyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// apply is a one-shot action. NITRO has no GET endpoint that reports
	// apply-state, so Read is a pure preserve-state no-op.
	var data NstestlicenseApplyResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for nstestlicense_apply; use the citrixadc_nstestlicense datasource for get(all)")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NstestlicenseApplyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO has no update endpoint for apply; the model has no writable
	// attributes, so Terraform never invokes Update for a real change.
	var data, state NstestlicenseApplyResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for nstestlicense_apply; it has no read/write attributes")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NstestlicenseApplyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// apply is a one-shot side-effect action. There is no inverse NITRO API
	// (no "un-apply"). Delete simply removes the resource from Terraform state.
	tflog.Debug(ctx, "Delete is a no-op for nstestlicense_apply; NITRO has no inverse of the apply action")
}

// nstestlicense_applyGetThePayloadFromthePlan builds the (empty) NITRO payload
// for the apply action. The apply action has no read/write attributes.
func nstestlicense_applyGetThePayloadFromthePlan(ctx context.Context, data *NstestlicenseApplyResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In nstestlicense_applyGetThePayloadFromthePlan Function")
	nstestlicense := make(map[string]interface{})
	return nstestlicense
}
