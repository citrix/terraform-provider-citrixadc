package dnsnsrec

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &DnsnsrecResource{}
var _ resource.ResourceWithConfigure = (*DnsnsrecResource)(nil)
var _ resource.ResourceWithImportState = (*DnsnsrecResource)(nil)

func NewDnsnsrecResource() resource.Resource {
	return &DnsnsrecResource{}
}

// DnsnsrecResource defines the resource implementation.
type DnsnsrecResource struct {
	client *service.NitroClient
}

func (r *DnsnsrecResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *DnsnsrecResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dnsnsrec"
}

func (r *DnsnsrecResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *DnsnsrecResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data DnsnsrecResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating dnsnsrec resource")

	dnsnsrec := dnsnsrecGetThePayloadFromthePlan(ctx, &data)

	// Named resource - the add URL takes no resource name (matches SDK v2).
	_, err := r.client.AddResource(service.Dnsnsrec.Type(), "", &dnsnsrec)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create dnsnsrec, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created dnsnsrec resource")

	// Set ID (legacy SDK v2 composite "domain,nameserver") before reading state back.
	data.Id = types.StringValue(fmt.Sprintf("%s,%s", data.Domain.ValueString(), data.Nameserver.ValueString()))

	// Read the created state back
	r.readDnsnsrecFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "dnsnsrec not found on the ADC immediately after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnsnsrecResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DnsnsrecResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading dnsnsrec resource")

	r.readDnsnsrecFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	// Record is gone on the ADC (readFromApi nulled the Id): drop it from state so a
	// subsequent apply recreates it, matching the SDK v2 provider's behaviour.
	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnsnsrecResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state DnsnsrecResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state.
	data.Id = state.Id

	// dnsnsrec has no updateable attributes: every attribute is ForceNew /
	// RequiresReplace (matching the SDK v2 resource, which defined no update path).
	// So Update never carries an in-place change; just re-read current state.
	tflog.Debug(ctx, "Updating dnsnsrec resource (no updateable attributes)")

	r.readDnsnsrecFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "dnsnsrec not found on the ADC during update")
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnsnsrecResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data DnsnsrecResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting dnsnsrec resource")

	// Delete keys on domain (URL) with nameserver as a mandatory query arg
	// (matches SDK v2 DeleteResourceWithArgsMap(type, domain, {nameserver})).
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"domain", "nameserver"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	domain_value, ok := idMap["domain"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Attribute 'domain' not found in ID")
		return
	}

	argsMap := make(map[string]string)
	if val, ok := idMap["nameserver"]; ok && val != "" {
		argsMap["nameserver"] = val
	}

	err = r.client.DeleteResourceWithArgsMap(service.Dnsnsrec.Type(), domain_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete dnsnsrec, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted dnsnsrec resource")
}

// Helper function to read dnsnsrec data from API. A domain can have multiple
// name server records, so the record is located by filtering the array on both
// domain and nameserver (matches the SDK v2 read). If it is not present the Id is
// set to null so callers can drop it from state.
func (r *DnsnsrecResource) readDnsnsrecFromApi(ctx context.Context, data *DnsnsrecResourceModel, diags *diag.Diagnostics) {
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"domain", "nameserver"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	domain_Name, ok := idMap["domain"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'domain' not found in ID string")
		return
	}
	nameserver_Value := idMap["nameserver"]

	findParams := service.FindParams{
		ResourceType:             service.Dnsnsrec.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err := r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read dnsnsrec, got error: %s", err))
		return
	}

	if len(dataArr) == 0 {
		data.Id = types.StringNull()
		return
	}

	foundIndex := -1
	for i, v := range dataArr {
		domainVal, ok := v["domain"].(string)
		if !ok || domainVal != domain_Name {
			continue
		}
		nsVal, ok := v["nameserver"].(string)
		if !ok || nsVal != nameserver_Value {
			continue
		}
		foundIndex = i
		break
	}

	if foundIndex == -1 {
		data.Id = types.StringNull()
		return
	}

	dnsnsrecSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
