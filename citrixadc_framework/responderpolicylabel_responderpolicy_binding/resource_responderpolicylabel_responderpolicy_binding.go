package responderpolicylabel_responderpolicy_binding

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
var _ resource.Resource = &ResponderpolicylabelResponderpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*ResponderpolicylabelResponderpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*ResponderpolicylabelResponderpolicyBindingResource)(nil)

func NewResponderpolicylabelResponderpolicyBindingResource() resource.Resource {
	return &ResponderpolicylabelResponderpolicyBindingResource{}
}

// ResponderpolicylabelResponderpolicyBindingResource defines the resource implementation.
type ResponderpolicylabelResponderpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *ResponderpolicylabelResponderpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ResponderpolicylabelResponderpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_responderpolicylabel_responderpolicy_binding"
}

func (r *ResponderpolicylabelResponderpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ResponderpolicylabelResponderpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ResponderpolicylabelResponderpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating responderpolicylabel_responderpolicy_binding resource")

	// responderpolicylabel_responderpolicy_binding := responderpolicylabel_responderpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Responderpolicylabel_responderpolicy_binding.Type(), &responderpolicylabel_responderpolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create responderpolicylabel_responderpolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("responderpolicylabel_responderpolicy_binding-config")

	tflog.Trace(ctx, "Created responderpolicylabel_responderpolicy_binding resource")

	// Read the updated state back
	r.readResponderpolicylabelResponderpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ResponderpolicylabelResponderpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ResponderpolicylabelResponderpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading responderpolicylabel_responderpolicy_binding resource")

	r.readResponderpolicylabelResponderpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ResponderpolicylabelResponderpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data ResponderpolicylabelResponderpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating responderpolicylabel_responderpolicy_binding resource")

	// Create API request body from the model
	// responderpolicylabel_responderpolicy_binding := responderpolicylabel_responderpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Responderpolicylabel_responderpolicy_binding.Type(), &responderpolicylabel_responderpolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update responderpolicylabel_responderpolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated responderpolicylabel_responderpolicy_binding resource")

	// Read the updated state back
	r.readResponderpolicylabelResponderpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ResponderpolicylabelResponderpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ResponderpolicylabelResponderpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting responderpolicylabel_responderpolicy_binding resource")

	// For responderpolicylabel_responderpolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted responderpolicylabel_responderpolicy_binding resource from state")
}

// Helper function to read responderpolicylabel_responderpolicy_binding data from API
func (r *ResponderpolicylabelResponderpolicyBindingResource) readResponderpolicylabelResponderpolicyBindingFromApi(ctx context.Context, data *ResponderpolicylabelResponderpolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Responderpolicylabel_responderpolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read responderpolicylabel_responderpolicy_binding, got error: %s", err))
		return
	}

	responderpolicylabel_responderpolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
