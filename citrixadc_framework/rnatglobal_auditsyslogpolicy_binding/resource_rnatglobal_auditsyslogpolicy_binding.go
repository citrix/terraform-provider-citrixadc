package rnatglobal_auditsyslogpolicy_binding

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
var _ resource.Resource = &RnatglobalAuditsyslogpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*RnatglobalAuditsyslogpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*RnatglobalAuditsyslogpolicyBindingResource)(nil)

func NewRnatglobalAuditsyslogpolicyBindingResource() resource.Resource {
	return &RnatglobalAuditsyslogpolicyBindingResource{}
}

// RnatglobalAuditsyslogpolicyBindingResource defines the resource implementation.
type RnatglobalAuditsyslogpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *RnatglobalAuditsyslogpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *RnatglobalAuditsyslogpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rnatglobal_auditsyslogpolicy_binding"
}

func (r *RnatglobalAuditsyslogpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *RnatglobalAuditsyslogpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data RnatglobalAuditsyslogpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating rnatglobal_auditsyslogpolicy_binding resource")

	// rnatglobal_auditsyslogpolicy_binding := rnatglobal_auditsyslogpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Rnatglobal_auditsyslogpolicy_binding.Type(), &rnatglobal_auditsyslogpolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create rnatglobal_auditsyslogpolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("rnatglobal_auditsyslogpolicy_binding-config")

	tflog.Trace(ctx, "Created rnatglobal_auditsyslogpolicy_binding resource")

	// Read the updated state back
	r.readRnatglobalAuditsyslogpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RnatglobalAuditsyslogpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data RnatglobalAuditsyslogpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading rnatglobal_auditsyslogpolicy_binding resource")

	r.readRnatglobalAuditsyslogpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RnatglobalAuditsyslogpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data RnatglobalAuditsyslogpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating rnatglobal_auditsyslogpolicy_binding resource")

	// Create API request body from the model
	// rnatglobal_auditsyslogpolicy_binding := rnatglobal_auditsyslogpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Rnatglobal_auditsyslogpolicy_binding.Type(), &rnatglobal_auditsyslogpolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update rnatglobal_auditsyslogpolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated rnatglobal_auditsyslogpolicy_binding resource")

	// Read the updated state back
	r.readRnatglobalAuditsyslogpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RnatglobalAuditsyslogpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data RnatglobalAuditsyslogpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting rnatglobal_auditsyslogpolicy_binding resource")

	// For rnatglobal_auditsyslogpolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted rnatglobal_auditsyslogpolicy_binding resource from state")
}

// Helper function to read rnatglobal_auditsyslogpolicy_binding data from API
func (r *RnatglobalAuditsyslogpolicyBindingResource) readRnatglobalAuditsyslogpolicyBindingFromApi(ctx context.Context, data *RnatglobalAuditsyslogpolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Rnatglobal_auditsyslogpolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read rnatglobal_auditsyslogpolicy_binding, got error: %s", err))
		return
	}

	rnatglobal_auditsyslogpolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
