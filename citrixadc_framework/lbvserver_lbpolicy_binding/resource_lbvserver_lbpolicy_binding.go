package lbvserver_lbpolicy_binding

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
var _ resource.Resource = &LbvserverLbpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*LbvserverLbpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*LbvserverLbpolicyBindingResource)(nil)

func NewLbvserverLbpolicyBindingResource() resource.Resource {
	return &LbvserverLbpolicyBindingResource{}
}

// LbvserverLbpolicyBindingResource defines the resource implementation.
type LbvserverLbpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *LbvserverLbpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LbvserverLbpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lbvserver_lbpolicy_binding"
}

func (r *LbvserverLbpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LbvserverLbpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LbvserverLbpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lbvserver_lbpolicy_binding resource")

	// lbvserver_lbpolicy_binding := lbvserver_lbpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Lbvserver_lbpolicy_binding.Type(), &lbvserver_lbpolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lbvserver_lbpolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("lbvserver_lbpolicy_binding-config")

	tflog.Trace(ctx, "Created lbvserver_lbpolicy_binding resource")

	// Read the updated state back
	r.readLbvserverLbpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbvserverLbpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LbvserverLbpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lbvserver_lbpolicy_binding resource")

	r.readLbvserverLbpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbvserverLbpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data LbvserverLbpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating lbvserver_lbpolicy_binding resource")

	// Create API request body from the model
	// lbvserver_lbpolicy_binding := lbvserver_lbpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Lbvserver_lbpolicy_binding.Type(), &lbvserver_lbpolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lbvserver_lbpolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated lbvserver_lbpolicy_binding resource")

	// Read the updated state back
	r.readLbvserverLbpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbvserverLbpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LbvserverLbpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lbvserver_lbpolicy_binding resource")

	// For lbvserver_lbpolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted lbvserver_lbpolicy_binding resource from state")
}

// Helper function to read lbvserver_lbpolicy_binding data from API
func (r *LbvserverLbpolicyBindingResource) readLbvserverLbpolicyBindingFromApi(ctx context.Context, data *LbvserverLbpolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Lbvserver_lbpolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lbvserver_lbpolicy_binding, got error: %s", err))
		return
	}

	lbvserver_lbpolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
