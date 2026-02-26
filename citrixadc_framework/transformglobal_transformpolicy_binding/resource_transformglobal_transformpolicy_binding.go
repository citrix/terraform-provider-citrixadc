package transformglobal_transformpolicy_binding

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
var _ resource.Resource = &TransformglobalTransformpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*TransformglobalTransformpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*TransformglobalTransformpolicyBindingResource)(nil)

func NewTransformglobalTransformpolicyBindingResource() resource.Resource {
	return &TransformglobalTransformpolicyBindingResource{}
}

// TransformglobalTransformpolicyBindingResource defines the resource implementation.
type TransformglobalTransformpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *TransformglobalTransformpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *TransformglobalTransformpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_transformglobal_transformpolicy_binding"
}

func (r *TransformglobalTransformpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *TransformglobalTransformpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data TransformglobalTransformpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating transformglobal_transformpolicy_binding resource")

	// transformglobal_transformpolicy_binding := transformglobal_transformpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Transformglobal_transformpolicy_binding.Type(), &transformglobal_transformpolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create transformglobal_transformpolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("transformglobal_transformpolicy_binding-config")

	tflog.Trace(ctx, "Created transformglobal_transformpolicy_binding resource")

	// Read the updated state back
	r.readTransformglobalTransformpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TransformglobalTransformpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data TransformglobalTransformpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading transformglobal_transformpolicy_binding resource")

	r.readTransformglobalTransformpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TransformglobalTransformpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data TransformglobalTransformpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating transformglobal_transformpolicy_binding resource")

	// Create API request body from the model
	// transformglobal_transformpolicy_binding := transformglobal_transformpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Transformglobal_transformpolicy_binding.Type(), &transformglobal_transformpolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update transformglobal_transformpolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated transformglobal_transformpolicy_binding resource")

	// Read the updated state back
	r.readTransformglobalTransformpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TransformglobalTransformpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data TransformglobalTransformpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting transformglobal_transformpolicy_binding resource")

	// For transformglobal_transformpolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted transformglobal_transformpolicy_binding resource from state")
}

// Helper function to read transformglobal_transformpolicy_binding data from API
func (r *TransformglobalTransformpolicyBindingResource) readTransformglobalTransformpolicyBindingFromApi(ctx context.Context, data *TransformglobalTransformpolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Transformglobal_transformpolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read transformglobal_transformpolicy_binding, got error: %s", err))
		return
	}

	transformglobal_transformpolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
