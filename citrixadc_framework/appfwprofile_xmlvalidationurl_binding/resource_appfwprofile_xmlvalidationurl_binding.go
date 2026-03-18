package appfwprofile_xmlvalidationurl_binding

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
var _ resource.Resource = &AppfwprofileXmlvalidationurlBindingResource{}
var _ resource.ResourceWithConfigure = (*AppfwprofileXmlvalidationurlBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwprofileXmlvalidationurlBindingResource)(nil)

func NewAppfwprofileXmlvalidationurlBindingResource() resource.Resource {
	return &AppfwprofileXmlvalidationurlBindingResource{}
}

// AppfwprofileXmlvalidationurlBindingResource defines the resource implementation.
type AppfwprofileXmlvalidationurlBindingResource struct {
	client *service.NitroClient
}

func (r *AppfwprofileXmlvalidationurlBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwprofileXmlvalidationurlBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_xmlvalidationurl_binding"
}

func (r *AppfwprofileXmlvalidationurlBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwprofileXmlvalidationurlBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwprofileXmlvalidationurlBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwprofile_xmlvalidationurl_binding resource")

	// appfwprofile_xmlvalidationurl_binding := appfwprofile_xmlvalidationurl_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwprofile_xmlvalidationurl_binding.Type(), &appfwprofile_xmlvalidationurl_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwprofile_xmlvalidationurl_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("appfwprofile_xmlvalidationurl_binding-config")

	tflog.Trace(ctx, "Created appfwprofile_xmlvalidationurl_binding resource")

	// Read the updated state back
	r.readAppfwprofileXmlvalidationurlBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileXmlvalidationurlBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwprofileXmlvalidationurlBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwprofile_xmlvalidationurl_binding resource")

	r.readAppfwprofileXmlvalidationurlBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileXmlvalidationurlBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AppfwprofileXmlvalidationurlBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating appfwprofile_xmlvalidationurl_binding resource")

	// Create API request body from the model
	// appfwprofile_xmlvalidationurl_binding := appfwprofile_xmlvalidationurl_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwprofile_xmlvalidationurl_binding.Type(), &appfwprofile_xmlvalidationurl_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appfwprofile_xmlvalidationurl_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated appfwprofile_xmlvalidationurl_binding resource")

	// Read the updated state back
	r.readAppfwprofileXmlvalidationurlBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileXmlvalidationurlBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwprofileXmlvalidationurlBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwprofile_xmlvalidationurl_binding resource")

	// For appfwprofile_xmlvalidationurl_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted appfwprofile_xmlvalidationurl_binding resource from state")
}

// Helper function to read appfwprofile_xmlvalidationurl_binding data from API
func (r *AppfwprofileXmlvalidationurlBindingResource) readAppfwprofileXmlvalidationurlBindingFromApi(ctx context.Context, data *AppfwprofileXmlvalidationurlBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Appfwprofile_xmlvalidationurl_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_xmlvalidationurl_binding, got error: %s", err))
		return
	}

	appfwprofile_xmlvalidationurl_bindingSetAttrFromGet(ctx, data, getResponseData)

}
