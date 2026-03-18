package authenticationloginschemapolicy

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
var _ resource.Resource = &AuthenticationloginschemapolicyResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationloginschemapolicyResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationloginschemapolicyResource)(nil)

func NewAuthenticationloginschemapolicyResource() resource.Resource {
	return &AuthenticationloginschemapolicyResource{}
}

// AuthenticationloginschemapolicyResource defines the resource implementation.
type AuthenticationloginschemapolicyResource struct {
	client *service.NitroClient
}

func (r *AuthenticationloginschemapolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationloginschemapolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationloginschemapolicy"
}

func (r *AuthenticationloginschemapolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationloginschemapolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuthenticationloginschemapolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationloginschemapolicy resource")

	// authenticationloginschemapolicy := authenticationloginschemapolicyGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Authenticationloginschemapolicy.Type(), &authenticationloginschemapolicy)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationloginschemapolicy, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("authenticationloginschemapolicy-config")

	tflog.Trace(ctx, "Created authenticationloginschemapolicy resource")

	// Read the updated state back
	r.readAuthenticationloginschemapolicyFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationloginschemapolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationloginschemapolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationloginschemapolicy resource")

	r.readAuthenticationloginschemapolicyFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationloginschemapolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AuthenticationloginschemapolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating authenticationloginschemapolicy resource")

	// Create API request body from the model
	// authenticationloginschemapolicy := authenticationloginschemapolicyGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Authenticationloginschemapolicy.Type(), &authenticationloginschemapolicy)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authenticationloginschemapolicy, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated authenticationloginschemapolicy resource")

	// Read the updated state back
	r.readAuthenticationloginschemapolicyFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationloginschemapolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationloginschemapolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationloginschemapolicy resource")

	// For authenticationloginschemapolicy, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted authenticationloginschemapolicy resource from state")
}

// Helper function to read authenticationloginschemapolicy data from API
func (r *AuthenticationloginschemapolicyResource) readAuthenticationloginschemapolicyFromApi(ctx context.Context, data *AuthenticationloginschemapolicyResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Authenticationloginschemapolicy.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationloginschemapolicy, got error: %s", err))
		return
	}

	authenticationloginschemapolicySetAttrFromGet(ctx, data, getResponseData)

}
