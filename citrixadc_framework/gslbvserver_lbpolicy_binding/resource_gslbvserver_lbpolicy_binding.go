package gslbvserver_lbpolicy_binding

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
var _ resource.Resource = &GslbvserverLbpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*GslbvserverLbpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*GslbvserverLbpolicyBindingResource)(nil)

func NewGslbvserverLbpolicyBindingResource() resource.Resource {
	return &GslbvserverLbpolicyBindingResource{}
}

// GslbvserverLbpolicyBindingResource defines the resource implementation.
type GslbvserverLbpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *GslbvserverLbpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *GslbvserverLbpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gslbvserver_lbpolicy_binding"
}

func (r *GslbvserverLbpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *GslbvserverLbpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data GslbvserverLbpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating gslbvserver_lbpolicy_binding resource")

	// gslbvserver_lbpolicy_binding := gslbvserver_lbpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Gslbvserver_lbpolicy_binding.Type(), &gslbvserver_lbpolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create gslbvserver_lbpolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("gslbvserver_lbpolicy_binding-config")

	tflog.Trace(ctx, "Created gslbvserver_lbpolicy_binding resource")

	// Read the updated state back
	r.readGslbvserverLbpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbvserverLbpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data GslbvserverLbpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading gslbvserver_lbpolicy_binding resource")

	r.readGslbvserverLbpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbvserverLbpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data GslbvserverLbpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating gslbvserver_lbpolicy_binding resource")

	// Create API request body from the model
	// gslbvserver_lbpolicy_binding := gslbvserver_lbpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Gslbvserver_lbpolicy_binding.Type(), &gslbvserver_lbpolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update gslbvserver_lbpolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated gslbvserver_lbpolicy_binding resource")

	// Read the updated state back
	r.readGslbvserverLbpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbvserverLbpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data GslbvserverLbpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting gslbvserver_lbpolicy_binding resource")

	// For gslbvserver_lbpolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted gslbvserver_lbpolicy_binding resource from state")
}

// Helper function to read gslbvserver_lbpolicy_binding data from API
func (r *GslbvserverLbpolicyBindingResource) readGslbvserverLbpolicyBindingFromApi(ctx context.Context, data *GslbvserverLbpolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Gslbvserver_lbpolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read gslbvserver_lbpolicy_binding, got error: %s", err))
		return
	}

	gslbvserver_lbpolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
