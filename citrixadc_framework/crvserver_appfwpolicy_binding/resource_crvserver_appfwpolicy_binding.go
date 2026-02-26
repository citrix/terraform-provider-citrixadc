package crvserver_appfwpolicy_binding

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
var _ resource.Resource = &CrvserverAppfwpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*CrvserverAppfwpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*CrvserverAppfwpolicyBindingResource)(nil)

func NewCrvserverAppfwpolicyBindingResource() resource.Resource {
	return &CrvserverAppfwpolicyBindingResource{}
}

// CrvserverAppfwpolicyBindingResource defines the resource implementation.
type CrvserverAppfwpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *CrvserverAppfwpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *CrvserverAppfwpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_crvserver_appfwpolicy_binding"
}

func (r *CrvserverAppfwpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *CrvserverAppfwpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data CrvserverAppfwpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating crvserver_appfwpolicy_binding resource")

	// crvserver_appfwpolicy_binding := crvserver_appfwpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Crvserver_appfwpolicy_binding.Type(), &crvserver_appfwpolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create crvserver_appfwpolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("crvserver_appfwpolicy_binding-config")

	tflog.Trace(ctx, "Created crvserver_appfwpolicy_binding resource")

	// Read the updated state back
	r.readCrvserverAppfwpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CrvserverAppfwpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data CrvserverAppfwpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading crvserver_appfwpolicy_binding resource")

	r.readCrvserverAppfwpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CrvserverAppfwpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data CrvserverAppfwpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating crvserver_appfwpolicy_binding resource")

	// Create API request body from the model
	// crvserver_appfwpolicy_binding := crvserver_appfwpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Crvserver_appfwpolicy_binding.Type(), &crvserver_appfwpolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update crvserver_appfwpolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated crvserver_appfwpolicy_binding resource")

	// Read the updated state back
	r.readCrvserverAppfwpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CrvserverAppfwpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data CrvserverAppfwpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting crvserver_appfwpolicy_binding resource")

	// For crvserver_appfwpolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted crvserver_appfwpolicy_binding resource from state")
}

// Helper function to read crvserver_appfwpolicy_binding data from API
func (r *CrvserverAppfwpolicyBindingResource) readCrvserverAppfwpolicyBindingFromApi(ctx context.Context, data *CrvserverAppfwpolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Crvserver_appfwpolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read crvserver_appfwpolicy_binding, got error: %s", err))
		return
	}

	crvserver_appfwpolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
