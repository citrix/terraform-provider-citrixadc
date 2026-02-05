package authenticationcitrixauthaction

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
var _ resource.Resource = &AuthenticationcitrixauthactionResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationcitrixauthactionResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationcitrixauthactionResource)(nil)

func NewAuthenticationcitrixauthactionResource() resource.Resource {
	return &AuthenticationcitrixauthactionResource{}
}

// AuthenticationcitrixauthactionResource defines the resource implementation.
type AuthenticationcitrixauthactionResource struct {
	client *service.NitroClient
}

func (r *AuthenticationcitrixauthactionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationcitrixauthactionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationcitrixauthaction"
}

func (r *AuthenticationcitrixauthactionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationcitrixauthactionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuthenticationcitrixauthactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationcitrixauthaction resource")

	// authenticationcitrixauthaction := authenticationcitrixauthactionGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Authenticationcitrixauthaction.Type(), &authenticationcitrixauthaction)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationcitrixauthaction, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("authenticationcitrixauthaction-config")

	tflog.Trace(ctx, "Created authenticationcitrixauthaction resource")

	// Read the updated state back
	r.readAuthenticationcitrixauthactionFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationcitrixauthactionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationcitrixauthactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationcitrixauthaction resource")

	r.readAuthenticationcitrixauthactionFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationcitrixauthactionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AuthenticationcitrixauthactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating authenticationcitrixauthaction resource")

	// Create API request body from the model
	// authenticationcitrixauthaction := authenticationcitrixauthactionGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Authenticationcitrixauthaction.Type(), &authenticationcitrixauthaction)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authenticationcitrixauthaction, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated authenticationcitrixauthaction resource")

	// Read the updated state back
	r.readAuthenticationcitrixauthactionFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationcitrixauthactionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationcitrixauthactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationcitrixauthaction resource")

	// For authenticationcitrixauthaction, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted authenticationcitrixauthaction resource from state")
}

// Helper function to read authenticationcitrixauthaction data from API
func (r *AuthenticationcitrixauthactionResource) readAuthenticationcitrixauthactionFromApi(ctx context.Context, data *AuthenticationcitrixauthactionResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Authenticationcitrixauthaction.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationcitrixauthaction, got error: %s", err))
		return
	}

	authenticationcitrixauthactionSetAttrFromGet(ctx, data, getResponseData)

}
