package lbmonitor

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &LbmonitorResource{}
var _ resource.ResourceWithConfigure = (*LbmonitorResource)(nil)
var _ resource.ResourceWithImportState = (*LbmonitorResource)(nil)

func NewLbmonitorResource() resource.Resource {
	return &LbmonitorResource{}
}

// LbmonitorResource defines the resource implementation.
type LbmonitorResource struct {
	client *service.NitroClient
}

func (r *LbmonitorResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LbmonitorResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lbmonitor"
}

func (r *LbmonitorResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LbmonitorResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config LbmonitorResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lbmonitor resource")
	// Get payload from plan (regular attributes)
	lbmonitor := lbmonitorGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	lbmonitorGetThePayloadFromtheConfig(ctx, &config, &lbmonitor)

	// Make API call
	_, err := r.client.AddResource(service.Lbmonitor.Type(), data.Monitorname.ValueString(), &lbmonitor)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lbmonitor, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created lbmonitor resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("monitorname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Monitorname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("type:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Type.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	if !r.readLbmonitorFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "lbmonitor not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbmonitorResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LbmonitorResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lbmonitor resource")

	found := r.readLbmonitorFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if !found {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbmonitorResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state LbmonitorResourceModel

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

	tflog.Debug(ctx, "Updating lbmonitor resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	attributesToUnset := []string{}
	if !data.Acctapplicationid.Equal(state.Acctapplicationid) {
		tflog.Debug(ctx, fmt.Sprintf("acctapplicationid has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Action.Equal(state.Action) {
		tflog.Debug(ctx, fmt.Sprintf("action has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Alertretries.Equal(state.Alertretries) {
		tflog.Debug(ctx, fmt.Sprintf("alertretries has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Application.Equal(state.Application) {
		tflog.Debug(ctx, fmt.Sprintf("application has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Attribute.Equal(state.Attribute) {
		tflog.Debug(ctx, fmt.Sprintf("attribute has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Authapplicationid.Equal(state.Authapplicationid) {
		tflog.Debug(ctx, fmt.Sprintf("authapplicationid has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Basedn.Equal(state.Basedn) {
		tflog.Debug(ctx, fmt.Sprintf("basedn has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Binddn.Equal(state.Binddn) {
		tflog.Debug(ctx, fmt.Sprintf("binddn has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Customheaders.Equal(state.Customheaders) {
		tflog.Debug(ctx, fmt.Sprintf("customheaders has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Database.Equal(state.Database) {
		tflog.Debug(ctx, fmt.Sprintf("database has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Destip.Equal(state.Destip) {
		tflog.Debug(ctx, fmt.Sprintf("destip has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Destport.Equal(state.Destport) {
		tflog.Debug(ctx, fmt.Sprintf("destport has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Deviation.Equal(state.Deviation) {
		tflog.Debug(ctx, fmt.Sprintf("deviation has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Dispatcherip.Equal(state.Dispatcherip) {
		tflog.Debug(ctx, fmt.Sprintf("dispatcherip has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Dispatcherport.Equal(state.Dispatcherport) {
		tflog.Debug(ctx, fmt.Sprintf("dispatcherport has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Domain.Equal(state.Domain) {
		tflog.Debug(ctx, fmt.Sprintf("domain has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Downtime.Equal(state.Downtime) {
		tflog.Debug(ctx, fmt.Sprintf("downtime has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Evalrule.Equal(state.Evalrule) {
		tflog.Debug(ctx, fmt.Sprintf("evalrule has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Failureretries.Equal(state.Failureretries) {
		tflog.Debug(ctx, fmt.Sprintf("failureretries has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Filename.Equal(state.Filename) {
		tflog.Debug(ctx, fmt.Sprintf("filename has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Filter.Equal(state.Filter) {
		tflog.Debug(ctx, fmt.Sprintf("filter has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Firmwarerevision.Equal(state.Firmwarerevision) {
		tflog.Debug(ctx, fmt.Sprintf("firmwarerevision has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Group.Equal(state.Group) {
		tflog.Debug(ctx, fmt.Sprintf("group has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Grpchealthcheck.Equal(state.Grpchealthcheck) {
		tflog.Debug(ctx, fmt.Sprintf("grpchealthcheck has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Grpcservicename.Equal(state.Grpcservicename) {
		tflog.Debug(ctx, fmt.Sprintf("grpcservicename has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Grpcstatuscode.Equal(state.Grpcstatuscode) {
		tflog.Debug(ctx, fmt.Sprintf("grpcstatuscode has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Hostipaddress.Equal(state.Hostipaddress) {
		tflog.Debug(ctx, fmt.Sprintf("hostipaddress has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Httprequest.Equal(state.Httprequest) {
		tflog.Debug(ctx, fmt.Sprintf("httprequest has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Inbandsecurityid.Equal(state.Inbandsecurityid) {
		tflog.Debug(ctx, fmt.Sprintf("inbandsecurityid has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Interval.Equal(state.Interval) {
		tflog.Debug(ctx, fmt.Sprintf("interval has changed for lbmonitor"))
		if config.Interval.IsNull() {
			attributesToUnset = append(attributesToUnset, "interval")
		} else {
			hasChange = true
		}
	}
	if !data.Ipaddress.Equal(state.Ipaddress) {
		tflog.Debug(ctx, fmt.Sprintf("ipaddress has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Iptunnel.Equal(state.Iptunnel) {
		tflog.Debug(ctx, fmt.Sprintf("iptunnel has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Kcdaccount.Equal(state.Kcdaccount) {
		tflog.Debug(ctx, fmt.Sprintf("kcdaccount has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Lasversion.Equal(state.Lasversion) {
		tflog.Debug(ctx, fmt.Sprintf("lasversion has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Logonpointname.Equal(state.Logonpointname) {
		tflog.Debug(ctx, fmt.Sprintf("logonpointname has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Lrtm.Equal(state.Lrtm) {
		tflog.Debug(ctx, fmt.Sprintf("lrtm has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Maxforwards.Equal(state.Maxforwards) {
		tflog.Debug(ctx, fmt.Sprintf("maxforwards has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Metric.Equal(state.Metric) {
		tflog.Debug(ctx, fmt.Sprintf("metric has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Metrictable.Equal(state.Metrictable) {
		tflog.Debug(ctx, fmt.Sprintf("metrictable has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Metricthreshold.Equal(state.Metricthreshold) {
		tflog.Debug(ctx, fmt.Sprintf("metricthreshold has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Metricweight.Equal(state.Metricweight) {
		tflog.Debug(ctx, fmt.Sprintf("metricweight has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Mqttclientidentifier.Equal(state.Mqttclientidentifier) {
		tflog.Debug(ctx, fmt.Sprintf("mqttclientidentifier has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Mqttversion.Equal(state.Mqttversion) {
		tflog.Debug(ctx, fmt.Sprintf("mqttversion has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Mssqlprotocolversion.Equal(state.Mssqlprotocolversion) {
		tflog.Debug(ctx, fmt.Sprintf("mssqlprotocolversion has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Netprofile.Equal(state.Netprofile) {
		tflog.Debug(ctx, fmt.Sprintf("netprofile has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Oraclesid.Equal(state.Oraclesid) {
		tflog.Debug(ctx, fmt.Sprintf("oraclesid has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Originhost.Equal(state.Originhost) {
		tflog.Debug(ctx, fmt.Sprintf("originhost has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Originrealm.Equal(state.Originrealm) {
		tflog.Debug(ctx, fmt.Sprintf("originrealm has changed for lbmonitor"))
		hasChange = true
	}
	// Check secret attribute password or its version tracker
	if !data.Password.Equal(state.Password) {
		tflog.Debug(ctx, fmt.Sprintf("password has changed for lbmonitor"))
		hasChange = true
	} else if !data.PasswordWoVersion.Equal(state.PasswordWoVersion) {
		tflog.Debug(ctx, fmt.Sprintf("password_wo_version has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Productname.Equal(state.Productname) {
		tflog.Debug(ctx, fmt.Sprintf("productname has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Query.Equal(state.Query) {
		tflog.Debug(ctx, fmt.Sprintf("query has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Querytype.Equal(state.Querytype) {
		tflog.Debug(ctx, fmt.Sprintf("querytype has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Radaccountsession.Equal(state.Radaccountsession) {
		tflog.Debug(ctx, fmt.Sprintf("radaccountsession has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Radaccounttype.Equal(state.Radaccounttype) {
		tflog.Debug(ctx, fmt.Sprintf("radaccounttype has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Radapn.Equal(state.Radapn) {
		tflog.Debug(ctx, fmt.Sprintf("radapn has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Radframedip.Equal(state.Radframedip) {
		tflog.Debug(ctx, fmt.Sprintf("radframedip has changed for lbmonitor"))
		hasChange = true
	}
	// Check secret attribute radkey or its version tracker
	if !data.Radkey.Equal(state.Radkey) {
		tflog.Debug(ctx, fmt.Sprintf("radkey has changed for lbmonitor"))
		hasChange = true
	} else if !data.RadkeyWoVersion.Equal(state.RadkeyWoVersion) {
		tflog.Debug(ctx, fmt.Sprintf("radkey_wo_version has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Radmsisdn.Equal(state.Radmsisdn) {
		tflog.Debug(ctx, fmt.Sprintf("radmsisdn has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Radnasid.Equal(state.Radnasid) {
		tflog.Debug(ctx, fmt.Sprintf("radnasid has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Radnasip.Equal(state.Radnasip) {
		tflog.Debug(ctx, fmt.Sprintf("radnasip has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Recv.Equal(state.Recv) {
		tflog.Debug(ctx, fmt.Sprintf("recv has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Respcode.Equal(state.Respcode) {
		tflog.Debug(ctx, fmt.Sprintf("respcode has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Resptimeout.Equal(state.Resptimeout) {
		tflog.Debug(ctx, fmt.Sprintf("resptimeout has changed for lbmonitor"))
		if config.Resptimeout.IsNull() {
			attributesToUnset = append(attributesToUnset, "resptimeout")
		} else {
			hasChange = true
		}
	}
	if !data.Resptimeoutthresh.Equal(state.Resptimeoutthresh) {
		tflog.Debug(ctx, fmt.Sprintf("resptimeoutthresh has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Retries.Equal(state.Retries) {
		tflog.Debug(ctx, fmt.Sprintf("retries has changed for lbmonitor"))
		if config.Retries.IsNull() {
			attributesToUnset = append(attributesToUnset, "retries")
		} else {
			hasChange = true
		}
	}
	if !data.Reverse.Equal(state.Reverse) {
		tflog.Debug(ctx, fmt.Sprintf("reverse has changed for lbmonitor"))
		if config.Reverse.IsNull() {
			attributesToUnset = append(attributesToUnset, "reverse")
		} else {
			hasChange = true
		}
	}
	if !data.Rtsprequest.Equal(state.Rtsprequest) {
		tflog.Debug(ctx, fmt.Sprintf("rtsprequest has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Scriptargs.Equal(state.Scriptargs) {
		tflog.Debug(ctx, fmt.Sprintf("scriptargs has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Scriptname.Equal(state.Scriptname) {
		tflog.Debug(ctx, fmt.Sprintf("scriptname has changed for lbmonitor"))
		hasChange = true
	}
	// Check secret attribute secondarypassword or its version tracker
	if !data.Secondarypassword.Equal(state.Secondarypassword) {
		tflog.Debug(ctx, fmt.Sprintf("secondarypassword has changed for lbmonitor"))
		hasChange = true
	} else if !data.SecondarypasswordWoVersion.Equal(state.SecondarypasswordWoVersion) {
		tflog.Debug(ctx, fmt.Sprintf("secondarypassword_wo_version has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Secure.Equal(state.Secure) {
		tflog.Debug(ctx, fmt.Sprintf("secure has changed for lbmonitor"))
		hasChange = true
	}
	// Check secret attribute secureargs or its version tracker
	if !data.Secureargs.Equal(state.Secureargs) {
		tflog.Debug(ctx, fmt.Sprintf("secureargs has changed for lbmonitor"))
		hasChange = true
	} else if !data.SecureargsWoVersion.Equal(state.SecureargsWoVersion) {
		tflog.Debug(ctx, fmt.Sprintf("secureargs_wo_version has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Send.Equal(state.Send) {
		tflog.Debug(ctx, fmt.Sprintf("send has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Sipmethod.Equal(state.Sipmethod) {
		tflog.Debug(ctx, fmt.Sprintf("sipmethod has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Sipreguri.Equal(state.Sipreguri) {
		tflog.Debug(ctx, fmt.Sprintf("sipreguri has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Sipuri.Equal(state.Sipuri) {
		tflog.Debug(ctx, fmt.Sprintf("sipuri has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Sitepath.Equal(state.Sitepath) {
		tflog.Debug(ctx, fmt.Sprintf("sitepath has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Snmpcommunity.Equal(state.Snmpcommunity) {
		tflog.Debug(ctx, fmt.Sprintf("snmpcommunity has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Snmpthreshold.Equal(state.Snmpthreshold) {
		tflog.Debug(ctx, fmt.Sprintf("snmpthreshold has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Snmpversion.Equal(state.Snmpversion) {
		tflog.Debug(ctx, fmt.Sprintf("snmpversion has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Sqlquery.Equal(state.Sqlquery) {
		tflog.Debug(ctx, fmt.Sprintf("sqlquery has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Sslprofile.Equal(state.Sslprofile) {
		tflog.Debug(ctx, fmt.Sprintf("sslprofile has changed for lbmonitor"))
		hasChange = true
	}
	if !data.State.Equal(state.State) {
		tflog.Debug(ctx, fmt.Sprintf("state has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Storedb.Equal(state.Storedb) {
		tflog.Debug(ctx, fmt.Sprintf("storedb has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Storefrontacctservice.Equal(state.Storefrontacctservice) {
		tflog.Debug(ctx, fmt.Sprintf("storefrontacctservice has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Storefrontcheckbackendservices.Equal(state.Storefrontcheckbackendservices) {
		tflog.Debug(ctx, fmt.Sprintf("storefrontcheckbackendservices has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Storename.Equal(state.Storename) {
		tflog.Debug(ctx, fmt.Sprintf("storename has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Successretries.Equal(state.Successretries) {
		tflog.Debug(ctx, fmt.Sprintf("successretries has changed for lbmonitor"))
		if config.Successretries.IsNull() {
			attributesToUnset = append(attributesToUnset, "successretries")
		} else {
			hasChange = true
		}
	}
	if !data.Supportedvendorids.Equal(state.Supportedvendorids) {
		tflog.Debug(ctx, fmt.Sprintf("supportedvendorids has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Tos.Equal(state.Tos) {
		tflog.Debug(ctx, fmt.Sprintf("tos has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Tosid.Equal(state.Tosid) {
		tflog.Debug(ctx, fmt.Sprintf("tosid has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Transparent.Equal(state.Transparent) {
		tflog.Debug(ctx, fmt.Sprintf("transparent has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Trofscode.Equal(state.Trofscode) {
		tflog.Debug(ctx, fmt.Sprintf("trofscode has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Trofsstring.Equal(state.Trofsstring) {
		tflog.Debug(ctx, fmt.Sprintf("trofsstring has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Type.Equal(state.Type) {
		tflog.Debug(ctx, fmt.Sprintf("type has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Units1.Equal(state.Units1) {
		tflog.Debug(ctx, fmt.Sprintf("units1 has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Units2.Equal(state.Units2) {
		tflog.Debug(ctx, fmt.Sprintf("units2 has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Units3.Equal(state.Units3) {
		tflog.Debug(ctx, fmt.Sprintf("units3 has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Units4.Equal(state.Units4) {
		tflog.Debug(ctx, fmt.Sprintf("units4 has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Username.Equal(state.Username) {
		tflog.Debug(ctx, fmt.Sprintf("username has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Validatecred.Equal(state.Validatecred) {
		tflog.Debug(ctx, fmt.Sprintf("validatecred has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Vendorid.Equal(state.Vendorid) {
		tflog.Debug(ctx, fmt.Sprintf("vendorid has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Vendorspecificacctapplicationids.Equal(state.Vendorspecificacctapplicationids) {
		tflog.Debug(ctx, fmt.Sprintf("vendorspecificacctapplicationids has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Vendorspecificauthapplicationids.Equal(state.Vendorspecificauthapplicationids) {
		tflog.Debug(ctx, fmt.Sprintf("vendorspecificauthapplicationids has changed for lbmonitor"))
		hasChange = true
	}
	if !data.Vendorspecificvendorid.Equal(state.Vendorspecificvendorid) {
		tflog.Debug(ctx, fmt.Sprintf("vendorspecificvendorid has changed for lbmonitor"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		// Get payload from plan (regular attributes)
		lbmonitor := lbmonitorGetThePayloadFromthePlan(ctx, &data)
		// Add write-only attributes from config to the payload
		lbmonitorGetThePayloadFromtheConfig(ctx, &config, &lbmonitor)
		// Make API call
		_, err := r.client.UpdateResource(service.Lbmonitor.Type(), data.Monitorname.ValueString(), &lbmonitor)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lbmonitor, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated lbmonitor resource")
	} else {
		tflog.Debug(ctx, "No changes detected for lbmonitor resource, skipping update")
	}

	// Clear attributes removed from configuration via NITRO unset
	unsetIdPayload := map[string]interface{}{
		"monitorname": data.Monitorname.ValueString(),
		"type":        data.Type.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Lbmonitor.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset lbmonitor attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readLbmonitorFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "lbmonitor not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbmonitorResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LbmonitorResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lbmonitor resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"monitorname"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	monitorname_value, ok := idMap["monitorname"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'monitorname' not found in ID")
		return
	}

	var argsMap map[string]string = make(map[string]string)
	if val, ok := idMap["type"]; ok && val != "" {
		argsMap["type"] = val
	}

	err = r.client.DeleteResourceWithArgsMap(service.Lbmonitor.Type(), monitorname_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete lbmonitor, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted lbmonitor binding")
}

// Helper function to read lbmonitor data from API
func (r *LbmonitorResource) readLbmonitorFromApi(ctx context.Context, data *LbmonitorResourceModel, diags *diag.Diagnostics) bool {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"monitorname"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return false
	}

	monitorname_Name, ok := idMap["monitorname"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'monitorname' not found in ID string")
		return false
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Lbmonitor.Type(),
		ResourceName:             monitorname_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lbmonitor, got error: %s", err))
		return false
	}

	// Resource is missing
	if len(dataArr) == 0 {
		return false
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true
		// Check type
		if val, ok := idMap["type"]; ok && val != "" {
			if v, ok := v["type"]; ok {
				if v.(string) != val {
					match = false
				}
			}
		}
		if match {
			foundIndex = i
			break
		}
	}

	//  Resource is missing
	if foundIndex == -1 {
		return false
	}

	lbmonitorSetAttrFromGet(ctx, data, dataArr[foundIndex])
	return true
}
