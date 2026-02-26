package appfwprofile_jsonsqlurl_binding

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
var _ resource.Resource = &AppfwprofileJsonsqlurlBindingResource{}
var _ resource.ResourceWithConfigure = (*AppfwprofileJsonsqlurlBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwprofileJsonsqlurlBindingResource)(nil)

func NewAppfwprofileJsonsqlurlBindingResource() resource.Resource {
	return &AppfwprofileJsonsqlurlBindingResource{}
}

// AppfwprofileJsonsqlurlBindingResource defines the resource implementation.
type AppfwprofileJsonsqlurlBindingResource struct {
	client *service.NitroClient
}

func (r *AppfwprofileJsonsqlurlBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwprofileJsonsqlurlBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_jsonsqlurl_binding"
}

func (r *AppfwprofileJsonsqlurlBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwprofileJsonsqlurlBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwprofileJsonsqlurlBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwprofile_jsonsqlurl_binding resource")

	// appfwprofile_jsonsqlurl_binding := appfwprofile_jsonsqlurl_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwprofile_jsonsqlurl_binding.Type(), &appfwprofile_jsonsqlurl_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwprofile_jsonsqlurl_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("appfwprofile_jsonsqlurl_binding-config")

	tflog.Trace(ctx, "Created appfwprofile_jsonsqlurl_binding resource")

	// Read the updated state back
	r.readAppfwprofileJsonsqlurlBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileJsonsqlurlBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwprofileJsonsqlurlBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwprofile_jsonsqlurl_binding resource")

	r.readAppfwprofileJsonsqlurlBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileJsonsqlurlBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AppfwprofileJsonsqlurlBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating appfwprofile_jsonsqlurl_binding resource")

	// Create API request body from the model
	// appfwprofile_jsonsqlurl_binding := appfwprofile_jsonsqlurl_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwprofile_jsonsqlurl_binding.Type(), &appfwprofile_jsonsqlurl_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appfwprofile_jsonsqlurl_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated appfwprofile_jsonsqlurl_binding resource")

	// Read the updated state back
	r.readAppfwprofileJsonsqlurlBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileJsonsqlurlBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwprofileJsonsqlurlBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwprofile_jsonsqlurl_binding resource")

	// For appfwprofile_jsonsqlurl_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted appfwprofile_jsonsqlurl_binding resource from state")
}

// Helper function to read appfwprofile_jsonsqlurl_binding data from API
func (r *AppfwprofileJsonsqlurlBindingResource) readAppfwprofileJsonsqlurlBindingFromApi(ctx context.Context, data *AppfwprofileJsonsqlurlBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Appfwprofile_jsonsqlurl_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_jsonsqlurl_binding, got error: %s", err))
		return
	}

	appfwprofile_jsonsqlurl_bindingSetAttrFromGet(ctx, data, getResponseData)

}
