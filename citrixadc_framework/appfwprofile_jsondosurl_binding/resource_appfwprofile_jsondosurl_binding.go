package appfwprofile_jsondosurl_binding

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
var _ resource.Resource = &AppfwprofileJsondosurlBindingResource{}
var _ resource.ResourceWithConfigure = (*AppfwprofileJsondosurlBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwprofileJsondosurlBindingResource)(nil)

func NewAppfwprofileJsondosurlBindingResource() resource.Resource {
	return &AppfwprofileJsondosurlBindingResource{}
}

// AppfwprofileJsondosurlBindingResource defines the resource implementation.
type AppfwprofileJsondosurlBindingResource struct {
	client *service.NitroClient
}

func (r *AppfwprofileJsondosurlBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwprofileJsondosurlBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_jsondosurl_binding"
}

func (r *AppfwprofileJsondosurlBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwprofileJsondosurlBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwprofileJsondosurlBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwprofile_jsondosurl_binding resource")

	// appfwprofile_jsondosurl_binding := appfwprofile_jsondosurl_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwprofile_jsondosurl_binding.Type(), &appfwprofile_jsondosurl_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwprofile_jsondosurl_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("appfwprofile_jsondosurl_binding-config")

	tflog.Trace(ctx, "Created appfwprofile_jsondosurl_binding resource")

	// Read the updated state back
	r.readAppfwprofileJsondosurlBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileJsondosurlBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwprofileJsondosurlBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwprofile_jsondosurl_binding resource")

	r.readAppfwprofileJsondosurlBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileJsondosurlBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AppfwprofileJsondosurlBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating appfwprofile_jsondosurl_binding resource")

	// Create API request body from the model
	// appfwprofile_jsondosurl_binding := appfwprofile_jsondosurl_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwprofile_jsondosurl_binding.Type(), &appfwprofile_jsondosurl_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appfwprofile_jsondosurl_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated appfwprofile_jsondosurl_binding resource")

	// Read the updated state back
	r.readAppfwprofileJsondosurlBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileJsondosurlBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwprofileJsondosurlBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwprofile_jsondosurl_binding resource")

	// For appfwprofile_jsondosurl_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted appfwprofile_jsondosurl_binding resource from state")
}

// Helper function to read appfwprofile_jsondosurl_binding data from API
func (r *AppfwprofileJsondosurlBindingResource) readAppfwprofileJsondosurlBindingFromApi(ctx context.Context, data *AppfwprofileJsondosurlBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Appfwprofile_jsondosurl_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_jsondosurl_binding, got error: %s", err))
		return
	}

	appfwprofile_jsondosurl_bindingSetAttrFromGet(ctx, data, getResponseData)

}
