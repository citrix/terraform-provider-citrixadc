package dnsproxyrecords

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/dns"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &DnsproxyrecordsFlushResource{}
var _ resource.ResourceWithConfigure = (*DnsproxyrecordsFlushResource)(nil)

func NewDnsproxyrecordsFlushResource() resource.Resource {
	return &DnsproxyrecordsFlushResource{}
}

// DnsproxyrecordsFlushResource defines the resource implementation.
type DnsproxyrecordsFlushResource struct {
	client *service.NitroClient
}

// DnsproxyrecordsFlushResourceModel describes the resource data model.
//
// This resource models the NITRO dnsproxyrecords `?action=flush` action. flush is a
// one-shot side-effect action with no GET endpoint and no inverse API, so
// Read/Update/Delete are no-ops. The flush payload carries the optional filter
// attributes negrectype and type.
type DnsproxyrecordsFlushResourceModel struct {
	Id         types.String `tfsdk:"id"`
	Negrectype types.String `tfsdk:"negrectype"`
	Type       types.String `tfsdk:"type"`
}

func (r *DnsproxyrecordsFlushResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dnsproxyrecords_flush"
}

func (r *DnsproxyrecordsFlushResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *DnsproxyrecordsFlushResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the dnsproxyrecords_flush resource.",
			},
			"negrectype": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Filter the Negative DNS records i.e NXDOMAIN and NODATA entries to be flushed. e.g flush dns proxyRecords NXDOMAIN will flush only the NXDOMAIN entries from the cache",
			},
			"type": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Filter the DNS records to be flushed.e.g flush dns proxyRecords -type A   will flush only the A records from the cache.",
			},
		},
	}
}

func (r *DnsproxyrecordsFlushResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data DnsproxyrecordsFlushResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Flushing dnsproxyrecords (action-only resource)")
	payload := dnsproxyrecords_flushGetThePayloadFromthePlan(ctx, &data)

	// NITRO exposes flush as POST ?action=flush. Use ActOnResource with the
	// case-sensitive "flush" verb (lower-case per the NITRO URL).
	err := r.client.ActOnResource(service.Dnsproxyrecords.Type(), &payload, "flush")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to flush dnsproxyrecords, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Flushed dnsproxyrecords")

	// Synthetic ID for the action-only resource; keeps Read/Delete no-ops
	// addressable by Terraform.
	data.Id = types.StringValue("dnsproxyrecords_flush")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnsproxyrecordsFlushResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// flush is a one-shot action. NITRO has no GET endpoint that reports
	// flush-state, so Read is a pure preserve-state no-op.
	var data DnsproxyrecordsFlushResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for dnsproxyrecords_flush; NITRO has no query endpoint for flush state")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnsproxyrecordsFlushResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO has no update endpoint for flush; every schema attribute is
	// RequiresReplace, so Terraform never invokes Update for a real change.
	var data, state DnsproxyrecordsFlushResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for dnsproxyrecords_flush; NITRO has no update endpoint and all attributes are RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnsproxyrecordsFlushResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// flush is a one-shot side-effect action. There is no inverse NITRO API
	// (no "un-flush"). Delete simply removes the resource from Terraform state.
	tflog.Debug(ctx, "Delete is a no-op for dnsproxyrecords_flush; NITRO has no inverse of the flush action")
}

func dnsproxyrecords_flushGetThePayloadFromthePlan(ctx context.Context, data *DnsproxyrecordsFlushResourceModel) dns.Dnsproxyrecords {
	tflog.Debug(ctx, "In dnsproxyrecords_flushGetThePayloadFromthePlan Function")

	// Create API request body from the model
	dnsproxyrecords := dns.Dnsproxyrecords{}
	if !data.Negrectype.IsNull() && !data.Negrectype.IsUnknown() {
		dnsproxyrecords.Negrectype = data.Negrectype.ValueString()
	}
	if !data.Type.IsNull() && !data.Type.IsUnknown() {
		dnsproxyrecords.Type = data.Type.ValueString()
	}

	return dnsproxyrecords
}
