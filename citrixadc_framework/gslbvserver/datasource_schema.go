package gslbvserver

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// GslbvserverDataSourceModel is the data-source-specific model, decoupled from
// GslbvserverResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only NITRO attributes the resource deliberately omits
// (curstate, status, health, totalservices, ...). Every non-key attribute is
// Computed; the Framework's per-attribute model <-> schema reflection requires
// this model to have exactly the attributes the data-source schema declares,
// which is why it cannot reuse the resource model.
type GslbvserverDataSourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Appflowlog             types.String `tfsdk:"appflowlog"`
	Backupip               types.String `tfsdk:"backupip"`
	Backuplbmethod         types.String `tfsdk:"backuplbmethod"`
	Backupsessiontimeout   types.Int64  `tfsdk:"backupsessiontimeout"`
	Backupvserver          types.String `tfsdk:"backupvserver"`
	Comment                types.String `tfsdk:"comment"`
	Considereffectivestate types.String `tfsdk:"considereffectivestate"`
	CookieDomain           types.String `tfsdk:"cookie_domain"`
	Cookietimeout          types.Int64  `tfsdk:"cookietimeout"`
	Disableprimaryondown   types.String `tfsdk:"disableprimaryondown"`
	Dnsrecordtype          types.String `tfsdk:"dnsrecordtype"`
	Domainname             types.String `tfsdk:"domainname"`
	Dynamicweight          types.String `tfsdk:"dynamicweight"`
	Ecs                    types.String `tfsdk:"ecs"`
	Ecsaddrvalidation      types.String `tfsdk:"ecsaddrvalidation"`
	Edr                    types.String `tfsdk:"edr"`
	Iptype                 types.String `tfsdk:"iptype"`
	Lbmethod               types.String `tfsdk:"lbmethod"`
	Mir                    types.String `tfsdk:"mir"`
	Name                   types.String `tfsdk:"name"` // Required lookup key
	Netmask                types.String `tfsdk:"netmask"`
	Newname                types.String `tfsdk:"newname"`
	Order                  types.Int64  `tfsdk:"order"`
	Orderthreshold         types.Int64  `tfsdk:"orderthreshold"`
	Persistenceid          types.Int64  `tfsdk:"persistenceid"`
	Persistencetype        types.String `tfsdk:"persistencetype"`
	Persistmask            types.String `tfsdk:"persistmask"`
	Rule                   types.String `tfsdk:"rule"`
	Servicegroupname       types.String `tfsdk:"servicegroupname"`
	Servicename            types.String `tfsdk:"servicename"`
	Servicetype            types.String `tfsdk:"servicetype"`
	Sitedomainttl          types.Int64  `tfsdk:"sitedomainttl"`
	Sobackupaction         types.String `tfsdk:"sobackupaction"`
	Somethod               types.String `tfsdk:"somethod"`
	Sopersistence          types.String `tfsdk:"sopersistence"`
	Sopersistencetimeout   types.Int64  `tfsdk:"sopersistencetimeout"`
	Sothreshold            types.Int64  `tfsdk:"sothreshold"`
	State                  types.String `tfsdk:"state"`
	Timeout                types.Int64  `tfsdk:"timeout"`
	Toggleorder            types.String `tfsdk:"toggleorder"`
	Tolerance              types.Int64  `tfsdk:"tolerance"`
	Ttl                    types.Int64  `tfsdk:"ttl"`
	V6netmasklen           types.Int64  `tfsdk:"v6netmasklen"`
	V6persistmasklen       types.Int64  `tfsdk:"v6persistmasklen"`
	Weight                 types.Int64  `tfsdk:"weight"`
	// Convenience binding blocks preserved from the SDK v2 resource.
	Domain  types.Set `tfsdk:"domain"`
	Service types.Set `tfsdk:"service"`

	// Read-only (GET-only) NITRO attributes from the read-only set
	// (zion73x_readonly/gslbvserver.json). Never settable; populated from GET.
	Curstate                  types.String `tfsdk:"curstate"`
	Status                    types.Int64  `tfsdk:"status"`
	Lbrrreason                types.Int64  `tfsdk:"lbrrreason"`
	Iscname                   types.String `tfsdk:"iscname"`
	Sitepersistence           types.String `tfsdk:"sitepersistence"`
	Totalservices             types.Int64  `tfsdk:"totalservices"`
	Activeservices            types.Int64  `tfsdk:"activeservices"`
	Statechangetimesec        types.String `tfsdk:"statechangetimesec"`
	Statechangetimemsec       types.Int64  `tfsdk:"statechangetimemsec"`
	Tickssincelaststatechange types.Int64  `tfsdk:"tickssincelaststatechange"`
	Health                    types.Int64  `tfsdk:"health"`
	Policyname                types.String `tfsdk:"policyname"`
	Priority                  types.Int64  `tfsdk:"priority"`
	Gotopriorityexpression    types.String `tfsdk:"gotopriorityexpression"`
	Type                      types.String `tfsdk:"type"`
	Vsvrbindsvcip             types.String `tfsdk:"vsvrbindsvcip"`
	Vsvrbindsvcport           types.Int64  `tfsdk:"vsvrbindsvcport"`
	Servername                types.String `tfsdk:"servername"`
	Nodefaultbindings         types.String `tfsdk:"nodefaultbindings"`
	Currentactiveorder        types.String `tfsdk:"currentactiveorder"`
}

func GslbvserverDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"appflowlog": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable logging appflow flow information",
			},
			"backupip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The IP address of the backup service for the specified domain name. Used when all the services bound to the domain are down, or when the backup chain of virtual servers is down.",
			},
			"backuplbmethod": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Backup load balancing method. Becomes operational if the primary load balancing method fails or cannot be used. Valid only if the primary method is based on either round-trip time (RTT) or static proximity.",
			},
			"backupsessiontimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "A non zero value enables the feature whose minimum value is 2 minutes. The feature can be disabled by setting the value to zero. The created session is in effect for a specific client per domain.",
			},
			"backupvserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the backup GSLB virtual server to which the appliance should to forward requests if the status of the primary GSLB virtual server is down or exceeds its spillover threshold.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments that you might want to associate with the GSLB virtual server.",
			},
			"considereffectivestate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If the primary state of all bound GSLB services is DOWN, consider the effective states of all the GSLB services, obtained through the Metrics Exchange Protocol (MEP), when determining the state of the GSLB virtual server. To consider the effective state, set the parameter to STATE_ONLY. To disregard the effective state, set the parameter to NONE.\n\nThe effective state of a GSLB service is the ability of the corresponding virtual server to serve traffic. The effective state of the load balancing virtual server, which is transferred to the GSLB service, is UP even if only one virtual server in the backup chain of virtual servers is in the UP state.",
			},
			"cookie_domain": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The cookie domain for the GSLB site. Used when inserting the GSLB site cookie in the HTTP response.",
			},
			"cookietimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Timeout, in minutes, for the GSLB site cookie.",
			},
			"disableprimaryondown": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Continue to direct traffic to the backup chain even after the primary GSLB virtual server returns to the UP state. Used when spillover is configured for the virtual server.",
			},
			"dnsrecordtype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "DNS record type to associate with the GSLB virtual server's domain name.",
			},
			"domainname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Domain name for which to change the time to live (TTL) and/or backup service IP address.",
			},
			"dynamicweight": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify if the appliance should consider the service count, service weights, or ignore both when using weight-based load balancing methods. The state of the number of services bound to the virtual server help the appliance to select the service.",
			},
			"ecs": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If enabled, respond with EDNS Client Subnet (ECS) option in the response for a DNS query with ECS. The ECS address will be used for persistence and spillover persistence (if enabled) instead of the LDNS address. Persistence mask is ignored if ECS is enabled.",
			},
			"ecsaddrvalidation": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Validate if ECS address is a private or unroutable address and in such cases, use the LDNS IP.",
			},
			"edr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Send clients an empty DNS response when the GSLB virtual server is DOWN.",
			},
			"iptype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The IP type for this GSLB vserver.",
			},
			"lbmethod": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Load balancing method for the GSLB virtual server.",
			},
			"mir": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Include multiple IP addresses in the DNS responses sent to clients.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the GSLB virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Can be changed after the virtual server is created.\n\nCLI Users:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my vserver\" or 'my vserver').",
			},
			"netmask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IPv4 network mask for use in the SOURCEIPHASH load balancing method.",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New name for the GSLB virtual server.",
			},
			"order": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Order number to be assigned to the service when it is bound to the lb vserver.",
			},
			"orderthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "This option is used to to specify the threshold of minimum number of services to be UP in an order, for it to be considered in Lb decision.",
			},
			"persistenceid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The persistence ID for the GSLB virtual server. The ID is a positive integer that enables GSLB sites to identify the GSLB virtual server, and is required if source IP address based or spill over based persistence is enabled on the virtual server.",
			},
			"persistencetype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use source IP address based persistence for the virtual server.\nAfter the load balancing method selects a service for the first packet, the IP address received in response to the DNS query is used for subsequent requests from the same client.",
			},
			"persistmask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The optional IPv4 network mask applied to IPv4 addresses to establish source IP address based persistence.",
			},
			"rule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression, or name of a named expression, against which traffic is evaluated.\nThis field is applicable only if gslb method or gslb backup method are set to API.\nThe following requirements apply only to the Citrix ADC CLI:\n* If the expression includes one or more spaces, enclose the entire expression in double quotation marks.\n* If the expression itself includes double quotation marks, escape the quotations by using the \\ character.\n* Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.",
			},
			"servicegroupname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The GSLB service group name bound to the selected GSLB virtual server.",
			},
			"servicename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the GSLB service for which to change the weight.",
			},
			"servicetype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Protocol used by services bound to the virtual server.",
			},
			"sitedomainttl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "TTL, in seconds, for all internally created site domains (created when a site prefix is configured on a GSLB service) that are associated with this virtual server.",
			},
			"sobackupaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Action to be performed if spillover is to take effect, but no backup chain to spillover is usable or exists",
			},
			"somethod": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of threshold that, when exceeded, triggers spillover. Available settings function as follows:\n* CONNECTION - Spillover occurs when the number of client connections exceeds the threshold.\n* DYNAMICCONNECTION - Spillover occurs when the number of client connections at the GSLB virtual server exceeds the sum of the maximum client (Max Clients) settings for bound GSLB services. Do not specify a spillover threshold for this setting, because the threshold is implied by the Max Clients settings of the bound GSLB services.\n* BANDWIDTH - Spillover occurs when the bandwidth consumed by the GSLB virtual server's incoming and outgoing traffic exceeds the threshold.\n* HEALTH - Spillover occurs when the percentage of weights of the GSLB services that are UP drops below the threshold. For example, if services gslbSvc1, gslbSvc2, and gslbSvc3 are bound to a virtual server, with weights 1, 2, and 3, and the spillover threshold is 50%, spillover occurs if gslbSvc1 and gslbSvc3 or gslbSvc2 and gslbSvc3 transition to DOWN.\n* NONE - Spillover does not occur.",
			},
			"sopersistence": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If spillover occurs, maintain source IP address based persistence for both primary and backup GSLB virtual servers.",
			},
			"sopersistencetimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Timeout for spillover persistence, in minutes.",
			},
			"sothreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Threshold at which spillover occurs. Specify an integer for the CONNECTION spillover method, a bandwidth value in kilobits per second for the BANDWIDTH method (do not enter the units), or a percentage for the HEALTH method (do not enter the percentage symbol).",
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of the GSLB virtual server.",
			},
			"timeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Idle time, in minutes, after which a persistence entry is cleared.",
			},
			"toggleorder": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Configure this option to toggle order preference",
			},
			"tolerance": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Tolerance in milliseconds. Tolerance value is used in deciding which sites in a GSLB configuration must be considered for implementing the RTT load balancing method. The sites having the RTT value less than or equal to the sum of the lowest RTT and tolerance value are considered. NetScaler implements the round robin method of global server load balancing among these considered sites. The sites that have RTT value greater than this value are not considered. The logic is applied for each LDNS and based on the LDNS, the sites that are considered might change. For example, a site that is considered for requests coming from LDNS1 might not be considered for requests coming from LDNS2.",
			},
			"ttl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time to live (TTL) for the domain.",
			},
			"v6netmasklen": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of bits to consider, in an IPv6 source IP address, for creating the hash that is required by the SOURCEIPHASH load balancing method.",
			},
			"v6persistmasklen": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of bits to consider in an IPv6 source IP address when creating source IP address based persistence sessions.",
			},
			"weight": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Weight for the service.",
			},

			// Read-only (GET-only) NITRO attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"curstate": schema.StringAttribute{
				Computed:    true,
				Description: "State of the gslb vserver (for example UP, DOWN, OUT OF SERVICE, DISABLED).",
			},
			"status": schema.Int64Attribute{
				Computed:    true,
				Description: "Current status of the gslb vserver. During the initial phase, if the configured lb method is not round robin, the vserver will adopt round robin to distribute traffic for a predefined number of requests.",
			},
			"lbrrreason": schema.Int64Attribute{
				Computed:    true,
				Description: "Reason why a vserver is in RR (round robin).",
			},
			"iscname": schema.StringAttribute{
				Computed:    true,
				Description: "Is cname feature set on vserver (ENABLED, DISABLED).",
			},
			"sitepersistence": schema.StringAttribute{
				Computed:    true,
				Description: "Type of Site Persistence set (ConnectionProxy, HTTPRedirect, NONE).",
			},
			"totalservices": schema.Int64Attribute{
				Computed:    true,
				Description: "Total number of services bound to the vserver.",
			},
			"activeservices": schema.Int64Attribute{
				Computed:    true,
				Description: "Total number of active services bound to the vserver.",
			},
			"statechangetimesec": schema.StringAttribute{
				Computed:    true,
				Description: "Time when last state change happened. Seconds part.",
			},
			"statechangetimemsec": schema.Int64Attribute{
				Computed:    true,
				Description: "Time at which last state change happened. Milliseconds part.",
			},
			"tickssincelaststatechange": schema.Int64Attribute{
				Computed:    true,
				Description: "Time in 10 millisecond ticks since the last state change.",
			},
			"health": schema.Int64Attribute{
				Computed:    true,
				Description: "Health of vserver based on percentage of weights of active svcs/all svcs. This does not consider administratively disabled svcs.",
			},
			"policyname": schema.StringAttribute{
				Computed:    true,
				Description: "Name of the policy bound to the GSLB vserver.",
			},
			"priority": schema.Int64Attribute{
				Computed:    true,
				Description: "Priority.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Computed:    true,
				Description: "Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"type": schema.StringAttribute{
				Computed:    true,
				Description: "The bindpoint to which the policy is bound (REQUEST, RESPONSE, MQTT_JUMBO_REQ, HTTP_EVENT_RESPONSE).",
			},
			"vsvrbindsvcip": schema.StringAttribute{
				Computed:    true,
				Description: "Used for showing the ip of bound entities.",
			},
			"vsvrbindsvcport": schema.Int64Attribute{
				Computed:    true,
				Description: "Used for showing ports of bound entities.",
			},
			"servername": schema.StringAttribute{
				Computed:    true,
				Description: "Used to display server name in case of GSLB servicegroup binding to GSLB vserver.",
			},
			"nodefaultbindings": schema.StringAttribute{
				Computed:    true,
				Description: "To determine if the configuration will have default ssl CIPHER and ECC curve bindings (YES, NO).",
			},
			"currentactiveorder": schema.StringAttribute{
				Computed:    true,
				Description: "Current order that takes the traffic in case service or servicegroup is bound with order.",
			},
		},
		Blocks: map[string]schema.Block{
			"domain": schema.SetNestedBlock{
				Description: "Domains bound to the GSLB virtual server.",
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"domainname": schema.StringAttribute{
							Optional:    true,
							Computed:    true,
							Description: "Domain name for which to change the time to live (TTL) and/or backup service IP address.",
						},
						"ttl": schema.Int64Attribute{
							Optional:    true,
							Computed:    true,
							Description: "Time to live (TTL) for the domain.",
						},
						"backupip": schema.StringAttribute{
							Optional:    true,
							Computed:    true,
							Description: "The IP address of the backup service for the specified domain name.",
						},
						"cookiedomain": schema.StringAttribute{
							Optional:    true,
							Computed:    true,
							Description: "The cookie domain for the GSLB site.",
						},
						"cookietimeout": schema.Int64Attribute{
							Optional:    true,
							Computed:    true,
							Description: "Timeout, in minutes, for the GSLB site cookie.",
						},
						"sitedomainttl": schema.Int64Attribute{
							Optional:    true,
							Computed:    true,
							Description: "TTL, in seconds, for all internally created site domains associated with this virtual server.",
						},
						"backupipflag": schema.BoolAttribute{
							Optional:    true,
							Computed:    true,
							Description: "The IP address of the backup service flag.",
						},
						"cookiedomainflag": schema.BoolAttribute{
							Optional:    true,
							Computed:    true,
							Description: "The cookie domain flag.",
						},
						"name": schema.StringAttribute{
							Optional:    true,
							Computed:    true,
							Description: "Name of the virtual server on which to perform the binding operation.",
						},
					},
				},
			},
			"service": schema.SetNestedBlock{
				Description: "GSLB services bound to the GSLB virtual server.",
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"servicename": schema.StringAttribute{
							Optional:    true,
							Computed:    true,
							Description: "Name of the GSLB service bound to the GSLB virtual server.",
						},
						"domainname": schema.StringAttribute{
							Optional:    true,
							Computed:    true,
							Description: "Domain name for which to change the time to live (TTL) and/or backup service IP address.",
						},
						"weight": schema.Int64Attribute{
							Optional:    true,
							Computed:    true,
							Description: "Weight to assign to the GSLB service.",
						},
					},
				},
			},
		},
	}
}

// gslbvserverDataSourceSetAttrFromGet projects a NITRO gslbvserver GET response
// onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them) — no unknown->null resolution or plan preservation is
// required. The shared utils.MapGet* helpers implement that projection.
func gslbvserverDataSourceSetAttrFromGet(ctx context.Context, data *GslbvserverDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In gslbvserverDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Appflowlog = utils.MapGetString(g, "appflowlog")
	data.Backupip = utils.MapGetString(g, "backupip")
	data.Backuplbmethod = utils.MapGetString(g, "backuplbmethod")
	data.Backupsessiontimeout = utils.MapGetInt64(g, "backupsessiontimeout")
	data.Backupvserver = utils.MapGetString(g, "backupvserver")
	data.Comment = utils.MapGetString(g, "comment")
	data.Considereffectivestate = utils.MapGetString(g, "considereffectivestate")
	data.CookieDomain = utils.MapGetString(g, "cookie_domain")
	data.Cookietimeout = utils.MapGetInt64(g, "cookietimeout")
	data.Disableprimaryondown = utils.MapGetString(g, "disableprimaryondown")
	data.Dnsrecordtype = utils.MapGetString(g, "dnsrecordtype")
	data.Domainname = utils.MapGetString(g, "domainname")
	data.Dynamicweight = utils.MapGetString(g, "dynamicweight")
	data.Ecs = utils.MapGetString(g, "ecs")
	data.Ecsaddrvalidation = utils.MapGetString(g, "ecsaddrvalidation")
	data.Edr = utils.MapGetString(g, "edr")
	data.Iptype = utils.MapGetString(g, "iptype")
	data.Lbmethod = utils.MapGetString(g, "lbmethod")
	data.Mir = utils.MapGetString(g, "mir")
	data.Netmask = utils.MapGetString(g, "netmask")
	data.Newname = utils.MapGetString(g, "newname")
	data.Order = utils.MapGetInt64(g, "order")
	data.Orderthreshold = utils.MapGetInt64(g, "orderthreshold")
	data.Persistenceid = utils.MapGetInt64(g, "persistenceid")
	data.Persistencetype = utils.MapGetString(g, "persistencetype")
	data.Persistmask = utils.MapGetString(g, "persistmask")
	data.Rule = utils.MapGetString(g, "rule")
	data.Servicegroupname = utils.MapGetString(g, "servicegroupname")
	data.Servicename = utils.MapGetString(g, "servicename")
	data.Servicetype = utils.MapGetString(g, "servicetype")
	data.Sitedomainttl = utils.MapGetInt64(g, "sitedomainttl")
	data.Sobackupaction = utils.MapGetString(g, "sobackupaction")
	data.Somethod = utils.MapGetString(g, "somethod")
	data.Sopersistence = utils.MapGetString(g, "sopersistence")
	data.Sopersistencetimeout = utils.MapGetInt64(g, "sopersistencetimeout")
	data.Sothreshold = utils.MapGetInt64(g, "sothreshold")
	data.State = utils.MapGetString(g, "state")
	data.Timeout = utils.MapGetInt64(g, "timeout")
	data.Toggleorder = utils.MapGetString(g, "toggleorder")
	data.Tolerance = utils.MapGetInt64(g, "tolerance")
	data.Ttl = utils.MapGetInt64(g, "ttl")
	data.V6netmasklen = utils.MapGetInt64(g, "v6netmasklen")
	data.V6persistmasklen = utils.MapGetInt64(g, "v6persistmasklen")
	data.Weight = utils.MapGetInt64(g, "weight")

	// The domain/service convenience blocks are populated from separate binding
	// GETs, not from the bare vserver GET, so leave them Null here.
	data.Domain = types.SetNull(types.ObjectType{AttrTypes: domainbindingAttrTypes})
	data.Service = types.SetNull(types.ObjectType{AttrTypes: servicebindingAttrTypes})

	// Read-only NITRO attributes.
	data.Curstate = utils.MapGetString(g, "curstate")
	data.Status = utils.MapGetInt64(g, "status")
	data.Lbrrreason = utils.MapGetInt64(g, "lbrrreason")
	data.Iscname = utils.MapGetString(g, "iscname")
	data.Sitepersistence = utils.MapGetString(g, "sitepersistence")
	data.Totalservices = utils.MapGetInt64(g, "totalservices")
	data.Activeservices = utils.MapGetInt64(g, "activeservices")
	data.Statechangetimesec = utils.MapGetString(g, "statechangetimesec")
	data.Statechangetimemsec = utils.MapGetInt64(g, "statechangetimemsec")
	data.Tickssincelaststatechange = utils.MapGetInt64(g, "tickssincelaststatechange")
	data.Health = utils.MapGetInt64(g, "health")
	data.Policyname = utils.MapGetString(g, "policyname")
	data.Priority = utils.MapGetInt64(g, "priority")
	data.Gotopriorityexpression = utils.MapGetString(g, "gotopriorityexpression")
	data.Type = utils.MapGetString(g, "type")
	data.Vsvrbindsvcip = utils.MapGetString(g, "vsvrbindsvcip")
	data.Vsvrbindsvcport = utils.MapGetInt64(g, "vsvrbindsvcport")
	data.Servername = utils.MapGetString(g, "servername")
	data.Nodefaultbindings = utils.MapGetString(g, "nodefaultbindings")
	data.Currentactiveorder = utils.MapGetString(g, "currentactiveorder")
}
