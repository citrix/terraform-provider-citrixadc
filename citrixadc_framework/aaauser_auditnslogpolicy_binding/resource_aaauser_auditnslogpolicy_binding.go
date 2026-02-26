package aaauser_auditnslogpolicy_binding

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
var _ resource.Resource = &AaauserAuditnslogpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*AaauserAuditnslogpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AaauserAuditnslogpolicyBindingResource)(nil)

func NewAaauserAuditnslogpolicyBindingResource() resource.Resource {
	return &AaauserAuditnslogpolicyBindingResource{}
}

// AaauserAuditnslogpolicyBindingResource defines the resource implementation.
type AaauserAuditnslogpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *AaauserAuditnslogpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AaauserAuditnslogpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_aaauser_auditnslogpolicy_binding"
}

func (r *AaauserAuditnslogpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AaauserAuditnslogpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AaauserAuditnslogpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating aaauser_auditnslogpolicy_binding resource")

	// aaauser_auditnslogpolicy_binding := aaauser_auditnslogpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Aaauser_auditnslogpolicy_binding.Type(), &aaauser_auditnslogpolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create aaauser_auditnslogpolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("aaauser_auditnslogpolicy_binding-config")

	tflog.Trace(ctx, "Created aaauser_auditnslogpolicy_binding resource")

	// Read the updated state back
	r.readAaauserAuditnslogpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaauserAuditnslogpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AaauserAuditnslogpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading aaauser_auditnslogpolicy_binding resource")

	r.readAaauserAuditnslogpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaauserAuditnslogpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AaauserAuditnslogpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating aaauser_auditnslogpolicy_binding resource")

	// Create API request body from the model
	// aaauser_auditnslogpolicy_binding := aaauser_auditnslogpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Aaauser_auditnslogpolicy_binding.Type(), &aaauser_auditnslogpolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update aaauser_auditnslogpolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated aaauser_auditnslogpolicy_binding resource")

	// Read the updated state back
	r.readAaauserAuditnslogpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaauserAuditnslogpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AaauserAuditnslogpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting aaauser_auditnslogpolicy_binding resource")

	// For aaauser_auditnslogpolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted aaauser_auditnslogpolicy_binding resource from state")
}

// Helper function to read aaauser_auditnslogpolicy_binding data from API
func (r *AaauserAuditnslogpolicyBindingResource) readAaauserAuditnslogpolicyBindingFromApi(ctx context.Context, data *AaauserAuditnslogpolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Aaauser_auditnslogpolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read aaauser_auditnslogpolicy_binding, got error: %s", err))
		return
	}

	aaauser_auditnslogpolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
