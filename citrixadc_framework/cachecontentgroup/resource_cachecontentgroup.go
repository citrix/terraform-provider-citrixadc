package cachecontentgroup

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
var _ resource.Resource = &CachecontentgroupResource{}
var _ resource.ResourceWithConfigure = (*CachecontentgroupResource)(nil)
var _ resource.ResourceWithImportState = (*CachecontentgroupResource)(nil)

func NewCachecontentgroupResource() resource.Resource {
	return &CachecontentgroupResource{}
}

// CachecontentgroupResource defines the resource implementation.
type CachecontentgroupResource struct {
	client *service.NitroClient
}

func (r *CachecontentgroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *CachecontentgroupResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cachecontentgroup"
}

func (r *CachecontentgroupResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *CachecontentgroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data CachecontentgroupResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating cachecontentgroup resource")

	// Create API request body from the model
	cachecontentgroup := cachecontentgroupGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	cachecontentgroupName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Cachecontentgroup.Type(), cachecontentgroupName, &cachecontentgroup)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create cachecontentgroup, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created cachecontentgroup resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(cachecontentgroupName)

	// Read the updated state back
	if !r.readCachecontentgroupFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "cachecontentgroup not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CachecontentgroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data CachecontentgroupResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading cachecontentgroup resource")

	found := r.readCachecontentgroupFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *CachecontentgroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state CachecontentgroupResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating cachecontentgroup resource")

	// Create API request body from the model
	cachecontentgroup := cachecontentgroupGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Named resource - use UpdateResource
	cachecontentgroupName := data.Name.ValueString()
	_, err := r.client.UpdateResource(service.Cachecontentgroup.Type(), cachecontentgroupName, &cachecontentgroup)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update cachecontentgroup, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Updated cachecontentgroup resource")

	// Read the updated state back
	if !r.readCachecontentgroupFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "cachecontentgroup not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CachecontentgroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data CachecontentgroupResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting cachecontentgroup resource")

	// Named resource - delete using DeleteResource
	cachecontentgroupName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Cachecontentgroup.Type(), cachecontentgroupName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete cachecontentgroup, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted cachecontentgroup resource")
}

// Helper function to read cachecontentgroup data from API
func (r *CachecontentgroupResource) readCachecontentgroupFromApi(ctx context.Context, data *CachecontentgroupResourceModel, diags *diag.Diagnostics) bool {
	// Case 2: Find with single ID attribute - ID is the plain value (name)
	cachecontentgroupName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Cachecontentgroup.Type(), cachecontentgroupName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read cachecontentgroup, got error: %s", err))
		return false
	}

	cachecontentgroupSetAttrFromGet(ctx, data, getResponseData)

	return true
}
