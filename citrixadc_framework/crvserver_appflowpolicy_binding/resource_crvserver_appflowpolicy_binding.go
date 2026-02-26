package crvserver_appflowpolicy_binding

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
var _ resource.Resource = &CrvserverAppflowpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*CrvserverAppflowpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*CrvserverAppflowpolicyBindingResource)(nil)

func NewCrvserverAppflowpolicyBindingResource() resource.Resource {
	return &CrvserverAppflowpolicyBindingResource{}
}

// CrvserverAppflowpolicyBindingResource defines the resource implementation.
type CrvserverAppflowpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *CrvserverAppflowpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *CrvserverAppflowpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_crvserver_appflowpolicy_binding"
}

func (r *CrvserverAppflowpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *CrvserverAppflowpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data CrvserverAppflowpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating crvserver_appflowpolicy_binding resource")

	// crvserver_appflowpolicy_binding := crvserver_appflowpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Crvserver_appflowpolicy_binding.Type(), &crvserver_appflowpolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create crvserver_appflowpolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("crvserver_appflowpolicy_binding-config")

	tflog.Trace(ctx, "Created crvserver_appflowpolicy_binding resource")

	// Read the updated state back
	r.readCrvserverAppflowpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CrvserverAppflowpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data CrvserverAppflowpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading crvserver_appflowpolicy_binding resource")

	r.readCrvserverAppflowpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CrvserverAppflowpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data CrvserverAppflowpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating crvserver_appflowpolicy_binding resource")

	// Create API request body from the model
	// crvserver_appflowpolicy_binding := crvserver_appflowpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Crvserver_appflowpolicy_binding.Type(), &crvserver_appflowpolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update crvserver_appflowpolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated crvserver_appflowpolicy_binding resource")

	// Read the updated state back
	r.readCrvserverAppflowpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CrvserverAppflowpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data CrvserverAppflowpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting crvserver_appflowpolicy_binding resource")

	// For crvserver_appflowpolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted crvserver_appflowpolicy_binding resource from state")
}

// Helper function to read crvserver_appflowpolicy_binding data from API
func (r *CrvserverAppflowpolicyBindingResource) readCrvserverAppflowpolicyBindingFromApi(ctx context.Context, data *CrvserverAppflowpolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Crvserver_appflowpolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read crvserver_appflowpolicy_binding, got error: %s", err))
		return
	}

	crvserver_appflowpolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
