package authenticationradiuspolicy

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
var _ resource.Resource = &AuthenticationradiuspolicyResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationradiuspolicyResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationradiuspolicyResource)(nil)

func NewAuthenticationradiuspolicyResource() resource.Resource {
	return &AuthenticationradiuspolicyResource{}
}

// AuthenticationradiuspolicyResource defines the resource implementation.
type AuthenticationradiuspolicyResource struct {
	client *service.NitroClient
}

func (r *AuthenticationradiuspolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationradiuspolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationradiuspolicy"
}

func (r *AuthenticationradiuspolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationradiuspolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuthenticationradiuspolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationradiuspolicy resource")

	// Create API request body from the model
	authenticationradiuspolicy := authenticationradiuspolicyGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource (NITRO add: POST /authenticationradiuspolicy)
	authenticationradiuspolicyName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Authenticationradiuspolicy.Type(), authenticationradiuspolicyName, &authenticationradiuspolicy)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationradiuspolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created authenticationradiuspolicy resource")

	// Set ID for the resource before reading state back (single unique attribute -> plain name value)
	data.Id = types.StringValue(authenticationradiuspolicyName)

	// Read the updated state back
	if !r.readAuthenticationradiuspolicyFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "authenticationradiuspolicy not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationradiuspolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationradiuspolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationradiuspolicy resource")

	found := r.readAuthenticationradiuspolicyFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AuthenticationradiuspolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AuthenticationradiuspolicyResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating authenticationradiuspolicy resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Reqaction.Equal(state.Reqaction) {
		tflog.Debug(ctx, "reqaction has changed for authenticationradiuspolicy")
		hasChange = true
	}
	if !data.Rule.Equal(state.Rule) {
		tflog.Debug(ctx, "rule has changed for authenticationradiuspolicy")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		authenticationradiuspolicy := authenticationradiuspolicyGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Named resource - use UpdateResource (NITRO update: PUT /authenticationradiuspolicy)
		authenticationradiuspolicyName := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Authenticationradiuspolicy.Type(), authenticationradiuspolicyName, &authenticationradiuspolicy)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authenticationradiuspolicy, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated authenticationradiuspolicy resource")
	} else {
		tflog.Debug(ctx, "No changes detected for authenticationradiuspolicy resource, skipping update")
	}

	// Read the updated state back
	if !r.readAuthenticationradiuspolicyFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "authenticationradiuspolicy not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationradiuspolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationradiuspolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationradiuspolicy resource")
	// Named resource - delete using DeleteResource (NITRO delete: DELETE /authenticationradiuspolicy/{name})
	authenticationradiuspolicyName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Authenticationradiuspolicy.Type(), authenticationradiuspolicyName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete authenticationradiuspolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted authenticationradiuspolicy resource")
}

// Helper function to read authenticationradiuspolicy data from API
func (r *AuthenticationradiuspolicyResource) readAuthenticationradiuspolicyFromApi(ctx context.Context, data *AuthenticationradiuspolicyResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain name value
	authenticationradiuspolicyName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Authenticationradiuspolicy.Type(), authenticationradiuspolicyName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationradiuspolicy, got error: %s", err))
		return false
	}

	authenticationradiuspolicySetAttrFromGet(ctx, data, getResponseData)

	return true
}
