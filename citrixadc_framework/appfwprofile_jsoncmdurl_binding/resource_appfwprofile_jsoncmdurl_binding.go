package appfwprofile_jsoncmdurl_binding

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
var _ resource.Resource = &AppfwprofileJsoncmdurlBindingResource{}
var _ resource.ResourceWithConfigure = (*AppfwprofileJsoncmdurlBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwprofileJsoncmdurlBindingResource)(nil)

func NewAppfwprofileJsoncmdurlBindingResource() resource.Resource {
	return &AppfwprofileJsoncmdurlBindingResource{}
}

// AppfwprofileJsoncmdurlBindingResource defines the resource implementation.
type AppfwprofileJsoncmdurlBindingResource struct {
	client *service.NitroClient
}

func (r *AppfwprofileJsoncmdurlBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwprofileJsoncmdurlBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_jsoncmdurl_binding"
}

func (r *AppfwprofileJsoncmdurlBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwprofileJsoncmdurlBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwprofileJsoncmdurlBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwprofile_jsoncmdurl_binding resource")

	// appfwprofile_jsoncmdurl_binding := appfwprofile_jsoncmdurl_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwprofile_jsoncmdurl_binding.Type(), &appfwprofile_jsoncmdurl_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwprofile_jsoncmdurl_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("appfwprofile_jsoncmdurl_binding-config")

	tflog.Trace(ctx, "Created appfwprofile_jsoncmdurl_binding resource")

	// Read the updated state back
	r.readAppfwprofileJsoncmdurlBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileJsoncmdurlBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwprofileJsoncmdurlBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwprofile_jsoncmdurl_binding resource")

	r.readAppfwprofileJsoncmdurlBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileJsoncmdurlBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AppfwprofileJsoncmdurlBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating appfwprofile_jsoncmdurl_binding resource")

	// Create API request body from the model
	// appfwprofile_jsoncmdurl_binding := appfwprofile_jsoncmdurl_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwprofile_jsoncmdurl_binding.Type(), &appfwprofile_jsoncmdurl_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appfwprofile_jsoncmdurl_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated appfwprofile_jsoncmdurl_binding resource")

	// Read the updated state back
	r.readAppfwprofileJsoncmdurlBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileJsoncmdurlBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwprofileJsoncmdurlBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwprofile_jsoncmdurl_binding resource")

	// For appfwprofile_jsoncmdurl_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted appfwprofile_jsoncmdurl_binding resource from state")
}

// Helper function to read appfwprofile_jsoncmdurl_binding data from API
func (r *AppfwprofileJsoncmdurlBindingResource) readAppfwprofileJsoncmdurlBindingFromApi(ctx context.Context, data *AppfwprofileJsoncmdurlBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Appfwprofile_jsoncmdurl_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_jsoncmdurl_binding, got error: %s", err))
		return
	}

	appfwprofile_jsoncmdurl_bindingSetAttrFromGet(ctx, data, getResponseData)

}
