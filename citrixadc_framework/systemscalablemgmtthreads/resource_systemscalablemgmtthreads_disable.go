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
var _ resource.Resource = &SystemscalablemgmtthreadsDisableResource{}
var _ resource.ResourceWithConfigure = (*SystemscalablemgmtthreadsDisableResource)(nil)

func NewSystemscalablemgmtthreadsDisableResource() resource.Resource {
	return &SystemscalablemgmtthreadsDisableResource{}
}

// SystemscalablemgmtthreadsDisableResource defines the resource implementation.
type SystemscalablemgmtthreadsDisableResource struct {
	client *service.NitroClient
}

// SystemscalablemgmtthreadsDisableResourceModel describes the resource data model.
//
// systemscalablemgmtthreads is a NITRO object that supports multiple actions
// (enable / disable) plus a get. Mirroring the appfwlearningdata package
// (appfwlearningdata_export / appfwlearningdata_reset), each action is modelled as
// its own action-only resource. This resource wraps the `?action=disable` action,
// which turns OFF the Scalable Management Threads feature.
//
// The disable action carries an EMPTY payload ({"systemscalablemgmtthreads":{}}):
// nodeid (the only read/write property) is a GET-only filter argument and is
// rejected in the action payload (NITRO errorcode 278 "Invalid argument
// [nodeid]"), so it is intentionally excluded here. There is no per-action GET, so
// Read/Update are no-ops; the live feature state is queryable via the
// citrixadc_systemscalablemgmtthreads data source. Delete is a state-only removal
// (the inverse of disable is the separate citrixadc_systemscalablemgmtthreads_enable
// resource).
type SystemscalablemgmtthreadsDisableResourceModel struct {
	Id types.String `tfsdk:"id"`
}

func (r *SystemscalablemgmtthreadsDisableResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_systemscalablemgmtthreads_disable"
}

func (r *SystemscalablemgmtthreadsDisableResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SystemscalablemgmtthreadsDisableResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the systemscalablemgmtthreads_disable resource.",
			},
			// The disable action accepts no attributes: the NITRO disable payload is
			// empty and nodeid is rejected in the action body (errorcode 278).
		},
	}
}

func (r *SystemscalablemgmtthreadsDisableResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SystemscalablemgmtthreadsDisableResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Disabling systemscalablemgmtthreads (action-only resource)")
	payload := systemscalablemgmtthreads_disableGetThePayloadFromthePlan(ctx, &data)

	// NITRO exposes disable as POST ?action=disable. Verb casing matches the URL.
	err := r.client.ActOnResource(service.Systemscalablemgmtthreads.Type(), &payload, "disable")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to disable systemscalablemgmtthreads, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Triggered systemscalablemgmtthreads disable")

	// Synthetic ID for the action-only resource; keeps Read/Delete no-ops
	// addressable by Terraform.
	data.Id = types.StringValue("systemscalablemgmtthreads_disable")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemscalablemgmtthreadsDisableResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// disable is a one-shot action. There is no per-action GET endpoint, so Read is
	// a pure preserve-state no-op. Query live state via the data source instead.
	var data SystemscalablemgmtthreadsDisableResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for systemscalablemgmtthreads_disable; NITRO has no per-action query endpoint")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemscalablemgmtthreadsDisableResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO has no update endpoint for the disable action; there are no config
	// attributes to change, so Terraform never invokes Update for a real change.
	var data, state SystemscalablemgmtthreadsDisableResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for systemscalablemgmtthreads_disable; NITRO has no update endpoint")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemscalablemgmtthreadsDisableResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// disable is a one-shot action. Its inverse is the separate
	// citrixadc_systemscalablemgmtthreads_enable resource, so Delete simply removes
	// the resource from Terraform state (it does NOT auto-enable the feature).
	tflog.Debug(ctx, "Delete is a no-op for systemscalablemgmtthreads_disable; use citrixadc_systemscalablemgmtthreads_enable to turn the feature on")
}

func systemscalablemgmtthreads_disableGetThePayloadFromthePlan(ctx context.Context, data *SystemscalablemgmtthreadsDisableResourceModel) system.Systemscalablemgmtthreads {
	tflog.Debug(ctx, "In systemscalablemgmtthreads_disableGetThePayloadFromthePlan Function")

	// The disable action carries an empty payload ({"systemscalablemgmtthreads":{}});
	// nodeid is a GET-only filter and is rejected in the action body.
	systemscalablemgmtthreads := system.Systemscalablemgmtthreads{}

	return systemscalablemgmtthreads
}
