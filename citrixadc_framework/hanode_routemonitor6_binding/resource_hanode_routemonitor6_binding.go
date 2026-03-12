package hanode_routemonitor6_binding

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
var _ resource.Resource = &HanodeRoutemonitor6BindingResource{}
var _ resource.ResourceWithConfigure = (*HanodeRoutemonitor6BindingResource)(nil)
var _ resource.ResourceWithImportState = (*HanodeRoutemonitor6BindingResource)(nil)

func NewHanodeRoutemonitor6BindingResource() resource.Resource {
	return &HanodeRoutemonitor6BindingResource{}
}

// HanodeRoutemonitor6BindingResource defines the resource implementation.
type HanodeRoutemonitor6BindingResource struct {
	client *service.NitroClient
}

func (r *HanodeRoutemonitor6BindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *HanodeRoutemonitor6BindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_hanode_routemonitor6_binding"
}

func (r *HanodeRoutemonitor6BindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *HanodeRoutemonitor6BindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data HanodeRoutemonitor6BindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating hanode_routemonitor6_binding resource")

	// hanode_routemonitor6_binding := hanode_routemonitor6_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Hanode_routemonitor6_binding.Type(), &hanode_routemonitor6_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create hanode_routemonitor6_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("hanode_routemonitor6_binding-config")

	tflog.Trace(ctx, "Created hanode_routemonitor6_binding resource")

	// Read the updated state back
	r.readHanodeRoutemonitor6BindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *HanodeRoutemonitor6BindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data HanodeRoutemonitor6BindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading hanode_routemonitor6_binding resource")

	r.readHanodeRoutemonitor6BindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *HanodeRoutemonitor6BindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data HanodeRoutemonitor6BindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating hanode_routemonitor6_binding resource")

	// Create API request body from the model
	// hanode_routemonitor6_binding := hanode_routemonitor6_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Hanode_routemonitor6_binding.Type(), &hanode_routemonitor6_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update hanode_routemonitor6_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated hanode_routemonitor6_binding resource")

	// Read the updated state back
	r.readHanodeRoutemonitor6BindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *HanodeRoutemonitor6BindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data HanodeRoutemonitor6BindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting hanode_routemonitor6_binding resource")

	// For hanode_routemonitor6_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted hanode_routemonitor6_binding resource from state")
}

// Helper function to read hanode_routemonitor6_binding data from API
func (r *HanodeRoutemonitor6BindingResource) readHanodeRoutemonitor6BindingFromApi(ctx context.Context, data *HanodeRoutemonitor6BindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Hanode_routemonitor6_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read hanode_routemonitor6_binding, got error: %s", err))
		return
	}

	hanode_routemonitor6_bindingSetAttrFromGet(ctx, data, getResponseData)

}
