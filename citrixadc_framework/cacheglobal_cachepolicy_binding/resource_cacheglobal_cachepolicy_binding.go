package cacheglobal_cachepolicy_binding

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
var _ resource.Resource = &CacheglobalCachepolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*CacheglobalCachepolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*CacheglobalCachepolicyBindingResource)(nil)

func NewCacheglobalCachepolicyBindingResource() resource.Resource {
	return &CacheglobalCachepolicyBindingResource{}
}

// CacheglobalCachepolicyBindingResource defines the resource implementation.
type CacheglobalCachepolicyBindingResource struct {
	client *service.NitroClient
}

func (r *CacheglobalCachepolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *CacheglobalCachepolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cacheglobal_cachepolicy_binding"
}

func (r *CacheglobalCachepolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *CacheglobalCachepolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data CacheglobalCachepolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating cacheglobal_cachepolicy_binding resource")

	// cacheglobal_cachepolicy_binding := cacheglobal_cachepolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Cacheglobal_cachepolicy_binding.Type(), &cacheglobal_cachepolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create cacheglobal_cachepolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("cacheglobal_cachepolicy_binding-config")

	tflog.Trace(ctx, "Created cacheglobal_cachepolicy_binding resource")

	// Read the updated state back
	r.readCacheglobalCachepolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CacheglobalCachepolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data CacheglobalCachepolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading cacheglobal_cachepolicy_binding resource")

	r.readCacheglobalCachepolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CacheglobalCachepolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data CacheglobalCachepolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating cacheglobal_cachepolicy_binding resource")

	// Create API request body from the model
	// cacheglobal_cachepolicy_binding := cacheglobal_cachepolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Cacheglobal_cachepolicy_binding.Type(), &cacheglobal_cachepolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update cacheglobal_cachepolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated cacheglobal_cachepolicy_binding resource")

	// Read the updated state back
	r.readCacheglobalCachepolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CacheglobalCachepolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data CacheglobalCachepolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting cacheglobal_cachepolicy_binding resource")

	// For cacheglobal_cachepolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted cacheglobal_cachepolicy_binding resource from state")
}

// Helper function to read cacheglobal_cachepolicy_binding data from API
func (r *CacheglobalCachepolicyBindingResource) readCacheglobalCachepolicyBindingFromApi(ctx context.Context, data *CacheglobalCachepolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Cacheglobal_cachepolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read cacheglobal_cachepolicy_binding, got error: %s", err))
		return
	}

	cacheglobal_cachepolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
