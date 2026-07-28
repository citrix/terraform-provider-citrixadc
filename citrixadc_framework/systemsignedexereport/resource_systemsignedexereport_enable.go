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

// systemsignedexereport_enable models the NITRO systemsignedexereport
// `?action=enable` action.
//
//   - NITRO exposes only the enable/disable actions:
//     POST /nitro/v1/config/systemsignedexereport?action=enable  (empty body)
//     There is NO add/set/get/delete endpoint.
//   - This resource models ONLY the "enable" verb. Create performs the enable
//     action; Read/Update/Delete are no-ops (there is nothing to reconcile and
//     no inverse API for delete).
//   - Because there is no GET endpoint, there is NO datasource for this resource
//     and it cannot be verified by reading it back.
//   - This resource is not registered in the vendored service.NitroResourceType
//     enum, so the literal NITRO name is used with ActOnResource.
var _ resource.Resource = &SystemsignedexereportEnableResource{}
var _ resource.ResourceWithConfigure = (*SystemsignedexereportEnableResource)(nil)

func NewSystemsignedexereportEnableResource() resource.Resource {
	return &SystemsignedexereportEnableResource{}
}

// SystemsignedexereportEnableResource defines the resource implementation.
type SystemsignedexereportEnableResource struct {
	client *service.NitroClient
}

// SystemsignedexereportEnableResourceModel describes the resource data model.
//
// The enable action accepts an empty payload (per NITRO doc and CLI: zero
// arguments), so the model carries only the synthetic id.
type SystemsignedexereportEnableResourceModel struct {
	Id types.String `tfsdk:"id"`
}

func (r *SystemsignedexereportEnableResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_systemsignedexereport_enable"
}

func (r *SystemsignedexereportEnableResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SystemsignedexereportEnableResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the systemsignedexereport_enable resource.",
			},
		},
	}
}

func (r *SystemsignedexereportEnableResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SystemsignedexereportEnableResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Enabling the systemsignedexereport (enable action)")
	payload := systemsignedexereport_enableGetThePayloadFromthePlan(ctx, &data)

	// Action-only resource - Create maps to the "enable" action. The literal
	// NITRO type name is used because this resource is not in the vendored enum.
	err := r.client.ActOnResource("systemsignedexereport", &payload, "enable")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to enable systemsignedexereport, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Enabled systemsignedexereport")

	// Synthetic ID - there is no GET endpoint to read back.
	data.Id = types.StringValue("systemsignedexereport_enable")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Read is a no-op. systemsignedexereport has no GET endpoint; there is nothing
// to reconcile.
func (r *SystemsignedexereportEnableResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SystemsignedexereportEnableResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for systemsignedexereport_enable; NITRO exposes no GET endpoint")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Update is a no-op. systemsignedexereport_enable has no attributes and no set
// endpoint.
func (r *SystemsignedexereportEnableResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SystemsignedexereportEnableResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Update is a no-op for systemsignedexereport_enable; it has no attributes and no set endpoint")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Delete is a no-op. The enable action has no inverse NITRO API on this
// resource; Delete simply removes the resource from Terraform state.
func (r *SystemsignedexereportEnableResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete is a no-op for systemsignedexereport_enable; NITRO has no inverse of the enable action on this resource")
}

// systemsignedexereport_enableGetThePayloadFromthePlan builds the (empty) NITRO
// payload for the enable action. This action has no read/write attributes, so
// the payload is an empty map.
func systemsignedexereport_enableGetThePayloadFromthePlan(ctx context.Context, data *SystemsignedexereportEnableResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In systemsignedexereport_enableGetThePayloadFromthePlan Function")
	payload := map[string]interface{}{}
	return payload
}
