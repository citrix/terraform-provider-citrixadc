package vpnicaconnection

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
var _ resource.Resource = &VpnicaconnectionKillResource{}
var _ resource.ResourceWithConfigure = (*VpnicaconnectionKillResource)(nil)

func NewVpnicaconnectionKillResource() resource.Resource {
	return &VpnicaconnectionKillResource{}
}

// VpnicaconnectionKillResource models the NITRO vpnicaconnection `?action=kill`
// runtime action. kill is a one-shot side-effect action that terminates active
// ICA connections; NITRO exposes no add/update/delete endpoint and no GET-backed
// object keyed by these inputs, so Read/Update/Delete are no-ops.
type VpnicaconnectionKillResource struct {
	client *service.NitroClient
}

// VpnicaconnectionKillResourceModel describes the resource data model.
// nodeid is intentionally absent: per the NITRO doc it is a GET-only filter
// argument, not a kill-payload property (Pattern 15).
type VpnicaconnectionKillResourceModel struct {
	Id         types.String `tfsdk:"id"`
	All        types.Bool   `tfsdk:"all"`
	Transproto types.String `tfsdk:"transproto"`
	Username   types.String `tfsdk:"username"`
}

func (r *VpnicaconnectionKillResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnicaconnection_kill"
}

func (r *VpnicaconnectionKillResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnicaconnectionKillResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnicaconnection_kill resource.",
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
				Description: "Terminate all active icaconnections.",
			},
			"transproto": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Transport type for the existing Existing ICA conenction.",
			},
			"username": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "User name for which ica connections needs to be terminated.",
			},
		},
	}
}

func (r *VpnicaconnectionKillResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnicaconnectionKillResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Killing vpnicaconnection (action-only resource)")
	payload := vpnicaconnection_killGetThePayloadFromthePlan(ctx, &data)

	// kill is a POST ?action=kill action (Pattern 1). There is no add endpoint;
	// AddResource/UpdateUnnamedResource would target a nonexistent verb.
	err := r.client.ActOnResource(service.Vpnicaconnection.Type(), payload, "kill")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to kill vpnicaconnection, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Killed vpnicaconnection")

	// Synthetic ID for the action-only resource; keeps Read/Delete no-ops
	// addressable by Terraform.
	data.Id = types.StringValue("vpnicaconnection_kill")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnicaconnectionKillResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnicaconnectionKillResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Read is a documented no-op. vpnicaconnection.kill terminates active ICA
	// connections; the killed connection is not a persistent object keyed by
	// these inputs, so a GET cannot stably re-resolve "this" record. Re-fetching
	// would cause perpetual drift / spurious resource removal. Preserve state.
	tflog.Debug(ctx, "Read is a no-op for vpnicaconnection_kill (kill action has no stable GET-backed object); preserving state")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnicaconnectionKillResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state VpnicaconnectionKillResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Update is a no-op: NITRO exposes no update/set endpoint for vpnicaconnection,
	// and every attribute is RequiresReplace, so Terraform never reaches here with
	// real diffs (Pattern 5).
	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for vpnicaconnection_kill; no NITRO update endpoint and all attributes are RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnicaconnectionKillResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// kill is a one-shot side-effect action. There is no inverse NITRO API
	// (no "un-kill"). Delete simply removes the resource from Terraform state.
	tflog.Debug(ctx, "Delete is a no-op for vpnicaconnection_kill; kill has no inverse NITRO endpoint")
}

// vpnicaconnection_killGetThePayloadFromthePlan builds the body for the kill action.
// The vendored vpn.Vpnicaconnection struct carries read-only fields and nodeid
// (a GET-only filter); build a map containing ONLY the kill arguments that are
// set so the action body never includes invalid arguments (Pattern 3 + 15).
func vpnicaconnection_killGetThePayloadFromthePlan(ctx context.Context, data *VpnicaconnectionKillResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In vpnicaconnection_killGetThePayloadFromthePlan Function")

	vpnicaconnection := map[string]interface{}{}
	if !data.All.IsNull() && !data.All.IsUnknown() {
		vpnicaconnection["all"] = data.All.ValueBool()
	}
	if !data.Transproto.IsNull() && !data.Transproto.IsUnknown() {
		vpnicaconnection["transproto"] = data.Transproto.ValueString()
	}
	if !data.Username.IsNull() && !data.Username.IsUnknown() {
		vpnicaconnection["username"] = data.Username.ValueString()
	}

	return vpnicaconnection
}
