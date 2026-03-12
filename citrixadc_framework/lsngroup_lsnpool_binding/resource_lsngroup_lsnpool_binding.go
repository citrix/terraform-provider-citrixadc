package lsngroup_lsnpool_binding

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
var _ resource.Resource = &LsngroupLsnpoolBindingResource{}
var _ resource.ResourceWithConfigure = (*LsngroupLsnpoolBindingResource)(nil)
var _ resource.ResourceWithImportState = (*LsngroupLsnpoolBindingResource)(nil)

func NewLsngroupLsnpoolBindingResource() resource.Resource {
	return &LsngroupLsnpoolBindingResource{}
}

// LsngroupLsnpoolBindingResource defines the resource implementation.
type LsngroupLsnpoolBindingResource struct {
	client *service.NitroClient
}

func (r *LsngroupLsnpoolBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LsngroupLsnpoolBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lsngroup_lsnpool_binding"
}

func (r *LsngroupLsnpoolBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LsngroupLsnpoolBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LsngroupLsnpoolBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lsngroup_lsnpool_binding resource")

	// lsngroup_lsnpool_binding := lsngroup_lsnpool_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Lsngroup_lsnpool_binding.Type(), &lsngroup_lsnpool_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lsngroup_lsnpool_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("lsngroup_lsnpool_binding-config")

	tflog.Trace(ctx, "Created lsngroup_lsnpool_binding resource")

	// Read the updated state back
	r.readLsngroupLsnpoolBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsngroupLsnpoolBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LsngroupLsnpoolBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lsngroup_lsnpool_binding resource")

	r.readLsngroupLsnpoolBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsngroupLsnpoolBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data LsngroupLsnpoolBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating lsngroup_lsnpool_binding resource")

	// Create API request body from the model
	// lsngroup_lsnpool_binding := lsngroup_lsnpool_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Lsngroup_lsnpool_binding.Type(), &lsngroup_lsnpool_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lsngroup_lsnpool_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated lsngroup_lsnpool_binding resource")

	// Read the updated state back
	r.readLsngroupLsnpoolBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsngroupLsnpoolBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LsngroupLsnpoolBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lsngroup_lsnpool_binding resource")

	// For lsngroup_lsnpool_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted lsngroup_lsnpool_binding resource from state")
}

// Helper function to read lsngroup_lsnpool_binding data from API
func (r *LsngroupLsnpoolBindingResource) readLsngroupLsnpoolBindingFromApi(ctx context.Context, data *LsngroupLsnpoolBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Lsngroup_lsnpool_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lsngroup_lsnpool_binding, got error: %s", err))
		return
	}

	lsngroup_lsnpool_bindingSetAttrFromGet(ctx, data, getResponseData)

}
