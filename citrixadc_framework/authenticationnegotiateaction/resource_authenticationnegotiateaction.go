package authenticationnegotiateaction

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
var _ resource.Resource = &AuthenticationnegotiateactionResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationnegotiateactionResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationnegotiateactionResource)(nil)

func NewAuthenticationnegotiateactionResource() resource.Resource {
	return &AuthenticationnegotiateactionResource{}
}

// AuthenticationnegotiateactionResource defines the resource implementation.
type AuthenticationnegotiateactionResource struct {
	client *service.NitroClient
}

func (r *AuthenticationnegotiateactionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationnegotiateactionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationnegotiateaction"
}

func (r *AuthenticationnegotiateactionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationnegotiateactionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config AuthenticationnegotiateactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationnegotiateaction resource")
	// Get payload from plan (regular attributes)
	authenticationnegotiateaction := authenticationnegotiateactionGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	authenticationnegotiateactionGetThePayloadFromtheConfig(ctx, &config, &authenticationnegotiateaction)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Authenticationnegotiateaction.Type(), name_value, &authenticationnegotiateaction)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationnegotiateaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created authenticationnegotiateaction resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	r.readAuthenticationnegotiateactionFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationnegotiateactionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationnegotiateactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationnegotiateaction resource")

	r.readAuthenticationnegotiateactionFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationnegotiateactionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state AuthenticationnegotiateactionResourceModel

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

	tflog.Debug(ctx, "Updating authenticationnegotiateaction resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Defaultauthenticationgroup.Equal(state.Defaultauthenticationgroup) {
		tflog.Debug(ctx, fmt.Sprintf("defaultauthenticationgroup has changed for authenticationnegotiateaction"))
		hasChange = true
	}
	if !data.Domain.Equal(state.Domain) {
		tflog.Debug(ctx, fmt.Sprintf("domain has changed for authenticationnegotiateaction"))
		hasChange = true
	}
	if !data.Domainuser.Equal(state.Domainuser) {
		tflog.Debug(ctx, fmt.Sprintf("domainuser has changed for authenticationnegotiateaction"))
		hasChange = true
	}
	// Check secret attribute domainuserpasswd or its version tracker
	if !data.Domainuserpasswd.Equal(state.Domainuserpasswd) {
		tflog.Debug(ctx, fmt.Sprintf("domainuserpasswd has changed for authenticationnegotiateaction"))
		hasChange = true
	} else if !data.DomainuserpasswdWoVersion.Equal(state.DomainuserpasswdWoVersion) {
		tflog.Debug(ctx, fmt.Sprintf("domainuserpasswd_wo_version has changed for authenticationnegotiateaction"))
		hasChange = true
	}
	if !data.Keytab.Equal(state.Keytab) {
		tflog.Debug(ctx, fmt.Sprintf("keytab has changed for authenticationnegotiateaction"))
		hasChange = true
	}
	if !data.Ntlmpath.Equal(state.Ntlmpath) {
		tflog.Debug(ctx, fmt.Sprintf("ntlmpath has changed for authenticationnegotiateaction"))
		hasChange = true
	}
	if !data.Ou.Equal(state.Ou) {
		tflog.Debug(ctx, fmt.Sprintf("ou has changed for authenticationnegotiateaction"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		// Get payload from plan (regular attributes)
		authenticationnegotiateaction := authenticationnegotiateactionGetThePayloadFromthePlan(ctx, &data)
		// Add write-only attributes from config to the payload
		authenticationnegotiateactionGetThePayloadFromtheConfig(ctx, &config, &authenticationnegotiateaction)
		// Make API call
		// Named resource - use UpdateResource
		name_value := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Authenticationnegotiateaction.Type(), name_value, &authenticationnegotiateaction)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authenticationnegotiateaction, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated authenticationnegotiateaction resource")
	} else {
		tflog.Debug(ctx, "No changes detected for authenticationnegotiateaction resource, skipping update")
	}

	// Read the updated state back
	r.readAuthenticationnegotiateactionFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationnegotiateactionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationnegotiateactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationnegotiateaction resource")
	// Named resource - delete using DeleteResource
	name_value := data.Name.ValueString()
	err := r.client.DeleteResource(service.Authenticationnegotiateaction.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete authenticationnegotiateaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted authenticationnegotiateaction resource")
}

// Helper function to read authenticationnegotiateaction data from API
func (r *AuthenticationnegotiateactionResource) readAuthenticationnegotiateactionFromApi(ctx context.Context, data *AuthenticationnegotiateactionResourceModel, diags *diag.Diagnostics) {

	// Case 2: Find with single ID attribute - ID is the plain value
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Authenticationnegotiateaction.Type(), name_Name)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationnegotiateaction, got error: %s", err))
		return
	}

	authenticationnegotiateactionSetAttrFromGet(ctx, data, getResponseData)

}
