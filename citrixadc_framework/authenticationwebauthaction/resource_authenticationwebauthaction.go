package authenticationwebauthaction

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
var _ resource.Resource = &AuthenticationwebauthactionResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationwebauthactionResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationwebauthactionResource)(nil)

func NewAuthenticationwebauthactionResource() resource.Resource {
	return &AuthenticationwebauthactionResource{}
}

// AuthenticationwebauthactionResource defines the resource implementation.
type AuthenticationwebauthactionResource struct {
	client *service.NitroClient
}

func (r *AuthenticationwebauthactionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationwebauthactionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationwebauthaction"
}

func (r *AuthenticationwebauthactionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationwebauthactionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuthenticationwebauthactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationwebauthaction resource")

	// Create API request body from the model
	authenticationwebauthaction := authenticationwebauthactionGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Authenticationwebauthaction.Type(), name_value, &authenticationwebauthaction)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationwebauthaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created authenticationwebauthaction resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(data.Name.ValueString())

	// Read the updated state back
	if !r.readAuthenticationwebauthactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "authenticationwebauthaction not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationwebauthactionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationwebauthactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationwebauthaction resource")

	found := r.readAuthenticationwebauthactionFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AuthenticationwebauthactionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AuthenticationwebauthactionResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating authenticationwebauthaction resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Attribute1.Equal(state.Attribute1) {
		hasChange = true
	}
	if !data.Attribute10.Equal(state.Attribute10) {
		hasChange = true
	}
	if !data.Attribute11.Equal(state.Attribute11) {
		hasChange = true
	}
	if !data.Attribute12.Equal(state.Attribute12) {
		hasChange = true
	}
	if !data.Attribute13.Equal(state.Attribute13) {
		hasChange = true
	}
	if !data.Attribute14.Equal(state.Attribute14) {
		hasChange = true
	}
	if !data.Attribute15.Equal(state.Attribute15) {
		hasChange = true
	}
	if !data.Attribute16.Equal(state.Attribute16) {
		hasChange = true
	}
	if !data.Attribute2.Equal(state.Attribute2) {
		hasChange = true
	}
	if !data.Attribute3.Equal(state.Attribute3) {
		hasChange = true
	}
	if !data.Attribute4.Equal(state.Attribute4) {
		hasChange = true
	}
	if !data.Attribute5.Equal(state.Attribute5) {
		hasChange = true
	}
	if !data.Attribute6.Equal(state.Attribute6) {
		hasChange = true
	}
	if !data.Attribute7.Equal(state.Attribute7) {
		hasChange = true
	}
	if !data.Attribute8.Equal(state.Attribute8) {
		hasChange = true
	}
	if !data.Attribute9.Equal(state.Attribute9) {
		hasChange = true
	}
	if !data.Defaultauthenticationgroup.Equal(state.Defaultauthenticationgroup) {
		hasChange = true
	}
	if !data.Fullreqexpr.Equal(state.Fullreqexpr) {
		hasChange = true
	}
	if !data.Scheme.Equal(state.Scheme) {
		hasChange = true
	}
	if !data.Serverip.Equal(state.Serverip) {
		hasChange = true
	}
	if !data.Serverport.Equal(state.Serverport) {
		hasChange = true
	}
	if !data.Successrule.Equal(state.Successrule) {
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		authenticationwebauthaction := authenticationwebauthactionGetThePayloadFromtheConfig(ctx, &data)
		// Make API call
		// Named resource - use UpdateResource
		name_value := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Authenticationwebauthaction.Type(), name_value, &authenticationwebauthaction)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authenticationwebauthaction, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated authenticationwebauthaction resource")
	} else {
		tflog.Debug(ctx, "No changes detected for authenticationwebauthaction resource, skipping update")
	}

	// Read the updated state back
	if !r.readAuthenticationwebauthactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "authenticationwebauthaction not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationwebauthactionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationwebauthactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationwebauthaction resource")

	// Named resource - delete using DeleteResource
	name_value := data.Name.ValueString()
	err := r.client.DeleteResource(service.Authenticationwebauthaction.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete authenticationwebauthaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted authenticationwebauthaction resource")
}

// Helper function to read authenticationwebauthaction data from API
func (r *AuthenticationwebauthactionResource) readAuthenticationwebauthactionFromApi(ctx context.Context, data *AuthenticationwebauthactionResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	name_Name := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Authenticationwebauthaction.Type(), name_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationwebauthaction, got error: %s", err))
		return false
	}

	authenticationwebauthactionSetAttrFromGet(ctx, data, getResponseData)

	return true
}
