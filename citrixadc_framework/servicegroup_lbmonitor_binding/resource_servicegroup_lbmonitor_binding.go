package servicegroup_lbmonitor_binding

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
var _ resource.Resource = &ServicegroupLbmonitorBindingResource{}
var _ resource.ResourceWithConfigure = (*ServicegroupLbmonitorBindingResource)(nil)
var _ resource.ResourceWithImportState = (*ServicegroupLbmonitorBindingResource)(nil)

func NewServicegroupLbmonitorBindingResource() resource.Resource {
	return &ServicegroupLbmonitorBindingResource{}
}

// ServicegroupLbmonitorBindingResource defines the resource implementation.
type ServicegroupLbmonitorBindingResource struct {
	client *service.NitroClient
}

func (r *ServicegroupLbmonitorBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ServicegroupLbmonitorBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_servicegroup_lbmonitor_binding"
}

func (r *ServicegroupLbmonitorBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ServicegroupLbmonitorBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ServicegroupLbmonitorBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating servicegroup_lbmonitor_binding resource")

	// servicegroup_lbmonitor_binding := servicegroup_lbmonitor_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Servicegroup_lbmonitor_binding.Type(), &servicegroup_lbmonitor_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create servicegroup_lbmonitor_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("servicegroup_lbmonitor_binding-config")

	tflog.Trace(ctx, "Created servicegroup_lbmonitor_binding resource")

	// Read the updated state back
	r.readServicegroupLbmonitorBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ServicegroupLbmonitorBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ServicegroupLbmonitorBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading servicegroup_lbmonitor_binding resource")

	r.readServicegroupLbmonitorBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ServicegroupLbmonitorBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data ServicegroupLbmonitorBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating servicegroup_lbmonitor_binding resource")

	// Create API request body from the model
	// servicegroup_lbmonitor_binding := servicegroup_lbmonitor_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Servicegroup_lbmonitor_binding.Type(), &servicegroup_lbmonitor_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update servicegroup_lbmonitor_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated servicegroup_lbmonitor_binding resource")

	// Read the updated state back
	r.readServicegroupLbmonitorBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ServicegroupLbmonitorBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ServicegroupLbmonitorBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting servicegroup_lbmonitor_binding resource")

	// For servicegroup_lbmonitor_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted servicegroup_lbmonitor_binding resource from state")
}

// Helper function to read servicegroup_lbmonitor_binding data from API
func (r *ServicegroupLbmonitorBindingResource) readServicegroupLbmonitorBindingFromApi(ctx context.Context, data *ServicegroupLbmonitorBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Servicegroup_lbmonitor_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read servicegroup_lbmonitor_binding, got error: %s", err))
		return
	}

	servicegroup_lbmonitor_bindingSetAttrFromGet(ctx, data, getResponseData)

}
