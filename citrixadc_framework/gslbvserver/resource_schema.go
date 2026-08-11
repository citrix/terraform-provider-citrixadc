package gslbvserver

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/gslb"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// DomainbindingModel is one element of the "domain" convenience set. It mirrors the
// SDK v2 "domain" block, which manages gslbvserver_domain_binding entries.
type DomainbindingModel struct {
	Backupip         types.String `tfsdk:"backupip"`
	Backupipflag     types.Bool   `tfsdk:"backupipflag"`
	Cookiedomain     types.String `tfsdk:"cookiedomain"`
	Cookiedomainflag types.Bool   `tfsdk:"cookiedomainflag"`
	Cookietimeout    types.Int64  `tfsdk:"cookietimeout"`
	Domainname       types.String `tfsdk:"domainname"`
	Name             types.String `tfsdk:"name"`
	Sitedomainttl    types.Int64  `tfsdk:"sitedomainttl"`
	Ttl              types.Int64  `tfsdk:"ttl"`
}

var domainbindingAttrTypes = map[string]attr.Type{
	"backupip":         types.StringType,
	"backupipflag":     types.BoolType,
	"cookiedomain":     types.StringType,
	"cookiedomainflag": types.BoolType,
	"cookietimeout":    types.Int64Type,
	"domainname":       types.StringType,
	"name":             types.StringType,
	"sitedomainttl":    types.Int64Type,
	"ttl":              types.Int64Type,
}

// ServicebindingModel is one element of the "service" convenience set. It mirrors the
// SDK v2 "service" block, which manages gslbvserver_gslbservice_binding entries.
type ServicebindingModel struct {
	Domainname  types.String `tfsdk:"domainname"`
	Servicename types.String `tfsdk:"servicename"`
	Weight      types.Int64  `tfsdk:"weight"`
}

var servicebindingAttrTypes = map[string]attr.Type{
	"domainname":  types.StringType,
	"servicename": types.StringType,
	"weight":      types.Int64Type,
}

// GslbvserverResourceModel describes the resource data model.
type GslbvserverResourceModel struct {
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
	Name                   types.String `tfsdk:"name"`
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
}

func (r *GslbvserverResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the gslbvserver resource.",
			},
			"appflowlog": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("ENABLED"),
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
				Default:     stringdefault.StaticString("NONE"),
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
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Continue to direct traffic to the backup chain even after the primary GSLB virtual server returns to the UP state. Used when spillover is configured for the virtual server.",
			},
			"dnsrecordtype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("A"),
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
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "If enabled, respond with EDNS Client Subnet (ECS) option in the response for a DNS query with ECS. The ECS address will be used for persistence and spillover persistence (if enabled) instead of the LDNS address. Persistence mask is ignored if ECS is enabled.",
			},
			"ecsaddrvalidation": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Validate if ECS address is a private or unroutable address and in such cases, use the LDNS IP.",
			},
			"edr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("DISABLED"),
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
				Default:     stringdefault.StaticString("LEASTCONNECTION"),
				Description: "Load balancing method for the GSLB virtual server.",
			},
			"mir": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Include multiple IP addresses in the DNS responses sent to clients.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the GSLB virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Can be changed after the virtual server is created.\n\nCLI Users:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my vserver\" or 'my vserver').",
			},
			"netmask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IPv4 network mask for use in the SOURCEIPHASH load balancing method.",
			},
			"newname": schema.StringAttribute{
				Optional: true,
				// newname is the rename trigger (NITRO ?action=rename). Changing it
				// must NOT force replacement - it drives an in-place rename via Update.
				// Not Computed: it is a pure user input, never echoed back by GET.
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
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
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
				Default:     stringdefault.StaticString("ASCENDING"),
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
		},
		Blocks: map[string]schema.Block{
			// "domain" convenience block preserved from SDK v2 - manages
			// gslbvserver_domain_binding entries bound to this vserver.
			"domain": schema.SetNestedBlock{
				Description: "Domains to bind to the GSLB virtual server.",
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
			// "service" convenience block preserved from SDK v2 - manages
			// gslbvserver_gslbservice_binding entries bound to this vserver.
			"service": schema.SetNestedBlock{
				Description: "GSLB services to bind to the GSLB virtual server.",
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"servicename": schema.StringAttribute{
							Required:    true,
							Description: "Name of the GSLB service to bind to the GSLB virtual server.",
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

// gslbvserverGetThePayloadFromthePlan builds the create (add) payload. The NITRO
// `add` verb does not accept `backupvserver` (it is applied by a follow-up update
// in Create, mirroring the SDK v2 behavior) and `newname` is rename-only, so both
// are intentionally excluded here.
func gslbvserverGetThePayloadFromthePlan(ctx context.Context, data *GslbvserverResourceModel) gslb.Gslbvserver {
	tflog.Debug(ctx, "In gslbvserverGetThePayloadFromthePlan Function")

	// Create API request body from the model
	gslbvserver := gslb.Gslbvserver{}
	if !data.Appflowlog.IsNull() && !data.Appflowlog.IsUnknown() {
		gslbvserver.Appflowlog = data.Appflowlog.ValueString()
	}
	if !data.Backupip.IsNull() && !data.Backupip.IsUnknown() {
		gslbvserver.Backupip = data.Backupip.ValueString()
	}
	if !data.Backuplbmethod.IsNull() && !data.Backuplbmethod.IsUnknown() {
		gslbvserver.Backuplbmethod = data.Backuplbmethod.ValueString()
	}
	if !data.Backupsessiontimeout.IsNull() && !data.Backupsessiontimeout.IsUnknown() {
		gslbvserver.Backupsessiontimeout = utils.IntPtr(int(data.Backupsessiontimeout.ValueInt64()))
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		gslbvserver.Comment = data.Comment.ValueString()
	}
	if !data.Considereffectivestate.IsNull() && !data.Considereffectivestate.IsUnknown() {
		gslbvserver.Considereffectivestate = data.Considereffectivestate.ValueString()
	}
	if !data.CookieDomain.IsNull() && !data.CookieDomain.IsUnknown() {
		gslbvserver.Cookiedomain = data.CookieDomain.ValueString()
	}
	if !data.Cookietimeout.IsNull() && !data.Cookietimeout.IsUnknown() {
		gslbvserver.Cookietimeout = utils.IntPtr(int(data.Cookietimeout.ValueInt64()))
	}
	if !data.Disableprimaryondown.IsNull() && !data.Disableprimaryondown.IsUnknown() {
		gslbvserver.Disableprimaryondown = data.Disableprimaryondown.ValueString()
	}
	if !data.Dnsrecordtype.IsNull() && !data.Dnsrecordtype.IsUnknown() {
		gslbvserver.Dnsrecordtype = data.Dnsrecordtype.ValueString()
	}
	if !data.Domainname.IsNull() && !data.Domainname.IsUnknown() {
		gslbvserver.Domainname = data.Domainname.ValueString()
	}
	if !data.Dynamicweight.IsNull() && !data.Dynamicweight.IsUnknown() {
		gslbvserver.Dynamicweight = data.Dynamicweight.ValueString()
	}
	if !data.Ecs.IsNull() && !data.Ecs.IsUnknown() {
		gslbvserver.Ecs = data.Ecs.ValueString()
	}
	if !data.Ecsaddrvalidation.IsNull() && !data.Ecsaddrvalidation.IsUnknown() {
		gslbvserver.Ecsaddrvalidation = data.Ecsaddrvalidation.ValueString()
	}
	if !data.Edr.IsNull() && !data.Edr.IsUnknown() {
		gslbvserver.Edr = data.Edr.ValueString()
	}
	if !data.Iptype.IsNull() && !data.Iptype.IsUnknown() {
		gslbvserver.Iptype = data.Iptype.ValueString()
	}
	if !data.Lbmethod.IsNull() && !data.Lbmethod.IsUnknown() {
		gslbvserver.Lbmethod = data.Lbmethod.ValueString()
	}
	if !data.Mir.IsNull() && !data.Mir.IsUnknown() {
		gslbvserver.Mir = data.Mir.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		gslbvserver.Name = data.Name.ValueString()
	}
	if !data.Netmask.IsNull() && !data.Netmask.IsUnknown() {
		gslbvserver.Netmask = data.Netmask.ValueString()
	}
	// newname is a rename-only argument (NITRO ?action=rename); excluded from add.
	if !data.Order.IsNull() && !data.Order.IsUnknown() {
		gslbvserver.Order = utils.IntPtr(int(data.Order.ValueInt64()))
	}
	if !data.Orderthreshold.IsNull() && !data.Orderthreshold.IsUnknown() {
		gslbvserver.Orderthreshold = utils.IntPtr(int(data.Orderthreshold.ValueInt64()))
	}
	if !data.Persistenceid.IsNull() && !data.Persistenceid.IsUnknown() {
		gslbvserver.Persistenceid = utils.IntPtr(int(data.Persistenceid.ValueInt64()))
	}
	if !data.Persistencetype.IsNull() && !data.Persistencetype.IsUnknown() {
		gslbvserver.Persistencetype = data.Persistencetype.ValueString()
	}
	if !data.Persistmask.IsNull() && !data.Persistmask.IsUnknown() {
		gslbvserver.Persistmask = data.Persistmask.ValueString()
	}
	if !data.Rule.IsNull() && !data.Rule.IsUnknown() {
		gslbvserver.Rule = data.Rule.ValueString()
	}
	if !data.Servicegroupname.IsNull() && !data.Servicegroupname.IsUnknown() {
		gslbvserver.Servicegroupname = data.Servicegroupname.ValueString()
	}
	if !data.Servicename.IsNull() && !data.Servicename.IsUnknown() {
		gslbvserver.Servicename = data.Servicename.ValueString()
	}
	if !data.Servicetype.IsNull() && !data.Servicetype.IsUnknown() {
		gslbvserver.Servicetype = data.Servicetype.ValueString()
	}
	if !data.Sitedomainttl.IsNull() && !data.Sitedomainttl.IsUnknown() {
		gslbvserver.Sitedomainttl = utils.IntPtr(int(data.Sitedomainttl.ValueInt64()))
	}
	if !data.Sobackupaction.IsNull() && !data.Sobackupaction.IsUnknown() {
		gslbvserver.Sobackupaction = data.Sobackupaction.ValueString()
	}
	if !data.Somethod.IsNull() && !data.Somethod.IsUnknown() {
		gslbvserver.Somethod = data.Somethod.ValueString()
	}
	if !data.Sopersistence.IsNull() && !data.Sopersistence.IsUnknown() {
		gslbvserver.Sopersistence = data.Sopersistence.ValueString()
	}
	if !data.Sopersistencetimeout.IsNull() && !data.Sopersistencetimeout.IsUnknown() {
		gslbvserver.Sopersistencetimeout = utils.IntPtr(int(data.Sopersistencetimeout.ValueInt64()))
	}
	if !data.Sothreshold.IsNull() && !data.Sothreshold.IsUnknown() {
		gslbvserver.Sothreshold = utils.IntPtr(int(data.Sothreshold.ValueInt64()))
	}
	if !data.State.IsNull() && !data.State.IsUnknown() {
		gslbvserver.State = data.State.ValueString()
	}
	if !data.Timeout.IsNull() && !data.Timeout.IsUnknown() {
		gslbvserver.Timeout = utils.IntPtr(int(data.Timeout.ValueInt64()))
	}
	if !data.Toggleorder.IsNull() && !data.Toggleorder.IsUnknown() {
		gslbvserver.Toggleorder = data.Toggleorder.ValueString()
	}
	if !data.Tolerance.IsNull() && !data.Tolerance.IsUnknown() {
		gslbvserver.Tolerance = utils.IntPtr(int(data.Tolerance.ValueInt64()))
	}
	if !data.Ttl.IsNull() && !data.Ttl.IsUnknown() {
		gslbvserver.Ttl = utils.IntPtr(int(data.Ttl.ValueInt64()))
	}
	if !data.V6netmasklen.IsNull() && !data.V6netmasklen.IsUnknown() {
		gslbvserver.V6netmasklen = utils.IntPtr(int(data.V6netmasklen.ValueInt64()))
	}
	if !data.V6persistmasklen.IsNull() && !data.V6persistmasklen.IsUnknown() {
		gslbvserver.V6persistmasklen = utils.IntPtr(int(data.V6persistmasklen.ValueInt64()))
	}
	if !data.Weight.IsNull() && !data.Weight.IsUnknown() {
		gslbvserver.Weight = utils.IntPtr(int(data.Weight.ValueInt64()))
	}

	return gslbvserver
}

// gslbvserverGetTheUpdatablePayloadFromThePlan builds the update (PUT) payload with
// only the fields that changed between prior state and plan. It excludes the key
// (name), the ForceNew attribute (servicetype), the enable/disable-driven attribute
// (state) and the rename-only attribute (newname). Returns the payload and whether
// any updatable field changed. The caller sets Name to the current live name.
func gslbvserverGetTheUpdatablePayloadFromThePlan(ctx context.Context, data *GslbvserverResourceModel, state *GslbvserverResourceModel) (gslb.Gslbvserver, bool) {
	tflog.Debug(ctx, "In gslbvserverGetTheUpdatablePayloadFromThePlan Function")

	gslbvserver := gslb.Gslbvserver{}
	hasChange := false

	// changed reports whether a planned attribute genuinely changed from prior state.
	// It intentionally ignores Unknown and Null plan values: Optional+Computed
	// attributes that the user did not configure are planned as Unknown on update, and
	// treating that as a change would emit spurious zero-valued fields (e.g. an int
	// IntPtr(0)) that NITRO rejects with prerequisite errors such as errorcode 1093
	// "Argument pre-requisite missing [cookieTimeout, domainName]". This mirrors the
	// SDK v2 d.HasChange() semantics of only sending real, user-driven changes.
	changed := func(plan, prior attr.Value) bool {
		return !plan.Equal(prior) && !plan.IsUnknown() && !plan.IsNull()
	}

	if changed(data.Appflowlog, state.Appflowlog) {
		gslbvserver.Appflowlog = data.Appflowlog.ValueString()
		hasChange = true
	}
	if changed(data.Backupip, state.Backupip) {
		gslbvserver.Backupip = data.Backupip.ValueString()
		hasChange = true
	}
	if changed(data.Backuplbmethod, state.Backuplbmethod) {
		gslbvserver.Backuplbmethod = data.Backuplbmethod.ValueString()
		hasChange = true
	}
	if changed(data.Backupsessiontimeout, state.Backupsessiontimeout) {
		gslbvserver.Backupsessiontimeout = utils.IntPtr(int(data.Backupsessiontimeout.ValueInt64()))
		hasChange = true
	}
	if changed(data.Backupvserver, state.Backupvserver) {
		gslbvserver.Backupvserver = data.Backupvserver.ValueString()
		hasChange = true
	}
	if changed(data.Comment, state.Comment) {
		gslbvserver.Comment = data.Comment.ValueString()
		hasChange = true
	}
	if changed(data.Considereffectivestate, state.Considereffectivestate) {
		gslbvserver.Considereffectivestate = data.Considereffectivestate.ValueString()
		hasChange = true
	}
	if changed(data.CookieDomain, state.CookieDomain) {
		gslbvserver.Cookiedomain = data.CookieDomain.ValueString()
		hasChange = true
	}
	if changed(data.Cookietimeout, state.Cookietimeout) {
		gslbvserver.Cookietimeout = utils.IntPtr(int(data.Cookietimeout.ValueInt64()))
		hasChange = true
	}
	if changed(data.Disableprimaryondown, state.Disableprimaryondown) {
		gslbvserver.Disableprimaryondown = data.Disableprimaryondown.ValueString()
		hasChange = true
	}
	if changed(data.Dnsrecordtype, state.Dnsrecordtype) {
		gslbvserver.Dnsrecordtype = data.Dnsrecordtype.ValueString()
		hasChange = true
	}
	if changed(data.Domainname, state.Domainname) {
		gslbvserver.Domainname = data.Domainname.ValueString()
		hasChange = true
	}
	if changed(data.Dynamicweight, state.Dynamicweight) {
		gslbvserver.Dynamicweight = data.Dynamicweight.ValueString()
		hasChange = true
	}
	if changed(data.Ecs, state.Ecs) {
		gslbvserver.Ecs = data.Ecs.ValueString()
		hasChange = true
	}
	if changed(data.Ecsaddrvalidation, state.Ecsaddrvalidation) {
		gslbvserver.Ecsaddrvalidation = data.Ecsaddrvalidation.ValueString()
		hasChange = true
	}
	if changed(data.Edr, state.Edr) {
		gslbvserver.Edr = data.Edr.ValueString()
		hasChange = true
	}
	if changed(data.Iptype, state.Iptype) {
		gslbvserver.Iptype = data.Iptype.ValueString()
		hasChange = true
	}
	if changed(data.Lbmethod, state.Lbmethod) {
		gslbvserver.Lbmethod = data.Lbmethod.ValueString()
		hasChange = true
	}
	if changed(data.Mir, state.Mir) {
		gslbvserver.Mir = data.Mir.ValueString()
		hasChange = true
	}
	if changed(data.Netmask, state.Netmask) {
		gslbvserver.Netmask = data.Netmask.ValueString()
		hasChange = true
	}
	if changed(data.Order, state.Order) {
		gslbvserver.Order = utils.IntPtr(int(data.Order.ValueInt64()))
		hasChange = true
	}
	if changed(data.Orderthreshold, state.Orderthreshold) {
		gslbvserver.Orderthreshold = utils.IntPtr(int(data.Orderthreshold.ValueInt64()))
		hasChange = true
	}
	if changed(data.Persistenceid, state.Persistenceid) {
		gslbvserver.Persistenceid = utils.IntPtr(int(data.Persistenceid.ValueInt64()))
		hasChange = true
	}
	if changed(data.Persistencetype, state.Persistencetype) {
		gslbvserver.Persistencetype = data.Persistencetype.ValueString()
		hasChange = true
	}
	if changed(data.Persistmask, state.Persistmask) {
		gslbvserver.Persistmask = data.Persistmask.ValueString()
		hasChange = true
	}
	if changed(data.Rule, state.Rule) {
		gslbvserver.Rule = data.Rule.ValueString()
		hasChange = true
	}
	if changed(data.Servicegroupname, state.Servicegroupname) {
		gslbvserver.Servicegroupname = data.Servicegroupname.ValueString()
		hasChange = true
	}
	if changed(data.Servicename, state.Servicename) {
		gslbvserver.Servicename = data.Servicename.ValueString()
		hasChange = true
	}
	if changed(data.Sitedomainttl, state.Sitedomainttl) {
		gslbvserver.Sitedomainttl = utils.IntPtr(int(data.Sitedomainttl.ValueInt64()))
		hasChange = true
	}
	if changed(data.Sobackupaction, state.Sobackupaction) {
		gslbvserver.Sobackupaction = data.Sobackupaction.ValueString()
		hasChange = true
	}
	if changed(data.Somethod, state.Somethod) {
		gslbvserver.Somethod = data.Somethod.ValueString()
		hasChange = true
	}
	if changed(data.Sopersistence, state.Sopersistence) {
		gslbvserver.Sopersistence = data.Sopersistence.ValueString()
		hasChange = true
	}
	if changed(data.Sopersistencetimeout, state.Sopersistencetimeout) {
		gslbvserver.Sopersistencetimeout = utils.IntPtr(int(data.Sopersistencetimeout.ValueInt64()))
		hasChange = true
	}
	if changed(data.Sothreshold, state.Sothreshold) {
		gslbvserver.Sothreshold = utils.IntPtr(int(data.Sothreshold.ValueInt64()))
		hasChange = true
	}
	if changed(data.Timeout, state.Timeout) {
		gslbvserver.Timeout = utils.IntPtr(int(data.Timeout.ValueInt64()))
		hasChange = true
	}
	if changed(data.Toggleorder, state.Toggleorder) {
		gslbvserver.Toggleorder = data.Toggleorder.ValueString()
		hasChange = true
	}
	if changed(data.Tolerance, state.Tolerance) {
		gslbvserver.Tolerance = utils.IntPtr(int(data.Tolerance.ValueInt64()))
		hasChange = true
	}
	if changed(data.Ttl, state.Ttl) {
		gslbvserver.Ttl = utils.IntPtr(int(data.Ttl.ValueInt64()))
		hasChange = true
	}
	if changed(data.V6netmasklen, state.V6netmasklen) {
		gslbvserver.V6netmasklen = utils.IntPtr(int(data.V6netmasklen.ValueInt64()))
		hasChange = true
	}
	if changed(data.V6persistmasklen, state.V6persistmasklen) {
		gslbvserver.V6persistmasklen = utils.IntPtr(int(data.V6persistmasklen.ValueInt64()))
		hasChange = true
	}
	if changed(data.Weight, state.Weight) {
		gslbvserver.Weight = utils.IntPtr(int(data.Weight.ValueInt64()))
		hasChange = true
	}

	return gslbvserver, hasChange
}

// gslbvserverSetAttrFromGet populates the resource model from a GET response while
// preserving prior plan/state for attributes the NITRO API omits. For Optional+
// Computed attributes NITRO frequently omits values equal to their server default
// (e.g. numeric 0), so the else-branch only nulls a value that is still Unknown -
// a configured value is preserved to avoid an "inconsistent result after apply".
func gslbvserverSetAttrFromGet(ctx context.Context, data *GslbvserverResourceModel, getResponseData map[string]interface{}) *GslbvserverResourceModel {
	tflog.Debug(ctx, "In gslbvserverSetAttrFromGet Function")

	if val, ok := getResponseData["appflowlog"]; ok && val != nil {
		data.Appflowlog = types.StringValue(val.(string))
	} else if data.Appflowlog.IsUnknown() {
		data.Appflowlog = types.StringNull()
	}
	if val, ok := getResponseData["backupip"]; ok && val != nil {
		data.Backupip = types.StringValue(val.(string))
	} else if data.Backupip.IsUnknown() {
		data.Backupip = types.StringNull()
	}
	if val, ok := getResponseData["backuplbmethod"]; ok && val != nil {
		data.Backuplbmethod = types.StringValue(val.(string))
	} else if data.Backuplbmethod.IsUnknown() {
		data.Backuplbmethod = types.StringNull()
	}
	if val, ok := getResponseData["backupsessiontimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Backupsessiontimeout = types.Int64Value(intVal)
		}
	} else if data.Backupsessiontimeout.IsUnknown() {
		data.Backupsessiontimeout = types.Int64Null()
	}
	if val, ok := getResponseData["backupvserver"]; ok && val != nil {
		data.Backupvserver = types.StringValue(val.(string))
	} else if data.Backupvserver.IsUnknown() {
		data.Backupvserver = types.StringNull()
	}
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else if data.Comment.IsUnknown() {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["considereffectivestate"]; ok && val != nil {
		data.Considereffectivestate = types.StringValue(val.(string))
	} else if data.Considereffectivestate.IsUnknown() {
		data.Considereffectivestate = types.StringNull()
	}
	if val, ok := getResponseData["cookie_domain"]; ok && val != nil {
		data.CookieDomain = types.StringValue(val.(string))
	} else if data.CookieDomain.IsUnknown() {
		data.CookieDomain = types.StringNull()
	}
	if val, ok := getResponseData["cookietimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Cookietimeout = types.Int64Value(intVal)
		}
	} else if data.Cookietimeout.IsUnknown() {
		data.Cookietimeout = types.Int64Null()
	}
	if val, ok := getResponseData["disableprimaryondown"]; ok && val != nil {
		data.Disableprimaryondown = types.StringValue(val.(string))
	} else if data.Disableprimaryondown.IsUnknown() {
		data.Disableprimaryondown = types.StringNull()
	}
	if val, ok := getResponseData["dnsrecordtype"]; ok && val != nil {
		data.Dnsrecordtype = types.StringValue(val.(string))
	} else if data.Dnsrecordtype.IsUnknown() {
		data.Dnsrecordtype = types.StringNull()
	}
	if val, ok := getResponseData["domainname"]; ok && val != nil {
		data.Domainname = types.StringValue(val.(string))
	} else if data.Domainname.IsUnknown() {
		data.Domainname = types.StringNull()
	}
	if val, ok := getResponseData["dynamicweight"]; ok && val != nil {
		data.Dynamicweight = types.StringValue(val.(string))
	} else if data.Dynamicweight.IsUnknown() {
		data.Dynamicweight = types.StringNull()
	}
	if val, ok := getResponseData["ecs"]; ok && val != nil {
		data.Ecs = types.StringValue(val.(string))
	} else if data.Ecs.IsUnknown() {
		data.Ecs = types.StringNull()
	}
	if val, ok := getResponseData["ecsaddrvalidation"]; ok && val != nil {
		data.Ecsaddrvalidation = types.StringValue(val.(string))
	} else if data.Ecsaddrvalidation.IsUnknown() {
		data.Ecsaddrvalidation = types.StringNull()
	}
	if val, ok := getResponseData["edr"]; ok && val != nil {
		data.Edr = types.StringValue(val.(string))
	} else if data.Edr.IsUnknown() {
		data.Edr = types.StringNull()
	}
	if val, ok := getResponseData["iptype"]; ok && val != nil {
		data.Iptype = types.StringValue(val.(string))
	} else if data.Iptype.IsUnknown() {
		data.Iptype = types.StringNull()
	}
	if val, ok := getResponseData["lbmethod"]; ok && val != nil {
		data.Lbmethod = types.StringValue(val.(string))
	} else if data.Lbmethod.IsUnknown() {
		data.Lbmethod = types.StringNull()
	}
	if val, ok := getResponseData["mir"]; ok && val != nil {
		data.Mir = types.StringValue(val.(string))
	} else if data.Mir.IsUnknown() {
		data.Mir = types.StringNull()
	}
	// name is the user-facing key. After a rename (via newname) the live object name
	// (tracked by data.Id) diverges from the configured name and GET returns the live
	// (new) name. Overwriting name from GET would clobber the user's configured value
	// and trigger a spurious RequiresReplace diff, so only adopt the GET value when we
	// don't already have one (e.g. on import, where state carries only the ID).
	if data.Name.IsNull() || data.Name.IsUnknown() || data.Name.ValueString() == "" {
		if val, ok := getResponseData["name"]; ok && val != nil {
			data.Name = types.StringValue(val.(string))
		}
	}
	if val, ok := getResponseData["netmask"]; ok && val != nil {
		data.Netmask = types.StringValue(val.(string))
	} else if data.Netmask.IsUnknown() {
		data.Netmask = types.StringNull()
	}
	// newname is rename-only and never echoed by GET; preserve plan/state value.
	if val, ok := getResponseData["order"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Order = types.Int64Value(intVal)
		}
	} else if data.Order.IsUnknown() {
		data.Order = types.Int64Null()
	}
	if val, ok := getResponseData["orderthreshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Orderthreshold = types.Int64Value(intVal)
		}
	} else if data.Orderthreshold.IsUnknown() {
		data.Orderthreshold = types.Int64Null()
	}
	if val, ok := getResponseData["persistenceid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Persistenceid = types.Int64Value(intVal)
		}
	} else if data.Persistenceid.IsUnknown() {
		data.Persistenceid = types.Int64Null()
	}
	if val, ok := getResponseData["persistencetype"]; ok && val != nil {
		data.Persistencetype = types.StringValue(val.(string))
	} else if data.Persistencetype.IsUnknown() {
		data.Persistencetype = types.StringNull()
	}
	if val, ok := getResponseData["persistmask"]; ok && val != nil {
		data.Persistmask = types.StringValue(val.(string))
	} else if data.Persistmask.IsUnknown() {
		data.Persistmask = types.StringNull()
	}
	if val, ok := getResponseData["rule"]; ok && val != nil {
		data.Rule = types.StringValue(val.(string))
	} else if data.Rule.IsUnknown() {
		data.Rule = types.StringNull()
	}
	if val, ok := getResponseData["servicegroupname"]; ok && val != nil {
		data.Servicegroupname = types.StringValue(val.(string))
	} else if data.Servicegroupname.IsUnknown() {
		data.Servicegroupname = types.StringNull()
	}
	if val, ok := getResponseData["servicename"]; ok && val != nil {
		data.Servicename = types.StringValue(val.(string))
	} else if data.Servicename.IsUnknown() {
		data.Servicename = types.StringNull()
	}
	if val, ok := getResponseData["servicetype"]; ok && val != nil {
		data.Servicetype = types.StringValue(val.(string))
	} else if data.Servicetype.IsUnknown() {
		data.Servicetype = types.StringNull()
	}
	if val, ok := getResponseData["sitedomainttl"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Sitedomainttl = types.Int64Value(intVal)
		}
	} else if data.Sitedomainttl.IsUnknown() {
		data.Sitedomainttl = types.Int64Null()
	}
	if val, ok := getResponseData["sobackupaction"]; ok && val != nil {
		data.Sobackupaction = types.StringValue(val.(string))
	} else if data.Sobackupaction.IsUnknown() {
		data.Sobackupaction = types.StringNull()
	}
	if val, ok := getResponseData["somethod"]; ok && val != nil {
		data.Somethod = types.StringValue(val.(string))
	} else if data.Somethod.IsUnknown() {
		data.Somethod = types.StringNull()
	}
	if val, ok := getResponseData["sopersistence"]; ok && val != nil {
		data.Sopersistence = types.StringValue(val.(string))
	} else if data.Sopersistence.IsUnknown() {
		data.Sopersistence = types.StringNull()
	}
	if val, ok := getResponseData["sopersistencetimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Sopersistencetimeout = types.Int64Value(intVal)
		}
	} else if data.Sopersistencetimeout.IsUnknown() {
		data.Sopersistencetimeout = types.Int64Null()
	}
	if val, ok := getResponseData["sothreshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Sothreshold = types.Int64Value(intVal)
		}
	} else if data.Sothreshold.IsUnknown() {
		data.Sothreshold = types.Int64Null()
	}
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	} else if data.State.IsUnknown() {
		data.State = types.StringNull()
	}
	if val, ok := getResponseData["timeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Timeout = types.Int64Value(intVal)
		}
	} else if data.Timeout.IsUnknown() {
		data.Timeout = types.Int64Null()
	}
	if val, ok := getResponseData["toggleorder"]; ok && val != nil {
		data.Toggleorder = types.StringValue(val.(string))
	} else if data.Toggleorder.IsUnknown() {
		data.Toggleorder = types.StringNull()
	}
	if val, ok := getResponseData["tolerance"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Tolerance = types.Int64Value(intVal)
		}
	} else if data.Tolerance.IsUnknown() {
		data.Tolerance = types.Int64Null()
	}
	if val, ok := getResponseData["ttl"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ttl = types.Int64Value(intVal)
		}
	} else if data.Ttl.IsUnknown() {
		data.Ttl = types.Int64Null()
	}
	if val, ok := getResponseData["v6netmasklen"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.V6netmasklen = types.Int64Value(intVal)
		}
	} else if data.V6netmasklen.IsUnknown() {
		data.V6netmasklen = types.Int64Null()
	}
	if val, ok := getResponseData["v6persistmasklen"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.V6persistmasklen = types.Int64Value(intVal)
		}
	} else if data.V6persistmasklen.IsUnknown() {
		data.V6persistmasklen = types.Int64Null()
	}
	if val, ok := getResponseData["weight"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Weight = types.Int64Value(intVal)
		}
	} else if data.Weight.IsUnknown() {
		data.Weight = types.Int64Null()
	}

	return data
}

// gslbvserverSetAttrFromGetForDatasource faithfully copies every field from the GET
// response. The datasource has no prior plan/state to preserve, so it populates the
// model directly from the API response and sets the ID itself.
func gslbvserverSetAttrFromGetForDatasource(ctx context.Context, data *GslbvserverResourceModel, getResponseData map[string]interface{}) *GslbvserverResourceModel {
	tflog.Debug(ctx, "In gslbvserverSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["appflowlog"]; ok && val != nil {
		data.Appflowlog = types.StringValue(val.(string))
	} else {
		data.Appflowlog = types.StringNull()
	}
	if val, ok := getResponseData["backupip"]; ok && val != nil {
		data.Backupip = types.StringValue(val.(string))
	} else {
		data.Backupip = types.StringNull()
	}
	if val, ok := getResponseData["backuplbmethod"]; ok && val != nil {
		data.Backuplbmethod = types.StringValue(val.(string))
	} else {
		data.Backuplbmethod = types.StringNull()
	}
	if val, ok := getResponseData["backupsessiontimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Backupsessiontimeout = types.Int64Value(intVal)
		}
	} else {
		data.Backupsessiontimeout = types.Int64Null()
	}
	if val, ok := getResponseData["backupvserver"]; ok && val != nil {
		data.Backupvserver = types.StringValue(val.(string))
	} else {
		data.Backupvserver = types.StringNull()
	}
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["considereffectivestate"]; ok && val != nil {
		data.Considereffectivestate = types.StringValue(val.(string))
	} else {
		data.Considereffectivestate = types.StringNull()
	}
	if val, ok := getResponseData["cookie_domain"]; ok && val != nil {
		data.CookieDomain = types.StringValue(val.(string))
	} else {
		data.CookieDomain = types.StringNull()
	}
	if val, ok := getResponseData["cookietimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Cookietimeout = types.Int64Value(intVal)
		}
	} else {
		data.Cookietimeout = types.Int64Null()
	}
	if val, ok := getResponseData["disableprimaryondown"]; ok && val != nil {
		data.Disableprimaryondown = types.StringValue(val.(string))
	} else {
		data.Disableprimaryondown = types.StringNull()
	}
	if val, ok := getResponseData["dnsrecordtype"]; ok && val != nil {
		data.Dnsrecordtype = types.StringValue(val.(string))
	} else {
		data.Dnsrecordtype = types.StringNull()
	}
	if val, ok := getResponseData["domainname"]; ok && val != nil {
		data.Domainname = types.StringValue(val.(string))
	} else {
		data.Domainname = types.StringNull()
	}
	if val, ok := getResponseData["dynamicweight"]; ok && val != nil {
		data.Dynamicweight = types.StringValue(val.(string))
	} else {
		data.Dynamicweight = types.StringNull()
	}
	if val, ok := getResponseData["ecs"]; ok && val != nil {
		data.Ecs = types.StringValue(val.(string))
	} else {
		data.Ecs = types.StringNull()
	}
	if val, ok := getResponseData["ecsaddrvalidation"]; ok && val != nil {
		data.Ecsaddrvalidation = types.StringValue(val.(string))
	} else {
		data.Ecsaddrvalidation = types.StringNull()
	}
	if val, ok := getResponseData["edr"]; ok && val != nil {
		data.Edr = types.StringValue(val.(string))
	} else {
		data.Edr = types.StringNull()
	}
	if val, ok := getResponseData["iptype"]; ok && val != nil {
		data.Iptype = types.StringValue(val.(string))
	} else {
		data.Iptype = types.StringNull()
	}
	if val, ok := getResponseData["lbmethod"]; ok && val != nil {
		data.Lbmethod = types.StringValue(val.(string))
	} else {
		data.Lbmethod = types.StringNull()
	}
	if val, ok := getResponseData["mir"]; ok && val != nil {
		data.Mir = types.StringValue(val.(string))
	} else {
		data.Mir = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["netmask"]; ok && val != nil {
		data.Netmask = types.StringValue(val.(string))
	} else {
		data.Netmask = types.StringNull()
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
	if val, ok := getResponseData["orderthreshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Orderthreshold = types.Int64Value(intVal)
		}
	} else {
		data.Orderthreshold = types.Int64Null()
	}
	if val, ok := getResponseData["persistenceid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Persistenceid = types.Int64Value(intVal)
		}
	} else {
		data.Persistenceid = types.Int64Null()
	}
	if val, ok := getResponseData["persistencetype"]; ok && val != nil {
		data.Persistencetype = types.StringValue(val.(string))
	} else {
		data.Persistencetype = types.StringNull()
	}
	if val, ok := getResponseData["persistmask"]; ok && val != nil {
		data.Persistmask = types.StringValue(val.(string))
	} else {
		data.Persistmask = types.StringNull()
	}
	if val, ok := getResponseData["rule"]; ok && val != nil {
		data.Rule = types.StringValue(val.(string))
	} else {
		data.Rule = types.StringNull()
	}
	if val, ok := getResponseData["servicegroupname"]; ok && val != nil {
		data.Servicegroupname = types.StringValue(val.(string))
	} else {
		data.Servicegroupname = types.StringNull()
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
	if val, ok := getResponseData["sitedomainttl"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Sitedomainttl = types.Int64Value(intVal)
		}
	} else {
		data.Sitedomainttl = types.Int64Null()
	}
	if val, ok := getResponseData["sobackupaction"]; ok && val != nil {
		data.Sobackupaction = types.StringValue(val.(string))
	} else {
		data.Sobackupaction = types.StringNull()
	}
	if val, ok := getResponseData["somethod"]; ok && val != nil {
		data.Somethod = types.StringValue(val.(string))
	} else {
		data.Somethod = types.StringNull()
	}
	if val, ok := getResponseData["sopersistence"]; ok && val != nil {
		data.Sopersistence = types.StringValue(val.(string))
	} else {
		data.Sopersistence = types.StringNull()
	}
	if val, ok := getResponseData["sopersistencetimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Sopersistencetimeout = types.Int64Value(intVal)
		}
	} else {
		data.Sopersistencetimeout = types.Int64Null()
	}
	if val, ok := getResponseData["sothreshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Sothreshold = types.Int64Value(intVal)
		}
	} else {
		data.Sothreshold = types.Int64Null()
	}
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	} else {
		data.State = types.StringNull()
	}
	if val, ok := getResponseData["timeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Timeout = types.Int64Value(intVal)
		}
	} else {
		data.Timeout = types.Int64Null()
	}
	if val, ok := getResponseData["toggleorder"]; ok && val != nil {
		data.Toggleorder = types.StringValue(val.(string))
	} else {
		data.Toggleorder = types.StringNull()
	}
	if val, ok := getResponseData["tolerance"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Tolerance = types.Int64Value(intVal)
		}
	} else {
		data.Tolerance = types.Int64Null()
	}
	if val, ok := getResponseData["ttl"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ttl = types.Int64Value(intVal)
		}
	} else {
		data.Ttl = types.Int64Null()
	}
	if val, ok := getResponseData["v6netmasklen"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.V6netmasklen = types.Int64Value(intVal)
		}
	} else {
		data.V6netmasklen = types.Int64Null()
	}
	if val, ok := getResponseData["v6persistmasklen"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.V6persistmasklen = types.Int64Value(intVal)
		}
	} else {
		data.V6persistmasklen = types.Int64Null()
	}
	if val, ok := getResponseData["weight"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Weight = types.Int64Value(intVal)
		}
	} else {
		data.Weight = types.Int64Null()
	}

	// Single unique attribute - use plain value as ID.
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}
