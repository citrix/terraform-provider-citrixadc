package appfwprofile_excluderescontenttype_binding

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
var _ resource.Resource = &AppfwprofileExcluderescontenttypeBindingResource{}
var _ resource.ResourceWithConfigure = (*AppfwprofileExcluderescontenttypeBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwprofileExcluderescontenttypeBindingResource)(nil)

func NewAppfwprofileExcluderescontenttypeBindingResource() resource.Resource {
	return &AppfwprofileExcluderescontenttypeBindingResource{}
}

// AppfwprofileExcluderescontenttypeBindingResource defines the resource implementation.
type AppfwprofileExcluderescontenttypeBindingResource struct {
	client *service.NitroClient
}

func (r *AppfwprofileExcluderescontenttypeBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwprofileExcluderescontenttypeBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_excluderescontenttype_binding"
}

func (r *AppfwprofileExcluderescontenttypeBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwprofileExcluderescontenttypeBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwprofileExcluderescontenttypeBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwprofile_excluderescontenttype_binding resource")

	// appfwprofile_excluderescontenttype_binding := appfwprofile_excluderescontenttype_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwprofile_excluderescontenttype_binding.Type(), &appfwprofile_excluderescontenttype_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwprofile_excluderescontenttype_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("appfwprofile_excluderescontenttype_binding-config")

	tflog.Trace(ctx, "Created appfwprofile_excluderescontenttype_binding resource")

	// Read the updated state back
	r.readAppfwprofileExcluderescontenttypeBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileExcluderescontenttypeBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwprofileExcluderescontenttypeBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwprofile_excluderescontenttype_binding resource")

	r.readAppfwprofileExcluderescontenttypeBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileExcluderescontenttypeBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AppfwprofileExcluderescontenttypeBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating appfwprofile_excluderescontenttype_binding resource")

	// Create API request body from the model
	// appfwprofile_excluderescontenttype_binding := appfwprofile_excluderescontenttype_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwprofile_excluderescontenttype_binding.Type(), &appfwprofile_excluderescontenttype_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appfwprofile_excluderescontenttype_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated appfwprofile_excluderescontenttype_binding resource")

	// Read the updated state back
	r.readAppfwprofileExcluderescontenttypeBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileExcluderescontenttypeBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwprofileExcluderescontenttypeBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwprofile_excluderescontenttype_binding resource")

	// For appfwprofile_excluderescontenttype_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted appfwprofile_excluderescontenttype_binding resource from state")
}

// Helper function to read appfwprofile_excluderescontenttype_binding data from API
func (r *AppfwprofileExcluderescontenttypeBindingResource) readAppfwprofileExcluderescontenttypeBindingFromApi(ctx context.Context, data *AppfwprofileExcluderescontenttypeBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Appfwprofile_excluderescontenttype_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_excluderescontenttype_binding, got error: %s", err))
		return
	}

	appfwprofile_excluderescontenttype_bindingSetAttrFromGet(ctx, data, getResponseData)

}
