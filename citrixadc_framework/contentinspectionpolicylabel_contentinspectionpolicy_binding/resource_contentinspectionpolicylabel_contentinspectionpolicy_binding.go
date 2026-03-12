package contentinspectionpolicylabel_contentinspectionpolicy_binding

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
var _ resource.Resource = &ContentinspectionpolicylabelContentinspectionpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*ContentinspectionpolicylabelContentinspectionpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*ContentinspectionpolicylabelContentinspectionpolicyBindingResource)(nil)

func NewContentinspectionpolicylabelContentinspectionpolicyBindingResource() resource.Resource {
	return &ContentinspectionpolicylabelContentinspectionpolicyBindingResource{}
}

// ContentinspectionpolicylabelContentinspectionpolicyBindingResource defines the resource implementation.
type ContentinspectionpolicylabelContentinspectionpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *ContentinspectionpolicylabelContentinspectionpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ContentinspectionpolicylabelContentinspectionpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_contentinspectionpolicylabel_contentinspectionpolicy_binding"
}

func (r *ContentinspectionpolicylabelContentinspectionpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ContentinspectionpolicylabelContentinspectionpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ContentinspectionpolicylabelContentinspectionpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating contentinspectionpolicylabel_contentinspectionpolicy_binding resource")

	// contentinspectionpolicylabel_contentinspectionpolicy_binding := contentinspectionpolicylabel_contentinspectionpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Contentinspectionpolicylabel_contentinspectionpolicy_binding.Type(), &contentinspectionpolicylabel_contentinspectionpolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create contentinspectionpolicylabel_contentinspectionpolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("contentinspectionpolicylabel_contentinspectionpolicy_binding-config")

	tflog.Trace(ctx, "Created contentinspectionpolicylabel_contentinspectionpolicy_binding resource")

	// Read the updated state back
	r.readContentinspectionpolicylabelContentinspectionpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ContentinspectionpolicylabelContentinspectionpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ContentinspectionpolicylabelContentinspectionpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading contentinspectionpolicylabel_contentinspectionpolicy_binding resource")

	r.readContentinspectionpolicylabelContentinspectionpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ContentinspectionpolicylabelContentinspectionpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data ContentinspectionpolicylabelContentinspectionpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating contentinspectionpolicylabel_contentinspectionpolicy_binding resource")

	// Create API request body from the model
	// contentinspectionpolicylabel_contentinspectionpolicy_binding := contentinspectionpolicylabel_contentinspectionpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Contentinspectionpolicylabel_contentinspectionpolicy_binding.Type(), &contentinspectionpolicylabel_contentinspectionpolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update contentinspectionpolicylabel_contentinspectionpolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated contentinspectionpolicylabel_contentinspectionpolicy_binding resource")

	// Read the updated state back
	r.readContentinspectionpolicylabelContentinspectionpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ContentinspectionpolicylabelContentinspectionpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ContentinspectionpolicylabelContentinspectionpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting contentinspectionpolicylabel_contentinspectionpolicy_binding resource")

	// For contentinspectionpolicylabel_contentinspectionpolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted contentinspectionpolicylabel_contentinspectionpolicy_binding resource from state")
}

// Helper function to read contentinspectionpolicylabel_contentinspectionpolicy_binding data from API
func (r *ContentinspectionpolicylabelContentinspectionpolicyBindingResource) readContentinspectionpolicylabelContentinspectionpolicyBindingFromApi(ctx context.Context, data *ContentinspectionpolicylabelContentinspectionpolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Contentinspectionpolicylabel_contentinspectionpolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read contentinspectionpolicylabel_contentinspectionpolicy_binding, got error: %s", err))
		return
	}

	contentinspectionpolicylabel_contentinspectionpolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
