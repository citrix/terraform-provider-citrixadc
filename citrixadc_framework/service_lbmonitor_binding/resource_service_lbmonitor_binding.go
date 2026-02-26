package service_lbmonitor_binding

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
var _ resource.Resource = &ServiceLbmonitorBindingResource{}
var _ resource.ResourceWithConfigure = (*ServiceLbmonitorBindingResource)(nil)
var _ resource.ResourceWithImportState = (*ServiceLbmonitorBindingResource)(nil)

func NewServiceLbmonitorBindingResource() resource.Resource {
	return &ServiceLbmonitorBindingResource{}
}

// ServiceLbmonitorBindingResource defines the resource implementation.
type ServiceLbmonitorBindingResource struct {
	client *service.NitroClient
}

func (r *ServiceLbmonitorBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ServiceLbmonitorBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_lbmonitor_binding"
}

func (r *ServiceLbmonitorBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ServiceLbmonitorBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ServiceLbmonitorBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating service_lbmonitor_binding resource")

	// service_lbmonitor_binding := service_lbmonitor_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Service_lbmonitor_binding.Type(), &service_lbmonitor_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create service_lbmonitor_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("service_lbmonitor_binding-config")

	tflog.Trace(ctx, "Created service_lbmonitor_binding resource")

	// Read the updated state back
	r.readServiceLbmonitorBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ServiceLbmonitorBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ServiceLbmonitorBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading service_lbmonitor_binding resource")

	r.readServiceLbmonitorBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ServiceLbmonitorBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data ServiceLbmonitorBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating service_lbmonitor_binding resource")

	// Create API request body from the model
	// service_lbmonitor_binding := service_lbmonitor_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Service_lbmonitor_binding.Type(), &service_lbmonitor_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update service_lbmonitor_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated service_lbmonitor_binding resource")

	// Read the updated state back
	r.readServiceLbmonitorBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ServiceLbmonitorBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ServiceLbmonitorBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting service_lbmonitor_binding resource")

	// For service_lbmonitor_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted service_lbmonitor_binding resource from state")
}

// Helper function to read service_lbmonitor_binding data from API
func (r *ServiceLbmonitorBindingResource) readServiceLbmonitorBindingFromApi(ctx context.Context, data *ServiceLbmonitorBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Service_lbmonitor_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read service_lbmonitor_binding, got error: %s", err))
		return
	}

	service_lbmonitor_bindingSetAttrFromGet(ctx, data, getResponseData)

}
