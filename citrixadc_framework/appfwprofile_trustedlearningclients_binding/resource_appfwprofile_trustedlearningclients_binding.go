package appfwprofile_trustedlearningclients_binding

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
var _ resource.Resource = &AppfwprofileTrustedlearningclientsBindingResource{}
var _ resource.ResourceWithConfigure = (*AppfwprofileTrustedlearningclientsBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwprofileTrustedlearningclientsBindingResource)(nil)

func NewAppfwprofileTrustedlearningclientsBindingResource() resource.Resource {
	return &AppfwprofileTrustedlearningclientsBindingResource{}
}

// AppfwprofileTrustedlearningclientsBindingResource defines the resource implementation.
type AppfwprofileTrustedlearningclientsBindingResource struct {
	client *service.NitroClient
}

func (r *AppfwprofileTrustedlearningclientsBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwprofileTrustedlearningclientsBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_trustedlearningclients_binding"
}

func (r *AppfwprofileTrustedlearningclientsBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwprofileTrustedlearningclientsBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwprofileTrustedlearningclientsBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwprofile_trustedlearningclients_binding resource")

	// appfwprofile_trustedlearningclients_binding := appfwprofile_trustedlearningclients_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwprofile_trustedlearningclients_binding.Type(), &appfwprofile_trustedlearningclients_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwprofile_trustedlearningclients_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("appfwprofile_trustedlearningclients_binding-config")

	tflog.Trace(ctx, "Created appfwprofile_trustedlearningclients_binding resource")

	// Read the updated state back
	r.readAppfwprofileTrustedlearningclientsBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileTrustedlearningclientsBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwprofileTrustedlearningclientsBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwprofile_trustedlearningclients_binding resource")

	r.readAppfwprofileTrustedlearningclientsBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileTrustedlearningclientsBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AppfwprofileTrustedlearningclientsBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating appfwprofile_trustedlearningclients_binding resource")

	// Create API request body from the model
	// appfwprofile_trustedlearningclients_binding := appfwprofile_trustedlearningclients_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwprofile_trustedlearningclients_binding.Type(), &appfwprofile_trustedlearningclients_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appfwprofile_trustedlearningclients_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated appfwprofile_trustedlearningclients_binding resource")

	// Read the updated state back
	r.readAppfwprofileTrustedlearningclientsBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileTrustedlearningclientsBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwprofileTrustedlearningclientsBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwprofile_trustedlearningclients_binding resource")

	// For appfwprofile_trustedlearningclients_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted appfwprofile_trustedlearningclients_binding resource from state")
}

// Helper function to read appfwprofile_trustedlearningclients_binding data from API
func (r *AppfwprofileTrustedlearningclientsBindingResource) readAppfwprofileTrustedlearningclientsBindingFromApi(ctx context.Context, data *AppfwprofileTrustedlearningclientsBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Appfwprofile_trustedlearningclients_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_trustedlearningclients_binding, got error: %s", err))
		return
	}

	appfwprofile_trustedlearningclients_bindingSetAttrFromGet(ctx, data, getResponseData)

}
