package responderglobal_responderpolicy_binding

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
var _ resource.Resource = &ResponderglobalResponderpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*ResponderglobalResponderpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*ResponderglobalResponderpolicyBindingResource)(nil)

func NewResponderglobalResponderpolicyBindingResource() resource.Resource {
	return &ResponderglobalResponderpolicyBindingResource{}
}

// ResponderglobalResponderpolicyBindingResource defines the resource implementation.
type ResponderglobalResponderpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *ResponderglobalResponderpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ResponderglobalResponderpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_responderglobal_responderpolicy_binding"
}

func (r *ResponderglobalResponderpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ResponderglobalResponderpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ResponderglobalResponderpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating responderglobal_responderpolicy_binding resource")

	// responderglobal_responderpolicy_binding := responderglobal_responderpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Responderglobal_responderpolicy_binding.Type(), &responderglobal_responderpolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create responderglobal_responderpolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("responderglobal_responderpolicy_binding-config")

	tflog.Trace(ctx, "Created responderglobal_responderpolicy_binding resource")

	// Read the updated state back
	r.readResponderglobalResponderpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ResponderglobalResponderpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ResponderglobalResponderpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading responderglobal_responderpolicy_binding resource")

	r.readResponderglobalResponderpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ResponderglobalResponderpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data ResponderglobalResponderpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating responderglobal_responderpolicy_binding resource")

	// Create API request body from the model
	// responderglobal_responderpolicy_binding := responderglobal_responderpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Responderglobal_responderpolicy_binding.Type(), &responderglobal_responderpolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update responderglobal_responderpolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated responderglobal_responderpolicy_binding resource")

	// Read the updated state back
	r.readResponderglobalResponderpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ResponderglobalResponderpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ResponderglobalResponderpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting responderglobal_responderpolicy_binding resource")

	// For responderglobal_responderpolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted responderglobal_responderpolicy_binding resource from state")
}

// Helper function to read responderglobal_responderpolicy_binding data from API
func (r *ResponderglobalResponderpolicyBindingResource) readResponderglobalResponderpolicyBindingFromApi(ctx context.Context, data *ResponderglobalResponderpolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Responderglobal_responderpolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read responderglobal_responderpolicy_binding, got error: %s", err))
		return
	}

	responderglobal_responderpolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
