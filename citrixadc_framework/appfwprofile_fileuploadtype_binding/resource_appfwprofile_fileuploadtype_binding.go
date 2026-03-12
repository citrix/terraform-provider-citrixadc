package appfwprofile_fileuploadtype_binding

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
var _ resource.Resource = &AppfwprofileFileuploadtypeBindingResource{}
var _ resource.ResourceWithConfigure = (*AppfwprofileFileuploadtypeBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwprofileFileuploadtypeBindingResource)(nil)

func NewAppfwprofileFileuploadtypeBindingResource() resource.Resource {
	return &AppfwprofileFileuploadtypeBindingResource{}
}

// AppfwprofileFileuploadtypeBindingResource defines the resource implementation.
type AppfwprofileFileuploadtypeBindingResource struct {
	client *service.NitroClient
}

func (r *AppfwprofileFileuploadtypeBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwprofileFileuploadtypeBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_fileuploadtype_binding"
}

func (r *AppfwprofileFileuploadtypeBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwprofileFileuploadtypeBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwprofileFileuploadtypeBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwprofile_fileuploadtype_binding resource")

	// appfwprofile_fileuploadtype_binding := appfwprofile_fileuploadtype_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwprofile_fileuploadtype_binding.Type(), &appfwprofile_fileuploadtype_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwprofile_fileuploadtype_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("appfwprofile_fileuploadtype_binding-config")

	tflog.Trace(ctx, "Created appfwprofile_fileuploadtype_binding resource")

	// Read the updated state back
	r.readAppfwprofileFileuploadtypeBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileFileuploadtypeBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwprofileFileuploadtypeBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwprofile_fileuploadtype_binding resource")

	r.readAppfwprofileFileuploadtypeBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileFileuploadtypeBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AppfwprofileFileuploadtypeBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating appfwprofile_fileuploadtype_binding resource")

	// Create API request body from the model
	// appfwprofile_fileuploadtype_binding := appfwprofile_fileuploadtype_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwprofile_fileuploadtype_binding.Type(), &appfwprofile_fileuploadtype_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appfwprofile_fileuploadtype_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated appfwprofile_fileuploadtype_binding resource")

	// Read the updated state back
	r.readAppfwprofileFileuploadtypeBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileFileuploadtypeBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwprofileFileuploadtypeBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwprofile_fileuploadtype_binding resource")

	// For appfwprofile_fileuploadtype_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted appfwprofile_fileuploadtype_binding resource from state")
}

// Helper function to read appfwprofile_fileuploadtype_binding data from API
func (r *AppfwprofileFileuploadtypeBindingResource) readAppfwprofileFileuploadtypeBindingFromApi(ctx context.Context, data *AppfwprofileFileuploadtypeBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Appfwprofile_fileuploadtype_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_fileuploadtype_binding, got error: %s", err))
		return
	}

	appfwprofile_fileuploadtype_bindingSetAttrFromGet(ctx, data, getResponseData)

}
