package authenticationcaptchaaction

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
var _ resource.Resource = &AuthenticationcaptchaactionResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationcaptchaactionResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationcaptchaactionResource)(nil)
var _ resource.ResourceWithValidateConfig = (*AuthenticationcaptchaactionResource)(nil)

func NewAuthenticationcaptchaactionResource() resource.Resource {
	return &AuthenticationcaptchaactionResource{}
}

// AuthenticationcaptchaactionResource defines the resource implementation.
type AuthenticationcaptchaactionResource struct {
	client *service.NitroClient
}

func (r *AuthenticationcaptchaactionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationcaptchaactionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationcaptchaaction"
}

func (r *AuthenticationcaptchaactionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationcaptchaactionResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data AuthenticationcaptchaactionResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Validate that either secretkey or secretkey_wo is specified
	if data.Secretkey.IsNull() && data.SecretkeyWo.IsNull() {
		resp.Diagnostics.AddAttributeError(
			path.Root("secretkey"),
			"Missing Required Attribute",
			"Either \"secretkey\" or \"secretkey_wo\" must be specified.",
		)
	}

	// Validate that either sitekey or sitekey_wo is specified
	if data.Sitekey.IsNull() && data.SitekeyWo.IsNull() {
		resp.Diagnostics.AddAttributeError(
			path.Root("sitekey"),
			"Missing Required Attribute",
			"Either \"sitekey\" or \"sitekey_wo\" must be specified.",
		)
	}
}

func (r *AuthenticationcaptchaactionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config AuthenticationcaptchaactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationcaptchaaction resource")
	// Get payload from plan (regular attributes)
	authenticationcaptchaaction := authenticationcaptchaactionGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	authenticationcaptchaactionGetThePayloadFromtheConfig(ctx, &config, &authenticationcaptchaaction)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Authenticationcaptchaaction.Type(), name_value, &authenticationcaptchaaction)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationcaptchaaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created authenticationcaptchaaction resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	if !r.readAuthenticationcaptchaactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "authenticationcaptchaaction not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationcaptchaactionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationcaptchaactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationcaptchaaction resource")

	found := r.readAuthenticationcaptchaactionFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AuthenticationcaptchaactionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state AuthenticationcaptchaactionResourceModel

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

	tflog.Debug(ctx, "Updating authenticationcaptchaaction resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	attributesToUnset := []string{}
	if !data.Defaultauthenticationgroup.Equal(state.Defaultauthenticationgroup) {
		tflog.Debug(ctx, fmt.Sprintf("defaultauthenticationgroup has changed for authenticationcaptchaaction"))
		hasChange = true
	}
	if !data.Scorethreshold.Equal(state.Scorethreshold) {
		tflog.Debug(ctx, fmt.Sprintf("scorethreshold has changed for authenticationcaptchaaction"))
		if config.Scorethreshold.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "scorethreshold")
		} else {
			hasChange = true
		}
	}
	// Check secret attribute secretkey or its version tracker
	if !data.Secretkey.Equal(state.Secretkey) {
		tflog.Debug(ctx, fmt.Sprintf("secretkey has changed for authenticationcaptchaaction"))
		hasChange = true
	} else if !data.SecretkeyWoVersion.Equal(state.SecretkeyWoVersion) {
		tflog.Debug(ctx, fmt.Sprintf("secretkey_wo_version has changed for authenticationcaptchaaction"))
		hasChange = true
	}
	if !data.Serverurl.Equal(state.Serverurl) {
		tflog.Debug(ctx, fmt.Sprintf("serverurl has changed for authenticationcaptchaaction"))
		hasChange = true
	}
	// Check secret attribute sitekey or its version tracker
	if !data.Sitekey.Equal(state.Sitekey) {
		tflog.Debug(ctx, fmt.Sprintf("sitekey has changed for authenticationcaptchaaction"))
		hasChange = true
	} else if !data.SitekeyWoVersion.Equal(state.SitekeyWoVersion) {
		tflog.Debug(ctx, fmt.Sprintf("sitekey_wo_version has changed for authenticationcaptchaaction"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		// Get payload from plan (regular attributes)
		authenticationcaptchaaction := authenticationcaptchaactionGetTheUpdatablePayloadFromThePlan(ctx, &data)
		// Add write-only attributes from config to the payload
		authenticationcaptchaactionGetThePayloadFromtheConfig(ctx, &config, &authenticationcaptchaaction)
		// Make API call
		// Named resource - use UpdateResource
		name_value := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Authenticationcaptchaaction.Type(), name_value, &authenticationcaptchaaction)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authenticationcaptchaaction, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated authenticationcaptchaaction resource")
	} else {
		tflog.Debug(ctx, "No changes detected for authenticationcaptchaaction resource, skipping update")
	}

	// Unset attributes that were removed from the configuration so the appliance
	// reverts them to their defaults. Update-then-unset ordering ensures any
	// default carried in the update payload is superseded by the unset.
	unsetIdPayload := map[string]interface{}{
		"name": data.Name.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Authenticationcaptchaaction.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset authenticationcaptchaaction attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readAuthenticationcaptchaactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "authenticationcaptchaaction not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationcaptchaactionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationcaptchaactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationcaptchaaction resource")
	// Named resource - delete using DeleteResource
	name_value := data.Name.ValueString()
	err := r.client.DeleteResource(service.Authenticationcaptchaaction.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete authenticationcaptchaaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted authenticationcaptchaaction resource")
}

// Helper function to read authenticationcaptchaaction data from API
func (r *AuthenticationcaptchaactionResource) readAuthenticationcaptchaactionFromApi(ctx context.Context, data *AuthenticationcaptchaactionResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Authenticationcaptchaaction.Type(), name_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationcaptchaaction, got error: %s", err))
		return false
	}

	authenticationcaptchaactionSetAttrFromGet(ctx, data, getResponseData)

	return true
}
