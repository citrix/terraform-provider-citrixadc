package gslbservice

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/gslb"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// GslbserviceResourceModel describes the resource data model.
type GslbserviceResourceModel struct {
	Id               types.String `tfsdk:"id"`
	Appflowlog       types.String `tfsdk:"appflowlog"`
	Cip              types.String `tfsdk:"cip"`
	Cipheader        types.String `tfsdk:"cipheader"`
	Clttimeout       types.Int64  `tfsdk:"clttimeout"`
	Cnameentry       types.String `tfsdk:"cnameentry"`
	Comment          types.String `tfsdk:"comment"`
	Cookietimeout    types.Int64  `tfsdk:"cookietimeout"`
	Downstateflush   types.String `tfsdk:"downstateflush"`
	Hashid           types.Int64  `tfsdk:"hashid"`
	Healthmonitor    types.String `tfsdk:"healthmonitor"`
	Ip               types.String `tfsdk:"ip"`
	Ipaddress        types.String `tfsdk:"ipaddress"`
	Maxaaausers      types.Int64  `tfsdk:"maxaaausers"`
	Maxbandwidth     types.Int64  `tfsdk:"maxbandwidth"`
	Maxclient        types.Int64  `tfsdk:"maxclient"`
	MonitorNameSvc   types.String `tfsdk:"monitor_name_svc"`
	Monthreshold     types.Int64  `tfsdk:"monthreshold"`
	Naptrdomainttl   types.Int64  `tfsdk:"naptrdomainttl"`
	Naptrorder       types.Int64  `tfsdk:"naptrorder"`
	Naptrpreference  types.Int64  `tfsdk:"naptrpreference"`
	Naptrreplacement types.String `tfsdk:"naptrreplacement"`
	Naptrservices    types.String `tfsdk:"naptrservices"`
	Newname          types.String `tfsdk:"newname"`
	Port             types.Int64  `tfsdk:"port"`
	Publicip         types.String `tfsdk:"publicip"`
	Publicport       types.Int64  `tfsdk:"publicport"`
	Servername       types.String `tfsdk:"servername"`
	Servicename      types.String `tfsdk:"servicename"`
	Servicetype      types.String `tfsdk:"servicetype"`
	Sitename         types.String `tfsdk:"sitename"`
	Sitepersistence  types.String `tfsdk:"sitepersistence"`
	Siteprefix       types.String `tfsdk:"siteprefix"`
	State            types.String `tfsdk:"state"`
	Svrtimeout       types.Int64  `tfsdk:"svrtimeout"`
	Viewip           types.String `tfsdk:"viewip"`
	Viewname         types.String `tfsdk:"viewname"`
	Weight           types.Int64  `tfsdk:"weight"`
}

func (r *GslbserviceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the gslbservice resource.",
			},
			"appflowlog": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Enable logging appflow flow information",
			},
			"cip": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "In the request that is forwarded to the GSLB service, insert a header that stores the client's IP address. Client IP header insertion is used in connection-proxy based site persistence.",
			},
			"cipheader": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name for the HTTP header that stores the client's IP address. Used with the Client IP option. If client IP header insertion is enabled on the service and a name is not specified for the header, the Citrix ADC uses the name specified by the cipHeader parameter in the set ns param command or, in the GUI, the Client IP Header parameter in the Configure HTTP Parameters dialog box.",
			},
			"clttimeout": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Idle time, in seconds, after which a client connection is terminated. Applicable if connection proxy based site persistence is used.",
			},
			"cnameentry": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Canonical name of the GSLB service. Used in CNAME-based GSLB.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments that you might want to associate with the GSLB service.",
			},
			"cookietimeout": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Timeout value, in minutes, for the cookie, when cookie based site persistence is enabled.",
			},
			"downstateflush": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Flush all active transactions associated with the GSLB service when its state transitions from UP to DOWN. Do not enable this option for services that must complete their transactions. Applicable if connection proxy based site persistence is used.",
			},
			"hashid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique hash identifier for the GSLB service, used by hash based load balancing methods.",
			},
			"healthmonitor": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("True"),
				Description: "Monitor the health of the GSLB service.",
			},
			"ip": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "IP address for the GSLB service. Should represent a load balancing, content switching, or VPN virtual server on the Citrix ADC, or the IP address of another load balancing device.",
			},
			"ipaddress": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The new IP address of the service.",
			},
			"maxaaausers": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of SSL VPN users that can be logged on concurrently to the VPN virtual server that is represented by this GSLB service. A GSLB service whose user count reaches the maximum is not considered when a GSLB decision is made, until the count drops below the maximum.",
			},
			"maxbandwidth": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer specifying the maximum bandwidth allowed for the service. A GSLB service whose bandwidth reaches the maximum is not considered when a GSLB decision is made, until its bandwidth consumption drops below the maximum.",
			},
			"maxclient": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The maximum number of open connections that the service can support at any given time. A GSLB service whose connection count reaches the maximum is not considered when a GSLB decision is made, until the connection count drops below the maximum.",
			},
			"monitor_name_svc": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the monitor to bind to the service.",
			},
			"monthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Monitoring threshold value for the GSLB service. If the sum of the weights of the monitors that are bound to this GSLB service and are in the UP state is not equal to or greater than this threshold value, the service is marked as DOWN.",
			},
			"naptrdomainttl": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(3600),
				Description: "Modify the TTL of the internally created naptr domain",
			},
			"naptrorder": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(1),
				Description: "An integer specifying the order in which the NAPTR records MUST be processed in order to accurately represent the ordered list of Rules. The ordering is from lowest to highest",
			},
			"naptrpreference": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(1),
				Description: "An integer specifying the preference of this NAPTR among NAPTR records having same order. lower the number, higher the preference.",
			},
			"naptrreplacement": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The replacement domain name for this NAPTR.",
			},
			"naptrservices": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Service Parameters applicable to this delegation path.",
			},
			"newname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "New name for the GSLB service.",
			},
			"port": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Port on which the load balancing entity represented by this GSLB service listens.",
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
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the server hosting the GSLB service.",
			},
			"servicename": schema.StringAttribute{
				Required:    true,
				Description: "Name for the GSLB service. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Can be changed after the GSLB service is created.\n\nCLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my gslbsvc\" or 'my gslbsvc').",
			},
			"servicetype": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("NSSVC_SERVICE_UNKNOWN"),
				Description: "Type of service to create.",
			},
			"sitename": schema.StringAttribute{
				Required:    true,
				Description: "Name of the GSLB site to which the service belongs.",
			},
			"sitepersistence": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use cookie-based site persistence. Applicable only to HTTP and SSL GSLB services.",
			},
			"siteprefix": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The site's prefix string. When the service is bound to a GSLB virtual server, a GSLB site domain is generated internally for each bound service-domain pair by concatenating the site prefix of the service and the name of the domain. If the special string NONE is specified, the site-prefix string is unset. When implementing HTTP redirect site persistence, the Citrix ADC redirects GSLB requests to GSLB services by using their site domains.",
			},
			"state": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Enable or disable the service.",
			},
			"svrtimeout": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Idle time, in seconds, after which a server connection is terminated. Applicable if connection proxy based site persistence is used.",
			},
			"viewip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address to be used for the given view",
			},
			"viewname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the DNS view of the service. A DNS view is used in global server load balancing (GSLB) to return a predetermined IP address to a specific group of clients, which are identified by using a DNS policy.",
			},
			"weight": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Weight to assign to the monitor-service binding. A larger number specifies a greater weight. Contributes to the monitoring threshold, which determines the state of the service.",
			},
		},
	}
}

func gslbserviceGetThePayloadFromtheConfig(ctx context.Context, data *GslbserviceResourceModel) gslb.Gslbservice {
	tflog.Debug(ctx, "In gslbserviceGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	gslbservice := gslb.Gslbservice{}
	if !data.Appflowlog.IsNull() {
		gslbservice.Appflowlog = data.Appflowlog.ValueString()
	}
	if !data.Cip.IsNull() {
		gslbservice.Cip = data.Cip.ValueString()
	}
	if !data.Cipheader.IsNull() {
		gslbservice.Cipheader = data.Cipheader.ValueString()
	}
	if !data.Clttimeout.IsNull() {
		gslbservice.Clttimeout = utils.IntPtr(int(data.Clttimeout.ValueInt64()))
	}
	if !data.Cnameentry.IsNull() {
		gslbservice.Cnameentry = data.Cnameentry.ValueString()
	}
	if !data.Comment.IsNull() {
		gslbservice.Comment = data.Comment.ValueString()
	}
	if !data.Cookietimeout.IsNull() {
		gslbservice.Cookietimeout = utils.IntPtr(int(data.Cookietimeout.ValueInt64()))
	}
	if !data.Downstateflush.IsNull() {
		gslbservice.Downstateflush = data.Downstateflush.ValueString()
	}
	if !data.Hashid.IsNull() {
		gslbservice.Hashid = utils.IntPtr(int(data.Hashid.ValueInt64()))
	}
	if !data.Healthmonitor.IsNull() {
		gslbservice.Healthmonitor = data.Healthmonitor.ValueString()
	}
	if !data.Ip.IsNull() {
		gslbservice.Ip = data.Ip.ValueString()
	}
	if !data.Ipaddress.IsNull() {
		gslbservice.Ipaddress = data.Ipaddress.ValueString()
	}
	if !data.Maxaaausers.IsNull() {
		gslbservice.Maxaaausers = utils.IntPtr(int(data.Maxaaausers.ValueInt64()))
	}
	if !data.Maxbandwidth.IsNull() {
		gslbservice.Maxbandwidth = utils.IntPtr(int(data.Maxbandwidth.ValueInt64()))
	}
	if !data.Maxclient.IsNull() {
		gslbservice.Maxclient = utils.IntPtr(int(data.Maxclient.ValueInt64()))
	}
	if !data.MonitorNameSvc.IsNull() {
		gslbservice.Monitornamesvc = data.MonitorNameSvc.ValueString()
	}
	if !data.Monthreshold.IsNull() {
		gslbservice.Monthreshold = utils.IntPtr(int(data.Monthreshold.ValueInt64()))
	}
	if !data.Naptrdomainttl.IsNull() {
		gslbservice.Naptrdomainttl = utils.IntPtr(int(data.Naptrdomainttl.ValueInt64()))
	}
	if !data.Naptrorder.IsNull() {
		gslbservice.Naptrorder = utils.IntPtr(int(data.Naptrorder.ValueInt64()))
	}
	if !data.Naptrpreference.IsNull() {
		gslbservice.Naptrpreference = utils.IntPtr(int(data.Naptrpreference.ValueInt64()))
	}
	if !data.Naptrreplacement.IsNull() {
		gslbservice.Naptrreplacement = data.Naptrreplacement.ValueString()
	}
	if !data.Naptrservices.IsNull() {
		gslbservice.Naptrservices = data.Naptrservices.ValueString()
	}
	if !data.Newname.IsNull() {
		gslbservice.Newname = data.Newname.ValueString()
	}
	if !data.Port.IsNull() {
		gslbservice.Port = utils.IntPtr(int(data.Port.ValueInt64()))
	}
	if !data.Publicip.IsNull() {
		gslbservice.Publicip = data.Publicip.ValueString()
	}
	if !data.Publicport.IsNull() {
		gslbservice.Publicport = utils.IntPtr(int(data.Publicport.ValueInt64()))
	}
	if !data.Servername.IsNull() {
		gslbservice.Servername = data.Servername.ValueString()
	}
	if !data.Servicename.IsNull() {
		gslbservice.Servicename = data.Servicename.ValueString()
	}
	if !data.Servicetype.IsNull() {
		gslbservice.Servicetype = data.Servicetype.ValueString()
	}
	if !data.Sitename.IsNull() {
		gslbservice.Sitename = data.Sitename.ValueString()
	}
	if !data.Sitepersistence.IsNull() {
		gslbservice.Sitepersistence = data.Sitepersistence.ValueString()
	}
	if !data.Siteprefix.IsNull() {
		gslbservice.Siteprefix = data.Siteprefix.ValueString()
	}
	if !data.State.IsNull() {
		gslbservice.State = data.State.ValueString()
	}
	if !data.Svrtimeout.IsNull() {
		gslbservice.Svrtimeout = utils.IntPtr(int(data.Svrtimeout.ValueInt64()))
	}
	if !data.Viewip.IsNull() {
		gslbservice.Viewip = data.Viewip.ValueString()
	}
	if !data.Viewname.IsNull() {
		gslbservice.Viewname = data.Viewname.ValueString()
	}
	if !data.Weight.IsNull() {
		gslbservice.Weight = utils.IntPtr(int(data.Weight.ValueInt64()))
	}

	return gslbservice
}

func gslbserviceSetAttrFromGet(ctx context.Context, data *GslbserviceResourceModel, getResponseData map[string]interface{}) *GslbserviceResourceModel {
	tflog.Debug(ctx, "In gslbserviceSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["appflowlog"]; ok && val != nil {
		data.Appflowlog = types.StringValue(val.(string))
	} else {
		data.Appflowlog = types.StringNull()
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
	if val, ok := getResponseData["cnameentry"]; ok && val != nil {
		data.Cnameentry = types.StringValue(val.(string))
	} else {
		data.Cnameentry = types.StringNull()
	}
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["cookietimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Cookietimeout = types.Int64Value(intVal)
		}
	} else {
		data.Cookietimeout = types.Int64Null()
	}
	if val, ok := getResponseData["downstateflush"]; ok && val != nil {
		data.Downstateflush = types.StringValue(val.(string))
	} else {
		data.Downstateflush = types.StringNull()
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
	if val, ok := getResponseData["ip"]; ok && val != nil {
		data.Ip = types.StringValue(val.(string))
	} else {
		data.Ip = types.StringNull()
	}
	if val, ok := getResponseData["ipaddress"]; ok && val != nil {
		data.Ipaddress = types.StringValue(val.(string))
	} else {
		data.Ipaddress = types.StringNull()
	}
	if val, ok := getResponseData["maxaaausers"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxaaausers = types.Int64Value(intVal)
		}
	} else {
		data.Maxaaausers = types.Int64Null()
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
	if val, ok := getResponseData["naptrdomainttl"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Naptrdomainttl = types.Int64Value(intVal)
		}
	} else {
		data.Naptrdomainttl = types.Int64Null()
	}
	if val, ok := getResponseData["naptrorder"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Naptrorder = types.Int64Value(intVal)
		}
	} else {
		data.Naptrorder = types.Int64Null()
	}
	if val, ok := getResponseData["naptrpreference"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Naptrpreference = types.Int64Value(intVal)
		}
	} else {
		data.Naptrpreference = types.Int64Null()
	}
	if val, ok := getResponseData["naptrreplacement"]; ok && val != nil {
		data.Naptrreplacement = types.StringValue(val.(string))
	} else {
		data.Naptrreplacement = types.StringNull()
	}
	if val, ok := getResponseData["naptrservices"]; ok && val != nil {
		data.Naptrservices = types.StringValue(val.(string))
	} else {
		data.Naptrservices = types.StringNull()
	}
	if val, ok := getResponseData["newname"]; ok && val != nil {
		data.Newname = types.StringValue(val.(string))
	} else {
		data.Newname = types.StringNull()
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
	if val, ok := getResponseData["servicename"]; ok && val != nil {
		data.Servicename = types.StringValue(val.(string))
	} else {
		data.Servicename = types.StringNull()
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
	if val, ok := getResponseData["viewip"]; ok && val != nil {
		data.Viewip = types.StringValue(val.(string))
	} else {
		data.Viewip = types.StringNull()
	}
	if val, ok := getResponseData["viewname"]; ok && val != nil {
		data.Viewname = types.StringValue(val.(string))
	} else {
		data.Viewname = types.StringNull()
	}
	if val, ok := getResponseData["weight"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Weight = types.Int64Value(intVal)
		}
	} else {
		data.Weight = types.Int64Null()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Servicename.ValueString())

	return data
}
