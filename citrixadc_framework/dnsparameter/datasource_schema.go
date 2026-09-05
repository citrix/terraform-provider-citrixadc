package dnsparameter

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// DnsparameterDataSourceModel is the data-source-specific model, decoupled from
// DnsparameterResourceModel. A data source is a pure read surface, so it exposes
// the full GET projection: the read/write attributes (as Computed outputs) AND
// the read-only attributes the resource deliberately omits (builtin, feature,
// nxdomainthresholdcrossed). Every non-key attribute is Computed.
type DnsparameterDataSourceModel struct {
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

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/dnsparameter.json). Never settable; populated from GET.
	Builtin                  types.List   `tfsdk:"builtin"`
	Feature                  types.String `tfsdk:"feature"`
	Nxdomainthresholdcrossed types.Int64  `tfsdk:"nxdomainthresholdcrossed"`
}

func DnsparameterDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"autosavekeyops": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Flag to enable/disable saving of rollover operations executed automatically to avoid config loss.\nApplicable only when autorollover option is enabled on a key. Note: when you enable this, full configuration will be saved",
			},
			"cacheecszeroprefix": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Cache ECS responses with a Scope Prefix length of zero. Such a cached response will be used for all queries with this domain name and any subnet. When disabled, ECS responses with Scope Prefix length of zero will be cached, but not tied to any subnet. This option has no effect if caching of ECS responses is disabled in the corresponding DNS profile.",
			},
			"cachehitbypass": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This parameter is applicable only in proxy mode and if this parameter is enabled  we will forward all the client requests to the backend DNS server and the response served will be cached on Citrix ADC",
			},
			"cachenoexpire": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If this flag is set to YES, the existing entries in cache do not age out. On reaching the max limit the cache records are frozen",
			},
			"cacherecords": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Cache resource records in the DNS cache. Applies to resource records obtained through proxy configurations only. End resolver and forwarder configurations always cache records in the DNS cache, and you cannot disable this behavior. When you disable record caching, the appliance stops caching server responses. However, cached records are not flushed. The appliance does not serve requests from the cache until record caching is enabled again.",
			},
			"dns64timeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "While doing DNS64 resolution, this parameter specifies the time to wait before sending an A query if no response is received from backend DNS server for AAAA query.",
			},
			"dnsrootreferral": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Send a root referral if a client queries a domain name that is unrelated to the domains configured/cached on the Citrix ADC. If the setting is disabled, the appliance sends a blank response instead of a root referral. Applicable to domains for which the appliance is authoritative. Disable the parameter when the appliance is under attack from a client that is sending a flood of queries for unrelated domains.",
			},
			"dnssec": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
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
				Computed:    true,
				Description: "Maximum time to live (TTL) for all negative records ( NXDONAIN and NODATA ) cached in the DNS cache by DNS proxy, end resolver, and forwarder configurations. If the TTL of a record that is to be cached is higher than the value configured for maxnegcacheTTL, the TTL of the record is set to the value of maxnegcacheTTL before caching. When you modify this setting, the new value is applied only to those records that are cached after the modification. The TTL values of existing records are not changed.",
			},
			"maxpipeline": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of concurrent DNS requests to allow on a single client connection, which is identified by the <clientip:port>-<vserver ip:port> tuple. A value of 0 (zero) applies no limit to the number of concurrent DNS requests allowed on a single client connection.",
			},
			"maxttl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum time to live (TTL) for all records cached in the DNS cache by DNS proxy, end resolver, and forwarder configurations. If the TTL of a record that is to be cached is higher than the value configured for maxTTL, the TTL of the record is set to the value of maxTTL before caching. When you modify this setting, the new value is applied only to those records that are cached after the modification. The TTL values of existing records are not changed.",
			},
			"maxudppacketsize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum UDP packet size that can be handled by Citrix ADC. This is the value advertised by Citrix ADC when responding as an authoritative server and it is also used when Citrix ADC queries other name servers as a forwarder. When acting as a proxy, requests from clients are limited by this parameter - if a request contains a size greater than this value in the OPT record, it will be replaced.",
			},
			"minttl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum permissible time to live (TTL) for all records cached in the DNS cache by DNS proxy, end resolver, and forwarder configurations. If the TTL of a record that is to be cached is lower than the value configured for minTTL, the TTL of the record is set to the value of minTTL before caching. When you modify this setting, the new value is applied only to those records that are cached after the modification. The TTL values of existing records are not changed.",
			},
			"namelookuppriority": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of lookup (DNS or WINS) to attempt first. If the first-priority lookup fails, the second-priority lookup is attempted. Used only by the SSL VPN feature.",
			},
			"nxdomainratelimitthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Rate limit threshold for Non-Existant domain (NXDOMAIN) responses generated from Citrix ADC. Once the threshold is breached , DNS queries leading to NXDOMAIN response will be dropped. This threshold will not be applied for NXDOMAIN responses got from the backend. The threshold will be applied per packet engine and per second.",
			},
			"recursion": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Function as an end resolver and recursively resolve queries for domains that are not hosted on the Citrix ADC. Also resolve queries recursively when the external name servers configured on the appliance (for a forwarder configuration) are unavailable. When external name servers are unavailable, the appliance queries a root server and resolves the request recursively, as it does for an end resolver configuration.",
			},
			"resolutionorder": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of DNS queries (A, AAAA, or both) to generate during the routine functioning of certain Citrix ADC features, such as SSL VPN, cache redirection, and the integrated cache. The queries are sent to the external name servers that are configured for the forwarder function. If you specify both query types, you can also specify the order. Available settings function as follows:\n* OnlyAQuery. Send queries for IPv4 address records (A records) only.\n* OnlyAAAAQuery. Send queries for IPv6 address records (AAAA records) instead of queries for IPv4 address records (A records).\n* AThenAAAAQuery. Send a query for an A record, and then send a query for an AAAA record if the query for the A record results in a NODATA response from the name server.\n* AAAAThenAQuery. Send a query for an AAAA record, and then send a query for an A record if the query for the AAAA record results in a NODATA response from the name server.",
			},
			"resolvermaxactiveresolutions": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of active concurrent DNS resolutions per Packet Engine",
			},
			"resolvermaxtcpconnections": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum DNS-TCP connections opened for recursive resolution per Packet Engine",
			},
			"resolvermaxtcptimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum wait time in seconds for the response on DNS-TCP connection for recursive resolution per Packet Engine",
			},
			"retries": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of retry attempts when no response is received for a query sent to a name server. Applies to end resolver and forwarder configurations.",
			},
			"splitpktqueryprocessing": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Processing requests split across multiple packets",
			},
			"zonetransfer": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Flag to enable/disable DNS zones configuration transfer to remote GSLB site nodes",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Flag to determine if dns param is built-in or not. Possible values: [ MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL ]",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
			"nxdomainthresholdcrossed": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of times requests has been dropped as number of DNS queries leading to NXDOMAIN response has crossed the threshold.",
			},
		},
	}
}

// dnsparameterDataSourceSetAttrFromGet projects a NITRO dnsparameter GET response
// onto the data-source model. Attributes are filled from the GET (or left Null
// when the GET omits them) via the shared utils.MapGet* helpers.
func dnsparameterDataSourceSetAttrFromGet(ctx context.Context, data *DnsparameterDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In dnsparameterDataSourceSetAttrFromGet Function")

	// Read/write attributes as read-back outputs.
	data.Autosavekeyops = utils.MapGetString(g, "autosavekeyops")
	data.Cacheecszeroprefix = utils.MapGetString(g, "cacheecszeroprefix")
	data.Cachehitbypass = utils.MapGetString(g, "cachehitbypass")
	data.Cachenoexpire = utils.MapGetString(g, "cachenoexpire")
	data.Cacherecords = utils.MapGetString(g, "cacherecords")
	data.Dns64timeout = utils.MapGetInt64(g, "dns64timeout")
	data.Dnsrootreferral = utils.MapGetString(g, "dnsrootreferral")
	data.Dnssec = utils.MapGetString(g, "dnssec")
	data.Ecsmaxsubnets = utils.MapGetInt64(g, "ecsmaxsubnets")
	data.Maxcachesize = utils.MapGetInt64(g, "maxcachesize")
	data.Maxnegativecachesize = utils.MapGetInt64(g, "maxnegativecachesize")
	data.Maxnegcachettl = utils.MapGetInt64(g, "maxnegcachettl")
	data.Maxpipeline = utils.MapGetInt64(g, "maxpipeline")
	data.Maxttl = utils.MapGetInt64(g, "maxttl")
	data.Maxudppacketsize = utils.MapGetInt64(g, "maxudppacketsize")
	data.Minttl = utils.MapGetInt64(g, "minttl")
	data.Namelookuppriority = utils.MapGetString(g, "namelookuppriority")
	data.Nxdomainratelimitthreshold = utils.MapGetInt64(g, "nxdomainratelimitthreshold")
	data.Recursion = utils.MapGetString(g, "recursion")
	data.Resolutionorder = utils.MapGetString(g, "resolutionorder")
	data.Resolvermaxactiveresolutions = utils.MapGetInt64(g, "resolvermaxactiveresolutions")
	data.Resolvermaxtcpconnections = utils.MapGetInt64(g, "resolvermaxtcpconnections")
	data.Resolvermaxtcptimeout = utils.MapGetInt64(g, "resolvermaxtcptimeout")
	data.Retries = utils.MapGetInt64(g, "retries")
	data.Splitpktqueryprocessing = utils.MapGetString(g, "splitpktqueryprocessing")
	data.Zonetransfer = utils.MapGetString(g, "zonetransfer")

	// Read-only attributes.
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
	data.Nxdomainthresholdcrossed = utils.MapGetInt64(g, "nxdomainthresholdcrossed")

	// Singleton resource - static ID.
	data.Id = types.StringValue("dnsparameter-config")
}
