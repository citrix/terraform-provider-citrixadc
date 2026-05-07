package gslbsite

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
var _ resource.Resource = &GslbsiteResource{}
var _ resource.ResourceWithConfigure = (*GslbsiteResource)(nil)
var _ resource.ResourceWithImportState = (*GslbsiteResource)(nil)

func NewGslbsiteResource() resource.Resource {
	return &GslbsiteResource{}
}

// GslbsiteResource defines the resource implementation.
type GslbsiteResource struct {
	client *service.NitroClient
}

func (r *GslbsiteResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *GslbsiteResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gslbsite"
}

func (r *GslbsiteResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *GslbsiteResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config GslbsiteResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating gslbsite resource")
	// Get payload from plan (regular attributes)
	gslbsite := gslbsiteGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	gslbsiteGetThePayloadFromtheConfig(ctx, &config, &gslbsite)

	// Make API call
	// Named resource - use AddResource
	sitename_value := data.Sitename.ValueString()
	_, err := r.client.AddResource(service.Gslbsite.Type(), sitename_value, &gslbsite)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create gslbsite, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created gslbsite resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Sitename.ValueString()))

	// Read the updated state back
	r.readGslbsiteFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbsiteResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data GslbsiteResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading gslbsite resource")

	r.readGslbsiteFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbsiteResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state GslbsiteResourceModel

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

	tflog.Debug(ctx, "Updating gslbsite resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Backupparentlist.Equal(state.Backupparentlist) {
		tflog.Debug(ctx, fmt.Sprintf("backupparentlist has changed for gslbsite"))
		hasChange = true
	}
	if !data.Metricexchange.Equal(state.Metricexchange) {
		tflog.Debug(ctx, fmt.Sprintf("metricexchange has changed for gslbsite"))
		hasChange = true
	}
	if !data.Naptrreplacementsuffix.Equal(state.Naptrreplacementsuffix) {
		tflog.Debug(ctx, fmt.Sprintf("naptrreplacementsuffix has changed for gslbsite"))
		hasChange = true
	}
	if !data.Nwmetricexchange.Equal(state.Nwmetricexchange) {
		tflog.Debug(ctx, fmt.Sprintf("nwmetricexchange has changed for gslbsite"))
		hasChange = true
	}
	if !data.Parentsite.Equal(state.Parentsite) {
		tflog.Debug(ctx, fmt.Sprintf("parentsite has changed for gslbsite"))
		hasChange = true
	}
	if !data.Publicip.Equal(state.Publicip) {
		tflog.Debug(ctx, fmt.Sprintf("publicip has changed for gslbsite"))
		hasChange = true
	}
	if !data.Sessionexchange.Equal(state.Sessionexchange) {
		tflog.Debug(ctx, fmt.Sprintf("sessionexchange has changed for gslbsite"))
		hasChange = true
	}
	if !data.Siteipaddress.Equal(state.Siteipaddress) {
		tflog.Debug(ctx, fmt.Sprintf("siteipaddress has changed for gslbsite"))
		hasChange = true
	}
	// Check secret attribute sitepassword or its version tracker
	if !data.Sitepassword.Equal(state.Sitepassword) {
		tflog.Debug(ctx, fmt.Sprintf("sitepassword has changed for gslbsite"))
		hasChange = true
	} else if !data.SitepasswordWoVersion.Equal(state.SitepasswordWoVersion) {
		tflog.Debug(ctx, fmt.Sprintf("sitepassword_wo_version has changed for gslbsite"))
		hasChange = true
	}
	if !data.Triggermonitor.Equal(state.Triggermonitor) {
		tflog.Debug(ctx, fmt.Sprintf("triggermonitor has changed for gslbsite"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		// Get payload from plan (regular attributes)
		gslbsite := gslbsiteGetThePayloadFromthePlan(ctx, &data)
		// Add write-only attributes from config to the payload
		gslbsiteGetThePayloadFromtheConfig(ctx, &config, &gslbsite)
		// Make API call
		// Named resource - use UpdateResource
		sitename_value := data.Sitename.ValueString()
		_, err := r.client.UpdateResource(service.Gslbsite.Type(), sitename_value, &gslbsite)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update gslbsite, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated gslbsite resource")
	} else {
		tflog.Debug(ctx, "No changes detected for gslbsite resource, skipping update")
	}

	// Read the updated state back
	r.readGslbsiteFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbsiteResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data GslbsiteResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting gslbsite resource")
	// Named resource - delete using DeleteResource
	sitename_value := data.Sitename.ValueString()
	err := r.client.DeleteResource(service.Gslbsite.Type(), sitename_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete gslbsite, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted gslbsite resource")
}

// Helper function to read gslbsite data from API
func (r *GslbsiteResource) readGslbsiteFromApi(ctx context.Context, data *GslbsiteResourceModel, diags *diag.Diagnostics) {

	// Case 2: Find with single ID attribute - ID is the plain value
	sitename_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Gslbsite.Type(), sitename_Name)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read gslbsite, got error: %s", err))
		return
	}

	gslbsiteSetAttrFromGet(ctx, data, getResponseData)

}
