package servicegroup_servicegroupmember_binding

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
var _ resource.Resource = &ServicegroupServicegroupmemberBindingResource{}
var _ resource.ResourceWithConfigure = (*ServicegroupServicegroupmemberBindingResource)(nil)
var _ resource.ResourceWithImportState = (*ServicegroupServicegroupmemberBindingResource)(nil)

func NewServicegroupServicegroupmemberBindingResource() resource.Resource {
	return &ServicegroupServicegroupmemberBindingResource{}
}

// ServicegroupServicegroupmemberBindingResource defines the resource implementation.
type ServicegroupServicegroupmemberBindingResource struct {
	client *service.NitroClient
}

func (r *ServicegroupServicegroupmemberBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ServicegroupServicegroupmemberBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_servicegroup_servicegroupmember_binding"
}

func (r *ServicegroupServicegroupmemberBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ServicegroupServicegroupmemberBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ServicegroupServicegroupmemberBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating servicegroup_servicegroupmember_binding resource")

	// servicegroup_servicegroupmember_binding := servicegroup_servicegroupmember_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Servicegroup_servicegroupmember_binding.Type(), &servicegroup_servicegroupmember_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create servicegroup_servicegroupmember_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("servicegroup_servicegroupmember_binding-config")

	tflog.Trace(ctx, "Created servicegroup_servicegroupmember_binding resource")

	// Read the updated state back
	r.readServicegroupServicegroupmemberBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ServicegroupServicegroupmemberBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ServicegroupServicegroupmemberBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading servicegroup_servicegroupmember_binding resource")

	r.readServicegroupServicegroupmemberBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ServicegroupServicegroupmemberBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data ServicegroupServicegroupmemberBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating servicegroup_servicegroupmember_binding resource")

	// Create API request body from the model
	// servicegroup_servicegroupmember_binding := servicegroup_servicegroupmember_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Servicegroup_servicegroupmember_binding.Type(), &servicegroup_servicegroupmember_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update servicegroup_servicegroupmember_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated servicegroup_servicegroupmember_binding resource")

	// Read the updated state back
	r.readServicegroupServicegroupmemberBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ServicegroupServicegroupmemberBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ServicegroupServicegroupmemberBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting servicegroup_servicegroupmember_binding resource")

	// For servicegroup_servicegroupmember_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted servicegroup_servicegroupmember_binding resource from state")
}

// Helper function to read servicegroup_servicegroupmember_binding data from API
func (r *ServicegroupServicegroupmemberBindingResource) readServicegroupServicegroupmemberBindingFromApi(ctx context.Context, data *ServicegroupServicegroupmemberBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Servicegroup_servicegroupmember_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read servicegroup_servicegroupmember_binding, got error: %s", err))
		return
	}

	servicegroup_servicegroupmember_bindingSetAttrFromGet(ctx, data, getResponseData)

}
