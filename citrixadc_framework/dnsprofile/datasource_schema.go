package dnsprofile

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func DnsprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"cacheecsresponses": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Cache DNS responses with EDNS Client Subnet(ECS) option in the DNS cache. When disabled, the appliance stops caching responses with ECS option. This is relevant to proxy configuration. Enabling/disabling support of ECS option when Citrix ADC is authoritative for a GSLB domain is supported using a knob in GSLB vserver. In all other modes, ECS option is ignored.",
			},
			"cachenegativeresponses": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Cache negative responses in the DNS cache. When disabled, the appliance stops caching negative responses except referral records. This applies to all configurations - proxy, end resolver, and forwarder. However, cached responses are not flushed. The appliance does not serve negative responses from the cache until this parameter is enabled again.",
			},
			"cacherecords": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Cache resource records in the DNS cache. Applies to resource records obtained through proxy configurations only. End resolver and forwarder configurations always cache records in the DNS cache, and you cannot disable this behavior. When you disable record caching, the appliance stops caching server responses. However, cached records are not flushed. The appliance does not serve requests from the cache until record caching is enabled again.",
			},
			"dnsanswerseclogging": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "DNS answer section; if enabled, answer section in the response will be logged.",
			},
			"dnserrorlogging": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "DNS error logging; if enabled, whenever error is encountered in DNS module reason for the error will be logged.",
			},
			"dnsextendedlogging": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "DNS extended logging; if enabled, authority and additional section in the response will be logged.",
			},
			"dnsprofilename": schema.StringAttribute{
				Required:    true,
				Description: "Name of the DNS profile",
			},
			"dnsquerylogging": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "DNS query logging; if enabled, DNS query information such as DNS query id, DNS query flags , DNS domain name and DNS query type will be logged",
			},
			"dropmultiqueryrequest": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Drop the DNS requests containing multiple queries. When enabled, DNS requests containing multiple queries will be dropped. In case of proxy configuration by default the DNS request containing multiple queries is forwarded to the backend and in case of ADNS and Resolver configuration NOCODE error response will be sent to the client.",
			},
			"insertecs": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Insert ECS Option on DNS query",
			},
			"maxcacheableecsprefixlength": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The maximum ecs prefix length that will be cached",
			},
			"maxcacheableecsprefixlength6": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The maximum ecs prefix length that will be cached for IPv6 subnets",
			},
			"recursiveresolution": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "DNS recursive resolution; if enabled, will do recursive resolution for DNS query when the profile is associated with ADNS service, CS Vserver and DNS action",
			},
			"replaceecs": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Replace ECS Option on DNS query",
			},
		},
	}
}
