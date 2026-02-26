package crvserver_responderpolicy_binding

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
var _ resource.Resource = &CrvserverResponderpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*CrvserverResponderpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*CrvserverResponderpolicyBindingResource)(nil)

func NewCrvserverResponderpolicyBindingResource() resource.Resource {
	return &CrvserverResponderpolicyBindingResource{}
}

// CrvserverResponderpolicyBindingResource defines the resource implementation.
type CrvserverResponderpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *CrvserverResponderpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *CrvserverResponderpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_crvserver_responderpolicy_binding"
}

func (r *CrvserverResponderpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *CrvserverResponderpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data CrvserverResponderpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating crvserver_responderpolicy_binding resource")

	// crvserver_responderpolicy_binding := crvserver_responderpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Crvserver_responderpolicy_binding.Type(), &crvserver_responderpolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create crvserver_responderpolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("crvserver_responderpolicy_binding-config")

	tflog.Trace(ctx, "Created crvserver_responderpolicy_binding resource")

	// Read the updated state back
	r.readCrvserverResponderpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CrvserverResponderpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data CrvserverResponderpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading crvserver_responderpolicy_binding resource")

	r.readCrvserverResponderpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CrvserverResponderpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data CrvserverResponderpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating crvserver_responderpolicy_binding resource")

	// Create API request body from the model
	// crvserver_responderpolicy_binding := crvserver_responderpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Crvserver_responderpolicy_binding.Type(), &crvserver_responderpolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update crvserver_responderpolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated crvserver_responderpolicy_binding resource")

	// Read the updated state back
	r.readCrvserverResponderpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CrvserverResponderpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data CrvserverResponderpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting crvserver_responderpolicy_binding resource")

	// For crvserver_responderpolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted crvserver_responderpolicy_binding resource from state")
}

// Helper function to read crvserver_responderpolicy_binding data from API
func (r *CrvserverResponderpolicyBindingResource) readCrvserverResponderpolicyBindingFromApi(ctx context.Context, data *CrvserverResponderpolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Crvserver_responderpolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read crvserver_responderpolicy_binding, got error: %s", err))
		return
	}

	crvserver_responderpolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
