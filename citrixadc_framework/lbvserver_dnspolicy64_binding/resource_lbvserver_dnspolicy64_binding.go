package lbvserver_dnspolicy64_binding

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
var _ resource.Resource = &LbvserverDnspolicy64BindingResource{}
var _ resource.ResourceWithConfigure = (*LbvserverDnspolicy64BindingResource)(nil)
var _ resource.ResourceWithImportState = (*LbvserverDnspolicy64BindingResource)(nil)

func NewLbvserverDnspolicy64BindingResource() resource.Resource {
	return &LbvserverDnspolicy64BindingResource{}
}

// LbvserverDnspolicy64BindingResource defines the resource implementation.
type LbvserverDnspolicy64BindingResource struct {
	client *service.NitroClient
}

func (r *LbvserverDnspolicy64BindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LbvserverDnspolicy64BindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lbvserver_dnspolicy64_binding"
}

func (r *LbvserverDnspolicy64BindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LbvserverDnspolicy64BindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LbvserverDnspolicy64BindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lbvserver_dnspolicy64_binding resource")

	// lbvserver_dnspolicy64_binding := lbvserver_dnspolicy64_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Lbvserver_dnspolicy64_binding.Type(), &lbvserver_dnspolicy64_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lbvserver_dnspolicy64_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("lbvserver_dnspolicy64_binding-config")

	tflog.Trace(ctx, "Created lbvserver_dnspolicy64_binding resource")

	// Read the updated state back
	r.readLbvserverDnspolicy64BindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbvserverDnspolicy64BindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LbvserverDnspolicy64BindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lbvserver_dnspolicy64_binding resource")

	r.readLbvserverDnspolicy64BindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbvserverDnspolicy64BindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data LbvserverDnspolicy64BindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating lbvserver_dnspolicy64_binding resource")

	// Create API request body from the model
	// lbvserver_dnspolicy64_binding := lbvserver_dnspolicy64_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Lbvserver_dnspolicy64_binding.Type(), &lbvserver_dnspolicy64_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lbvserver_dnspolicy64_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated lbvserver_dnspolicy64_binding resource")

	// Read the updated state back
	r.readLbvserverDnspolicy64BindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbvserverDnspolicy64BindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LbvserverDnspolicy64BindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lbvserver_dnspolicy64_binding resource")

	// For lbvserver_dnspolicy64_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted lbvserver_dnspolicy64_binding resource from state")
}

// Helper function to read lbvserver_dnspolicy64_binding data from API
func (r *LbvserverDnspolicy64BindingResource) readLbvserverDnspolicy64BindingFromApi(ctx context.Context, data *LbvserverDnspolicy64BindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Lbvserver_dnspolicy64_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lbvserver_dnspolicy64_binding, got error: %s", err))
		return
	}

	lbvserver_dnspolicy64_bindingSetAttrFromGet(ctx, data, getResponseData)

}
