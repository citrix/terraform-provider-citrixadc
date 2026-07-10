package authenticationdfaaction

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
var _ resource.Resource = &AuthenticationdfaactionResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationdfaactionResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationdfaactionResource)(nil)
var _ resource.ResourceWithValidateConfig = (*AuthenticationdfaactionResource)(nil)

func NewAuthenticationdfaactionResource() resource.Resource {
	return &AuthenticationdfaactionResource{}
}

// AuthenticationdfaactionResource defines the resource implementation.
type AuthenticationdfaactionResource struct {
	client *service.NitroClient
}

func (r *AuthenticationdfaactionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationdfaactionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationdfaaction"
}

func (r *AuthenticationdfaactionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationdfaactionResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data AuthenticationdfaactionResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Validate that either passphrase or passphrase_wo is specified
	if data.Passphrase.IsNull() && data.PassphraseWo.IsNull() {
		resp.Diagnostics.AddAttributeError(
			path.Root("passphrase"),
			"Missing Required Attribute",
			"Either \"passphrase\" or \"passphrase_wo\" must be specified.",
		)
	}
}

func (r *AuthenticationdfaactionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config AuthenticationdfaactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationdfaaction resource")
	// Get payload from plan (regular attributes)
	authenticationdfaaction := authenticationdfaactionGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	authenticationdfaactionGetThePayloadFromtheConfig(ctx, &config, &authenticationdfaaction)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Authenticationdfaaction.Type(), name_value, &authenticationdfaaction)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationdfaaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created authenticationdfaaction resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	if !r.readAuthenticationdfaactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "authenticationdfaaction not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationdfaactionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationdfaactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationdfaaction resource")

	found := r.readAuthenticationdfaactionFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AuthenticationdfaactionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state AuthenticationdfaactionResourceModel

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

	tflog.Debug(ctx, "Updating authenticationdfaaction resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Clientid.Equal(state.Clientid) {
		tflog.Debug(ctx, fmt.Sprintf("clientid has changed for authenticationdfaaction"))
		hasChange = true
	}
	if !data.Defaultauthenticationgroup.Equal(state.Defaultauthenticationgroup) {
		tflog.Debug(ctx, fmt.Sprintf("defaultauthenticationgroup has changed for authenticationdfaaction"))
		hasChange = true
	}
	// Check secret attribute passphrase or its version tracker
	if !data.Passphrase.Equal(state.Passphrase) {
		tflog.Debug(ctx, fmt.Sprintf("passphrase has changed for authenticationdfaaction"))
		hasChange = true
	} else if !data.PassphraseWoVersion.Equal(state.PassphraseWoVersion) {
		tflog.Debug(ctx, fmt.Sprintf("passphrase_wo_version has changed for authenticationdfaaction"))
		hasChange = true
	}
	if !data.Serverurl.Equal(state.Serverurl) {
		tflog.Debug(ctx, fmt.Sprintf("serverurl has changed for authenticationdfaaction"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		// Get payload from plan (regular attributes)
		authenticationdfaaction := authenticationdfaactionGetThePayloadFromthePlan(ctx, &data)
		// Add write-only attributes from config to the payload
		authenticationdfaactionGetThePayloadFromtheConfig(ctx, &config, &authenticationdfaaction)
		// Make API call
		// Named resource - use UpdateResource
		name_value := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Authenticationdfaaction.Type(), name_value, &authenticationdfaaction)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authenticationdfaaction, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated authenticationdfaaction resource")
	} else {
		tflog.Debug(ctx, "No changes detected for authenticationdfaaction resource, skipping update")
	}

	// Read the updated state back
	if !r.readAuthenticationdfaactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "authenticationdfaaction not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationdfaactionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationdfaactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationdfaaction resource")
	// Named resource - delete using DeleteResource
	name_value := data.Name.ValueString()
	err := r.client.DeleteResource(service.Authenticationdfaaction.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete authenticationdfaaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted authenticationdfaaction resource")
}

// Helper function to read authenticationdfaaction data from API
func (r *AuthenticationdfaactionResource) readAuthenticationdfaactionFromApi(ctx context.Context, data *AuthenticationdfaactionResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Authenticationdfaaction.Type(), name_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationdfaaction, got error: %s", err))
		return false
	}

	authenticationdfaactionSetAttrFromGet(ctx, data, getResponseData)

	return true
}
