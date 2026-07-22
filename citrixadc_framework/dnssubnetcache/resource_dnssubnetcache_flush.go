package dnssubnetcache

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
var _ resource.Resource = &DnssubnetcacheFlushResource{}
var _ resource.ResourceWithConfigure = (*DnssubnetcacheFlushResource)(nil)
var _ resource.ResourceWithValidateConfig = (*DnssubnetcacheFlushResource)(nil)

func NewDnssubnetcacheFlushResource() resource.Resource {
	return &DnssubnetcacheFlushResource{}
}

// DnssubnetcacheFlushResource defines the resource implementation.
type DnssubnetcacheFlushResource struct {
	client *service.NitroClient
}

// DnssubnetcacheFlushResourceModel describes the resource data model.
//
// This resource models the NITRO dnssubnetcache `?action=flush` action (POST).
// flush is a one-shot side-effect action that flushes ECS subnet(s) from the
// runtime DNS cache; there is no GET endpoint reporting flush-state and no
// inverse API, so Read/Update/Delete are no-ops. The flush payload carries only
// the mutually-exclusive filter attributes ecssubnet/all.
//
// nodeid is a NITRO GET-only cluster filter (Pattern 15) and is intentionally
// not modelled: it is never part of the flush write payload.
type DnssubnetcacheFlushResourceModel struct {
	Id        types.String `tfsdk:"id"`
	All       types.Bool   `tfsdk:"all"`
	Ecssubnet types.String `tfsdk:"ecssubnet"`
}

func (r *DnssubnetcacheFlushResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dnssubnetcache_flush"
}

func (r *DnssubnetcacheFlushResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

// ValidateConfig enforces the NITRO flush mandatory mutually-exclusive choice.
// The CLI synopsis is `flush dns subnetcache (<ecsSubnet> | -all)`: the group is
// outside the optional brackets, so exactly one of ecssubnet/all is required.
// A config supplying neither builds an empty flush payload and the appliance
// returns "Required argument missing [all, ecsSubnet]"; supplying both violates
// the mutual exclusivity (Pattern 8/17).
func (r *DnssubnetcacheFlushResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data DnssubnetcacheFlushResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	hasAll := !data.All.IsNull() && !data.All.IsUnknown()
	hasEcssubnet := !data.Ecssubnet.IsNull() && !data.Ecssubnet.IsUnknown()

	if hasAll && hasEcssubnet {
		resp.Diagnostics.AddError(
			"Invalid dnssubnetcache_flush configuration",
			"Specify exactly one of \"ecssubnet\" OR \"all\", not both.",
		)
		return
	}
	if !hasAll && !hasEcssubnet {
		resp.Diagnostics.AddError(
			"Invalid dnssubnetcache_flush configuration",
			"You must specify exactly one of \"ecssubnet\" OR \"all\".",
		)
		return
	}
}

func (r *DnssubnetcacheFlushResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the dnssubnetcache_flush resource.",
			},
			"all": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Flush all the ECS subnets from the DNS cache.",
			},
			"ecssubnet": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "ECS Subnet.",
			},
		},
	}
}

func (r *DnssubnetcacheFlushResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data DnssubnetcacheFlushResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Flushing dnssubnetcache (action-only resource)")
	payload := dnssubnetcache_flushGetThePayloadFromthePlan(ctx, &data)

	// NITRO exposes flush as POST ?action=flush. Use ActOnResource with the
	// case-sensitive "flush" verb (lower-case per the NITRO URL).
	err := r.client.ActOnResource(service.Dnssubnetcache.Type(), payload, "flush")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to flush dnssubnetcache, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Flushed dnssubnetcache")

	// Synthetic ID for the action-only resource; keeps Read/Delete no-ops
	// addressable by Terraform.
	data.Id = types.StringValue("dnssubnetcache_flush")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnssubnetcacheFlushResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// flush is a one-shot action. NITRO has no GET endpoint that reports
	// flush-state, so Read is a pure preserve-state no-op.
	var data DnssubnetcacheFlushResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for dnssubnetcache_flush; NITRO has no query endpoint for flush state")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnssubnetcacheFlushResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO has no update endpoint for flush; every schema attribute is
	// RequiresReplace, so Terraform never invokes Update for a real change.
	var data, state DnssubnetcacheFlushResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for dnssubnetcache_flush; NITRO has no update endpoint and all attributes are RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnssubnetcacheFlushResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// flush is a one-shot side-effect action. There is no inverse NITRO API
	// (no "un-flush"). Delete simply removes the resource from Terraform state.
	tflog.Debug(ctx, "Delete is a no-op for dnssubnetcache_flush; NITRO has no inverse of the flush action")
}

func dnssubnetcache_flushGetThePayloadFromthePlan(ctx context.Context, data *DnssubnetcacheFlushResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In dnssubnetcache_flushGetThePayloadFromthePlan Function")

	// Build the flush action payload. Only ecssubnet/all are valid flush inputs;
	// nodeid is a GET-only filter and must never be sent.
	dnssubnetcache := make(map[string]interface{})
	if !data.All.IsNull() && !data.All.IsUnknown() {
		dnssubnetcache["all"] = data.All.ValueBool()
	}
	if !data.Ecssubnet.IsNull() && !data.Ecssubnet.IsUnknown() {
		dnssubnetcache["ecssubnet"] = data.Ecssubnet.ValueString()
	}

	return dnssubnetcache
}
