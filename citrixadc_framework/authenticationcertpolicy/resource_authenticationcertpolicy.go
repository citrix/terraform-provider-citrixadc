package authenticationcertpolicy

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
var _ resource.Resource = &AuthenticationcertpolicyResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationcertpolicyResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationcertpolicyResource)(nil)

func NewAuthenticationcertpolicyResource() resource.Resource {
	return &AuthenticationcertpolicyResource{}
}

// AuthenticationcertpolicyResource defines the resource implementation.
type AuthenticationcertpolicyResource struct {
	client *service.NitroClient
}

func (r *AuthenticationcertpolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationcertpolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationcertpolicy"
}

func (r *AuthenticationcertpolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationcertpolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuthenticationcertpolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationcertpolicy resource")
	// Get payload from plan
	authenticationcertpolicy := authenticationcertpolicyGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	authenticationcertpolicyName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Authenticationcertpolicy.Type(), authenticationcertpolicyName, &authenticationcertpolicy)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationcertpolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created authenticationcertpolicy resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	if !r.readAuthenticationcertpolicyFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "authenticationcertpolicy not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationcertpolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationcertpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationcertpolicy resource")

	found := r.readAuthenticationcertpolicyFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AuthenticationcertpolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AuthenticationcertpolicyResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating authenticationcertpolicy resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Reqaction.Equal(state.Reqaction) {
		tflog.Debug(ctx, "reqaction has changed for authenticationcertpolicy")
		hasChange = true
	}
	if !data.Rule.Equal(state.Rule) {
		tflog.Debug(ctx, "rule has changed for authenticationcertpolicy")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		authenticationcertpolicy := authenticationcertpolicyGetThePayloadFromtheConfig(ctx, &data)
		// Make API call
		// Named resource - use UpdateResource
		authenticationcertpolicyName := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Authenticationcertpolicy.Type(), authenticationcertpolicyName, &authenticationcertpolicy)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authenticationcertpolicy, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated authenticationcertpolicy resource")
	} else {
		tflog.Debug(ctx, "No changes detected for authenticationcertpolicy resource, skipping update")
	}

	// Read the updated state back
	if !r.readAuthenticationcertpolicyFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "authenticationcertpolicy not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationcertpolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationcertpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationcertpolicy resource")
	// Named resource - delete using DeleteResource
	authenticationcertpolicyName := data.Name.ValueString()
	err := r.client.DeleteResource(service.Authenticationcertpolicy.Type(), authenticationcertpolicyName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete authenticationcertpolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted authenticationcertpolicy resource")
}

// Helper function to read authenticationcertpolicy data from API
func (r *AuthenticationcertpolicyResource) readAuthenticationcertpolicyFromApi(ctx context.Context, data *AuthenticationcertpolicyResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	authenticationcertpolicyName := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Authenticationcertpolicy.Type(), authenticationcertpolicyName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationcertpolicy, got error: %s", err))
		return false
	}

	authenticationcertpolicySetAttrFromGet(ctx, data, getResponseData)

	return true
}
