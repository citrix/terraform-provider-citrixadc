package auditsyslogaction

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
var _ resource.Resource = &AuditsyslogactionResource{}
var _ resource.ResourceWithConfigure = (*AuditsyslogactionResource)(nil)
var _ resource.ResourceWithImportState = (*AuditsyslogactionResource)(nil)

func NewAuditsyslogactionResource() resource.Resource {
	return &AuditsyslogactionResource{}
}

// AuditsyslogactionResource defines the resource implementation.
type AuditsyslogactionResource struct {
	client *service.NitroClient
}

func (r *AuditsyslogactionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuditsyslogactionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_auditsyslogaction"
}

func (r *AuditsyslogactionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuditsyslogactionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config AuditsyslogactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating auditsyslogaction resource")
	// Get payload from plan (regular attributes)
	auditsyslogaction := auditsyslogactionGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	auditsyslogactionGetThePayloadFromtheConfig(ctx, &config, &auditsyslogaction)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Auditsyslogaction.Type(), name_value, &auditsyslogaction)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create auditsyslogaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created auditsyslogaction resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	r.readAuditsyslogactionFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuditsyslogactionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuditsyslogactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading auditsyslogaction resource")

	r.readAuditsyslogactionFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuditsyslogactionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state AuditsyslogactionResourceModel

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

	tflog.Debug(ctx, "Updating auditsyslogaction resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Acl.Equal(state.Acl) {
		tflog.Debug(ctx, fmt.Sprintf("acl has changed for auditsyslogaction"))
		hasChange = true
	}
	if !data.Alg.Equal(state.Alg) {
		tflog.Debug(ctx, fmt.Sprintf("alg has changed for auditsyslogaction"))
		hasChange = true
	}
	if !data.Appflowexport.Equal(state.Appflowexport) {
		tflog.Debug(ctx, fmt.Sprintf("appflowexport has changed for auditsyslogaction"))
		hasChange = true
	}
	if !data.Contentinspectionlog.Equal(state.Contentinspectionlog) {
		tflog.Debug(ctx, fmt.Sprintf("contentinspectionlog has changed for auditsyslogaction"))
		hasChange = true
	}
	if !data.Dateformat.Equal(state.Dateformat) {
		tflog.Debug(ctx, fmt.Sprintf("dateformat has changed for auditsyslogaction"))
		hasChange = true
	}
	if !data.Dns.Equal(state.Dns) {
		tflog.Debug(ctx, fmt.Sprintf("dns has changed for auditsyslogaction"))
		hasChange = true
	}
	if !data.Domainresolvenow.Equal(state.Domainresolvenow) {
		tflog.Debug(ctx, fmt.Sprintf("domainresolvenow has changed for auditsyslogaction"))
		hasChange = true
	}
	if !data.Domainresolveretry.Equal(state.Domainresolveretry) {
		tflog.Debug(ctx, fmt.Sprintf("domainresolveretry has changed for auditsyslogaction"))
		hasChange = true
	}
	// Check secret attribute httpauthtoken or its version tracker
	if !data.Httpauthtoken.Equal(state.Httpauthtoken) {
		tflog.Debug(ctx, fmt.Sprintf("httpauthtoken has changed for auditsyslogaction"))
		hasChange = true
	} else if !data.HttpauthtokenWoVersion.Equal(state.HttpauthtokenWoVersion) {
		tflog.Debug(ctx, fmt.Sprintf("httpauthtoken_wo_version has changed for auditsyslogaction"))
		hasChange = true
	}
	if !data.Httpendpointurl.Equal(state.Httpendpointurl) {
		tflog.Debug(ctx, fmt.Sprintf("httpendpointurl has changed for auditsyslogaction"))
		hasChange = true
	}
	if !data.Lbvservername.Equal(state.Lbvservername) {
		tflog.Debug(ctx, fmt.Sprintf("lbvservername has changed for auditsyslogaction"))
		hasChange = true
	}
	if !data.Logfacility.Equal(state.Logfacility) {
		tflog.Debug(ctx, fmt.Sprintf("logfacility has changed for auditsyslogaction"))
		hasChange = true
	}
	if !data.Loglevel.Equal(state.Loglevel) {
		tflog.Debug(ctx, fmt.Sprintf("loglevel has changed for auditsyslogaction"))
		hasChange = true
	}
	if !data.Lsn.Equal(state.Lsn) {
		tflog.Debug(ctx, fmt.Sprintf("lsn has changed for auditsyslogaction"))
		hasChange = true
	}
	if !data.Managementlog.Equal(state.Managementlog) {
		tflog.Debug(ctx, fmt.Sprintf("managementlog has changed for auditsyslogaction"))
		hasChange = true
	}
	if !data.Maxlogdatasizetohold.Equal(state.Maxlogdatasizetohold) {
		tflog.Debug(ctx, fmt.Sprintf("maxlogdatasizetohold has changed for auditsyslogaction"))
		hasChange = true
	}
	if !data.Mgmtloglevel.Equal(state.Mgmtloglevel) {
		tflog.Debug(ctx, fmt.Sprintf("mgmtloglevel has changed for auditsyslogaction"))
		hasChange = true
	}
	if !data.Netprofile.Equal(state.Netprofile) {
		tflog.Debug(ctx, fmt.Sprintf("netprofile has changed for auditsyslogaction"))
		hasChange = true
	}
	if !data.Protocolviolations.Equal(state.Protocolviolations) {
		tflog.Debug(ctx, fmt.Sprintf("protocolviolations has changed for auditsyslogaction"))
		hasChange = true
	}
	if !data.Serverdomainname.Equal(state.Serverdomainname) {
		tflog.Debug(ctx, fmt.Sprintf("serverdomainname has changed for auditsyslogaction"))
		hasChange = true
	}
	if !data.Serverip.Equal(state.Serverip) {
		tflog.Debug(ctx, fmt.Sprintf("serverip has changed for auditsyslogaction"))
		hasChange = true
	}
	if !data.Serverport.Equal(state.Serverport) {
		tflog.Debug(ctx, fmt.Sprintf("serverport has changed for auditsyslogaction"))
		hasChange = true
	}
	if !data.Sslinterception.Equal(state.Sslinterception) {
		tflog.Debug(ctx, fmt.Sprintf("sslinterception has changed for auditsyslogaction"))
		hasChange = true
	}
	if !data.Streamanalytics.Equal(state.Streamanalytics) {
		tflog.Debug(ctx, fmt.Sprintf("streamanalytics has changed for auditsyslogaction"))
		hasChange = true
	}
	if !data.Subscriberlog.Equal(state.Subscriberlog) {
		tflog.Debug(ctx, fmt.Sprintf("subscriberlog has changed for auditsyslogaction"))
		hasChange = true
	}
	if !data.Syslogcompliance.Equal(state.Syslogcompliance) {
		tflog.Debug(ctx, fmt.Sprintf("syslogcompliance has changed for auditsyslogaction"))
		hasChange = true
	}
	if !data.Tcp.Equal(state.Tcp) {
		tflog.Debug(ctx, fmt.Sprintf("tcp has changed for auditsyslogaction"))
		hasChange = true
	}
	if !data.Tcpprofilename.Equal(state.Tcpprofilename) {
		tflog.Debug(ctx, fmt.Sprintf("tcpprofilename has changed for auditsyslogaction"))
		hasChange = true
	}
	if !data.Timezone.Equal(state.Timezone) {
		tflog.Debug(ctx, fmt.Sprintf("timezone has changed for auditsyslogaction"))
		hasChange = true
	}
	if !data.Urlfiltering.Equal(state.Urlfiltering) {
		tflog.Debug(ctx, fmt.Sprintf("urlfiltering has changed for auditsyslogaction"))
		hasChange = true
	}
	if !data.Userdefinedauditlog.Equal(state.Userdefinedauditlog) {
		tflog.Debug(ctx, fmt.Sprintf("userdefinedauditlog has changed for auditsyslogaction"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		// Get payload from plan (regular attributes)
		auditsyslogaction := auditsyslogactionGetThePayloadFromthePlan(ctx, &data)
		// Add write-only attributes from config to the payload
		auditsyslogactionGetThePayloadFromtheConfig(ctx, &config, &auditsyslogaction)
		// Clear immutable fields: transport cannot be sent in update requests (ADC error 278)
		auditsyslogaction.Transport = ""
		// Make API call
		// Named resource - use UpdateResource
		name_value := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Auditsyslogaction.Type(), name_value, &auditsyslogaction)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update auditsyslogaction, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated auditsyslogaction resource")
	} else {
		tflog.Debug(ctx, "No changes detected for auditsyslogaction resource, skipping update")
	}

	// Read the updated state back
	r.readAuditsyslogactionFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuditsyslogactionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuditsyslogactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting auditsyslogaction resource")
	// Named resource - delete using DeleteResource
	name_value := data.Name.ValueString()
	err := r.client.DeleteResource(service.Auditsyslogaction.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete auditsyslogaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted auditsyslogaction resource")
}

// Helper function to read auditsyslogaction data from API
func (r *AuditsyslogactionResource) readAuditsyslogactionFromApi(ctx context.Context, data *AuditsyslogactionResourceModel, diags *diag.Diagnostics) {

	// Case 2: Find with single ID attribute - ID is the plain value
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Auditsyslogaction.Type(), name_Name)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read auditsyslogaction, got error: %s", err))
		return
	}

	auditsyslogactionSetAttrFromGet(ctx, data, getResponseData)

}
