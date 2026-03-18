package rewriteglobal_rewritepolicy_binding

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
var _ resource.Resource = &RewriteglobalRewritepolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*RewriteglobalRewritepolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*RewriteglobalRewritepolicyBindingResource)(nil)

func NewRewriteglobalRewritepolicyBindingResource() resource.Resource {
	return &RewriteglobalRewritepolicyBindingResource{}
}

// RewriteglobalRewritepolicyBindingResource defines the resource implementation.
type RewriteglobalRewritepolicyBindingResource struct {
	client *service.NitroClient
}

func (r *RewriteglobalRewritepolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *RewriteglobalRewritepolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rewriteglobal_rewritepolicy_binding"
}

func (r *RewriteglobalRewritepolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *RewriteglobalRewritepolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data RewriteglobalRewritepolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating rewriteglobal_rewritepolicy_binding resource")

	// rewriteglobal_rewritepolicy_binding := rewriteglobal_rewritepolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Rewriteglobal_rewritepolicy_binding.Type(), &rewriteglobal_rewritepolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create rewriteglobal_rewritepolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("rewriteglobal_rewritepolicy_binding-config")

	tflog.Trace(ctx, "Created rewriteglobal_rewritepolicy_binding resource")

	// Read the updated state back
	r.readRewriteglobalRewritepolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RewriteglobalRewritepolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data RewriteglobalRewritepolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading rewriteglobal_rewritepolicy_binding resource")

	r.readRewriteglobalRewritepolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RewriteglobalRewritepolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data RewriteglobalRewritepolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating rewriteglobal_rewritepolicy_binding resource")

	// Create API request body from the model
	// rewriteglobal_rewritepolicy_binding := rewriteglobal_rewritepolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Rewriteglobal_rewritepolicy_binding.Type(), &rewriteglobal_rewritepolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update rewriteglobal_rewritepolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated rewriteglobal_rewritepolicy_binding resource")

	// Read the updated state back
	r.readRewriteglobalRewritepolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RewriteglobalRewritepolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data RewriteglobalRewritepolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting rewriteglobal_rewritepolicy_binding resource")

	// For rewriteglobal_rewritepolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted rewriteglobal_rewritepolicy_binding resource from state")
}

// Helper function to read rewriteglobal_rewritepolicy_binding data from API
func (r *RewriteglobalRewritepolicyBindingResource) readRewriteglobalRewritepolicyBindingFromApi(ctx context.Context, data *RewriteglobalRewritepolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Rewriteglobal_rewritepolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read rewriteglobal_rewritepolicy_binding, got error: %s", err))
		return
	}

	rewriteglobal_rewritepolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
