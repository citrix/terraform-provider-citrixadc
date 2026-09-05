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
	var data, config, state CachecontentgroupResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from configuration (for unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating cachecontentgroup resource")

	// Detect attributes removed from config so they can be unset (reverted to
	// their NITRO defaults) after the update. An attribute that changed relative
	// to prior state AND is null in config was removed by the user.
	attributesToUnset := []string{}
	if !data.Alwaysevalpolicies.Equal(state.Alwaysevalpolicies) && config.Alwaysevalpolicies.IsNull() {
		attributesToUnset = append(attributesToUnset, "alwaysevalpolicies")
	}
	if !data.Expireatlastbyte.Equal(state.Expireatlastbyte) && config.Expireatlastbyte.IsNull() {
		attributesToUnset = append(attributesToUnset, "expireatlastbyte")
	}
	if !data.Ignorereloadreq.Equal(state.Ignorereloadreq) && config.Ignorereloadreq.IsNull() {
		attributesToUnset = append(attributesToUnset, "ignorereloadreq")
	}
	if !data.Ignorereqcachinghdrs.Equal(state.Ignorereqcachinghdrs) && config.Ignorereqcachinghdrs.IsNull() {
		attributesToUnset = append(attributesToUnset, "ignorereqcachinghdrs")
	}
	if !data.Insertage.Equal(state.Insertage) && config.Insertage.IsNull() {
		attributesToUnset = append(attributesToUnset, "insertage")
	}
	if !data.Insertetag.Equal(state.Insertetag) && config.Insertetag.IsNull() {
		attributesToUnset = append(attributesToUnset, "insertetag")
	}
	if !data.Insertvia.Equal(state.Insertvia) && config.Insertvia.IsNull() {
		attributesToUnset = append(attributesToUnset, "insertvia")
	}
	if !data.Lazydnsresolve.Equal(state.Lazydnsresolve) && config.Lazydnsresolve.IsNull() {
		attributesToUnset = append(attributesToUnset, "lazydnsresolve")
	}
	if !data.Maxressize.Equal(state.Maxressize) && config.Maxressize.IsNull() {
		attributesToUnset = append(attributesToUnset, "maxressize")
	}
	if !data.Memlimit.Equal(state.Memlimit) && config.Memlimit.IsNull() {
		attributesToUnset = append(attributesToUnset, "memlimit")
	}
	if !data.Minhits.Equal(state.Minhits) && config.Minhits.IsNull() {
		attributesToUnset = append(attributesToUnset, "minhits")
	}
	if !data.Minressize.Equal(state.Minressize) && config.Minressize.IsNull() {
		attributesToUnset = append(attributesToUnset, "minressize")
	}
	if !data.Persistha.Equal(state.Persistha) && config.Persistha.IsNull() {
		attributesToUnset = append(attributesToUnset, "persistha")
	}
	if !data.Pinned.Equal(state.Pinned) && config.Pinned.IsNull() {
		attributesToUnset = append(attributesToUnset, "pinned")
	}
	if !data.Polleverytime.Equal(state.Polleverytime) && config.Polleverytime.IsNull() {
		attributesToUnset = append(attributesToUnset, "polleverytime")
	}
	if !data.Prefetch.Equal(state.Prefetch) && config.Prefetch.IsNull() {
		attributesToUnset = append(attributesToUnset, "prefetch")
	}
	if !data.Quickabortsize.Equal(state.Quickabortsize) && config.Quickabortsize.IsNull() {
		attributesToUnset = append(attributesToUnset, "quickabortsize")
	}
	if !data.Removecookies.Equal(state.Removecookies) && config.Removecookies.IsNull() {
		attributesToUnset = append(attributesToUnset, "removecookies")
	}

	// Build the SET (PUT) payload. Unlike the create payload it EXCLUDES the
	// create-only attr (type) and the flush/GET-only filter attrs so the PUT does
	// not leak create-only params (Pattern 9: add-vs-set payload drift).
	cachecontentgroup := cachecontentgroupGetTheUpdatePayloadFromThePlan(ctx, &data)

	// Make API call
	// Named resource - use UpdateResource. Address by the live name (data.Id).
	cachecontentgroupName := data.Id.ValueString()
	cachecontentgroup.Name = cachecontentgroupName
	_, err := r.client.UpdateResource(service.Cachecontentgroup.Type(), cachecontentgroupName, &cachecontentgroup)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update cachecontentgroup, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Updated cachecontentgroup resource")

	// Unset attributes removed from config so the appliance reverts them to
	// their defaults. Done after the update so any default value the update
	// payload carried for a removed attribute is superseded by the unset.
	unsetIdPayload := map[string]interface{}{
		"name": data.Name.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Cachecontentgroup.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset cachecontentgroup attributes, got error: %s", err))
		return
	}

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
