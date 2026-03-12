package rewritepolicylabel_rewritepolicy_binding

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
var _ resource.Resource = &RewritepolicylabelRewritepolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*RewritepolicylabelRewritepolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*RewritepolicylabelRewritepolicyBindingResource)(nil)

func NewRewritepolicylabelRewritepolicyBindingResource() resource.Resource {
	return &RewritepolicylabelRewritepolicyBindingResource{}
}

// RewritepolicylabelRewritepolicyBindingResource defines the resource implementation.
type RewritepolicylabelRewritepolicyBindingResource struct {
	client *service.NitroClient
}

func (r *RewritepolicylabelRewritepolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *RewritepolicylabelRewritepolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rewritepolicylabel_rewritepolicy_binding"
}

func (r *RewritepolicylabelRewritepolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *RewritepolicylabelRewritepolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data RewritepolicylabelRewritepolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating rewritepolicylabel_rewritepolicy_binding resource")

	// rewritepolicylabel_rewritepolicy_binding := rewritepolicylabel_rewritepolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Rewritepolicylabel_rewritepolicy_binding.Type(), &rewritepolicylabel_rewritepolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create rewritepolicylabel_rewritepolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("rewritepolicylabel_rewritepolicy_binding-config")

	tflog.Trace(ctx, "Created rewritepolicylabel_rewritepolicy_binding resource")

	// Read the updated state back
	r.readRewritepolicylabelRewritepolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RewritepolicylabelRewritepolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data RewritepolicylabelRewritepolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading rewritepolicylabel_rewritepolicy_binding resource")

	r.readRewritepolicylabelRewritepolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RewritepolicylabelRewritepolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data RewritepolicylabelRewritepolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating rewritepolicylabel_rewritepolicy_binding resource")

	// Create API request body from the model
	// rewritepolicylabel_rewritepolicy_binding := rewritepolicylabel_rewritepolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Rewritepolicylabel_rewritepolicy_binding.Type(), &rewritepolicylabel_rewritepolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update rewritepolicylabel_rewritepolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated rewritepolicylabel_rewritepolicy_binding resource")

	// Read the updated state back
	r.readRewritepolicylabelRewritepolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RewritepolicylabelRewritepolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data RewritepolicylabelRewritepolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting rewritepolicylabel_rewritepolicy_binding resource")

	// For rewritepolicylabel_rewritepolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted rewritepolicylabel_rewritepolicy_binding resource from state")
}

// Helper function to read rewritepolicylabel_rewritepolicy_binding data from API
func (r *RewritepolicylabelRewritepolicyBindingResource) readRewritepolicylabelRewritepolicyBindingFromApi(ctx context.Context, data *RewritepolicylabelRewritepolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Rewritepolicylabel_rewritepolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read rewritepolicylabel_rewritepolicy_binding, got error: %s", err))
		return
	}

	rewritepolicylabel_rewritepolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
