package autoscaleprofile

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
var _ resource.Resource = &AutoscaleprofileResource{}
var _ resource.ResourceWithConfigure = (*AutoscaleprofileResource)(nil)
var _ resource.ResourceWithImportState = (*AutoscaleprofileResource)(nil)
var _ resource.ResourceWithValidateConfig = (*AutoscaleprofileResource)(nil)

func NewAutoscaleprofileResource() resource.Resource {
	return &AutoscaleprofileResource{}
}

// AutoscaleprofileResource defines the resource implementation.
type AutoscaleprofileResource struct {
	client *service.NitroClient
}

func (r *AutoscaleprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AutoscaleprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_autoscaleprofile"
}

func (r *AutoscaleprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AutoscaleprofileResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data AutoscaleprofileResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Validate that either apikey or apikey_wo is specified
	if data.Apikey.IsNull() && data.ApikeyWo.IsNull() {
		resp.Diagnostics.AddAttributeError(
			path.Root("apikey"),
			"Missing Required Attribute",
			"Either \"apikey\" or \"apikey_wo\" must be specified.",
		)
	}

	// Validate that either sharedsecret or sharedsecret_wo is specified
	if data.Sharedsecret.IsNull() && data.SharedsecretWo.IsNull() {
		resp.Diagnostics.AddAttributeError(
			path.Root("sharedsecret"),
			"Missing Required Attribute",
			"Either \"sharedsecret\" or \"sharedsecret_wo\" must be specified.",
		)
	}
}

func (r *AutoscaleprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config AutoscaleprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating autoscaleprofile resource")
	// Get payload from plan (regular attributes)
	autoscaleprofile := autoscaleprofileGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	autoscaleprofileGetThePayloadFromtheConfig(ctx, &config, &autoscaleprofile)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Autoscaleprofile.Type(), name_value, &autoscaleprofile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create autoscaleprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created autoscaleprofile resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	r.readAutoscaleprofileFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AutoscaleprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AutoscaleprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading autoscaleprofile resource")

	r.readAutoscaleprofileFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AutoscaleprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state AutoscaleprofileResourceModel

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

	tflog.Debug(ctx, "Updating autoscaleprofile resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	// Check secret attribute apikey or its version tracker
	if !data.Apikey.Equal(state.Apikey) {
		tflog.Debug(ctx, fmt.Sprintf("apikey has changed for autoscaleprofile"))
		hasChange = true
	} else if !data.ApikeyWoVersion.Equal(state.ApikeyWoVersion) {
		tflog.Debug(ctx, fmt.Sprintf("apikey_wo_version has changed for autoscaleprofile"))
		hasChange = true
	}
	// Check secret attribute sharedsecret or its version tracker
	if !data.Sharedsecret.Equal(state.Sharedsecret) {
		tflog.Debug(ctx, fmt.Sprintf("sharedsecret has changed for autoscaleprofile"))
		hasChange = true
	} else if !data.SharedsecretWoVersion.Equal(state.SharedsecretWoVersion) {
		tflog.Debug(ctx, fmt.Sprintf("sharedsecret_wo_version has changed for autoscaleprofile"))
		hasChange = true
	}
	if !data.Url.Equal(state.Url) {
		tflog.Debug(ctx, fmt.Sprintf("url has changed for autoscaleprofile"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		// Get payload from plan (regular attributes)
		autoscaleprofile := autoscaleprofileGetThePayloadFromthePlan(ctx, &data)
		// Add write-only attributes from config to the payload
		autoscaleprofileGetThePayloadFromtheConfig(ctx, &config, &autoscaleprofile)

		autoscaleprofile.Type = ""
		// Make API call
		// Named resource - use UpdateResource
		name_value := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Autoscaleprofile.Type(), name_value, &autoscaleprofile)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update autoscaleprofile, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated autoscaleprofile resource")
	} else {
		tflog.Debug(ctx, "No changes detected for autoscaleprofile resource, skipping update")
	}

	// Read the updated state back
	r.readAutoscaleprofileFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AutoscaleprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AutoscaleprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting autoscaleprofile resource")
	// Named resource - delete using DeleteResource
	name_value := data.Name.ValueString()
	err := r.client.DeleteResource(service.Autoscaleprofile.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete autoscaleprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted autoscaleprofile resource")
}

// Helper function to read autoscaleprofile data from API
func (r *AutoscaleprofileResource) readAutoscaleprofileFromApi(ctx context.Context, data *AutoscaleprofileResourceModel, diags *diag.Diagnostics) {

	// Case 2: Find with single ID attribute - ID is the plain value
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Autoscaleprofile.Type(), name_Name)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read autoscaleprofile, got error: %s", err))
		return
	}

	autoscaleprofileSetAttrFromGet(ctx, data, getResponseData)

}
