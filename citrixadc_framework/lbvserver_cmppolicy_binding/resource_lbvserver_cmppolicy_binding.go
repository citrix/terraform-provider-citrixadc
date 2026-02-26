package lbvserver_cmppolicy_binding

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
var _ resource.Resource = &LbvserverCmppolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*LbvserverCmppolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*LbvserverCmppolicyBindingResource)(nil)

func NewLbvserverCmppolicyBindingResource() resource.Resource {
	return &LbvserverCmppolicyBindingResource{}
}

// LbvserverCmppolicyBindingResource defines the resource implementation.
type LbvserverCmppolicyBindingResource struct {
	client *service.NitroClient
}

func (r *LbvserverCmppolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LbvserverCmppolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lbvserver_cmppolicy_binding"
}

func (r *LbvserverCmppolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LbvserverCmppolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LbvserverCmppolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lbvserver_cmppolicy_binding resource")

	// lbvserver_cmppolicy_binding := lbvserver_cmppolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Lbvserver_cmppolicy_binding.Type(), &lbvserver_cmppolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lbvserver_cmppolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("lbvserver_cmppolicy_binding-config")

	tflog.Trace(ctx, "Created lbvserver_cmppolicy_binding resource")

	// Read the updated state back
	r.readLbvserverCmppolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbvserverCmppolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LbvserverCmppolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lbvserver_cmppolicy_binding resource")

	r.readLbvserverCmppolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbvserverCmppolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data LbvserverCmppolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating lbvserver_cmppolicy_binding resource")

	// Create API request body from the model
	// lbvserver_cmppolicy_binding := lbvserver_cmppolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Lbvserver_cmppolicy_binding.Type(), &lbvserver_cmppolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lbvserver_cmppolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated lbvserver_cmppolicy_binding resource")

	// Read the updated state back
	r.readLbvserverCmppolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbvserverCmppolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LbvserverCmppolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lbvserver_cmppolicy_binding resource")

	// For lbvserver_cmppolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted lbvserver_cmppolicy_binding resource from state")
}

// Helper function to read lbvserver_cmppolicy_binding data from API
func (r *LbvserverCmppolicyBindingResource) readLbvserverCmppolicyBindingFromApi(ctx context.Context, data *LbvserverCmppolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Lbvserver_cmppolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lbvserver_cmppolicy_binding, got error: %s", err))
		return
	}

	lbvserver_cmppolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
