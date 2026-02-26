package appfwprofile_jsonxssurl_binding

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
var _ resource.Resource = &AppfwprofileJsonxssurlBindingResource{}
var _ resource.ResourceWithConfigure = (*AppfwprofileJsonxssurlBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwprofileJsonxssurlBindingResource)(nil)

func NewAppfwprofileJsonxssurlBindingResource() resource.Resource {
	return &AppfwprofileJsonxssurlBindingResource{}
}

// AppfwprofileJsonxssurlBindingResource defines the resource implementation.
type AppfwprofileJsonxssurlBindingResource struct {
	client *service.NitroClient
}

func (r *AppfwprofileJsonxssurlBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwprofileJsonxssurlBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_jsonxssurl_binding"
}

func (r *AppfwprofileJsonxssurlBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwprofileJsonxssurlBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwprofileJsonxssurlBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwprofile_jsonxssurl_binding resource")

	// appfwprofile_jsonxssurl_binding := appfwprofile_jsonxssurl_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwprofile_jsonxssurl_binding.Type(), &appfwprofile_jsonxssurl_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwprofile_jsonxssurl_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("appfwprofile_jsonxssurl_binding-config")

	tflog.Trace(ctx, "Created appfwprofile_jsonxssurl_binding resource")

	// Read the updated state back
	r.readAppfwprofileJsonxssurlBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileJsonxssurlBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwprofileJsonxssurlBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwprofile_jsonxssurl_binding resource")

	r.readAppfwprofileJsonxssurlBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileJsonxssurlBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AppfwprofileJsonxssurlBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating appfwprofile_jsonxssurl_binding resource")

	// Create API request body from the model
	// appfwprofile_jsonxssurl_binding := appfwprofile_jsonxssurl_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwprofile_jsonxssurl_binding.Type(), &appfwprofile_jsonxssurl_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appfwprofile_jsonxssurl_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated appfwprofile_jsonxssurl_binding resource")

	// Read the updated state back
	r.readAppfwprofileJsonxssurlBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileJsonxssurlBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwprofileJsonxssurlBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwprofile_jsonxssurl_binding resource")

	// For appfwprofile_jsonxssurl_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted appfwprofile_jsonxssurl_binding resource from state")
}

// Helper function to read appfwprofile_jsonxssurl_binding data from API
func (r *AppfwprofileJsonxssurlBindingResource) readAppfwprofileJsonxssurlBindingFromApi(ctx context.Context, data *AppfwprofileJsonxssurlBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Appfwprofile_jsonxssurl_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_jsonxssurl_binding, got error: %s", err))
		return
	}

	appfwprofile_jsonxssurl_bindingSetAttrFromGet(ctx, data, getResponseData)

}
