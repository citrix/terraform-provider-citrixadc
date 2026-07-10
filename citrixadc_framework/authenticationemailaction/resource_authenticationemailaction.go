package authenticationemailaction

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
var _ resource.Resource = &AuthenticationemailactionResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationemailactionResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationemailactionResource)(nil)
var _ resource.ResourceWithValidateConfig = (*AuthenticationemailactionResource)(nil)

func NewAuthenticationemailactionResource() resource.Resource {
	return &AuthenticationemailactionResource{}
}

// AuthenticationemailactionResource defines the resource implementation.
type AuthenticationemailactionResource struct {
	client *service.NitroClient
}

func (r *AuthenticationemailactionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationemailactionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationemailaction"
}

func (r *AuthenticationemailactionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationemailactionResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data AuthenticationemailactionResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Validate that either password or password_wo is specified
	if data.Password.IsNull() && data.PasswordWo.IsNull() {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Missing Required Attribute",
			"Either \"password\" or \"password_wo\" must be specified.",
		)
	}
}

func (r *AuthenticationemailactionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config AuthenticationemailactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationemailaction resource")
	// Get payload from plan (regular attributes)
	authenticationemailaction := authenticationemailactionGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	authenticationemailactionGetThePayloadFromtheConfig(ctx, &config, &authenticationemailaction)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Authenticationemailaction.Type(), name_value, &authenticationemailaction)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationemailaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created authenticationemailaction resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	if !r.readAuthenticationemailactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "authenticationemailaction not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationemailactionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationemailactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationemailaction resource")

	found := r.readAuthenticationemailactionFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AuthenticationemailactionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state AuthenticationemailactionResourceModel

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

	tflog.Debug(ctx, "Updating authenticationemailaction resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Content.Equal(state.Content) {
		tflog.Debug(ctx, fmt.Sprintf("content has changed for authenticationemailaction"))
		hasChange = true
	}
	if !data.Defaultauthenticationgroup.Equal(state.Defaultauthenticationgroup) {
		tflog.Debug(ctx, fmt.Sprintf("defaultauthenticationgroup has changed for authenticationemailaction"))
		hasChange = true
	}
	if !data.Emailaddress.Equal(state.Emailaddress) {
		tflog.Debug(ctx, fmt.Sprintf("emailaddress has changed for authenticationemailaction"))
		hasChange = true
	}
	// Check secret attribute password or its version tracker
	if !data.Password.Equal(state.Password) {
		tflog.Debug(ctx, fmt.Sprintf("password has changed for authenticationemailaction"))
		hasChange = true
	} else if !data.PasswordWoVersion.Equal(state.PasswordWoVersion) {
		tflog.Debug(ctx, fmt.Sprintf("password_wo_version has changed for authenticationemailaction"))
		hasChange = true
	}
	if !data.Serverurl.Equal(state.Serverurl) {
		tflog.Debug(ctx, fmt.Sprintf("serverurl has changed for authenticationemailaction"))
		hasChange = true
	}
	if !data.Timeout.Equal(state.Timeout) {
		tflog.Debug(ctx, fmt.Sprintf("timeout has changed for authenticationemailaction"))
		hasChange = true
	}
	if !data.Type.Equal(state.Type) {
		tflog.Debug(ctx, fmt.Sprintf("type has changed for authenticationemailaction"))
		hasChange = true
	}
	if !data.Username.Equal(state.Username) {
		tflog.Debug(ctx, fmt.Sprintf("username has changed for authenticationemailaction"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		// Get payload from plan (regular attributes)
		authenticationemailaction := authenticationemailactionGetTheUpdatablePayloadFromThePlan(ctx, &data)
		// Add write-only attributes from config to the payload
		authenticationemailactionGetThePayloadFromtheConfig(ctx, &config, &authenticationemailaction)
		// Make API call
		// Named resource - use UpdateResource
		name_value := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Authenticationemailaction.Type(), name_value, &authenticationemailaction)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authenticationemailaction, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated authenticationemailaction resource")
	} else {
		tflog.Debug(ctx, "No changes detected for authenticationemailaction resource, skipping update")
	}

	// Read the updated state back
	if !r.readAuthenticationemailactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "authenticationemailaction not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationemailactionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationemailactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationemailaction resource")
	// Named resource - delete using DeleteResource
	name_value := data.Name.ValueString()
	err := r.client.DeleteResource(service.Authenticationemailaction.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete authenticationemailaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted authenticationemailaction resource")
}

// Helper function to read authenticationemailaction data from API
func (r *AuthenticationemailactionResource) readAuthenticationemailactionFromApi(ctx context.Context, data *AuthenticationemailactionResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Authenticationemailaction.Type(), name_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationemailaction, got error: %s", err))
		return false
	}

	authenticationemailactionSetAttrFromGet(ctx, data, getResponseData)

	return true
}
