package hanode_routemonitor_binding

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
var _ resource.Resource = &HanodeRoutemonitorBindingResource{}
var _ resource.ResourceWithConfigure = (*HanodeRoutemonitorBindingResource)(nil)
var _ resource.ResourceWithImportState = (*HanodeRoutemonitorBindingResource)(nil)

func NewHanodeRoutemonitorBindingResource() resource.Resource {
	return &HanodeRoutemonitorBindingResource{}
}

// HanodeRoutemonitorBindingResource defines the resource implementation.
type HanodeRoutemonitorBindingResource struct {
	client *service.NitroClient
}

func (r *HanodeRoutemonitorBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *HanodeRoutemonitorBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_hanode_routemonitor_binding"
}

func (r *HanodeRoutemonitorBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *HanodeRoutemonitorBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data HanodeRoutemonitorBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating hanode_routemonitor_binding resource")

	// hanode_routemonitor_binding := hanode_routemonitor_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Hanode_routemonitor_binding.Type(), &hanode_routemonitor_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create hanode_routemonitor_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("hanode_routemonitor_binding-config")

	tflog.Trace(ctx, "Created hanode_routemonitor_binding resource")

	// Read the updated state back
	r.readHanodeRoutemonitorBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *HanodeRoutemonitorBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data HanodeRoutemonitorBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading hanode_routemonitor_binding resource")

	r.readHanodeRoutemonitorBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *HanodeRoutemonitorBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data HanodeRoutemonitorBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating hanode_routemonitor_binding resource")

	// Create API request body from the model
	// hanode_routemonitor_binding := hanode_routemonitor_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Hanode_routemonitor_binding.Type(), &hanode_routemonitor_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update hanode_routemonitor_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated hanode_routemonitor_binding resource")

	// Read the updated state back
	r.readHanodeRoutemonitorBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *HanodeRoutemonitorBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data HanodeRoutemonitorBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting hanode_routemonitor_binding resource")

	// For hanode_routemonitor_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted hanode_routemonitor_binding resource from state")
}

// Helper function to read hanode_routemonitor_binding data from API
func (r *HanodeRoutemonitorBindingResource) readHanodeRoutemonitorBindingFromApi(ctx context.Context, data *HanodeRoutemonitorBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Hanode_routemonitor_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read hanode_routemonitor_binding, got error: %s", err))
		return
	}

	hanode_routemonitor_bindingSetAttrFromGet(ctx, data, getResponseData)

}
