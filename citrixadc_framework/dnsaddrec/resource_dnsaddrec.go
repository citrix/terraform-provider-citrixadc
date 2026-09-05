package dnsaddrec

import (
	"context"
	"fmt"
	"net/url"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &DnsaddrecResource{}
var _ resource.ResourceWithConfigure = (*DnsaddrecResource)(nil)
var _ resource.ResourceWithImportState = (*DnsaddrecResource)(nil)

func NewDnsaddrecResource() resource.Resource {
	return &DnsaddrecResource{}
}

// DnsaddrecResource defines the resource implementation.
type DnsaddrecResource struct {
	client *service.NitroClient
}

func (r *DnsaddrecResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *DnsaddrecResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dnsaddrec"
}

func (r *DnsaddrecResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *DnsaddrecResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data DnsaddrecResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating dnsaddrec resource")

	dnsaddrec := dnsaddrecGetThePayloadFromtheConfig(ctx, &data)

	// Named resource - use AddResource.
	_, err := r.client.AddResource(service.Dnsaddrec.Type(), data.Hostname.ValueString(), &dnsaddrec)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create dnsaddrec, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created dnsaddrec resource")

	// Set ID for the resource before reading state back.
	// Backward-compatible SDK v2 ID format: "hostname,ipaddress".
	data.Id = types.StringValue(fmt.Sprintf("%s,%s", data.Hostname.ValueString(), data.Ipaddress.ValueString()))

	// Read the updated state back
	if !r.readDnsaddrecFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "dnsaddrec not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnsaddrecResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DnsaddrecResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading dnsaddrec resource")

	found := r.readDnsaddrecFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if !found {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnsaddrecResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state DnsaddrecResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state.
	data.Id = state.Id

	tflog.Debug(ctx, "Updating dnsaddrec resource")

	// dnsaddrec has no NITRO-updatable attributes: every attribute is ForceNew
	// (RequiresReplace), so any attribute change is handled by destroy+create and
	// Update is only reached when nothing meaningful changed. No write call is made.
	tflog.Trace(ctx, "No updatable attributes for dnsaddrec resource, skipping update")

	// Read the current state back
	if !r.readDnsaddrecFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "dnsaddrec not found during update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnsaddrecResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data DnsaddrecResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting dnsaddrec resource")

	// Named resource with delete args. The DELETE URL is keyed on hostname and the
	// specific record is selected by ipaddress (and optionally ecssubnet), matching
	// the SDK v2 behavior and the NITRO delete signature
	// (args=ecssubnet:<v>,ipaddress:<v>).
	argsMap := make(map[string]string)
	if !data.Ecssubnet.IsNull() && data.Ecssubnet.ValueString() != "" {
		argsMap["ecssubnet"] = url.QueryEscape(data.Ecssubnet.ValueString())
	}
	argsMap["ipaddress"] = url.QueryEscape(data.Ipaddress.ValueString())

	err := r.client.DeleteResourceWithArgsMap(service.Dnsaddrec.Type(), data.Hostname.ValueString(), argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete dnsaddrec, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted dnsaddrec resource")
}

// Helper function to read dnsaddrec data from API. Returns true when the record
// was found, false when it no longer exists (so the caller can drop it from state).
func (r *DnsaddrecResource) readDnsaddrecFromApi(ctx context.Context, data *DnsaddrecResourceModel, diags *diag.Diagnostics) bool {
	// The ID uniquely identifies the record by (hostname, ipaddress). Parse it so
	// both new and legacy SDK v2 IDs resolve correctly (also covers import).
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"hostname", "ipaddress"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse dnsaddrec ID %s, got error: %s", data.Id.ValueString(), err))
		return false
	}
	hostname := idMap["hostname"]
	ipaddress := idMap["ipaddress"]

	// A single hostname can hold multiple address records (one per ipaddress), so
	// fetch all records and filter by both hostname and ipaddress.
	dataArr, err := r.client.FindAllResources(service.Dnsaddrec.Type())
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read dnsaddrec, got error: %s", err))
		return false
	}
	if len(dataArr) == 0 {
		return false
	}

	foundIndex := -1
	for i, v := range dataArr {
		hostnameVal, hok := v["hostname"].(string)
		ipaddressVal, iok := v["ipaddress"].(string)
		if hok && iok && hostnameVal == hostname && ipaddressVal == ipaddress {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		return false
	}

	dnsaddrecSetAttrFromGet(ctx, data, dataArr[foundIndex])

	return true
}
