package cachecontentgroup

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// CachecontentgroupDataSourceModel is the data-source-specific model, decoupled
// from CachecontentgroupResourceModel. A data source is a pure read surface, so
// it can expose the full GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits.
type CachecontentgroupDataSourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Absexpiry              types.List   `tfsdk:"absexpiry"`
	Absexpirygmt           types.List   `tfsdk:"absexpirygmt"`
	Alwaysevalpolicies     types.String `tfsdk:"alwaysevalpolicies"`
	Cachecontrol           types.String `tfsdk:"cachecontrol"`
	Expireatlastbyte       types.String `tfsdk:"expireatlastbyte"`
	Flashcache             types.String `tfsdk:"flashcache"`
	Heurexpiryparam        types.Int64  `tfsdk:"heurexpiryparam"`
	Hitparams              types.List   `tfsdk:"hitparams"`
	Hitselector            types.String `tfsdk:"hitselector"`
	Host                   types.String `tfsdk:"host"`
	Ignoreparamvaluecase   types.String `tfsdk:"ignoreparamvaluecase"`
	Ignorereloadreq        types.String `tfsdk:"ignorereloadreq"`
	Ignorereqcachinghdrs   types.String `tfsdk:"ignorereqcachinghdrs"`
	Insertage              types.String `tfsdk:"insertage"`
	Insertetag             types.String `tfsdk:"insertetag"`
	Insertvia              types.String `tfsdk:"insertvia"`
	Invalparams            types.List   `tfsdk:"invalparams"`
	Invalrestrictedtohost  types.String `tfsdk:"invalrestrictedtohost"`
	Invalselector          types.String `tfsdk:"invalselector"`
	Lazydnsresolve         types.String `tfsdk:"lazydnsresolve"`
	Matchcookies           types.String `tfsdk:"matchcookies"`
	Maxressize             types.Int64  `tfsdk:"maxressize"`
	Memlimit               types.Int64  `tfsdk:"memlimit"`
	Minhits                types.Int64  `tfsdk:"minhits"`
	Minressize             types.Int64  `tfsdk:"minressize"`
	Name                   types.String `tfsdk:"name"`
	Persistha              types.String `tfsdk:"persistha"`
	Pinned                 types.String `tfsdk:"pinned"`
	Polleverytime          types.String `tfsdk:"polleverytime"`
	Prefetch               types.String `tfsdk:"prefetch"`
	Prefetchmaxpending     types.Int64  `tfsdk:"prefetchmaxpending"`
	Prefetchperiod         types.Int64  `tfsdk:"prefetchperiod"`
	Prefetchperiodmillisec types.Int64  `tfsdk:"prefetchperiodmillisec"`
	Query                  types.String `tfsdk:"query"`
	Quickabortsize         types.Int64  `tfsdk:"quickabortsize"`
	Relexpiry              types.Int64  `tfsdk:"relexpiry"`
	Relexpirymillisec      types.Int64  `tfsdk:"relexpirymillisec"`
	Removecookies          types.String `tfsdk:"removecookies"`
	Selectorvalue          types.String `tfsdk:"selectorvalue"`
	Tosecondary            types.String `tfsdk:"tosecondary"`
	Type                   types.String `tfsdk:"type"`
	Weaknegrelexpiry       types.Int64  `tfsdk:"weaknegrelexpiry"`
	Weakposrelexpiry       types.Int64  `tfsdk:"weakposrelexpiry"`

	// Read-only (GET-only) metadata from the NITRO read-only set
	// (zion73x_readonly/cachecontentgroup.json). Never settable; from GET.
	Flags                 types.Int64  `tfsdk:"flags"`
	Prefetchcur           types.Int64  `tfsdk:"prefetchcur"`
	Memusage              types.Int64  `tfsdk:"memusage"`
	Memdusage             types.Int64  `tfsdk:"memdusage"`
	Disklimit             types.Int64  `tfsdk:"disklimit"`
	Cachenon304hits       types.Int64  `tfsdk:"cachenon304hits"`
	Cache304hits          types.Int64  `tfsdk:"cache304hits"`
	Cachecells            types.Int64  `tfsdk:"cachecells"`
	Cachegroupincarnation types.Int64  `tfsdk:"cachegroupincarnation"`
	Persist               types.String `tfsdk:"persist"`
	Policyname            types.List   `tfsdk:"policyname"`
	Cachenuminvalpolicy   types.Int64  `tfsdk:"cachenuminvalpolicy"`
	Markercells           types.Int64  `tfsdk:"markercells"`
	Builtin               types.List   `tfsdk:"builtin"`
	Feature               types.String `tfsdk:"feature"`
}

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

			// Read-only (GET-only) metadata surfaced by the data source
			// (intentionally NOT modeled on the resource). All Computed.
			"flags": schema.Int64Attribute{
				Computed:    true,
				Description: "Flags.",
			},
			"prefetchcur": schema.Int64Attribute{
				Computed:    true,
				Description: "Current outstanding prefetches.",
			},
			"memusage": schema.Int64Attribute{
				Computed:    true,
				Description: "Current memory usage.",
			},
			"memdusage": schema.Int64Attribute{
				Computed:    true,
				Description: "Current disk memory usage.",
			},
			"disklimit": schema.Int64Attribute{
				Computed:    true,
				Description: "Maximum amount of disk that the cache can use. The effective limit is based on the available memory of the Citrix ADC.",
			},
			"cachenon304hits": schema.Int64Attribute{
				Computed:    true,
				Description: "Cache non 304 hits.",
			},
			"cache304hits": schema.Int64Attribute{
				Computed:    true,
				Description: "Cache 304 hits.",
			},
			"cachecells": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of cells.",
			},
			"cachegroupincarnation": schema.Int64Attribute{
				Computed:    true,
				Description: "Cache group incarnation.",
			},
			"persist": schema.StringAttribute{
				Computed:    true,
				Description: "Setting persist to YES causes IC to save objects in contentgroup to disk. Possible values = YES, NO",
			},
			"policyname": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Active cache policies referring to this group.",
			},
			"cachenuminvalpolicy": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of active Invalidation policies referring to this group.",
			},
			"markercells": schema.Int64Attribute{
				Computed:    true,
				Description: "Numbers of marker cells in this group.",
			},
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Flag to determine whether the content group is built-in. Possible values = MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
		},
	}
}

// cachecontentgroupDataSourceSetAttrFromGet projects a NITRO cachecontentgroup
// GET response onto the data-source model. Attributes are filled from the GET
// (or left Null when the GET omits them) via the shared utils.MapGet* helpers.
func cachecontentgroupDataSourceSetAttrFromGet(ctx context.Context, data *CachecontentgroupDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In cachecontentgroupDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Absexpiry = utils.MapGetStringList(g, "absexpiry")
	data.Absexpirygmt = utils.MapGetStringList(g, "absexpirygmt")
	data.Alwaysevalpolicies = utils.MapGetString(g, "alwaysevalpolicies")
	data.Cachecontrol = utils.MapGetString(g, "cachecontrol")
	data.Expireatlastbyte = utils.MapGetString(g, "expireatlastbyte")
	data.Flashcache = utils.MapGetString(g, "flashcache")
	data.Heurexpiryparam = utils.MapGetInt64(g, "heurexpiryparam")
	data.Hitparams = utils.MapGetStringList(g, "hitparams")
	data.Hitselector = utils.MapGetString(g, "hitselector")
	data.Host = utils.MapGetString(g, "host")
	data.Ignoreparamvaluecase = utils.MapGetString(g, "ignoreparamvaluecase")
	data.Ignorereloadreq = utils.MapGetString(g, "ignorereloadreq")
	data.Ignorereqcachinghdrs = utils.MapGetString(g, "ignorereqcachinghdrs")
	data.Insertage = utils.MapGetString(g, "insertage")
	data.Insertetag = utils.MapGetString(g, "insertetag")
	data.Insertvia = utils.MapGetString(g, "insertvia")
	data.Invalparams = utils.MapGetStringList(g, "invalparams")
	data.Invalrestrictedtohost = utils.MapGetString(g, "invalrestrictedtohost")
	data.Invalselector = utils.MapGetString(g, "invalselector")
	data.Lazydnsresolve = utils.MapGetString(g, "lazydnsresolve")
	data.Matchcookies = utils.MapGetString(g, "matchcookies")
	data.Maxressize = utils.MapGetInt64(g, "maxressize")
	data.Memlimit = utils.MapGetInt64(g, "memlimit")
	data.Minhits = utils.MapGetInt64(g, "minhits")
	data.Minressize = utils.MapGetInt64(g, "minressize")
	data.Persistha = utils.MapGetString(g, "persistha")
	data.Pinned = utils.MapGetString(g, "pinned")
	data.Polleverytime = utils.MapGetString(g, "polleverytime")
	data.Prefetch = utils.MapGetString(g, "prefetch")
	data.Prefetchmaxpending = utils.MapGetInt64(g, "prefetchmaxpending")
	data.Prefetchperiod = utils.MapGetInt64(g, "prefetchperiod")
	data.Prefetchperiodmillisec = utils.MapGetInt64(g, "prefetchperiodmillisec")
	data.Query = utils.MapGetString(g, "query")
	data.Quickabortsize = utils.MapGetInt64(g, "quickabortsize")
	data.Relexpiry = utils.MapGetInt64(g, "relexpiry")
	data.Relexpirymillisec = utils.MapGetInt64(g, "relexpirymillisec")
	data.Removecookies = utils.MapGetString(g, "removecookies")
	data.Selectorvalue = utils.MapGetString(g, "selectorvalue")
	data.Tosecondary = utils.MapGetString(g, "tosecondary")
	data.Type = utils.MapGetString(g, "type")
	data.Weaknegrelexpiry = utils.MapGetInt64(g, "weaknegrelexpiry")
	data.Weakposrelexpiry = utils.MapGetInt64(g, "weakposrelexpiry")

	// Read-only metadata.
	data.Flags = utils.MapGetInt64(g, "flags")
	data.Prefetchcur = utils.MapGetInt64(g, "prefetchcur")
	data.Memusage = utils.MapGetInt64(g, "memusage")
	data.Memdusage = utils.MapGetInt64(g, "memdusage")
	data.Disklimit = utils.MapGetInt64(g, "disklimit")
	data.Cachenon304hits = utils.MapGetInt64(g, "cachenon304hits")
	data.Cache304hits = utils.MapGetInt64(g, "cache304hits")
	data.Cachecells = utils.MapGetInt64(g, "cachecells")
	data.Cachegroupincarnation = utils.MapGetInt64(g, "cachegroupincarnation")
	data.Persist = utils.MapGetString(g, "persist")
	data.Policyname = utils.MapGetStringList(g, "policyname")
	data.Cachenuminvalpolicy = utils.MapGetInt64(g, "cachenuminvalpolicy")
	data.Markercells = utils.MapGetInt64(g, "markercells")
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
}
