package rnatsession

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/network"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &RnatsessionFlushResource{}
var _ resource.ResourceWithConfigure = (*RnatsessionFlushResource)(nil)
var _ resource.ResourceWithImportState = (*RnatsessionFlushResource)(nil)

func NewRnatsessionFlushResource() resource.Resource {
	return &RnatsessionFlushResource{}
}

// RnatsessionFlushResource defines the resource implementation.
type RnatsessionFlushResource struct {
	client *service.NitroClient
}

// RnatsessionFlushResourceModel describes the resource data model.
//
// This resource models the NITRO rnatsession `?action=flush` action. flush is a
// one-shot side-effect action with no GET endpoint and no inverse API, so
// Read/Update/Delete are no-ops. The flush payload carries the optional filter
// attributes network, netmask, natip and aclname.
type RnatsessionFlushResourceModel struct {
	Id      types.String `tfsdk:"id"`
	Aclname types.String `tfsdk:"aclname"`
	Natip   types.String `tfsdk:"natip"`
	Netmask types.String `tfsdk:"netmask"`
	Network types.String `tfsdk:"network"`
}

func (r *RnatsessionFlushResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *RnatsessionFlushResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rnatsession_flush"
}

func (r *RnatsessionFlushResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *RnatsessionFlushResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the rnatsession_flush resource.",
			},
			"aclname": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of any configured extended ACL whose action is ALLOW.",
			},
			"natip": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The NAT IP address defined for the RNAT entry.",
			},
			"netmask": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Subnet mask associated with the network address.",
			},
			"network": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "IPv4 network address on whose traffic you want the Citrix ADC to do RNAT processing.",
			},
		},
	}
}

func (r *RnatsessionFlushResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data RnatsessionFlushResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Flushing rnatsession (action-only resource)")
	payload := rnatsession_flushGetThePayloadFromthePlan(ctx, &data)

	// NITRO exposes flush as POST ?action=flush. Use ActOnResource with the
	// case-sensitive "flush" verb (lower-case per the NITRO URL).
	err := r.client.ActOnResource(service.Rnatsession.Type(), &payload, "flush")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to flush rnatsession, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Flushed rnatsession")

	// Synthetic ID for the action-only resource; keeps Read/Delete no-ops
	// addressable by Terraform.
	data.Id = types.StringValue("rnatsession_flush")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RnatsessionFlushResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// flush is a one-shot action. NITRO has no GET endpoint that reports
	// flush-state, so Read is a pure preserve-state no-op.
	var data RnatsessionFlushResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for rnatsession_flush; NITRO has no query endpoint for flush state")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RnatsessionFlushResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO has no update endpoint for flush; every schema attribute is
	// RequiresReplace, so Terraform never invokes Update for a real change.
	var data, state RnatsessionFlushResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for rnatsession_flush; NITRO has no update endpoint and all attributes are RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RnatsessionFlushResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// flush is a one-shot side-effect action. There is no inverse NITRO API
	// (no "un-flush"). Delete simply removes the resource from Terraform state.
	tflog.Debug(ctx, "Delete is a no-op for rnatsession_flush; NITRO has no inverse of the flush action")
}

func rnatsession_flushGetThePayloadFromthePlan(ctx context.Context, data *RnatsessionFlushResourceModel) network.Rnatsession {
	tflog.Debug(ctx, "In rnatsession_flushGetThePayloadFromthePlan Function")

	// Create API request body from the model
	rnatsession := network.Rnatsession{}
	if !data.Aclname.IsNull() && !data.Aclname.IsUnknown() {
		rnatsession.Aclname = data.Aclname.ValueString()
	}
	if !data.Natip.IsNull() && !data.Natip.IsUnknown() {
		rnatsession.Natip = data.Natip.ValueString()
	}
	if !data.Netmask.IsNull() && !data.Netmask.IsUnknown() {
		rnatsession.Netmask = data.Netmask.ValueString()
	}
	if !data.Network.IsNull() && !data.Network.IsUnknown() {
		rnatsession.Network = data.Network.ValueString()
	}

	return rnatsession
}
