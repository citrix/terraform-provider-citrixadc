package lsnsession

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &LsnsessionFlushResource{}
var _ resource.ResourceWithConfigure = (*LsnsessionFlushResource)(nil)
var _ resource.ResourceWithImportState = (*LsnsessionFlushResource)(nil)

func NewLsnsessionFlushResource() resource.Resource {
	return &LsnsessionFlushResource{}
}

// LsnsessionFlushResource defines the resource implementation.
type LsnsessionFlushResource struct {
	client *service.NitroClient
}

// LsnsessionFlushResourceModel describes the resource data model.
//
// This resource models the NITRO lsnsession `?action=flush` action. flush is a
// one-shot side-effect action with no GET endpoint and no inverse API, so
// Read/Update/Delete are no-ops. The flush payload carries the optional filter
// attributes that select which LSN sessions to flush.
type LsnsessionFlushResourceModel struct {
	Id         types.String `tfsdk:"id"`
	Clientname types.String `tfsdk:"clientname"`
	Natip      types.String `tfsdk:"natip"`
	Natport2   types.Int64  `tfsdk:"natport2"`
	Nattype    types.String `tfsdk:"nattype"`
	Netmask    types.String `tfsdk:"netmask"`
	Network    types.String `tfsdk:"network"`
	Network6   types.String `tfsdk:"network6"`
	Nodeid     types.Int64  `tfsdk:"nodeid"`
	Td         types.Int64  `tfsdk:"td"`
}

func (r *LsnsessionFlushResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LsnsessionFlushResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lsnsession_flush"
}

func (r *LsnsessionFlushResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LsnsessionFlushResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lsnsession_flush resource.",
			},
			"clientname": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the LSN Client entity.",
			},
			"natip": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Mapped NAT IP address used in LSN sessions.",
			},
			"natport2": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Mapped NAT port used in the LSN sessions.",
			},
			"nattype": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Type of sessions to be flushed. If omitted, NITRO applies its server-side default of NAT44.",
			},
			"netmask": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Subnet mask for the IP address specified by the network parameter. Must be supplied together with network.",
			},
			"network": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "IP address or network address of subscriber(s).",
			},
			"network6": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "IPv6 address of the LSN subscriber or B4 device.",
			},
			"nodeid": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Unique number that identifies the cluster node.",
			},
			"td": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Traffic domain ID of the LSN client entity.",
			},
		},
	}
}

func (r *LsnsessionFlushResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LsnsessionFlushResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Flushing lsnsession (action-only resource)")
	payload := lsnsession_flushGetThePayloadFromthePlan(ctx, &data)

	// NITRO exposes flush as POST ?action=flush. Use ActOnResource with the
	// case-sensitive "flush" verb (lower-case per the NITRO URL).
	err := r.client.ActOnResource(service.Lsnsession.Type(), payload, "flush")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to flush lsnsession, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Flushed lsnsession")

	// Synthetic ID for the action-only resource; keeps Read/Delete no-ops
	// addressable by Terraform.
	data.Id = types.StringValue("lsnsession_flush")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnsessionFlushResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// flush is a one-shot action. NITRO has no GET endpoint that reports
	// flush-state, so Read is a pure preserve-state no-op.
	var data LsnsessionFlushResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for lsnsession_flush; NITRO has no query endpoint for flush state")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnsessionFlushResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO has no update endpoint for flush; every schema attribute is
	// RequiresReplace, so Terraform never invokes Update for a real change.
	var data, state LsnsessionFlushResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for lsnsession_flush; NITRO has no update endpoint and all attributes are RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnsessionFlushResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// flush is a one-shot side-effect action. There is no inverse NITRO API
	// (no "un-flush"). Delete simply removes the resource from Terraform state.
	tflog.Debug(ctx, "Delete is a no-op for lsnsession_flush; NITRO has no inverse of the flush action")
}

func lsnsession_flushGetThePayloadFromthePlan(ctx context.Context, data *LsnsessionFlushResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In lsnsession_flushGetThePayloadFromthePlan Function")

	// Build the flush payload from whichever filter selectors are set.
	lsnsession := make(map[string]interface{})
	if !data.Clientname.IsNull() && !data.Clientname.IsUnknown() {
		lsnsession["clientname"] = data.Clientname.ValueString()
	}
	if !data.Natip.IsNull() && !data.Natip.IsUnknown() {
		lsnsession["natip"] = data.Natip.ValueString()
	}
	if !data.Natport2.IsNull() && !data.Natport2.IsUnknown() {
		lsnsession["natport2"] = int(data.Natport2.ValueInt64())
	}
	if !data.Nattype.IsNull() && !data.Nattype.IsUnknown() {
		lsnsession["nattype"] = data.Nattype.ValueString()
	}
	if !data.Netmask.IsNull() && !data.Netmask.IsUnknown() {
		lsnsession["netmask"] = data.Netmask.ValueString()
	}
	if !data.Network.IsNull() && !data.Network.IsUnknown() {
		lsnsession["network"] = data.Network.ValueString()
	}
	if !data.Network6.IsNull() && !data.Network6.IsUnknown() {
		lsnsession["network6"] = data.Network6.ValueString()
	}
	if !data.Nodeid.IsNull() && !data.Nodeid.IsUnknown() {
		lsnsession["nodeid"] = int(data.Nodeid.ValueInt64())
	}
	if !data.Td.IsNull() && !data.Td.IsUnknown() {
		lsnsession["td"] = int(data.Td.ValueInt64())
	}

	return lsnsession
}
