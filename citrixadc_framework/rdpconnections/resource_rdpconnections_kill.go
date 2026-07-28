package rdpconnections

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &RdpconnectionsKillResource{}
var _ resource.ResourceWithConfigure = (*RdpconnectionsKillResource)(nil)

func NewRdpconnectionsKillResource() resource.Resource {
	return &RdpconnectionsKillResource{}
}

// RdpconnectionsKillResource defines the resource implementation.
//
// This resource models the NITRO rdpconnections `?action=kill` action (POST).
// kill is a one-shot side-effect action: NITRO exposes no add, no set/update,
// and no delete endpoint for it, so Read/Update/Delete are no-ops. The kill
// payload carries the optional selectors username and all. The read-only
// telemetry (endpointip, endpointport, targetip, targetport, peid) belongs to
// the get/count read side and lives on the citrixadc_rdpconnections datasource.
type RdpconnectionsKillResource struct {
	client *service.NitroClient
}

// RdpconnectionsKillResourceModel describes the resource data model. Only the
// two kill selectors (username, all) plus a synthetic id are modelled here.
type RdpconnectionsKillResourceModel struct {
	Id       types.String `tfsdk:"id"`
	All      types.Bool   `tfsdk:"all"`
	Username types.String `tfsdk:"username"`
}

func (r *RdpconnectionsKillResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rdpconnections_kill"
}

func (r *RdpconnectionsKillResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *RdpconnectionsKillResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the rdpconnections_kill resource.",
			},
			// kill selectors. Both Optional (one-of-or-none semantics, not
			// enforced as a group). Not Computed: Read is a no-op, so no server
			// value ever resolves them. RequiresReplace: any change re-fires the
			// kill.
			"all": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Terminate all active rdpconnections.",
			},
			"username": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "User name for which to display connections.",
			},
		},
	}
}

func (r *RdpconnectionsKillResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data RdpconnectionsKillResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating rdpconnections_kill resource (firing ?action=kill)")
	payload := rdpconnections_killGetThePayloadFromthePlan(ctx, &data)

	// rdpconnections has no add/set endpoint. Fire the kill action.
	// NITRO verb is case-sensitive: "kill" (POST ?action=kill).
	err := r.client.ActOnResource(service.Rdpconnections.Type(), payload, "kill")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to kill rdpconnections, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Killed rdpconnections")

	// Synthetic ID for the action-only resource; keeps Read/Delete no-ops
	// addressable by Terraform.
	data.Id = types.StringValue("rdpconnections_kill")

	// Save data into Terraform state (no Read: transient table, no GET-by-key).
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Read is a no-op. rdpconnections kill is a one-shot action with no GET-by-key
// endpoint; the killed connections are not a persistent object, so there is
// nothing to reconcile. Preserve prior state unchanged.
func (r *RdpconnectionsKillResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data RdpconnectionsKillResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for rdpconnections_kill (action-only kill, transient table)")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Update is a no-op. All attributes (username, all) are RequiresReplace, so
// Terraform never routes a change through Update. There is no set endpoint.
func (r *RdpconnectionsKillResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state RdpconnectionsKillResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for rdpconnections_kill; all attributes are RequiresReplace and there is no set endpoint")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Delete is a no-op. rdpconnections has no delete endpoint; the kill action is
// fire-and-forget. Removing the resource simply drops it from state.
func (r *RdpconnectionsKillResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete is a no-op for rdpconnections_kill (no inverse of the kill action on NITRO side)")
}

// rdpconnections_killGetThePayloadFromthePlan builds the ?action=kill payload.
// Null selectors are omitted so a bare kill and all=true are both expressible.
// Returns map[string]interface{} because the kill action needs only the two
// selectors, not the full vendored struct (mirrors the prior impl).
func rdpconnections_killGetThePayloadFromthePlan(ctx context.Context, data *RdpconnectionsKillResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In rdpconnections_killGetThePayloadFromthePlan Function")

	payload := map[string]interface{}{}
	if !data.Username.IsNull() && !data.Username.IsUnknown() {
		payload["username"] = data.Username.ValueString()
	}
	if !data.All.IsNull() && !data.All.IsUnknown() {
		payload["all"] = data.All.ValueBool()
	}

	return payload
}
