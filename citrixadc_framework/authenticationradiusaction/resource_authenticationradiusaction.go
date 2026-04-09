package authenticationradiusaction

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
var _ resource.Resource = &AuthenticationradiusactionResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationradiusactionResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationradiusactionResource)(nil)
var _ resource.ResourceWithValidateConfig = (*AuthenticationradiusactionResource)(nil)

func NewAuthenticationradiusactionResource() resource.Resource {
	return &AuthenticationradiusactionResource{}
}

// AuthenticationradiusactionResource defines the resource implementation.
type AuthenticationradiusactionResource struct {
	client *service.NitroClient
}

func (r *AuthenticationradiusactionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationradiusactionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationradiusaction"
}

func (r *AuthenticationradiusactionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationradiusactionResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data AuthenticationradiusactionResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Validate that either radkey or radkey_wo is specified
	if data.Radkey.IsNull() && data.RadkeyWo.IsNull() {
		resp.Diagnostics.AddAttributeError(
			path.Root("radkey"),
			"Missing Required Attribute",
			"Either \"radkey\" or \"radkey_wo\" must be specified.",
		)
	}
}

func (r *AuthenticationradiusactionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config AuthenticationradiusactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationradiusaction resource")
	// Get payload from plan (regular attributes)
	authenticationradiusaction := authenticationradiusactionGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	authenticationradiusactionGetThePayloadFromtheConfig(ctx, &config, &authenticationradiusaction)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Authenticationradiusaction.Type(), name_value, &authenticationradiusaction)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationradiusaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created authenticationradiusaction resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	r.readAuthenticationradiusactionFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationradiusactionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationradiusactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationradiusaction resource")

	r.readAuthenticationradiusactionFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationradiusactionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state AuthenticationradiusactionResourceModel

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

	tflog.Debug(ctx, "Updating authenticationradiusaction resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Accounting.Equal(state.Accounting) {
		tflog.Debug(ctx, fmt.Sprintf("accounting has changed for authenticationradiusaction"))
		hasChange = true
	}
	if !data.Authentication.Equal(state.Authentication) {
		tflog.Debug(ctx, fmt.Sprintf("authentication has changed for authenticationradiusaction"))
		hasChange = true
	}
	if !data.Authservretry.Equal(state.Authservretry) {
		tflog.Debug(ctx, fmt.Sprintf("authservretry has changed for authenticationradiusaction"))
		hasChange = true
	}
	if !data.Authtimeout.Equal(state.Authtimeout) {
		tflog.Debug(ctx, fmt.Sprintf("authtimeout has changed for authenticationradiusaction"))
		hasChange = true
	}
	if !data.Callingstationid.Equal(state.Callingstationid) {
		tflog.Debug(ctx, fmt.Sprintf("callingstationid has changed for authenticationradiusaction"))
		hasChange = true
	}
	if !data.Defaultauthenticationgroup.Equal(state.Defaultauthenticationgroup) {
		tflog.Debug(ctx, fmt.Sprintf("defaultauthenticationgroup has changed for authenticationradiusaction"))
		hasChange = true
	}
	if !data.Ipattributetype.Equal(state.Ipattributetype) {
		tflog.Debug(ctx, fmt.Sprintf("ipattributetype has changed for authenticationradiusaction"))
		hasChange = true
	}
	if !data.Ipvendorid.Equal(state.Ipvendorid) {
		tflog.Debug(ctx, fmt.Sprintf("ipvendorid has changed for authenticationradiusaction"))
		hasChange = true
	}
	if !data.Messageauthenticator.Equal(state.Messageauthenticator) {
		tflog.Debug(ctx, fmt.Sprintf("messageauthenticator has changed for authenticationradiusaction"))
		hasChange = true
	}
	if !data.Passencoding.Equal(state.Passencoding) {
		tflog.Debug(ctx, fmt.Sprintf("passencoding has changed for authenticationradiusaction"))
		hasChange = true
	}
	if !data.Pwdattributetype.Equal(state.Pwdattributetype) {
		tflog.Debug(ctx, fmt.Sprintf("pwdattributetype has changed for authenticationradiusaction"))
		hasChange = true
	}
	if !data.Pwdvendorid.Equal(state.Pwdvendorid) {
		tflog.Debug(ctx, fmt.Sprintf("pwdvendorid has changed for authenticationradiusaction"))
		hasChange = true
	}
	if !data.Radattributetype.Equal(state.Radattributetype) {
		tflog.Debug(ctx, fmt.Sprintf("radattributetype has changed for authenticationradiusaction"))
		hasChange = true
	}
	if !data.Radgroupseparator.Equal(state.Radgroupseparator) {
		tflog.Debug(ctx, fmt.Sprintf("radgroupseparator has changed for authenticationradiusaction"))
		hasChange = true
	}
	if !data.Radgroupsprefix.Equal(state.Radgroupsprefix) {
		tflog.Debug(ctx, fmt.Sprintf("radgroupsprefix has changed for authenticationradiusaction"))
		hasChange = true
	}
	// Check secret attribute radkey or its version tracker
	if !data.Radkey.Equal(state.Radkey) {
		tflog.Debug(ctx, fmt.Sprintf("radkey has changed for authenticationradiusaction"))
		hasChange = true
	} else if !data.RadkeyWoVersion.Equal(state.RadkeyWoVersion) {
		tflog.Debug(ctx, fmt.Sprintf("radkey_wo_version has changed for authenticationradiusaction"))
		hasChange = true
	}
	if !data.Radnasid.Equal(state.Radnasid) {
		tflog.Debug(ctx, fmt.Sprintf("radnasid has changed for authenticationradiusaction"))
		hasChange = true
	}
	if !data.Radnasip.Equal(state.Radnasip) {
		tflog.Debug(ctx, fmt.Sprintf("radnasip has changed for authenticationradiusaction"))
		hasChange = true
	}
	if !data.Radvendorid.Equal(state.Radvendorid) {
		tflog.Debug(ctx, fmt.Sprintf("radvendorid has changed for authenticationradiusaction"))
		hasChange = true
	}
	if !data.Serverip.Equal(state.Serverip) {
		tflog.Debug(ctx, fmt.Sprintf("serverip has changed for authenticationradiusaction"))
		hasChange = true
	}
	if !data.Servername.Equal(state.Servername) {
		tflog.Debug(ctx, fmt.Sprintf("servername has changed for authenticationradiusaction"))
		hasChange = true
	}
	if !data.Serverport.Equal(state.Serverport) {
		tflog.Debug(ctx, fmt.Sprintf("serverport has changed for authenticationradiusaction"))
		hasChange = true
	}
	if !data.Targetlbvserver.Equal(state.Targetlbvserver) {
		tflog.Debug(ctx, fmt.Sprintf("targetlbvserver has changed for authenticationradiusaction"))
		hasChange = true
	}
	if !data.Transport.Equal(state.Transport) {
		tflog.Debug(ctx, fmt.Sprintf("transport has changed for authenticationradiusaction"))
		hasChange = true
	}
	if !data.Tunnelendpointclientip.Equal(state.Tunnelendpointclientip) {
		tflog.Debug(ctx, fmt.Sprintf("tunnelendpointclientip has changed for authenticationradiusaction"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		// Get payload from plan (regular attributes)
		authenticationradiusaction := authenticationradiusactionGetThePayloadFromthePlan(ctx, &data)
		// Add write-only attributes from config to the payload
		authenticationradiusactionGetThePayloadFromtheConfig(ctx, &config, &authenticationradiusaction)
		// Make API call
		// Named resource - use UpdateResource
		name_value := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Authenticationradiusaction.Type(), name_value, &authenticationradiusaction)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authenticationradiusaction, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated authenticationradiusaction resource")
	} else {
		tflog.Debug(ctx, "No changes detected for authenticationradiusaction resource, skipping update")
	}

	// Read the updated state back
	r.readAuthenticationradiusactionFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationradiusactionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationradiusactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationradiusaction resource")
	// Named resource - delete using DeleteResource
	name_value := data.Name.ValueString()
	err := r.client.DeleteResource(service.Authenticationradiusaction.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete authenticationradiusaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted authenticationradiusaction resource")
}

// Helper function to read authenticationradiusaction data from API
func (r *AuthenticationradiusactionResource) readAuthenticationradiusactionFromApi(ctx context.Context, data *AuthenticationradiusactionResourceModel, diags *diag.Diagnostics) {

	// Case 2: Find with single ID attribute - ID is the plain value
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Authenticationradiusaction.Type(), name_Name)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationradiusaction, got error: %s", err))
		return
	}

	authenticationradiusactionSetAttrFromGet(ctx, data, getResponseData)

}
