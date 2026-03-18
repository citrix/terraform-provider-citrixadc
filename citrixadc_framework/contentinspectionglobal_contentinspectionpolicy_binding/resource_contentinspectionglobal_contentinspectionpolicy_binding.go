package contentinspectionglobal_contentinspectionpolicy_binding

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
var _ resource.Resource = &ContentinspectionglobalContentinspectionpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*ContentinspectionglobalContentinspectionpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*ContentinspectionglobalContentinspectionpolicyBindingResource)(nil)

func NewContentinspectionglobalContentinspectionpolicyBindingResource() resource.Resource {
	return &ContentinspectionglobalContentinspectionpolicyBindingResource{}
}

// ContentinspectionglobalContentinspectionpolicyBindingResource defines the resource implementation.
type ContentinspectionglobalContentinspectionpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *ContentinspectionglobalContentinspectionpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ContentinspectionglobalContentinspectionpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_contentinspectionglobal_contentinspectionpolicy_binding"
}

func (r *ContentinspectionglobalContentinspectionpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ContentinspectionglobalContentinspectionpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ContentinspectionglobalContentinspectionpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating contentinspectionglobal_contentinspectionpolicy_binding resource")

	// contentinspectionglobal_contentinspectionpolicy_binding := contentinspectionglobal_contentinspectionpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Contentinspectionglobal_contentinspectionpolicy_binding.Type(), &contentinspectionglobal_contentinspectionpolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create contentinspectionglobal_contentinspectionpolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("contentinspectionglobal_contentinspectionpolicy_binding-config")

	tflog.Trace(ctx, "Created contentinspectionglobal_contentinspectionpolicy_binding resource")

	// Read the updated state back
	r.readContentinspectionglobalContentinspectionpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ContentinspectionglobalContentinspectionpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ContentinspectionglobalContentinspectionpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading contentinspectionglobal_contentinspectionpolicy_binding resource")

	r.readContentinspectionglobalContentinspectionpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ContentinspectionglobalContentinspectionpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data ContentinspectionglobalContentinspectionpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating contentinspectionglobal_contentinspectionpolicy_binding resource")

	// Create API request body from the model
	// contentinspectionglobal_contentinspectionpolicy_binding := contentinspectionglobal_contentinspectionpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Contentinspectionglobal_contentinspectionpolicy_binding.Type(), &contentinspectionglobal_contentinspectionpolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update contentinspectionglobal_contentinspectionpolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated contentinspectionglobal_contentinspectionpolicy_binding resource")

	// Read the updated state back
	r.readContentinspectionglobalContentinspectionpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ContentinspectionglobalContentinspectionpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ContentinspectionglobalContentinspectionpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting contentinspectionglobal_contentinspectionpolicy_binding resource")

	// For contentinspectionglobal_contentinspectionpolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted contentinspectionglobal_contentinspectionpolicy_binding resource from state")
}

// Helper function to read contentinspectionglobal_contentinspectionpolicy_binding data from API
func (r *ContentinspectionglobalContentinspectionpolicyBindingResource) readContentinspectionglobalContentinspectionpolicyBindingFromApi(ctx context.Context, data *ContentinspectionglobalContentinspectionpolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Contentinspectionglobal_contentinspectionpolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read contentinspectionglobal_contentinspectionpolicy_binding, got error: %s", err))
		return
	}

	contentinspectionglobal_contentinspectionpolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
