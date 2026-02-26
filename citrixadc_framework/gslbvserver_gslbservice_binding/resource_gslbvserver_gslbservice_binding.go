package gslbvserver_gslbservice_binding

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
var _ resource.Resource = &GslbvserverGslbserviceBindingResource{}
var _ resource.ResourceWithConfigure = (*GslbvserverGslbserviceBindingResource)(nil)
var _ resource.ResourceWithImportState = (*GslbvserverGslbserviceBindingResource)(nil)

func NewGslbvserverGslbserviceBindingResource() resource.Resource {
	return &GslbvserverGslbserviceBindingResource{}
}

// GslbvserverGslbserviceBindingResource defines the resource implementation.
type GslbvserverGslbserviceBindingResource struct {
	client *service.NitroClient
}

func (r *GslbvserverGslbserviceBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *GslbvserverGslbserviceBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gslbvserver_gslbservice_binding"
}

func (r *GslbvserverGslbserviceBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *GslbvserverGslbserviceBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data GslbvserverGslbserviceBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating gslbvserver_gslbservice_binding resource")

	// gslbvserver_gslbservice_binding := gslbvserver_gslbservice_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Gslbvserver_gslbservice_binding.Type(), &gslbvserver_gslbservice_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create gslbvserver_gslbservice_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("gslbvserver_gslbservice_binding-config")

	tflog.Trace(ctx, "Created gslbvserver_gslbservice_binding resource")

	// Read the updated state back
	r.readGslbvserverGslbserviceBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbvserverGslbserviceBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data GslbvserverGslbserviceBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading gslbvserver_gslbservice_binding resource")

	r.readGslbvserverGslbserviceBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbvserverGslbserviceBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data GslbvserverGslbserviceBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating gslbvserver_gslbservice_binding resource")

	// Create API request body from the model
	// gslbvserver_gslbservice_binding := gslbvserver_gslbservice_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Gslbvserver_gslbservice_binding.Type(), &gslbvserver_gslbservice_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update gslbvserver_gslbservice_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated gslbvserver_gslbservice_binding resource")

	// Read the updated state back
	r.readGslbvserverGslbserviceBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbvserverGslbserviceBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data GslbvserverGslbserviceBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting gslbvserver_gslbservice_binding resource")

	// For gslbvserver_gslbservice_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted gslbvserver_gslbservice_binding resource from state")
}

// Helper function to read gslbvserver_gslbservice_binding data from API
func (r *GslbvserverGslbserviceBindingResource) readGslbvserverGslbserviceBindingFromApi(ctx context.Context, data *GslbvserverGslbserviceBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Gslbvserver_gslbservice_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read gslbvserver_gslbservice_binding, got error: %s", err))
		return
	}

	gslbvserver_gslbservice_bindingSetAttrFromGet(ctx, data, getResponseData)

}
