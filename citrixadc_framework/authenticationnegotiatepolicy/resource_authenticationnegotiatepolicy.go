package authenticationnegotiatepolicy

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
var _ resource.Resource = &AuthenticationnegotiatepolicyResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationnegotiatepolicyResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationnegotiatepolicyResource)(nil)

func NewAuthenticationnegotiatepolicyResource() resource.Resource {
	return &AuthenticationnegotiatepolicyResource{}
}

// AuthenticationnegotiatepolicyResource defines the resource implementation.
type AuthenticationnegotiatepolicyResource struct {
	client *service.NitroClient
}

func (r *AuthenticationnegotiatepolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationnegotiatepolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationnegotiatepolicy"
}

func (r *AuthenticationnegotiatepolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationnegotiatepolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuthenticationnegotiatepolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationnegotiatepolicy resource")

	// authenticationnegotiatepolicy := authenticationnegotiatepolicyGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Authenticationnegotiatepolicy.Type(), &authenticationnegotiatepolicy)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationnegotiatepolicy, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("authenticationnegotiatepolicy-config")

	tflog.Trace(ctx, "Created authenticationnegotiatepolicy resource")

	// Read the updated state back
	r.readAuthenticationnegotiatepolicyFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationnegotiatepolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationnegotiatepolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationnegotiatepolicy resource")

	r.readAuthenticationnegotiatepolicyFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationnegotiatepolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AuthenticationnegotiatepolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating authenticationnegotiatepolicy resource")

	// Create API request body from the model
	// authenticationnegotiatepolicy := authenticationnegotiatepolicyGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Authenticationnegotiatepolicy.Type(), &authenticationnegotiatepolicy)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authenticationnegotiatepolicy, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated authenticationnegotiatepolicy resource")

	// Read the updated state back
	r.readAuthenticationnegotiatepolicyFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationnegotiatepolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationnegotiatepolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationnegotiatepolicy resource")

	// For authenticationnegotiatepolicy, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted authenticationnegotiatepolicy resource from state")
}

// Helper function to read authenticationnegotiatepolicy data from API
func (r *AuthenticationnegotiatepolicyResource) readAuthenticationnegotiatepolicyFromApi(ctx context.Context, data *AuthenticationnegotiatepolicyResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Authenticationnegotiatepolicy.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationnegotiatepolicy, got error: %s", err))
		return
	}

	authenticationnegotiatepolicySetAttrFromGet(ctx, data, getResponseData)

}
