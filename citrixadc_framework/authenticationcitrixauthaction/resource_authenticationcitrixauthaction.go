package authenticationcitrixauthaction

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
var _ resource.Resource = &AuthenticationcitrixauthactionResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationcitrixauthactionResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationcitrixauthactionResource)(nil)

func NewAuthenticationcitrixauthactionResource() resource.Resource {
	return &AuthenticationcitrixauthactionResource{}
}

// AuthenticationcitrixauthactionResource defines the resource implementation.
type AuthenticationcitrixauthactionResource struct {
	client *service.NitroClient
}

func (r *AuthenticationcitrixauthactionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationcitrixauthactionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationcitrixauthaction"
}

func (r *AuthenticationcitrixauthactionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationcitrixauthactionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuthenticationcitrixauthactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationcitrixauthaction resource")
	// Get payload from plan
	authenticationcitrixauthaction := authenticationcitrixauthactionGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Authenticationcitrixauthaction.Type(), name_value, &authenticationcitrixauthaction)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationcitrixauthaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created authenticationcitrixauthaction resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	if !r.readAuthenticationcitrixauthactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "authenticationcitrixauthaction not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationcitrixauthactionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationcitrixauthactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationcitrixauthaction resource")

	found := r.readAuthenticationcitrixauthactionFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AuthenticationcitrixauthactionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AuthenticationcitrixauthactionResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating authenticationcitrixauthaction resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Authentication.Equal(state.Authentication) {
		tflog.Debug(ctx, "authentication has changed for authenticationcitrixauthaction")
		hasChange = true
	}
	if !data.Authenticationtype.Equal(state.Authenticationtype) {
		tflog.Debug(ctx, "authenticationtype has changed for authenticationcitrixauthaction")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		authenticationcitrixauthaction := authenticationcitrixauthactionGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Named resource - use UpdateResource
		name_value := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Authenticationcitrixauthaction.Type(), name_value, &authenticationcitrixauthaction)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authenticationcitrixauthaction, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated authenticationcitrixauthaction resource")
	} else {
		tflog.Debug(ctx, "No changes detected for authenticationcitrixauthaction resource, skipping update")
	}

	// Read the updated state back
	if !r.readAuthenticationcitrixauthactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "authenticationcitrixauthaction not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationcitrixauthactionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationcitrixauthactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationcitrixauthaction resource")
	// Named resource - delete using DeleteResource
	name_value := data.Name.ValueString()
	err := r.client.DeleteResource(service.Authenticationcitrixauthaction.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete authenticationcitrixauthaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted authenticationcitrixauthaction resource")
}

// Helper function to read authenticationcitrixauthaction data from API
func (r *AuthenticationcitrixauthactionResource) readAuthenticationcitrixauthactionFromApi(ctx context.Context, data *AuthenticationcitrixauthactionResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Authenticationcitrixauthaction.Type(), name_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationcitrixauthaction, got error: %s", err))
		return false
	}

	authenticationcitrixauthactionSetAttrFromGet(ctx, data, getResponseData)

	return true
}
