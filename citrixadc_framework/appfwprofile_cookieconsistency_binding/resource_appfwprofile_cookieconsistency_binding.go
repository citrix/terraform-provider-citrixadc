package appfwprofile_cookieconsistency_binding

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
var _ resource.Resource = &AppfwprofileCookieconsistencyBindingResource{}
var _ resource.ResourceWithConfigure = (*AppfwprofileCookieconsistencyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwprofileCookieconsistencyBindingResource)(nil)

func NewAppfwprofileCookieconsistencyBindingResource() resource.Resource {
	return &AppfwprofileCookieconsistencyBindingResource{}
}

// AppfwprofileCookieconsistencyBindingResource defines the resource implementation.
type AppfwprofileCookieconsistencyBindingResource struct {
	client *service.NitroClient
}

func (r *AppfwprofileCookieconsistencyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwprofileCookieconsistencyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_cookieconsistency_binding"
}

func (r *AppfwprofileCookieconsistencyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwprofileCookieconsistencyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwprofileCookieconsistencyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwprofile_cookieconsistency_binding resource")

	// appfwprofile_cookieconsistency_binding := appfwprofile_cookieconsistency_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwprofile_cookieconsistency_binding.Type(), &appfwprofile_cookieconsistency_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwprofile_cookieconsistency_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("appfwprofile_cookieconsistency_binding-config")

	tflog.Trace(ctx, "Created appfwprofile_cookieconsistency_binding resource")

	// Read the updated state back
	r.readAppfwprofileCookieconsistencyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileCookieconsistencyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwprofileCookieconsistencyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwprofile_cookieconsistency_binding resource")

	r.readAppfwprofileCookieconsistencyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileCookieconsistencyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AppfwprofileCookieconsistencyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating appfwprofile_cookieconsistency_binding resource")

	// Create API request body from the model
	// appfwprofile_cookieconsistency_binding := appfwprofile_cookieconsistency_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwprofile_cookieconsistency_binding.Type(), &appfwprofile_cookieconsistency_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appfwprofile_cookieconsistency_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated appfwprofile_cookieconsistency_binding resource")

	// Read the updated state back
	r.readAppfwprofileCookieconsistencyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileCookieconsistencyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwprofileCookieconsistencyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwprofile_cookieconsistency_binding resource")

	// For appfwprofile_cookieconsistency_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted appfwprofile_cookieconsistency_binding resource from state")
}

// Helper function to read appfwprofile_cookieconsistency_binding data from API
func (r *AppfwprofileCookieconsistencyBindingResource) readAppfwprofileCookieconsistencyBindingFromApi(ctx context.Context, data *AppfwprofileCookieconsistencyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Appfwprofile_cookieconsistency_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_cookieconsistency_binding, got error: %s", err))
		return
	}

	appfwprofile_cookieconsistency_bindingSetAttrFromGet(ctx, data, getResponseData)

}
