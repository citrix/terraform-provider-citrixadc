package gslbservice

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/gslb"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// LbmonitorbindingModel is one element of the lbmonitorbinding set.
type LbmonitorbindingModel struct {
	Weight      types.Int64  `tfsdk:"weight"`
	MonitorName types.String `tfsdk:"monitor_name"`
	Monstate    types.String `tfsdk:"monstate"`
}

var lbmonitorbindingAttrTypes = map[string]attr.Type{
	"weight":       types.Int64Type,
	"monitor_name": types.StringType,
	"monstate":     types.StringType,
}

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
	Delay            types.Int64  `tfsdk:"delay"`
	Downstateflush   types.String `tfsdk:"downstateflush"`
	Hashid           types.Int64  `tfsdk:"hashid"`
	Healthmonitor    types.String `tfsdk:"healthmonitor"`
	Ip               types.String `tfsdk:"ip"`
	Ipaddress        types.String `tfsdk:"ipaddress"`
	Maxaaausers      types.Int64  `tfsdk:"maxaaausers"`
	Maxbandwidth     types.Int64  `tfsdk:"maxbandwidth"`
	Maxclient        types.Int64  `tfsdk:"maxclient"`
	Monitornamesvc   types.String `tfsdk:"monitornamesvc"`
	Monthreshold     types.Int64  `tfsdk:"monthreshold"`
	Naptrdomainttl   types.Int64  `tfsdk:"naptrdomainttl"`
	Naptrorder       types.Int64  `tfsdk:"naptrorder"`
	Naptrpreference  types.Int64  `tfsdk:"naptrpreference"`
	Naptrreplacement types.String `tfsdk:"naptrreplacement"`
	Naptrservices    types.String `tfsdk:"naptrservices"`
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
	Lbmonitorbinding types.Set    `tfsdk:"lbmonitorbinding"`
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
				Computed:    true,
				Description: "Enable logging appflow flow information.",
			},
			"cip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Insert the client's IP address header in the request forwarded to the GSLB service.",
			},
			"cipheader": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name for the HTTP header that stores the client's IP address.",
			},
			"clttimeout": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Idle time, in seconds, after which a client connection is terminated.",
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
			"delay": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The time, in seconds, after which the GSLB service is disabled when disabling with -delay.",
			},
			"downstateflush": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Flush all active transactions associated with the GSLB service when its state transitions from UP to DOWN.",
			},
			"hashid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique hash identifier for the GSLB service, used by hash based load balancing methods.",
			},
			"healthmonitor": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Monitor the health of the GSLB service.",
			},
			"ip": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "IP address for the GSLB service.",
			},
			"ipaddress": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The new IP address of the service.",
			},
			"maxaaausers": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of SSL VPN users that can be logged on concurrently to the VPN virtual server represented by this GSLB service.",
			},
			"maxbandwidth": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum bandwidth, in Kbps, allowed for the service.",
			},
			"maxclient": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of open connections that the service can support at any given time.",
			},
			"monitornamesvc": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the monitor to bind to the service.",
			},
			"monthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Monitoring threshold value for the GSLB service.",
			},
			"naptrdomainttl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Modify the TTL of the internally created naptr domain.",
			},
			"naptrorder": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Order in which the NAPTR records MUST be processed.",
			},
			"naptrpreference": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Preference of this NAPTR among NAPTR records having same order.",
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
				Description: "The public IP address that a NAT device translates to the GSLB service's private IP address.",
			},
			"publicport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The public port associated with the GSLB service's public IP address.",
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
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the GSLB service. Cannot be changed after the GSLB service is created.",
			},
			"servicetype": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Type of service to create.",
			},
			"sitename": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
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
				Description: "The site's prefix string.",
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable the service.",
			},
			"svrtimeout": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Idle time, in seconds, after which a server connection is terminated.",
			},
			"viewip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address to be used for the given view.",
			},
			"viewname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the DNS view of the service.",
			},
			"weight": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Weight to assign to the monitor-service binding.",
			},
		},
		Blocks: map[string]schema.Block{
			"lbmonitorbinding": schema.SetNestedBlock{
				Description: "Monitors to bind to the GSLB service.",
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"weight": schema.Int64Attribute{
							Optional:    true,
							Computed:    true,
							Description: "Weight to assign to the monitor-service binding.",
						},
						"monitor_name": schema.StringAttribute{
							Optional:    true,
							Computed:    true,
							Description: "Name of the monitor bound to the GSLB service.",
						},
						"monstate": schema.StringAttribute{
							Optional:    true,
							Computed:    true,
							Description: "State of the monitor bound to the GSLB service.",
						},
					},
				},
			},
		},
	}
}

// gslbserviceGetThePayloadFromthePlan builds the full add payload (used on create).
func gslbserviceGetThePayloadFromthePlan(ctx context.Context, data *GslbserviceResourceModel) gslb.Gslbservice {
	tflog.Debug(ctx, "In gslbserviceGetThePayloadFromthePlan Function")

	gslbservice := gslb.Gslbservice{}
	if !data.Appflowlog.IsNull() && !data.Appflowlog.IsUnknown() {
		gslbservice.Appflowlog = data.Appflowlog.ValueString()
	}
	if !data.Cip.IsNull() && !data.Cip.IsUnknown() {
		gslbservice.Cip = data.Cip.ValueString()
	}
	if !data.Cipheader.IsNull() && !data.Cipheader.IsUnknown() {
		gslbservice.Cipheader = data.Cipheader.ValueString()
	}
	if !data.Clttimeout.IsNull() && !data.Clttimeout.IsUnknown() {
		gslbservice.Clttimeout = utils.IntPtr(int(data.Clttimeout.ValueInt64()))
	}
	if !data.Cnameentry.IsNull() && !data.Cnameentry.IsUnknown() {
		gslbservice.Cnameentry = data.Cnameentry.ValueString()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		gslbservice.Comment = data.Comment.ValueString()
	}
	if !data.Cookietimeout.IsNull() && !data.Cookietimeout.IsUnknown() {
		gslbservice.Cookietimeout = utils.IntPtr(int(data.Cookietimeout.ValueInt64()))
	}
	if !data.Downstateflush.IsNull() && !data.Downstateflush.IsUnknown() {
		gslbservice.Downstateflush = data.Downstateflush.ValueString()
	}
	if !data.Hashid.IsNull() && !data.Hashid.IsUnknown() {
		gslbservice.Hashid = utils.IntPtr(int(data.Hashid.ValueInt64()))
	}
	if !data.Healthmonitor.IsNull() && !data.Healthmonitor.IsUnknown() {
		gslbservice.Healthmonitor = data.Healthmonitor.ValueString()
	}
	if !data.Ip.IsNull() && !data.Ip.IsUnknown() {
		gslbservice.Ip = data.Ip.ValueString()
	}
	if !data.Ipaddress.IsNull() && !data.Ipaddress.IsUnknown() {
		gslbservice.Ipaddress = data.Ipaddress.ValueString()
	}
	if !data.Maxaaausers.IsNull() && !data.Maxaaausers.IsUnknown() {
		gslbservice.Maxaaausers = utils.IntPtr(int(data.Maxaaausers.ValueInt64()))
	}
	if !data.Maxbandwidth.IsNull() && !data.Maxbandwidth.IsUnknown() {
		gslbservice.Maxbandwidth = utils.IntPtr(int(data.Maxbandwidth.ValueInt64()))
	}
	if !data.Maxclient.IsNull() && !data.Maxclient.IsUnknown() {
		gslbservice.Maxclient = utils.IntPtr(int(data.Maxclient.ValueInt64()))
	}
	if !data.Monitornamesvc.IsNull() && !data.Monitornamesvc.IsUnknown() {
		gslbservice.Monitornamesvc = data.Monitornamesvc.ValueString()
	}
	if !data.Monthreshold.IsNull() && !data.Monthreshold.IsUnknown() {
		gslbservice.Monthreshold = utils.IntPtr(int(data.Monthreshold.ValueInt64()))
	}
	if !data.Naptrdomainttl.IsNull() && !data.Naptrdomainttl.IsUnknown() {
		gslbservice.Naptrdomainttl = utils.IntPtr(int(data.Naptrdomainttl.ValueInt64()))
	}
	if !data.Naptrorder.IsNull() && !data.Naptrorder.IsUnknown() {
		gslbservice.Naptrorder = utils.IntPtr(int(data.Naptrorder.ValueInt64()))
	}
	if !data.Naptrpreference.IsNull() && !data.Naptrpreference.IsUnknown() {
		gslbservice.Naptrpreference = utils.IntPtr(int(data.Naptrpreference.ValueInt64()))
	}
	if !data.Naptrreplacement.IsNull() && !data.Naptrreplacement.IsUnknown() {
		gslbservice.Naptrreplacement = data.Naptrreplacement.ValueString()
	}
	if !data.Naptrservices.IsNull() && !data.Naptrservices.IsUnknown() {
		gslbservice.Naptrservices = data.Naptrservices.ValueString()
	}
	if !data.Port.IsNull() && !data.Port.IsUnknown() {
		gslbservice.Port = utils.IntPtr(int(data.Port.ValueInt64()))
	}
	if !data.Publicip.IsNull() && !data.Publicip.IsUnknown() {
		gslbservice.Publicip = data.Publicip.ValueString()
	}
	if !data.Publicport.IsNull() && !data.Publicport.IsUnknown() {
		gslbservice.Publicport = utils.IntPtr(int(data.Publicport.ValueInt64()))
	}
	if !data.Servername.IsNull() && !data.Servername.IsUnknown() {
		gslbservice.Servername = data.Servername.ValueString()
	}
	if !data.Servicename.IsNull() && !data.Servicename.IsUnknown() {
		gslbservice.Servicename = data.Servicename.ValueString()
	}
	if !data.Servicetype.IsNull() && !data.Servicetype.IsUnknown() {
		gslbservice.Servicetype = data.Servicetype.ValueString()
	}
	if !data.Sitename.IsNull() && !data.Sitename.IsUnknown() {
		gslbservice.Sitename = data.Sitename.ValueString()
	}
	if !data.Sitepersistence.IsNull() && !data.Sitepersistence.IsUnknown() {
		gslbservice.Sitepersistence = data.Sitepersistence.ValueString()
	}
	if !data.Siteprefix.IsNull() && !data.Siteprefix.IsUnknown() {
		gslbservice.Siteprefix = data.Siteprefix.ValueString()
	}
	if !data.State.IsNull() && !data.State.IsUnknown() {
		gslbservice.State = data.State.ValueString()
	}
	if !data.Svrtimeout.IsNull() && !data.Svrtimeout.IsUnknown() {
		gslbservice.Svrtimeout = utils.IntPtr(int(data.Svrtimeout.ValueInt64()))
	}
	if !data.Viewip.IsNull() && !data.Viewip.IsUnknown() {
		gslbservice.Viewip = data.Viewip.ValueString()
	}
	if !data.Viewname.IsNull() && !data.Viewname.IsUnknown() {
		gslbservice.Viewname = data.Viewname.ValueString()
	}
	if !data.Weight.IsNull() && !data.Weight.IsUnknown() {
		gslbservice.Weight = utils.IntPtr(int(data.Weight.ValueInt64()))
	}

	return gslbservice
}

func gslbserviceSetAttrFromGet(ctx context.Context, data *GslbserviceResourceModel, getResponseData map[string]interface{}) *GslbserviceResourceModel {
	tflog.Debug(ctx, "In gslbserviceSetAttrFromGet Function")

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
	} else if data.Clttimeout.IsUnknown() {
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
	} else if data.Cookietimeout.IsUnknown() {
		data.Cookietimeout = types.Int64Null()
	}
	// delay is config-only (never returned by NITRO); preserve a known configured value,
	// resolve unknown to null.
	if data.Delay.IsUnknown() {
		data.Delay = types.Int64Null()
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
	} else if data.Hashid.IsUnknown() {
		data.Hashid = types.Int64Null()
	}
	if val, ok := getResponseData["healthmonitor"]; ok && val != nil {
		data.Healthmonitor = types.StringValue(val.(string))
	} else {
		data.Healthmonitor = types.StringNull()
	}
	// ip is not returned by NITRO; SDK v2 maps it from ipaddress.
	if val, ok := getResponseData["ipaddress"]; ok && val != nil {
		data.Ip = types.StringValue(val.(string))
		data.Ipaddress = types.StringValue(val.(string))
	} else {
		if data.Ip.IsUnknown() {
			data.Ip = types.StringNull()
		}
		data.Ipaddress = types.StringNull()
	}
	if val, ok := getResponseData["maxaaausers"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxaaausers = types.Int64Value(intVal)
		}
	} else if data.Maxaaausers.IsUnknown() {
		data.Maxaaausers = types.Int64Null()
	}
	if val, ok := getResponseData["maxbandwidth"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxbandwidth = types.Int64Value(intVal)
		}
	} else if data.Maxbandwidth.IsUnknown() {
		data.Maxbandwidth = types.Int64Null()
	}
	if val, ok := getResponseData["maxclient"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxclient = types.Int64Value(intVal)
		}
	} else if data.Maxclient.IsUnknown() {
		data.Maxclient = types.Int64Null()
	}
	if val, ok := getResponseData["monitornamesvc"]; ok && val != nil {
		data.Monitornamesvc = types.StringValue(val.(string))
	} else {
		data.Monitornamesvc = types.StringNull()
	}
	if val, ok := getResponseData["monthreshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Monthreshold = types.Int64Value(intVal)
		}
	} else if data.Monthreshold.IsUnknown() {
		data.Monthreshold = types.Int64Null()
	}
	if val, ok := getResponseData["naptrdomainttl"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Naptrdomainttl = types.Int64Value(intVal)
		}
	} else if data.Naptrdomainttl.IsUnknown() {
		data.Naptrdomainttl = types.Int64Null()
	}
	if val, ok := getResponseData["naptrorder"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Naptrorder = types.Int64Value(intVal)
		}
	} else if data.Naptrorder.IsUnknown() {
		data.Naptrorder = types.Int64Null()
	}
	if val, ok := getResponseData["naptrpreference"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Naptrpreference = types.Int64Value(intVal)
		}
	} else if data.Naptrpreference.IsUnknown() {
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
	if val, ok := getResponseData["port"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Port = types.Int64Value(intVal)
		}
	} else if data.Port.IsUnknown() {
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
	} else if data.Publicport.IsUnknown() {
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
	} else if data.Svrtimeout.IsUnknown() {
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
	} else if data.Weight.IsUnknown() {
		data.Weight = types.Int64Null()
	}

	// Set ID for the resource: single unique attribute (servicename)
	data.Id = types.StringValue(data.Servicename.ValueString())

	return data
}
