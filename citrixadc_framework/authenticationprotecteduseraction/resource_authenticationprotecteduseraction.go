package authenticationprotecteduseraction

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AuthenticationprotecteduseractionResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationprotecteduseractionResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationprotecteduseractionResource)(nil)

func NewAuthenticationprotecteduseractionResource() resource.Resource {
	return &AuthenticationprotecteduseractionResource{}
}

// AuthenticationprotecteduseractionResource defines the resource implementation.
type AuthenticationprotecteduseractionResource struct {
	client *service.NitroClient
}

func (r *AuthenticationprotecteduseractionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationprotecteduseractionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationprotecteduseraction"
}

func (r *AuthenticationprotecteduseractionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationprotecteduseractionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuthenticationprotecteduseractionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationprotecteduseraction resource")
	authenticationprotecteduseraction := authenticationprotecteduseractionGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Authenticationprotecteduseraction.Type(), name_value, &authenticationprotecteduseraction)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationprotecteduseraction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created authenticationprotecteduseraction resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	r.readAuthenticationprotecteduseractionFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationprotecteduseractionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationprotecteduseractionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationprotecteduseraction resource")

	r.readAuthenticationprotecteduseractionFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationprotecteduseractionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AuthenticationprotecteduseractionResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating authenticationprotecteduseraction resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Maxconcurrentusers.Equal(state.Maxconcurrentusers) {
		tflog.Debug(ctx, fmt.Sprintf("maxconcurrentusers has changed for authenticationprotecteduseraction"))
		hasChange = true
	}
	if !data.Realmstr.Equal(state.Realmstr) {
		tflog.Debug(ctx, fmt.Sprintf("realmstr has changed for authenticationprotecteduseraction"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		authenticationprotecteduseraction := authenticationprotecteduseractionGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Named resource - use UpdateResource
		name_value := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Authenticationprotecteduseraction.Type(), name_value, &authenticationprotecteduseraction)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authenticationprotecteduseraction, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated authenticationprotecteduseraction resource")
	} else {
		tflog.Debug(ctx, "No changes detected for authenticationprotecteduseraction resource, skipping update")
	}

	// Read the updated state back
	r.readAuthenticationprotecteduseractionFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationprotecteduseractionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationprotecteduseractionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationprotecteduseraction resource")
	// Named resource - delete using DeleteResource
	name_value := data.Name.ValueString()
	err := r.client.DeleteResource(service.Authenticationprotecteduseraction.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete authenticationprotecteduseraction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted authenticationprotecteduseraction resource")
}

// Helper function to read authenticationprotecteduseraction data from API
func (r *AuthenticationprotecteduseractionResource) readAuthenticationprotecteduseractionFromApi(ctx context.Context, data *AuthenticationprotecteduseractionResourceModel, diags *diag.Diagnostics) {

	// Case 2: Find with single ID attribute - ID is the plain value
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Authenticationprotecteduseraction.Type(), name_Name)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationprotecteduseraction, got error: %s", err))
		return
	}

	authenticationprotecteduseractionSetAttrFromGet(ctx, data, getResponseData)

}
