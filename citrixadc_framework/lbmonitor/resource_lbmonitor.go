package lbmonitor

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/lb"
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
	r.readLbmonitorFromApi(ctx, &data, &resp.Diagnostics)

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

	r.readLbmonitorFromApi(ctx, &data, &resp.Diagnostics)

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

	// Delta-payload update. `full` is the complete payload built from the plan (value
	// conversion + write-only secrets); `lbmonitor` is a fresh payload seeded with only the
	// identity (monitorname + the mandatory type). For each attribute that genuinely changed we
	// copy just that field from `full`, so the PUT carries monitorname + type + changed fields
	// only. Rebuilding the whole struct instead re-sends type-specific/interdependent args and
	// breaks both the v2 -> Framework upgrade (ec1094 "too few arguments" when computed attrs are
	// unknown) and steady-state updates. type is Required (the monitor's identity) so NITRO needs
	// it on every lbmonitor PUT.
	full := lbmonitorGetThePayloadFromthePlan(ctx, &data)
	lbmonitorGetThePayloadFromtheConfig(ctx, &config, &full)
	lbmonitor := lb.Lbmonitor{
		Monitorname: data.Monitorname.ValueString(),
		Type:        data.Type.ValueString(),
	}

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Acctapplicationid.IsUnknown() && !data.Acctapplicationid.IsNull() && !data.Acctapplicationid.Equal(state.Acctapplicationid) {
		tflog.Debug(ctx, fmt.Sprintf("acctapplicationid has changed for lbmonitor"))
		lbmonitor.Acctapplicationid = full.Acctapplicationid
		hasChange = true
	}
	if !data.Action.IsUnknown() && !data.Action.IsNull() && !data.Action.Equal(state.Action) {
		tflog.Debug(ctx, fmt.Sprintf("action has changed for lbmonitor"))
		lbmonitor.Action = full.Action
		hasChange = true
	}
	if !data.Alertretries.IsUnknown() && !data.Alertretries.IsNull() && !data.Alertretries.Equal(state.Alertretries) {
		tflog.Debug(ctx, fmt.Sprintf("alertretries has changed for lbmonitor"))
		lbmonitor.Alertretries = full.Alertretries
		hasChange = true
	}
	if !data.Application.IsUnknown() && !data.Application.IsNull() && !data.Application.Equal(state.Application) {
		tflog.Debug(ctx, fmt.Sprintf("application has changed for lbmonitor"))
		lbmonitor.Application = full.Application
		hasChange = true
	}
	if !data.Attribute.IsUnknown() && !data.Attribute.IsNull() && !data.Attribute.Equal(state.Attribute) {
		tflog.Debug(ctx, fmt.Sprintf("attribute has changed for lbmonitor"))
		lbmonitor.Attribute = full.Attribute
		hasChange = true
	}
	if !data.Authapplicationid.IsUnknown() && !data.Authapplicationid.IsNull() && !data.Authapplicationid.Equal(state.Authapplicationid) {
		tflog.Debug(ctx, fmt.Sprintf("authapplicationid has changed for lbmonitor"))
		lbmonitor.Authapplicationid = full.Authapplicationid
		hasChange = true
	}
	if !data.Basedn.IsUnknown() && !data.Basedn.IsNull() && !data.Basedn.Equal(state.Basedn) {
		tflog.Debug(ctx, fmt.Sprintf("basedn has changed for lbmonitor"))
		lbmonitor.Basedn = full.Basedn
		hasChange = true
	}
	if !data.Binddn.IsUnknown() && !data.Binddn.IsNull() && !data.Binddn.Equal(state.Binddn) {
		tflog.Debug(ctx, fmt.Sprintf("binddn has changed for lbmonitor"))
		lbmonitor.Binddn = full.Binddn
		hasChange = true
	}
	if !data.Customheaders.IsUnknown() && !data.Customheaders.IsNull() && !data.Customheaders.Equal(state.Customheaders) {
		tflog.Debug(ctx, fmt.Sprintf("customheaders has changed for lbmonitor"))
		lbmonitor.Customheaders = full.Customheaders
		hasChange = true
	}
	if !data.Database.IsUnknown() && !data.Database.IsNull() && !data.Database.Equal(state.Database) {
		tflog.Debug(ctx, fmt.Sprintf("database has changed for lbmonitor"))
		lbmonitor.Database = full.Database
		hasChange = true
	}
	if !data.Destip.IsUnknown() && !data.Destip.IsNull() && !data.Destip.Equal(state.Destip) {
		tflog.Debug(ctx, fmt.Sprintf("destip has changed for lbmonitor"))
		lbmonitor.Destip = full.Destip
		hasChange = true
	}
	if !data.Destport.IsUnknown() && !data.Destport.IsNull() && !data.Destport.Equal(state.Destport) {
		tflog.Debug(ctx, fmt.Sprintf("destport has changed for lbmonitor"))
		lbmonitor.Destport = full.Destport
		hasChange = true
	}
	if !data.Deviation.IsUnknown() && !data.Deviation.IsNull() && !data.Deviation.Equal(state.Deviation) {
		tflog.Debug(ctx, fmt.Sprintf("deviation has changed for lbmonitor"))
		lbmonitor.Deviation = full.Deviation
		hasChange = true
	}
	if !data.Dispatcherip.IsUnknown() && !data.Dispatcherip.IsNull() && !data.Dispatcherip.Equal(state.Dispatcherip) {
		tflog.Debug(ctx, fmt.Sprintf("dispatcherip has changed for lbmonitor"))
		lbmonitor.Dispatcherip = full.Dispatcherip
		hasChange = true
	}
	if !data.Dispatcherport.IsUnknown() && !data.Dispatcherport.IsNull() && !data.Dispatcherport.Equal(state.Dispatcherport) {
		tflog.Debug(ctx, fmt.Sprintf("dispatcherport has changed for lbmonitor"))
		lbmonitor.Dispatcherport = full.Dispatcherport
		hasChange = true
	}
	if !data.Domain.IsUnknown() && !data.Domain.IsNull() && !data.Domain.Equal(state.Domain) {
		tflog.Debug(ctx, fmt.Sprintf("domain has changed for lbmonitor"))
		lbmonitor.Domain = full.Domain
		hasChange = true
	}
	if !data.Downtime.IsUnknown() && !data.Downtime.IsNull() && !data.Downtime.Equal(state.Downtime) {
		tflog.Debug(ctx, fmt.Sprintf("downtime has changed for lbmonitor"))
		lbmonitor.Downtime = full.Downtime
		hasChange = true
	}
	if !data.Evalrule.IsUnknown() && !data.Evalrule.IsNull() && !data.Evalrule.Equal(state.Evalrule) {
		tflog.Debug(ctx, fmt.Sprintf("evalrule has changed for lbmonitor"))
		lbmonitor.Evalrule = full.Evalrule
		hasChange = true
	}
	if !data.Failureretries.IsUnknown() && !data.Failureretries.IsNull() && !data.Failureretries.Equal(state.Failureretries) {
		tflog.Debug(ctx, fmt.Sprintf("failureretries has changed for lbmonitor"))
		lbmonitor.Failureretries = full.Failureretries
		hasChange = true
	}
	if !data.Filename.IsUnknown() && !data.Filename.IsNull() && !data.Filename.Equal(state.Filename) {
		tflog.Debug(ctx, fmt.Sprintf("filename has changed for lbmonitor"))
		lbmonitor.Filename = full.Filename
		hasChange = true
	}
	if !data.Filter.IsUnknown() && !data.Filter.IsNull() && !data.Filter.Equal(state.Filter) {
		tflog.Debug(ctx, fmt.Sprintf("filter has changed for lbmonitor"))
		lbmonitor.Filter = full.Filter
		hasChange = true
	}
	if !data.Firmwarerevision.IsUnknown() && !data.Firmwarerevision.IsNull() && !data.Firmwarerevision.Equal(state.Firmwarerevision) {
		tflog.Debug(ctx, fmt.Sprintf("firmwarerevision has changed for lbmonitor"))
		lbmonitor.Firmwarerevision = full.Firmwarerevision
		hasChange = true
	}
	if !data.Group.IsUnknown() && !data.Group.IsNull() && !data.Group.Equal(state.Group) {
		tflog.Debug(ctx, fmt.Sprintf("group has changed for lbmonitor"))
		lbmonitor.Group = full.Group
		hasChange = true
	}
	if !data.Grpchealthcheck.IsUnknown() && !data.Grpchealthcheck.IsNull() && !data.Grpchealthcheck.Equal(state.Grpchealthcheck) {
		tflog.Debug(ctx, fmt.Sprintf("grpchealthcheck has changed for lbmonitor"))
		lbmonitor.Grpchealthcheck = full.Grpchealthcheck
		hasChange = true
	}
	if !data.Grpcservicename.IsUnknown() && !data.Grpcservicename.IsNull() && !data.Grpcservicename.Equal(state.Grpcservicename) {
		tflog.Debug(ctx, fmt.Sprintf("grpcservicename has changed for lbmonitor"))
		lbmonitor.Grpcservicename = full.Grpcservicename
		hasChange = true
	}
	if !data.Grpcstatuscode.IsUnknown() && !data.Grpcstatuscode.IsNull() && !data.Grpcstatuscode.Equal(state.Grpcstatuscode) {
		tflog.Debug(ctx, fmt.Sprintf("grpcstatuscode has changed for lbmonitor"))
		lbmonitor.Grpcstatuscode = full.Grpcstatuscode
		hasChange = true
	}
	if !data.Hostipaddress.IsUnknown() && !data.Hostipaddress.IsNull() && !data.Hostipaddress.Equal(state.Hostipaddress) {
		tflog.Debug(ctx, fmt.Sprintf("hostipaddress has changed for lbmonitor"))
		lbmonitor.Hostipaddress = full.Hostipaddress
		hasChange = true
	}
	if !data.Httprequest.IsUnknown() && !data.Httprequest.IsNull() && !data.Httprequest.Equal(state.Httprequest) {
		tflog.Debug(ctx, fmt.Sprintf("httprequest has changed for lbmonitor"))
		lbmonitor.Httprequest = full.Httprequest
		hasChange = true
	}
	if !data.Inbandsecurityid.IsUnknown() && !data.Inbandsecurityid.IsNull() && !data.Inbandsecurityid.Equal(state.Inbandsecurityid) {
		tflog.Debug(ctx, fmt.Sprintf("inbandsecurityid has changed for lbmonitor"))
		lbmonitor.Inbandsecurityid = full.Inbandsecurityid
		hasChange = true
	}
	if !data.Interval.IsUnknown() && !data.Interval.IsNull() && !data.Interval.Equal(state.Interval) {
		tflog.Debug(ctx, fmt.Sprintf("interval has changed for lbmonitor"))
		lbmonitor.Interval = full.Interval
		hasChange = true
	}
	if !data.Ipaddress.IsUnknown() && !data.Ipaddress.IsNull() && !data.Ipaddress.Equal(state.Ipaddress) {
		tflog.Debug(ctx, fmt.Sprintf("ipaddress has changed for lbmonitor"))
		lbmonitor.Ipaddress = full.Ipaddress
		hasChange = true
	}
	if !data.Iptunnel.IsUnknown() && !data.Iptunnel.IsNull() && !data.Iptunnel.Equal(state.Iptunnel) {
		tflog.Debug(ctx, fmt.Sprintf("iptunnel has changed for lbmonitor"))
		lbmonitor.Iptunnel = full.Iptunnel
		hasChange = true
	}
	if !data.Kcdaccount.IsUnknown() && !data.Kcdaccount.IsNull() && !data.Kcdaccount.Equal(state.Kcdaccount) {
		tflog.Debug(ctx, fmt.Sprintf("kcdaccount has changed for lbmonitor"))
		lbmonitor.Kcdaccount = full.Kcdaccount
		hasChange = true
	}
	if !data.Lasversion.IsUnknown() && !data.Lasversion.IsNull() && !data.Lasversion.Equal(state.Lasversion) {
		tflog.Debug(ctx, fmt.Sprintf("lasversion has changed for lbmonitor"))
		lbmonitor.Lasversion = full.Lasversion
		hasChange = true
	}
	if !data.Logonpointname.IsUnknown() && !data.Logonpointname.IsNull() && !data.Logonpointname.Equal(state.Logonpointname) {
		tflog.Debug(ctx, fmt.Sprintf("logonpointname has changed for lbmonitor"))
		lbmonitor.Logonpointname = full.Logonpointname
		hasChange = true
	}
	if !data.Lrtm.IsUnknown() && !data.Lrtm.IsNull() && !data.Lrtm.Equal(state.Lrtm) {
		tflog.Debug(ctx, fmt.Sprintf("lrtm has changed for lbmonitor"))
		lbmonitor.Lrtm = full.Lrtm
		hasChange = true
	}
	if !data.Maxforwards.IsUnknown() && !data.Maxforwards.IsNull() && !data.Maxforwards.Equal(state.Maxforwards) {
		tflog.Debug(ctx, fmt.Sprintf("maxforwards has changed for lbmonitor"))
		lbmonitor.Maxforwards = full.Maxforwards
		hasChange = true
	}
	if !data.Metric.IsUnknown() && !data.Metric.IsNull() && !data.Metric.Equal(state.Metric) {
		tflog.Debug(ctx, fmt.Sprintf("metric has changed for lbmonitor"))
		lbmonitor.Metric = full.Metric
		hasChange = true
	}
	if !data.Metrictable.IsUnknown() && !data.Metrictable.IsNull() && !data.Metrictable.Equal(state.Metrictable) {
		tflog.Debug(ctx, fmt.Sprintf("metrictable has changed for lbmonitor"))
		lbmonitor.Metrictable = full.Metrictable
		hasChange = true
	}
	if !data.Metricthreshold.IsUnknown() && !data.Metricthreshold.IsNull() && !data.Metricthreshold.Equal(state.Metricthreshold) {
		tflog.Debug(ctx, fmt.Sprintf("metricthreshold has changed for lbmonitor"))
		lbmonitor.Metricthreshold = full.Metricthreshold
		hasChange = true
	}
	if !data.Metricweight.IsUnknown() && !data.Metricweight.IsNull() && !data.Metricweight.Equal(state.Metricweight) {
		tflog.Debug(ctx, fmt.Sprintf("metricweight has changed for lbmonitor"))
		lbmonitor.Metricweight = full.Metricweight
		hasChange = true
	}
	if !data.Mqttclientidentifier.IsUnknown() && !data.Mqttclientidentifier.IsNull() && !data.Mqttclientidentifier.Equal(state.Mqttclientidentifier) {
		tflog.Debug(ctx, fmt.Sprintf("mqttclientidentifier has changed for lbmonitor"))
		lbmonitor.Mqttclientidentifier = full.Mqttclientidentifier
		hasChange = true
	}
	if !data.Mqttversion.IsUnknown() && !data.Mqttversion.IsNull() && !data.Mqttversion.Equal(state.Mqttversion) {
		tflog.Debug(ctx, fmt.Sprintf("mqttversion has changed for lbmonitor"))
		lbmonitor.Mqttversion = full.Mqttversion
		hasChange = true
	}
	if !data.Mssqlprotocolversion.IsUnknown() && !data.Mssqlprotocolversion.IsNull() && !data.Mssqlprotocolversion.Equal(state.Mssqlprotocolversion) {
		tflog.Debug(ctx, fmt.Sprintf("mssqlprotocolversion has changed for lbmonitor"))
		lbmonitor.Mssqlprotocolversion = full.Mssqlprotocolversion
		hasChange = true
	}
	if !data.Netprofile.IsUnknown() && !data.Netprofile.IsNull() && !data.Netprofile.Equal(state.Netprofile) {
		tflog.Debug(ctx, fmt.Sprintf("netprofile has changed for lbmonitor"))
		lbmonitor.Netprofile = full.Netprofile
		hasChange = true
	}
	if !data.Oraclesid.IsUnknown() && !data.Oraclesid.IsNull() && !data.Oraclesid.Equal(state.Oraclesid) {
		tflog.Debug(ctx, fmt.Sprintf("oraclesid has changed for lbmonitor"))
		lbmonitor.Oraclesid = full.Oraclesid
		hasChange = true
	}
	if !data.Originhost.IsUnknown() && !data.Originhost.IsNull() && !data.Originhost.Equal(state.Originhost) {
		tflog.Debug(ctx, fmt.Sprintf("originhost has changed for lbmonitor"))
		lbmonitor.Originhost = full.Originhost
		hasChange = true
	}
	if !data.Originrealm.IsUnknown() && !data.Originrealm.IsNull() && !data.Originrealm.Equal(state.Originrealm) {
		tflog.Debug(ctx, fmt.Sprintf("originrealm has changed for lbmonitor"))
		lbmonitor.Originrealm = full.Originrealm
		hasChange = true
	}
	// Check secret attribute password or its version tracker
	// Only send the secret when config actually supplies it (full.Password != ""): guards the
	// legacy null/"" mismatch and the _wo_version Default drift from firing a spurious
	// update on the v2 -> Framework upgrade.
	if !data.Password.IsUnknown() && !data.Password.IsNull() && !data.Password.Equal(state.Password) && full.Password != "" {
		tflog.Debug(ctx, fmt.Sprintf("password has changed for lbmonitor"))
		lbmonitor.Password = full.Password
		hasChange = true
	} else if !data.PasswordWoVersion.Equal(state.PasswordWoVersion) && full.Password != "" {
		tflog.Debug(ctx, fmt.Sprintf("password_wo_version has changed for lbmonitor"))
		lbmonitor.Password = full.Password
		hasChange = true
	}
	if !data.Productname.IsUnknown() && !data.Productname.IsNull() && !data.Productname.Equal(state.Productname) {
		tflog.Debug(ctx, fmt.Sprintf("productname has changed for lbmonitor"))
		lbmonitor.Productname = full.Productname
		hasChange = true
	}
	if !data.Query.IsUnknown() && !data.Query.IsNull() && !data.Query.Equal(state.Query) {
		tflog.Debug(ctx, fmt.Sprintf("query has changed for lbmonitor"))
		lbmonitor.Query = full.Query
		hasChange = true
	}
	if !data.Querytype.IsUnknown() && !data.Querytype.IsNull() && !data.Querytype.Equal(state.Querytype) {
		tflog.Debug(ctx, fmt.Sprintf("querytype has changed for lbmonitor"))
		lbmonitor.Querytype = full.Querytype
		hasChange = true
	}
	if !data.Radaccountsession.IsUnknown() && !data.Radaccountsession.IsNull() && !data.Radaccountsession.Equal(state.Radaccountsession) {
		tflog.Debug(ctx, fmt.Sprintf("radaccountsession has changed for lbmonitor"))
		lbmonitor.Radaccountsession = full.Radaccountsession
		hasChange = true
	}
	if !data.Radaccounttype.IsUnknown() && !data.Radaccounttype.IsNull() && !data.Radaccounttype.Equal(state.Radaccounttype) {
		tflog.Debug(ctx, fmt.Sprintf("radaccounttype has changed for lbmonitor"))
		lbmonitor.Radaccounttype = full.Radaccounttype
		hasChange = true
	}
	if !data.Radapn.IsUnknown() && !data.Radapn.IsNull() && !data.Radapn.Equal(state.Radapn) {
		tflog.Debug(ctx, fmt.Sprintf("radapn has changed for lbmonitor"))
		lbmonitor.Radapn = full.Radapn
		hasChange = true
	}
	if !data.Radframedip.IsUnknown() && !data.Radframedip.IsNull() && !data.Radframedip.Equal(state.Radframedip) {
		tflog.Debug(ctx, fmt.Sprintf("radframedip has changed for lbmonitor"))
		lbmonitor.Radframedip = full.Radframedip
		hasChange = true
	}
	// Check secret attribute radkey or its version tracker
	// Only send the secret when config actually supplies it (full.Radkey != ""): guards the
	// legacy null/"" mismatch and the _wo_version Default drift from firing a spurious
	// update on the v2 -> Framework upgrade.
	if !data.Radkey.IsUnknown() && !data.Radkey.IsNull() && !data.Radkey.Equal(state.Radkey) && full.Radkey != "" {
		tflog.Debug(ctx, fmt.Sprintf("radkey has changed for lbmonitor"))
		lbmonitor.Radkey = full.Radkey
		hasChange = true
	} else if !data.RadkeyWoVersion.Equal(state.RadkeyWoVersion) && full.Radkey != "" {
		tflog.Debug(ctx, fmt.Sprintf("radkey_wo_version has changed for lbmonitor"))
		lbmonitor.Radkey = full.Radkey
		hasChange = true
	}
	if !data.Radmsisdn.IsUnknown() && !data.Radmsisdn.IsNull() && !data.Radmsisdn.Equal(state.Radmsisdn) {
		tflog.Debug(ctx, fmt.Sprintf("radmsisdn has changed for lbmonitor"))
		lbmonitor.Radmsisdn = full.Radmsisdn
		hasChange = true
	}
	if !data.Radnasid.IsUnknown() && !data.Radnasid.IsNull() && !data.Radnasid.Equal(state.Radnasid) {
		tflog.Debug(ctx, fmt.Sprintf("radnasid has changed for lbmonitor"))
		lbmonitor.Radnasid = full.Radnasid
		hasChange = true
	}
	if !data.Radnasip.IsUnknown() && !data.Radnasip.IsNull() && !data.Radnasip.Equal(state.Radnasip) {
		tflog.Debug(ctx, fmt.Sprintf("radnasip has changed for lbmonitor"))
		lbmonitor.Radnasip = full.Radnasip
		hasChange = true
	}
	if !data.Recv.IsUnknown() && !data.Recv.IsNull() && !data.Recv.Equal(state.Recv) {
		tflog.Debug(ctx, fmt.Sprintf("recv has changed for lbmonitor"))
		lbmonitor.Recv = full.Recv
		hasChange = true
	}
	if !data.Respcode.IsUnknown() && !data.Respcode.IsNull() && !data.Respcode.Equal(state.Respcode) {
		tflog.Debug(ctx, fmt.Sprintf("respcode has changed for lbmonitor"))
		lbmonitor.Respcode = full.Respcode
		hasChange = true
	}
	if !data.Resptimeout.IsUnknown() && !data.Resptimeout.IsNull() && !data.Resptimeout.Equal(state.Resptimeout) {
		tflog.Debug(ctx, fmt.Sprintf("resptimeout has changed for lbmonitor"))
		lbmonitor.Resptimeout = full.Resptimeout
		hasChange = true
	}
	if !data.Resptimeoutthresh.IsUnknown() && !data.Resptimeoutthresh.IsNull() && !data.Resptimeoutthresh.Equal(state.Resptimeoutthresh) {
		tflog.Debug(ctx, fmt.Sprintf("resptimeoutthresh has changed for lbmonitor"))
		lbmonitor.Resptimeoutthresh = full.Resptimeoutthresh
		hasChange = true
	}
	if !data.Retries.IsUnknown() && !data.Retries.IsNull() && !data.Retries.Equal(state.Retries) {
		tflog.Debug(ctx, fmt.Sprintf("retries has changed for lbmonitor"))
		lbmonitor.Retries = full.Retries
		hasChange = true
	}
	if !data.Reverse.IsUnknown() && !data.Reverse.IsNull() && !data.Reverse.Equal(state.Reverse) {
		tflog.Debug(ctx, fmt.Sprintf("reverse has changed for lbmonitor"))
		lbmonitor.Reverse = full.Reverse
		hasChange = true
	}
	if !data.Rtsprequest.IsUnknown() && !data.Rtsprequest.IsNull() && !data.Rtsprequest.Equal(state.Rtsprequest) {
		tflog.Debug(ctx, fmt.Sprintf("rtsprequest has changed for lbmonitor"))
		lbmonitor.Rtsprequest = full.Rtsprequest
		hasChange = true
	}
	if !data.Scriptargs.IsUnknown() && !data.Scriptargs.IsNull() && !data.Scriptargs.Equal(state.Scriptargs) {
		tflog.Debug(ctx, fmt.Sprintf("scriptargs has changed for lbmonitor"))
		lbmonitor.Scriptargs = full.Scriptargs
		hasChange = true
	}
	if !data.Scriptname.IsUnknown() && !data.Scriptname.IsNull() && !data.Scriptname.Equal(state.Scriptname) {
		tflog.Debug(ctx, fmt.Sprintf("scriptname has changed for lbmonitor"))
		lbmonitor.Scriptname = full.Scriptname
		hasChange = true
	}
	// Check secret attribute secondarypassword or its version tracker
	// Only send the secret when config actually supplies it (full.Secondarypassword != ""): guards the
	// legacy null/"" mismatch and the _wo_version Default drift from firing a spurious
	// update on the v2 -> Framework upgrade.
	if !data.Secondarypassword.IsUnknown() && !data.Secondarypassword.IsNull() && !data.Secondarypassword.Equal(state.Secondarypassword) && full.Secondarypassword != "" {
		tflog.Debug(ctx, fmt.Sprintf("secondarypassword has changed for lbmonitor"))
		lbmonitor.Secondarypassword = full.Secondarypassword
		hasChange = true
	} else if !data.SecondarypasswordWoVersion.Equal(state.SecondarypasswordWoVersion) && full.Secondarypassword != "" {
		tflog.Debug(ctx, fmt.Sprintf("secondarypassword_wo_version has changed for lbmonitor"))
		lbmonitor.Secondarypassword = full.Secondarypassword
		hasChange = true
	}
	if !data.Secure.IsUnknown() && !data.Secure.IsNull() && !data.Secure.Equal(state.Secure) {
		tflog.Debug(ctx, fmt.Sprintf("secure has changed for lbmonitor"))
		lbmonitor.Secure = full.Secure
		hasChange = true
	}
	// Check secret attribute secureargs or its version tracker
	// Only send the secret when config actually supplies it (full.Secureargs != ""): guards the
	// legacy null/"" mismatch and the _wo_version Default drift from firing a spurious
	// update on the v2 -> Framework upgrade.
	if !data.Secureargs.IsUnknown() && !data.Secureargs.IsNull() && !data.Secureargs.Equal(state.Secureargs) && full.Secureargs != "" {
		tflog.Debug(ctx, fmt.Sprintf("secureargs has changed for lbmonitor"))
		lbmonitor.Secureargs = full.Secureargs
		hasChange = true
	} else if !data.SecureargsWoVersion.Equal(state.SecureargsWoVersion) && full.Secureargs != "" {
		tflog.Debug(ctx, fmt.Sprintf("secureargs_wo_version has changed for lbmonitor"))
		lbmonitor.Secureargs = full.Secureargs
		hasChange = true
	}
	if !data.Send.IsUnknown() && !data.Send.IsNull() && !data.Send.Equal(state.Send) {
		tflog.Debug(ctx, fmt.Sprintf("send has changed for lbmonitor"))
		lbmonitor.Send = full.Send
		hasChange = true
	}
	if !data.Sipmethod.IsUnknown() && !data.Sipmethod.IsNull() && !data.Sipmethod.Equal(state.Sipmethod) {
		tflog.Debug(ctx, fmt.Sprintf("sipmethod has changed for lbmonitor"))
		lbmonitor.Sipmethod = full.Sipmethod
		hasChange = true
	}
	if !data.Sipreguri.IsUnknown() && !data.Sipreguri.IsNull() && !data.Sipreguri.Equal(state.Sipreguri) {
		tflog.Debug(ctx, fmt.Sprintf("sipreguri has changed for lbmonitor"))
		lbmonitor.Sipreguri = full.Sipreguri
		hasChange = true
	}
	if !data.Sipuri.IsUnknown() && !data.Sipuri.IsNull() && !data.Sipuri.Equal(state.Sipuri) {
		tflog.Debug(ctx, fmt.Sprintf("sipuri has changed for lbmonitor"))
		lbmonitor.Sipuri = full.Sipuri
		hasChange = true
	}
	if !data.Sitepath.IsUnknown() && !data.Sitepath.IsNull() && !data.Sitepath.Equal(state.Sitepath) {
		tflog.Debug(ctx, fmt.Sprintf("sitepath has changed for lbmonitor"))
		lbmonitor.Sitepath = full.Sitepath
		hasChange = true
	}
	if !data.Snmpcommunity.IsUnknown() && !data.Snmpcommunity.IsNull() && !data.Snmpcommunity.Equal(state.Snmpcommunity) {
		tflog.Debug(ctx, fmt.Sprintf("snmpcommunity has changed for lbmonitor"))
		lbmonitor.Snmpcommunity = full.Snmpcommunity
		hasChange = true
	}
	if !data.Snmpthreshold.IsUnknown() && !data.Snmpthreshold.IsNull() && !data.Snmpthreshold.Equal(state.Snmpthreshold) {
		tflog.Debug(ctx, fmt.Sprintf("snmpthreshold has changed for lbmonitor"))
		lbmonitor.Snmpthreshold = full.Snmpthreshold
		hasChange = true
	}
	if !data.Snmpversion.IsUnknown() && !data.Snmpversion.IsNull() && !data.Snmpversion.Equal(state.Snmpversion) {
		tflog.Debug(ctx, fmt.Sprintf("snmpversion has changed for lbmonitor"))
		lbmonitor.Snmpversion = full.Snmpversion
		hasChange = true
	}
	if !data.Sqlquery.IsUnknown() && !data.Sqlquery.IsNull() && !data.Sqlquery.Equal(state.Sqlquery) {
		tflog.Debug(ctx, fmt.Sprintf("sqlquery has changed for lbmonitor"))
		lbmonitor.Sqlquery = full.Sqlquery
		hasChange = true
	}
	if !data.Sslprofile.IsUnknown() && !data.Sslprofile.IsNull() && !data.Sslprofile.Equal(state.Sslprofile) {
		tflog.Debug(ctx, fmt.Sprintf("sslprofile has changed for lbmonitor"))
		lbmonitor.Sslprofile = full.Sslprofile
		hasChange = true
	}
	if !data.State.IsUnknown() && !data.State.IsNull() && !data.State.Equal(state.State) {
		tflog.Debug(ctx, fmt.Sprintf("state has changed for lbmonitor"))
		lbmonitor.State = full.State
		hasChange = true
	}
	if !data.Storedb.IsUnknown() && !data.Storedb.IsNull() && !data.Storedb.Equal(state.Storedb) {
		tflog.Debug(ctx, fmt.Sprintf("storedb has changed for lbmonitor"))
		lbmonitor.Storedb = full.Storedb
		hasChange = true
	}
	if !data.Storefrontacctservice.IsUnknown() && !data.Storefrontacctservice.IsNull() && !data.Storefrontacctservice.Equal(state.Storefrontacctservice) {
		tflog.Debug(ctx, fmt.Sprintf("storefrontacctservice has changed for lbmonitor"))
		lbmonitor.Storefrontacctservice = full.Storefrontacctservice
		hasChange = true
	}
	if !data.Storefrontcheckbackendservices.IsUnknown() && !data.Storefrontcheckbackendservices.IsNull() && !data.Storefrontcheckbackendservices.Equal(state.Storefrontcheckbackendservices) {
		tflog.Debug(ctx, fmt.Sprintf("storefrontcheckbackendservices has changed for lbmonitor"))
		lbmonitor.Storefrontcheckbackendservices = full.Storefrontcheckbackendservices
		hasChange = true
	}
	if !data.Storename.IsUnknown() && !data.Storename.IsNull() && !data.Storename.Equal(state.Storename) {
		tflog.Debug(ctx, fmt.Sprintf("storename has changed for lbmonitor"))
		lbmonitor.Storename = full.Storename
		hasChange = true
	}
	if !data.Successretries.IsUnknown() && !data.Successretries.IsNull() && !data.Successretries.Equal(state.Successretries) {
		tflog.Debug(ctx, fmt.Sprintf("successretries has changed for lbmonitor"))
		lbmonitor.Successretries = full.Successretries
		hasChange = true
	}
	if !data.Supportedvendorids.IsUnknown() && !data.Supportedvendorids.IsNull() && !data.Supportedvendorids.Equal(state.Supportedvendorids) {
		tflog.Debug(ctx, fmt.Sprintf("supportedvendorids has changed for lbmonitor"))
		lbmonitor.Supportedvendorids = full.Supportedvendorids
		hasChange = true
	}
	if !data.Tos.IsUnknown() && !data.Tos.IsNull() && !data.Tos.Equal(state.Tos) {
		tflog.Debug(ctx, fmt.Sprintf("tos has changed for lbmonitor"))
		lbmonitor.Tos = full.Tos
		hasChange = true
	}
	if !data.Tosid.IsUnknown() && !data.Tosid.IsNull() && !data.Tosid.Equal(state.Tosid) {
		tflog.Debug(ctx, fmt.Sprintf("tosid has changed for lbmonitor"))
		lbmonitor.Tosid = full.Tosid
		hasChange = true
	}
	if !data.Transparent.IsUnknown() && !data.Transparent.IsNull() && !data.Transparent.Equal(state.Transparent) {
		tflog.Debug(ctx, fmt.Sprintf("transparent has changed for lbmonitor"))
		lbmonitor.Transparent = full.Transparent
		hasChange = true
	}
	if !data.Trofscode.IsUnknown() && !data.Trofscode.IsNull() && !data.Trofscode.Equal(state.Trofscode) {
		tflog.Debug(ctx, fmt.Sprintf("trofscode has changed for lbmonitor"))
		lbmonitor.Trofscode = full.Trofscode
		hasChange = true
	}
	if !data.Trofsstring.IsUnknown() && !data.Trofsstring.IsNull() && !data.Trofsstring.Equal(state.Trofsstring) {
		tflog.Debug(ctx, fmt.Sprintf("trofsstring has changed for lbmonitor"))
		lbmonitor.Trofsstring = full.Trofsstring
		hasChange = true
	}
	if !data.Type.IsUnknown() && !data.Type.IsNull() && !data.Type.Equal(state.Type) {
		tflog.Debug(ctx, fmt.Sprintf("type has changed for lbmonitor"))
		lbmonitor.Type = full.Type
		hasChange = true
	}
	if !data.Units1.IsUnknown() && !data.Units1.IsNull() && !data.Units1.Equal(state.Units1) {
		tflog.Debug(ctx, fmt.Sprintf("units1 has changed for lbmonitor"))
		lbmonitor.Units1 = full.Units1
		hasChange = true
	}
	if !data.Units2.IsUnknown() && !data.Units2.IsNull() && !data.Units2.Equal(state.Units2) {
		tflog.Debug(ctx, fmt.Sprintf("units2 has changed for lbmonitor"))
		lbmonitor.Units2 = full.Units2
		hasChange = true
	}
	if !data.Units3.IsUnknown() && !data.Units3.IsNull() && !data.Units3.Equal(state.Units3) {
		tflog.Debug(ctx, fmt.Sprintf("units3 has changed for lbmonitor"))
		lbmonitor.Units3 = full.Units3
		hasChange = true
	}
	if !data.Units4.IsUnknown() && !data.Units4.IsNull() && !data.Units4.Equal(state.Units4) {
		tflog.Debug(ctx, fmt.Sprintf("units4 has changed for lbmonitor"))
		lbmonitor.Units4 = full.Units4
		hasChange = true
	}
	if !data.Username.IsUnknown() && !data.Username.IsNull() && !data.Username.Equal(state.Username) {
		tflog.Debug(ctx, fmt.Sprintf("username has changed for lbmonitor"))
		lbmonitor.Username = full.Username
		hasChange = true
	}
	if !data.Validatecred.IsUnknown() && !data.Validatecred.IsNull() && !data.Validatecred.Equal(state.Validatecred) {
		tflog.Debug(ctx, fmt.Sprintf("validatecred has changed for lbmonitor"))
		lbmonitor.Validatecred = full.Validatecred
		hasChange = true
	}
	if !data.Vendorid.IsUnknown() && !data.Vendorid.IsNull() && !data.Vendorid.Equal(state.Vendorid) {
		tflog.Debug(ctx, fmt.Sprintf("vendorid has changed for lbmonitor"))
		lbmonitor.Vendorid = full.Vendorid
		hasChange = true
	}
	if !data.Vendorspecificacctapplicationids.IsUnknown() && !data.Vendorspecificacctapplicationids.IsNull() && !data.Vendorspecificacctapplicationids.Equal(state.Vendorspecificacctapplicationids) {
		tflog.Debug(ctx, fmt.Sprintf("vendorspecificacctapplicationids has changed for lbmonitor"))
		lbmonitor.Vendorspecificacctapplicationids = full.Vendorspecificacctapplicationids
		hasChange = true
	}
	if !data.Vendorspecificauthapplicationids.IsUnknown() && !data.Vendorspecificauthapplicationids.IsNull() && !data.Vendorspecificauthapplicationids.Equal(state.Vendorspecificauthapplicationids) {
		tflog.Debug(ctx, fmt.Sprintf("vendorspecificauthapplicationids has changed for lbmonitor"))
		lbmonitor.Vendorspecificauthapplicationids = full.Vendorspecificauthapplicationids
		hasChange = true
	}
	if !data.Vendorspecificvendorid.IsUnknown() && !data.Vendorspecificvendorid.IsNull() && !data.Vendorspecificvendorid.Equal(state.Vendorspecificvendorid) {
		tflog.Debug(ctx, fmt.Sprintf("vendorspecificvendorid has changed for lbmonitor"))
		lbmonitor.Vendorspecificvendorid = full.Vendorspecificvendorid
		hasChange = true
	}

	if hasChange {
		// `lbmonitor` already holds monitorname + type + only the changed fields (delta payload built above).
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

	// Read the updated state back
	r.readLbmonitorFromApi(ctx, &data, &resp.Diagnostics)

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
	// Legacy SDK v2 id is the plain monitorname (d.SetId(monitorname)); pass the
	// legacyOrder so a legacy id parses. New-format ids ("monitorname:..,type:..")
	// are auto-detected and parsed by key regardless.
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
func (r *LbmonitorResource) readLbmonitorFromApi(ctx context.Context, data *LbmonitorResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	// Legacy SDK v2 id is the plain monitorname (d.SetId(monitorname)); pass the
	// legacyOrder so a legacy id parses. New-format ids ("monitorname:..,type:..")
	// are auto-detected and parsed by key regardless.
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"monitorname"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	monitorname_Name, ok := idMap["monitorname"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'monitorname' not found in ID string")
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Lbmonitor.Type(),
		ResourceName:             monitorname_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lbmonitor, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "lbmonitor returned empty array.")
		return
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
		diags.AddError("Client Error", fmt.Sprintf("lbmonitor not found with the provided ID attributes"))
		return
	}

	lbmonitorSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
