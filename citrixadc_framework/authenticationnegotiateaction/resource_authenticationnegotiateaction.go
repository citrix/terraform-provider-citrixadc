package authenticationnegotiateaction

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
var _ resource.Resource = &AuthenticationnegotiateactionResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationnegotiateactionResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationnegotiateactionResource)(nil)

func NewAuthenticationnegotiateactionResource() resource.Resource {
	return &AuthenticationnegotiateactionResource{}
}

// AuthenticationnegotiateactionResource defines the resource implementation.
type AuthenticationnegotiateactionResource struct {
	client *service.NitroClient
}

func (r *AuthenticationnegotiateactionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationnegotiateactionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationnegotiateaction"
}

func (r *AuthenticationnegotiateactionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationnegotiateactionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuthenticationnegotiateactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationnegotiateaction resource")

	// authenticationnegotiateaction := authenticationnegotiateactionGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Authenticationnegotiateaction.Type(), &authenticationnegotiateaction)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationnegotiateaction, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("authenticationnegotiateaction-config")

	tflog.Trace(ctx, "Created authenticationnegotiateaction resource")

	// Read the updated state back
	r.readAuthenticationnegotiateactionFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationnegotiateactionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationnegotiateactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationnegotiateaction resource")

	r.readAuthenticationnegotiateactionFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationnegotiateactionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AuthenticationnegotiateactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating authenticationnegotiateaction resource")

	// Create API request body from the model
	// authenticationnegotiateaction := authenticationnegotiateactionGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Authenticationnegotiateaction.Type(), &authenticationnegotiateaction)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authenticationnegotiateaction, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated authenticationnegotiateaction resource")

	// Read the updated state back
	r.readAuthenticationnegotiateactionFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationnegotiateactionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationnegotiateactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationnegotiateaction resource")

	// For authenticationnegotiateaction, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted authenticationnegotiateaction resource from state")
}

// Helper function to read authenticationnegotiateaction data from API
func (r *AuthenticationnegotiateactionResource) readAuthenticationnegotiateactionFromApi(ctx context.Context, data *AuthenticationnegotiateactionResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Authenticationnegotiateaction.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationnegotiateaction, got error: %s", err))
		return
	}

	authenticationnegotiateactionSetAttrFromGet(ctx, data, getResponseData)

}
