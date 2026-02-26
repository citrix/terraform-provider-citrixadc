package transformpolicylabel_transformpolicy_binding

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
var _ resource.Resource = &TransformpolicylabelTransformpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*TransformpolicylabelTransformpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*TransformpolicylabelTransformpolicyBindingResource)(nil)

func NewTransformpolicylabelTransformpolicyBindingResource() resource.Resource {
	return &TransformpolicylabelTransformpolicyBindingResource{}
}

// TransformpolicylabelTransformpolicyBindingResource defines the resource implementation.
type TransformpolicylabelTransformpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *TransformpolicylabelTransformpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *TransformpolicylabelTransformpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_transformpolicylabel_transformpolicy_binding"
}

func (r *TransformpolicylabelTransformpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *TransformpolicylabelTransformpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data TransformpolicylabelTransformpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating transformpolicylabel_transformpolicy_binding resource")

	// transformpolicylabel_transformpolicy_binding := transformpolicylabel_transformpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Transformpolicylabel_transformpolicy_binding.Type(), &transformpolicylabel_transformpolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create transformpolicylabel_transformpolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("transformpolicylabel_transformpolicy_binding-config")

	tflog.Trace(ctx, "Created transformpolicylabel_transformpolicy_binding resource")

	// Read the updated state back
	r.readTransformpolicylabelTransformpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TransformpolicylabelTransformpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data TransformpolicylabelTransformpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading transformpolicylabel_transformpolicy_binding resource")

	r.readTransformpolicylabelTransformpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TransformpolicylabelTransformpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data TransformpolicylabelTransformpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating transformpolicylabel_transformpolicy_binding resource")

	// Create API request body from the model
	// transformpolicylabel_transformpolicy_binding := transformpolicylabel_transformpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Transformpolicylabel_transformpolicy_binding.Type(), &transformpolicylabel_transformpolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update transformpolicylabel_transformpolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated transformpolicylabel_transformpolicy_binding resource")

	// Read the updated state back
	r.readTransformpolicylabelTransformpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TransformpolicylabelTransformpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data TransformpolicylabelTransformpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting transformpolicylabel_transformpolicy_binding resource")

	// For transformpolicylabel_transformpolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted transformpolicylabel_transformpolicy_binding resource from state")
}

// Helper function to read transformpolicylabel_transformpolicy_binding data from API
func (r *TransformpolicylabelTransformpolicyBindingResource) readTransformpolicylabelTransformpolicyBindingFromApi(ctx context.Context, data *TransformpolicylabelTransformpolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Transformpolicylabel_transformpolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read transformpolicylabel_transformpolicy_binding, got error: %s", err))
		return
	}

	transformpolicylabel_transformpolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
