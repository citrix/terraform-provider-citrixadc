package dnsnaptrrec

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/dns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// DnsnaptrrecResourceModel describes the resource data model.
type DnsnaptrrecResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Domain      types.String `tfsdk:"domain"`
	Ecssubnet   types.String `tfsdk:"ecssubnet"`
	Flags       types.String `tfsdk:"flags"`
	Nodeid      types.Int64  `tfsdk:"nodeid"`
	Order       types.Int64  `tfsdk:"order"`
	Preference  types.Int64  `tfsdk:"preference"`
	Recordid    types.Int64  `tfsdk:"recordid"`
	Regexp      types.String `tfsdk:"regexp"`
	Replacement types.String `tfsdk:"replacement"`
	Services    types.String `tfsdk:"services"`
	Ttl         types.Int64  `tfsdk:"ttl"`
}

func (r *DnsnaptrrecResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the dnsnaptrrec resource.",
			},
			"domain": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the domain for the NAPTR record.",
			},
			"ecssubnet": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Subnet for which the cached NAPTR record need to be removed.",
			},
			"flags": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "flags for this NAPTR.",
			},
			"nodeid": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Unique number that identifies the cluster node.",
			},
			"order": schema.Int64Attribute{
				Required: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "An integer specifying the order in which the NAPTR records MUST be processed in order to accurately represent the ordered list of Rules. The ordering is from lowest to highest",
			},
			"preference": schema.Int64Attribute{
				Required: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "An integer specifying the preference of this NAPTR among NAPTR records having same order. lower the number, higher the preference.",
			},
			"recordid": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Unique, internally generated record ID. View the details of the naptr record to obtain its record ID. Records can be removed by either specifying the domain name and record id OR by specifying\ndomain name and all other naptr record attributes as was supplied during the add command.",
			},
			"regexp": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The regular expression, that specifies the substitution expression for this NAPTR",
			},
			"replacement": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The replacement domain name for this NAPTR.",
			},
			"services": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Service Parameters applicable to this delegation path.",
			},
			"ttl": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Default:     int64default.StaticInt64(3600),
				Description: "Time to Live (TTL), in seconds, for the record. TTL is the time for which the record must be cached by DNS proxies. The specified TTL is applied to all the resource records that are of the same record type and belong to the specified domain name. For example, if you add an address record, with a TTL of 36000, to the domain name example.com, the TTLs of all the address records of example.com are changed to 36000. If the TTL is not specified, the Citrix ADC uses either the DNS zone's minimum TTL or, if the SOA record is not available on the appliance, the default value of 3600.",
			},
		},
	}
}

func dnsnaptrrecGetThePayloadFromtheConfig(ctx context.Context, data *DnsnaptrrecResourceModel) dns.Dnsnaptrrec {
	tflog.Debug(ctx, "In dnsnaptrrecGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	dnsnaptrrec := dns.Dnsnaptrrec{}
	if !data.Domain.IsNull() {
		dnsnaptrrec.Domain = data.Domain.ValueString()
	}
	if !data.Ecssubnet.IsNull() {
		dnsnaptrrec.Ecssubnet = data.Ecssubnet.ValueString()
	}
	if !data.Flags.IsNull() {
		dnsnaptrrec.Flags = data.Flags.ValueString()
	}
	if !data.Nodeid.IsNull() {
		dnsnaptrrec.Nodeid = utils.IntPtr(int(data.Nodeid.ValueInt64()))
	}
	if !data.Order.IsNull() {
		dnsnaptrrec.Order = utils.IntPtr(int(data.Order.ValueInt64()))
	}
	if !data.Preference.IsNull() {
		dnsnaptrrec.Preference = utils.IntPtr(int(data.Preference.ValueInt64()))
	}
	if !data.Recordid.IsNull() {
		dnsnaptrrec.Recordid = utils.IntPtr(int(data.Recordid.ValueInt64()))
	}
	if !data.Regexp.IsNull() {
		dnsnaptrrec.Regexp = data.Regexp.ValueString()
	}
	if !data.Replacement.IsNull() {
		dnsnaptrrec.Replacement = data.Replacement.ValueString()
	}
	if !data.Services.IsNull() {
		dnsnaptrrec.Services = data.Services.ValueString()
	}
	if !data.Ttl.IsNull() {
		dnsnaptrrec.Ttl = utils.IntPtr(int(data.Ttl.ValueInt64()))
	}

	return dnsnaptrrec
}

func dnsnaptrrecSetAttrFromGet(ctx context.Context, data *DnsnaptrrecResourceModel, getResponseData map[string]interface{}) *DnsnaptrrecResourceModel {
	tflog.Debug(ctx, "In dnsnaptrrecSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["domain"]; ok && val != nil {
		data.Domain = types.StringValue(val.(string))
	} else {
		data.Domain = types.StringNull()
	}
	if val, ok := getResponseData["ecssubnet"]; ok && val != nil {
		data.Ecssubnet = types.StringValue(val.(string))
	} else {
		data.Ecssubnet = types.StringNull()
	}
	if val, ok := getResponseData["flags"]; ok && val != nil {
		data.Flags = types.StringValue(val.(string))
	} else {
		data.Flags = types.StringNull()
	}
	if val, ok := getResponseData["nodeid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Nodeid = types.Int64Value(intVal)
		}
	} else {
		data.Nodeid = types.Int64Null()
	}
	if val, ok := getResponseData["order"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Order = types.Int64Value(intVal)
		}
	} else {
		data.Order = types.Int64Null()
	}
	if val, ok := getResponseData["preference"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Preference = types.Int64Value(intVal)
		}
	} else {
		data.Preference = types.Int64Null()
	}
	if val, ok := getResponseData["recordid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Recordid = types.Int64Value(intVal)
		}
	} else {
		data.Recordid = types.Int64Null()
	}
	if val, ok := getResponseData["regexp"]; ok && val != nil {
		data.Regexp = types.StringValue(val.(string))
	} else {
		data.Regexp = types.StringNull()
	}
	if val, ok := getResponseData["replacement"]; ok && val != nil {
		data.Replacement = types.StringValue(val.(string))
	} else {
		data.Replacement = types.StringNull()
	}
	if val, ok := getResponseData["services"]; ok && val != nil {
		data.Services = types.StringValue(val.(string))
	} else {
		data.Services = types.StringNull()
	}
	if val, ok := getResponseData["ttl"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ttl = types.Int64Value(intVal)
		}
	} else {
		data.Ttl = types.Int64Null()
	}
	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated
	data.Id = types.StringValue(fmt.Sprintf("%s", data.Domain.ValueString()))

	return data
}
