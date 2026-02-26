package systemglobal_auditnslogpolicy_binding

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
var _ resource.Resource = &SystemglobalAuditnslogpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*SystemglobalAuditnslogpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*SystemglobalAuditnslogpolicyBindingResource)(nil)

func NewSystemglobalAuditnslogpolicyBindingResource() resource.Resource {
	return &SystemglobalAuditnslogpolicyBindingResource{}
}

// SystemglobalAuditnslogpolicyBindingResource defines the resource implementation.
type SystemglobalAuditnslogpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *SystemglobalAuditnslogpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SystemglobalAuditnslogpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_systemglobal_auditnslogpolicy_binding"
}

func (r *SystemglobalAuditnslogpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SystemglobalAuditnslogpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SystemglobalAuditnslogpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating systemglobal_auditnslogpolicy_binding resource")

	// systemglobal_auditnslogpolicy_binding := systemglobal_auditnslogpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Systemglobal_auditnslogpolicy_binding.Type(), &systemglobal_auditnslogpolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create systemglobal_auditnslogpolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("systemglobal_auditnslogpolicy_binding-config")

	tflog.Trace(ctx, "Created systemglobal_auditnslogpolicy_binding resource")

	// Read the updated state back
	r.readSystemglobalAuditnslogpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemglobalAuditnslogpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SystemglobalAuditnslogpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading systemglobal_auditnslogpolicy_binding resource")

	r.readSystemglobalAuditnslogpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemglobalAuditnslogpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data SystemglobalAuditnslogpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating systemglobal_auditnslogpolicy_binding resource")

	// Create API request body from the model
	// systemglobal_auditnslogpolicy_binding := systemglobal_auditnslogpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Systemglobal_auditnslogpolicy_binding.Type(), &systemglobal_auditnslogpolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update systemglobal_auditnslogpolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated systemglobal_auditnslogpolicy_binding resource")

	// Read the updated state back
	r.readSystemglobalAuditnslogpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemglobalAuditnslogpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SystemglobalAuditnslogpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting systemglobal_auditnslogpolicy_binding resource")

	// For systemglobal_auditnslogpolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted systemglobal_auditnslogpolicy_binding resource from state")
}

// Helper function to read systemglobal_auditnslogpolicy_binding data from API
func (r *SystemglobalAuditnslogpolicyBindingResource) readSystemglobalAuditnslogpolicyBindingFromApi(ctx context.Context, data *SystemglobalAuditnslogpolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Systemglobal_auditnslogpolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read systemglobal_auditnslogpolicy_binding, got error: %s", err))
		return
	}

	systemglobal_auditnslogpolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
