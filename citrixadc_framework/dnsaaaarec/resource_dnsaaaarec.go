package dnsaaaarec

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

	// dnsaaaarec := dnsaaaarecGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Dnsaaaarec.Type(), &dnsaaaarec)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create dnsaaaarec, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("dnsaaaarec-config")

	tflog.Trace(ctx, "Created dnsaaaarec resource")

	// Read the updated state back
	r.readDnsaaaarecFromApi(ctx, &data, &resp.Diagnostics)

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

	r.readDnsaaaarecFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnsaaaarecResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data DnsaaaarecResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating dnsaaaarec resource")

	// Create API request body from the model
	// dnsaaaarec := dnsaaaarecGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Dnsaaaarec.Type(), &dnsaaaarec)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update dnsaaaarec, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated dnsaaaarec resource")

	// Read the updated state back
	r.readDnsaaaarecFromApi(ctx, &data, &resp.Diagnostics)

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

	// For dnsaaaarec, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted dnsaaaarec resource from state")
}

// Helper function to read dnsaaaarec data from API
func (r *DnsaaaarecResource) readDnsaaaarecFromApi(ctx context.Context, data *DnsaaaarecResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Dnsaaaarec.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read dnsaaaarec, got error: %s", err))
		return
	}

	dnsaaaarecSetAttrFromGet(ctx, data, getResponseData)

}
