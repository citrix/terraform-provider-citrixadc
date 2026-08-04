package authenticationepaaction

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
var _ resource.Resource = &AuthenticationepaactionResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationepaactionResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationepaactionResource)(nil)

func NewAuthenticationepaactionResource() resource.Resource {
	return &AuthenticationepaactionResource{}
}

// AuthenticationepaactionResource defines the resource implementation.
type AuthenticationepaactionResource struct {
	client *service.NitroClient
}

func (r *AuthenticationepaactionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationepaactionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationepaaction"
}

func (r *AuthenticationepaactionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationepaactionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuthenticationepaactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationepaaction resource")
	// Get payload from plan
	authenticationepaaction := authenticationepaactionGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Authenticationepaaction.Type(), name_value, &authenticationepaaction)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationepaaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created authenticationepaaction resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	if !r.readAuthenticationepaactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "authenticationepaaction not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationepaactionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationepaactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationepaaction resource")

	found := r.readAuthenticationepaactionFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AuthenticationepaactionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AuthenticationepaactionResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating authenticationepaaction resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Csecexpr.Equal(state.Csecexpr) {
		tflog.Debug(ctx, "csecexpr has changed for authenticationepaaction")
		hasChange = true
	}
	if !data.Defaultepagroup.Equal(state.Defaultepagroup) {
		tflog.Debug(ctx, "defaultepagroup has changed for authenticationepaaction")
		hasChange = true
	}
	if !data.Deletefiles.Equal(state.Deletefiles) {
		tflog.Debug(ctx, "deletefiles has changed for authenticationepaaction")
		hasChange = true
	}
	if !data.Deviceposture.Equal(state.Deviceposture) {
		tflog.Debug(ctx, "deviceposture has changed for authenticationepaaction")
		hasChange = true
	}
	if !data.Killprocess.Equal(state.Killprocess) {
		tflog.Debug(ctx, "killprocess has changed for authenticationepaaction")
		hasChange = true
	}
	if !data.Quarantinegroup.Equal(state.Quarantinegroup) {
		tflog.Debug(ctx, "quarantinegroup has changed for authenticationepaaction")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		authenticationepaaction := authenticationepaactionGetTheUpdatablePayloadFromThePlan(ctx, &data)
		// Make API call
		// Named resource - use UpdateResource
		name_value := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Authenticationepaaction.Type(), name_value, &authenticationepaaction)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authenticationepaaction, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated authenticationepaaction resource")
	} else {
		tflog.Debug(ctx, "No changes detected for authenticationepaaction resource, skipping update")
	}

	// Read the updated state back
	if !r.readAuthenticationepaactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "authenticationepaaction not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationepaactionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationepaactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationepaaction resource")
	// Named resource - delete using DeleteResource
	name_value := data.Name.ValueString()
	err := r.client.DeleteResource(service.Authenticationepaaction.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete authenticationepaaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted authenticationepaaction resource")
}

// Helper function to read authenticationepaaction data from API
func (r *AuthenticationepaactionResource) readAuthenticationepaactionFromApi(ctx context.Context, data *AuthenticationepaactionResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	authenticationepaaction_Name := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Authenticationepaaction.Type(), authenticationepaaction_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationepaaction, got error: %s", err))
		return false
	}

	authenticationepaactionSetAttrFromGet(ctx, data, getResponseData)

	return true
}
