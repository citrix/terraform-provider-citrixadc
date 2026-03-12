package cachepolicylabel_cachepolicy_binding

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
var _ resource.Resource = &CachepolicylabelCachepolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*CachepolicylabelCachepolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*CachepolicylabelCachepolicyBindingResource)(nil)

func NewCachepolicylabelCachepolicyBindingResource() resource.Resource {
	return &CachepolicylabelCachepolicyBindingResource{}
}

// CachepolicylabelCachepolicyBindingResource defines the resource implementation.
type CachepolicylabelCachepolicyBindingResource struct {
	client *service.NitroClient
}

func (r *CachepolicylabelCachepolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *CachepolicylabelCachepolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cachepolicylabel_cachepolicy_binding"
}

func (r *CachepolicylabelCachepolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *CachepolicylabelCachepolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data CachepolicylabelCachepolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating cachepolicylabel_cachepolicy_binding resource")

	// cachepolicylabel_cachepolicy_binding := cachepolicylabel_cachepolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Cachepolicylabel_cachepolicy_binding.Type(), &cachepolicylabel_cachepolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create cachepolicylabel_cachepolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("cachepolicylabel_cachepolicy_binding-config")

	tflog.Trace(ctx, "Created cachepolicylabel_cachepolicy_binding resource")

	// Read the updated state back
	r.readCachepolicylabelCachepolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CachepolicylabelCachepolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data CachepolicylabelCachepolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading cachepolicylabel_cachepolicy_binding resource")

	r.readCachepolicylabelCachepolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CachepolicylabelCachepolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data CachepolicylabelCachepolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating cachepolicylabel_cachepolicy_binding resource")

	// Create API request body from the model
	// cachepolicylabel_cachepolicy_binding := cachepolicylabel_cachepolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Cachepolicylabel_cachepolicy_binding.Type(), &cachepolicylabel_cachepolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update cachepolicylabel_cachepolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated cachepolicylabel_cachepolicy_binding resource")

	// Read the updated state back
	r.readCachepolicylabelCachepolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CachepolicylabelCachepolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data CachepolicylabelCachepolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting cachepolicylabel_cachepolicy_binding resource")

	// For cachepolicylabel_cachepolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted cachepolicylabel_cachepolicy_binding resource from state")
}

// Helper function to read cachepolicylabel_cachepolicy_binding data from API
func (r *CachepolicylabelCachepolicyBindingResource) readCachepolicylabelCachepolicyBindingFromApi(ctx context.Context, data *CachepolicylabelCachepolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Cachepolicylabel_cachepolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read cachepolicylabel_cachepolicy_binding, got error: %s", err))
		return
	}

	cachepolicylabel_cachepolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
