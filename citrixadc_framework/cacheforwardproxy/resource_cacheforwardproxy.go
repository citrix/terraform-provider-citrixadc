package cacheforwardproxy

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
var _ resource.Resource = &CacheforwardproxyResource{}
var _ resource.ResourceWithConfigure = (*CacheforwardproxyResource)(nil)
var _ resource.ResourceWithImportState = (*CacheforwardproxyResource)(nil)

func NewCacheforwardproxyResource() resource.Resource {
	return &CacheforwardproxyResource{}
}

// CacheforwardproxyResource defines the resource implementation.
type CacheforwardproxyResource struct {
	client *service.NitroClient
}

func (r *CacheforwardproxyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *CacheforwardproxyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cacheforwardproxy"
}

func (r *CacheforwardproxyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *CacheforwardproxyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data CacheforwardproxyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating cacheforwardproxy resource")

	// cacheforwardproxy := cacheforwardproxyGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Cacheforwardproxy.Type(), &cacheforwardproxy)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create cacheforwardproxy, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("cacheforwardproxy-config")

	tflog.Trace(ctx, "Created cacheforwardproxy resource")

	// Read the updated state back
	r.readCacheforwardproxyFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CacheforwardproxyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data CacheforwardproxyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading cacheforwardproxy resource")

	r.readCacheforwardproxyFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CacheforwardproxyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data CacheforwardproxyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating cacheforwardproxy resource")

	// Create API request body from the model
	// cacheforwardproxy := cacheforwardproxyGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Cacheforwardproxy.Type(), &cacheforwardproxy)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update cacheforwardproxy, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated cacheforwardproxy resource")

	// Read the updated state back
	r.readCacheforwardproxyFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CacheforwardproxyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data CacheforwardproxyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting cacheforwardproxy resource")

	// For cacheforwardproxy, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted cacheforwardproxy resource from state")
}

// Helper function to read cacheforwardproxy data from API
func (r *CacheforwardproxyResource) readCacheforwardproxyFromApi(ctx context.Context, data *CacheforwardproxyResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Cacheforwardproxy.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read cacheforwardproxy, got error: %s", err))
		return
	}

	cacheforwardproxySetAttrFromGet(ctx, data, getResponseData)

}
