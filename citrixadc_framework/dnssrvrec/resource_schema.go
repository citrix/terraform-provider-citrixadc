package dnssrvrec

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/dns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// DnssrvrecResourceModel describes the resource data model.
type DnssrvrecResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Domain    types.String `tfsdk:"domain"`
	Ecssubnet types.String `tfsdk:"ecssubnet"`
	Nodeid    types.Int64  `tfsdk:"nodeid"`
	Port      types.Int64  `tfsdk:"port"`
	Priority  types.Int64  `tfsdk:"priority"`
	Target    types.String `tfsdk:"target"`
	Ttl       types.Int64  `tfsdk:"ttl"`
	Weight    types.Int64  `tfsdk:"weight"`
}

func (r *DnssrvrecResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the dnssrvrec resource.",
			},
			"domain": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Domain name, which, by convention, is prefixed by the symbolic name of the desired service and the symbolic name of the desired protocol, each with an underscore (_) prepended. For example, if an SRV-aware client wants to discover a SIP service that is provided over UDP, in the domain example.com, the client performs a lookup for _sip._udp.example.com.",
			},
			"ecssubnet": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Subnet for which the cached SRV record need to be removed.",
			},
			"nodeid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique number that identifies the cluster node.",
			},
			"port": schema.Int64Attribute{
				Required:    true,
				Description: "Port on which the target host listens for client requests.",
			},
			"priority": schema.Int64Attribute{
				Required:    true,
				Description: "Integer specifying the priority of the target host. The lower the number, the higher the priority. If multiple target hosts have the same priority, selection is based on the Weight parameter.",
			},
			"target": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Target host for the specified service.",
			},
			"ttl": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				// NITRO default (spec: Default value 3600). An explicit Default is
				// required so removing ttl from config produces a plan diff, which
				// lets Update fire the unset (revert to 3600).
				Default:     int64default.StaticInt64(3600),
				Description: "Time to Live (TTL), in seconds, for the record. TTL is the time for which the record must be cached by DNS proxies. The specified TTL is applied to all the resource records that are of the same record type and belong to the specified domain name. For example, if you add an address record, with a TTL of 36000, to the domain name example.com, the TTLs of all the address records of example.com are changed to 36000. If the TTL is not specified, the Citrix ADC uses either the DNS zone's minimum TTL or, if the SOA record is not available on the appliance, the default value of 3600.",
			},
			"weight": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Weight for the target host. Aids host selection when two or more hosts have the same priority. A larger number indicates greater weight.",
			},
		},
	}
}

func dnssrvrecGetThePayloadFromthePlan(ctx context.Context, data *DnssrvrecResourceModel) dns.Dnssrvrec {
	tflog.Debug(ctx, "In dnssrvrecGetThePayloadFromthePlan Function")

	// Create API request body from the model
	dnssrvrec := dns.Dnssrvrec{}
	if !data.Domain.IsNull() && !data.Domain.IsUnknown() {
		dnssrvrec.Domain = data.Domain.ValueString()
	}
	if !data.Ecssubnet.IsNull() && !data.Ecssubnet.IsUnknown() {
		dnssrvrec.Ecssubnet = data.Ecssubnet.ValueString()
	}
	if !data.Nodeid.IsNull() && !data.Nodeid.IsUnknown() {
		dnssrvrec.Nodeid = utils.IntPtr(int(data.Nodeid.ValueInt64()))
	}
	if !data.Port.IsNull() && !data.Port.IsUnknown() {
		dnssrvrec.Port = utils.IntPtr(int(data.Port.ValueInt64()))
	}
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		dnssrvrec.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}
	if !data.Target.IsNull() && !data.Target.IsUnknown() {
		dnssrvrec.Target = data.Target.ValueString()
	}
	if !data.Ttl.IsNull() && !data.Ttl.IsUnknown() {
		dnssrvrec.Ttl = utils.IntPtr(int(data.Ttl.ValueInt64()))
	}
	if !data.Weight.IsNull() && !data.Weight.IsUnknown() {
		dnssrvrec.Weight = utils.IntPtr(int(data.Weight.ValueInt64()))
	}

	return dnssrvrec
}

func dnssrvrecSetAttrFromGet(ctx context.Context, data *DnssrvrecResourceModel, getResponseData map[string]interface{}) *DnssrvrecResourceModel {
	tflog.Debug(ctx, "In dnssrvrecSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["domain"]; ok && val != nil {
		data.Domain = types.StringValue(val.(string))
	} else if data.Domain.IsUnknown() {
		data.Domain = types.StringNull()
	}
	if val, ok := getResponseData["ecssubnet"]; ok && val != nil {
		data.Ecssubnet = types.StringValue(val.(string))
	} else if data.Ecssubnet.IsUnknown() {
		// NITRO omits ecssubnet from GET when unset; only null it when it was
		// unknown (never configured) so a user-supplied value is preserved.
		data.Ecssubnet = types.StringNull()
	}
	if val, ok := getResponseData["nodeid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Nodeid = types.Int64Value(intVal)
		}
	} else if data.Nodeid.IsUnknown() {
		// NITRO omits nodeid (default 0) from GET; preserve a configured value.
		data.Nodeid = types.Int64Null()
	}
	if val, ok := getResponseData["port"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Port = types.Int64Value(intVal)
		}
	} else if data.Port.IsUnknown() {
		data.Port = types.Int64Null()
	}
	if val, ok := getResponseData["priority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Priority = types.Int64Value(intVal)
		}
	} else if data.Priority.IsUnknown() {
		data.Priority = types.Int64Null()
	}
	if val, ok := getResponseData["target"]; ok && val != nil {
		data.Target = types.StringValue(val.(string))
	} else if data.Target.IsUnknown() {
		data.Target = types.StringNull()
	}
	if val, ok := getResponseData["ttl"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ttl = types.Int64Value(intVal)
		}
	} else if data.Ttl.IsUnknown() {
		data.Ttl = types.Int64Null()
	}
	if val, ok := getResponseData["weight"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Weight = types.Int64Value(intVal)
		}
	} else if data.Weight.IsUnknown() {
		data.Weight = types.Int64Null()
	}

	// Set ID for the resource - matches the SDK v2 composite "domain,target" format.
	data.Id = types.StringValue(fmt.Sprintf("%s,%s", data.Domain.ValueString(), data.Target.ValueString()))

	return data
}
