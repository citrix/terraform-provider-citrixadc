package dnsparameter

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/dns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// DnsparameterResourceModel describes the resource data model.
type DnsparameterResourceModel struct {
	Id                           types.String `tfsdk:"id"`
	Autosavekeyops               types.String `tfsdk:"autosavekeyops"`
	Cacheecszeroprefix           types.String `tfsdk:"cacheecszeroprefix"`
	Cachehitbypass               types.String `tfsdk:"cachehitbypass"`
	Cachenoexpire                types.String `tfsdk:"cachenoexpire"`
	Cacherecords                 types.String `tfsdk:"cacherecords"`
	Dns64timeout                 types.Int64  `tfsdk:"dns64timeout"`
	Dnsrootreferral              types.String `tfsdk:"dnsrootreferral"`
	Dnssec                       types.String `tfsdk:"dnssec"`
	Ecsmaxsubnets                types.Int64  `tfsdk:"ecsmaxsubnets"`
	Maxcachesize                 types.Int64  `tfsdk:"maxcachesize"`
	Maxnegativecachesize         types.Int64  `tfsdk:"maxnegativecachesize"`
	Maxnegcachettl               types.Int64  `tfsdk:"maxnegcachettl"`
	Maxpipeline                  types.Int64  `tfsdk:"maxpipeline"`
	Maxttl                       types.Int64  `tfsdk:"maxttl"`
	Maxudppacketsize             types.Int64  `tfsdk:"maxudppacketsize"`
	Minttl                       types.Int64  `tfsdk:"minttl"`
	Namelookuppriority           types.String `tfsdk:"namelookuppriority"`
	Nxdomainratelimitthreshold   types.Int64  `tfsdk:"nxdomainratelimitthreshold"`
	Recursion                    types.String `tfsdk:"recursion"`
	Resolutionorder              types.String `tfsdk:"resolutionorder"`
	Resolvermaxactiveresolutions types.Int64  `tfsdk:"resolvermaxactiveresolutions"`
	Resolvermaxtcpconnections    types.Int64  `tfsdk:"resolvermaxtcpconnections"`
	Resolvermaxtcptimeout        types.Int64  `tfsdk:"resolvermaxtcptimeout"`
	Retries                      types.Int64  `tfsdk:"retries"`
	Splitpktqueryprocessing      types.String `tfsdk:"splitpktqueryprocessing"`
	Zonetransfer                 types.String `tfsdk:"zonetransfer"`
}

func (r *DnsparameterResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the dnsparameter resource.",
			},
			"autosavekeyops": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Flag to enable/disable saving of rollover operations executed automatically to avoid config loss.\nApplicable only when autorollover option is enabled on a key. Note: when you enable this, full configuration will be saved",
			},
			"cacheecszeroprefix": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Cache ECS responses with a Scope Prefix length of zero. Such a cached response will be used for all queries with this domain name and any subnet. When disabled, ECS responses with Scope Prefix length of zero will be cached, but not tied to any subnet. This option has no effect if caching of ECS responses is disabled in the corresponding DNS profile.",
			},
			"cachehitbypass": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "This parameter is applicable only in proxy mode and if this parameter is enabled  we will forward all the client requests to the backend DNS server and the response served will be cached on Citrix ADC",
			},
			"cachenoexpire": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "If this flag is set to YES, the existing entries in cache do not age out. On reaching the max limit the cache records are frozen",
			},
			"cacherecords": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("True"),
				Description: "Cache resource records in the DNS cache. Applies to resource records obtained through proxy configurations only. End resolver and forwarder configurations always cache records in the DNS cache, and you cannot disable this behavior. When you disable record caching, the appliance stops caching server responses. However, cached records are not flushed. The appliance does not serve requests from the cache until record caching is enabled again.",
			},
			"dns64timeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "While doing DNS64 resolution, this parameter specifies the time to wait before sending an A query if no response is received from backend DNS server for AAAA query.",
			},
			"dnsrootreferral": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Send a root referral if a client queries a domain name that is unrelated to the domains configured/cached on the Citrix ADC. If the setting is disabled, the appliance sends a blank response instead of a root referral. Applicable to domains for which the appliance is authoritative. Disable the parameter when the appliance is under attack from a client that is sending a flood of queries for unrelated domains.",
			},
			"dnssec": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Enable or disable the Domain Name System Security Extensions (DNSSEC) feature on the appliance. Note: Even when the DNSSEC feature is enabled, forwarder configurations (used by internal Citrix ADC features such as SSL VPN and Cache Redirection for name resolution) do not support the DNSSEC OK (DO) bit in the EDNS0 OPT header.",
			},
			"ecsmaxsubnets": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of subnets that can be cached corresponding to a single domain. Subnet caching will occur for responses with EDNS Client Subnet (ECS) option. Caching of such responses can be disabled using DNS profile settings. A value of zero indicates that the number of subnets cached is limited only by existing memory constraints. The default value is zero.",
			},
			"maxcachesize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum memory, in megabytes, that can be used for dns caching per Packet Engine.",
			},
			"maxnegativecachesize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum memory, in megabytes, that can be used for caching of negative DNS responses per packet engine.",
			},
			"maxnegcachettl": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(604800),
				Description: "Maximum time to live (TTL) for all negative records ( NXDONAIN and NODATA ) cached in the DNS cache by DNS proxy, end resolver, and forwarder configurations. If the TTL of a record that is to be cached is higher than the value configured for maxnegcacheTTL, the TTL of the record is set to the value of maxnegcacheTTL before caching. When you modify this setting, the new value is applied only to those records that are cached after the modification. The TTL values of existing records are not changed.",
			},
			"maxpipeline": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of concurrent DNS requests to allow on a single client connection, which is identified by the <clientip:port>-<vserver ip:port> tuple. A value of 0 (zero) applies no limit to the number of concurrent DNS requests allowed on a single client connection.",
			},
			"maxttl": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(604800),
				Description: "Maximum time to live (TTL) for all records cached in the DNS cache by DNS proxy, end resolver, and forwarder configurations. If the TTL of a record that is to be cached is higher than the value configured for maxTTL, the TTL of the record is set to the value of maxTTL before caching. When you modify this setting, the new value is applied only to those records that are cached after the modification. The TTL values of existing records are not changed.",
			},
			"maxudppacketsize": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(1280),
				Description: "Maximum UDP packet size that can be handled by Citrix ADC. This is the value advertised by Citrix ADC when responding as an authoritative server and it is also used when Citrix ADC queries other name servers as a forwarder. When acting as a proxy, requests from clients are limited by this parameter - if a request contains a size greater than this value in the OPT record, it will be replaced.",
			},
			"minttl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum permissible time to live (TTL) for all records cached in the DNS cache by DNS proxy, end resolver, and forwarder configurations. If the TTL of a record that is to be cached is lower than the value configured for minTTL, the TTL of the record is set to the value of minTTL before caching. When you modify this setting, the new value is applied only to those records that are cached after the modification. The TTL values of existing records are not changed.",
			},
			"namelookuppriority": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("WINS"),
				Description: "Type of lookup (DNS or WINS) to attempt first. If the first-priority lookup fails, the second-priority lookup is attempted. Used only by the SSL VPN feature.",
			},
			"nxdomainratelimitthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Rate limit threshold for Non-Existant domain (NXDOMAIN) responses generated from Citrix ADC. Once the threshold is breached , DNS queries leading to NXDOMAIN response will be dropped. This threshold will not be applied for NXDOMAIN responses got from the backend. The threshold will be applied per packet engine and per second.",
			},
			"recursion": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Function as an end resolver and recursively resolve queries for domains that are not hosted on the Citrix ADC. Also resolve queries recursively when the external name servers configured on the appliance (for a forwarder configuration) are unavailable. When external name servers are unavailable, the appliance queries a root server and resolves the request recursively, as it does for an end resolver configuration.",
			},
			"resolutionorder": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("OnlyAQuery"),
				Description: "Type of DNS queries (A, AAAA, or both) to generate during the routine functioning of certain Citrix ADC features, such as SSL VPN, cache redirection, and the integrated cache. The queries are sent to the external name servers that are configured for the forwarder function. If you specify both query types, you can also specify the order. Available settings function as follows:\n* OnlyAQuery. Send queries for IPv4 address records (A records) only.\n* OnlyAAAAQuery. Send queries for IPv6 address records (AAAA records) instead of queries for IPv4 address records (A records).\n* AThenAAAAQuery. Send a query for an A record, and then send a query for an AAAA record if the query for the A record results in a NODATA response from the name server.\n* AAAAThenAQuery. Send a query for an AAAA record, and then send a query for an A record if the query for the AAAA record results in a NODATA response from the name server.",
			},
			"resolvermaxactiveresolutions": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of active concurrent DNS resolutions per Packet Engine",
			},
			"resolvermaxtcpconnections": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(1000),
				Description: "Maximum DNS-TCP connections opened for recursive resolution per Packet Engine",
			},
			"resolvermaxtcptimeout": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(5),
				Description: "Maximum wait time in seconds for the response on DNS-TCP connection for recursive resolution per Packet Engine",
			},
			"retries": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(5),
				Description: "Maximum number of retry attempts when no response is received for a query sent to a name server. Applies to end resolver and forwarder configurations.",
			},
			"splitpktqueryprocessing": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ALLOW"),
				Description: "Processing requests split across multiple packets",
			},
			"zonetransfer": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Flag to enable/disable DNS zones configuration transfer to remote GSLB site nodes",
			},
		},
	}
}

func dnsparameterGetThePayloadFromtheConfig(ctx context.Context, data *DnsparameterResourceModel) dns.Dnsparameter {
	tflog.Debug(ctx, "In dnsparameterGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	dnsparameter := dns.Dnsparameter{}
	if !data.Autosavekeyops.IsNull() {
		dnsparameter.Autosavekeyops = data.Autosavekeyops.ValueString()
	}
	if !data.Cacheecszeroprefix.IsNull() {
		dnsparameter.Cacheecszeroprefix = data.Cacheecszeroprefix.ValueString()
	}
	if !data.Cachehitbypass.IsNull() {
		dnsparameter.Cachehitbypass = data.Cachehitbypass.ValueString()
	}
	if !data.Cachenoexpire.IsNull() {
		dnsparameter.Cachenoexpire = data.Cachenoexpire.ValueString()
	}
	if !data.Cacherecords.IsNull() {
		dnsparameter.Cacherecords = data.Cacherecords.ValueString()
	}
	if !data.Dns64timeout.IsNull() {
		dnsparameter.Dns64timeout = utils.IntPtr(int(data.Dns64timeout.ValueInt64()))
	}
	if !data.Dnsrootreferral.IsNull() {
		dnsparameter.Dnsrootreferral = data.Dnsrootreferral.ValueString()
	}
	if !data.Dnssec.IsNull() {
		dnsparameter.Dnssec = data.Dnssec.ValueString()
	}
	if !data.Ecsmaxsubnets.IsNull() {
		dnsparameter.Ecsmaxsubnets = utils.IntPtr(int(data.Ecsmaxsubnets.ValueInt64()))
	}
	if !data.Maxcachesize.IsNull() {
		dnsparameter.Maxcachesize = utils.IntPtr(int(data.Maxcachesize.ValueInt64()))
	}
	if !data.Maxnegativecachesize.IsNull() {
		dnsparameter.Maxnegativecachesize = utils.IntPtr(int(data.Maxnegativecachesize.ValueInt64()))
	}
	if !data.Maxnegcachettl.IsNull() {
		dnsparameter.Maxnegcachettl = utils.IntPtr(int(data.Maxnegcachettl.ValueInt64()))
	}
	if !data.Maxpipeline.IsNull() {
		dnsparameter.Maxpipeline = utils.IntPtr(int(data.Maxpipeline.ValueInt64()))
	}
	if !data.Maxttl.IsNull() {
		dnsparameter.Maxttl = utils.IntPtr(int(data.Maxttl.ValueInt64()))
	}
	if !data.Maxudppacketsize.IsNull() {
		dnsparameter.Maxudppacketsize = utils.IntPtr(int(data.Maxudppacketsize.ValueInt64()))
	}
	if !data.Minttl.IsNull() {
		dnsparameter.Minttl = utils.IntPtr(int(data.Minttl.ValueInt64()))
	}
	if !data.Namelookuppriority.IsNull() {
		dnsparameter.Namelookuppriority = data.Namelookuppriority.ValueString()
	}
	if !data.Nxdomainratelimitthreshold.IsNull() {
		dnsparameter.Nxdomainratelimitthreshold = utils.IntPtr(int(data.Nxdomainratelimitthreshold.ValueInt64()))
	}
	if !data.Recursion.IsNull() {
		dnsparameter.Recursion = data.Recursion.ValueString()
	}
	if !data.Resolutionorder.IsNull() {
		dnsparameter.Resolutionorder = data.Resolutionorder.ValueString()
	}
	if !data.Resolvermaxactiveresolutions.IsNull() {
		dnsparameter.Resolvermaxactiveresolutions = utils.IntPtr(int(data.Resolvermaxactiveresolutions.ValueInt64()))
	}
	if !data.Resolvermaxtcpconnections.IsNull() {
		dnsparameter.Resolvermaxtcpconnections = utils.IntPtr(int(data.Resolvermaxtcpconnections.ValueInt64()))
	}
	if !data.Resolvermaxtcptimeout.IsNull() {
		dnsparameter.Resolvermaxtcptimeout = utils.IntPtr(int(data.Resolvermaxtcptimeout.ValueInt64()))
	}
	if !data.Retries.IsNull() {
		dnsparameter.Retries = utils.IntPtr(int(data.Retries.ValueInt64()))
	}
	if !data.Splitpktqueryprocessing.IsNull() {
		dnsparameter.Splitpktqueryprocessing = data.Splitpktqueryprocessing.ValueString()
	}
	if !data.Zonetransfer.IsNull() {
		dnsparameter.Zonetransfer = data.Zonetransfer.ValueString()
	}

	return dnsparameter
}

func dnsparameterSetAttrFromGet(ctx context.Context, data *DnsparameterResourceModel, getResponseData map[string]interface{}) *DnsparameterResourceModel {
	tflog.Debug(ctx, "In dnsparameterSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["autosavekeyops"]; ok && val != nil {
		data.Autosavekeyops = types.StringValue(val.(string))
	} else {
		data.Autosavekeyops = types.StringNull()
	}
	if val, ok := getResponseData["cacheecszeroprefix"]; ok && val != nil {
		data.Cacheecszeroprefix = types.StringValue(val.(string))
	} else {
		data.Cacheecszeroprefix = types.StringNull()
	}
	if val, ok := getResponseData["cachehitbypass"]; ok && val != nil {
		data.Cachehitbypass = types.StringValue(val.(string))
	} else {
		data.Cachehitbypass = types.StringNull()
	}
	if val, ok := getResponseData["cachenoexpire"]; ok && val != nil {
		data.Cachenoexpire = types.StringValue(val.(string))
	} else {
		data.Cachenoexpire = types.StringNull()
	}
	if val, ok := getResponseData["cacherecords"]; ok && val != nil {
		data.Cacherecords = types.StringValue(val.(string))
	} else {
		data.Cacherecords = types.StringNull()
	}
	if val, ok := getResponseData["dns64timeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Dns64timeout = types.Int64Value(intVal)
		}
	} else {
		data.Dns64timeout = types.Int64Null()
	}
	if val, ok := getResponseData["dnsrootreferral"]; ok && val != nil {
		data.Dnsrootreferral = types.StringValue(val.(string))
	} else {
		data.Dnsrootreferral = types.StringNull()
	}
	if val, ok := getResponseData["dnssec"]; ok && val != nil {
		data.Dnssec = types.StringValue(val.(string))
	} else {
		data.Dnssec = types.StringNull()
	}
	if val, ok := getResponseData["ecsmaxsubnets"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ecsmaxsubnets = types.Int64Value(intVal)
		}
	} else {
		data.Ecsmaxsubnets = types.Int64Null()
	}
	if val, ok := getResponseData["maxcachesize"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxcachesize = types.Int64Value(intVal)
		}
	} else {
		data.Maxcachesize = types.Int64Null()
	}
	if val, ok := getResponseData["maxnegativecachesize"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxnegativecachesize = types.Int64Value(intVal)
		}
	} else {
		data.Maxnegativecachesize = types.Int64Null()
	}
	if val, ok := getResponseData["maxnegcachettl"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxnegcachettl = types.Int64Value(intVal)
		}
	} else {
		data.Maxnegcachettl = types.Int64Null()
	}
	if val, ok := getResponseData["maxpipeline"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxpipeline = types.Int64Value(intVal)
		}
	} else {
		data.Maxpipeline = types.Int64Null()
	}
	if val, ok := getResponseData["maxttl"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxttl = types.Int64Value(intVal)
		}
	} else {
		data.Maxttl = types.Int64Null()
	}
	if val, ok := getResponseData["maxudppacketsize"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxudppacketsize = types.Int64Value(intVal)
		}
	} else {
		data.Maxudppacketsize = types.Int64Null()
	}
	if val, ok := getResponseData["minttl"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Minttl = types.Int64Value(intVal)
		}
	} else {
		data.Minttl = types.Int64Null()
	}
	if val, ok := getResponseData["namelookuppriority"]; ok && val != nil {
		data.Namelookuppriority = types.StringValue(val.(string))
	} else {
		data.Namelookuppriority = types.StringNull()
	}
	if val, ok := getResponseData["nxdomainratelimitthreshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Nxdomainratelimitthreshold = types.Int64Value(intVal)
		}
	} else {
		data.Nxdomainratelimitthreshold = types.Int64Null()
	}
	if val, ok := getResponseData["recursion"]; ok && val != nil {
		data.Recursion = types.StringValue(val.(string))
	} else {
		data.Recursion = types.StringNull()
	}
	if val, ok := getResponseData["resolutionorder"]; ok && val != nil {
		data.Resolutionorder = types.StringValue(val.(string))
	} else {
		data.Resolutionorder = types.StringNull()
	}
	if val, ok := getResponseData["resolvermaxactiveresolutions"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Resolvermaxactiveresolutions = types.Int64Value(intVal)
		}
	} else {
		data.Resolvermaxactiveresolutions = types.Int64Null()
	}
	if val, ok := getResponseData["resolvermaxtcpconnections"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Resolvermaxtcpconnections = types.Int64Value(intVal)
		}
	} else {
		data.Resolvermaxtcpconnections = types.Int64Null()
	}
	if val, ok := getResponseData["resolvermaxtcptimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Resolvermaxtcptimeout = types.Int64Value(intVal)
		}
	} else {
		data.Resolvermaxtcptimeout = types.Int64Null()
	}
	if val, ok := getResponseData["retries"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Retries = types.Int64Value(intVal)
		}
	} else {
		data.Retries = types.Int64Null()
	}
	if val, ok := getResponseData["splitpktqueryprocessing"]; ok && val != nil {
		data.Splitpktqueryprocessing = types.StringValue(val.(string))
	} else {
		data.Splitpktqueryprocessing = types.StringNull()
	}
	if val, ok := getResponseData["zonetransfer"]; ok && val != nil {
		data.Zonetransfer = types.StringValue(val.(string))
	} else {
		data.Zonetransfer = types.StringNull()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("dnsparameter-config")

	return data
}
