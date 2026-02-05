package authenticationoauthidpprofile

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
var _ resource.Resource = &AuthenticationoauthidpprofileResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationoauthidpprofileResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationoauthidpprofileResource)(nil)

func NewAuthenticationoauthidpprofileResource() resource.Resource {
	return &AuthenticationoauthidpprofileResource{}
}

// AuthenticationoauthidpprofileResource defines the resource implementation.
type AuthenticationoauthidpprofileResource struct {
	client *service.NitroClient
}

func (r *AuthenticationoauthidpprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationoauthidpprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationoauthidpprofile"
}

func (r *AuthenticationoauthidpprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationoauthidpprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuthenticationoauthidpprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationoauthidpprofile resource")

	// authenticationoauthidpprofile := authenticationoauthidpprofileGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Authenticationoauthidpprofile.Type(), &authenticationoauthidpprofile)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationoauthidpprofile, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("authenticationoauthidpprofile-config")

	tflog.Trace(ctx, "Created authenticationoauthidpprofile resource")

	// Read the updated state back
	r.readAuthenticationoauthidpprofileFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationoauthidpprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationoauthidpprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationoauthidpprofile resource")

	r.readAuthenticationoauthidpprofileFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationoauthidpprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AuthenticationoauthidpprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating authenticationoauthidpprofile resource")

	// Create API request body from the model
	// authenticationoauthidpprofile := authenticationoauthidpprofileGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Authenticationoauthidpprofile.Type(), &authenticationoauthidpprofile)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authenticationoauthidpprofile, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated authenticationoauthidpprofile resource")

	// Read the updated state back
	r.readAuthenticationoauthidpprofileFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationoauthidpprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationoauthidpprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationoauthidpprofile resource")

	// For authenticationoauthidpprofile, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted authenticationoauthidpprofile resource from state")
}

// Helper function to read authenticationoauthidpprofile data from API
func (r *AuthenticationoauthidpprofileResource) readAuthenticationoauthidpprofileFromApi(ctx context.Context, data *AuthenticationoauthidpprofileResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Authenticationoauthidpprofile.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationoauthidpprofile, got error: %s", err))
		return
	}

	authenticationoauthidpprofileSetAttrFromGet(ctx, data, getResponseData)

}
