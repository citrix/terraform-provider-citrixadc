package dnsaction64

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
var _ resource.Resource = &Dnsaction64Resource{}
var _ resource.ResourceWithConfigure = (*Dnsaction64Resource)(nil)
var _ resource.ResourceWithImportState = (*Dnsaction64Resource)(nil)

func NewDnsaction64Resource() resource.Resource {
	return &Dnsaction64Resource{}
}

// Dnsaction64Resource defines the resource implementation.
type Dnsaction64Resource struct {
	client *service.NitroClient
}

func (r *Dnsaction64Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *Dnsaction64Resource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dnsaction64"
}

func (r *Dnsaction64Resource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *Dnsaction64Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data Dnsaction64ResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating dnsaction64 resource")

	// dnsaction64 := dnsaction64GetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Dnsaction64.Type(), &dnsaction64)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create dnsaction64, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("dnsaction64-config")

	tflog.Trace(ctx, "Created dnsaction64 resource")

	// Read the updated state back
	r.readDnsaction64FromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Dnsaction64Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data Dnsaction64ResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading dnsaction64 resource")

	r.readDnsaction64FromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Dnsaction64Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data Dnsaction64ResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating dnsaction64 resource")

	// Create API request body from the model
	// dnsaction64 := dnsaction64GetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Dnsaction64.Type(), &dnsaction64)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update dnsaction64, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated dnsaction64 resource")

	// Read the updated state back
	r.readDnsaction64FromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Dnsaction64Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data Dnsaction64ResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting dnsaction64 resource")

	// For dnsaction64, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted dnsaction64 resource from state")
}

// Helper function to read dnsaction64 data from API
func (r *Dnsaction64Resource) readDnsaction64FromApi(ctx context.Context, data *Dnsaction64ResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Dnsaction64.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read dnsaction64, got error: %s", err))
		return
	}

	dnsaction64SetAttrFromGet(ctx, data, getResponseData)

}
