package authenticationazurekeyvault

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AuthenticationazurekeyvaultResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationazurekeyvaultResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationazurekeyvaultResource)(nil)
var _ resource.ResourceWithValidateConfig = (*AuthenticationazurekeyvaultResource)(nil)

func NewAuthenticationazurekeyvaultResource() resource.Resource {
	return &AuthenticationazurekeyvaultResource{}
}

// AuthenticationazurekeyvaultResource defines the resource implementation.
type AuthenticationazurekeyvaultResource struct {
	client *service.NitroClient
}

func (r *AuthenticationazurekeyvaultResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationazurekeyvaultResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationazurekeyvault"
}

func (r *AuthenticationazurekeyvaultResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationazurekeyvaultResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data AuthenticationazurekeyvaultResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Validate that either clientsecret or clientsecret_wo is specified
	if data.Clientsecret.IsNull() && data.ClientsecretWo.IsNull() {
		resp.Diagnostics.AddAttributeError(
			path.Root("clientsecret"),
			"Missing Required Attribute",
			"Either \"clientsecret\" or \"clientsecret_wo\" must be specified.",
		)
	}
}

func (r *AuthenticationazurekeyvaultResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config AuthenticationazurekeyvaultResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationazurekeyvault resource")
	// Get payload from plan (regular attributes)
	authenticationazurekeyvault := authenticationazurekeyvaultGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	authenticationazurekeyvaultGetThePayloadFromtheConfig(ctx, &config, &authenticationazurekeyvault)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Authenticationazurekeyvault.Type(), name_value, &authenticationazurekeyvault)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationazurekeyvault, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created authenticationazurekeyvault resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	if !r.readAuthenticationazurekeyvaultFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "authenticationazurekeyvault not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationazurekeyvaultResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationazurekeyvaultResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationazurekeyvault resource")

	found := r.readAuthenticationazurekeyvaultFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if !found {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationazurekeyvaultResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state AuthenticationazurekeyvaultResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating authenticationazurekeyvault resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Authentication.Equal(state.Authentication) {
		tflog.Debug(ctx, fmt.Sprintf("authentication has changed for authenticationazurekeyvault"))
		hasChange = true
	}
	if !data.Clientid.Equal(state.Clientid) {
		tflog.Debug(ctx, fmt.Sprintf("clientid has changed for authenticationazurekeyvault"))
		hasChange = true
	}
	// Check secret attribute clientsecret or its version tracker
	if !data.Clientsecret.Equal(state.Clientsecret) {
		tflog.Debug(ctx, fmt.Sprintf("clientsecret has changed for authenticationazurekeyvault"))
		hasChange = true
	} else if !data.ClientsecretWoVersion.Equal(state.ClientsecretWoVersion) {
		tflog.Debug(ctx, fmt.Sprintf("clientsecret_wo_version has changed for authenticationazurekeyvault"))
		hasChange = true
	}
	if !data.Defaultauthenticationgroup.Equal(state.Defaultauthenticationgroup) {
		tflog.Debug(ctx, fmt.Sprintf("defaultauthenticationgroup has changed for authenticationazurekeyvault"))
		hasChange = true
	}
	if !data.Pushservice.Equal(state.Pushservice) {
		tflog.Debug(ctx, fmt.Sprintf("pushservice has changed for authenticationazurekeyvault"))
		hasChange = true
	}
	if !data.Refreshinterval.Equal(state.Refreshinterval) {
		tflog.Debug(ctx, fmt.Sprintf("refreshinterval has changed for authenticationazurekeyvault"))
		hasChange = true
	}
	if !data.Servicekeyname.Equal(state.Servicekeyname) {
		tflog.Debug(ctx, fmt.Sprintf("servicekeyname has changed for authenticationazurekeyvault"))
		hasChange = true
	}
	if !data.Signaturealg.Equal(state.Signaturealg) {
		tflog.Debug(ctx, fmt.Sprintf("signaturealg has changed for authenticationazurekeyvault"))
		hasChange = true
	}
	if !data.Tenantid.Equal(state.Tenantid) {
		tflog.Debug(ctx, fmt.Sprintf("tenantid has changed for authenticationazurekeyvault"))
		hasChange = true
	}
	if !data.Tokenendpoint.Equal(state.Tokenendpoint) {
		tflog.Debug(ctx, fmt.Sprintf("tokenendpoint has changed for authenticationazurekeyvault"))
		hasChange = true
	}
	if !data.Vaultname.Equal(state.Vaultname) {
		tflog.Debug(ctx, fmt.Sprintf("vaultname has changed for authenticationazurekeyvault"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		// Get payload from plan (regular attributes)
		authenticationazurekeyvault := authenticationazurekeyvaultGetThePayloadFromthePlan(ctx, &data)
		// Add write-only attributes from config to the payload
		authenticationazurekeyvaultGetThePayloadFromtheConfig(ctx, &config, &authenticationazurekeyvault)
		// Make API call
		// Named resource - use UpdateResource
		name_value := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Authenticationazurekeyvault.Type(), name_value, &authenticationazurekeyvault)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authenticationazurekeyvault, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated authenticationazurekeyvault resource")
	} else {
		tflog.Debug(ctx, "No changes detected for authenticationazurekeyvault resource, skipping update")
	}

	// Read the updated state back
	if !r.readAuthenticationazurekeyvaultFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "authenticationazurekeyvault not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationazurekeyvaultResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationazurekeyvaultResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationazurekeyvault resource")
	// Named resource - delete using DeleteResource
	name_value := data.Name.ValueString()
	err := r.client.DeleteResource(service.Authenticationazurekeyvault.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete authenticationazurekeyvault, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted authenticationazurekeyvault resource")
}

// Helper function to read authenticationazurekeyvault data from API
func (r *AuthenticationazurekeyvaultResource) readAuthenticationazurekeyvaultFromApi(ctx context.Context, data *AuthenticationazurekeyvaultResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Authenticationazurekeyvault.Type(), name_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationazurekeyvault, got error: %s", err))
		return false
	}

	authenticationazurekeyvaultSetAttrFromGet(ctx, data, getResponseData)

	return true
}
