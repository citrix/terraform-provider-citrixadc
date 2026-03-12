package appfwglobal_auditsyslogpolicy_binding

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
var _ resource.Resource = &AppfwglobalAuditsyslogpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*AppfwglobalAuditsyslogpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwglobalAuditsyslogpolicyBindingResource)(nil)

func NewAppfwglobalAuditsyslogpolicyBindingResource() resource.Resource {
	return &AppfwglobalAuditsyslogpolicyBindingResource{}
}

// AppfwglobalAuditsyslogpolicyBindingResource defines the resource implementation.
type AppfwglobalAuditsyslogpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *AppfwglobalAuditsyslogpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwglobalAuditsyslogpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwglobal_auditsyslogpolicy_binding"
}

func (r *AppfwglobalAuditsyslogpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwglobalAuditsyslogpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwglobalAuditsyslogpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwglobal_auditsyslogpolicy_binding resource")

	// appfwglobal_auditsyslogpolicy_binding := appfwglobal_auditsyslogpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwglobal_auditsyslogpolicy_binding.Type(), &appfwglobal_auditsyslogpolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwglobal_auditsyslogpolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("appfwglobal_auditsyslogpolicy_binding-config")

	tflog.Trace(ctx, "Created appfwglobal_auditsyslogpolicy_binding resource")

	// Read the updated state back
	r.readAppfwglobalAuditsyslogpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwglobalAuditsyslogpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwglobalAuditsyslogpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwglobal_auditsyslogpolicy_binding resource")

	r.readAppfwglobalAuditsyslogpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwglobalAuditsyslogpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AppfwglobalAuditsyslogpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating appfwglobal_auditsyslogpolicy_binding resource")

	// Create API request body from the model
	// appfwglobal_auditsyslogpolicy_binding := appfwglobal_auditsyslogpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwglobal_auditsyslogpolicy_binding.Type(), &appfwglobal_auditsyslogpolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appfwglobal_auditsyslogpolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated appfwglobal_auditsyslogpolicy_binding resource")

	// Read the updated state back
	r.readAppfwglobalAuditsyslogpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwglobalAuditsyslogpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwglobalAuditsyslogpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwglobal_auditsyslogpolicy_binding resource")

	// For appfwglobal_auditsyslogpolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted appfwglobal_auditsyslogpolicy_binding resource from state")
}

// Helper function to read appfwglobal_auditsyslogpolicy_binding data from API
func (r *AppfwglobalAuditsyslogpolicyBindingResource) readAppfwglobalAuditsyslogpolicyBindingFromApi(ctx context.Context, data *AppfwglobalAuditsyslogpolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Appfwglobal_auditsyslogpolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwglobal_auditsyslogpolicy_binding, got error: %s", err))
		return
	}

	appfwglobal_auditsyslogpolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
