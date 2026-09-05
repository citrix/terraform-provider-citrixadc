package authenticationsamlaction

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
var _ resource.Resource = &AuthenticationsamlactionResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationsamlactionResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationsamlactionResource)(nil)

func NewAuthenticationsamlactionResource() resource.Resource {
	return &AuthenticationsamlactionResource{}
}

// AuthenticationsamlactionResource defines the resource implementation.
type AuthenticationsamlactionResource struct {
	client *service.NitroClient
}

func (r *AuthenticationsamlactionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationsamlactionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationsamlaction"
}

func (r *AuthenticationsamlactionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationsamlactionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuthenticationsamlactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationsamlaction resource")

	// Create API request body from the model
	authenticationsamlaction := authenticationsamlactionGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource
	authenticationsamlactionName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Authenticationsamlaction.Type(), authenticationsamlactionName, &authenticationsamlaction)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationsamlaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created authenticationsamlaction resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(authenticationsamlactionName)

	// Read the updated state back
	if !r.readAuthenticationsamlactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "authenticationsamlaction not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationsamlactionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationsamlactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationsamlaction resource")

	found := r.readAuthenticationsamlactionFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AuthenticationsamlactionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state AuthenticationsamlactionResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from config (to be unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating authenticationsamlaction resource")

	// Determine attributes removed from config so they can be unset on the
	// appliance (revert to NITRO defaults).
	hasChange := false
	attributesToUnset := []string{}
	if !data.Attributeconsumingserviceindex.Equal(state.Attributeconsumingserviceindex) {
		if config.Attributeconsumingserviceindex.IsNull() {
			attributesToUnset = append(attributesToUnset, "attributeconsumingserviceindex")
		} else {
			hasChange = true
		}
	}
	if !data.Digestmethod.Equal(state.Digestmethod) {
		if config.Digestmethod.IsNull() {
			attributesToUnset = append(attributesToUnset, "digestmethod")
		} else {
			hasChange = true
		}
	}
	if !data.Enforceusername.Equal(state.Enforceusername) {
		if config.Enforceusername.IsNull() {
			attributesToUnset = append(attributesToUnset, "enforceusername")
		} else {
			hasChange = true
		}
	}
	if !data.Forceauthn.Equal(state.Forceauthn) {
		if config.Forceauthn.IsNull() {
			attributesToUnset = append(attributesToUnset, "forceauthn")
		} else {
			hasChange = true
		}
	}
	if !data.Logoutbinding.Equal(state.Logoutbinding) {
		if config.Logoutbinding.IsNull() {
			attributesToUnset = append(attributesToUnset, "logoutbinding")
		} else {
			hasChange = true
		}
	}
	if !data.Requestedauthncontext.Equal(state.Requestedauthncontext) {
		if config.Requestedauthncontext.IsNull() {
			attributesToUnset = append(attributesToUnset, "requestedauthncontext")
		} else {
			hasChange = true
		}
	}
	if !data.Samlacsindex.Equal(state.Samlacsindex) {
		if config.Samlacsindex.IsNull() {
			attributesToUnset = append(attributesToUnset, "samlacsindex")
		} else {
			hasChange = true
		}
	}
	if !data.Samlbinding.Equal(state.Samlbinding) {
		if config.Samlbinding.IsNull() {
			attributesToUnset = append(attributesToUnset, "samlbinding")
		} else {
			hasChange = true
		}
	}
	if !data.Samlrejectunsignedassertion.Equal(state.Samlrejectunsignedassertion) {
		if config.Samlrejectunsignedassertion.IsNull() {
			attributesToUnset = append(attributesToUnset, "samlrejectunsignedassertion")
		} else {
			hasChange = true
		}
	}
	if !data.Samltwofactor.Equal(state.Samltwofactor) {
		if config.Samltwofactor.IsNull() {
			attributesToUnset = append(attributesToUnset, "samltwofactor")
		} else {
			hasChange = true
		}
	}
	if !data.Sendthumbprint.Equal(state.Sendthumbprint) {
		if config.Sendthumbprint.IsNull() {
			attributesToUnset = append(attributesToUnset, "sendthumbprint")
		} else {
			hasChange = true
		}
	}
	if !data.Signaturealg.Equal(state.Signaturealg) {
		if config.Signaturealg.IsNull() {
			attributesToUnset = append(attributesToUnset, "signaturealg")
		} else {
			hasChange = true
		}
	}
	if !data.Skewtime.Equal(state.Skewtime) {
		if config.Skewtime.IsNull() {
			attributesToUnset = append(attributesToUnset, "skewtime")
		} else {
			hasChange = true
		}
	}
	if !data.Storesamlresponse.Equal(state.Storesamlresponse) {
		if config.Storesamlresponse.IsNull() {
			attributesToUnset = append(attributesToUnset, "storesamlresponse")
		} else {
			hasChange = true
		}
	}
	_ = hasChange

	// Create API request body from the model
	authenticationsamlaction := authenticationsamlactionGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use UpdateResource
	authenticationsamlactionName := data.Name.ValueString()
	_, err := r.client.UpdateResource(service.Authenticationsamlaction.Type(), authenticationsamlactionName, &authenticationsamlaction)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authenticationsamlaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Updated authenticationsamlaction resource")

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults.
	unsetIdPayload := map[string]interface{}{
		"name": data.Name.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Authenticationsamlaction.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset authenticationsamlaction attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readAuthenticationsamlactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "authenticationsamlaction not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationsamlactionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationsamlactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationsamlaction resource")

	// Named resource - delete using DeleteResource
	authenticationsamlactionName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Authenticationsamlaction.Type(), authenticationsamlactionName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete authenticationsamlaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted authenticationsamlaction resource")
}

// Helper function to read authenticationsamlaction data from API
func (r *AuthenticationsamlactionResource) readAuthenticationsamlactionFromApi(ctx context.Context, data *AuthenticationsamlactionResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value (name)
	authenticationsamlactionName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Authenticationsamlaction.Type(), authenticationsamlactionName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationsamlaction, got error: %s", err))
		return false
	}

	authenticationsamlactionSetAttrFromGet(ctx, data, getResponseData)

	return true
}
