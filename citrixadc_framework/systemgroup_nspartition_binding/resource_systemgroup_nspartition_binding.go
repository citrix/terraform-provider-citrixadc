package systemgroup_nspartition_binding

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
var _ resource.Resource = &SystemgroupNspartitionBindingResource{}
var _ resource.ResourceWithConfigure = (*SystemgroupNspartitionBindingResource)(nil)
var _ resource.ResourceWithImportState = (*SystemgroupNspartitionBindingResource)(nil)

func NewSystemgroupNspartitionBindingResource() resource.Resource {
	return &SystemgroupNspartitionBindingResource{}
}

// SystemgroupNspartitionBindingResource defines the resource implementation.
type SystemgroupNspartitionBindingResource struct {
	client *service.NitroClient
}

func (r *SystemgroupNspartitionBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SystemgroupNspartitionBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_systemgroup_nspartition_binding"
}

func (r *SystemgroupNspartitionBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SystemgroupNspartitionBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SystemgroupNspartitionBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating systemgroup_nspartition_binding resource")

	// systemgroup_nspartition_binding := systemgroup_nspartition_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Systemgroup_nspartition_binding.Type(), &systemgroup_nspartition_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create systemgroup_nspartition_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("systemgroup_nspartition_binding-config")

	tflog.Trace(ctx, "Created systemgroup_nspartition_binding resource")

	// Read the updated state back
	r.readSystemgroupNspartitionBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemgroupNspartitionBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SystemgroupNspartitionBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading systemgroup_nspartition_binding resource")

	r.readSystemgroupNspartitionBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemgroupNspartitionBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data SystemgroupNspartitionBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating systemgroup_nspartition_binding resource")

	// Create API request body from the model
	// systemgroup_nspartition_binding := systemgroup_nspartition_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Systemgroup_nspartition_binding.Type(), &systemgroup_nspartition_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update systemgroup_nspartition_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated systemgroup_nspartition_binding resource")

	// Read the updated state back
	r.readSystemgroupNspartitionBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemgroupNspartitionBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SystemgroupNspartitionBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting systemgroup_nspartition_binding resource")

	// For systemgroup_nspartition_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted systemgroup_nspartition_binding resource from state")
}

// Helper function to read systemgroup_nspartition_binding data from API
func (r *SystemgroupNspartitionBindingResource) readSystemgroupNspartitionBindingFromApi(ctx context.Context, data *SystemgroupNspartitionBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Systemgroup_nspartition_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read systemgroup_nspartition_binding, got error: %s", err))
		return
	}

	systemgroup_nspartition_bindingSetAttrFromGet(ctx, data, getResponseData)

}
