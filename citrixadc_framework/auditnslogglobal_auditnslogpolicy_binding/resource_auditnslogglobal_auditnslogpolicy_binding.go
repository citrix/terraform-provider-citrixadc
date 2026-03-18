package auditnslogglobal_auditnslogpolicy_binding

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
var _ resource.Resource = &AuditnslogglobalAuditnslogpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*AuditnslogglobalAuditnslogpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AuditnslogglobalAuditnslogpolicyBindingResource)(nil)

func NewAuditnslogglobalAuditnslogpolicyBindingResource() resource.Resource {
	return &AuditnslogglobalAuditnslogpolicyBindingResource{}
}

// AuditnslogglobalAuditnslogpolicyBindingResource defines the resource implementation.
type AuditnslogglobalAuditnslogpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *AuditnslogglobalAuditnslogpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuditnslogglobalAuditnslogpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_auditnslogglobal_auditnslogpolicy_binding"
}

func (r *AuditnslogglobalAuditnslogpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuditnslogglobalAuditnslogpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuditnslogglobalAuditnslogpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating auditnslogglobal_auditnslogpolicy_binding resource")

	// auditnslogglobal_auditnslogpolicy_binding := auditnslogglobal_auditnslogpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Auditnslogglobal_auditnslogpolicy_binding.Type(), &auditnslogglobal_auditnslogpolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create auditnslogglobal_auditnslogpolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("auditnslogglobal_auditnslogpolicy_binding-config")

	tflog.Trace(ctx, "Created auditnslogglobal_auditnslogpolicy_binding resource")

	// Read the updated state back
	r.readAuditnslogglobalAuditnslogpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuditnslogglobalAuditnslogpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuditnslogglobalAuditnslogpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading auditnslogglobal_auditnslogpolicy_binding resource")

	r.readAuditnslogglobalAuditnslogpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuditnslogglobalAuditnslogpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AuditnslogglobalAuditnslogpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating auditnslogglobal_auditnslogpolicy_binding resource")

	// Create API request body from the model
	// auditnslogglobal_auditnslogpolicy_binding := auditnslogglobal_auditnslogpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Auditnslogglobal_auditnslogpolicy_binding.Type(), &auditnslogglobal_auditnslogpolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update auditnslogglobal_auditnslogpolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated auditnslogglobal_auditnslogpolicy_binding resource")

	// Read the updated state back
	r.readAuditnslogglobalAuditnslogpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuditnslogglobalAuditnslogpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuditnslogglobalAuditnslogpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting auditnslogglobal_auditnslogpolicy_binding resource")

	// For auditnslogglobal_auditnslogpolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted auditnslogglobal_auditnslogpolicy_binding resource from state")
}

// Helper function to read auditnslogglobal_auditnslogpolicy_binding data from API
func (r *AuditnslogglobalAuditnslogpolicyBindingResource) readAuditnslogglobalAuditnslogpolicyBindingFromApi(ctx context.Context, data *AuditnslogglobalAuditnslogpolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Auditnslogglobal_auditnslogpolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read auditnslogglobal_auditnslogpolicy_binding, got error: %s", err))
		return
	}

	auditnslogglobal_auditnslogpolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
