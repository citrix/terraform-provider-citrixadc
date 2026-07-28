package authenticationoauthaction

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
var _ resource.Resource = &AuthenticationoauthactionResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationoauthactionResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationoauthactionResource)(nil)

func NewAuthenticationoauthactionResource() resource.Resource {
	return &AuthenticationoauthactionResource{}
}

// AuthenticationoauthactionResource defines the resource implementation.
type AuthenticationoauthactionResource struct {
	client *service.NitroClient
}

func (r *AuthenticationoauthactionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationoauthactionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationoauthaction"
}

func (r *AuthenticationoauthactionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationoauthactionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config AuthenticationoauthactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationoauthaction resource")
	// Get payload from plan (regular attributes)
	authenticationoauthaction := authenticationoauthactionGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	authenticationoauthactionGetThePayloadFromtheConfig(ctx, &config, &authenticationoauthaction)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Authenticationoauthaction.Type(), name_value, &authenticationoauthaction)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationoauthaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created authenticationoauthaction resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	if !r.readAuthenticationoauthactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "authenticationoauthaction not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationoauthactionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationoauthactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationoauthaction resource")

	found := r.readAuthenticationoauthactionFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AuthenticationoauthactionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state AuthenticationoauthactionResourceModel

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

	tflog.Debug(ctx, "Updating authenticationoauthaction resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	attributesToUnset := []string{}
	if !data.Allowedalgorithms.Equal(state.Allowedalgorithms) {
		tflog.Debug(ctx, fmt.Sprintf("allowedalgorithms has changed for authenticationoauthaction"))
		hasChange = true
	}
	if !data.Attribute1.Equal(state.Attribute1) {
		tflog.Debug(ctx, fmt.Sprintf("attribute1 has changed for authenticationoauthaction"))
		hasChange = true
	}
	if !data.Attribute10.Equal(state.Attribute10) {
		tflog.Debug(ctx, fmt.Sprintf("attribute10 has changed for authenticationoauthaction"))
		hasChange = true
	}
	if !data.Attribute11.Equal(state.Attribute11) {
		tflog.Debug(ctx, fmt.Sprintf("attribute11 has changed for authenticationoauthaction"))
		hasChange = true
	}
	if !data.Attribute12.Equal(state.Attribute12) {
		tflog.Debug(ctx, fmt.Sprintf("attribute12 has changed for authenticationoauthaction"))
		hasChange = true
	}
	if !data.Attribute13.Equal(state.Attribute13) {
		tflog.Debug(ctx, fmt.Sprintf("attribute13 has changed for authenticationoauthaction"))
		hasChange = true
	}
	if !data.Attribute14.Equal(state.Attribute14) {
		tflog.Debug(ctx, fmt.Sprintf("attribute14 has changed for authenticationoauthaction"))
		hasChange = true
	}
	if !data.Attribute15.Equal(state.Attribute15) {
		tflog.Debug(ctx, fmt.Sprintf("attribute15 has changed for authenticationoauthaction"))
		hasChange = true
	}
	if !data.Attribute16.Equal(state.Attribute16) {
		tflog.Debug(ctx, fmt.Sprintf("attribute16 has changed for authenticationoauthaction"))
		hasChange = true
	}
	if !data.Attribute2.Equal(state.Attribute2) {
		tflog.Debug(ctx, fmt.Sprintf("attribute2 has changed for authenticationoauthaction"))
		hasChange = true
	}
	if !data.Attribute3.Equal(state.Attribute3) {
		tflog.Debug(ctx, fmt.Sprintf("attribute3 has changed for authenticationoauthaction"))
		hasChange = true
	}
	if !data.Attribute4.Equal(state.Attribute4) {
		tflog.Debug(ctx, fmt.Sprintf("attribute4 has changed for authenticationoauthaction"))
		hasChange = true
	}
	if !data.Attribute5.Equal(state.Attribute5) {
		tflog.Debug(ctx, fmt.Sprintf("attribute5 has changed for authenticationoauthaction"))
		hasChange = true
	}
	if !data.Attribute6.Equal(state.Attribute6) {
		tflog.Debug(ctx, fmt.Sprintf("attribute6 has changed for authenticationoauthaction"))
		hasChange = true
	}
	if !data.Attribute7.Equal(state.Attribute7) {
		tflog.Debug(ctx, fmt.Sprintf("attribute7 has changed for authenticationoauthaction"))
		hasChange = true
	}
	if !data.Attribute8.Equal(state.Attribute8) {
		tflog.Debug(ctx, fmt.Sprintf("attribute8 has changed for authenticationoauthaction"))
		hasChange = true
	}
	if !data.Attribute9.Equal(state.Attribute9) {
		tflog.Debug(ctx, fmt.Sprintf("attribute9 has changed for authenticationoauthaction"))
		hasChange = true
	}
	if !data.Attributes.Equal(state.Attributes) {
		tflog.Debug(ctx, fmt.Sprintf("attributes has changed for authenticationoauthaction"))
		hasChange = true
	}
	if !data.Audience.Equal(state.Audience) {
		tflog.Debug(ctx, fmt.Sprintf("audience has changed for authenticationoauthaction"))
		hasChange = true
	}
	if !data.Authentication.Equal(state.Authentication) {
		tflog.Debug(ctx, fmt.Sprintf("authentication has changed for authenticationoauthaction"))
		if config.Authentication.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "authentication")
		} else {
			hasChange = true
		}
	}
	if !data.Authorizationendpoint.Equal(state.Authorizationendpoint) {
		tflog.Debug(ctx, fmt.Sprintf("authorizationendpoint has changed for authenticationoauthaction"))
		hasChange = true
	}
	if !data.Certendpoint.Equal(state.Certendpoint) {
		tflog.Debug(ctx, fmt.Sprintf("certendpoint has changed for authenticationoauthaction"))
		hasChange = true
	}
	if !data.Certfilepath.Equal(state.Certfilepath) {
		tflog.Debug(ctx, fmt.Sprintf("certfilepath has changed for authenticationoauthaction"))
		hasChange = true
	}
	if !data.Clientid.Equal(state.Clientid) {
		tflog.Debug(ctx, fmt.Sprintf("clientid has changed for authenticationoauthaction"))
		hasChange = true
	}
	// Check secret attribute clientsecret or its version tracker
	if !data.Clientsecret.Equal(state.Clientsecret) {
		tflog.Debug(ctx, fmt.Sprintf("clientsecret has changed for authenticationoauthaction"))
		hasChange = true
	} else if !data.ClientsecretWoVersion.Equal(state.ClientsecretWoVersion) {
		tflog.Debug(ctx, fmt.Sprintf("clientsecret_wo_version has changed for authenticationoauthaction"))
		hasChange = true
	}
	if !data.Defaultauthenticationgroup.Equal(state.Defaultauthenticationgroup) {
		tflog.Debug(ctx, fmt.Sprintf("defaultauthenticationgroup has changed for authenticationoauthaction"))
		hasChange = true
	}
	if !data.Granttype.Equal(state.Granttype) {
		tflog.Debug(ctx, fmt.Sprintf("granttype has changed for authenticationoauthaction"))
		hasChange = true
	}
	if !data.Graphendpoint.Equal(state.Graphendpoint) {
		tflog.Debug(ctx, fmt.Sprintf("graphendpoint has changed for authenticationoauthaction"))
		hasChange = true
	}
	if !data.Idtokendecryptendpoint.Equal(state.Idtokendecryptendpoint) {
		tflog.Debug(ctx, fmt.Sprintf("idtokendecryptendpoint has changed for authenticationoauthaction"))
		hasChange = true
	}
	if !data.Introspecturl.Equal(state.Introspecturl) {
		tflog.Debug(ctx, fmt.Sprintf("introspecturl has changed for authenticationoauthaction"))
		hasChange = true
	}
	if !data.Intunedeviceidexpression.Equal(state.Intunedeviceidexpression) {
		tflog.Debug(ctx, fmt.Sprintf("intunedeviceidexpression has changed for authenticationoauthaction"))
		hasChange = true
	}
	if !data.Issuer.Equal(state.Issuer) {
		tflog.Debug(ctx, fmt.Sprintf("issuer has changed for authenticationoauthaction"))
		hasChange = true
	}
	if !data.Metadataurl.Equal(state.Metadataurl) {
		tflog.Debug(ctx, fmt.Sprintf("metadataurl has changed for authenticationoauthaction"))
		hasChange = true
	}
	if !data.Oauthmiscflags.Equal(state.Oauthmiscflags) {
		tflog.Debug(ctx, fmt.Sprintf("oauthmiscflags has changed for authenticationoauthaction"))
		hasChange = true
	}
	if !data.Oauthtype.Equal(state.Oauthtype) {
		tflog.Debug(ctx, fmt.Sprintf("oauthtype has changed for authenticationoauthaction"))
		if config.Oauthtype.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "oauthtype")
		} else {
			hasChange = true
		}
	}
	if !data.Pkce.Equal(state.Pkce) {
		tflog.Debug(ctx, fmt.Sprintf("pkce has changed for authenticationoauthaction"))
		if config.Pkce.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "pkce")
		} else {
			hasChange = true
		}
	}
	if !data.Refreshinterval.Equal(state.Refreshinterval) {
		tflog.Debug(ctx, fmt.Sprintf("refreshinterval has changed for authenticationoauthaction"))
		if config.Refreshinterval.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "refreshinterval")
		} else {
			hasChange = true
		}
	}
	if !data.Requestattribute.Equal(state.Requestattribute) {
		tflog.Debug(ctx, fmt.Sprintf("requestattribute has changed for authenticationoauthaction"))
		hasChange = true
	}
	if !data.Resourceuri.Equal(state.Resourceuri) {
		tflog.Debug(ctx, fmt.Sprintf("resourceuri has changed for authenticationoauthaction"))
		hasChange = true
	}
	if !data.Skewtime.Equal(state.Skewtime) {
		tflog.Debug(ctx, fmt.Sprintf("skewtime has changed for authenticationoauthaction"))
		if config.Skewtime.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "skewtime")
		} else {
			hasChange = true
		}
	}
	if !data.Tenantid.Equal(state.Tenantid) {
		tflog.Debug(ctx, fmt.Sprintf("tenantid has changed for authenticationoauthaction"))
		hasChange = true
	}
	if !data.Tokenendpoint.Equal(state.Tokenendpoint) {
		tflog.Debug(ctx, fmt.Sprintf("tokenendpoint has changed for authenticationoauthaction"))
		hasChange = true
	}
	if !data.Tokenendpointauthmethod.Equal(state.Tokenendpointauthmethod) {
		tflog.Debug(ctx, fmt.Sprintf("tokenendpointauthmethod has changed for authenticationoauthaction"))
		if config.Tokenendpointauthmethod.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "tokenendpointauthmethod")
		} else {
			hasChange = true
		}
	}
	if !data.Userinfourl.Equal(state.Userinfourl) {
		tflog.Debug(ctx, fmt.Sprintf("userinfourl has changed for authenticationoauthaction"))
		hasChange = true
	}
	if !data.Usernamefield.Equal(state.Usernamefield) {
		tflog.Debug(ctx, fmt.Sprintf("usernamefield has changed for authenticationoauthaction"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		// Get payload from plan (regular attributes)
		authenticationoauthaction := authenticationoauthactionGetThePayloadFromthePlan(ctx, &data)
		// Add write-only attributes from config to the payload
		authenticationoauthactionGetThePayloadFromtheConfig(ctx, &config, &authenticationoauthaction)
		// Make API call
		// Named resource - use UpdateResource
		name_value := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Authenticationoauthaction.Type(), name_value, &authenticationoauthaction)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authenticationoauthaction, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated authenticationoauthaction resource")
	} else {
		tflog.Debug(ctx, "No changes detected for authenticationoauthaction resource, skipping update")
	}

	// Issue a single batched unset for attributes removed from config so the
	// appliance reverts them to their defaults. Update-then-unset ordering
	// ensures any default the update payload carried is superseded.
	unsetIdPayload := map[string]interface{}{
		"name": data.Name.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Authenticationoauthaction.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset authenticationoauthaction attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readAuthenticationoauthactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "authenticationoauthaction not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationoauthactionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationoauthactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationoauthaction resource")
	// Named resource - delete using DeleteResource
	name_value := data.Name.ValueString()
	err := r.client.DeleteResource(service.Authenticationoauthaction.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete authenticationoauthaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted authenticationoauthaction resource")
}

// Helper function to read authenticationoauthaction data from API
func (r *AuthenticationoauthactionResource) readAuthenticationoauthactionFromApi(ctx context.Context, data *AuthenticationoauthactionResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Authenticationoauthaction.Type(), name_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationoauthaction, got error: %s", err))
		return false
	}

	authenticationoauthactionSetAttrFromGet(ctx, data, getResponseData)

	return true
}
