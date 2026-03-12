package systemuser_nspartition_binding

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
var _ resource.Resource = &SystemuserNspartitionBindingResource{}
var _ resource.ResourceWithConfigure = (*SystemuserNspartitionBindingResource)(nil)
var _ resource.ResourceWithImportState = (*SystemuserNspartitionBindingResource)(nil)

func NewSystemuserNspartitionBindingResource() resource.Resource {
	return &SystemuserNspartitionBindingResource{}
}

// SystemuserNspartitionBindingResource defines the resource implementation.
type SystemuserNspartitionBindingResource struct {
	client *service.NitroClient
}

func (r *SystemuserNspartitionBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SystemuserNspartitionBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_systemuser_nspartition_binding"
}

func (r *SystemuserNspartitionBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SystemuserNspartitionBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SystemuserNspartitionBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating systemuser_nspartition_binding resource")

	// systemuser_nspartition_binding := systemuser_nspartition_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Systemuser_nspartition_binding.Type(), &systemuser_nspartition_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create systemuser_nspartition_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("systemuser_nspartition_binding-config")

	tflog.Trace(ctx, "Created systemuser_nspartition_binding resource")

	// Read the updated state back
	r.readSystemuserNspartitionBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemuserNspartitionBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SystemuserNspartitionBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading systemuser_nspartition_binding resource")

	r.readSystemuserNspartitionBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemuserNspartitionBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data SystemuserNspartitionBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating systemuser_nspartition_binding resource")

	// Create API request body from the model
	// systemuser_nspartition_binding := systemuser_nspartition_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Systemuser_nspartition_binding.Type(), &systemuser_nspartition_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update systemuser_nspartition_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated systemuser_nspartition_binding resource")

	// Read the updated state back
	r.readSystemuserNspartitionBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemuserNspartitionBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SystemuserNspartitionBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting systemuser_nspartition_binding resource")

	// For systemuser_nspartition_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted systemuser_nspartition_binding resource from state")
}

// Helper function to read systemuser_nspartition_binding data from API
func (r *SystemuserNspartitionBindingResource) readSystemuserNspartitionBindingFromApi(ctx context.Context, data *SystemuserNspartitionBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Systemuser_nspartition_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read systemuser_nspartition_binding, got error: %s", err))
		return
	}

	systemuser_nspartition_bindingSetAttrFromGet(ctx, data, getResponseData)

}
