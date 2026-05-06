package analyticsprofile

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
var _ resource.Resource = &AnalyticsprofileResource{}
var _ resource.ResourceWithConfigure = (*AnalyticsprofileResource)(nil)
var _ resource.ResourceWithImportState = (*AnalyticsprofileResource)(nil)

func NewAnalyticsprofileResource() resource.Resource {
	return &AnalyticsprofileResource{}
}

// AnalyticsprofileResource defines the resource implementation.
type AnalyticsprofileResource struct {
	client *service.NitroClient
}

func (r *AnalyticsprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AnalyticsprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_analyticsprofile"
}

func (r *AnalyticsprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AnalyticsprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config AnalyticsprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating analyticsprofile resource")
	// Get payload from plan (regular attributes)
	analyticsprofile := analyticsprofileGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	analyticsprofileGetThePayloadFromtheConfig(ctx, &config, &analyticsprofile)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Analyticsprofile.Type(), name_value, &analyticsprofile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create analyticsprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created analyticsprofile resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	r.readAnalyticsprofileFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AnalyticsprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AnalyticsprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading analyticsprofile resource")

	r.readAnalyticsprofileFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AnalyticsprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state AnalyticsprofileResourceModel

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

	tflog.Debug(ctx, "Updating analyticsprofile resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Allhttpheaders.Equal(state.Allhttpheaders) {
		tflog.Debug(ctx, fmt.Sprintf("allhttpheaders has changed for analyticsprofile"))
		hasChange = true
	}
	// Check secret attribute analyticsauthtoken or its version tracker
	if !data.Analyticsauthtoken.Equal(state.Analyticsauthtoken) {
		tflog.Debug(ctx, fmt.Sprintf("analyticsauthtoken has changed for analyticsprofile"))
		hasChange = true
	} else if !data.AnalyticsauthtokenWoVersion.Equal(state.AnalyticsauthtokenWoVersion) {
		tflog.Debug(ctx, fmt.Sprintf("analyticsauthtoken_wo_version has changed for analyticsprofile"))
		hasChange = true
	}
	if !data.Analyticsendpointcontenttype.Equal(state.Analyticsendpointcontenttype) {
		tflog.Debug(ctx, fmt.Sprintf("analyticsendpointcontenttype has changed for analyticsprofile"))
		hasChange = true
	}
	if !data.Analyticsendpointmetadata.Equal(state.Analyticsendpointmetadata) {
		tflog.Debug(ctx, fmt.Sprintf("analyticsendpointmetadata has changed for analyticsprofile"))
		hasChange = true
	}
	if !data.Analyticsendpointurl.Equal(state.Analyticsendpointurl) {
		tflog.Debug(ctx, fmt.Sprintf("analyticsendpointurl has changed for analyticsprofile"))
		hasChange = true
	}
	if !data.Auditlogs.Equal(state.Auditlogs) {
		tflog.Debug(ctx, fmt.Sprintf("auditlogs has changed for analyticsprofile"))
		hasChange = true
	}
	if !data.Collectors.Equal(state.Collectors) {
		tflog.Debug(ctx, fmt.Sprintf("collectors has changed for analyticsprofile"))
		hasChange = true
	}
	if !data.Cqareporting.Equal(state.Cqareporting) {
		tflog.Debug(ctx, fmt.Sprintf("cqareporting has changed for analyticsprofile"))
		hasChange = true
	}
	if !data.Dataformatfile.Equal(state.Dataformatfile) {
		tflog.Debug(ctx, fmt.Sprintf("dataformatfile has changed for analyticsprofile"))
		hasChange = true
	}
	if !data.Events.Equal(state.Events) {
		tflog.Debug(ctx, fmt.Sprintf("events has changed for analyticsprofile"))
		hasChange = true
	}
	if !data.Grpcstatus.Equal(state.Grpcstatus) {
		tflog.Debug(ctx, fmt.Sprintf("grpcstatus has changed for analyticsprofile"))
		hasChange = true
	}
	if !data.Httpauthentication.Equal(state.Httpauthentication) {
		tflog.Debug(ctx, fmt.Sprintf("httpauthentication has changed for analyticsprofile"))
		hasChange = true
	}
	if !data.Httpclientsidemeasurements.Equal(state.Httpclientsidemeasurements) {
		tflog.Debug(ctx, fmt.Sprintf("httpclientsidemeasurements has changed for analyticsprofile"))
		hasChange = true
	}
	if !data.Httpcontenttype.Equal(state.Httpcontenttype) {
		tflog.Debug(ctx, fmt.Sprintf("httpcontenttype has changed for analyticsprofile"))
		hasChange = true
	}
	if !data.Httpcookie.Equal(state.Httpcookie) {
		tflog.Debug(ctx, fmt.Sprintf("httpcookie has changed for analyticsprofile"))
		hasChange = true
	}
	if !data.Httpcustomheaders.Equal(state.Httpcustomheaders) {
		tflog.Debug(ctx, fmt.Sprintf("httpcustomheaders has changed for analyticsprofile"))
		hasChange = true
	}
	if !data.Httpdomainname.Equal(state.Httpdomainname) {
		tflog.Debug(ctx, fmt.Sprintf("httpdomainname has changed for analyticsprofile"))
		hasChange = true
	}
	if !data.Httphost.Equal(state.Httphost) {
		tflog.Debug(ctx, fmt.Sprintf("httphost has changed for analyticsprofile"))
		hasChange = true
	}
	if !data.Httplocation.Equal(state.Httplocation) {
		tflog.Debug(ctx, fmt.Sprintf("httplocation has changed for analyticsprofile"))
		hasChange = true
	}
	if !data.Httpmethod.Equal(state.Httpmethod) {
		tflog.Debug(ctx, fmt.Sprintf("httpmethod has changed for analyticsprofile"))
		hasChange = true
	}
	if !data.Httppagetracking.Equal(state.Httppagetracking) {
		tflog.Debug(ctx, fmt.Sprintf("httppagetracking has changed for analyticsprofile"))
		hasChange = true
	}
	if !data.Httpreferer.Equal(state.Httpreferer) {
		tflog.Debug(ctx, fmt.Sprintf("httpreferer has changed for analyticsprofile"))
		hasChange = true
	}
	if !data.Httpsetcookie.Equal(state.Httpsetcookie) {
		tflog.Debug(ctx, fmt.Sprintf("httpsetcookie has changed for analyticsprofile"))
		hasChange = true
	}
	if !data.Httpsetcookie2.Equal(state.Httpsetcookie2) {
		tflog.Debug(ctx, fmt.Sprintf("httpsetcookie2 has changed for analyticsprofile"))
		hasChange = true
	}
	if !data.Httpurl.Equal(state.Httpurl) {
		tflog.Debug(ctx, fmt.Sprintf("httpurl has changed for analyticsprofile"))
		hasChange = true
	}
	if !data.Httpurlquery.Equal(state.Httpurlquery) {
		tflog.Debug(ctx, fmt.Sprintf("httpurlquery has changed for analyticsprofile"))
		hasChange = true
	}
	if !data.Httpuseragent.Equal(state.Httpuseragent) {
		tflog.Debug(ctx, fmt.Sprintf("httpuseragent has changed for analyticsprofile"))
		hasChange = true
	}
	if !data.Httpvia.Equal(state.Httpvia) {
		tflog.Debug(ctx, fmt.Sprintf("httpvia has changed for analyticsprofile"))
		hasChange = true
	}
	if !data.Httpxforwardedforheader.Equal(state.Httpxforwardedforheader) {
		tflog.Debug(ctx, fmt.Sprintf("httpxforwardedforheader has changed for analyticsprofile"))
		hasChange = true
	}
	if !data.Integratedcache.Equal(state.Integratedcache) {
		tflog.Debug(ctx, fmt.Sprintf("integratedcache has changed for analyticsprofile"))
		hasChange = true
	}
	if !data.Managementlog.Equal(state.Managementlog) {
		tflog.Debug(ctx, fmt.Sprintf("managementlog has changed for analyticsprofile"))
		hasChange = true
	}
	if !data.Metrics.Equal(state.Metrics) {
		tflog.Debug(ctx, fmt.Sprintf("metrics has changed for analyticsprofile"))
		hasChange = true
	}
	if !data.Metricsexportfrequency.Equal(state.Metricsexportfrequency) {
		tflog.Debug(ctx, fmt.Sprintf("metricsexportfrequency has changed for analyticsprofile"))
		hasChange = true
	}
	if !data.Outputmode.Equal(state.Outputmode) {
		tflog.Debug(ctx, fmt.Sprintf("outputmode has changed for analyticsprofile"))
		hasChange = true
	}
	if !data.Schemafile.Equal(state.Schemafile) {
		tflog.Debug(ctx, fmt.Sprintf("schemafile has changed for analyticsprofile"))
		hasChange = true
	}
	if !data.Servemode.Equal(state.Servemode) {
		tflog.Debug(ctx, fmt.Sprintf("servemode has changed for analyticsprofile"))
		hasChange = true
	}
	if !data.Tcpburstreporting.Equal(state.Tcpburstreporting) {
		tflog.Debug(ctx, fmt.Sprintf("tcpburstreporting has changed for analyticsprofile"))
		hasChange = true
	}
	if !data.Topn.Equal(state.Topn) {
		tflog.Debug(ctx, fmt.Sprintf("topn has changed for analyticsprofile"))
		hasChange = true
	}
	if !data.Type.Equal(state.Type) {
		tflog.Debug(ctx, fmt.Sprintf("type has changed for analyticsprofile"))
		hasChange = true
	}
	if !data.Urlcategory.Equal(state.Urlcategory) {
		tflog.Debug(ctx, fmt.Sprintf("urlcategory has changed for analyticsprofile"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		// Get payload from plan (regular attributes)
		analyticsprofile := analyticsprofileGetThePayloadFromthePlan(ctx, &data)
		// Add write-only attributes from config to the payload
		analyticsprofileGetThePayloadFromtheConfig(ctx, &config, &analyticsprofile)
		// Make API call
		// Named resource - use UpdateResource
		name_value := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Analyticsprofile.Type(), name_value, &analyticsprofile)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update analyticsprofile, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated analyticsprofile resource")
	} else {
		tflog.Debug(ctx, "No changes detected for analyticsprofile resource, skipping update")
	}

	// Read the updated state back
	r.readAnalyticsprofileFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AnalyticsprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AnalyticsprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting analyticsprofile resource")
	// Named resource - delete using DeleteResource
	name_value := data.Name.ValueString()
	err := r.client.DeleteResource(service.Analyticsprofile.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete analyticsprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted analyticsprofile resource")
}

// Helper function to read analyticsprofile data from API
func (r *AnalyticsprofileResource) readAnalyticsprofileFromApi(ctx context.Context, data *AnalyticsprofileResourceModel, diags *diag.Diagnostics) {

	// Case 2: Find with single ID attribute - ID is the plain value
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Analyticsprofile.Type(), name_Name)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read analyticsprofile, got error: %s", err))
		return
	}

	analyticsprofileSetAttrFromGet(ctx, data, getResponseData)

}
