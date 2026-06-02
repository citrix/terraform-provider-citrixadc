package metricsprofile

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
var _ resource.Resource = &MetricsprofileResource{}
var _ resource.ResourceWithConfigure = (*MetricsprofileResource)(nil)
var _ resource.ResourceWithImportState = (*MetricsprofileResource)(nil)

func NewMetricsprofileResource() resource.Resource {
	return &MetricsprofileResource{}
}

// MetricsprofileResource defines the resource implementation.
type MetricsprofileResource struct {
	client *service.NitroClient
}

func (r *MetricsprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *MetricsprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_metricsprofile"
}

func (r *MetricsprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *MetricsprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config MetricsprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating metricsprofile resource")
	// Get payload from plan (regular attributes)
	metricsprofile := metricsprofileGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	metricsprofileGetThePayloadFromtheConfig(ctx, &config, metricsprofile)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Metricsprofile.Type(), name_value, &metricsprofile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create metricsprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created metricsprofile resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	r.readMetricsprofileFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MetricsprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data MetricsprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading metricsprofile resource")

	r.readMetricsprofileFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MetricsprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state MetricsprofileResourceModel

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

	tflog.Debug(ctx, "Updating metricsprofile resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Collector.Equal(state.Collector) {
		tflog.Debug(ctx, fmt.Sprintf("collector has changed for metricsprofile"))
		hasChange = true
	}
	if !data.Metrics.Equal(state.Metrics) {
		tflog.Debug(ctx, fmt.Sprintf("metrics has changed for metricsprofile"))
		hasChange = true
	}
	// Check secret attribute metricsauthtoken or its version tracker
	if !data.Metricsauthtoken.Equal(state.Metricsauthtoken) {
		tflog.Debug(ctx, fmt.Sprintf("metricsauthtoken has changed for metricsprofile"))
		hasChange = true
	} else if !data.MetricsauthtokenWoVersion.Equal(state.MetricsauthtokenWoVersion) {
		tflog.Debug(ctx, fmt.Sprintf("metricsauthtoken_wo_version has changed for metricsprofile"))
		hasChange = true
	}
	if !data.Metricsendpointurl.Equal(state.Metricsendpointurl) {
		tflog.Debug(ctx, fmt.Sprintf("metricsendpointurl has changed for metricsprofile"))
		hasChange = true
	}
	if !data.Metricsexportfrequency.Equal(state.Metricsexportfrequency) {
		tflog.Debug(ctx, fmt.Sprintf("metricsexportfrequency has changed for metricsprofile"))
		hasChange = true
	}
	if !data.Outputmode.Equal(state.Outputmode) {
		tflog.Debug(ctx, fmt.Sprintf("outputmode has changed for metricsprofile"))
		hasChange = true
	}
	if !data.Schemafile.Equal(state.Schemafile) {
		tflog.Debug(ctx, fmt.Sprintf("schemafile has changed for metricsprofile"))
		hasChange = true
	}
	if !data.Servemode.Equal(state.Servemode) {
		tflog.Debug(ctx, fmt.Sprintf("servemode has changed for metricsprofile"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		// Get payload from plan (regular attributes)
		metricsprofile := metricsprofileGetThePayloadFromthePlan(ctx, &data)
		// Add write-only attributes from config to the payload
		metricsprofileGetThePayloadFromtheConfig(ctx, &config, metricsprofile)
		// Make API call
		// Named resource - use UpdateResource
		name_value := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Metricsprofile.Type(), name_value, &metricsprofile)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update metricsprofile, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated metricsprofile resource")
	} else {
		tflog.Debug(ctx, "No changes detected for metricsprofile resource, skipping update")
	}

	// Read the updated state back
	r.readMetricsprofileFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MetricsprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data MetricsprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting metricsprofile resource")
	// Named resource - delete using DeleteResource
	name_value := data.Name.ValueString()
	err := r.client.DeleteResource(service.Metricsprofile.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete metricsprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted metricsprofile resource")
}

// Helper function to read metricsprofile data from API
func (r *MetricsprofileResource) readMetricsprofileFromApi(ctx context.Context, data *MetricsprofileResourceModel, diags *diag.Diagnostics) {

	// Case 2: Find with single ID attribute - ID is the plain value
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Metricsprofile.Type(), name_Name)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read metricsprofile, got error: %s", err))
		return
	}

	metricsprofileSetAttrFromGet(ctx, data, getResponseData)

}
