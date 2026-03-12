package appfwprofile_xmlattachmenturl_binding

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
var _ resource.Resource = &AppfwprofileXmlattachmenturlBindingResource{}
var _ resource.ResourceWithConfigure = (*AppfwprofileXmlattachmenturlBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwprofileXmlattachmenturlBindingResource)(nil)

func NewAppfwprofileXmlattachmenturlBindingResource() resource.Resource {
	return &AppfwprofileXmlattachmenturlBindingResource{}
}

// AppfwprofileXmlattachmenturlBindingResource defines the resource implementation.
type AppfwprofileXmlattachmenturlBindingResource struct {
	client *service.NitroClient
}

func (r *AppfwprofileXmlattachmenturlBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwprofileXmlattachmenturlBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_xmlattachmenturl_binding"
}

func (r *AppfwprofileXmlattachmenturlBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwprofileXmlattachmenturlBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwprofileXmlattachmenturlBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwprofile_xmlattachmenturl_binding resource")

	// appfwprofile_xmlattachmenturl_binding := appfwprofile_xmlattachmenturl_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwprofile_xmlattachmenturl_binding.Type(), &appfwprofile_xmlattachmenturl_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwprofile_xmlattachmenturl_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("appfwprofile_xmlattachmenturl_binding-config")

	tflog.Trace(ctx, "Created appfwprofile_xmlattachmenturl_binding resource")

	// Read the updated state back
	r.readAppfwprofileXmlattachmenturlBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileXmlattachmenturlBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwprofileXmlattachmenturlBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwprofile_xmlattachmenturl_binding resource")

	r.readAppfwprofileXmlattachmenturlBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileXmlattachmenturlBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AppfwprofileXmlattachmenturlBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating appfwprofile_xmlattachmenturl_binding resource")

	// Create API request body from the model
	// appfwprofile_xmlattachmenturl_binding := appfwprofile_xmlattachmenturl_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwprofile_xmlattachmenturl_binding.Type(), &appfwprofile_xmlattachmenturl_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appfwprofile_xmlattachmenturl_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated appfwprofile_xmlattachmenturl_binding resource")

	// Read the updated state back
	r.readAppfwprofileXmlattachmenturlBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileXmlattachmenturlBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwprofileXmlattachmenturlBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwprofile_xmlattachmenturl_binding resource")

	// For appfwprofile_xmlattachmenturl_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted appfwprofile_xmlattachmenturl_binding resource from state")
}

// Helper function to read appfwprofile_xmlattachmenturl_binding data from API
func (r *AppfwprofileXmlattachmenturlBindingResource) readAppfwprofileXmlattachmenturlBindingFromApi(ctx context.Context, data *AppfwprofileXmlattachmenturlBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Appfwprofile_xmlattachmenturl_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_xmlattachmenturl_binding, got error: %s", err))
		return
	}

	appfwprofile_xmlattachmenturl_bindingSetAttrFromGet(ctx, data, getResponseData)

}
