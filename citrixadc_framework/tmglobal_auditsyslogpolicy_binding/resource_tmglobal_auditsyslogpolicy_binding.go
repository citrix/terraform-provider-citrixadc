package tmglobal_auditsyslogpolicy_binding

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
var _ resource.Resource = &TmglobalAuditsyslogpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*TmglobalAuditsyslogpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*TmglobalAuditsyslogpolicyBindingResource)(nil)

func NewTmglobalAuditsyslogpolicyBindingResource() resource.Resource {
	return &TmglobalAuditsyslogpolicyBindingResource{}
}

// TmglobalAuditsyslogpolicyBindingResource defines the resource implementation.
type TmglobalAuditsyslogpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *TmglobalAuditsyslogpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *TmglobalAuditsyslogpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tmglobal_auditsyslogpolicy_binding"
}

func (r *TmglobalAuditsyslogpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *TmglobalAuditsyslogpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data TmglobalAuditsyslogpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating tmglobal_auditsyslogpolicy_binding resource")

	// tmglobal_auditsyslogpolicy_binding := tmglobal_auditsyslogpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Tmglobal_auditsyslogpolicy_binding.Type(), &tmglobal_auditsyslogpolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create tmglobal_auditsyslogpolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("tmglobal_auditsyslogpolicy_binding-config")

	tflog.Trace(ctx, "Created tmglobal_auditsyslogpolicy_binding resource")

	// Read the updated state back
	r.readTmglobalAuditsyslogpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TmglobalAuditsyslogpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data TmglobalAuditsyslogpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading tmglobal_auditsyslogpolicy_binding resource")

	r.readTmglobalAuditsyslogpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TmglobalAuditsyslogpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data TmglobalAuditsyslogpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating tmglobal_auditsyslogpolicy_binding resource")

	// Create API request body from the model
	// tmglobal_auditsyslogpolicy_binding := tmglobal_auditsyslogpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Tmglobal_auditsyslogpolicy_binding.Type(), &tmglobal_auditsyslogpolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update tmglobal_auditsyslogpolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated tmglobal_auditsyslogpolicy_binding resource")

	// Read the updated state back
	r.readTmglobalAuditsyslogpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TmglobalAuditsyslogpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data TmglobalAuditsyslogpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting tmglobal_auditsyslogpolicy_binding resource")

	// For tmglobal_auditsyslogpolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted tmglobal_auditsyslogpolicy_binding resource from state")
}

// Helper function to read tmglobal_auditsyslogpolicy_binding data from API
func (r *TmglobalAuditsyslogpolicyBindingResource) readTmglobalAuditsyslogpolicyBindingFromApi(ctx context.Context, data *TmglobalAuditsyslogpolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Tmglobal_auditsyslogpolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read tmglobal_auditsyslogpolicy_binding, got error: %s", err))
		return
	}

	tmglobal_auditsyslogpolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
