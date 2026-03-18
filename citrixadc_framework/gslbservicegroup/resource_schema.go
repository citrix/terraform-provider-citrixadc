package gslbservicegroup

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/gslb"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
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
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Enable logging of AppFlow information for the specified GSLB service group.",
			},
			"autodelayedtrofs": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Indicates graceful movement of the service to TROFS. System will wait for monitor response time out before moving to TROFS",
			},
			"autoscale": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Auto scale option for a GSLB servicegroup",
			},
			"cip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Insert the Client IP header in requests forwarded to the GSLB service.",
			},
			"cipheader": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the HTTP header whose value must be set to the IP address of the client. Used with the Client IP parameter. If client IP insertion is enabled, and the client IP header is not specified, the value of Client IP Header parameter or the value set by the set ns config command is used as client's IP header name.",
			},
			"clttimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in seconds, after which to terminate an idle client connection.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any information about the GSLB service group.",
			},
			"delay": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "The time allowed (in seconds) for a graceful shutdown. During this period, new connections or requests will continue to be sent to this service for clients who already have a persistent session on the system. Connections or requests from fresh or new clients who do not yet have a persistence sessions on the system will not be sent to the service. Instead, they will be load balanced among other available services. After the delay time expires, no new requests or connections will be sent to the service.",
			},
			"downstateflush": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Flush all active transactions associated with all the services in the GSLB service group whose state transitions from UP to DOWN. Do not enable this option for applications that must complete their transactions.",
			},
			"dup_weight": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "weight of the monitor that is bound to GSLB servicegroup.",
			},
			"graceful": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Wait for all existing connections to the service to terminate before shutting down the service.",
			},
			"hashid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The hash identifier for the service. This must be unique for each service. This parameter is used by hash based load balancing methods.",
			},
			"healthmonitor": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("True"),
				Description: "Monitor the health of this GSLB service.Available settings function are as follows:\nYES - Send probes to check the health of the GSLB service.\nNO - Do not send probes to check the health of the GSLB service. With the NO option, the appliance shows the service as UP at all times.",
			},
			"includemembers": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Display the members of the listed GSLB service groups in addition to their settings. Can be specified when no service group name is provided in the command. In that case, the details displayed for each service group are identical to the details displayed when a service group name is provided, except that bound monitors are not displayed.",
			},
			"maxbandwidth": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum bandwidth, in Kbps, allocated for all the services in the GSLB service group.",
			},
			"maxclient": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of simultaneous open connections for the GSLB service group.",
			},
			"monitor_name_svc": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the monitor bound to the GSLB service group. Used to assign a weight to the monitor.",
			},
			"monthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum sum of weights of the monitors that are bound to this GSLB service. Used to determine whether to mark a GSLB service as UP or DOWN.",
			},
			"newname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "New name for the GSLB service group.",
			},
			"order": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Order number to be assigned to the gslb servicegroup member",
			},
			"port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Server port number.",
			},
			"publicip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The public IP address that a NAT device translates to the GSLB service's private IP address. Optional.",
			},
			"publicport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The public port associated with the GSLB service's public IP address. The port is mapped to the service's private port number. Applicable to the local GSLB service. Optional.",
			},
			"servername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the server to which to bind the service group.",
			},
			"servicegroupname": schema.StringAttribute{
				Required:    true,
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
				Optional:    true,
				Computed:    true,
				Description: "Use cookie-based site persistence. Applicable only to HTTP and SSL non-autoscale enabled GSLB servicegroups.",
			},
			"siteprefix": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The site's prefix string. When the GSLB service group is bound to a GSLB virtual server, a GSLB site domain is generated internally for each bound serviceitem-domain pair by concatenating the site prefix of the service item and the name of the domain. If the special string NONE is specified, the site-prefix string is unset. When implementing HTTP redirect site persistence, the Citrix ADC redirects GSLB requests to GSLB services by using their site domains.",
			},
			"state": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Initial state of the GSLB service group.",
			},
			"svrtimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in seconds, after which to terminate an idle server connection.",
			},
			"weight": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Weight to assign to the servers in the service group. Specifies the capacity of the servers relative to the other servers in the load balancing configuration. The higher the weight, the higher the percentage of requests sent to the service.",
			},
		},
	}
}

func gslbservicegroupGetThePayloadFromtheConfig(ctx context.Context, data *GslbservicegroupResourceModel) gslb.Gslbservicegroup {
	tflog.Debug(ctx, "In gslbservicegroupGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	gslbservicegroup := gslb.Gslbservicegroup{}
	if !data.Appflowlog.IsNull() {
		gslbservicegroup.Appflowlog = data.Appflowlog.ValueString()
	}
	if !data.Autodelayedtrofs.IsNull() {
		gslbservicegroup.Autodelayedtrofs = data.Autodelayedtrofs.ValueString()
	}
	if !data.Autoscale.IsNull() {
		gslbservicegroup.Autoscale = data.Autoscale.ValueString()
	}
	if !data.Cip.IsNull() {
		gslbservicegroup.Cip = data.Cip.ValueString()
	}
	if !data.Cipheader.IsNull() {
		gslbservicegroup.Cipheader = data.Cipheader.ValueString()
	}
	if !data.Clttimeout.IsNull() {
		gslbservicegroup.Clttimeout = utils.IntPtr(int(data.Clttimeout.ValueInt64()))
	}
	if !data.Comment.IsNull() {
		gslbservicegroup.Comment = data.Comment.ValueString()
	}
	if !data.Delay.IsNull() {
		gslbservicegroup.Delay = utils.IntPtr(int(data.Delay.ValueInt64()))
	}
	if !data.Downstateflush.IsNull() {
		gslbservicegroup.Downstateflush = data.Downstateflush.ValueString()
	}
	if !data.DupWeight.IsNull() {
		gslbservicegroup.Dupweight = utils.IntPtr(int(data.DupWeight.ValueInt64()))
	}
	if !data.Graceful.IsNull() {
		gslbservicegroup.Graceful = data.Graceful.ValueString()
	}
	if !data.Hashid.IsNull() {
		gslbservicegroup.Hashid = utils.IntPtr(int(data.Hashid.ValueInt64()))
	}
	if !data.Healthmonitor.IsNull() {
		gslbservicegroup.Healthmonitor = data.Healthmonitor.ValueString()
	}
	if !data.Includemembers.IsNull() {
		gslbservicegroup.Includemembers = data.Includemembers.ValueBool()
	}
	if !data.Maxbandwidth.IsNull() {
		gslbservicegroup.Maxbandwidth = utils.IntPtr(int(data.Maxbandwidth.ValueInt64()))
	}
	if !data.Maxclient.IsNull() {
		gslbservicegroup.Maxclient = utils.IntPtr(int(data.Maxclient.ValueInt64()))
	}
	if !data.MonitorNameSvc.IsNull() {
		gslbservicegroup.Monitornamesvc = data.MonitorNameSvc.ValueString()
	}
	if !data.Monthreshold.IsNull() {
		gslbservicegroup.Monthreshold = utils.IntPtr(int(data.Monthreshold.ValueInt64()))
	}
	if !data.Newname.IsNull() {
		gslbservicegroup.Newname = data.Newname.ValueString()
	}
	if !data.Order.IsNull() {
		gslbservicegroup.Order = utils.IntPtr(int(data.Order.ValueInt64()))
	}
	if !data.Port.IsNull() {
		gslbservicegroup.Port = utils.IntPtr(int(data.Port.ValueInt64()))
	}
	if !data.Publicip.IsNull() {
		gslbservicegroup.Publicip = data.Publicip.ValueString()
	}
	if !data.Publicport.IsNull() {
		gslbservicegroup.Publicport = utils.IntPtr(int(data.Publicport.ValueInt64()))
	}
	if !data.Servername.IsNull() {
		gslbservicegroup.Servername = data.Servername.ValueString()
	}
	if !data.Servicegroupname.IsNull() {
		gslbservicegroup.Servicegroupname = data.Servicegroupname.ValueString()
	}
	if !data.Servicetype.IsNull() {
		gslbservicegroup.Servicetype = data.Servicetype.ValueString()
	}
	if !data.Sitename.IsNull() {
		gslbservicegroup.Sitename = data.Sitename.ValueString()
	}
	if !data.Sitepersistence.IsNull() {
		gslbservicegroup.Sitepersistence = data.Sitepersistence.ValueString()
	}
	if !data.Siteprefix.IsNull() {
		gslbservicegroup.Siteprefix = data.Siteprefix.ValueString()
	}
	if !data.State.IsNull() {
		gslbservicegroup.State = data.State.ValueString()
	}
	if !data.Svrtimeout.IsNull() {
		gslbservicegroup.Svrtimeout = utils.IntPtr(int(data.Svrtimeout.ValueInt64()))
	}
	if !data.Weight.IsNull() {
		gslbservicegroup.Weight = utils.IntPtr(int(data.Weight.ValueInt64()))
	}

	return gslbservicegroup
}

func gslbservicegroupSetAttrFromGet(ctx context.Context, data *GslbservicegroupResourceModel, getResponseData map[string]interface{}) *GslbservicegroupResourceModel {
	tflog.Debug(ctx, "In gslbservicegroupSetAttrFromGet Function")

	// Convert API response to model
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
	if val, ok := getResponseData["delay"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Delay = types.Int64Value(intVal)
		}
	} else {
		data.Delay = types.Int64Null()
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
	if val, ok := getResponseData["graceful"]; ok && val != nil {
		data.Graceful = types.StringValue(val.(string))
	} else {
		data.Graceful = types.StringNull()
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
	if val, ok := getResponseData["includemembers"]; ok && val != nil {
		data.Includemembers = types.BoolValue(val.(bool))
	} else {
		data.Includemembers = types.BoolNull()
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
	if val, ok := getResponseData["newname"]; ok && val != nil {
		data.Newname = types.StringValue(val.(string))
	} else {
		data.Newname = types.StringNull()
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

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated
	data.Id = types.StringValue(data.Servicegroupname.ValueString())

	return data
}
