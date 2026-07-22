package systemautorestorefeature

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// systemautorestorefeature_enable models the NITRO systemautorestorefeature
// `?action=enable` action.
//
//   - NITRO exposes only the enable/disable actions (empty body):
//     POST /nitro/v1/config/systemautorestorefeature?action=enable
//     There is NO add/set/get/delete endpoint.
//   - This resource models ONLY the "enable" verb. Create performs the enable
//     action; Read/Update/Delete are no-ops (there is no GET endpoint to
//     reconcile against and no inverse API bound to this resource).
//   - Because there is no GET endpoint, there is NO datasource for this resource.
//
// The NITRO resource name "systemautorestorefeature" is not registered in the
// vendored service.NitroResourceType enum, so the literal name is used with
// ActOnResource.
var _ resource.Resource = &SystemautorestorefeatureEnableResource{}
var _ resource.ResourceWithConfigure = (*SystemautorestorefeatureEnableResource)(nil)

func NewSystemautorestorefeatureEnableResource() resource.Resource {
	return &SystemautorestorefeatureEnableResource{}
}

// SystemautorestorefeatureEnableResource defines the resource implementation.
type SystemautorestorefeatureEnableResource struct {
	client *service.NitroClient
}

// SystemautorestorefeatureEnableResourceModel describes the resource data model.
//
// systemautorestorefeature is a ZERO-ATTRIBUTE, ACTION-ONLY resource: the NITRO
// object exposes no read/write properties and the enable action takes an empty
// payload. The model therefore carries only the synthetic id.
type SystemautorestorefeatureEnableResourceModel struct {
	Id types.String `tfsdk:"id"`
}

func (r *SystemautorestorefeatureEnableResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_systemautorestorefeature_enable"
}

func (r *SystemautorestorefeatureEnableResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SystemautorestorefeatureEnableResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the systemautorestorefeature_enable resource.",
			},
		},
	}
}

func (r *SystemautorestorefeatureEnableResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SystemautorestorefeatureEnableResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Enabling the systemautorestorefeature (enable action)")
	payload := systemautorestorefeature_enableGetThePayloadFromthePlan(ctx, &data)

	// Action-only resource - Create maps to the case-sensitive "enable" verb.
	err := r.client.ActOnResource("systemautorestorefeature", payload, "enable")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to enable systemautorestorefeature, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Enabled systemautorestorefeature")

	// Synthetic ID - there is no GET endpoint to read back.
	data.Id = types.StringValue("systemautorestorefeature_enable")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Read is a no-op. systemautorestorefeature has no GET endpoint; there is
// nothing to reconcile.
func (r *SystemautorestorefeatureEnableResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SystemautorestorefeatureEnableResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for systemautorestorefeature_enable; NITRO exposes no GET endpoint")

	// Preserve prior state unchanged.
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Update is a no-op. systemautorestorefeature has no attributes and no set
// endpoint.
func (r *SystemautorestorefeatureEnableResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SystemautorestorefeatureEnableResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for systemautorestorefeature_enable; it has no attributes and no set endpoint")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Delete is a no-op. The enable action has no inverse API bound to this
// resource; Delete simply removes the resource from Terraform state.
func (r *SystemautorestorefeatureEnableResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete is a no-op for systemautorestorefeature_enable; NITRO has no inverse bound to the enable action")
}

// systemautorestorefeature_enableGetThePayloadFromthePlan builds the (empty)
// NITRO payload for the enable action. This resource has no read/write
// attributes, so the payload is an empty map.
func systemautorestorefeature_enableGetThePayloadFromthePlan(ctx context.Context, data *SystemautorestorefeatureEnableResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In systemautorestorefeature_enableGetThePayloadFromthePlan Function")
	payload := map[string]interface{}{}
	return payload
}
