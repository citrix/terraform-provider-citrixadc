package lbvserver_tmtrafficpolicy_binding

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
var _ resource.Resource = &LbvserverTmtrafficpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*LbvserverTmtrafficpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*LbvserverTmtrafficpolicyBindingResource)(nil)

func NewLbvserverTmtrafficpolicyBindingResource() resource.Resource {
	return &LbvserverTmtrafficpolicyBindingResource{}
}

// LbvserverTmtrafficpolicyBindingResource defines the resource implementation.
type LbvserverTmtrafficpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *LbvserverTmtrafficpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LbvserverTmtrafficpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lbvserver_tmtrafficpolicy_binding"
}

func (r *LbvserverTmtrafficpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LbvserverTmtrafficpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LbvserverTmtrafficpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lbvserver_tmtrafficpolicy_binding resource")

	// lbvserver_tmtrafficpolicy_binding := lbvserver_tmtrafficpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Lbvserver_tmtrafficpolicy_binding.Type(), &lbvserver_tmtrafficpolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lbvserver_tmtrafficpolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("lbvserver_tmtrafficpolicy_binding-config")

	tflog.Trace(ctx, "Created lbvserver_tmtrafficpolicy_binding resource")

	// Read the updated state back
	r.readLbvserverTmtrafficpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbvserverTmtrafficpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LbvserverTmtrafficpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lbvserver_tmtrafficpolicy_binding resource")

	r.readLbvserverTmtrafficpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbvserverTmtrafficpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data LbvserverTmtrafficpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating lbvserver_tmtrafficpolicy_binding resource")

	// Create API request body from the model
	// lbvserver_tmtrafficpolicy_binding := lbvserver_tmtrafficpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Lbvserver_tmtrafficpolicy_binding.Type(), &lbvserver_tmtrafficpolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lbvserver_tmtrafficpolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated lbvserver_tmtrafficpolicy_binding resource")

	// Read the updated state back
	r.readLbvserverTmtrafficpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbvserverTmtrafficpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LbvserverTmtrafficpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lbvserver_tmtrafficpolicy_binding resource")

	// For lbvserver_tmtrafficpolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted lbvserver_tmtrafficpolicy_binding resource from state")
}

// Helper function to read lbvserver_tmtrafficpolicy_binding data from API
func (r *LbvserverTmtrafficpolicyBindingResource) readLbvserverTmtrafficpolicyBindingFromApi(ctx context.Context, data *LbvserverTmtrafficpolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Lbvserver_tmtrafficpolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lbvserver_tmtrafficpolicy_binding, got error: %s", err))
		return
	}

	lbvserver_tmtrafficpolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
