package dnsview

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
var _ resource.Resource = &DnsviewResource{}
var _ resource.ResourceWithConfigure = (*DnsviewResource)(nil)
var _ resource.ResourceWithImportState = (*DnsviewResource)(nil)

func NewDnsviewResource() resource.Resource {
	return &DnsviewResource{}
}

// DnsviewResource defines the resource implementation.
type DnsviewResource struct {
	client *service.NitroClient
}

func (r *DnsviewResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *DnsviewResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dnsview"
}

func (r *DnsviewResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *DnsviewResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data DnsviewResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating dnsview resource")

	// dnsview := dnsviewGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Dnsview.Type(), &dnsview)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create dnsview, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("dnsview-config")

	tflog.Trace(ctx, "Created dnsview resource")

	// Read the updated state back
	r.readDnsviewFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnsviewResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DnsviewResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading dnsview resource")

	r.readDnsviewFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnsviewResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data DnsviewResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating dnsview resource")

	// Create API request body from the model
	// dnsview := dnsviewGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Dnsview.Type(), &dnsview)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update dnsview, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated dnsview resource")

	// Read the updated state back
	r.readDnsviewFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnsviewResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data DnsviewResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting dnsview resource")

	// For dnsview, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted dnsview resource from state")
}

// Helper function to read dnsview data from API
func (r *DnsviewResource) readDnsviewFromApi(ctx context.Context, data *DnsviewResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Dnsview.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read dnsview, got error: %s", err))
		return
	}

	dnsviewSetAttrFromGet(ctx, data, getResponseData)

}
