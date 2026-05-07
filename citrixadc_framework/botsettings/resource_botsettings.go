package botsettings

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
var _ resource.Resource = &BotsettingsResource{}
var _ resource.ResourceWithConfigure = (*BotsettingsResource)(nil)
var _ resource.ResourceWithImportState = (*BotsettingsResource)(nil)

func NewBotsettingsResource() resource.Resource {
	return &BotsettingsResource{}
}

// BotsettingsResource defines the resource implementation.
type BotsettingsResource struct {
	client *service.NitroClient
}

func (r *BotsettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *BotsettingsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_botsettings"
}

func (r *BotsettingsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *BotsettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config BotsettingsResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating botsettings resource")
	// Get payload from plan (regular attributes)
	botsettings := botsettingsGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	botsettingsGetThePayloadFromtheConfig(ctx, &config, &botsettings)

	// Make API call
	// Singleton resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Botsettings.Type(), &botsettings)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create botsettings, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created botsettings resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue("botsettings-config")

	// Read the updated state back
	r.readBotsettingsFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BotsettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data BotsettingsResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading botsettings resource")

	r.readBotsettingsFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BotsettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state BotsettingsResourceModel

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

	tflog.Debug(ctx, "Updating botsettings resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Defaultnonintrusiveprofile.Equal(state.Defaultnonintrusiveprofile) {
		tflog.Debug(ctx, fmt.Sprintf("defaultnonintrusiveprofile has changed for botsettings"))
		hasChange = true
	}
	if !data.Defaultprofile.Equal(state.Defaultprofile) {
		tflog.Debug(ctx, fmt.Sprintf("defaultprofile has changed for botsettings"))
		hasChange = true
	}
	if !data.Dfprequestlimit.Equal(state.Dfprequestlimit) {
		tflog.Debug(ctx, fmt.Sprintf("dfprequestlimit has changed for botsettings"))
		hasChange = true
	}
	if !data.Javascriptname.Equal(state.Javascriptname) {
		tflog.Debug(ctx, fmt.Sprintf("javascriptname has changed for botsettings"))
		hasChange = true
	}
	// Check secret attribute proxypassword or its version tracker
	if !data.Proxypassword.Equal(state.Proxypassword) {
		tflog.Debug(ctx, fmt.Sprintf("proxypassword has changed for botsettings"))
		hasChange = true
	} else if !data.ProxypasswordWoVersion.Equal(state.ProxypasswordWoVersion) {
		tflog.Debug(ctx, fmt.Sprintf("proxypassword_wo_version has changed for botsettings"))
		hasChange = true
	}
	if !data.Proxyport.Equal(state.Proxyport) {
		tflog.Debug(ctx, fmt.Sprintf("proxyport has changed for botsettings"))
		hasChange = true
	}
	if !data.Proxyserver.Equal(state.Proxyserver) {
		tflog.Debug(ctx, fmt.Sprintf("proxyserver has changed for botsettings"))
		hasChange = true
	}
	if !data.Proxyusername.Equal(state.Proxyusername) {
		tflog.Debug(ctx, fmt.Sprintf("proxyusername has changed for botsettings"))
		hasChange = true
	}
	if !data.Sessioncookiename.Equal(state.Sessioncookiename) {
		tflog.Debug(ctx, fmt.Sprintf("sessioncookiename has changed for botsettings"))
		hasChange = true
	}
	if !data.Sessiontimeout.Equal(state.Sessiontimeout) {
		tflog.Debug(ctx, fmt.Sprintf("sessiontimeout has changed for botsettings"))
		hasChange = true
	}
	if !data.Signatureautoupdate.Equal(state.Signatureautoupdate) {
		tflog.Debug(ctx, fmt.Sprintf("signatureautoupdate has changed for botsettings"))
		hasChange = true
	}
	if !data.Signatureurl.Equal(state.Signatureurl) {
		tflog.Debug(ctx, fmt.Sprintf("signatureurl has changed for botsettings"))
		hasChange = true
	}
	if !data.Trapurlautogenerate.Equal(state.Trapurlautogenerate) {
		tflog.Debug(ctx, fmt.Sprintf("trapurlautogenerate has changed for botsettings"))
		hasChange = true
	}
	if !data.Trapurlinterval.Equal(state.Trapurlinterval) {
		tflog.Debug(ctx, fmt.Sprintf("trapurlinterval has changed for botsettings"))
		hasChange = true
	}
	if !data.Trapurllength.Equal(state.Trapurllength) {
		tflog.Debug(ctx, fmt.Sprintf("trapurllength has changed for botsettings"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		// Get payload from plan (regular attributes)
		botsettings := botsettingsGetThePayloadFromthePlan(ctx, &data)
		// Add write-only attributes from config to the payload
		botsettingsGetThePayloadFromtheConfig(ctx, &config, &botsettings)
		// Make API call
		// Singleton resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Botsettings.Type(), &botsettings)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update botsettings, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated botsettings resource")
	} else {
		tflog.Debug(ctx, "No changes detected for botsettings resource, skipping update")
	}

	// Read the updated state back
	r.readBotsettingsFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BotsettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data BotsettingsResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting botsettings resource")
	// Singleton resource - no delete operation on ADC, just remove from state
	tflog.Trace(ctx, "Removed botsettings from Terraform state")
}

// Helper function to read botsettings data from API
func (r *BotsettingsResource) readBotsettingsFromApi(ctx context.Context, data *BotsettingsResourceModel, diags *diag.Diagnostics) {

	// Case 1: Simple find without ID
	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Botsettings.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read botsettings, got error: %s", err))
		return
	}

	botsettingsSetAttrFromGet(ctx, data, getResponseData)

}
