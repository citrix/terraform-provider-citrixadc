package dnsproxyrecords

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
var _ resource.Resource = &DnsproxyrecordsResource{}
var _ resource.ResourceWithConfigure = (*DnsproxyrecordsResource)(nil)
var _ resource.ResourceWithImportState = (*DnsproxyrecordsResource)(nil)

func NewDnsproxyrecordsResource() resource.Resource {
	return &DnsproxyrecordsResource{}
}

// DnsproxyrecordsResource defines the resource implementation.
type DnsproxyrecordsResource struct {
	client *service.NitroClient
}

func (r *DnsproxyrecordsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *DnsproxyrecordsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dnsproxyrecords"
}

func (r *DnsproxyrecordsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *DnsproxyrecordsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data DnsproxyrecordsResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Flushing dnsproxyrecords (action=flush)")
	dnsproxyrecords := dnsproxyrecordsGetThePayloadFromthePlan(ctx, &data)

	// Action-only resource - the only NITRO operation is the flush action (POST ?action=flush).
	err := r.client.ActOnResource(service.Dnsproxyrecords.Type(), &dnsproxyrecords, "flush")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to flush dnsproxyrecords, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Flushed dnsproxyrecords")

	// Set the synthetic ID for the resource. There is no GET endpoint to read back.
	data.Id = types.StringValue("dnsproxyrecords-config")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnsproxyrecordsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DnsproxyrecordsResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read is a no-op for dnsproxyrecords: the only NITRO operation is the flush
	// action; there is no GET endpoint, so nothing is persisted to reconcile.
	tflog.Debug(ctx, "Read is a no-op for dnsproxyrecords; no GET endpoint on NITRO side")

	// Save prior state back unchanged
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnsproxyrecordsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state DnsproxyrecordsResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Update is a no-op for dnsproxyrecords; all attributes are RequiresReplace and
	// the only NITRO operation is the flush action.
	tflog.Debug(ctx, "Update is a no-op for dnsproxyrecords; all attributes are RequiresReplace")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnsproxyrecordsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data DnsproxyrecordsResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Action-only resource - nothing is persisted on the ADC, so Delete just removes
	// the object from Terraform state.
	tflog.Debug(ctx, "Delete is a no-op for dnsproxyrecords; removed from Terraform state")
}
