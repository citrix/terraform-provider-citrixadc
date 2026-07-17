package aaasession

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AaasessionKillResource{}
var _ resource.ResourceWithConfigure = (*AaasessionKillResource)(nil)
var _ resource.ResourceWithImportState = (*AaasessionKillResource)(nil)

func NewAaasessionKillResource() resource.Resource {
	return &AaasessionKillResource{}
}

// AaasessionKillResource defines the resource implementation.
//
// This resource models the NITRO aaasession `?action=kill` action. kill is a
// one-shot side-effect action with no GET endpoint and no inverse API, so
// Read/Update/Delete are no-ops. The kill payload carries the optional filter
// attributes all, groupname, iip, netmask, sessionkey and username. nodeid is a
// GET-only cluster filter (Pattern 15) and is therefore excluded from the kill
// payload (see the payload builder).
type AaasessionKillResource struct {
	client *service.NitroClient
}

// AaasessionKillResourceModel describes the resource data model.
type AaasessionKillResourceModel struct {
	Id         types.String `tfsdk:"id"`
	All        types.Bool   `tfsdk:"all"`
	Groupname  types.String `tfsdk:"groupname"`
	Iip        types.String `tfsdk:"iip"`
	Netmask    types.String `tfsdk:"netmask"`
	Nodeid     types.Int64  `tfsdk:"nodeid"`
	Sessionkey types.String `tfsdk:"sessionkey"`
	Username   types.String `tfsdk:"username"`
}

func (r *AaasessionKillResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AaasessionKillResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_aaasession_kill"
}

func (r *AaasessionKillResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AaasessionKillResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the aaasession_kill resource.",
			},
			// All kill arguments are optional filters. Read is a no-op (kill is an
			// action, the killed session is not a persistent managed object), so
			// these must NOT be Computed or Terraform reports an unknown value
			// after apply (Pattern 13 schema-flag implication).
			"all": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Terminate all active AAA-TM/VPN sessions.",
			},
			"groupname": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the AAA group.",
			},
			"iip": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "IP address or the first address in the intranet IP range.",
			},
			"netmask": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Subnet mask for the intranet IP range.",
			},
			"nodeid": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Unique number that identifies the cluster node.",
			},
			"sessionkey": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Show aaa session associated with given session key",
			},
			"username": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the AAA user.",
			},
		},
	}
}

func (r *AaasessionKillResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AaasessionKillResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Killing aaasession (action-only resource)")
	payload := aaasession_killGetThePayloadFromthePlan(ctx, &data)

	// kill is a POST ?action=kill action (Pattern 1). There is no add endpoint;
	// AddResource/UpdateUnnamedResource would target a nonexistent verb. The
	// verb casing is lower-case per the NITRO URL.
	err := r.client.ActOnResource(service.Aaasession.Type(), payload, "kill")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to kill aaasession, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Killed aaasession")

	// Synthetic ID for the action-only resource; keeps Read/Delete no-ops
	// addressable by Terraform. The killed session is not a queryable managed
	// object, so this ID is purely a Terraform state handle.
	data.Id = types.StringValue("aaasession_kill")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaasessionKillResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// kill is a one-shot action. NITRO has no GET endpoint that reports
	// kill-state, so Read is a pure preserve-state no-op.
	var data AaasessionKillResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for aaasession_kill; kill has no stable GET-backed object")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaasessionKillResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO has no update endpoint for kill; every schema attribute is
	// RequiresReplace, so Terraform never invokes Update for a real change.
	var data, state AaasessionKillResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for aaasession_kill; NITRO has no update endpoint and all attributes are RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaasessionKillResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// kill is a one-shot side-effect action. There is no inverse NITRO API
	// (no "un-kill"). Delete simply removes the resource from Terraform state.
	tflog.Debug(ctx, "Delete is a no-op for aaasession_kill; NITRO has no inverse of the kill action")
}

// aaasession_killGetThePayloadFromthePlan builds the body for the kill action.
// Build a map containing ONLY the kill arguments that are set so the action body
// never includes invalid arguments. nodeid is a GET-only cluster filter and is
// intentionally excluded from the kill payload (Pattern 15).
func aaasession_killGetThePayloadFromthePlan(ctx context.Context, data *AaasessionKillResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In aaasession_killGetThePayloadFromthePlan Function")

	aaasession := map[string]interface{}{}
	if !data.All.IsNull() && !data.All.IsUnknown() {
		aaasession["all"] = data.All.ValueBool()
	}
	if !data.Groupname.IsNull() && !data.Groupname.IsUnknown() {
		aaasession["groupname"] = data.Groupname.ValueString()
	}
	if !data.Iip.IsNull() && !data.Iip.IsUnknown() {
		aaasession["iip"] = data.Iip.ValueString()
	}
	if !data.Netmask.IsNull() && !data.Netmask.IsUnknown() {
		aaasession["netmask"] = data.Netmask.ValueString()
	}
	if !data.Sessionkey.IsNull() && !data.Sessionkey.IsUnknown() {
		aaasession["sessionkey"] = data.Sessionkey.ValueString()
	}
	if !data.Username.IsNull() && !data.Username.IsUnknown() {
		aaasession["username"] = data.Username.ValueString()
	}

	return aaasession
}
