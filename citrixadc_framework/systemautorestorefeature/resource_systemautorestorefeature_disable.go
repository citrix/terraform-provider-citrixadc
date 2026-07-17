package systemautorestorefeature

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

// systemautorestorefeature_disable models the NITRO systemautorestorefeature
// `?action=disable` action.
//
//   - NITRO exposes only the enable/disable actions (empty body):
//     POST /nitro/v1/config/systemautorestorefeature?action=disable
//     There is NO add/set/get/delete endpoint.
//   - This resource models ONLY the "disable" verb. Create performs the disable
//     action; Read/Update/Delete are no-ops (there is no GET endpoint to
//     reconcile against and no inverse API bound to this resource).
//   - Because there is no GET endpoint, there is NO datasource for this resource.
//
// The NITRO resource name "systemautorestorefeature" is not registered in the
// vendored service.NitroResourceType enum, so the literal name is used with
// ActOnResource.
var _ resource.Resource = &SystemautorestorefeatureDisableResource{}
var _ resource.ResourceWithConfigure = (*SystemautorestorefeatureDisableResource)(nil)
var _ resource.ResourceWithImportState = (*SystemautorestorefeatureDisableResource)(nil)

func NewSystemautorestorefeatureDisableResource() resource.Resource {
	return &SystemautorestorefeatureDisableResource{}
}

// SystemautorestorefeatureDisableResource defines the resource implementation.
type SystemautorestorefeatureDisableResource struct {
	client *service.NitroClient
}

// SystemautorestorefeatureDisableResourceModel describes the resource data model.
//
// systemautorestorefeature is a ZERO-ATTRIBUTE, ACTION-ONLY resource: the NITRO
// object exposes no read/write properties and the disable action takes an empty
// payload. The model therefore carries only the synthetic id.
type SystemautorestorefeatureDisableResourceModel struct {
	Id types.String `tfsdk:"id"`
}

func (r *SystemautorestorefeatureDisableResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SystemautorestorefeatureDisableResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_systemautorestorefeature_disable"
}

func (r *SystemautorestorefeatureDisableResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SystemautorestorefeatureDisableResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the systemautorestorefeature_disable resource.",
			},
		},
	}
}

func (r *SystemautorestorefeatureDisableResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SystemautorestorefeatureDisableResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Disabling the systemautorestorefeature (disable action)")
	payload := systemautorestorefeature_disableGetThePayloadFromthePlan(ctx, &data)

	// Action-only resource - Create maps to the case-sensitive "disable" verb.
	err := r.client.ActOnResource("systemautorestorefeature", payload, "disable")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to disable systemautorestorefeature, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Disabled systemautorestorefeature")

	// Synthetic ID - there is no GET endpoint to read back.
	data.Id = types.StringValue("systemautorestorefeature_disable")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Read is a no-op. systemautorestorefeature has no GET endpoint; there is
// nothing to reconcile.
func (r *SystemautorestorefeatureDisableResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SystemautorestorefeatureDisableResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for systemautorestorefeature_disable; NITRO exposes no GET endpoint")

	// Preserve prior state unchanged.
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Update is a no-op. systemautorestorefeature has no attributes and no set
// endpoint.
func (r *SystemautorestorefeatureDisableResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SystemautorestorefeatureDisableResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for systemautorestorefeature_disable; it has no attributes and no set endpoint")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Delete is a no-op. The disable action has no inverse API bound to this
// resource; Delete simply removes the resource from Terraform state.
func (r *SystemautorestorefeatureDisableResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete is a no-op for systemautorestorefeature_disable; NITRO has no inverse bound to the disable action")
}

// systemautorestorefeature_disableGetThePayloadFromthePlan builds the (empty)
// NITRO payload for the disable action. This resource has no read/write
// attributes, so the payload is an empty map.
func systemautorestorefeature_disableGetThePayloadFromthePlan(ctx context.Context, data *SystemautorestorefeatureDisableResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In systemautorestorefeature_disableGetThePayloadFromthePlan Function")
	payload := map[string]interface{}{}
	return payload
}
