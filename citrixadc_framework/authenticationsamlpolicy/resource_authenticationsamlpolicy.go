package authenticationsamlpolicy

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
var _ resource.Resource = &AuthenticationsamlpolicyResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationsamlpolicyResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationsamlpolicyResource)(nil)

func NewAuthenticationsamlpolicyResource() resource.Resource {
	return &AuthenticationsamlpolicyResource{}
}

// AuthenticationsamlpolicyResource defines the resource implementation.
type AuthenticationsamlpolicyResource struct {
	client *service.NitroClient
}

func (r *AuthenticationsamlpolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationsamlpolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationsamlpolicy"
}

func (r *AuthenticationsamlpolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationsamlpolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuthenticationsamlpolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationsamlpolicy resource")

	// Create API request body from the model
	authenticationsamlpolicy := authenticationsamlpolicyGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	authenticationsamlpolicyName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Authenticationsamlpolicy.Type(), authenticationsamlpolicyName, &authenticationsamlpolicy)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationsamlpolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created authenticationsamlpolicy resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(authenticationsamlpolicyName)

	// Read the updated state back
	if !r.readAuthenticationsamlpolicyFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "authenticationsamlpolicy not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationsamlpolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationsamlpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationsamlpolicy resource")

	found := r.readAuthenticationsamlpolicyFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AuthenticationsamlpolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AuthenticationsamlpolicyResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating authenticationsamlpolicy resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Reqaction.Equal(state.Reqaction) {
		tflog.Debug(ctx, "reqaction has changed for authenticationsamlpolicy, starting update")
		hasChange = true
	}
	if !data.Rule.Equal(state.Rule) {
		tflog.Debug(ctx, "rule has changed for authenticationsamlpolicy, starting update")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		authenticationsamlpolicy := authenticationsamlpolicyGetThePayloadFromtheConfig(ctx, &data)
		// Make API call
		// Named resource - use UpdateResource
		authenticationsamlpolicyName := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Authenticationsamlpolicy.Type(), authenticationsamlpolicyName, &authenticationsamlpolicy)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authenticationsamlpolicy, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated authenticationsamlpolicy resource")
	} else {
		tflog.Debug(ctx, "No changes detected for authenticationsamlpolicy resource, skipping update")
	}

	// Read the updated state back
	if !r.readAuthenticationsamlpolicyFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "authenticationsamlpolicy not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationsamlpolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationsamlpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationsamlpolicy resource")

	// Named resource - delete using DeleteResource
	authenticationsamlpolicyName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Authenticationsamlpolicy.Type(), authenticationsamlpolicyName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete authenticationsamlpolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted authenticationsamlpolicy resource")
}

// Helper function to read authenticationsamlpolicy data from API
func (r *AuthenticationsamlpolicyResource) readAuthenticationsamlpolicyFromApi(ctx context.Context, data *AuthenticationsamlpolicyResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value (name)
	authenticationsamlpolicyName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Authenticationsamlpolicy.Type(), authenticationsamlpolicyName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationsamlpolicy, got error: %s", err))
		return false
	}

	authenticationsamlpolicySetAttrFromGet(ctx, data, getResponseData)

	return true
}
