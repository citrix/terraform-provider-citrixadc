package authenticationwebauthpolicy

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
var _ resource.Resource = &AuthenticationwebauthpolicyResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationwebauthpolicyResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationwebauthpolicyResource)(nil)

func NewAuthenticationwebauthpolicyResource() resource.Resource {
	return &AuthenticationwebauthpolicyResource{}
}

// AuthenticationwebauthpolicyResource defines the resource implementation.
type AuthenticationwebauthpolicyResource struct {
	client *service.NitroClient
}

func (r *AuthenticationwebauthpolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationwebauthpolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationwebauthpolicy"
}

func (r *AuthenticationwebauthpolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationwebauthpolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuthenticationwebauthpolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationwebauthpolicy resource")

	// Create API request body from the model
	authenticationwebauthpolicy := authenticationwebauthpolicyGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	authenticationwebauthpolicyName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Authenticationwebauthpolicy.Type(), authenticationwebauthpolicyName, &authenticationwebauthpolicy)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationwebauthpolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created authenticationwebauthpolicy resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", authenticationwebauthpolicyName))

	// Read the updated state back
	if !r.readAuthenticationwebauthpolicyFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "authenticationwebauthpolicy not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationwebauthpolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationwebauthpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationwebauthpolicy resource")

	found := r.readAuthenticationwebauthpolicyFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AuthenticationwebauthpolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AuthenticationwebauthpolicyResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating authenticationwebauthpolicy resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Action.Equal(state.Action) {
		tflog.Debug(ctx, "action has changed for authenticationwebauthpolicy")
		hasChange = true
	}
	if !data.Rule.Equal(state.Rule) {
		tflog.Debug(ctx, "rule has changed for authenticationwebauthpolicy")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		authenticationwebauthpolicy := authenticationwebauthpolicyGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Named resource - use UpdateResource
		authenticationwebauthpolicyName := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Authenticationwebauthpolicy.Type(), authenticationwebauthpolicyName, &authenticationwebauthpolicy)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authenticationwebauthpolicy, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated authenticationwebauthpolicy resource")
	} else {
		tflog.Debug(ctx, "No changes detected for authenticationwebauthpolicy resource, skipping update")
	}

	// Read the updated state back
	if !r.readAuthenticationwebauthpolicyFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "authenticationwebauthpolicy not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationwebauthpolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationwebauthpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationwebauthpolicy resource")
	// Named resource - delete using DeleteResource
	authenticationwebauthpolicyName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Authenticationwebauthpolicy.Type(), authenticationwebauthpolicyName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete authenticationwebauthpolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted authenticationwebauthpolicy resource")
}

// Helper function to read authenticationwebauthpolicy data from API
func (r *AuthenticationwebauthpolicyResource) readAuthenticationwebauthpolicyFromApi(ctx context.Context, data *AuthenticationwebauthpolicyResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	authenticationwebauthpolicyName := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Authenticationwebauthpolicy.Type(), authenticationwebauthpolicyName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationwebauthpolicy, got error: %s", err))
		return false
	}

	authenticationwebauthpolicySetAttrFromGet(ctx, data, getResponseData)

	return true
}
