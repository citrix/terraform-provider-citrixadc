package crvserver_feopolicy_binding

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
var _ resource.Resource = &CrvserverFeopolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*CrvserverFeopolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*CrvserverFeopolicyBindingResource)(nil)

func NewCrvserverFeopolicyBindingResource() resource.Resource {
	return &CrvserverFeopolicyBindingResource{}
}

// CrvserverFeopolicyBindingResource defines the resource implementation.
type CrvserverFeopolicyBindingResource struct {
	client *service.NitroClient
}

func (r *CrvserverFeopolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *CrvserverFeopolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_crvserver_feopolicy_binding"
}

func (r *CrvserverFeopolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *CrvserverFeopolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data CrvserverFeopolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating crvserver_feopolicy_binding resource")

	// crvserver_feopolicy_binding := crvserver_feopolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Crvserver_feopolicy_binding.Type(), &crvserver_feopolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create crvserver_feopolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("crvserver_feopolicy_binding-config")

	tflog.Trace(ctx, "Created crvserver_feopolicy_binding resource")

	// Read the updated state back
	r.readCrvserverFeopolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CrvserverFeopolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data CrvserverFeopolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading crvserver_feopolicy_binding resource")

	r.readCrvserverFeopolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CrvserverFeopolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data CrvserverFeopolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating crvserver_feopolicy_binding resource")

	// Create API request body from the model
	// crvserver_feopolicy_binding := crvserver_feopolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Crvserver_feopolicy_binding.Type(), &crvserver_feopolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update crvserver_feopolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated crvserver_feopolicy_binding resource")

	// Read the updated state back
	r.readCrvserverFeopolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CrvserverFeopolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data CrvserverFeopolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting crvserver_feopolicy_binding resource")

	// For crvserver_feopolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted crvserver_feopolicy_binding resource from state")
}

// Helper function to read crvserver_feopolicy_binding data from API
func (r *CrvserverFeopolicyBindingResource) readCrvserverFeopolicyBindingFromApi(ctx context.Context, data *CrvserverFeopolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Crvserver_feopolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read crvserver_feopolicy_binding, got error: %s", err))
		return
	}

	crvserver_feopolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
