package cacheparameter

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func CacheparameterDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"cacheevictionpolicy": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The cacheEvictionPolicy determines the threshold for preemptive eviction of cache objects using the LRU (Least Recently Used) algorithm. If set to AGGRESSIVE, eviction is triggered when free cache memory drops to 40%. MODERATE triggers eviction at 25%, and RELAXED triggers eviction at 10%.",
			},
			"enablebypass": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Evaluate the request-time policies before attempting hit selection. If set to NO, an incoming request for which a matching object is found in cache storage results in a response regardless of the policy configuration.\nIf the request matches a policy with a NOCACHE action, the request bypasses all cache processing.\nThis parameter does not affect processing of requests that match any invalidation policy.",
			},
			"enablehaobjpersist": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The HA object persisting parameter. When this value is set to YES, cache objects can be synced to Secondary in a HA deployment.  If set to NO, objects will never be synced to Secondary node.",
			},
			"maxpostlen": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of POST body bytes to consider when evaluating parameters for a content group for which you have configured hit parameters and invalidation parameters.",
			},
			"memlimit": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Amount of memory available for storing the cache objects. In practice, the amount of memory available for caching can be less than half the total memory of the Citrix ADC.",
			},
			"prefetchmaxpending": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of outstanding prefetches in the Integrated Cache.",
			},
			"undefaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Action to take when a policy cannot be evaluated.",
			},
			"verifyusing": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Criteria for deciding whether a cached object can be served for an incoming HTTP request. Available settings function as follows:\nHOSTNAME - The URL, host name, and host port values in the incoming HTTP request header must match the cache policy. The IP address and the TCP port of the destination host are not evaluated. Do not use the HOSTNAME setting unless you are certain that no rogue client can access a rogue server through the cache.\nHOSTNAME_AND_IP - The URL, host name, host port in the incoming HTTP request header, and the IP address and TCP port of\nthe destination server, must match the cache policy.\nDNS - The URL, host name and host port in the incoming HTTP request, and the TCP port must match the cache policy. The host name is used for DNS lookup of the destination server's IP address, and is compared with the set of addresses returned by the DNS lookup.",
			},
			"via": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "String to include in the Via header. A Via header is inserted into all responses served from a content group if its Insert Via flag is set.",
			},
		},
	}
}
