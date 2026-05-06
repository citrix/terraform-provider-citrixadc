package appflowparam

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
var _ resource.Resource = &AppflowparamResource{}
var _ resource.ResourceWithConfigure = (*AppflowparamResource)(nil)
var _ resource.ResourceWithImportState = (*AppflowparamResource)(nil)

func NewAppflowparamResource() resource.Resource {
	return &AppflowparamResource{}
}

// AppflowparamResource defines the resource implementation.
type AppflowparamResource struct {
	client *service.NitroClient
}

func (r *AppflowparamResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppflowparamResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appflowparam"
}

func (r *AppflowparamResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppflowparamResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config AppflowparamResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appflowparam resource")
	// Get payload from plan (regular attributes)
	appflowparam := appflowparamGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	appflowparamGetThePayloadFromtheConfig(ctx, &config, &appflowparam)

	// Make API call
	// Singleton resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Appflowparam.Type(), &appflowparam)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appflowparam, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created appflowparam resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue("appflowparam-config")

	// Read the updated state back
	r.readAppflowparamFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppflowparamResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppflowparamResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appflowparam resource")

	r.readAppflowparamFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppflowparamResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state AppflowparamResourceModel

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

	tflog.Debug(ctx, "Updating appflowparam resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Aaausername.Equal(state.Aaausername) {
		tflog.Debug(ctx, fmt.Sprintf("aaausername has changed for appflowparam"))
		hasChange = true
	}
	// Check secret attribute analyticsauthtoken or its version tracker
	if !data.Analyticsauthtoken.Equal(state.Analyticsauthtoken) {
		tflog.Debug(ctx, fmt.Sprintf("analyticsauthtoken has changed for appflowparam"))
		hasChange = true
	} else if !data.AnalyticsauthtokenWoVersion.Equal(state.AnalyticsauthtokenWoVersion) {
		tflog.Debug(ctx, fmt.Sprintf("analyticsauthtoken_wo_version has changed for appflowparam"))
		hasChange = true
	}
	if !data.Appnamerefresh.Equal(state.Appnamerefresh) {
		tflog.Debug(ctx, fmt.Sprintf("appnamerefresh has changed for appflowparam"))
		hasChange = true
	}
	if !data.Auditlogs.Equal(state.Auditlogs) {
		tflog.Debug(ctx, fmt.Sprintf("auditlogs has changed for appflowparam"))
		hasChange = true
	}
	if !data.Cacheinsight.Equal(state.Cacheinsight) {
		tflog.Debug(ctx, fmt.Sprintf("cacheinsight has changed for appflowparam"))
		hasChange = true
	}
	if !data.Clienttrafficonly.Equal(state.Clienttrafficonly) {
		tflog.Debug(ctx, fmt.Sprintf("clienttrafficonly has changed for appflowparam"))
		hasChange = true
	}
	if !data.Connectionchaining.Equal(state.Connectionchaining) {
		tflog.Debug(ctx, fmt.Sprintf("connectionchaining has changed for appflowparam"))
		hasChange = true
	}
	if !data.Cqareporting.Equal(state.Cqareporting) {
		tflog.Debug(ctx, fmt.Sprintf("cqareporting has changed for appflowparam"))
		hasChange = true
	}
	if !data.Distributedtracing.Equal(state.Distributedtracing) {
		tflog.Debug(ctx, fmt.Sprintf("distributedtracing has changed for appflowparam"))
		hasChange = true
	}
	if !data.Disttracingsamplingrate.Equal(state.Disttracingsamplingrate) {
		tflog.Debug(ctx, fmt.Sprintf("disttracingsamplingrate has changed for appflowparam"))
		hasChange = true
	}
	if !data.Emailaddress.Equal(state.Emailaddress) {
		tflog.Debug(ctx, fmt.Sprintf("emailaddress has changed for appflowparam"))
		hasChange = true
	}
	if !data.Events.Equal(state.Events) {
		tflog.Debug(ctx, fmt.Sprintf("events has changed for appflowparam"))
		hasChange = true
	}
	if !data.Flowrecordinterval.Equal(state.Flowrecordinterval) {
		tflog.Debug(ctx, fmt.Sprintf("flowrecordinterval has changed for appflowparam"))
		hasChange = true
	}
	if !data.Gxsessionreporting.Equal(state.Gxsessionreporting) {
		tflog.Debug(ctx, fmt.Sprintf("gxsessionreporting has changed for appflowparam"))
		hasChange = true
	}
	if !data.Httpauthorization.Equal(state.Httpauthorization) {
		tflog.Debug(ctx, fmt.Sprintf("httpauthorization has changed for appflowparam"))
		hasChange = true
	}
	if !data.Httpcontenttype.Equal(state.Httpcontenttype) {
		tflog.Debug(ctx, fmt.Sprintf("httpcontenttype has changed for appflowparam"))
		hasChange = true
	}
	if !data.Httpcookie.Equal(state.Httpcookie) {
		tflog.Debug(ctx, fmt.Sprintf("httpcookie has changed for appflowparam"))
		hasChange = true
	}
	if !data.Httpdomain.Equal(state.Httpdomain) {
		tflog.Debug(ctx, fmt.Sprintf("httpdomain has changed for appflowparam"))
		hasChange = true
	}
	if !data.Httphost.Equal(state.Httphost) {
		tflog.Debug(ctx, fmt.Sprintf("httphost has changed for appflowparam"))
		hasChange = true
	}
	if !data.Httplocation.Equal(state.Httplocation) {
		tflog.Debug(ctx, fmt.Sprintf("httplocation has changed for appflowparam"))
		hasChange = true
	}
	if !data.Httpmethod.Equal(state.Httpmethod) {
		tflog.Debug(ctx, fmt.Sprintf("httpmethod has changed for appflowparam"))
		hasChange = true
	}
	if !data.Httpquerywithurl.Equal(state.Httpquerywithurl) {
		tflog.Debug(ctx, fmt.Sprintf("httpquerywithurl has changed for appflowparam"))
		hasChange = true
	}
	if !data.Httpreferer.Equal(state.Httpreferer) {
		tflog.Debug(ctx, fmt.Sprintf("httpreferer has changed for appflowparam"))
		hasChange = true
	}
	if !data.Httpsetcookie.Equal(state.Httpsetcookie) {
		tflog.Debug(ctx, fmt.Sprintf("httpsetcookie has changed for appflowparam"))
		hasChange = true
	}
	if !data.Httpsetcookie2.Equal(state.Httpsetcookie2) {
		tflog.Debug(ctx, fmt.Sprintf("httpsetcookie2 has changed for appflowparam"))
		hasChange = true
	}
	if !data.Httpurl.Equal(state.Httpurl) {
		tflog.Debug(ctx, fmt.Sprintf("httpurl has changed for appflowparam"))
		hasChange = true
	}
	if !data.Httpuseragent.Equal(state.Httpuseragent) {
		tflog.Debug(ctx, fmt.Sprintf("httpuseragent has changed for appflowparam"))
		hasChange = true
	}
	if !data.Httpvia.Equal(state.Httpvia) {
		tflog.Debug(ctx, fmt.Sprintf("httpvia has changed for appflowparam"))
		hasChange = true
	}
	if !data.Httpxforwardedfor.Equal(state.Httpxforwardedfor) {
		tflog.Debug(ctx, fmt.Sprintf("httpxforwardedfor has changed for appflowparam"))
		hasChange = true
	}
	if !data.Identifiername.Equal(state.Identifiername) {
		tflog.Debug(ctx, fmt.Sprintf("identifiername has changed for appflowparam"))
		hasChange = true
	}
	if !data.Identifiersessionname.Equal(state.Identifiersessionname) {
		tflog.Debug(ctx, fmt.Sprintf("identifiersessionname has changed for appflowparam"))
		hasChange = true
	}
	if !data.Logstreamovernsip.Equal(state.Logstreamovernsip) {
		tflog.Debug(ctx, fmt.Sprintf("logstreamovernsip has changed for appflowparam"))
		hasChange = true
	}
	if !data.Lsnlogging.Equal(state.Lsnlogging) {
		tflog.Debug(ctx, fmt.Sprintf("lsnlogging has changed for appflowparam"))
		hasChange = true
	}
	if !data.Metrics.Equal(state.Metrics) {
		tflog.Debug(ctx, fmt.Sprintf("metrics has changed for appflowparam"))
		hasChange = true
	}
	if !data.Observationdomainid.Equal(state.Observationdomainid) {
		tflog.Debug(ctx, fmt.Sprintf("observationdomainid has changed for appflowparam"))
		hasChange = true
	}
	if !data.Observationdomainname.Equal(state.Observationdomainname) {
		tflog.Debug(ctx, fmt.Sprintf("observationdomainname has changed for appflowparam"))
		hasChange = true
	}
	if !data.Observationpointid.Equal(state.Observationpointid) {
		tflog.Debug(ctx, fmt.Sprintf("observationpointid has changed for appflowparam"))
		hasChange = true
	}
	if !data.Securityinsightrecordinterval.Equal(state.Securityinsightrecordinterval) {
		tflog.Debug(ctx, fmt.Sprintf("securityinsightrecordinterval has changed for appflowparam"))
		hasChange = true
	}
	if !data.Securityinsighttraffic.Equal(state.Securityinsighttraffic) {
		tflog.Debug(ctx, fmt.Sprintf("securityinsighttraffic has changed for appflowparam"))
		hasChange = true
	}
	if !data.Skipcacheredirectionhttptransaction.Equal(state.Skipcacheredirectionhttptransaction) {
		tflog.Debug(ctx, fmt.Sprintf("skipcacheredirectionhttptransaction has changed for appflowparam"))
		hasChange = true
	}
	if !data.Subscriberawareness.Equal(state.Subscriberawareness) {
		tflog.Debug(ctx, fmt.Sprintf("subscriberawareness has changed for appflowparam"))
		hasChange = true
	}
	if !data.Subscriberidobfuscation.Equal(state.Subscriberidobfuscation) {
		tflog.Debug(ctx, fmt.Sprintf("subscriberidobfuscation has changed for appflowparam"))
		hasChange = true
	}
	if !data.Subscriberidobfuscationalgo.Equal(state.Subscriberidobfuscationalgo) {
		tflog.Debug(ctx, fmt.Sprintf("subscriberidobfuscationalgo has changed for appflowparam"))
		hasChange = true
	}
	if !data.Tcpattackcounterinterval.Equal(state.Tcpattackcounterinterval) {
		tflog.Debug(ctx, fmt.Sprintf("tcpattackcounterinterval has changed for appflowparam"))
		hasChange = true
	}
	if !data.Templaterefresh.Equal(state.Templaterefresh) {
		tflog.Debug(ctx, fmt.Sprintf("templaterefresh has changed for appflowparam"))
		hasChange = true
	}
	if !data.Timeseriesovernsip.Equal(state.Timeseriesovernsip) {
		tflog.Debug(ctx, fmt.Sprintf("timeseriesovernsip has changed for appflowparam"))
		hasChange = true
	}
	if !data.Udppmtu.Equal(state.Udppmtu) {
		tflog.Debug(ctx, fmt.Sprintf("udppmtu has changed for appflowparam"))
		hasChange = true
	}
	if !data.Urlcategory.Equal(state.Urlcategory) {
		tflog.Debug(ctx, fmt.Sprintf("urlcategory has changed for appflowparam"))
		hasChange = true
	}
	if !data.Usagerecordinterval.Equal(state.Usagerecordinterval) {
		tflog.Debug(ctx, fmt.Sprintf("usagerecordinterval has changed for appflowparam"))
		hasChange = true
	}
	if !data.Videoinsight.Equal(state.Videoinsight) {
		tflog.Debug(ctx, fmt.Sprintf("videoinsight has changed for appflowparam"))
		hasChange = true
	}
	if !data.Websaasappusagereporting.Equal(state.Websaasappusagereporting) {
		tflog.Debug(ctx, fmt.Sprintf("websaasappusagereporting has changed for appflowparam"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		// Get payload from plan (regular attributes)
		appflowparam := appflowparamGetThePayloadFromthePlan(ctx, &data)
		// Add write-only attributes from config to the payload
		appflowparamGetThePayloadFromtheConfig(ctx, &config, &appflowparam)
		// Make API call
		// Singleton resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Appflowparam.Type(), &appflowparam)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appflowparam, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated appflowparam resource")
	} else {
		tflog.Debug(ctx, "No changes detected for appflowparam resource, skipping update")
	}

	// Read the updated state back
	r.readAppflowparamFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppflowparamResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppflowparamResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appflowparam resource")
	// Singleton resource - no delete operation on ADC, just remove from state
	tflog.Trace(ctx, "Removed appflowparam from Terraform state")
}

// Helper function to read appflowparam data from API
func (r *AppflowparamResource) readAppflowparamFromApi(ctx context.Context, data *AppflowparamResourceModel, diags *diag.Diagnostics) {

	// Case 1: Simple find without ID
	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Appflowparam.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appflowparam, got error: %s", err))
		return
	}

	appflowparamSetAttrFromGet(ctx, data, getResponseData)

}
