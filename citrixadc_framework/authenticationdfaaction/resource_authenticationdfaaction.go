package authenticationdfaaction

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
var _ resource.Resource = &AuthenticationdfaactionResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationdfaactionResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationdfaactionResource)(nil)

func NewAuthenticationdfaactionResource() resource.Resource {
	return &AuthenticationdfaactionResource{}
}

// AuthenticationdfaactionResource defines the resource implementation.
type AuthenticationdfaactionResource struct {
	client *service.NitroClient
}

func (r *AuthenticationdfaactionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationdfaactionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationdfaaction"
}

func (r *AuthenticationdfaactionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationdfaactionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuthenticationdfaactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationdfaaction resource")

	// authenticationdfaaction := authenticationdfaactionGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Authenticationdfaaction.Type(), &authenticationdfaaction)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationdfaaction, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("authenticationdfaaction-config")

	tflog.Trace(ctx, "Created authenticationdfaaction resource")

	// Read the updated state back
	r.readAuthenticationdfaactionFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationdfaactionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationdfaactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationdfaaction resource")

	r.readAuthenticationdfaactionFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationdfaactionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AuthenticationdfaactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating authenticationdfaaction resource")

	// Create API request body from the model
	// authenticationdfaaction := authenticationdfaactionGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Authenticationdfaaction.Type(), &authenticationdfaaction)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authenticationdfaaction, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated authenticationdfaaction resource")

	// Read the updated state back
	r.readAuthenticationdfaactionFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationdfaactionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationdfaactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationdfaaction resource")

	// For authenticationdfaaction, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted authenticationdfaaction resource from state")
}

// Helper function to read authenticationdfaaction data from API
func (r *AuthenticationdfaactionResource) readAuthenticationdfaactionFromApi(ctx context.Context, data *AuthenticationdfaactionResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Authenticationdfaaction.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationdfaaction, got error: %s", err))
		return
	}

	authenticationdfaactionSetAttrFromGet(ctx, data, getResponseData)

}
