package ipsecalgsession

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/ipsecalg"
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
var _ resource.Resource = &IpsecalgsessionFlushResource{}
var _ resource.ResourceWithConfigure = (*IpsecalgsessionFlushResource)(nil)
var _ resource.ResourceWithImportState = (*IpsecalgsessionFlushResource)(nil)

func NewIpsecalgsessionFlushResource() resource.Resource {
	return &IpsecalgsessionFlushResource{}
}

// IpsecalgsessionFlushResource defines the resource implementation.
//
// This resource models the NITRO ipsecalgsession `?action=flush` action. flush is
// a one-shot side-effect action with no add/set/delete endpoint and no inverse
// API, so Read/Update/Delete are no-ops. The flush payload carries the optional
// scope attributes sourceip/natip/destip (a bare flush flushes all sessions).
type IpsecalgsessionFlushResource struct {
	client *service.NitroClient
}

// IpsecalgsessionFlushResourceModel describes the resource data model.
type IpsecalgsessionFlushResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Destip      types.String `tfsdk:"destip"`
	DestipAlg   types.String `tfsdk:"destip_alg"`
	Natip       types.String `tfsdk:"natip"`
	NatipAlg    types.String `tfsdk:"natip_alg"`
	Sourceip    types.String `tfsdk:"sourceip"`
	SourceipAlg types.String `tfsdk:"sourceip_alg"`
}

func (r *IpsecalgsessionFlushResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *IpsecalgsessionFlushResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ipsecalgsession_flush"
}

func (r *IpsecalgsessionFlushResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *IpsecalgsessionFlushResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the ipsecalgsession_flush resource.",
			},
			"destip": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Destination IP address (flush scope).",
			},
			"destip_alg": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Destination IP address.",
			},
			"natip": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Natted Source IP address (flush scope).",
			},
			"natip_alg": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Natted Source IP address.",
			},
			"sourceip": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Original Source IP address (flush scope).",
			},
			"sourceip_alg": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Original Source IP address.",
			},
		},
	}
}

func (r *IpsecalgsessionFlushResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data IpsecalgsessionFlushResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Flushing ipsecalgsession (action-only resource)")
	payload := ipsecalgsession_flushGetThePayloadFromthePlan(ctx, &data)

	// NITRO exposes flush as POST ?action=flush. Use ActOnResource with the
	// case-sensitive "flush" verb (lower-case per the NITRO URL).
	err := r.client.ActOnResource(service.Ipsecalgsession.Type(), &payload, "flush")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to flush ipsecalgsession, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Flushed ipsecalgsession")

	// Synthetic ID for the action-only resource; keeps Read/Delete no-ops
	// addressable by Terraform.
	data.Id = types.StringValue("ipsecalgsession_flush")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IpsecalgsessionFlushResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// flush is a one-shot action. NITRO has no GET endpoint that reports
	// flush-state, so Read is a pure preserve-state no-op.
	var data IpsecalgsessionFlushResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for ipsecalgsession_flush; NITRO has no query endpoint for flush state")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IpsecalgsessionFlushResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO has no update endpoint for flush; every schema attribute is
	// RequiresReplace, so Terraform never invokes Update for a real change.
	var data, state IpsecalgsessionFlushResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for ipsecalgsession_flush; NITRO has no update endpoint and all attributes are RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IpsecalgsessionFlushResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// flush is a one-shot side-effect action. There is no inverse NITRO API
	// (no "un-flush"). Delete simply removes the resource from Terraform state.
	tflog.Debug(ctx, "Delete is a no-op for ipsecalgsession_flush; NITRO has no inverse of the flush action")
}

// ipsecalgsession_flushGetThePayloadFromthePlan builds the ?action=flush payload.
// Per the NITRO doc, flush accepts only the canonical scope fields
// sourceip/natip/destip (a bare flush with no fields flushes all sessions). The
// *_alg twins are GET filter names and are NOT sent, to avoid emitting
// conflicting duplicate keys (NitroValidator Pattern 15).
func ipsecalgsession_flushGetThePayloadFromthePlan(ctx context.Context, data *IpsecalgsessionFlushResourceModel) ipsecalg.Ipsecalgsession {
	tflog.Debug(ctx, "In ipsecalgsession_flushGetThePayloadFromthePlan Function")

	ipsecalgsession := ipsecalg.Ipsecalgsession{}
	if !data.Sourceip.IsNull() && !data.Sourceip.IsUnknown() {
		ipsecalgsession.Sourceip = data.Sourceip.ValueString()
	}
	if !data.Natip.IsNull() && !data.Natip.IsUnknown() {
		ipsecalgsession.Natip = data.Natip.ValueString()
	}
	if !data.Destip.IsNull() && !data.Destip.IsUnknown() {
		ipsecalgsession.Destip = data.Destip.ValueString()
	}

	return ipsecalgsession
}
