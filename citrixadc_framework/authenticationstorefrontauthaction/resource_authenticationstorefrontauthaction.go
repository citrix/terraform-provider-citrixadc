package authenticationstorefrontauthaction

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
var _ resource.Resource = &AuthenticationstorefrontauthactionResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationstorefrontauthactionResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationstorefrontauthactionResource)(nil)

func NewAuthenticationstorefrontauthactionResource() resource.Resource {
	return &AuthenticationstorefrontauthactionResource{}
}

// AuthenticationstorefrontauthactionResource defines the resource implementation.
type AuthenticationstorefrontauthactionResource struct {
	client *service.NitroClient
}

func (r *AuthenticationstorefrontauthactionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationstorefrontauthactionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationstorefrontauthaction"
}

func (r *AuthenticationstorefrontauthactionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationstorefrontauthactionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuthenticationstorefrontauthactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationstorefrontauthaction resource")

	// authenticationstorefrontauthaction := authenticationstorefrontauthactionGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Authenticationstorefrontauthaction.Type(), &authenticationstorefrontauthaction)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationstorefrontauthaction, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("authenticationstorefrontauthaction-config")

	tflog.Trace(ctx, "Created authenticationstorefrontauthaction resource")

	// Read the updated state back
	r.readAuthenticationstorefrontauthactionFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationstorefrontauthactionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationstorefrontauthactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationstorefrontauthaction resource")

	r.readAuthenticationstorefrontauthactionFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationstorefrontauthactionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AuthenticationstorefrontauthactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating authenticationstorefrontauthaction resource")

	// Create API request body from the model
	// authenticationstorefrontauthaction := authenticationstorefrontauthactionGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Authenticationstorefrontauthaction.Type(), &authenticationstorefrontauthaction)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authenticationstorefrontauthaction, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated authenticationstorefrontauthaction resource")

	// Read the updated state back
	r.readAuthenticationstorefrontauthactionFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationstorefrontauthactionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationstorefrontauthactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationstorefrontauthaction resource")

	// For authenticationstorefrontauthaction, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted authenticationstorefrontauthaction resource from state")
}

// Helper function to read authenticationstorefrontauthaction data from API
func (r *AuthenticationstorefrontauthactionResource) readAuthenticationstorefrontauthactionFromApi(ctx context.Context, data *AuthenticationstorefrontauthactionResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Authenticationstorefrontauthaction.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationstorefrontauthaction, got error: %s", err))
		return
	}

	authenticationstorefrontauthactionSetAttrFromGet(ctx, data, getResponseData)

}
