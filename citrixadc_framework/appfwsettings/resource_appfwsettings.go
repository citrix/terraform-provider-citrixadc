package appfwsettings

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
var _ resource.Resource = &AppfwsettingsResource{}
var _ resource.ResourceWithConfigure = (*AppfwsettingsResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwsettingsResource)(nil)

func NewAppfwsettingsResource() resource.Resource {
	return &AppfwsettingsResource{}
}

// AppfwsettingsResource defines the resource implementation.
type AppfwsettingsResource struct {
	client *service.NitroClient
}

func (r *AppfwsettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwsettingsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwsettings"
}

func (r *AppfwsettingsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwsettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config AppfwsettingsResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwsettings resource")
	// Get payload from plan (regular attributes)
	appfwsettings := appfwsettingsGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	appfwsettingsGetThePayloadFromtheConfig(ctx, &config, &appfwsettings)

	// Make API call
	// Singleton resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Appfwsettings.Type(), &appfwsettings)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwsettings, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created appfwsettings resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue("appfwsettings-config")

	// Read the updated state back
	r.readAppfwsettingsFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwsettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwsettingsResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwsettings resource")

	r.readAppfwsettingsFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwsettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state AppfwsettingsResourceModel

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

	tflog.Debug(ctx, "Updating appfwsettings resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Ceflogging.Equal(state.Ceflogging) {
		tflog.Debug(ctx, fmt.Sprintf("ceflogging has changed for appfwsettings"))
		hasChange = true
	}
	if !data.Centralizedlearning.Equal(state.Centralizedlearning) {
		tflog.Debug(ctx, fmt.Sprintf("centralizedlearning has changed for appfwsettings"))
		hasChange = true
	}
	if !data.Clientiploggingheader.Equal(state.Clientiploggingheader) {
		tflog.Debug(ctx, fmt.Sprintf("clientiploggingheader has changed for appfwsettings"))
		hasChange = true
	}
	if !data.Cookieflags.Equal(state.Cookieflags) {
		tflog.Debug(ctx, fmt.Sprintf("cookieflags has changed for appfwsettings"))
		hasChange = true
	}
	if !data.Cookiepostencryptprefix.Equal(state.Cookiepostencryptprefix) {
		tflog.Debug(ctx, fmt.Sprintf("cookiepostencryptprefix has changed for appfwsettings"))
		hasChange = true
	}
	if !data.Defaultprofile.Equal(state.Defaultprofile) {
		tflog.Debug(ctx, fmt.Sprintf("defaultprofile has changed for appfwsettings"))
		hasChange = true
	}
	if !data.Entitydecoding.Equal(state.Entitydecoding) {
		tflog.Debug(ctx, fmt.Sprintf("entitydecoding has changed for appfwsettings"))
		hasChange = true
	}
	if !data.Geolocationlogging.Equal(state.Geolocationlogging) {
		tflog.Debug(ctx, fmt.Sprintf("geolocationlogging has changed for appfwsettings"))
		hasChange = true
	}
	if !data.Importsizelimit.Equal(state.Importsizelimit) {
		tflog.Debug(ctx, fmt.Sprintf("importsizelimit has changed for appfwsettings"))
		hasChange = true
	}
	if !data.Learnratelimit.Equal(state.Learnratelimit) {
		tflog.Debug(ctx, fmt.Sprintf("learnratelimit has changed for appfwsettings"))
		hasChange = true
	}
	if !data.Malformedreqaction.Equal(state.Malformedreqaction) {
		tflog.Debug(ctx, fmt.Sprintf("malformedreqaction has changed for appfwsettings"))
		hasChange = true
	}
	// Check secret attribute proxypassword or its version tracker
	if !data.Proxypassword.Equal(state.Proxypassword) {
		tflog.Debug(ctx, fmt.Sprintf("proxypassword has changed for appfwsettings"))
		hasChange = true
	} else if !data.ProxypasswordWoVersion.Equal(state.ProxypasswordWoVersion) {
		tflog.Debug(ctx, fmt.Sprintf("proxypassword_wo_version has changed for appfwsettings"))
		hasChange = true
	}
	if !data.Proxyport.Equal(state.Proxyport) {
		tflog.Debug(ctx, fmt.Sprintf("proxyport has changed for appfwsettings"))
		hasChange = true
	}
	if !data.Proxyserver.Equal(state.Proxyserver) {
		tflog.Debug(ctx, fmt.Sprintf("proxyserver has changed for appfwsettings"))
		hasChange = true
	}
	if !data.Proxyusername.Equal(state.Proxyusername) {
		tflog.Debug(ctx, fmt.Sprintf("proxyusername has changed for appfwsettings"))
		hasChange = true
	}
	if !data.Sessioncookiename.Equal(state.Sessioncookiename) {
		tflog.Debug(ctx, fmt.Sprintf("sessioncookiename has changed for appfwsettings"))
		hasChange = true
	}
	if !data.Sessionlifetime.Equal(state.Sessionlifetime) {
		tflog.Debug(ctx, fmt.Sprintf("sessionlifetime has changed for appfwsettings"))
		hasChange = true
	}
	if !data.Sessionlimit.Equal(state.Sessionlimit) {
		tflog.Debug(ctx, fmt.Sprintf("sessionlimit has changed for appfwsettings"))
		hasChange = true
	}
	if !data.Sessiontimeout.Equal(state.Sessiontimeout) {
		tflog.Debug(ctx, fmt.Sprintf("sessiontimeout has changed for appfwsettings"))
		hasChange = true
	}
	if !data.Signatureautoupdate.Equal(state.Signatureautoupdate) {
		tflog.Debug(ctx, fmt.Sprintf("signatureautoupdate has changed for appfwsettings"))
		hasChange = true
	}
	if !data.Signatureurl.Equal(state.Signatureurl) {
		tflog.Debug(ctx, fmt.Sprintf("signatureurl has changed for appfwsettings"))
		hasChange = true
	}
	if !data.Undefaction.Equal(state.Undefaction) {
		tflog.Debug(ctx, fmt.Sprintf("undefaction has changed for appfwsettings"))
		hasChange = true
	}
	if !data.Useconfigurablesecretkey.Equal(state.Useconfigurablesecretkey) {
		tflog.Debug(ctx, fmt.Sprintf("useconfigurablesecretkey has changed for appfwsettings"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		// Get payload from plan (regular attributes)
		appfwsettings := appfwsettingsGetThePayloadFromthePlan(ctx, &data)
		// Add write-only attributes from config to the payload
		appfwsettingsGetThePayloadFromtheConfig(ctx, &config, &appfwsettings)
		// Make API call
		// Singleton resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Appfwsettings.Type(), &appfwsettings)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appfwsettings, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated appfwsettings resource")
	} else {
		tflog.Debug(ctx, "No changes detected for appfwsettings resource, skipping update")
	}

	// Read the updated state back
	r.readAppfwsettingsFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwsettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwsettingsResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwsettings resource")
	// Singleton resource - no delete operation on ADC, just remove from state
	tflog.Trace(ctx, "Removed appfwsettings from Terraform state")
}

// Helper function to read appfwsettings data from API
func (r *AppfwsettingsResource) readAppfwsettingsFromApi(ctx context.Context, data *AppfwsettingsResourceModel, diags *diag.Diagnostics) {

	// Case 1: Simple find without ID
	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Appfwsettings.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwsettings, got error: %s", err))
		return
	}

	appfwsettingsSetAttrFromGet(ctx, data, getResponseData)

}
