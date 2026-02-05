package authenticationwebauthaction

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
var _ resource.Resource = &AuthenticationwebauthactionResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationwebauthactionResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationwebauthactionResource)(nil)

func NewAuthenticationwebauthactionResource() resource.Resource {
	return &AuthenticationwebauthactionResource{}
}

// AuthenticationwebauthactionResource defines the resource implementation.
type AuthenticationwebauthactionResource struct {
	client *service.NitroClient
}

func (r *AuthenticationwebauthactionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationwebauthactionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationwebauthaction"
}

func (r *AuthenticationwebauthactionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationwebauthactionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuthenticationwebauthactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationwebauthaction resource")

	// authenticationwebauthaction := authenticationwebauthactionGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Authenticationwebauthaction.Type(), &authenticationwebauthaction)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationwebauthaction, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("authenticationwebauthaction-config")

	tflog.Trace(ctx, "Created authenticationwebauthaction resource")

	// Read the updated state back
	r.readAuthenticationwebauthactionFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationwebauthactionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationwebauthactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationwebauthaction resource")

	r.readAuthenticationwebauthactionFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationwebauthactionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AuthenticationwebauthactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating authenticationwebauthaction resource")

	// Create API request body from the model
	// authenticationwebauthaction := authenticationwebauthactionGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Authenticationwebauthaction.Type(), &authenticationwebauthaction)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authenticationwebauthaction, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated authenticationwebauthaction resource")

	// Read the updated state back
	r.readAuthenticationwebauthactionFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationwebauthactionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationwebauthactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationwebauthaction resource")

	// For authenticationwebauthaction, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted authenticationwebauthaction resource from state")
}

// Helper function to read authenticationwebauthaction data from API
func (r *AuthenticationwebauthactionResource) readAuthenticationwebauthactionFromApi(ctx context.Context, data *AuthenticationwebauthactionResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Authenticationwebauthaction.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationwebauthaction, got error: %s", err))
		return
	}

	authenticationwebauthactionSetAttrFromGet(ctx, data, getResponseData)

}
