package lbvserver_videooptimizationpacingpolicy_binding

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
var _ resource.Resource = &LbvserverVideooptimizationpacingpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*LbvserverVideooptimizationpacingpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*LbvserverVideooptimizationpacingpolicyBindingResource)(nil)

func NewLbvserverVideooptimizationpacingpolicyBindingResource() resource.Resource {
	return &LbvserverVideooptimizationpacingpolicyBindingResource{}
}

// LbvserverVideooptimizationpacingpolicyBindingResource defines the resource implementation.
type LbvserverVideooptimizationpacingpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *LbvserverVideooptimizationpacingpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LbvserverVideooptimizationpacingpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lbvserver_videooptimizationpacingpolicy_binding"
}

func (r *LbvserverVideooptimizationpacingpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LbvserverVideooptimizationpacingpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LbvserverVideooptimizationpacingpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lbvserver_videooptimizationpacingpolicy_binding resource")

	// lbvserver_videooptimizationpacingpolicy_binding := lbvserver_videooptimizationpacingpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Lbvserver_videooptimizationpacingpolicy_binding.Type(), &lbvserver_videooptimizationpacingpolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lbvserver_videooptimizationpacingpolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("lbvserver_videooptimizationpacingpolicy_binding-config")

	tflog.Trace(ctx, "Created lbvserver_videooptimizationpacingpolicy_binding resource")

	// Read the updated state back
	r.readLbvserverVideooptimizationpacingpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbvserverVideooptimizationpacingpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LbvserverVideooptimizationpacingpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lbvserver_videooptimizationpacingpolicy_binding resource")

	r.readLbvserverVideooptimizationpacingpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbvserverVideooptimizationpacingpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data LbvserverVideooptimizationpacingpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating lbvserver_videooptimizationpacingpolicy_binding resource")

	// Create API request body from the model
	// lbvserver_videooptimizationpacingpolicy_binding := lbvserver_videooptimizationpacingpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Lbvserver_videooptimizationpacingpolicy_binding.Type(), &lbvserver_videooptimizationpacingpolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lbvserver_videooptimizationpacingpolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated lbvserver_videooptimizationpacingpolicy_binding resource")

	// Read the updated state back
	r.readLbvserverVideooptimizationpacingpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbvserverVideooptimizationpacingpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LbvserverVideooptimizationpacingpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lbvserver_videooptimizationpacingpolicy_binding resource")

	// For lbvserver_videooptimizationpacingpolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted lbvserver_videooptimizationpacingpolicy_binding resource from state")
}

// Helper function to read lbvserver_videooptimizationpacingpolicy_binding data from API
func (r *LbvserverVideooptimizationpacingpolicyBindingResource) readLbvserverVideooptimizationpacingpolicyBindingFromApi(ctx context.Context, data *LbvserverVideooptimizationpacingpolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Lbvserver_videooptimizationpacingpolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lbvserver_videooptimizationpacingpolicy_binding, got error: %s", err))
		return
	}

	lbvserver_videooptimizationpacingpolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
