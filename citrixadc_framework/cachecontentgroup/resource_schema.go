package cachecontentgroup

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/cache"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// CachecontentgroupResourceModel describes the resource data model.
type CachecontentgroupResourceModel struct {
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
}

func (r *CachecontentgroupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the cachecontentgroup resource.",
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
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Flush only objects that belong to the specified host. Do not use except with parameterized invalidation. Also, the Invalidation Restricted to Host parameter for the group must be set to YES.",
			},
			"ignoreparamvaluecase": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Ignore case when comparing parameter values during parameterized hit evaluation. (Parameter value case is ignored by default during parameterized invalidation.)",
			},
			"ignorereloadreq": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("True"),
				Description: "Ignore any request to reload a cached object from the origin server.\nTo guard against Denial of Service attacks, set this parameter to YES. For RFC-compliant behavior, set it to NO.",
			},
			"ignorereqcachinghdrs": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("True"),
				Description: "Ignore Cache-Control and Pragma headers in the incoming request.",
			},
			"insertage": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("True"),
				Description: "Insert an Age header into the response. An Age header contains information about the age of the object, in seconds, as calculated by the integrated cache.",
			},
			"insertetag": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("True"),
				Description: "Insert an ETag header in the response. With ETag header insertion, the integrated cache does not serve full responses on repeat requests.",
			},
			"insertvia": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("True"),
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
				Default:     stringdefault.StaticString("True"),
				Description: "Perform DNS resolution for responses only if the destination IP address in the request does not match the destination IP address of the cached response.",
			},
			"matchcookies": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Evaluate for parameters in the cookie header also.",
			},
			"maxressize": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(80),
				Description: "Maximum size of a response that can be cached in this content group.",
			},
			"memlimit": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(65536),
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
				Default:     stringdefault.StaticString("True"),
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
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Query string specifying individual objects to flush from this group by using parameterized invalidation. If this parameter is not set, all objects are flushed from the group.",
			},
			"quickabortsize": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(4194303),
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
				Default:     stringdefault.StaticString("True"),
				Description: "Remove cookies from responses.",
			},
			"selectorvalue": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Value of the selector to be used for flushing objects from the content group. Requires that an invalidation selector be configured for the content group.",
			},
			"tosecondary": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "content group whose objects are to be sent to secondary.",
			},
			"type": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("HTTP"),
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

func cachecontentgroupGetThePayloadFromtheConfig(ctx context.Context, data *CachecontentgroupResourceModel) cache.Cachecontentgroup {
	tflog.Debug(ctx, "In cachecontentgroupGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	cachecontentgroup := cache.Cachecontentgroup{}
	if !data.Alwaysevalpolicies.IsNull() {
		cachecontentgroup.Alwaysevalpolicies = data.Alwaysevalpolicies.ValueString()
	}
	if !data.Cachecontrol.IsNull() {
		cachecontentgroup.Cachecontrol = data.Cachecontrol.ValueString()
	}
	if !data.Expireatlastbyte.IsNull() {
		cachecontentgroup.Expireatlastbyte = data.Expireatlastbyte.ValueString()
	}
	if !data.Flashcache.IsNull() {
		cachecontentgroup.Flashcache = data.Flashcache.ValueString()
	}
	if !data.Heurexpiryparam.IsNull() {
		cachecontentgroup.Heurexpiryparam = utils.IntPtr(int(data.Heurexpiryparam.ValueInt64()))
	}
	if !data.Hitselector.IsNull() {
		cachecontentgroup.Hitselector = data.Hitselector.ValueString()
	}
	if !data.Host.IsNull() {
		cachecontentgroup.Host = data.Host.ValueString()
	}
	if !data.Ignoreparamvaluecase.IsNull() {
		cachecontentgroup.Ignoreparamvaluecase = data.Ignoreparamvaluecase.ValueString()
	}
	if !data.Ignorereloadreq.IsNull() {
		cachecontentgroup.Ignorereloadreq = data.Ignorereloadreq.ValueString()
	}
	if !data.Ignorereqcachinghdrs.IsNull() {
		cachecontentgroup.Ignorereqcachinghdrs = data.Ignorereqcachinghdrs.ValueString()
	}
	if !data.Insertage.IsNull() {
		cachecontentgroup.Insertage = data.Insertage.ValueString()
	}
	if !data.Insertetag.IsNull() {
		cachecontentgroup.Insertetag = data.Insertetag.ValueString()
	}
	if !data.Insertvia.IsNull() {
		cachecontentgroup.Insertvia = data.Insertvia.ValueString()
	}
	if !data.Invalrestrictedtohost.IsNull() {
		cachecontentgroup.Invalrestrictedtohost = data.Invalrestrictedtohost.ValueString()
	}
	if !data.Invalselector.IsNull() {
		cachecontentgroup.Invalselector = data.Invalselector.ValueString()
	}
	if !data.Lazydnsresolve.IsNull() {
		cachecontentgroup.Lazydnsresolve = data.Lazydnsresolve.ValueString()
	}
	if !data.Matchcookies.IsNull() {
		cachecontentgroup.Matchcookies = data.Matchcookies.ValueString()
	}
	if !data.Maxressize.IsNull() {
		cachecontentgroup.Maxressize = utils.IntPtr(int(data.Maxressize.ValueInt64()))
	}
	if !data.Memlimit.IsNull() {
		cachecontentgroup.Memlimit = utils.IntPtr(int(data.Memlimit.ValueInt64()))
	}
	if !data.Minhits.IsNull() {
		cachecontentgroup.Minhits = utils.IntPtr(int(data.Minhits.ValueInt64()))
	}
	if !data.Minressize.IsNull() {
		cachecontentgroup.Minressize = utils.IntPtr(int(data.Minressize.ValueInt64()))
	}
	if !data.Name.IsNull() {
		cachecontentgroup.Name = data.Name.ValueString()
	}
	if !data.Persistha.IsNull() {
		cachecontentgroup.Persistha = data.Persistha.ValueString()
	}
	if !data.Pinned.IsNull() {
		cachecontentgroup.Pinned = data.Pinned.ValueString()
	}
	if !data.Polleverytime.IsNull() {
		cachecontentgroup.Polleverytime = data.Polleverytime.ValueString()
	}
	if !data.Prefetch.IsNull() {
		cachecontentgroup.Prefetch = data.Prefetch.ValueString()
	}
	if !data.Prefetchmaxpending.IsNull() {
		cachecontentgroup.Prefetchmaxpending = utils.IntPtr(int(data.Prefetchmaxpending.ValueInt64()))
	}
	if !data.Prefetchperiod.IsNull() {
		cachecontentgroup.Prefetchperiod = utils.IntPtr(int(data.Prefetchperiod.ValueInt64()))
	}
	if !data.Prefetchperiodmillisec.IsNull() {
		cachecontentgroup.Prefetchperiodmillisec = utils.IntPtr(int(data.Prefetchperiodmillisec.ValueInt64()))
	}
	if !data.Query.IsNull() {
		cachecontentgroup.Query = data.Query.ValueString()
	}
	if !data.Quickabortsize.IsNull() {
		cachecontentgroup.Quickabortsize = utils.IntPtr(int(data.Quickabortsize.ValueInt64()))
	}
	if !data.Relexpiry.IsNull() {
		cachecontentgroup.Relexpiry = utils.IntPtr(int(data.Relexpiry.ValueInt64()))
	}
	if !data.Relexpirymillisec.IsNull() {
		cachecontentgroup.Relexpirymillisec = utils.IntPtr(int(data.Relexpirymillisec.ValueInt64()))
	}
	if !data.Removecookies.IsNull() {
		cachecontentgroup.Removecookies = data.Removecookies.ValueString()
	}
	if !data.Selectorvalue.IsNull() {
		cachecontentgroup.Selectorvalue = data.Selectorvalue.ValueString()
	}
	if !data.Tosecondary.IsNull() {
		cachecontentgroup.Tosecondary = data.Tosecondary.ValueString()
	}
	if !data.Type.IsNull() {
		cachecontentgroup.Type = data.Type.ValueString()
	}
	if !data.Weaknegrelexpiry.IsNull() {
		cachecontentgroup.Weaknegrelexpiry = utils.IntPtr(int(data.Weaknegrelexpiry.ValueInt64()))
	}
	if !data.Weakposrelexpiry.IsNull() {
		cachecontentgroup.Weakposrelexpiry = utils.IntPtr(int(data.Weakposrelexpiry.ValueInt64()))
	}

	return cachecontentgroup
}

func cachecontentgroupSetAttrFromGet(ctx context.Context, data *CachecontentgroupResourceModel, getResponseData map[string]interface{}) *CachecontentgroupResourceModel {
	tflog.Debug(ctx, "In cachecontentgroupSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["alwaysevalpolicies"]; ok && val != nil {
		data.Alwaysevalpolicies = types.StringValue(val.(string))
	} else {
		data.Alwaysevalpolicies = types.StringNull()
	}
	if val, ok := getResponseData["cachecontrol"]; ok && val != nil {
		data.Cachecontrol = types.StringValue(val.(string))
	} else {
		data.Cachecontrol = types.StringNull()
	}
	if val, ok := getResponseData["expireatlastbyte"]; ok && val != nil {
		data.Expireatlastbyte = types.StringValue(val.(string))
	} else {
		data.Expireatlastbyte = types.StringNull()
	}
	if val, ok := getResponseData["flashcache"]; ok && val != nil {
		data.Flashcache = types.StringValue(val.(string))
	} else {
		data.Flashcache = types.StringNull()
	}
	if val, ok := getResponseData["heurexpiryparam"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Heurexpiryparam = types.Int64Value(intVal)
		}
	} else {
		data.Heurexpiryparam = types.Int64Null()
	}
	if val, ok := getResponseData["hitselector"]; ok && val != nil {
		data.Hitselector = types.StringValue(val.(string))
	} else {
		data.Hitselector = types.StringNull()
	}
	if val, ok := getResponseData["host"]; ok && val != nil {
		data.Host = types.StringValue(val.(string))
	} else {
		data.Host = types.StringNull()
	}
	if val, ok := getResponseData["ignoreparamvaluecase"]; ok && val != nil {
		data.Ignoreparamvaluecase = types.StringValue(val.(string))
	} else {
		data.Ignoreparamvaluecase = types.StringNull()
	}
	if val, ok := getResponseData["ignorereloadreq"]; ok && val != nil {
		data.Ignorereloadreq = types.StringValue(val.(string))
	} else {
		data.Ignorereloadreq = types.StringNull()
	}
	if val, ok := getResponseData["ignorereqcachinghdrs"]; ok && val != nil {
		data.Ignorereqcachinghdrs = types.StringValue(val.(string))
	} else {
		data.Ignorereqcachinghdrs = types.StringNull()
	}
	if val, ok := getResponseData["insertage"]; ok && val != nil {
		data.Insertage = types.StringValue(val.(string))
	} else {
		data.Insertage = types.StringNull()
	}
	if val, ok := getResponseData["insertetag"]; ok && val != nil {
		data.Insertetag = types.StringValue(val.(string))
	} else {
		data.Insertetag = types.StringNull()
	}
	if val, ok := getResponseData["insertvia"]; ok && val != nil {
		data.Insertvia = types.StringValue(val.(string))
	} else {
		data.Insertvia = types.StringNull()
	}
	if val, ok := getResponseData["invalrestrictedtohost"]; ok && val != nil {
		data.Invalrestrictedtohost = types.StringValue(val.(string))
	} else {
		data.Invalrestrictedtohost = types.StringNull()
	}
	if val, ok := getResponseData["invalselector"]; ok && val != nil {
		data.Invalselector = types.StringValue(val.(string))
	} else {
		data.Invalselector = types.StringNull()
	}
	if val, ok := getResponseData["lazydnsresolve"]; ok && val != nil {
		data.Lazydnsresolve = types.StringValue(val.(string))
	} else {
		data.Lazydnsresolve = types.StringNull()
	}
	if val, ok := getResponseData["matchcookies"]; ok && val != nil {
		data.Matchcookies = types.StringValue(val.(string))
	} else {
		data.Matchcookies = types.StringNull()
	}
	if val, ok := getResponseData["maxressize"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxressize = types.Int64Value(intVal)
		}
	} else {
		data.Maxressize = types.Int64Null()
	}
	if val, ok := getResponseData["memlimit"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Memlimit = types.Int64Value(intVal)
		}
	} else {
		data.Memlimit = types.Int64Null()
	}
	if val, ok := getResponseData["minhits"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Minhits = types.Int64Value(intVal)
		}
	} else {
		data.Minhits = types.Int64Null()
	}
	if val, ok := getResponseData["minressize"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Minressize = types.Int64Value(intVal)
		}
	} else {
		data.Minressize = types.Int64Null()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["persistha"]; ok && val != nil {
		data.Persistha = types.StringValue(val.(string))
	} else {
		data.Persistha = types.StringNull()
	}
	if val, ok := getResponseData["pinned"]; ok && val != nil {
		data.Pinned = types.StringValue(val.(string))
	} else {
		data.Pinned = types.StringNull()
	}
	if val, ok := getResponseData["polleverytime"]; ok && val != nil {
		data.Polleverytime = types.StringValue(val.(string))
	} else {
		data.Polleverytime = types.StringNull()
	}
	if val, ok := getResponseData["prefetch"]; ok && val != nil {
		data.Prefetch = types.StringValue(val.(string))
	} else {
		data.Prefetch = types.StringNull()
	}
	if val, ok := getResponseData["prefetchmaxpending"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Prefetchmaxpending = types.Int64Value(intVal)
		}
	} else {
		data.Prefetchmaxpending = types.Int64Null()
	}
	if val, ok := getResponseData["prefetchperiod"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Prefetchperiod = types.Int64Value(intVal)
		}
	} else {
		data.Prefetchperiod = types.Int64Null()
	}
	if val, ok := getResponseData["prefetchperiodmillisec"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Prefetchperiodmillisec = types.Int64Value(intVal)
		}
	} else {
		data.Prefetchperiodmillisec = types.Int64Null()
	}
	if val, ok := getResponseData["query"]; ok && val != nil {
		data.Query = types.StringValue(val.(string))
	} else {
		data.Query = types.StringNull()
	}
	if val, ok := getResponseData["quickabortsize"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Quickabortsize = types.Int64Value(intVal)
		}
	} else {
		data.Quickabortsize = types.Int64Null()
	}
	if val, ok := getResponseData["relexpiry"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Relexpiry = types.Int64Value(intVal)
		}
	} else {
		data.Relexpiry = types.Int64Null()
	}
	if val, ok := getResponseData["relexpirymillisec"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Relexpirymillisec = types.Int64Value(intVal)
		}
	} else {
		data.Relexpirymillisec = types.Int64Null()
	}
	if val, ok := getResponseData["removecookies"]; ok && val != nil {
		data.Removecookies = types.StringValue(val.(string))
	} else {
		data.Removecookies = types.StringNull()
	}
	if val, ok := getResponseData["selectorvalue"]; ok && val != nil {
		data.Selectorvalue = types.StringValue(val.(string))
	} else {
		data.Selectorvalue = types.StringNull()
	}
	if val, ok := getResponseData["tosecondary"]; ok && val != nil {
		data.Tosecondary = types.StringValue(val.(string))
	} else {
		data.Tosecondary = types.StringNull()
	}
	if val, ok := getResponseData["type"]; ok && val != nil {
		data.Type = types.StringValue(val.(string))
	} else {
		data.Type = types.StringNull()
	}
	if val, ok := getResponseData["weaknegrelexpiry"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Weaknegrelexpiry = types.Int64Value(intVal)
		}
	} else {
		data.Weaknegrelexpiry = types.Int64Null()
	}
	if val, ok := getResponseData["weakposrelexpiry"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Weakposrelexpiry = types.Int64Value(intVal)
		}
	} else {
		data.Weakposrelexpiry = types.Int64Null()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
