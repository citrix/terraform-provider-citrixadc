package dnsprofile

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/dns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// DnsprofileResourceModel describes the resource data model.
type DnsprofileResourceModel struct {
	Id                           types.String `tfsdk:"id"`
	Cacheecsresponses            types.String `tfsdk:"cacheecsresponses"`
	Cachenegativeresponses       types.String `tfsdk:"cachenegativeresponses"`
	Cacherecords                 types.String `tfsdk:"cacherecords"`
	Dnsanswerseclogging          types.String `tfsdk:"dnsanswerseclogging"`
	Dnserrorlogging              types.String `tfsdk:"dnserrorlogging"`
	Dnsextendedlogging           types.String `tfsdk:"dnsextendedlogging"`
	Dnsprofilename               types.String `tfsdk:"dnsprofilename"`
	Dnsquerylogging              types.String `tfsdk:"dnsquerylogging"`
	Dropmultiqueryrequest        types.String `tfsdk:"dropmultiqueryrequest"`
	Insertecs                    types.String `tfsdk:"insertecs"`
	Maxcacheableecsprefixlength  types.Int64  `tfsdk:"maxcacheableecsprefixlength"`
	Maxcacheableecsprefixlength6 types.Int64  `tfsdk:"maxcacheableecsprefixlength6"`
	Recursiveresolution          types.String `tfsdk:"recursiveresolution"`
	Replaceecs                   types.String `tfsdk:"replaceecs"`
}

func (r *DnsprofileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the dnsprofile resource.",
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
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
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

func dnsprofileGetThePayloadFromthePlan(ctx context.Context, data *DnsprofileResourceModel) dns.Dnsprofile {
	tflog.Debug(ctx, "In dnsprofileGetThePayloadFromthePlan Function")

	// Create API request body from the model
	dnsprofile := dns.Dnsprofile{}
	if !data.Cacheecsresponses.IsNull() && !data.Cacheecsresponses.IsUnknown() {
		dnsprofile.Cacheecsresponses = data.Cacheecsresponses.ValueString()
	}
	if !data.Cachenegativeresponses.IsNull() && !data.Cachenegativeresponses.IsUnknown() {
		dnsprofile.Cachenegativeresponses = data.Cachenegativeresponses.ValueString()
	}
	if !data.Cacherecords.IsNull() && !data.Cacherecords.IsUnknown() {
		dnsprofile.Cacherecords = data.Cacherecords.ValueString()
	}
	if !data.Dnsanswerseclogging.IsNull() && !data.Dnsanswerseclogging.IsUnknown() {
		dnsprofile.Dnsanswerseclogging = data.Dnsanswerseclogging.ValueString()
	}
	if !data.Dnserrorlogging.IsNull() && !data.Dnserrorlogging.IsUnknown() {
		dnsprofile.Dnserrorlogging = data.Dnserrorlogging.ValueString()
	}
	if !data.Dnsextendedlogging.IsNull() && !data.Dnsextendedlogging.IsUnknown() {
		dnsprofile.Dnsextendedlogging = data.Dnsextendedlogging.ValueString()
	}
	if !data.Dnsprofilename.IsNull() && !data.Dnsprofilename.IsUnknown() {
		dnsprofile.Dnsprofilename = data.Dnsprofilename.ValueString()
	}
	if !data.Dnsquerylogging.IsNull() && !data.Dnsquerylogging.IsUnknown() {
		dnsprofile.Dnsquerylogging = data.Dnsquerylogging.ValueString()
	}
	if !data.Dropmultiqueryrequest.IsNull() && !data.Dropmultiqueryrequest.IsUnknown() {
		dnsprofile.Dropmultiqueryrequest = data.Dropmultiqueryrequest.ValueString()
	}
	if !data.Insertecs.IsNull() && !data.Insertecs.IsUnknown() {
		dnsprofile.Insertecs = data.Insertecs.ValueString()
	}
	if !data.Maxcacheableecsprefixlength.IsNull() && !data.Maxcacheableecsprefixlength.IsUnknown() {
		dnsprofile.Maxcacheableecsprefixlength = utils.IntPtr(int(data.Maxcacheableecsprefixlength.ValueInt64()))
	}
	if !data.Maxcacheableecsprefixlength6.IsNull() && !data.Maxcacheableecsprefixlength6.IsUnknown() {
		dnsprofile.Maxcacheableecsprefixlength6 = utils.IntPtr(int(data.Maxcacheableecsprefixlength6.ValueInt64()))
	}
	if !data.Recursiveresolution.IsNull() && !data.Recursiveresolution.IsUnknown() {
		dnsprofile.Recursiveresolution = data.Recursiveresolution.ValueString()
	}
	if !data.Replaceecs.IsNull() && !data.Replaceecs.IsUnknown() {
		dnsprofile.Replaceecs = data.Replaceecs.ValueString()
	}

	return dnsprofile
}

func dnsprofileSetAttrFromGet(ctx context.Context, data *DnsprofileResourceModel, getResponseData map[string]interface{}) *DnsprofileResourceModel {
	tflog.Debug(ctx, "In dnsprofileSetAttrFromGet Function")

	// Convert API response to model. For Optional+Computed attributes whose value
	// may be omitted from the GET response (e.g. a configured 0/false), only null
	// the field when it was unknown so a configured value is preserved and we avoid
	// "inconsistent result after apply".
	if val, ok := getResponseData["cacheecsresponses"]; ok && val != nil {
		data.Cacheecsresponses = types.StringValue(val.(string))
	} else if data.Cacheecsresponses.IsUnknown() {
		data.Cacheecsresponses = types.StringNull()
	}
	if val, ok := getResponseData["cachenegativeresponses"]; ok && val != nil {
		data.Cachenegativeresponses = types.StringValue(val.(string))
	} else if data.Cachenegativeresponses.IsUnknown() {
		data.Cachenegativeresponses = types.StringNull()
	}
	if val, ok := getResponseData["cacherecords"]; ok && val != nil {
		data.Cacherecords = types.StringValue(val.(string))
	} else if data.Cacherecords.IsUnknown() {
		data.Cacherecords = types.StringNull()
	}
	if val, ok := getResponseData["dnsanswerseclogging"]; ok && val != nil {
		data.Dnsanswerseclogging = types.StringValue(val.(string))
	} else if data.Dnsanswerseclogging.IsUnknown() {
		data.Dnsanswerseclogging = types.StringNull()
	}
	if val, ok := getResponseData["dnserrorlogging"]; ok && val != nil {
		data.Dnserrorlogging = types.StringValue(val.(string))
	} else if data.Dnserrorlogging.IsUnknown() {
		data.Dnserrorlogging = types.StringNull()
	}
	if val, ok := getResponseData["dnsextendedlogging"]; ok && val != nil {
		data.Dnsextendedlogging = types.StringValue(val.(string))
	} else if data.Dnsextendedlogging.IsUnknown() {
		data.Dnsextendedlogging = types.StringNull()
	}
	if val, ok := getResponseData["dnsprofilename"]; ok && val != nil {
		data.Dnsprofilename = types.StringValue(val.(string))
	} else if data.Dnsprofilename.IsUnknown() {
		data.Dnsprofilename = types.StringNull()
	}
	if val, ok := getResponseData["dnsquerylogging"]; ok && val != nil {
		data.Dnsquerylogging = types.StringValue(val.(string))
	} else if data.Dnsquerylogging.IsUnknown() {
		data.Dnsquerylogging = types.StringNull()
	}
	if val, ok := getResponseData["dropmultiqueryrequest"]; ok && val != nil {
		data.Dropmultiqueryrequest = types.StringValue(val.(string))
	} else if data.Dropmultiqueryrequest.IsUnknown() {
		data.Dropmultiqueryrequest = types.StringNull()
	}
	if val, ok := getResponseData["insertecs"]; ok && val != nil {
		data.Insertecs = types.StringValue(val.(string))
	} else if data.Insertecs.IsUnknown() {
		data.Insertecs = types.StringNull()
	}
	if val, ok := getResponseData["maxcacheableecsprefixlength"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxcacheableecsprefixlength = types.Int64Value(intVal)
		}
	} else if data.Maxcacheableecsprefixlength.IsUnknown() {
		data.Maxcacheableecsprefixlength = types.Int64Null()
	}
	if val, ok := getResponseData["maxcacheableecsprefixlength6"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxcacheableecsprefixlength6 = types.Int64Value(intVal)
		}
	} else if data.Maxcacheableecsprefixlength6.IsUnknown() {
		data.Maxcacheableecsprefixlength6 = types.Int64Null()
	}
	if val, ok := getResponseData["recursiveresolution"]; ok && val != nil {
		data.Recursiveresolution = types.StringValue(val.(string))
	} else if data.Recursiveresolution.IsUnknown() {
		data.Recursiveresolution = types.StringNull()
	}
	if val, ok := getResponseData["replaceecs"]; ok && val != nil {
		data.Replaceecs = types.StringValue(val.(string))
	} else if data.Replaceecs.IsUnknown() {
		data.Replaceecs = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(data.Dnsprofilename.ValueString())

	return data
}
