package authenticationloginschema

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
var _ resource.Resource = &AuthenticationloginschemaResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationloginschemaResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationloginschemaResource)(nil)

func NewAuthenticationloginschemaResource() resource.Resource {
	return &AuthenticationloginschemaResource{}
}

// AuthenticationloginschemaResource defines the resource implementation.
type AuthenticationloginschemaResource struct {
	client *service.NitroClient
}

func (r *AuthenticationloginschemaResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationloginschemaResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationloginschema"
}

func (r *AuthenticationloginschemaResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationloginschemaResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuthenticationloginschemaResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationloginschema resource")
	// Get payload from plan
	authenticationloginschema := authenticationloginschemaGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Authenticationloginschema.Type(), name_value, &authenticationloginschema)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationloginschema, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created authenticationloginschema resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	if !r.readAuthenticationloginschemaFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "authenticationloginschema not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationloginschemaResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationloginschemaResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationloginschema resource")

	found := r.readAuthenticationloginschemaFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AuthenticationloginschemaResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AuthenticationloginschemaResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating authenticationloginschema resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Authenticationschema.Equal(state.Authenticationschema) {
		tflog.Debug(ctx, "authenticationschema has changed for authenticationloginschema")
		hasChange = true
	}
	if !data.Authenticationstrength.Equal(state.Authenticationstrength) {
		tflog.Debug(ctx, "authenticationstrength has changed for authenticationloginschema")
		hasChange = true
	}
	if !data.Passwdexpression.Equal(state.Passwdexpression) {
		tflog.Debug(ctx, "passwdexpression has changed for authenticationloginschema")
		hasChange = true
	}
	if !data.Passwordcredentialindex.Equal(state.Passwordcredentialindex) {
		tflog.Debug(ctx, "passwordcredentialindex has changed for authenticationloginschema")
		hasChange = true
	}
	if !data.Ssocredentials.Equal(state.Ssocredentials) {
		tflog.Debug(ctx, "ssocredentials has changed for authenticationloginschema")
		hasChange = true
	}
	if !data.Usercredentialindex.Equal(state.Usercredentialindex) {
		tflog.Debug(ctx, "usercredentialindex has changed for authenticationloginschema")
		hasChange = true
	}
	if !data.Userexpression.Equal(state.Userexpression) {
		tflog.Debug(ctx, "userexpression has changed for authenticationloginschema")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		authenticationloginschema := authenticationloginschemaGetThePayloadFromtheConfig(ctx, &data)
		// Make API call
		// Named resource - use UpdateResource
		name_value := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Authenticationloginschema.Type(), name_value, &authenticationloginschema)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authenticationloginschema, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated authenticationloginschema resource")
	} else {
		tflog.Debug(ctx, "No changes detected for authenticationloginschema resource, skipping update")
	}

	// Read the updated state back
	if !r.readAuthenticationloginschemaFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "authenticationloginschema not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationloginschemaResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationloginschemaResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationloginschema resource")
	// Named resource - delete using DeleteResource
	name_value := data.Name.ValueString()
	err := r.client.DeleteResource(service.Authenticationloginschema.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete authenticationloginschema, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted authenticationloginschema resource")
}

// Helper function to read authenticationloginschema data from API
func (r *AuthenticationloginschemaResource) readAuthenticationloginschemaFromApi(ctx context.Context, data *AuthenticationloginschemaResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	authenticationloginschema_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Authenticationloginschema.Type(), authenticationloginschema_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationloginschema, got error: %s", err))
		return false
	}

	authenticationloginschemaSetAttrFromGet(ctx, data, getResponseData)

	return true
}
