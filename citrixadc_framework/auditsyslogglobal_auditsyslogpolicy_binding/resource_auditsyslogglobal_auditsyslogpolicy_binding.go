package auditsyslogglobal_auditsyslogpolicy_binding

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
var _ resource.Resource = &AuditsyslogglobalAuditsyslogpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*AuditsyslogglobalAuditsyslogpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AuditsyslogglobalAuditsyslogpolicyBindingResource)(nil)

func NewAuditsyslogglobalAuditsyslogpolicyBindingResource() resource.Resource {
	return &AuditsyslogglobalAuditsyslogpolicyBindingResource{}
}

// AuditsyslogglobalAuditsyslogpolicyBindingResource defines the resource implementation.
type AuditsyslogglobalAuditsyslogpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *AuditsyslogglobalAuditsyslogpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuditsyslogglobalAuditsyslogpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_auditsyslogglobal_auditsyslogpolicy_binding"
}

func (r *AuditsyslogglobalAuditsyslogpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuditsyslogglobalAuditsyslogpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuditsyslogglobalAuditsyslogpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating auditsyslogglobal_auditsyslogpolicy_binding resource")

	// auditsyslogglobal_auditsyslogpolicy_binding := auditsyslogglobal_auditsyslogpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Auditsyslogglobal_auditsyslogpolicy_binding.Type(), &auditsyslogglobal_auditsyslogpolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create auditsyslogglobal_auditsyslogpolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("auditsyslogglobal_auditsyslogpolicy_binding-config")

	tflog.Trace(ctx, "Created auditsyslogglobal_auditsyslogpolicy_binding resource")

	// Read the updated state back
	r.readAuditsyslogglobalAuditsyslogpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuditsyslogglobalAuditsyslogpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuditsyslogglobalAuditsyslogpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading auditsyslogglobal_auditsyslogpolicy_binding resource")

	r.readAuditsyslogglobalAuditsyslogpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuditsyslogglobalAuditsyslogpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AuditsyslogglobalAuditsyslogpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating auditsyslogglobal_auditsyslogpolicy_binding resource")

	// Create API request body from the model
	// auditsyslogglobal_auditsyslogpolicy_binding := auditsyslogglobal_auditsyslogpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Auditsyslogglobal_auditsyslogpolicy_binding.Type(), &auditsyslogglobal_auditsyslogpolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update auditsyslogglobal_auditsyslogpolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated auditsyslogglobal_auditsyslogpolicy_binding resource")

	// Read the updated state back
	r.readAuditsyslogglobalAuditsyslogpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuditsyslogglobalAuditsyslogpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuditsyslogglobalAuditsyslogpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting auditsyslogglobal_auditsyslogpolicy_binding resource")

	// For auditsyslogglobal_auditsyslogpolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted auditsyslogglobal_auditsyslogpolicy_binding resource from state")
}

// Helper function to read auditsyslogglobal_auditsyslogpolicy_binding data from API
func (r *AuditsyslogglobalAuditsyslogpolicyBindingResource) readAuditsyslogglobalAuditsyslogpolicyBindingFromApi(ctx context.Context, data *AuditsyslogglobalAuditsyslogpolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Auditsyslogglobal_auditsyslogpolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read auditsyslogglobal_auditsyslogpolicy_binding, got error: %s", err))
		return
	}

	auditsyslogglobal_auditsyslogpolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
