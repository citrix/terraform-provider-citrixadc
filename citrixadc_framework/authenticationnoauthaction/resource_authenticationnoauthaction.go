package authenticationnoauthaction

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
var _ resource.Resource = &AuthenticationnoauthactionResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationnoauthactionResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationnoauthactionResource)(nil)

func NewAuthenticationnoauthactionResource() resource.Resource {
	return &AuthenticationnoauthactionResource{}
}

// AuthenticationnoauthactionResource defines the resource implementation.
type AuthenticationnoauthactionResource struct {
	client *service.NitroClient
}

func (r *AuthenticationnoauthactionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationnoauthactionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationnoauthaction"
}

func (r *AuthenticationnoauthactionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationnoauthactionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuthenticationnoauthactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationnoauthaction resource")

	authenticationnoauthaction := authenticationnoauthactionGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Authenticationnoauthaction.Type(), name_value, &authenticationnoauthaction)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationnoauthaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created authenticationnoauthaction resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	if !r.readAuthenticationnoauthactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "authenticationnoauthaction not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationnoauthactionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationnoauthactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationnoauthaction resource")

	found := r.readAuthenticationnoauthactionFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AuthenticationnoauthactionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state AuthenticationnoauthactionResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from config (for unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating authenticationnoauthaction resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	attributesToUnset := []string{}
	if !data.Defaultauthenticationgroup.Equal(state.Defaultauthenticationgroup) {
		tflog.Debug(ctx, "defaultauthenticationgroup has changed for authenticationnoauthaction")
		if config.Defaultauthenticationgroup.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "defaultauthenticationgroup")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		// Create API request body from the model
		authenticationnoauthaction := authenticationnoauthactionGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Named resource - use UpdateResource
		name_value := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Authenticationnoauthaction.Type(), name_value, &authenticationnoauthaction)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authenticationnoauthaction, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated authenticationnoauthaction resource")
	} else {
		tflog.Debug(ctx, "No changes detected for authenticationnoauthaction resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults.
	unsetIdPayload := map[string]interface{}{
		"name": data.Name.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Authenticationnoauthaction.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset authenticationnoauthaction attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readAuthenticationnoauthactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "authenticationnoauthaction not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationnoauthactionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationnoauthactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationnoauthaction resource")
	// Named resource - delete using DeleteResource
	name_value := data.Id.ValueString()
	err := r.client.DeleteResource(service.Authenticationnoauthaction.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete authenticationnoauthaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted authenticationnoauthaction resource")
}

// Helper function to read authenticationnoauthaction data from API
func (r *AuthenticationnoauthactionResource) readAuthenticationnoauthactionFromApi(ctx context.Context, data *AuthenticationnoauthactionResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Authenticationnoauthaction.Type(), name_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationnoauthaction, got error: %s", err))
		return false
	}

	authenticationnoauthactionSetAttrFromGet(ctx, data, getResponseData)

	return true
}
