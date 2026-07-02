package dnssubnetcache

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &DnssubnetcacheResource{}
var _ resource.ResourceWithConfigure = (*DnssubnetcacheResource)(nil)
var _ resource.ResourceWithImportState = (*DnssubnetcacheResource)(nil)

func NewDnssubnetcacheResource() resource.Resource {
	return &DnssubnetcacheResource{}
}

// DnssubnetcacheResource defines the resource implementation.
type DnssubnetcacheResource struct {
	client *service.NitroClient
}

func (r *DnssubnetcacheResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *DnssubnetcacheResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dnssubnetcache"
}

func (r *DnssubnetcacheResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

// Create performs the dnssubnetcache flush action. dnssubnetcache has NO add/set/delete
// endpoints on NITRO; the only mutating operation is the "flush" action (POST ?action=flush),
// which flushes ECS subnet(s) from the runtime DNS cache. There is no persistent object to
// reconcile, so Read/Update/Delete are no-ops (Pattern 13).
func (r *DnssubnetcacheResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data DnssubnetcacheResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Flushing dnssubnetcache (action=flush)")
	payload := dnssubnetcacheGetThePayloadFromthePlan(ctx, &data)

	// Action-only resource - flush via ActOnResource (POST ?action=flush)
	err := r.client.ActOnResource(service.Dnssubnetcache.Type(), payload, "flush")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to flush dnssubnetcache, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Flushed dnssubnetcache")

	// Synthetic ID - there is no persistent object.
	data.Id = types.StringValue("dnssubnetcache-flush")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Read is a no-op. dnssubnetcache is a runtime DNS cache with no persistent
// configuration object to reconcile against (Pattern 13).
func (r *DnssubnetcacheResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DnssubnetcacheResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for dnssubnetcache; it is a runtime cache with no persistent object")

	// Preserve prior state unchanged.
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Update is a no-op. All trigger attributes are RequiresReplace and there is no
// NITRO set/update endpoint for dnssubnetcache (Pattern 13).
func (r *DnssubnetcacheResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state DnssubnetcacheResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Update is a no-op for dnssubnetcache; all attributes are RequiresReplace and NITRO has no set endpoint")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Delete is a no-op. dnssubnetcache has no NITRO delete endpoint; the flush action
// is not reversible and there is no persistent object to remove (Pattern 13).
func (r *DnssubnetcacheResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete is a no-op for dnssubnetcache; NITRO has no delete endpoint")
}
