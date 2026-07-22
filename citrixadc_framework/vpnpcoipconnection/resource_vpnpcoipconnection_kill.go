package vpnpcoipconnection

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
var _ resource.Resource = &VpnpcoipconnectionKillResource{}
var _ resource.ResourceWithConfigure = (*VpnpcoipconnectionKillResource)(nil)

func NewVpnpcoipconnectionKillResource() resource.Resource {
	return &VpnpcoipconnectionKillResource{}
}

// VpnpcoipconnectionKillResource models the NITRO vpnpcoipconnection
// `?action=kill` action. kill is a one-shot side-effect action (terminates
// active PCoIP connections) with no GET endpoint, no add/update/set endpoint and
// no inverse API, so Read/Update/Delete are no-ops. The kill payload carries the
// optional filter attributes all and username.
type VpnpcoipconnectionKillResource struct {
	client *service.NitroClient
}

// VpnpcoipconnectionKillResourceModel describes the resource data model.
// nodeid is intentionally absent: per the NITRO doc it is a GET-only filter
// argument, not a kill-payload property (Pattern 15).
type VpnpcoipconnectionKillResourceModel struct {
	Id       types.String `tfsdk:"id"`
	All      types.Bool   `tfsdk:"all"`
	Username types.String `tfsdk:"username"`
}

func (r *VpnpcoipconnectionKillResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnpcoipconnection_kill"
}

func (r *VpnpcoipconnectionKillResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnpcoipconnectionKillResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnpcoipconnection_kill resource.",
			},
			// User-facing kill arguments. Read is a no-op (kill is an action,
			// the killed connection is not a persistent managed object), so
			// these must NOT be Computed or Terraform reports an unknown value
			// after apply (Pattern 13 schema-flag implication).
			"all": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "All active pcoip connections.",
			},
			"username": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "User name for the PCOIP connections.",
			},
		},
	}
}

func (r *VpnpcoipconnectionKillResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnpcoipconnectionKillResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Killing vpnpcoipconnection (action-only resource)")
	payload := vpnpcoipconnection_killGetThePayloadFromthePlan(ctx, &data)

	// kill is a POST ?action=kill action (Pattern 1). There is no add endpoint;
	// AddResource/UpdateUnnamedResource would target a nonexistent verb.
	err := r.client.ActOnResource(service.Vpnpcoipconnection.Type(), payload, "kill")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to kill vpnpcoipconnection, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Killed vpnpcoipconnection")

	// Synthetic ID for the action-only resource; keeps Read/Delete no-ops
	// addressable by Terraform.
	data.Id = types.StringValue("vpnpcoipconnection_kill")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnpcoipconnectionKillResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// kill is a one-shot action. NITRO has no GET endpoint that reports
	// kill-state, so Read is a pure preserve-state no-op.
	var data VpnpcoipconnectionKillResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for vpnpcoipconnection_kill; NITRO has no query endpoint for kill state")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnpcoipconnectionKillResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO has no update endpoint for kill; every schema attribute is
	// RequiresReplace, so Terraform never invokes Update for a real change.
	var data, state VpnpcoipconnectionKillResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for vpnpcoipconnection_kill; NITRO has no update endpoint and all attributes are RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnpcoipconnectionKillResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// kill is a one-shot side-effect action. There is no inverse NITRO API
	// (no "un-kill"). Delete simply removes the resource from Terraform state.
	tflog.Debug(ctx, "Delete is a no-op for vpnpcoipconnection_kill; NITRO has no inverse of the kill action")
}

// vpnpcoipconnection_killGetThePayloadFromthePlan builds the body for the kill
// action. The vendored vpn.Vpnpcoipconnection struct carries read-only fields
// and nodeid (a GET-only filter); build a map containing ONLY the kill arguments
// that are set so the action body never includes invalid arguments (Pattern 3 + 15).
func vpnpcoipconnection_killGetThePayloadFromthePlan(ctx context.Context, data *VpnpcoipconnectionKillResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In vpnpcoipconnection_killGetThePayloadFromthePlan Function")

	vpnpcoipconnection := map[string]interface{}{}
	if !data.All.IsNull() && !data.All.IsUnknown() {
		vpnpcoipconnection["all"] = data.All.ValueBool()
	}
	if !data.Username.IsNull() && !data.Username.IsUnknown() {
		vpnpcoipconnection["username"] = data.Username.ValueString()
	}

	return vpnpcoipconnection
}
