package lbgroup_lbvserver_binding

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
var _ resource.Resource = &LbgroupLbvserverBindingResource{}
var _ resource.ResourceWithConfigure = (*LbgroupLbvserverBindingResource)(nil)
var _ resource.ResourceWithImportState = (*LbgroupLbvserverBindingResource)(nil)

func NewLbgroupLbvserverBindingResource() resource.Resource {
	return &LbgroupLbvserverBindingResource{}
}

// LbgroupLbvserverBindingResource defines the resource implementation.
type LbgroupLbvserverBindingResource struct {
	client *service.NitroClient
}

func (r *LbgroupLbvserverBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LbgroupLbvserverBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lbgroup_lbvserver_binding"
}

func (r *LbgroupLbvserverBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LbgroupLbvserverBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LbgroupLbvserverBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lbgroup_lbvserver_binding resource")

	// lbgroup_lbvserver_binding := lbgroup_lbvserver_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Lbgroup_lbvserver_binding.Type(), &lbgroup_lbvserver_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lbgroup_lbvserver_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("lbgroup_lbvserver_binding-config")

	tflog.Trace(ctx, "Created lbgroup_lbvserver_binding resource")

	// Read the updated state back
	r.readLbgroupLbvserverBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbgroupLbvserverBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LbgroupLbvserverBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lbgroup_lbvserver_binding resource")

	r.readLbgroupLbvserverBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbgroupLbvserverBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data LbgroupLbvserverBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating lbgroup_lbvserver_binding resource")

	// Create API request body from the model
	// lbgroup_lbvserver_binding := lbgroup_lbvserver_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Lbgroup_lbvserver_binding.Type(), &lbgroup_lbvserver_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lbgroup_lbvserver_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated lbgroup_lbvserver_binding resource")

	// Read the updated state back
	r.readLbgroupLbvserverBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbgroupLbvserverBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LbgroupLbvserverBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lbgroup_lbvserver_binding resource")

	// For lbgroup_lbvserver_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted lbgroup_lbvserver_binding resource from state")
}

// Helper function to read lbgroup_lbvserver_binding data from API
func (r *LbgroupLbvserverBindingResource) readLbgroupLbvserverBindingFromApi(ctx context.Context, data *LbgroupLbvserverBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Lbgroup_lbvserver_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lbgroup_lbvserver_binding, got error: %s", err))
		return
	}

	lbgroup_lbvserver_bindingSetAttrFromGet(ctx, data, getResponseData)

}
