package authenticationradiuspolicy

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
var _ resource.Resource = &AuthenticationradiuspolicyResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationradiuspolicyResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationradiuspolicyResource)(nil)

func NewAuthenticationradiuspolicyResource() resource.Resource {
	return &AuthenticationradiuspolicyResource{}
}

// AuthenticationradiuspolicyResource defines the resource implementation.
type AuthenticationradiuspolicyResource struct {
	client *service.NitroClient
}

func (r *AuthenticationradiuspolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationradiuspolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationradiuspolicy"
}

func (r *AuthenticationradiuspolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationradiuspolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuthenticationradiuspolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationradiuspolicy resource")

	// authenticationradiuspolicy := authenticationradiuspolicyGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Authenticationradiuspolicy.Type(), &authenticationradiuspolicy)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationradiuspolicy, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("authenticationradiuspolicy-config")

	tflog.Trace(ctx, "Created authenticationradiuspolicy resource")

	// Read the updated state back
	r.readAuthenticationradiuspolicyFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationradiuspolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationradiuspolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationradiuspolicy resource")

	r.readAuthenticationradiuspolicyFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationradiuspolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AuthenticationradiuspolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating authenticationradiuspolicy resource")

	// Create API request body from the model
	// authenticationradiuspolicy := authenticationradiuspolicyGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Authenticationradiuspolicy.Type(), &authenticationradiuspolicy)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authenticationradiuspolicy, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated authenticationradiuspolicy resource")

	// Read the updated state back
	r.readAuthenticationradiuspolicyFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationradiuspolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationradiuspolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationradiuspolicy resource")

	// For authenticationradiuspolicy, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted authenticationradiuspolicy resource from state")
}

// Helper function to read authenticationradiuspolicy data from API
func (r *AuthenticationradiuspolicyResource) readAuthenticationradiuspolicyFromApi(ctx context.Context, data *AuthenticationradiuspolicyResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Authenticationradiuspolicy.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationradiuspolicy, got error: %s", err))
		return
	}

	authenticationradiuspolicySetAttrFromGet(ctx, data, getResponseData)

}
