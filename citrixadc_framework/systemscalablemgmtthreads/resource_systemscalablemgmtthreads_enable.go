package systemscalablemgmtthreads

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/system"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SystemscalablemgmtthreadsEnableResource{}
var _ resource.ResourceWithConfigure = (*SystemscalablemgmtthreadsEnableResource)(nil)

func NewSystemscalablemgmtthreadsEnableResource() resource.Resource {
	return &SystemscalablemgmtthreadsEnableResource{}
}

// SystemscalablemgmtthreadsEnableResource defines the resource implementation.
type SystemscalablemgmtthreadsEnableResource struct {
	client *service.NitroClient
}

// SystemscalablemgmtthreadsEnableResourceModel describes the resource data model.
//
// systemscalablemgmtthreads is a NITRO object that supports multiple actions
// (enable / disable) plus a get. Mirroring the appfwlearningdata package
// (appfwlearningdata_export / appfwlearningdata_reset), each action is modelled as
// its own action-only resource. This resource wraps the `?action=enable` action,
// which turns ON the Scalable Management Threads feature.
//
// The enable action carries an EMPTY payload ({"systemscalablemgmtthreads":{}}):
// nodeid (the only read/write property) is a GET-only filter argument and is
// rejected in the action payload (NITRO errorcode 278 "Invalid argument
// [nodeid]"), so it is intentionally excluded here. There is no per-action GET, so
// Read/Update are no-ops; the live feature state is queryable via the
// citrixadc_systemscalablemgmtthreads data source. Delete is a state-only removal
// (the inverse of enable is the separate citrixadc_systemscalablemgmtthreads_disable
// resource).
type SystemscalablemgmtthreadsEnableResourceModel struct {
	Id types.String `tfsdk:"id"`
}

func (r *SystemscalablemgmtthreadsEnableResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_systemscalablemgmtthreads_enable"
}

func (r *SystemscalablemgmtthreadsEnableResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SystemscalablemgmtthreadsEnableResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the systemscalablemgmtthreads_enable resource.",
			},
			// The enable action accepts no attributes: the NITRO enable payload is
			// empty and nodeid is rejected in the action body (errorcode 278).
		},
	}
}

func (r *SystemscalablemgmtthreadsEnableResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SystemscalablemgmtthreadsEnableResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Enabling systemscalablemgmtthreads (action-only resource)")
	payload := systemscalablemgmtthreads_enableGetThePayloadFromthePlan(ctx, &data)

	// NITRO exposes enable as POST ?action=enable. Verb casing matches the URL.
	err := r.client.ActOnResource(service.Systemscalablemgmtthreads.Type(), &payload, "enable")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to enable systemscalablemgmtthreads, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Triggered systemscalablemgmtthreads enable")

	// Synthetic ID for the action-only resource; keeps Read/Delete no-ops
	// addressable by Terraform.
	data.Id = types.StringValue("systemscalablemgmtthreads_enable")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemscalablemgmtthreadsEnableResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// enable is a one-shot action. There is no per-action GET endpoint, so Read is
	// a pure preserve-state no-op. Query live state via the data source instead.
	var data SystemscalablemgmtthreadsEnableResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for systemscalablemgmtthreads_enable; NITRO has no per-action query endpoint")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemscalablemgmtthreadsEnableResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO has no update endpoint for the enable action; there are no config
	// attributes to change, so Terraform never invokes Update for a real change.
	var data, state SystemscalablemgmtthreadsEnableResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for systemscalablemgmtthreads_enable; NITRO has no update endpoint")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemscalablemgmtthreadsEnableResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// enable is a one-shot action. Its inverse is the separate
	// citrixadc_systemscalablemgmtthreads_disable resource, so Delete simply removes
	// the resource from Terraform state (it does NOT auto-disable the feature).
	tflog.Debug(ctx, "Delete is a no-op for systemscalablemgmtthreads_enable; use citrixadc_systemscalablemgmtthreads_disable to turn the feature off")
}

func systemscalablemgmtthreads_enableGetThePayloadFromthePlan(ctx context.Context, data *SystemscalablemgmtthreadsEnableResourceModel) system.Systemscalablemgmtthreads {
	tflog.Debug(ctx, "In systemscalablemgmtthreads_enableGetThePayloadFromthePlan Function")

	// The enable action carries an empty payload ({"systemscalablemgmtthreads":{}});
	// nodeid is a GET-only filter and is rejected in the action body.
	systemscalablemgmtthreads := system.Systemscalablemgmtthreads{}

	return systemscalablemgmtthreads
}
