package dnsaddrec

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
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

	// dnsaddrec := dnsaddrecGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Dnsaddrec.Type(), &dnsaddrec)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create dnsaddrec, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("dnsaddrec-config")

	tflog.Trace(ctx, "Created dnsaddrec resource")

	// Read the updated state back
	r.readDnsaddrecFromApi(ctx, &data, &resp.Diagnostics)

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

	r.readDnsaddrecFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnsaddrecResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data DnsaddrecResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating dnsaddrec resource")

	// Create API request body from the model
	// dnsaddrec := dnsaddrecGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Dnsaddrec.Type(), &dnsaddrec)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update dnsaddrec, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated dnsaddrec resource")

	// Read the updated state back
	r.readDnsaddrecFromApi(ctx, &data, &resp.Diagnostics)

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

	// For dnsaddrec, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted dnsaddrec resource from state")
}

// Helper function to read dnsaddrec data from API
func (r *DnsaddrecResource) readDnsaddrecFromApi(ctx context.Context, data *DnsaddrecResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Dnsaddrec.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read dnsaddrec, got error: %s", err))
		return
	}

	dnsaddrecSetAttrFromGet(ctx, data, getResponseData)

}
