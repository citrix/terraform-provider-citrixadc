package lbvserver_authorizationpolicy_binding

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
var _ resource.Resource = &LbvserverAuthorizationpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*LbvserverAuthorizationpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*LbvserverAuthorizationpolicyBindingResource)(nil)

func NewLbvserverAuthorizationpolicyBindingResource() resource.Resource {
	return &LbvserverAuthorizationpolicyBindingResource{}
}

// LbvserverAuthorizationpolicyBindingResource defines the resource implementation.
type LbvserverAuthorizationpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *LbvserverAuthorizationpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LbvserverAuthorizationpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lbvserver_authorizationpolicy_binding"
}

func (r *LbvserverAuthorizationpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LbvserverAuthorizationpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LbvserverAuthorizationpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lbvserver_authorizationpolicy_binding resource")

	// lbvserver_authorizationpolicy_binding := lbvserver_authorizationpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Lbvserver_authorizationpolicy_binding.Type(), &lbvserver_authorizationpolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lbvserver_authorizationpolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("lbvserver_authorizationpolicy_binding-config")

	tflog.Trace(ctx, "Created lbvserver_authorizationpolicy_binding resource")

	// Read the updated state back
	r.readLbvserverAuthorizationpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbvserverAuthorizationpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LbvserverAuthorizationpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lbvserver_authorizationpolicy_binding resource")

	r.readLbvserverAuthorizationpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbvserverAuthorizationpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data LbvserverAuthorizationpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating lbvserver_authorizationpolicy_binding resource")

	// Create API request body from the model
	// lbvserver_authorizationpolicy_binding := lbvserver_authorizationpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Lbvserver_authorizationpolicy_binding.Type(), &lbvserver_authorizationpolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lbvserver_authorizationpolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated lbvserver_authorizationpolicy_binding resource")

	// Read the updated state back
	r.readLbvserverAuthorizationpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbvserverAuthorizationpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LbvserverAuthorizationpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lbvserver_authorizationpolicy_binding resource")

	// For lbvserver_authorizationpolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted lbvserver_authorizationpolicy_binding resource from state")
}

// Helper function to read lbvserver_authorizationpolicy_binding data from API
func (r *LbvserverAuthorizationpolicyBindingResource) readLbvserverAuthorizationpolicyBindingFromApi(ctx context.Context, data *LbvserverAuthorizationpolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Lbvserver_authorizationpolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lbvserver_authorizationpolicy_binding, got error: %s", err))
		return
	}

	lbvserver_authorizationpolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
