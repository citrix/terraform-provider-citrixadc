package dnsptrrec

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
var _ resource.Resource = &DnsptrrecResource{}
var _ resource.ResourceWithConfigure = (*DnsptrrecResource)(nil)
var _ resource.ResourceWithImportState = (*DnsptrrecResource)(nil)

func NewDnsptrrecResource() resource.Resource {
	return &DnsptrrecResource{}
}

// DnsptrrecResource defines the resource implementation.
type DnsptrrecResource struct {
	client *service.NitroClient
}

func (r *DnsptrrecResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *DnsptrrecResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dnsptrrec"
}

func (r *DnsptrrecResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *DnsptrrecResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data DnsptrrecResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating dnsptrrec resource")

	dnsptrrec := dnsptrrecGetThePayloadFromtheConfig(ctx, &data)

	dnsptrrecName := data.Reversedomain.ValueString()

	// Make API call
	_, err := r.client.AddResource(service.Dnsptrrec.Type(), dnsptrrecName, &dnsptrrec)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create dnsptrrec, got error: %s", err))
		return
	}

	// Set ID to reversedomain
	data.Id = types.StringValue(dnsptrrecName)

	tflog.Trace(ctx, "Created dnsptrrec resource")

	// Read the updated state back
	r.readDnsptrrecFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnsptrrecResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DnsptrrecResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading dnsptrrec resource")

	r.readDnsptrrecFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnsptrrecResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data DnsptrrecResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating dnsptrrec resource")

	// Create API request body from the model
	// dnsptrrec := dnsptrrecGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Dnsptrrec.Type(), &dnsptrrec)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update dnsptrrec, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated dnsptrrec resource")

	// Read the updated state back
	r.readDnsptrrecFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnsptrrecResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data DnsptrrecResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting dnsptrrec resource")

	dnsptrrecName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Dnsptrrec.Type(), dnsptrrecName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete dnsptrrec, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted dnsptrrec resource")
}

// Helper function to read dnsptrrec data from API
func (r *DnsptrrecResource) readDnsptrrecFromApi(ctx context.Context, data *DnsptrrecResourceModel, diags *diag.Diagnostics) {
	dnsptrrecName := data.Id.ValueString()
	getResponseData, err := r.client.FindResource(service.Dnsptrrec.Type(), dnsptrrecName)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read dnsptrrec, got error: %s", err))
		return
	}

	dnsptrrecSetAttrFromGet(ctx, data, getResponseData)

}
