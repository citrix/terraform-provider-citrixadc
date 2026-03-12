package appfwprofile_xmlwsiurl_binding

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
var _ resource.Resource = &AppfwprofileXmlwsiurlBindingResource{}
var _ resource.ResourceWithConfigure = (*AppfwprofileXmlwsiurlBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwprofileXmlwsiurlBindingResource)(nil)

func NewAppfwprofileXmlwsiurlBindingResource() resource.Resource {
	return &AppfwprofileXmlwsiurlBindingResource{}
}

// AppfwprofileXmlwsiurlBindingResource defines the resource implementation.
type AppfwprofileXmlwsiurlBindingResource struct {
	client *service.NitroClient
}

func (r *AppfwprofileXmlwsiurlBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwprofileXmlwsiurlBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_xmlwsiurl_binding"
}

func (r *AppfwprofileXmlwsiurlBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwprofileXmlwsiurlBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwprofileXmlwsiurlBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwprofile_xmlwsiurl_binding resource")

	// appfwprofile_xmlwsiurl_binding := appfwprofile_xmlwsiurl_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwprofile_xmlwsiurl_binding.Type(), &appfwprofile_xmlwsiurl_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwprofile_xmlwsiurl_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("appfwprofile_xmlwsiurl_binding-config")

	tflog.Trace(ctx, "Created appfwprofile_xmlwsiurl_binding resource")

	// Read the updated state back
	r.readAppfwprofileXmlwsiurlBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileXmlwsiurlBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwprofileXmlwsiurlBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwprofile_xmlwsiurl_binding resource")

	r.readAppfwprofileXmlwsiurlBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileXmlwsiurlBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AppfwprofileXmlwsiurlBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating appfwprofile_xmlwsiurl_binding resource")

	// Create API request body from the model
	// appfwprofile_xmlwsiurl_binding := appfwprofile_xmlwsiurl_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwprofile_xmlwsiurl_binding.Type(), &appfwprofile_xmlwsiurl_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appfwprofile_xmlwsiurl_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated appfwprofile_xmlwsiurl_binding resource")

	// Read the updated state back
	r.readAppfwprofileXmlwsiurlBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileXmlwsiurlBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwprofileXmlwsiurlBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwprofile_xmlwsiurl_binding resource")

	// For appfwprofile_xmlwsiurl_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted appfwprofile_xmlwsiurl_binding resource from state")
}

// Helper function to read appfwprofile_xmlwsiurl_binding data from API
func (r *AppfwprofileXmlwsiurlBindingResource) readAppfwprofileXmlwsiurlBindingFromApi(ctx context.Context, data *AppfwprofileXmlwsiurlBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Appfwprofile_xmlwsiurl_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_xmlwsiurl_binding, got error: %s", err))
		return
	}

	appfwprofile_xmlwsiurl_bindingSetAttrFromGet(ctx, data, getResponseData)

}
