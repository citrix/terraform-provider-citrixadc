package cmpglobal_cmppolicy_binding

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
var _ resource.Resource = &CmpglobalCmppolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*CmpglobalCmppolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*CmpglobalCmppolicyBindingResource)(nil)

func NewCmpglobalCmppolicyBindingResource() resource.Resource {
	return &CmpglobalCmppolicyBindingResource{}
}

// CmpglobalCmppolicyBindingResource defines the resource implementation.
type CmpglobalCmppolicyBindingResource struct {
	client *service.NitroClient
}

func (r *CmpglobalCmppolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *CmpglobalCmppolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cmpglobal_cmppolicy_binding"
}

func (r *CmpglobalCmppolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *CmpglobalCmppolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data CmpglobalCmppolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating cmpglobal_cmppolicy_binding resource")

	// cmpglobal_cmppolicy_binding := cmpglobal_cmppolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Cmpglobal_cmppolicy_binding.Type(), &cmpglobal_cmppolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create cmpglobal_cmppolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("cmpglobal_cmppolicy_binding-config")

	tflog.Trace(ctx, "Created cmpglobal_cmppolicy_binding resource")

	// Read the updated state back
	r.readCmpglobalCmppolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CmpglobalCmppolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data CmpglobalCmppolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading cmpglobal_cmppolicy_binding resource")

	r.readCmpglobalCmppolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CmpglobalCmppolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data CmpglobalCmppolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating cmpglobal_cmppolicy_binding resource")

	// Create API request body from the model
	// cmpglobal_cmppolicy_binding := cmpglobal_cmppolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Cmpglobal_cmppolicy_binding.Type(), &cmpglobal_cmppolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update cmpglobal_cmppolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated cmpglobal_cmppolicy_binding resource")

	// Read the updated state back
	r.readCmpglobalCmppolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CmpglobalCmppolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data CmpglobalCmppolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting cmpglobal_cmppolicy_binding resource")

	// For cmpglobal_cmppolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted cmpglobal_cmppolicy_binding resource from state")
}

// Helper function to read cmpglobal_cmppolicy_binding data from API
func (r *CmpglobalCmppolicyBindingResource) readCmpglobalCmppolicyBindingFromApi(ctx context.Context, data *CmpglobalCmppolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Cmpglobal_cmppolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read cmpglobal_cmppolicy_binding, got error: %s", err))
		return
	}

	cmpglobal_cmppolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
