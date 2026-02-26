package authenticationvserver_auditsyslogpolicy_binding

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
var _ resource.Resource = &AuthenticationvserverAuditsyslogpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationvserverAuditsyslogpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationvserverAuditsyslogpolicyBindingResource)(nil)

func NewAuthenticationvserverAuditsyslogpolicyBindingResource() resource.Resource {
	return &AuthenticationvserverAuditsyslogpolicyBindingResource{}
}

// AuthenticationvserverAuditsyslogpolicyBindingResource defines the resource implementation.
type AuthenticationvserverAuditsyslogpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *AuthenticationvserverAuditsyslogpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationvserverAuditsyslogpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationvserver_auditsyslogpolicy_binding"
}

func (r *AuthenticationvserverAuditsyslogpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationvserverAuditsyslogpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuthenticationvserverAuditsyslogpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationvserver_auditsyslogpolicy_binding resource")

	// authenticationvserver_auditsyslogpolicy_binding := authenticationvserver_auditsyslogpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Authenticationvserver_auditsyslogpolicy_binding.Type(), &authenticationvserver_auditsyslogpolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationvserver_auditsyslogpolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("authenticationvserver_auditsyslogpolicy_binding-config")

	tflog.Trace(ctx, "Created authenticationvserver_auditsyslogpolicy_binding resource")

	// Read the updated state back
	r.readAuthenticationvserverAuditsyslogpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationvserverAuditsyslogpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationvserverAuditsyslogpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationvserver_auditsyslogpolicy_binding resource")

	r.readAuthenticationvserverAuditsyslogpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationvserverAuditsyslogpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AuthenticationvserverAuditsyslogpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating authenticationvserver_auditsyslogpolicy_binding resource")

	// Create API request body from the model
	// authenticationvserver_auditsyslogpolicy_binding := authenticationvserver_auditsyslogpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Authenticationvserver_auditsyslogpolicy_binding.Type(), &authenticationvserver_auditsyslogpolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authenticationvserver_auditsyslogpolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated authenticationvserver_auditsyslogpolicy_binding resource")

	// Read the updated state back
	r.readAuthenticationvserverAuditsyslogpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationvserverAuditsyslogpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationvserverAuditsyslogpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationvserver_auditsyslogpolicy_binding resource")

	// For authenticationvserver_auditsyslogpolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted authenticationvserver_auditsyslogpolicy_binding resource from state")
}

// Helper function to read authenticationvserver_auditsyslogpolicy_binding data from API
func (r *AuthenticationvserverAuditsyslogpolicyBindingResource) readAuthenticationvserverAuditsyslogpolicyBindingFromApi(ctx context.Context, data *AuthenticationvserverAuditsyslogpolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Authenticationvserver_auditsyslogpolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationvserver_auditsyslogpolicy_binding, got error: %s", err))
		return
	}

	authenticationvserver_auditsyslogpolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
