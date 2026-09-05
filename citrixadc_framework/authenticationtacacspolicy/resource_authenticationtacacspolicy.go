package authenticationtacacspolicy

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
var _ resource.Resource = &AuthenticationtacacspolicyResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationtacacspolicyResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationtacacspolicyResource)(nil)

func NewAuthenticationtacacspolicyResource() resource.Resource {
	return &AuthenticationtacacspolicyResource{}
}

// AuthenticationtacacspolicyResource defines the resource implementation.
type AuthenticationtacacspolicyResource struct {
	client *service.NitroClient
}

func (r *AuthenticationtacacspolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationtacacspolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationtacacspolicy"
}

func (r *AuthenticationtacacspolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationtacacspolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuthenticationtacacspolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationtacacspolicy resource")

	authenticationtacacspolicy := authenticationtacacspolicyGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	authenticationtacacspolicyName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Authenticationtacacspolicy.Type(), authenticationtacacspolicyName, &authenticationtacacspolicy)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationtacacspolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created authenticationtacacspolicy resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(authenticationtacacspolicyName)

	// Read the updated state back
	if !r.readAuthenticationtacacspolicyFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "authenticationtacacspolicy not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationtacacspolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationtacacspolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationtacacspolicy resource")

	found := r.readAuthenticationtacacspolicyFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AuthenticationtacacspolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AuthenticationtacacspolicyResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating authenticationtacacspolicy resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Reqaction.Equal(state.Reqaction) {
		tflog.Debug(ctx, "reqaction has changed for authenticationtacacspolicy")
		hasChange = true
	}
	if !data.Rule.Equal(state.Rule) {
		tflog.Debug(ctx, "rule has changed for authenticationtacacspolicy")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		authenticationtacacspolicy := authenticationtacacspolicyGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Named resource - use UpdateResource
		authenticationtacacspolicyName := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Authenticationtacacspolicy.Type(), authenticationtacacspolicyName, &authenticationtacacspolicy)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authenticationtacacspolicy, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated authenticationtacacspolicy resource")
	} else {
		tflog.Debug(ctx, "No changes detected for authenticationtacacspolicy resource, skipping update")
	}

	// Read the updated state back
	if !r.readAuthenticationtacacspolicyFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "authenticationtacacspolicy not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationtacacspolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationtacacspolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationtacacspolicy resource")

	// Named resource - delete using DeleteResource
	authenticationtacacspolicyName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Authenticationtacacspolicy.Type(), authenticationtacacspolicyName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete authenticationtacacspolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted authenticationtacacspolicy resource")
}

// Helper function to read authenticationtacacspolicy data from API
func (r *AuthenticationtacacspolicyResource) readAuthenticationtacacspolicyFromApi(ctx context.Context, data *AuthenticationtacacspolicyResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	authenticationtacacspolicyName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Authenticationtacacspolicy.Type(), authenticationtacacspolicyName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationtacacspolicy, got error: %s", err))
		return false
	}

	authenticationtacacspolicySetAttrFromGet(ctx, data, getResponseData)

	return true
}
