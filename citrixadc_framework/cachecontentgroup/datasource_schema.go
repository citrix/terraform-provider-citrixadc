package cachecontentgroup

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func CachecontentgroupDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"absexpiry": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Local time, up to 4 times a day, at which all objects in the content group must expire.\n\nCLI Users:\nFor example, to specify that the objects in the content group should expire by 11:00 PM, type the following command: add cache contentgroup <contentgroup name> -absexpiry 23:00\nTo specify that the objects in the content group should expire at 10:00 AM, 3 PM, 6 PM, and 11:00 PM, type: add cache contentgroup <contentgroup name> -absexpiry 10:00 15:00 18:00 23:00",
			},
			"absexpirygmt": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Coordinated Universal Time (GMT), up to 4 times a day, when all objects in the content group must expire.",
			},
			"alwaysevalpolicies": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Force policy evaluation for each response arriving from the origin server. Cannot be set to YES if the Prefetch parameter is also set to YES.",
			},
			"cachecontrol": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Insert a Cache-Control header into the response.",
			},
			"expireatlastbyte": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Force expiration of the content immediately after the response is downloaded (upon receipt of the last byte of the response body). Applicable only to positive responses.",
			},
			"flashcache": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Perform flash cache. Mutually exclusive with Poll Every Time (PET) on the same content group.",
			},
			"heurexpiryparam": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Heuristic expiry time, in percent of the duration, since the object was last modified.",
			},
			"hitparams": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Parameters to use for parameterized hit evaluation of an object. Up to 128 parameters can be specified. Mutually exclusive with the Hit Selector parameter.",
			},
			"hitselector": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Selector for evaluating whether an object gets stored in a particular content group. A selector is an abstraction for a collection of PIXL expressions.",
			},
			"host": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Flush only objects that belong to the specified host. Do not use except with parameterized invalidation. Also, the Invalidation Restricted to Host parameter for the group must be set to YES.",
			},
			"ignoreparamvaluecase": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Ignore case when comparing parameter values during parameterized hit evaluation. (Parameter value case is ignored by default during parameterized invalidation.)",
			},
			"ignorereloadreq": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Ignore any request to reload a cached object from the origin server.\nTo guard against Denial of Service attacks, set this parameter to YES. For RFC-compliant behavior, set it to NO.",
			},
			"ignorereqcachinghdrs": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Ignore Cache-Control and Pragma headers in the incoming request.",
			},
			"insertage": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Insert an Age header into the response. An Age header contains information about the age of the object, in seconds, as calculated by the integrated cache.",
			},
			"insertetag": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Insert an ETag header in the response. With ETag header insertion, the integrated cache does not serve full responses on repeat requests.",
			},
			"insertvia": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Insert a Via header into the response.",
			},
			"invalparams": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Parameters for parameterized invalidation of an object. You can specify up to 8 parameters. Mutually exclusive with invalSelector.",
			},
			"invalrestrictedtohost": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Take the host header into account during parameterized invalidation.",
			},
			"invalselector": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Selector for invalidating objects in the content group. A selector is an abstraction for a collection of PIXL expressions.",
			},
			"lazydnsresolve": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Perform DNS resolution for responses only if the destination IP address in the request does not match the destination IP address of the cached response.",
			},
			"matchcookies": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Evaluate for parameters in the cookie header also.",
			},
			"maxressize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum size of a response that can be cached in this content group.",
			},
			"memlimit": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum amount of memory that the cache can use. The effective limit is based on the available memory of the Citrix ADC.",
			},
			"minhits": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of hits that qualifies a response for storage in this content group.",
			},
			"minressize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum size of a response that can be cached in this content group.\n Default minimum response size is 0.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the content group.  Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the content group is created.",
			},
			"persistha": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Setting persistHA to YES causes IC to save objects in contentgroup to Secondary node in HA deployment.",
			},
			"pinned": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Do not flush objects from this content group under memory pressure.",
			},
			"polleverytime": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Always poll for the objects in this content group. That is, retrieve the objects from the origin server whenever they are requested.",
			},
			"prefetch": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Attempt to refresh objects that are about to go stale.",
			},
			"prefetchmaxpending": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of outstanding prefetches that can be queued for the content group.",
			},
			"prefetchperiod": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time period, in seconds before an object's calculated expiry time, during which to attempt prefetch.",
			},
			"prefetchperiodmillisec": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time period, in milliseconds before an object's calculated expiry time, during which to attempt prefetch.",
			},
			"query": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Query string specifying individual objects to flush from this group by using parameterized invalidation. If this parameter is not set, all objects are flushed from the group.",
			},
			"quickabortsize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "If the size of an object that is being downloaded is less than or equal to the quick abort value, and a client aborts during the download, the cache stops downloading the response. If the object is larger than the quick abort size, the cache continues to download the response.",
			},
			"relexpiry": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Relative expiry time, in seconds, after which to expire an object cached in this content group.",
			},
			"relexpirymillisec": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Relative expiry time, in milliseconds, after which to expire an object cached in this content group.",
			},
			"removecookies": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Remove cookies from responses.",
			},
			"selectorvalue": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Value of the selector to be used for flushing objects from the content group. Requires that an invalidation selector be configured for the content group.",
			},
			"tosecondary": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "content group whose objects are to be sent to secondary.",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The type of the content group.",
			},
			"weaknegrelexpiry": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Relative expiry time, in seconds, for expiring negative responses. This value is used only if the expiry time cannot be determined from any other source. It is applicable only to the following status codes: 307, 403, 404, and 410.",
			},
			"weakposrelexpiry": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Relative expiry time, in seconds, for expiring positive responses with response codes between 200 and 399. Cannot be used in combination with other Expiry attributes. Similar to -relExpiry but has lower precedence.",
			},
		},
	}
}
