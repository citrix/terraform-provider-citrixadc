package authenticationadfsproxyprofile

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
var _ resource.Resource = &AuthenticationadfsproxyprofileResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationadfsproxyprofileResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationadfsproxyprofileResource)(nil)
var _ resource.ResourceWithValidateConfig = (*AuthenticationadfsproxyprofileResource)(nil)

func NewAuthenticationadfsproxyprofileResource() resource.Resource {
	return &AuthenticationadfsproxyprofileResource{}
}

// AuthenticationadfsproxyprofileResource defines the resource implementation.
type AuthenticationadfsproxyprofileResource struct {
	client *service.NitroClient
}

func (r *AuthenticationadfsproxyprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationadfsproxyprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationadfsproxyprofile"
}

func (r *AuthenticationadfsproxyprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationadfsproxyprofileResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data AuthenticationadfsproxyprofileResourceModel
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

func (r *AuthenticationadfsproxyprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config AuthenticationadfsproxyprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationadfsproxyprofile resource")
	// Get payload from plan (regular attributes)
	authenticationadfsproxyprofile := authenticationadfsproxyprofileGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	authenticationadfsproxyprofileGetThePayloadFromtheConfig(ctx, &config, &authenticationadfsproxyprofile)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Authenticationadfsproxyprofile.Type(), name_value, &authenticationadfsproxyprofile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationadfsproxyprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created authenticationadfsproxyprofile resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	r.readAuthenticationadfsproxyprofileFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationadfsproxyprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationadfsproxyprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationadfsproxyprofile resource")

	r.readAuthenticationadfsproxyprofileFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationadfsproxyprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state AuthenticationadfsproxyprofileResourceModel

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

	tflog.Debug(ctx, "Updating authenticationadfsproxyprofile resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Certkeyname.Equal(state.Certkeyname) {
		tflog.Debug(ctx, fmt.Sprintf("certkeyname has changed for authenticationadfsproxyprofile"))
		hasChange = true
	}
	// Check secret attribute password or its version tracker
	if !data.Password.Equal(state.Password) {
		tflog.Debug(ctx, fmt.Sprintf("password has changed for authenticationadfsproxyprofile"))
		hasChange = true
	} else if !data.PasswordWoVersion.Equal(state.PasswordWoVersion) {
		tflog.Debug(ctx, fmt.Sprintf("password_wo_version has changed for authenticationadfsproxyprofile"))
		hasChange = true
	}
	if !data.Serverurl.Equal(state.Serverurl) {
		tflog.Debug(ctx, fmt.Sprintf("serverurl has changed for authenticationadfsproxyprofile"))
		hasChange = true
	}
	if !data.Username.Equal(state.Username) {
		tflog.Debug(ctx, fmt.Sprintf("username has changed for authenticationadfsproxyprofile"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		// Get payload from plan (regular attributes)
		authenticationadfsproxyprofile := authenticationadfsproxyprofileGetThePayloadFromthePlan(ctx, &data)
		// Add write-only attributes from config to the payload
		authenticationadfsproxyprofileGetThePayloadFromtheConfig(ctx, &config, &authenticationadfsproxyprofile)
		// Make API call
		// Named resource - use UpdateResource
		name_value := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Authenticationadfsproxyprofile.Type(), name_value, &authenticationadfsproxyprofile)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authenticationadfsproxyprofile, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated authenticationadfsproxyprofile resource")
	} else {
		tflog.Debug(ctx, "No changes detected for authenticationadfsproxyprofile resource, skipping update")
	}

	// Read the updated state back
	r.readAuthenticationadfsproxyprofileFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationadfsproxyprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationadfsproxyprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationadfsproxyprofile resource")
	// Named resource - delete using DeleteResource
	name_value := data.Name.ValueString()
	err := r.client.DeleteResource(service.Authenticationadfsproxyprofile.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete authenticationadfsproxyprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted authenticationadfsproxyprofile resource")
}

// Helper function to read authenticationadfsproxyprofile data from API
func (r *AuthenticationadfsproxyprofileResource) readAuthenticationadfsproxyprofileFromApi(ctx context.Context, data *AuthenticationadfsproxyprofileResourceModel, diags *diag.Diagnostics) {

	// Case 2: Find with single ID attribute - ID is the plain value
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Authenticationadfsproxyprofile.Type(), name_Name)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationadfsproxyprofile, got error: %s", err))
		return
	}

	authenticationadfsproxyprofileSetAttrFromGet(ctx, data, getResponseData)

}
