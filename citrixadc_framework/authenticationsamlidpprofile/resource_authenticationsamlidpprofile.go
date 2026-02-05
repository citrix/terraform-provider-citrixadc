package authenticationsamlidpprofile

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
var _ resource.Resource = &AuthenticationsamlidpprofileResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationsamlidpprofileResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationsamlidpprofileResource)(nil)

func NewAuthenticationsamlidpprofileResource() resource.Resource {
	return &AuthenticationsamlidpprofileResource{}
}

// AuthenticationsamlidpprofileResource defines the resource implementation.
type AuthenticationsamlidpprofileResource struct {
	client *service.NitroClient
}

func (r *AuthenticationsamlidpprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationsamlidpprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationsamlidpprofile"
}

func (r *AuthenticationsamlidpprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationsamlidpprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuthenticationsamlidpprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationsamlidpprofile resource")

	// authenticationsamlidpprofile := authenticationsamlidpprofileGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Authenticationsamlidpprofile.Type(), &authenticationsamlidpprofile)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationsamlidpprofile, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("authenticationsamlidpprofile-config")

	tflog.Trace(ctx, "Created authenticationsamlidpprofile resource")

	// Read the updated state back
	r.readAuthenticationsamlidpprofileFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationsamlidpprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationsamlidpprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationsamlidpprofile resource")

	r.readAuthenticationsamlidpprofileFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationsamlidpprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AuthenticationsamlidpprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating authenticationsamlidpprofile resource")

	// Create API request body from the model
	// authenticationsamlidpprofile := authenticationsamlidpprofileGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Authenticationsamlidpprofile.Type(), &authenticationsamlidpprofile)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authenticationsamlidpprofile, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated authenticationsamlidpprofile resource")

	// Read the updated state back
	r.readAuthenticationsamlidpprofileFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationsamlidpprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationsamlidpprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationsamlidpprofile resource")

	// For authenticationsamlidpprofile, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted authenticationsamlidpprofile resource from state")
}

// Helper function to read authenticationsamlidpprofile data from API
func (r *AuthenticationsamlidpprofileResource) readAuthenticationsamlidpprofileFromApi(ctx context.Context, data *AuthenticationsamlidpprofileResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Authenticationsamlidpprofile.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationsamlidpprofile, got error: %s", err))
		return
	}

	authenticationsamlidpprofileSetAttrFromGet(ctx, data, getResponseData)

}
