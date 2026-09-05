package authenticationnegotiatepolicy

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
var _ resource.Resource = &AuthenticationnegotiatepolicyResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationnegotiatepolicyResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationnegotiatepolicyResource)(nil)

func NewAuthenticationnegotiatepolicyResource() resource.Resource {
	return &AuthenticationnegotiatepolicyResource{}
}

// AuthenticationnegotiatepolicyResource defines the resource implementation.
type AuthenticationnegotiatepolicyResource struct {
	client *service.NitroClient
}

func (r *AuthenticationnegotiatepolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationnegotiatepolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationnegotiatepolicy"
}

func (r *AuthenticationnegotiatepolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationnegotiatepolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuthenticationnegotiatepolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationnegotiatepolicy resource")

	// Create API request body from the model
	authenticationnegotiatepolicy := authenticationnegotiatepolicyGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	authenticationnegotiatepolicyName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Authenticationnegotiatepolicy.Type(), authenticationnegotiatepolicyName, &authenticationnegotiatepolicy)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationnegotiatepolicy, got error: %s", err))
		return
	}

	// Set ID for the resource before reading state
	data.Id = types.StringValue(authenticationnegotiatepolicyName)

	tflog.Trace(ctx, "Created authenticationnegotiatepolicy resource")

	// Read the updated state back
	if !r.readAuthenticationnegotiatepolicyFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "authenticationnegotiatepolicy not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationnegotiatepolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationnegotiatepolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationnegotiatepolicy resource")

	found := r.readAuthenticationnegotiatepolicyFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AuthenticationnegotiatepolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AuthenticationnegotiatepolicyResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating authenticationnegotiatepolicy resource")

	// Check if there are any changes in updateable attributes
	// name is ForceNew/RequiresReplace and never reaches Update
	hasChange := false
	if !data.Reqaction.Equal(state.Reqaction) {
		tflog.Debug(ctx, "reqaction has changed for authenticationnegotiatepolicy")
		hasChange = true
	}
	if !data.Rule.Equal(state.Rule) {
		tflog.Debug(ctx, "rule has changed for authenticationnegotiatepolicy")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		authenticationnegotiatepolicy := authenticationnegotiatepolicyGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Named resource - use UpdateResource
		authenticationnegotiatepolicyName := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Authenticationnegotiatepolicy.Type(), authenticationnegotiatepolicyName, &authenticationnegotiatepolicy)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authenticationnegotiatepolicy, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated authenticationnegotiatepolicy resource")
	} else {
		tflog.Debug(ctx, "No changes detected for authenticationnegotiatepolicy resource, skipping update")
	}

	// Read the updated state back
	if !r.readAuthenticationnegotiatepolicyFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "authenticationnegotiatepolicy not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationnegotiatepolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationnegotiatepolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationnegotiatepolicy resource")
	// Named resource - delete using DeleteResource
	authenticationnegotiatepolicyName := data.Name.ValueString()
	err := r.client.DeleteResource(service.Authenticationnegotiatepolicy.Type(), authenticationnegotiatepolicyName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete authenticationnegotiatepolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted authenticationnegotiatepolicy resource")
}

// Helper function to read authenticationnegotiatepolicy data from API
func (r *AuthenticationnegotiatepolicyResource) readAuthenticationnegotiatepolicyFromApi(ctx context.Context, data *AuthenticationnegotiatepolicyResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value (name)
	authenticationnegotiatepolicyName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Authenticationnegotiatepolicy.Type(), authenticationnegotiatepolicyName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationnegotiatepolicy, got error: %s", err))
		return false
	}

	authenticationnegotiatepolicySetAttrFromGet(ctx, data, getResponseData)

	return true
}
