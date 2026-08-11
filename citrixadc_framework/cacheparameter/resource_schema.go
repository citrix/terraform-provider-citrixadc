package cacheparameter

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/cache"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// CacheparameterResourceModel describes the resource data model.
type CacheparameterResourceModel struct {
	Id                  types.String `tfsdk:"id"`
	Cacheevictionpolicy types.String `tfsdk:"cacheevictionpolicy"`
	Enablebypass        types.String `tfsdk:"enablebypass"`
	Enablehaobjpersist  types.String `tfsdk:"enablehaobjpersist"`
	Maxpostlen          types.Int64  `tfsdk:"maxpostlen"`
	Memlimit            types.Int64  `tfsdk:"memlimit"`
	Prefetchmaxpending  types.Int64  `tfsdk:"prefetchmaxpending"`
	Undefaction         types.String `tfsdk:"undefaction"`
	Verifyusing         types.String `tfsdk:"verifyusing"`
	Via                 types.String `tfsdk:"via"`
}

func (r *CacheparameterResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the cacheparameter resource.",
			},
			"cacheevictionpolicy": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("RELAXED"),
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
				Default:     int64default.StaticInt64(4096),
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

func cacheparameterGetThePayloadFromtheConfig(ctx context.Context, data *CacheparameterResourceModel) cache.Cacheparameter {
	tflog.Debug(ctx, "In cacheparameterGetThePayloadFromtheConfig Function")

	// Create API request body from the model.
	// Skip Null and Unknown values: an unconfigured Optional+Computed attribute is
	// Unknown in the create plan and must NOT be sent (it would push an empty
	// string / zero to the ADC).
	cacheparameter := cache.Cacheparameter{}
	if !data.Cacheevictionpolicy.IsNull() && !data.Cacheevictionpolicy.IsUnknown() {
		cacheparameter.Cacheevictionpolicy = data.Cacheevictionpolicy.ValueString()
	}
	if !data.Enablebypass.IsNull() && !data.Enablebypass.IsUnknown() {
		cacheparameter.Enablebypass = data.Enablebypass.ValueString()
	}
	if !data.Enablehaobjpersist.IsNull() && !data.Enablehaobjpersist.IsUnknown() {
		cacheparameter.Enablehaobjpersist = data.Enablehaobjpersist.ValueString()
	}
	if !data.Maxpostlen.IsNull() && !data.Maxpostlen.IsUnknown() {
		cacheparameter.Maxpostlen = utils.IntPtr(int(data.Maxpostlen.ValueInt64()))
	}
	if !data.Memlimit.IsNull() && !data.Memlimit.IsUnknown() {
		cacheparameter.Memlimit = utils.IntPtr(int(data.Memlimit.ValueInt64()))
	}
	if !data.Prefetchmaxpending.IsNull() && !data.Prefetchmaxpending.IsUnknown() {
		cacheparameter.Prefetchmaxpending = utils.IntPtr(int(data.Prefetchmaxpending.ValueInt64()))
	}
	if !data.Undefaction.IsNull() && !data.Undefaction.IsUnknown() {
		cacheparameter.Undefaction = data.Undefaction.ValueString()
	}
	if !data.Verifyusing.IsNull() && !data.Verifyusing.IsUnknown() {
		cacheparameter.Verifyusing = data.Verifyusing.ValueString()
	}
	if !data.Via.IsNull() && !data.Via.IsUnknown() {
		cacheparameter.Via = data.Via.ValueString()
	}

	return cacheparameter
}

// cacheparameterSetAttrFromGet maps the GET response onto the RESOURCE state.
// For each Optional+Computed attribute, when the ADC omits the field from GET
// (which happens for NITRO defaults such as 0/false/empty), we only reset it to
// Null if the model value is still Unknown (i.e. it was never configured). This
// preserves a value the user explicitly configured (e.g. maxpostlen = 0) and
// avoids "inconsistent result after apply" errors.
func cacheparameterSetAttrFromGet(ctx context.Context, data *CacheparameterResourceModel, getResponseData map[string]interface{}) *CacheparameterResourceModel {
	tflog.Debug(ctx, "In cacheparameterSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["cacheevictionpolicy"]; ok && val != nil {
		data.Cacheevictionpolicy = types.StringValue(val.(string))
	} else if data.Cacheevictionpolicy.IsUnknown() {
		data.Cacheevictionpolicy = types.StringNull()
	}
	if val, ok := getResponseData["enablebypass"]; ok && val != nil {
		data.Enablebypass = types.StringValue(val.(string))
	} else if data.Enablebypass.IsUnknown() {
		data.Enablebypass = types.StringNull()
	}
	if val, ok := getResponseData["enablehaobjpersist"]; ok && val != nil {
		data.Enablehaobjpersist = types.StringValue(val.(string))
	} else if data.Enablehaobjpersist.IsUnknown() {
		data.Enablehaobjpersist = types.StringNull()
	}
	if val, ok := getResponseData["maxpostlen"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxpostlen = types.Int64Value(intVal)
		}
	} else if data.Maxpostlen.IsUnknown() {
		data.Maxpostlen = types.Int64Null()
	}
	if val, ok := getResponseData["memlimit"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Memlimit = types.Int64Value(intVal)
		}
	} else if data.Memlimit.IsUnknown() {
		data.Memlimit = types.Int64Null()
	}
	if val, ok := getResponseData["prefetchmaxpending"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Prefetchmaxpending = types.Int64Value(intVal)
		}
	} else if data.Prefetchmaxpending.IsUnknown() {
		data.Prefetchmaxpending = types.Int64Null()
	}
	if val, ok := getResponseData["undefaction"]; ok && val != nil {
		data.Undefaction = types.StringValue(val.(string))
	} else if data.Undefaction.IsUnknown() {
		data.Undefaction = types.StringNull()
	}
	if val, ok := getResponseData["verifyusing"]; ok && val != nil {
		data.Verifyusing = types.StringValue(val.(string))
	} else if data.Verifyusing.IsUnknown() {
		data.Verifyusing = types.StringNull()
	}
	if val, ok := getResponseData["via"]; ok && val != nil {
		data.Via = types.StringValue(val.(string))
	} else if data.Via.IsUnknown() {
		data.Via = types.StringNull()
	}

	// Set ID for the resource
	// Case 1: No unique attributes (singleton) - static ID
	data.Id = types.StringValue("cacheparameter-config")

	return data
}

// cacheparameterSetAttrFromGetForDatasource maps the GET response onto the
// DATASOURCE model. A datasource has no prior state to preserve, so every
// attribute is copied straight from the GET response (value when present, Null
// otherwise) and the synthetic singleton ID is set.
func cacheparameterSetAttrFromGetForDatasource(ctx context.Context, data *CacheparameterResourceModel, getResponseData map[string]interface{}) *CacheparameterResourceModel {
	tflog.Debug(ctx, "In cacheparameterSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["cacheevictionpolicy"]; ok && val != nil {
		data.Cacheevictionpolicy = types.StringValue(val.(string))
	} else {
		data.Cacheevictionpolicy = types.StringNull()
	}
	if val, ok := getResponseData["enablebypass"]; ok && val != nil {
		data.Enablebypass = types.StringValue(val.(string))
	} else {
		data.Enablebypass = types.StringNull()
	}
	if val, ok := getResponseData["enablehaobjpersist"]; ok && val != nil {
		data.Enablehaobjpersist = types.StringValue(val.(string))
	} else {
		data.Enablehaobjpersist = types.StringNull()
	}
	if val, ok := getResponseData["maxpostlen"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxpostlen = types.Int64Value(intVal)
		} else {
			data.Maxpostlen = types.Int64Null()
		}
	} else {
		data.Maxpostlen = types.Int64Null()
	}
	if val, ok := getResponseData["memlimit"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Memlimit = types.Int64Value(intVal)
		} else {
			data.Memlimit = types.Int64Null()
		}
	} else {
		data.Memlimit = types.Int64Null()
	}
	if val, ok := getResponseData["prefetchmaxpending"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Prefetchmaxpending = types.Int64Value(intVal)
		} else {
			data.Prefetchmaxpending = types.Int64Null()
		}
	} else {
		data.Prefetchmaxpending = types.Int64Null()
	}
	if val, ok := getResponseData["undefaction"]; ok && val != nil {
		data.Undefaction = types.StringValue(val.(string))
	} else {
		data.Undefaction = types.StringNull()
	}
	if val, ok := getResponseData["verifyusing"]; ok && val != nil {
		data.Verifyusing = types.StringValue(val.(string))
	} else {
		data.Verifyusing = types.StringNull()
	}
	if val, ok := getResponseData["via"]; ok && val != nil {
		data.Via = types.StringValue(val.(string))
	} else {
		data.Via = types.StringNull()
	}

	// Set ID for the datasource
	data.Id = types.StringValue("cacheparameter-config")

	return data
}
