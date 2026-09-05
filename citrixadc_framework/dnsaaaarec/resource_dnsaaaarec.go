package dnsaaaarec

import (
	"context"
	"fmt"
	"net/url"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &DnsaaaarecResource{}
var _ resource.ResourceWithConfigure = (*DnsaaaarecResource)(nil)
var _ resource.ResourceWithImportState = (*DnsaaaarecResource)(nil)

func NewDnsaaaarecResource() resource.Resource {
	return &DnsaaaarecResource{}
}

// DnsaaaarecResource defines the resource implementation.
type DnsaaaarecResource struct {
	client *service.NitroClient
}

func (r *DnsaaaarecResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *DnsaaaarecResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dnsaaaarec"
}

func (r *DnsaaaarecResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *DnsaaaarecResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data DnsaaaarecResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating dnsaaaarec resource")

	dnsaaaarec := dnsaaaarecGetThePayloadFromtheConfig(ctx, &data)

	// Named resource keyed on the primary attribute (hostname) - use AddResource
	hostname := data.Hostname.ValueString()
	_, err := r.client.AddResource(service.Dnsaaaarec.Type(), hostname, &dnsaaaarec)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create dnsaaaarec, got error: %s", err))
		return
	}

	// ID matches the SDK v2 resource format (d.SetId(hostname))
	data.Id = types.StringValue(hostname)

	tflog.Trace(ctx, "Created dnsaaaarec resource")

	// Read the updated state back
	if !r.readDnsaaaarecFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "dnsaaaarec not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnsaaaarecResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DnsaaaarecResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading dnsaaaarec resource")

	found := r.readDnsaaaarecFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *DnsaaaarecResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state DnsaaaarecResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating dnsaaaarec resource")

	// dnsaaaarec has no NITRO-updatable attributes; every attribute is ForceNew
	// (RequiresReplace), so any change triggers a destroy/create rather than an
	// in-place update. There is nothing to push to the ADC here.
	tflog.Trace(ctx, "No updatable attributes for dnsaaaarec resource, refreshing state")

	// Read the updated state back
	if !r.readDnsaaaarecFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "dnsaaaarec not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnsaaaarecResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data DnsaaaarecResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting dnsaaaarec resource")

	// Named resource - delete keyed on hostname, disambiguated by ipv6address
	// (and optionally ecssubnet) exactly as the SDK v2 resource did.
	argsMap := make(map[string]string)
	if !data.Ecssubnet.IsNull() && !data.Ecssubnet.IsUnknown() && data.Ecssubnet.ValueString() != "" {
		argsMap["ecssubnet"] = url.QueryEscape(data.Ecssubnet.ValueString())
	}
	argsMap["ipv6address"] = url.QueryEscape(data.Ipv6address.ValueString())

	err := r.client.DeleteResourceWithArgsMap(service.Dnsaaaarec.Type(), data.Id.ValueString(), argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete dnsaaaarec, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted dnsaaaarec resource")
}

// Helper function to read dnsaaaarec data from API.
// Returns false when the record no longer exists on the ADC.
func (r *DnsaaaarecResource) readDnsaaaarecFromApi(ctx context.Context, data *DnsaaaarecResourceModel, diags *diag.Diagnostics) bool {
	// hostname is the primary key / ID. On import only the ID is populated.
	hostname := data.Hostname.ValueString()
	if hostname == "" {
		hostname = data.Id.ValueString()
	}
	ipv6address := data.Ipv6address.ValueString()

	findParams := service.FindParams{
		ResourceType: service.Dnsaaaarec.Type(),
	}
	dataArr, err := r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read dnsaaaarec, got error: %s", err))
		return false
	}

	if len(dataArr) == 0 {
		return false
	}

	// Match on hostname (and ipv6address when known) to select the right record
	// among possibly many records sharing the same hostname.
	foundIndex := -1
	for i, v := range dataArr {
		hn, _ := v["hostname"].(string)
		if hn != hostname {
			continue
		}
		if ipv6address == "" {
			foundIndex = i
			break
		}
		ip, _ := v["ipv6address"].(string)
		if ip == ipv6address {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		return false
	}

	dnsaaaarecSetAttrFromGet(ctx, data, dataArr[foundIndex])

	return true
}
