package aaauser_auditsyslogpolicy_binding

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
var _ resource.Resource = &AaauserAuditsyslogpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*AaauserAuditsyslogpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AaauserAuditsyslogpolicyBindingResource)(nil)

func NewAaauserAuditsyslogpolicyBindingResource() resource.Resource {
	return &AaauserAuditsyslogpolicyBindingResource{}
}

// AaauserAuditsyslogpolicyBindingResource defines the resource implementation.
type AaauserAuditsyslogpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *AaauserAuditsyslogpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AaauserAuditsyslogpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_aaauser_auditsyslogpolicy_binding"
}

func (r *AaauserAuditsyslogpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AaauserAuditsyslogpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AaauserAuditsyslogpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating aaauser_auditsyslogpolicy_binding resource")

	// aaauser_auditsyslogpolicy_binding := aaauser_auditsyslogpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Aaauser_auditsyslogpolicy_binding.Type(), &aaauser_auditsyslogpolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create aaauser_auditsyslogpolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("aaauser_auditsyslogpolicy_binding-config")

	tflog.Trace(ctx, "Created aaauser_auditsyslogpolicy_binding resource")

	// Read the updated state back
	r.readAaauserAuditsyslogpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaauserAuditsyslogpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AaauserAuditsyslogpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading aaauser_auditsyslogpolicy_binding resource")

	r.readAaauserAuditsyslogpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaauserAuditsyslogpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AaauserAuditsyslogpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating aaauser_auditsyslogpolicy_binding resource")

	// Create API request body from the model
	// aaauser_auditsyslogpolicy_binding := aaauser_auditsyslogpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Aaauser_auditsyslogpolicy_binding.Type(), &aaauser_auditsyslogpolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update aaauser_auditsyslogpolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated aaauser_auditsyslogpolicy_binding resource")

	// Read the updated state back
	r.readAaauserAuditsyslogpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaauserAuditsyslogpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AaauserAuditsyslogpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting aaauser_auditsyslogpolicy_binding resource")

	// For aaauser_auditsyslogpolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted aaauser_auditsyslogpolicy_binding resource from state")
}

// Helper function to read aaauser_auditsyslogpolicy_binding data from API
func (r *AaauserAuditsyslogpolicyBindingResource) readAaauserAuditsyslogpolicyBindingFromApi(ctx context.Context, data *AaauserAuditsyslogpolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Aaauser_auditsyslogpolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read aaauser_auditsyslogpolicy_binding, got error: %s", err))
		return
	}

	aaauser_auditsyslogpolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
