package cachepolicylabel

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
var _ resource.Resource = &CachepolicylabelResource{}
var _ resource.ResourceWithConfigure = (*CachepolicylabelResource)(nil)
var _ resource.ResourceWithImportState = (*CachepolicylabelResource)(nil)

func NewCachepolicylabelResource() resource.Resource {
	return &CachepolicylabelResource{}
}

// CachepolicylabelResource defines the resource implementation.
type CachepolicylabelResource struct {
	client *service.NitroClient
}

func (r *CachepolicylabelResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *CachepolicylabelResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cachepolicylabel"
}

func (r *CachepolicylabelResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *CachepolicylabelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data CachepolicylabelResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating cachepolicylabel resource")

	// cachepolicylabel := cachepolicylabelGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Cachepolicylabel.Type(), &cachepolicylabel)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create cachepolicylabel, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("cachepolicylabel-config")

	tflog.Trace(ctx, "Created cachepolicylabel resource")

	// Read the updated state back
	r.readCachepolicylabelFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CachepolicylabelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data CachepolicylabelResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading cachepolicylabel resource")

	r.readCachepolicylabelFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CachepolicylabelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data CachepolicylabelResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating cachepolicylabel resource")

	// Create API request body from the model
	// cachepolicylabel := cachepolicylabelGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Cachepolicylabel.Type(), &cachepolicylabel)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update cachepolicylabel, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated cachepolicylabel resource")

	// Read the updated state back
	r.readCachepolicylabelFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CachepolicylabelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data CachepolicylabelResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting cachepolicylabel resource")

	// For cachepolicylabel, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted cachepolicylabel resource from state")
}

// Helper function to read cachepolicylabel data from API
func (r *CachepolicylabelResource) readCachepolicylabelFromApi(ctx context.Context, data *CachepolicylabelResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Cachepolicylabel.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read cachepolicylabel, got error: %s", err))
		return
	}

	cachepolicylabelSetAttrFromGet(ctx, data, getResponseData)

}
