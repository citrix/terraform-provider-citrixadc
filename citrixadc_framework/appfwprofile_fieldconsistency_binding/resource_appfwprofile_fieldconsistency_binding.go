package appfwprofile_fieldconsistency_binding

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
var _ resource.Resource = &AppfwprofileFieldconsistencyBindingResource{}
var _ resource.ResourceWithConfigure = (*AppfwprofileFieldconsistencyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwprofileFieldconsistencyBindingResource)(nil)

func NewAppfwprofileFieldconsistencyBindingResource() resource.Resource {
	return &AppfwprofileFieldconsistencyBindingResource{}
}

// AppfwprofileFieldconsistencyBindingResource defines the resource implementation.
type AppfwprofileFieldconsistencyBindingResource struct {
	client *service.NitroClient
}

func (r *AppfwprofileFieldconsistencyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwprofileFieldconsistencyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_fieldconsistency_binding"
}

func (r *AppfwprofileFieldconsistencyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwprofileFieldconsistencyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwprofileFieldconsistencyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwprofile_fieldconsistency_binding resource")

	// appfwprofile_fieldconsistency_binding := appfwprofile_fieldconsistency_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwprofile_fieldconsistency_binding.Type(), &appfwprofile_fieldconsistency_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwprofile_fieldconsistency_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("appfwprofile_fieldconsistency_binding-config")

	tflog.Trace(ctx, "Created appfwprofile_fieldconsistency_binding resource")

	// Read the updated state back
	r.readAppfwprofileFieldconsistencyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileFieldconsistencyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwprofileFieldconsistencyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwprofile_fieldconsistency_binding resource")

	r.readAppfwprofileFieldconsistencyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileFieldconsistencyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AppfwprofileFieldconsistencyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating appfwprofile_fieldconsistency_binding resource")

	// Create API request body from the model
	// appfwprofile_fieldconsistency_binding := appfwprofile_fieldconsistency_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwprofile_fieldconsistency_binding.Type(), &appfwprofile_fieldconsistency_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appfwprofile_fieldconsistency_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated appfwprofile_fieldconsistency_binding resource")

	// Read the updated state back
	r.readAppfwprofileFieldconsistencyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileFieldconsistencyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwprofileFieldconsistencyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwprofile_fieldconsistency_binding resource")

	// For appfwprofile_fieldconsistency_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted appfwprofile_fieldconsistency_binding resource from state")
}

// Helper function to read appfwprofile_fieldconsistency_binding data from API
func (r *AppfwprofileFieldconsistencyBindingResource) readAppfwprofileFieldconsistencyBindingFromApi(ctx context.Context, data *AppfwprofileFieldconsistencyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Appfwprofile_fieldconsistency_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_fieldconsistency_binding, got error: %s", err))
		return
	}

	appfwprofile_fieldconsistency_bindingSetAttrFromGet(ctx, data, getResponseData)

}
