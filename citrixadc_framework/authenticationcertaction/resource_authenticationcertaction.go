package authenticationcertaction

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
var _ resource.Resource = &AuthenticationcertactionResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationcertactionResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationcertactionResource)(nil)

func NewAuthenticationcertactionResource() resource.Resource {
	return &AuthenticationcertactionResource{}
}

// AuthenticationcertactionResource defines the resource implementation.
type AuthenticationcertactionResource struct {
	client *service.NitroClient
}

func (r *AuthenticationcertactionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationcertactionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationcertaction"
}

func (r *AuthenticationcertactionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationcertactionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuthenticationcertactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationcertaction resource")
	// Get payload from plan
	authenticationcertaction := authenticationcertactionGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	authenticationcertactionName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Authenticationcertaction.Type(), authenticationcertactionName, &authenticationcertaction)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationcertaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created authenticationcertaction resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	if !r.readAuthenticationcertactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "authenticationcertaction not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationcertactionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationcertactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationcertaction resource")

	found := r.readAuthenticationcertactionFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AuthenticationcertactionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state AuthenticationcertactionResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from configuration
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating authenticationcertaction resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	attributesToUnset := []string{}
	if !data.Defaultauthenticationgroup.Equal(state.Defaultauthenticationgroup) {
		tflog.Debug(ctx, "defaultauthenticationgroup has changed for authenticationcertaction")
		hasChange = true
	}
	if !data.Groupnamefield.Equal(state.Groupnamefield) {
		tflog.Debug(ctx, "groupnamefield has changed for authenticationcertaction")
		hasChange = true
	}
	if !data.Twofactor.Equal(state.Twofactor) {
		tflog.Debug(ctx, "twofactor has changed for authenticationcertaction")
		if config.Twofactor.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "twofactor")
		} else {
			hasChange = true
		}
	}
	if !data.Usernamefield.Equal(state.Usernamefield) {
		tflog.Debug(ctx, "usernamefield has changed for authenticationcertaction")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		authenticationcertaction := authenticationcertactionGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Named resource - use UpdateResource
		authenticationcertactionName := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Authenticationcertaction.Type(), authenticationcertactionName, &authenticationcertaction)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authenticationcertaction, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated authenticationcertaction resource")
	} else {
		tflog.Debug(ctx, "No changes detected for authenticationcertaction resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults.
	unsetIdPayload := map[string]interface{}{
		"name": data.Name.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Authenticationcertaction.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset authenticationcertaction attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readAuthenticationcertactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "authenticationcertaction not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationcertactionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationcertactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationcertaction resource")
	// Named resource - delete using DeleteResource
	authenticationcertactionName := data.Name.ValueString()
	err := r.client.DeleteResource(service.Authenticationcertaction.Type(), authenticationcertactionName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete authenticationcertaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted authenticationcertaction resource")
}

// Helper function to read authenticationcertaction data from API
func (r *AuthenticationcertactionResource) readAuthenticationcertactionFromApi(ctx context.Context, data *AuthenticationcertactionResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	authenticationcertactionName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Authenticationcertaction.Type(), authenticationcertactionName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationcertaction, got error: %s", err))
		return false
	}

	authenticationcertactionSetAttrFromGet(ctx, data, getResponseData)

	return true
}
