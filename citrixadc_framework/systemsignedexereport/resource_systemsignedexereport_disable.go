package systemsignedexereport

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// systemsignedexereport_disable models the NITRO systemsignedexereport
// `?action=disable` action.
//
//   - NITRO exposes only the enable/disable actions:
//     POST /nitro/v1/config/systemsignedexereport?action=disable  (empty body)
//     There is NO add/set/get/delete endpoint.
//   - This resource models ONLY the "disable" verb. Create performs the disable
//     action; Read/Update/Delete are no-ops (there is nothing to reconcile and
//     no inverse API for delete).
//   - Because there is no GET endpoint, there is NO datasource for this resource
//     and it cannot be verified by reading it back.
//   - This resource is not registered in the vendored service.NitroResourceType
//     enum, so the literal NITRO name is used with ActOnResource.
var _ resource.Resource = &SystemsignedexereportDisableResource{}
var _ resource.ResourceWithConfigure = (*SystemsignedexereportDisableResource)(nil)

func NewSystemsignedexereportDisableResource() resource.Resource {
	return &SystemsignedexereportDisableResource{}
}

// SystemsignedexereportDisableResource defines the resource implementation.
type SystemsignedexereportDisableResource struct {
	client *service.NitroClient
}

// SystemsignedexereportDisableResourceModel describes the resource data model.
//
// The disable action accepts an empty payload (per NITRO doc and CLI: zero
// arguments), so the model carries only the synthetic id.
type SystemsignedexereportDisableResourceModel struct {
	Id types.String `tfsdk:"id"`
}

func (r *SystemsignedexereportDisableResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_systemsignedexereport_disable"
}

func (r *SystemsignedexereportDisableResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SystemsignedexereportDisableResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the systemsignedexereport_disable resource.",
			},
		},
	}
}

func (r *SystemsignedexereportDisableResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SystemsignedexereportDisableResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Disabling the systemsignedexereport (disable action)")
	payload := systemsignedexereport_disableGetThePayloadFromthePlan(ctx, &data)

	// Action-only resource - Create maps to the "disable" action. The literal
	// NITRO type name is used because this resource is not in the vendored enum.
	err := r.client.ActOnResource("systemsignedexereport", &payload, "disable")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to disable systemsignedexereport, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Disabled systemsignedexereport")

	// Synthetic ID - there is no GET endpoint to read back.
	data.Id = types.StringValue("systemsignedexereport_disable")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Read is a no-op. systemsignedexereport has no GET endpoint; there is nothing
// to reconcile.
func (r *SystemsignedexereportDisableResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SystemsignedexereportDisableResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for systemsignedexereport_disable; NITRO exposes no GET endpoint")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Update is a no-op. systemsignedexereport_disable has no attributes and no set
// endpoint.
func (r *SystemsignedexereportDisableResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SystemsignedexereportDisableResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Update is a no-op for systemsignedexereport_disable; it has no attributes and no set endpoint")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Delete is a no-op. The disable action has no inverse NITRO API on this
// resource; Delete simply removes the resource from Terraform state.
func (r *SystemsignedexereportDisableResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete is a no-op for systemsignedexereport_disable; NITRO has no inverse of the disable action on this resource")
}

// systemsignedexereport_disableGetThePayloadFromthePlan builds the (empty) NITRO
// payload for the disable action. This action has no read/write attributes, so
// the payload is an empty map.
func systemsignedexereport_disableGetThePayloadFromthePlan(ctx context.Context, data *SystemsignedexereportDisableResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In systemsignedexereport_disableGetThePayloadFromthePlan Function")
	payload := map[string]interface{}{}
	return payload
}
