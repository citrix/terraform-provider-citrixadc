package lbvserver_contentinspectionpolicy_binding

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
var _ resource.Resource = &LbvserverContentinspectionpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*LbvserverContentinspectionpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*LbvserverContentinspectionpolicyBindingResource)(nil)

func NewLbvserverContentinspectionpolicyBindingResource() resource.Resource {
	return &LbvserverContentinspectionpolicyBindingResource{}
}

// LbvserverContentinspectionpolicyBindingResource defines the resource implementation.
type LbvserverContentinspectionpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *LbvserverContentinspectionpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LbvserverContentinspectionpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lbvserver_contentinspectionpolicy_binding"
}

func (r *LbvserverContentinspectionpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LbvserverContentinspectionpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LbvserverContentinspectionpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lbvserver_contentinspectionpolicy_binding resource")

	// lbvserver_contentinspectionpolicy_binding := lbvserver_contentinspectionpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Lbvserver_contentinspectionpolicy_binding.Type(), &lbvserver_contentinspectionpolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lbvserver_contentinspectionpolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("lbvserver_contentinspectionpolicy_binding-config")

	tflog.Trace(ctx, "Created lbvserver_contentinspectionpolicy_binding resource")

	// Read the updated state back
	r.readLbvserverContentinspectionpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbvserverContentinspectionpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LbvserverContentinspectionpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lbvserver_contentinspectionpolicy_binding resource")

	r.readLbvserverContentinspectionpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbvserverContentinspectionpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data LbvserverContentinspectionpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating lbvserver_contentinspectionpolicy_binding resource")

	// Create API request body from the model
	// lbvserver_contentinspectionpolicy_binding := lbvserver_contentinspectionpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Lbvserver_contentinspectionpolicy_binding.Type(), &lbvserver_contentinspectionpolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lbvserver_contentinspectionpolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated lbvserver_contentinspectionpolicy_binding resource")

	// Read the updated state back
	r.readLbvserverContentinspectionpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbvserverContentinspectionpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LbvserverContentinspectionpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lbvserver_contentinspectionpolicy_binding resource")

	// For lbvserver_contentinspectionpolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted lbvserver_contentinspectionpolicy_binding resource from state")
}

// Helper function to read lbvserver_contentinspectionpolicy_binding data from API
func (r *LbvserverContentinspectionpolicyBindingResource) readLbvserverContentinspectionpolicyBindingFromApi(ctx context.Context, data *LbvserverContentinspectionpolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Lbvserver_contentinspectionpolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lbvserver_contentinspectionpolicy_binding, got error: %s", err))
		return
	}

	lbvserver_contentinspectionpolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
