package gslbservicegroup

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/gslb"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// GslbservicegroupResourceModel describes the resource data model.
type GslbservicegroupResourceModel struct {
	Id               types.String `tfsdk:"id"`
	Appflowlog       types.String `tfsdk:"appflowlog"`
	Autodelayedtrofs types.String `tfsdk:"autodelayedtrofs"`
	Autoscale        types.String `tfsdk:"autoscale"`
	Cip              types.String `tfsdk:"cip"`
	Cipheader        types.String `tfsdk:"cipheader"`
	Clttimeout       types.Int64  `tfsdk:"clttimeout"`
	Comment          types.String `tfsdk:"comment"`
	Delay            types.Int64  `tfsdk:"delay"`
	Downstateflush   types.String `tfsdk:"downstateflush"`
	DupWeight        types.Int64  `tfsdk:"dup_weight"`
	Graceful         types.String `tfsdk:"graceful"`
	Hashid           types.Int64  `tfsdk:"hashid"`
	Healthmonitor    types.String `tfsdk:"healthmonitor"`
	Includemembers   types.Bool   `tfsdk:"includemembers"`
	Maxbandwidth     types.Int64  `tfsdk:"maxbandwidth"`
	Maxclient        types.Int64  `tfsdk:"maxclient"`
	MonitorNameSvc   types.String `tfsdk:"monitor_name_svc"`
	Monthreshold     types.Int64  `tfsdk:"monthreshold"`
	Newname          types.String `tfsdk:"newname"`
	Order            types.Int64  `tfsdk:"order"`
	Port             types.Int64  `tfsdk:"port"`
	Publicip         types.String `tfsdk:"publicip"`
	Publicport       types.Int64  `tfsdk:"publicport"`
	Servername       types.String `tfsdk:"servername"`
	Servicegroupname types.String `tfsdk:"servicegroupname"`
	Servicetype      types.String `tfsdk:"servicetype"`
	Sitename         types.String `tfsdk:"sitename"`
	Sitepersistence  types.String `tfsdk:"sitepersistence"`
	Siteprefix       types.String `tfsdk:"siteprefix"`
	State            types.String `tfsdk:"state"`
	Svrtimeout       types.Int64  `tfsdk:"svrtimeout"`
	Weight           types.Int64  `tfsdk:"weight"`
}

func (r *GslbservicegroupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the gslbservicegroup resource.",
			},
			"appflowlog": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable logging of AppFlow information for the specified GSLB service group.",
			},
			"autodelayedtrofs": schema.StringAttribute{
				// NITRO does not echo this field in a bare GET, so it cannot be
				// Computed (it would stay unknown-after-apply). Optional only.
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Indicates graceful movement of the service to TROFS. System will wait for monitor response time out before moving to TROFS",
			},
			"autoscale": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Auto scale option for a GSLB servicegroup",
			},
			"cip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Insert the Client IP header in requests forwarded to the GSLB service.",
			},
			"cipheader": schema.StringAttribute{
				// NITRO does not echo this field in a bare GET, so it cannot be
				// Computed (it would stay unknown-after-apply). Optional only.
				Optional:    true,
				Description: "Name of the HTTP header whose value must be set to the IP address of the client. Used with the Client IP parameter. If client IP insertion is enabled, and the client IP header is not specified, the value of Client IP Header parameter or the value set by the set ns config command is used as client's IP header name.",
			},
			"clttimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in seconds, after which to terminate an idle client connection.",
			},
			"comment": schema.StringAttribute{
				// NITRO does not echo this field in a bare GET, so it cannot be
				// Computed (it would stay unknown-after-apply). Optional only.
				Optional:    true,
				Description: "Any information about the GSLB service group.",
			},
			"delay": schema.Int64Attribute{
				Optional:    true,
				Description: "The time allowed (in seconds) for a graceful shutdown. During this period, new connections or requests will continue to be sent to this service for clients who already have a persistent session on the system. Connections or requests from fresh or new clients who do not yet have a persistence sessions on the system will not be sent to the service. Instead, they will be load balanced among other available services. After the delay time expires, no new requests or connections will be sent to the service.",
			},
			"downstateflush": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Flush all active transactions associated with all the services in the GSLB service group whose state transitions from UP to DOWN. Do not enable this option for applications that must complete their transactions.",
			},
			"dup_weight": schema.Int64Attribute{
				Optional:    true,
				Description: "weight of the monitor that is bound to GSLB servicegroup.",
			},
			"graceful": schema.StringAttribute{
				Optional:    true,
				Description: "Wait for all existing connections to the service to terminate before shutting down the service.",
			},
			"hashid": schema.Int64Attribute{
				Optional:    true,
				Description: "The hash identifier for the service. This must be unique for each service. This parameter is used by hash based load balancing methods.",
			},
			"healthmonitor": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Monitor the health of this GSLB service.Available settings function are as follows:\nYES - Send probes to check the health of the GSLB service.\nNO - Do not send probes to check the health of the GSLB service. With the NO option, the appliance shows the service as UP at all times.",
			},
			"includemembers": schema.BoolAttribute{
				Optional:    true,
				Description: "Display the members of the listed GSLB service groups in addition to their settings. Can be specified when no service group name is provided in the command. In that case, the details displayed for each service group are identical to the details displayed when a service group name is provided, except that bound monitors are not displayed.",
			},
			"maxbandwidth": schema.Int64Attribute{
				// Echoed by GET with a server default; Computed so null->server-value is legal.
				Optional:    true,
				Computed:    true,
				Description: "Maximum bandwidth, in Kbps, allocated for all the services in the GSLB service group.",
			},
			"maxclient": schema.Int64Attribute{
				// Echoed by GET with a server default; Computed so null->server-value is legal.
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of simultaneous open connections for the GSLB service group.",
			},
			"monitor_name_svc": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the monitor bound to the GSLB service group. Used to assign a weight to the monitor.",
			},
			"monthreshold": schema.Int64Attribute{
				// Echoed by GET with a server default; Computed so null->server-value is legal.
				Optional:    true,
				Computed:    true,
				Description: "Minimum sum of weights of the monitors that are bound to this GSLB service. Used to determine whether to mark a GSLB service as UP or DOWN.",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Description: "New name for the GSLB service group.",
			},
			"order": schema.Int64Attribute{
				Optional:    true,
				Description: "Order number to be assigned to the gslb servicegroup member",
			},
			"port": schema.Int64Attribute{
				Optional:    true,
				Description: "Server port number.",
			},
			"publicip": schema.StringAttribute{
				// Echoed by GET with a server default (0.0.0.0); Computed so null->server-value is legal.
				Optional:    true,
				Computed:    true,
				Description: "The public IP address that a NAT device translates to the GSLB service's private IP address. Optional.",
			},
			"publicport": schema.Int64Attribute{
				Optional:    true,
				Description: "The public port associated with the GSLB service's public IP address. The port is mapped to the service's private port number. Applicable to the local GSLB service. Optional.",
			},
			"servername": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the server to which to bind the service group.",
			},
			"servicegroupname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the GSLB service group. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Can be changed after the name is created.",
			},
			"servicetype": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Protocol used to exchange data with the GSLB service.",
			},
			"sitename": schema.StringAttribute{
				Required:    true,
				Description: "Name of the GSLB site to which the service group belongs.",
			},
			"sitepersistence": schema.StringAttribute{
				// NITRO does not echo this field in a bare GET for HTTP non-autoscale
				// groups, so it cannot be Computed (would stay unknown-after-apply).
				Optional:    true,
				Description: "Use cookie-based site persistence. Applicable only to HTTP and SSL non-autoscale enabled GSLB servicegroups.",
			},
			"siteprefix": schema.StringAttribute{
				Optional:    true,
				Description: "The site's prefix string. When the GSLB service group is bound to a GSLB virtual server, a GSLB site domain is generated internally for each bound serviceitem-domain pair by concatenating the site prefix of the service item and the name of the domain. If the special string NONE is specified, the site-prefix string is unset. When implementing HTTP redirect site persistence, the Citrix ADC redirects GSLB requests to GSLB services by using their site domains.",
			},
			"state": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Initial state of the GSLB service group.",
			},
			"svrtimeout": schema.Int64Attribute{
				// Echoed by GET with a server default (2000); Computed so null->server-value is legal.
				Optional:    true,
				Computed:    true,
				Description: "Time, in seconds, after which to terminate an idle server connection.",
			},
			"weight": schema.Int64Attribute{
				Optional:    true,
				Description: "Weight to assign to the servers in the service group. Specifies the capacity of the servers relative to the other servers in the load balancing configuration. The higher the weight, the higher the percentage of requests sent to the service.",
			},
		},
	}
}

// gslbservicegroupGetThePayloadFromthePlan builds the ADD (create) payload.
// It carries the full create-time attribute set including the create-only attrs
// (servicetype, autoscale, autodelayedtrofs, state). It EXCLUDES the rename-only
// newname (Pattern: Rename support), the GET-only filter includemembers (Pattern
// 15), and the disable-action-only delay/graceful attrs.
func gslbservicegroupGetThePayloadFromthePlan(ctx context.Context, data *GslbservicegroupResourceModel) gslb.Gslbservicegroup {
	tflog.Debug(ctx, "In gslbservicegroupGetThePayloadFromthePlan Function")

	gslbservicegroup := gslb.Gslbservicegroup{}
	if !data.Appflowlog.IsNull() && !data.Appflowlog.IsUnknown() {
		gslbservicegroup.Appflowlog = data.Appflowlog.ValueString()
	}
	// autodelayedtrofs is create-only (RequiresReplace) - included in add payload only.
	if !data.Autodelayedtrofs.IsNull() && !data.Autodelayedtrofs.IsUnknown() {
		gslbservicegroup.Autodelayedtrofs = data.Autodelayedtrofs.ValueString()
	}
	// autoscale is create-only (RequiresReplace) - included in add payload only.
	if !data.Autoscale.IsNull() && !data.Autoscale.IsUnknown() {
		gslbservicegroup.Autoscale = data.Autoscale.ValueString()
	}
	if !data.Cip.IsNull() && !data.Cip.IsUnknown() {
		gslbservicegroup.Cip = data.Cip.ValueString()
	}
	if !data.Cipheader.IsNull() && !data.Cipheader.IsUnknown() {
		gslbservicegroup.Cipheader = data.Cipheader.ValueString()
	}
	if !data.Clttimeout.IsNull() && !data.Clttimeout.IsUnknown() {
		gslbservicegroup.Clttimeout = utils.IntPtr(int(data.Clttimeout.ValueInt64()))
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		gslbservicegroup.Comment = data.Comment.ValueString()
	}
	// delay is disable-action-only - excluded from add/set payloads.
	if !data.Downstateflush.IsNull() && !data.Downstateflush.IsUnknown() {
		gslbservicegroup.Downstateflush = data.Downstateflush.ValueString()
	}
	if !data.DupWeight.IsNull() && !data.DupWeight.IsUnknown() {
		gslbservicegroup.Dupweight = utils.IntPtr(int(data.DupWeight.ValueInt64()))
	}
	// graceful is disable-action-only - excluded from add/set payloads.
	if !data.Hashid.IsNull() && !data.Hashid.IsUnknown() {
		gslbservicegroup.Hashid = utils.IntPtr(int(data.Hashid.ValueInt64()))
	}
	if !data.Healthmonitor.IsNull() && !data.Healthmonitor.IsUnknown() {
		gslbservicegroup.Healthmonitor = data.Healthmonitor.ValueString()
	}
	// includemembers is a GET-only filter (Pattern 15) - excluded from add/set payloads.
	if !data.Maxbandwidth.IsNull() && !data.Maxbandwidth.IsUnknown() {
		gslbservicegroup.Maxbandwidth = utils.IntPtr(int(data.Maxbandwidth.ValueInt64()))
	}
	if !data.Maxclient.IsNull() && !data.Maxclient.IsUnknown() {
		gslbservicegroup.Maxclient = utils.IntPtr(int(data.Maxclient.ValueInt64()))
	}
	if !data.MonitorNameSvc.IsNull() && !data.MonitorNameSvc.IsUnknown() {
		gslbservicegroup.Monitornamesvc = data.MonitorNameSvc.ValueString()
	}
	if !data.Monthreshold.IsNull() && !data.Monthreshold.IsUnknown() {
		gslbservicegroup.Monthreshold = utils.IntPtr(int(data.Monthreshold.ValueInt64()))
	}
	// newname is rename-only - excluded from add/set payloads (wired in Update via ?action=rename).
	if !data.Order.IsNull() && !data.Order.IsUnknown() {
		gslbservicegroup.Order = utils.IntPtr(int(data.Order.ValueInt64()))
	}
	if !data.Port.IsNull() && !data.Port.IsUnknown() {
		gslbservicegroup.Port = utils.IntPtr(int(data.Port.ValueInt64()))
	}
	if !data.Publicip.IsNull() && !data.Publicip.IsUnknown() {
		gslbservicegroup.Publicip = data.Publicip.ValueString()
	}
	if !data.Publicport.IsNull() && !data.Publicport.IsUnknown() {
		gslbservicegroup.Publicport = utils.IntPtr(int(data.Publicport.ValueInt64()))
	}
	if !data.Servername.IsNull() && !data.Servername.IsUnknown() {
		gslbservicegroup.Servername = data.Servername.ValueString()
	}
	if !data.Servicegroupname.IsNull() && !data.Servicegroupname.IsUnknown() {
		gslbservicegroup.Servicegroupname = data.Servicegroupname.ValueString()
	}
	// servicetype is create-only (RequiresReplace) - included in add payload only.
	if !data.Servicetype.IsNull() && !data.Servicetype.IsUnknown() {
		gslbservicegroup.Servicetype = data.Servicetype.ValueString()
	}
	if !data.Sitename.IsNull() && !data.Sitename.IsUnknown() {
		gslbservicegroup.Sitename = data.Sitename.ValueString()
	}
	if !data.Sitepersistence.IsNull() && !data.Sitepersistence.IsUnknown() {
		gslbservicegroup.Sitepersistence = data.Sitepersistence.ValueString()
	}
	if !data.Siteprefix.IsNull() && !data.Siteprefix.IsUnknown() {
		gslbservicegroup.Siteprefix = data.Siteprefix.ValueString()
	}
	// state is create-only via add (NOT accepted in PUT update) - included in add payload only.
	if !data.State.IsNull() && !data.State.IsUnknown() {
		gslbservicegroup.State = data.State.ValueString()
	}
	if !data.Svrtimeout.IsNull() && !data.Svrtimeout.IsUnknown() {
		gslbservicegroup.Svrtimeout = utils.IntPtr(int(data.Svrtimeout.ValueInt64()))
	}
	if !data.Weight.IsNull() && !data.Weight.IsUnknown() {
		gslbservicegroup.Weight = utils.IntPtr(int(data.Weight.ValueInt64()))
	}

	return gslbservicegroup
}

// gslbservicegroupGetTheUpdatePayloadFromthePlan builds the SET (update/PUT) payload.
// Pattern 9 (add-vs-set payload drift): it EXCLUDES the create-only attrs
// (servicetype, autoscale, autodelayedtrofs, state) and the non-write attrs
// (newname, includemembers, delay, graceful). It carries only the genuinely
// updateable attributes.
func gslbservicegroupGetTheUpdatePayloadFromthePlan(ctx context.Context, data *GslbservicegroupResourceModel) gslb.Gslbservicegroup {
	tflog.Debug(ctx, "In gslbservicegroupGetTheUpdatePayloadFromthePlan Function")

	gslbservicegroup := gslb.Gslbservicegroup{}
	// servicegroupname is the name key, required to address the resource in PUT.
	if !data.Servicegroupname.IsNull() && !data.Servicegroupname.IsUnknown() {
		gslbservicegroup.Servicegroupname = data.Servicegroupname.ValueString()
	}
	if !data.Appflowlog.IsNull() && !data.Appflowlog.IsUnknown() {
		gslbservicegroup.Appflowlog = data.Appflowlog.ValueString()
	}
	if !data.Cip.IsNull() && !data.Cip.IsUnknown() {
		gslbservicegroup.Cip = data.Cip.ValueString()
	}
	if !data.Cipheader.IsNull() && !data.Cipheader.IsUnknown() {
		gslbservicegroup.Cipheader = data.Cipheader.ValueString()
	}
	if !data.Clttimeout.IsNull() && !data.Clttimeout.IsUnknown() {
		gslbservicegroup.Clttimeout = utils.IntPtr(int(data.Clttimeout.ValueInt64()))
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		gslbservicegroup.Comment = data.Comment.ValueString()
	}
	if !data.Downstateflush.IsNull() && !data.Downstateflush.IsUnknown() {
		gslbservicegroup.Downstateflush = data.Downstateflush.ValueString()
	}
	if !data.DupWeight.IsNull() && !data.DupWeight.IsUnknown() {
		gslbservicegroup.Dupweight = utils.IntPtr(int(data.DupWeight.ValueInt64()))
	}
	if !data.Hashid.IsNull() && !data.Hashid.IsUnknown() {
		gslbservicegroup.Hashid = utils.IntPtr(int(data.Hashid.ValueInt64()))
	}
	if !data.Healthmonitor.IsNull() && !data.Healthmonitor.IsUnknown() {
		gslbservicegroup.Healthmonitor = data.Healthmonitor.ValueString()
	}
	if !data.Maxbandwidth.IsNull() && !data.Maxbandwidth.IsUnknown() {
		gslbservicegroup.Maxbandwidth = utils.IntPtr(int(data.Maxbandwidth.ValueInt64()))
	}
	if !data.Maxclient.IsNull() && !data.Maxclient.IsUnknown() {
		gslbservicegroup.Maxclient = utils.IntPtr(int(data.Maxclient.ValueInt64()))
	}
	if !data.MonitorNameSvc.IsNull() && !data.MonitorNameSvc.IsUnknown() {
		gslbservicegroup.Monitornamesvc = data.MonitorNameSvc.ValueString()
	}
	if !data.Monthreshold.IsNull() && !data.Monthreshold.IsUnknown() {
		gslbservicegroup.Monthreshold = utils.IntPtr(int(data.Monthreshold.ValueInt64()))
	}
	if !data.Order.IsNull() && !data.Order.IsUnknown() {
		gslbservicegroup.Order = utils.IntPtr(int(data.Order.ValueInt64()))
	}
	if !data.Port.IsNull() && !data.Port.IsUnknown() {
		gslbservicegroup.Port = utils.IntPtr(int(data.Port.ValueInt64()))
	}
	if !data.Publicip.IsNull() && !data.Publicip.IsUnknown() {
		gslbservicegroup.Publicip = data.Publicip.ValueString()
	}
	if !data.Publicport.IsNull() && !data.Publicport.IsUnknown() {
		gslbservicegroup.Publicport = utils.IntPtr(int(data.Publicport.ValueInt64()))
	}
	if !data.Servername.IsNull() && !data.Servername.IsUnknown() {
		gslbservicegroup.Servername = data.Servername.ValueString()
	}
	if !data.Sitename.IsNull() && !data.Sitename.IsUnknown() {
		gslbservicegroup.Sitename = data.Sitename.ValueString()
	}
	if !data.Sitepersistence.IsNull() && !data.Sitepersistence.IsUnknown() {
		gslbservicegroup.Sitepersistence = data.Sitepersistence.ValueString()
	}
	if !data.Siteprefix.IsNull() && !data.Siteprefix.IsUnknown() {
		gslbservicegroup.Siteprefix = data.Siteprefix.ValueString()
	}
	if !data.Svrtimeout.IsNull() && !data.Svrtimeout.IsUnknown() {
		gslbservicegroup.Svrtimeout = utils.IntPtr(int(data.Svrtimeout.ValueInt64()))
	}
	if !data.Weight.IsNull() && !data.Weight.IsUnknown() {
		gslbservicegroup.Weight = utils.IntPtr(int(data.Weight.ValueInt64()))
	}

	return gslbservicegroup
}

// gslbservicegroupSetAttrFromGet populates the resource model from a GET response.
// It preserves the user-configured servicegroupname (only adopting the GET value
// when the model key is empty, e.g. on import) so a rename does not clobber the
// configured value. It reads back the server-defaulted echoed attributes
// (healthmonitor, state, downstateflush, appflowlog, autoscale, autodelayedtrofs)
// and the other echoed attributes. Attributes that are not reliably echoed by a
// bare group GET (Optional-only in the schema) are preserved from the plan/state
// and NOT touched here, to avoid unknown-after-apply / inconsistent-result churn.
func gslbservicegroupSetAttrFromGet(ctx context.Context, data *GslbservicegroupResourceModel, getResponseData map[string]interface{}) *GslbservicegroupResourceModel {
	tflog.Debug(ctx, "In gslbservicegroupSetAttrFromGet Function")

	if val, ok := getResponseData["appflowlog"]; ok && val != nil {
		data.Appflowlog = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["autodelayedtrofs"]; ok && val != nil {
		data.Autodelayedtrofs = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["autoscale"]; ok && val != nil {
		data.Autoscale = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["cip"]; ok && val != nil {
		data.Cip = types.StringValue(val.(string))
	}
	// cipheader is Optional-only and not always echoed (the server auto-populates it
	// to a default like "client-ip" once cip is ENABLED). Only adopt the GET value
	// when the user configured it; otherwise preserve plan/state to avoid an
	// inconsistent-result-after-apply on a null->server-default transition.
	if !data.Cipheader.IsNull() && !data.Cipheader.IsUnknown() {
		if val, ok := getResponseData["cipheader"]; ok && val != nil {
			data.Cipheader = types.StringValue(val.(string))
		}
	}
	if val, ok := getResponseData["clttimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Clttimeout = types.Int64Value(intVal)
		}
	}
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["downstateflush"]; ok && val != nil {
		data.Downstateflush = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["healthmonitor"]; ok && val != nil {
		data.Healthmonitor = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["maxbandwidth"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxbandwidth = types.Int64Value(intVal)
		}
	}
	if val, ok := getResponseData["maxclient"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxclient = types.Int64Value(intVal)
		}
	}
	if val, ok := getResponseData["monthreshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Monthreshold = types.Int64Value(intVal)
		}
	}
	if val, ok := getResponseData["publicip"]; ok && val != nil {
		data.Publicip = types.StringValue(val.(string))
	}
	// servicegroupname: preserve the configured value; only adopt the GET value
	// when the model key is empty (import). This prevents a rename from clobbering
	// the user-configured name.
	if data.Servicegroupname.IsNull() || data.Servicegroupname.ValueString() == "" {
		if val, ok := getResponseData["servicegroupname"]; ok && val != nil {
			data.Servicegroupname = types.StringValue(val.(string))
		}
	}
	if val, ok := getResponseData["servicetype"]; ok && val != nil {
		data.Servicetype = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["sitename"]; ok && val != nil {
		data.Sitename = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["sitepersistence"]; ok && val != nil {
		data.Sitepersistence = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["svrtimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Svrtimeout = types.Int64Value(intVal)
		}
	}

	return data
}

// gslbservicegroupSetAttrFromGetForDatasource faithfully copies every field from
// the GET response into the model and composes the datasource ID. Unlike the
// resource setter it does not preserve any prior plan/state (a datasource has none).
func gslbservicegroupSetAttrFromGetForDatasource(ctx context.Context, data *GslbservicegroupResourceModel, getResponseData map[string]interface{}) *GslbservicegroupResourceModel {
	tflog.Debug(ctx, "In gslbservicegroupSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["appflowlog"]; ok && val != nil {
		data.Appflowlog = types.StringValue(val.(string))
	} else {
		data.Appflowlog = types.StringNull()
	}
	if val, ok := getResponseData["autodelayedtrofs"]; ok && val != nil {
		data.Autodelayedtrofs = types.StringValue(val.(string))
	} else {
		data.Autodelayedtrofs = types.StringNull()
	}
	if val, ok := getResponseData["autoscale"]; ok && val != nil {
		data.Autoscale = types.StringValue(val.(string))
	} else {
		data.Autoscale = types.StringNull()
	}
	if val, ok := getResponseData["cip"]; ok && val != nil {
		data.Cip = types.StringValue(val.(string))
	} else {
		data.Cip = types.StringNull()
	}
	if val, ok := getResponseData["cipheader"]; ok && val != nil {
		data.Cipheader = types.StringValue(val.(string))
	} else {
		data.Cipheader = types.StringNull()
	}
	if val, ok := getResponseData["clttimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Clttimeout = types.Int64Value(intVal)
		}
	} else {
		data.Clttimeout = types.Int64Null()
	}
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["downstateflush"]; ok && val != nil {
		data.Downstateflush = types.StringValue(val.(string))
	} else {
		data.Downstateflush = types.StringNull()
	}
	if val, ok := getResponseData["dup_weight"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.DupWeight = types.Int64Value(intVal)
		}
	} else {
		data.DupWeight = types.Int64Null()
	}
	if val, ok := getResponseData["hashid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Hashid = types.Int64Value(intVal)
		}
	} else {
		data.Hashid = types.Int64Null()
	}
	if val, ok := getResponseData["healthmonitor"]; ok && val != nil {
		data.Healthmonitor = types.StringValue(val.(string))
	} else {
		data.Healthmonitor = types.StringNull()
	}
	if val, ok := getResponseData["maxbandwidth"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxbandwidth = types.Int64Value(intVal)
		}
	} else {
		data.Maxbandwidth = types.Int64Null()
	}
	if val, ok := getResponseData["maxclient"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxclient = types.Int64Value(intVal)
		}
	} else {
		data.Maxclient = types.Int64Null()
	}
	if val, ok := getResponseData["monitor_name_svc"]; ok && val != nil {
		data.MonitorNameSvc = types.StringValue(val.(string))
	} else {
		data.MonitorNameSvc = types.StringNull()
	}
	if val, ok := getResponseData["monthreshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Monthreshold = types.Int64Value(intVal)
		}
	} else {
		data.Monthreshold = types.Int64Null()
	}
	if val, ok := getResponseData["order"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Order = types.Int64Value(intVal)
		}
	} else {
		data.Order = types.Int64Null()
	}
	if val, ok := getResponseData["port"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Port = types.Int64Value(intVal)
		}
	} else {
		data.Port = types.Int64Null()
	}
	if val, ok := getResponseData["publicip"]; ok && val != nil {
		data.Publicip = types.StringValue(val.(string))
	} else {
		data.Publicip = types.StringNull()
	}
	if val, ok := getResponseData["publicport"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Publicport = types.Int64Value(intVal)
		}
	} else {
		data.Publicport = types.Int64Null()
	}
	if val, ok := getResponseData["servername"]; ok && val != nil {
		data.Servername = types.StringValue(val.(string))
	} else {
		data.Servername = types.StringNull()
	}
	if val, ok := getResponseData["servicegroupname"]; ok && val != nil {
		data.Servicegroupname = types.StringValue(val.(string))
	} else {
		data.Servicegroupname = types.StringNull()
	}
	if val, ok := getResponseData["servicetype"]; ok && val != nil {
		data.Servicetype = types.StringValue(val.(string))
	} else {
		data.Servicetype = types.StringNull()
	}
	if val, ok := getResponseData["sitename"]; ok && val != nil {
		data.Sitename = types.StringValue(val.(string))
	} else {
		data.Sitename = types.StringNull()
	}
	if val, ok := getResponseData["sitepersistence"]; ok && val != nil {
		data.Sitepersistence = types.StringValue(val.(string))
	} else {
		data.Sitepersistence = types.StringNull()
	}
	if val, ok := getResponseData["siteprefix"]; ok && val != nil {
		data.Siteprefix = types.StringValue(val.(string))
	} else {
		data.Siteprefix = types.StringNull()
	}
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	} else {
		data.State = types.StringNull()
	}
	if val, ok := getResponseData["svrtimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Svrtimeout = types.Int64Value(intVal)
		}
	} else {
		data.Svrtimeout = types.Int64Null()
	}
	if val, ok := getResponseData["weight"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Weight = types.Int64Value(intVal)
		}
	} else {
		data.Weight = types.Int64Null()
	}

	// Set ID for the datasource (named resource - plain servicegroupname value).
	data.Id = types.StringValue(data.Servicegroupname.ValueString())

	return data
}
