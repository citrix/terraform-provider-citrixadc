package vpnglobal_appcontroller_binding

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
var _ resource.Resource = &VpnglobalAppcontrollerBindingResource{}
var _ resource.ResourceWithConfigure = (*VpnglobalAppcontrollerBindingResource)(nil)
var _ resource.ResourceWithImportState = (*VpnglobalAppcontrollerBindingResource)(nil)

func NewVpnglobalAppcontrollerBindingResource() resource.Resource {
	return &VpnglobalAppcontrollerBindingResource{}
}

// VpnglobalAppcontrollerBindingResource defines the resource implementation.
type VpnglobalAppcontrollerBindingResource struct {
	client *service.NitroClient
}

func (r *VpnglobalAppcontrollerBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnglobalAppcontrollerBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnglobal_appcontroller_binding"
}

func (r *VpnglobalAppcontrollerBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnglobalAppcontrollerBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnglobalAppcontrollerBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnglobal_appcontroller_binding resource")

	// vpnglobal_appcontroller_binding := vpnglobal_appcontroller_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnglobal_appcontroller_binding.Type(), &vpnglobal_appcontroller_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnglobal_appcontroller_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("vpnglobal_appcontroller_binding-config")

	tflog.Trace(ctx, "Created vpnglobal_appcontroller_binding resource")

	// Read the updated state back
	r.readVpnglobalAppcontrollerBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalAppcontrollerBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnglobalAppcontrollerBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnglobal_appcontroller_binding resource")

	r.readVpnglobalAppcontrollerBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalAppcontrollerBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data VpnglobalAppcontrollerBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating vpnglobal_appcontroller_binding resource")

	// Create API request body from the model
	// vpnglobal_appcontroller_binding := vpnglobal_appcontroller_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnglobal_appcontroller_binding.Type(), &vpnglobal_appcontroller_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnglobal_appcontroller_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated vpnglobal_appcontroller_binding resource")

	// Read the updated state back
	r.readVpnglobalAppcontrollerBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalAppcontrollerBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnglobalAppcontrollerBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnglobal_appcontroller_binding resource")

	// For vpnglobal_appcontroller_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted vpnglobal_appcontroller_binding resource from state")
}

// Helper function to read vpnglobal_appcontroller_binding data from API
func (r *VpnglobalAppcontrollerBindingResource) readVpnglobalAppcontrollerBindingFromApi(ctx context.Context, data *VpnglobalAppcontrollerBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Vpnglobal_appcontroller_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnglobal_appcontroller_binding, got error: %s", err))
		return
	}

	vpnglobal_appcontroller_bindingSetAttrFromGet(ctx, data, getResponseData)

}
