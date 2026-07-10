package authenticationoauthidpprofile

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
var _ resource.Resource = &AuthenticationoauthidpprofileResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationoauthidpprofileResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationoauthidpprofileResource)(nil)

func NewAuthenticationoauthidpprofileResource() resource.Resource {
	return &AuthenticationoauthidpprofileResource{}
}

// AuthenticationoauthidpprofileResource defines the resource implementation.
type AuthenticationoauthidpprofileResource struct {
	client *service.NitroClient
}

func (r *AuthenticationoauthidpprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationoauthidpprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationoauthidpprofile"
}

func (r *AuthenticationoauthidpprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationoauthidpprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config AuthenticationoauthidpprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationoauthidpprofile resource")
	// Get payload from plan (regular attributes)
	authenticationoauthidpprofile := authenticationoauthidpprofileGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	authenticationoauthidpprofileGetThePayloadFromtheConfig(ctx, &config, &authenticationoauthidpprofile)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Authenticationoauthidpprofile.Type(), name_value, &authenticationoauthidpprofile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationoauthidpprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created authenticationoauthidpprofile resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	if !r.readAuthenticationoauthidpprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "authenticationoauthidpprofile not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationoauthidpprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationoauthidpprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationoauthidpprofile resource")

	found := r.readAuthenticationoauthidpprofileFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AuthenticationoauthidpprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state AuthenticationoauthidpprofileResourceModel

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

	tflog.Debug(ctx, "Updating authenticationoauthidpprofile resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Attributes.Equal(state.Attributes) {
		tflog.Debug(ctx, fmt.Sprintf("attributes has changed for authenticationoauthidpprofile"))
		hasChange = true
	}
	if !data.Audience.Equal(state.Audience) {
		tflog.Debug(ctx, fmt.Sprintf("audience has changed for authenticationoauthidpprofile"))
		hasChange = true
	}
	if !data.Clientid.Equal(state.Clientid) {
		tflog.Debug(ctx, fmt.Sprintf("clientid has changed for authenticationoauthidpprofile"))
		hasChange = true
	}
	// Check secret attribute clientsecret or its version tracker
	if !data.Clientsecret.Equal(state.Clientsecret) {
		tflog.Debug(ctx, fmt.Sprintf("clientsecret has changed for authenticationoauthidpprofile"))
		hasChange = true
	} else if !data.ClientsecretWoVersion.Equal(state.ClientsecretWoVersion) {
		tflog.Debug(ctx, fmt.Sprintf("clientsecret_wo_version has changed for authenticationoauthidpprofile"))
		hasChange = true
	}
	if !data.Configservice.Equal(state.Configservice) {
		tflog.Debug(ctx, fmt.Sprintf("configservice has changed for authenticationoauthidpprofile"))
		hasChange = true
	}
	if !data.Defaultauthenticationgroup.Equal(state.Defaultauthenticationgroup) {
		tflog.Debug(ctx, fmt.Sprintf("defaultauthenticationgroup has changed for authenticationoauthidpprofile"))
		hasChange = true
	}
	if !data.Encrypttoken.Equal(state.Encrypttoken) {
		tflog.Debug(ctx, fmt.Sprintf("encrypttoken has changed for authenticationoauthidpprofile"))
		hasChange = true
	}
	if !data.Issuer.Equal(state.Issuer) {
		tflog.Debug(ctx, fmt.Sprintf("issuer has changed for authenticationoauthidpprofile"))
		hasChange = true
	}
	if !data.Redirecturl.Equal(state.Redirecturl) {
		tflog.Debug(ctx, fmt.Sprintf("redirecturl has changed for authenticationoauthidpprofile"))
		hasChange = true
	}
	if !data.Refreshinterval.Equal(state.Refreshinterval) {
		tflog.Debug(ctx, fmt.Sprintf("refreshinterval has changed for authenticationoauthidpprofile"))
		hasChange = true
	}
	if !data.Relyingpartymetadataurl.Equal(state.Relyingpartymetadataurl) {
		tflog.Debug(ctx, fmt.Sprintf("relyingpartymetadataurl has changed for authenticationoauthidpprofile"))
		hasChange = true
	}
	if !data.Sendpassword.Equal(state.Sendpassword) {
		tflog.Debug(ctx, fmt.Sprintf("sendpassword has changed for authenticationoauthidpprofile"))
		hasChange = true
	}
	if !data.Signaturealg.Equal(state.Signaturealg) {
		tflog.Debug(ctx, fmt.Sprintf("signaturealg has changed for authenticationoauthidpprofile"))
		hasChange = true
	}
	if !data.Signatureservice.Equal(state.Signatureservice) {
		tflog.Debug(ctx, fmt.Sprintf("signatureservice has changed for authenticationoauthidpprofile"))
		hasChange = true
	}
	if !data.Skewtime.Equal(state.Skewtime) {
		tflog.Debug(ctx, fmt.Sprintf("skewtime has changed for authenticationoauthidpprofile"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		// Get payload from plan (regular attributes)
		authenticationoauthidpprofile := authenticationoauthidpprofileGetThePayloadFromthePlan(ctx, &data)
		// Add write-only attributes from config to the payload
		authenticationoauthidpprofileGetThePayloadFromtheConfig(ctx, &config, &authenticationoauthidpprofile)
		// Make API call
		// Named resource - use UpdateResource
		name_value := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Authenticationoauthidpprofile.Type(), name_value, &authenticationoauthidpprofile)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authenticationoauthidpprofile, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated authenticationoauthidpprofile resource")
	} else {
		tflog.Debug(ctx, "No changes detected for authenticationoauthidpprofile resource, skipping update")
	}

	// Read the updated state back
	if !r.readAuthenticationoauthidpprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "authenticationoauthidpprofile not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationoauthidpprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationoauthidpprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationoauthidpprofile resource")
	// Named resource - delete using DeleteResource
	name_value := data.Name.ValueString()
	err := r.client.DeleteResource(service.Authenticationoauthidpprofile.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete authenticationoauthidpprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted authenticationoauthidpprofile resource")
}

// Helper function to read authenticationoauthidpprofile data from API
func (r *AuthenticationoauthidpprofileResource) readAuthenticationoauthidpprofileFromApi(ctx context.Context, data *AuthenticationoauthidpprofileResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Authenticationoauthidpprofile.Type(), name_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationoauthidpprofile, got error: %s", err))
		return false
	}

	authenticationoauthidpprofileSetAttrFromGet(ctx, data, getResponseData)

	return true
}
