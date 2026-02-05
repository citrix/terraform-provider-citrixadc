package authenticationoauthidppolicy

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
var _ resource.Resource = &AuthenticationoauthidppolicyResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationoauthidppolicyResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationoauthidppolicyResource)(nil)

func NewAuthenticationoauthidppolicyResource() resource.Resource {
	return &AuthenticationoauthidppolicyResource{}
}

// AuthenticationoauthidppolicyResource defines the resource implementation.
type AuthenticationoauthidppolicyResource struct {
	client *service.NitroClient
}

func (r *AuthenticationoauthidppolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationoauthidppolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationoauthidppolicy"
}

func (r *AuthenticationoauthidppolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationoauthidppolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuthenticationoauthidppolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationoauthidppolicy resource")

	// authenticationoauthidppolicy := authenticationoauthidppolicyGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Authenticationoauthidppolicy.Type(), &authenticationoauthidppolicy)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationoauthidppolicy, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("authenticationoauthidppolicy-config")

	tflog.Trace(ctx, "Created authenticationoauthidppolicy resource")

	// Read the updated state back
	r.readAuthenticationoauthidppolicyFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationoauthidppolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationoauthidppolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationoauthidppolicy resource")

	r.readAuthenticationoauthidppolicyFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationoauthidppolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AuthenticationoauthidppolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating authenticationoauthidppolicy resource")

	// Create API request body from the model
	// authenticationoauthidppolicy := authenticationoauthidppolicyGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Authenticationoauthidppolicy.Type(), &authenticationoauthidppolicy)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authenticationoauthidppolicy, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated authenticationoauthidppolicy resource")

	// Read the updated state back
	r.readAuthenticationoauthidppolicyFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationoauthidppolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationoauthidppolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationoauthidppolicy resource")

	// For authenticationoauthidppolicy, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted authenticationoauthidppolicy resource from state")
}

// Helper function to read authenticationoauthidppolicy data from API
func (r *AuthenticationoauthidppolicyResource) readAuthenticationoauthidppolicyFromApi(ctx context.Context, data *AuthenticationoauthidppolicyResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Authenticationoauthidppolicy.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationoauthidppolicy, got error: %s", err))
		return
	}

	authenticationoauthidppolicySetAttrFromGet(ctx, data, getResponseData)

}
