package authenticationstorefrontauthaction

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
var _ resource.Resource = &AuthenticationstorefrontauthactionResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationstorefrontauthactionResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationstorefrontauthactionResource)(nil)

func NewAuthenticationstorefrontauthactionResource() resource.Resource {
	return &AuthenticationstorefrontauthactionResource{}
}

// AuthenticationstorefrontauthactionResource defines the resource implementation.
type AuthenticationstorefrontauthactionResource struct {
	client *service.NitroClient
}

func (r *AuthenticationstorefrontauthactionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationstorefrontauthactionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationstorefrontauthaction"
}

func (r *AuthenticationstorefrontauthactionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationstorefrontauthactionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuthenticationstorefrontauthactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationstorefrontauthaction resource")

	authenticationstorefrontauthaction := authenticationstorefrontauthactionGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Authenticationstorefrontauthaction.Type(), name_value, &authenticationstorefrontauthaction)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationstorefrontauthaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created authenticationstorefrontauthaction resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(data.Name.ValueString())

	// Read the updated state back
	if !r.readAuthenticationstorefrontauthactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "authenticationstorefrontauthaction not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationstorefrontauthactionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationstorefrontauthactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationstorefrontauthaction resource")

	found := r.readAuthenticationstorefrontauthactionFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AuthenticationstorefrontauthactionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state AuthenticationstorefrontauthactionResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from configuration (to unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating authenticationstorefrontauthaction resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	attributesToUnset := []string{}
	if !data.Serverurl.Equal(state.Serverurl) {
		tflog.Debug(ctx, "serverurl has changed for authenticationstorefrontauthaction")
		hasChange = true
	}
	if !data.Domain.Equal(state.Domain) {
		tflog.Debug(ctx, "domain has changed for authenticationstorefrontauthaction")
		if config.Domain.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "domain")
		} else {
			hasChange = true
		}
	}
	if !data.Defaultauthenticationgroup.Equal(state.Defaultauthenticationgroup) {
		tflog.Debug(ctx, "defaultauthenticationgroup has changed for authenticationstorefrontauthaction")
		if config.Defaultauthenticationgroup.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "defaultauthenticationgroup")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		// Create API request body from the model
		authenticationstorefrontauthaction := authenticationstorefrontauthactionGetTheUpdatablePayloadFromThePlan(ctx, &data)
		// Make API call
		// Named resource - use UpdateResource
		name_value := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Authenticationstorefrontauthaction.Type(), name_value, &authenticationstorefrontauthaction)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authenticationstorefrontauthaction, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated authenticationstorefrontauthaction resource")
	} else {
		tflog.Debug(ctx, "No changes detected for authenticationstorefrontauthaction resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults.
	unsetIdPayload := map[string]interface{}{
		"name": data.Name.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Authenticationstorefrontauthaction.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset authenticationstorefrontauthaction attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readAuthenticationstorefrontauthactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "authenticationstorefrontauthaction not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationstorefrontauthactionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationstorefrontauthactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationstorefrontauthaction resource")
	// Named resource - delete using DeleteResource
	name_value := data.Name.ValueString()
	err := r.client.DeleteResource(service.Authenticationstorefrontauthaction.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete authenticationstorefrontauthaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted authenticationstorefrontauthaction resource")
}

// Helper function to read authenticationstorefrontauthaction data from API
func (r *AuthenticationstorefrontauthactionResource) readAuthenticationstorefrontauthactionFromApi(ctx context.Context, data *AuthenticationstorefrontauthactionResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	authenticationstorefrontauthaction_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Authenticationstorefrontauthaction.Type(), authenticationstorefrontauthaction_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationstorefrontauthaction, got error: %s", err))
		return false
	}

	authenticationstorefrontauthactionSetAttrFromGet(ctx, data, getResponseData)

	return true
}
