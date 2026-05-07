package authenticationpushservice

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
var _ resource.Resource = &AuthenticationpushserviceResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationpushserviceResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationpushserviceResource)(nil)

func NewAuthenticationpushserviceResource() resource.Resource {
	return &AuthenticationpushserviceResource{}
}

// AuthenticationpushserviceResource defines the resource implementation.
type AuthenticationpushserviceResource struct {
	client *service.NitroClient
}

func (r *AuthenticationpushserviceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationpushserviceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationpushservice"
}

func (r *AuthenticationpushserviceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationpushserviceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config AuthenticationpushserviceResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationpushservice resource")
	// Get payload from plan (regular attributes)
	authenticationpushservice := authenticationpushserviceGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	authenticationpushserviceGetThePayloadFromtheConfig(ctx, &config, &authenticationpushservice)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Authenticationpushservice.Type(), name_value, &authenticationpushservice)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationpushservice, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created authenticationpushservice resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	r.readAuthenticationpushserviceFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationpushserviceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationpushserviceResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationpushservice resource")

	r.readAuthenticationpushserviceFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationpushserviceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state AuthenticationpushserviceResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating authenticationpushservice resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Clientid.Equal(state.Clientid) {
		tflog.Debug(ctx, fmt.Sprintf("clientid has changed for authenticationpushservice"))
		hasChange = true
	}
	// Check secret attribute clientsecret or its version tracker
	if !data.Clientsecret.Equal(state.Clientsecret) {
		tflog.Debug(ctx, fmt.Sprintf("clientsecret has changed for authenticationpushservice"))
		hasChange = true
	} else if !data.ClientsecretWoVersion.Equal(state.ClientsecretWoVersion) {
		tflog.Debug(ctx, fmt.Sprintf("clientsecret_wo_version has changed for authenticationpushservice"))
		hasChange = true
	}
	if !data.Customerid.Equal(state.Customerid) {
		tflog.Debug(ctx, fmt.Sprintf("customerid has changed for authenticationpushservice"))
		hasChange = true
	}
	if !data.Refreshinterval.Equal(state.Refreshinterval) {
		tflog.Debug(ctx, fmt.Sprintf("refreshinterval has changed for authenticationpushservice"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		// Get payload from plan (regular attributes)
		authenticationpushservice := authenticationpushserviceGetThePayloadFromthePlan(ctx, &data)
		// Add write-only attributes from config to the payload
		authenticationpushserviceGetThePayloadFromtheConfig(ctx, &config, &authenticationpushservice)
		// Make API call
		// Named resource - use UpdateResource
		name_value := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Authenticationpushservice.Type(), name_value, &authenticationpushservice)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authenticationpushservice, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated authenticationpushservice resource")
	} else {
		tflog.Debug(ctx, "No changes detected for authenticationpushservice resource, skipping update")
	}

	// Read the updated state back
	r.readAuthenticationpushserviceFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationpushserviceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationpushserviceResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationpushservice resource")
	// Named resource - delete using DeleteResource
	name_value := data.Name.ValueString()
	err := r.client.DeleteResource(service.Authenticationpushservice.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete authenticationpushservice, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted authenticationpushservice resource")
}

// Helper function to read authenticationpushservice data from API
func (r *AuthenticationpushserviceResource) readAuthenticationpushserviceFromApi(ctx context.Context, data *AuthenticationpushserviceResourceModel, diags *diag.Diagnostics) {

	// Case 2: Find with single ID attribute - ID is the plain value
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Authenticationpushservice.Type(), name_Name)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationpushservice, got error: %s", err))
		return
	}

	authenticationpushserviceSetAttrFromGet(ctx, data, getResponseData)

}
