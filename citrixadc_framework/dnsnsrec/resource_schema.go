package dnsnsrec

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

// DnsnsrecResourceModel describes the resource data model.
type DnsnsrecResourceModel struct {
	Id         types.String `tfsdk:"id"`
	Domain     types.String `tfsdk:"domain"`
	Ecssubnet  types.String `tfsdk:"ecssubnet"`
	Nameserver types.String `tfsdk:"nameserver"`
	Nodeid     types.Int64  `tfsdk:"nodeid"`
	Ttl        types.Int64  `tfsdk:"ttl"`
}

func (r *DnsnsrecResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the dnsnsrec resource.",
			},
			"domain": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Domain name.",
			},
			"ecssubnet": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Subnet for which the cached name server record need to be removed.",
			},
			"nameserver": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Host name of the name server to add to the domain.",
			},
			"nodeid": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Unique number that identifies the cluster node.",
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

func dnsnsrecGetThePayloadFromtheConfig(ctx context.Context, data *DnsnsrecResourceModel) dns.Dnsnsrec {
	tflog.Debug(ctx, "In dnsnsrecGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	dnsnsrec := dns.Dnsnsrec{}
	if !data.Domain.IsNull() {
		dnsnsrec.Domain = data.Domain.ValueString()
	}
	if !data.Ecssubnet.IsNull() {
		dnsnsrec.Ecssubnet = data.Ecssubnet.ValueString()
	}
	if !data.Nameserver.IsNull() {
		dnsnsrec.Nameserver = data.Nameserver.ValueString()
	}
	if !data.Nodeid.IsNull() {
		dnsnsrec.Nodeid = utils.IntPtr(int(data.Nodeid.ValueInt64()))
	}
	if !data.Ttl.IsNull() {
		dnsnsrec.Ttl = utils.IntPtr(int(data.Ttl.ValueInt64()))
	}

	return dnsnsrec
}

func dnsnsrecSetAttrFromGet(ctx context.Context, data *DnsnsrecResourceModel, getResponseData map[string]interface{}) *DnsnsrecResourceModel {
	tflog.Debug(ctx, "In dnsnsrecSetAttrFromGet Function")

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
	if val, ok := getResponseData["nameserver"]; ok && val != nil {
		data.Nameserver = types.StringValue(val.(string))
	} else {
		data.Nameserver = types.StringNull()
	}
	if val, ok := getResponseData["nodeid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Nodeid = types.Int64Value(intVal)
		}
	} else {
		data.Nodeid = types.Int64Null()
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
