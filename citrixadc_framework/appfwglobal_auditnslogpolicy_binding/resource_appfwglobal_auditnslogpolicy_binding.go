package appfwglobal_auditnslogpolicy_binding

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
var _ resource.Resource = &AppfwglobalAuditnslogpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*AppfwglobalAuditnslogpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwglobalAuditnslogpolicyBindingResource)(nil)

func NewAppfwglobalAuditnslogpolicyBindingResource() resource.Resource {
	return &AppfwglobalAuditnslogpolicyBindingResource{}
}

// AppfwglobalAuditnslogpolicyBindingResource defines the resource implementation.
type AppfwglobalAuditnslogpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *AppfwglobalAuditnslogpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwglobalAuditnslogpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwglobal_auditnslogpolicy_binding"
}

func (r *AppfwglobalAuditnslogpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwglobalAuditnslogpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwglobalAuditnslogpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwglobal_auditnslogpolicy_binding resource")

	// appfwglobal_auditnslogpolicy_binding := appfwglobal_auditnslogpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwglobal_auditnslogpolicy_binding.Type(), &appfwglobal_auditnslogpolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwglobal_auditnslogpolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("appfwglobal_auditnslogpolicy_binding-config")

	tflog.Trace(ctx, "Created appfwglobal_auditnslogpolicy_binding resource")

	// Read the updated state back
	r.readAppfwglobalAuditnslogpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwglobalAuditnslogpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwglobalAuditnslogpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwglobal_auditnslogpolicy_binding resource")

	r.readAppfwglobalAuditnslogpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwglobalAuditnslogpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AppfwglobalAuditnslogpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating appfwglobal_auditnslogpolicy_binding resource")

	// Create API request body from the model
	// appfwglobal_auditnslogpolicy_binding := appfwglobal_auditnslogpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwglobal_auditnslogpolicy_binding.Type(), &appfwglobal_auditnslogpolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appfwglobal_auditnslogpolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated appfwglobal_auditnslogpolicy_binding resource")

	// Read the updated state back
	r.readAppfwglobalAuditnslogpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwglobalAuditnslogpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwglobalAuditnslogpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwglobal_auditnslogpolicy_binding resource")

	// For appfwglobal_auditnslogpolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted appfwglobal_auditnslogpolicy_binding resource from state")
}

// Helper function to read appfwglobal_auditnslogpolicy_binding data from API
func (r *AppfwglobalAuditnslogpolicyBindingResource) readAppfwglobalAuditnslogpolicyBindingFromApi(ctx context.Context, data *AppfwglobalAuditnslogpolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Appfwglobal_auditnslogpolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwglobal_auditnslogpolicy_binding, got error: %s", err))
		return
	}

	appfwglobal_auditnslogpolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
