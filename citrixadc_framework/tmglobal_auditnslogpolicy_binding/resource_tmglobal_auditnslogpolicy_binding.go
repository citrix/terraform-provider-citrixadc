package tmglobal_auditnslogpolicy_binding

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
var _ resource.Resource = &TmglobalAuditnslogpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*TmglobalAuditnslogpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*TmglobalAuditnslogpolicyBindingResource)(nil)

func NewTmglobalAuditnslogpolicyBindingResource() resource.Resource {
	return &TmglobalAuditnslogpolicyBindingResource{}
}

// TmglobalAuditnslogpolicyBindingResource defines the resource implementation.
type TmglobalAuditnslogpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *TmglobalAuditnslogpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *TmglobalAuditnslogpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tmglobal_auditnslogpolicy_binding"
}

func (r *TmglobalAuditnslogpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *TmglobalAuditnslogpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data TmglobalAuditnslogpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating tmglobal_auditnslogpolicy_binding resource")

	// tmglobal_auditnslogpolicy_binding := tmglobal_auditnslogpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Tmglobal_auditnslogpolicy_binding.Type(), &tmglobal_auditnslogpolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create tmglobal_auditnslogpolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("tmglobal_auditnslogpolicy_binding-config")

	tflog.Trace(ctx, "Created tmglobal_auditnslogpolicy_binding resource")

	// Read the updated state back
	r.readTmglobalAuditnslogpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TmglobalAuditnslogpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data TmglobalAuditnslogpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading tmglobal_auditnslogpolicy_binding resource")

	r.readTmglobalAuditnslogpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TmglobalAuditnslogpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data TmglobalAuditnslogpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating tmglobal_auditnslogpolicy_binding resource")

	// Create API request body from the model
	// tmglobal_auditnslogpolicy_binding := tmglobal_auditnslogpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Tmglobal_auditnslogpolicy_binding.Type(), &tmglobal_auditnslogpolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update tmglobal_auditnslogpolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated tmglobal_auditnslogpolicy_binding resource")

	// Read the updated state back
	r.readTmglobalAuditnslogpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TmglobalAuditnslogpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data TmglobalAuditnslogpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting tmglobal_auditnslogpolicy_binding resource")

	// For tmglobal_auditnslogpolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted tmglobal_auditnslogpolicy_binding resource from state")
}

// Helper function to read tmglobal_auditnslogpolicy_binding data from API
func (r *TmglobalAuditnslogpolicyBindingResource) readTmglobalAuditnslogpolicyBindingFromApi(ctx context.Context, data *TmglobalAuditnslogpolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Tmglobal_auditnslogpolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read tmglobal_auditnslogpolicy_binding, got error: %s", err))
		return
	}

	tmglobal_auditnslogpolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
