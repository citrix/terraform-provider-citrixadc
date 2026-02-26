package systemgroup_systemcmdpolicy_binding

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
var _ resource.Resource = &SystemgroupSystemcmdpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*SystemgroupSystemcmdpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*SystemgroupSystemcmdpolicyBindingResource)(nil)

func NewSystemgroupSystemcmdpolicyBindingResource() resource.Resource {
	return &SystemgroupSystemcmdpolicyBindingResource{}
}

// SystemgroupSystemcmdpolicyBindingResource defines the resource implementation.
type SystemgroupSystemcmdpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *SystemgroupSystemcmdpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SystemgroupSystemcmdpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_systemgroup_systemcmdpolicy_binding"
}

func (r *SystemgroupSystemcmdpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SystemgroupSystemcmdpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SystemgroupSystemcmdpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating systemgroup_systemcmdpolicy_binding resource")

	// systemgroup_systemcmdpolicy_binding := systemgroup_systemcmdpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Systemgroup_systemcmdpolicy_binding.Type(), &systemgroup_systemcmdpolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create systemgroup_systemcmdpolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("systemgroup_systemcmdpolicy_binding-config")

	tflog.Trace(ctx, "Created systemgroup_systemcmdpolicy_binding resource")

	// Read the updated state back
	r.readSystemgroupSystemcmdpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemgroupSystemcmdpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SystemgroupSystemcmdpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading systemgroup_systemcmdpolicy_binding resource")

	r.readSystemgroupSystemcmdpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemgroupSystemcmdpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data SystemgroupSystemcmdpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating systemgroup_systemcmdpolicy_binding resource")

	// Create API request body from the model
	// systemgroup_systemcmdpolicy_binding := systemgroup_systemcmdpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Systemgroup_systemcmdpolicy_binding.Type(), &systemgroup_systemcmdpolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update systemgroup_systemcmdpolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated systemgroup_systemcmdpolicy_binding resource")

	// Read the updated state back
	r.readSystemgroupSystemcmdpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemgroupSystemcmdpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SystemgroupSystemcmdpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting systemgroup_systemcmdpolicy_binding resource")

	// For systemgroup_systemcmdpolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted systemgroup_systemcmdpolicy_binding resource from state")
}

// Helper function to read systemgroup_systemcmdpolicy_binding data from API
func (r *SystemgroupSystemcmdpolicyBindingResource) readSystemgroupSystemcmdpolicyBindingFromApi(ctx context.Context, data *SystemgroupSystemcmdpolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Systemgroup_systemcmdpolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read systemgroup_systemcmdpolicy_binding, got error: %s", err))
		return
	}

	systemgroup_systemcmdpolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
