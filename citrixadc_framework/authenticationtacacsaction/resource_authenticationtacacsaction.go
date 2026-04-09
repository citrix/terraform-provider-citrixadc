package authenticationtacacsaction

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
var _ resource.Resource = &AuthenticationtacacsactionResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationtacacsactionResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationtacacsactionResource)(nil)

func NewAuthenticationtacacsactionResource() resource.Resource {
	return &AuthenticationtacacsactionResource{}
}

// AuthenticationtacacsactionResource defines the resource implementation.
type AuthenticationtacacsactionResource struct {
	client *service.NitroClient
}

func (r *AuthenticationtacacsactionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationtacacsactionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationtacacsaction"
}

func (r *AuthenticationtacacsactionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationtacacsactionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config AuthenticationtacacsactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationtacacsaction resource")
	// Get payload from plan (regular attributes)
	authenticationtacacsaction := authenticationtacacsactionGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	authenticationtacacsactionGetThePayloadFromtheConfig(ctx, &config, &authenticationtacacsaction)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Authenticationtacacsaction.Type(), name_value, &authenticationtacacsaction)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationtacacsaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created authenticationtacacsaction resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	r.readAuthenticationtacacsactionFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationtacacsactionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationtacacsactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationtacacsaction resource")

	r.readAuthenticationtacacsactionFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationtacacsactionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state AuthenticationtacacsactionResourceModel

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

	tflog.Debug(ctx, "Updating authenticationtacacsaction resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Accounting.Equal(state.Accounting) {
		tflog.Debug(ctx, fmt.Sprintf("accounting has changed for authenticationtacacsaction"))
		hasChange = true
	}
	if !data.Attribute1.Equal(state.Attribute1) {
		tflog.Debug(ctx, fmt.Sprintf("attribute1 has changed for authenticationtacacsaction"))
		hasChange = true
	}
	if !data.Attribute10.Equal(state.Attribute10) {
		tflog.Debug(ctx, fmt.Sprintf("attribute10 has changed for authenticationtacacsaction"))
		hasChange = true
	}
	if !data.Attribute11.Equal(state.Attribute11) {
		tflog.Debug(ctx, fmt.Sprintf("attribute11 has changed for authenticationtacacsaction"))
		hasChange = true
	}
	if !data.Attribute12.Equal(state.Attribute12) {
		tflog.Debug(ctx, fmt.Sprintf("attribute12 has changed for authenticationtacacsaction"))
		hasChange = true
	}
	if !data.Attribute13.Equal(state.Attribute13) {
		tflog.Debug(ctx, fmt.Sprintf("attribute13 has changed for authenticationtacacsaction"))
		hasChange = true
	}
	if !data.Attribute14.Equal(state.Attribute14) {
		tflog.Debug(ctx, fmt.Sprintf("attribute14 has changed for authenticationtacacsaction"))
		hasChange = true
	}
	if !data.Attribute15.Equal(state.Attribute15) {
		tflog.Debug(ctx, fmt.Sprintf("attribute15 has changed for authenticationtacacsaction"))
		hasChange = true
	}
	if !data.Attribute16.Equal(state.Attribute16) {
		tflog.Debug(ctx, fmt.Sprintf("attribute16 has changed for authenticationtacacsaction"))
		hasChange = true
	}
	if !data.Attribute2.Equal(state.Attribute2) {
		tflog.Debug(ctx, fmt.Sprintf("attribute2 has changed for authenticationtacacsaction"))
		hasChange = true
	}
	if !data.Attribute3.Equal(state.Attribute3) {
		tflog.Debug(ctx, fmt.Sprintf("attribute3 has changed for authenticationtacacsaction"))
		hasChange = true
	}
	if !data.Attribute4.Equal(state.Attribute4) {
		tflog.Debug(ctx, fmt.Sprintf("attribute4 has changed for authenticationtacacsaction"))
		hasChange = true
	}
	if !data.Attribute5.Equal(state.Attribute5) {
		tflog.Debug(ctx, fmt.Sprintf("attribute5 has changed for authenticationtacacsaction"))
		hasChange = true
	}
	if !data.Attribute6.Equal(state.Attribute6) {
		tflog.Debug(ctx, fmt.Sprintf("attribute6 has changed for authenticationtacacsaction"))
		hasChange = true
	}
	if !data.Attribute7.Equal(state.Attribute7) {
		tflog.Debug(ctx, fmt.Sprintf("attribute7 has changed for authenticationtacacsaction"))
		hasChange = true
	}
	if !data.Attribute8.Equal(state.Attribute8) {
		tflog.Debug(ctx, fmt.Sprintf("attribute8 has changed for authenticationtacacsaction"))
		hasChange = true
	}
	if !data.Attribute9.Equal(state.Attribute9) {
		tflog.Debug(ctx, fmt.Sprintf("attribute9 has changed for authenticationtacacsaction"))
		hasChange = true
	}
	if !data.Attributes.Equal(state.Attributes) {
		tflog.Debug(ctx, fmt.Sprintf("attributes has changed for authenticationtacacsaction"))
		hasChange = true
	}
	if !data.Auditfailedcmds.Equal(state.Auditfailedcmds) {
		tflog.Debug(ctx, fmt.Sprintf("auditfailedcmds has changed for authenticationtacacsaction"))
		hasChange = true
	}
	if !data.Authorization.Equal(state.Authorization) {
		tflog.Debug(ctx, fmt.Sprintf("authorization has changed for authenticationtacacsaction"))
		hasChange = true
	}
	if !data.Authtimeout.Equal(state.Authtimeout) {
		tflog.Debug(ctx, fmt.Sprintf("authtimeout has changed for authenticationtacacsaction"))
		hasChange = true
	}
	if !data.Defaultauthenticationgroup.Equal(state.Defaultauthenticationgroup) {
		tflog.Debug(ctx, fmt.Sprintf("defaultauthenticationgroup has changed for authenticationtacacsaction"))
		hasChange = true
	}
	if !data.Groupattrname.Equal(state.Groupattrname) {
		tflog.Debug(ctx, fmt.Sprintf("groupattrname has changed for authenticationtacacsaction"))
		hasChange = true
	}
	if !data.Serverip.Equal(state.Serverip) {
		tflog.Debug(ctx, fmt.Sprintf("serverip has changed for authenticationtacacsaction"))
		hasChange = true
	}
	if !data.Serverport.Equal(state.Serverport) {
		tflog.Debug(ctx, fmt.Sprintf("serverport has changed for authenticationtacacsaction"))
		hasChange = true
	}
	// Check secret attribute tacacssecret or its version tracker
	if !data.Tacacssecret.Equal(state.Tacacssecret) {
		tflog.Debug(ctx, fmt.Sprintf("tacacssecret has changed for authenticationtacacsaction"))
		hasChange = true
	} else if !data.TacacssecretWoVersion.Equal(state.TacacssecretWoVersion) {
		tflog.Debug(ctx, fmt.Sprintf("tacacssecret_wo_version has changed for authenticationtacacsaction"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		// Get payload from plan (regular attributes)
		authenticationtacacsaction := authenticationtacacsactionGetThePayloadFromthePlan(ctx, &data)
		// Add write-only attributes from config to the payload
		authenticationtacacsactionGetThePayloadFromtheConfig(ctx, &config, &authenticationtacacsaction)
		// Make API call
		// Named resource - use UpdateResource
		name_value := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Authenticationtacacsaction.Type(), name_value, &authenticationtacacsaction)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authenticationtacacsaction, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated authenticationtacacsaction resource")
	} else {
		tflog.Debug(ctx, "No changes detected for authenticationtacacsaction resource, skipping update")
	}

	// Read the updated state back
	r.readAuthenticationtacacsactionFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationtacacsactionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationtacacsactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationtacacsaction resource")
	// Named resource - delete using DeleteResource
	name_value := data.Name.ValueString()
	err := r.client.DeleteResource(service.Authenticationtacacsaction.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete authenticationtacacsaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted authenticationtacacsaction resource")
}

// Helper function to read authenticationtacacsaction data from API
func (r *AuthenticationtacacsactionResource) readAuthenticationtacacsactionFromApi(ctx context.Context, data *AuthenticationtacacsactionResourceModel, diags *diag.Diagnostics) {

	// Case 2: Find with single ID attribute - ID is the plain value
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Authenticationtacacsaction.Type(), name_Name)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationtacacsaction, got error: %s", err))
		return
	}

	authenticationtacacsactionSetAttrFromGet(ctx, data, getResponseData)

}
