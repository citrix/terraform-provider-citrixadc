package lbvserver_videooptimizationdetectionpolicy_binding

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
var _ resource.Resource = &LbvserverVideooptimizationdetectionpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*LbvserverVideooptimizationdetectionpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*LbvserverVideooptimizationdetectionpolicyBindingResource)(nil)

func NewLbvserverVideooptimizationdetectionpolicyBindingResource() resource.Resource {
	return &LbvserverVideooptimizationdetectionpolicyBindingResource{}
}

// LbvserverVideooptimizationdetectionpolicyBindingResource defines the resource implementation.
type LbvserverVideooptimizationdetectionpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *LbvserverVideooptimizationdetectionpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LbvserverVideooptimizationdetectionpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lbvserver_videooptimizationdetectionpolicy_binding"
}

func (r *LbvserverVideooptimizationdetectionpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LbvserverVideooptimizationdetectionpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LbvserverVideooptimizationdetectionpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lbvserver_videooptimizationdetectionpolicy_binding resource")

	// lbvserver_videooptimizationdetectionpolicy_binding := lbvserver_videooptimizationdetectionpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Lbvserver_videooptimizationdetectionpolicy_binding.Type(), &lbvserver_videooptimizationdetectionpolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lbvserver_videooptimizationdetectionpolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("lbvserver_videooptimizationdetectionpolicy_binding-config")

	tflog.Trace(ctx, "Created lbvserver_videooptimizationdetectionpolicy_binding resource")

	// Read the updated state back
	r.readLbvserverVideooptimizationdetectionpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbvserverVideooptimizationdetectionpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LbvserverVideooptimizationdetectionpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lbvserver_videooptimizationdetectionpolicy_binding resource")

	r.readLbvserverVideooptimizationdetectionpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbvserverVideooptimizationdetectionpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data LbvserverVideooptimizationdetectionpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating lbvserver_videooptimizationdetectionpolicy_binding resource")

	// Create API request body from the model
	// lbvserver_videooptimizationdetectionpolicy_binding := lbvserver_videooptimizationdetectionpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Lbvserver_videooptimizationdetectionpolicy_binding.Type(), &lbvserver_videooptimizationdetectionpolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lbvserver_videooptimizationdetectionpolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated lbvserver_videooptimizationdetectionpolicy_binding resource")

	// Read the updated state back
	r.readLbvserverVideooptimizationdetectionpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbvserverVideooptimizationdetectionpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LbvserverVideooptimizationdetectionpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lbvserver_videooptimizationdetectionpolicy_binding resource")

	// For lbvserver_videooptimizationdetectionpolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted lbvserver_videooptimizationdetectionpolicy_binding resource from state")
}

// Helper function to read lbvserver_videooptimizationdetectionpolicy_binding data from API
func (r *LbvserverVideooptimizationdetectionpolicyBindingResource) readLbvserverVideooptimizationdetectionpolicyBindingFromApi(ctx context.Context, data *LbvserverVideooptimizationdetectionpolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Lbvserver_videooptimizationdetectionpolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lbvserver_videooptimizationdetectionpolicy_binding, got error: %s", err))
		return
	}

	lbvserver_videooptimizationdetectionpolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
