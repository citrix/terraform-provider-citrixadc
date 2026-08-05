package dnsnsrec

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/dns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
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
			"nameserver": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Host name of the name server to add to the domain.",
			},
			"ttl": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Time to Live (TTL), in seconds, for the record. TTL is the time for which the record must be cached by DNS proxies. The specified TTL is applied to all the resource records that are of the same record type and belong to the specified domain name. For example, if you add an address record, with a TTL of 36000, to the domain name example.com, the TTLs of all the address records of example.com are changed to 36000. If the TTL is not specified, the Citrix ADC uses either the DNS zone's minimum TTL or, if the SOA record is not available on the appliance, the default value of 3600.",
			},
			// ecssubnet is a delete-only removal parameter in NITRO (it is not part of
			// the add payload and was not exposed by the SDK v2 resource). It is kept as
			// a read-only computed output populated from GET.
			"ecssubnet": schema.StringAttribute{
				Computed:    true,
				Description: "Subnet for which the cached name server record need to be removed.",
			},
			// nodeid is a read-only cluster attribute (not part of the add payload and
			// not exposed by the SDK v2 resource). Kept as a computed output.
			"nodeid": schema.Int64Attribute{
				Computed:    true,
				Description: "Unique number that identifies the cluster node.",
			},
		},
	}
}

func dnsnsrecGetThePayloadFromthePlan(ctx context.Context, data *DnsnsrecResourceModel) dns.Dnsnsrec {
	tflog.Debug(ctx, "In dnsnsrecGetThePayloadFromthePlan Function")

	// Create API request body from the model. Only the add-supported fields are
	// sent (domain, nameserver, ttl) - matching the SDK v2 resource and the NITRO
	// add contract. ecssubnet/nodeid are read-only/removal-only and are excluded.
	dnsnsrec := dns.Dnsnsrec{}
	if !data.Domain.IsNull() && !data.Domain.IsUnknown() {
		dnsnsrec.Domain = data.Domain.ValueString()
	}
	if !data.Nameserver.IsNull() && !data.Nameserver.IsUnknown() {
		dnsnsrec.Nameserver = data.Nameserver.ValueString()
	}
	if !data.Ttl.IsNull() && !data.Ttl.IsUnknown() {
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
	if val, ok := getResponseData["nameserver"]; ok && val != nil {
		data.Nameserver = types.StringValue(val.(string))
	} else {
		data.Nameserver = types.StringNull()
	}
	if val, ok := getResponseData["ttl"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ttl = types.Int64Value(intVal)
		}
	} else if data.Ttl.IsUnknown() {
		// NITRO omitted ttl and the plan had no configured value: only then null it.
		// This preserves a configured ttl (prevents "inconsistent result after apply").
		data.Ttl = types.Int64Null()
	}
	if val, ok := getResponseData["ecssubnet"]; ok && val != nil {
		data.Ecssubnet = types.StringValue(val.(string))
	} else {
		data.Ecssubnet = types.StringNull()
	}
	if val, ok := getResponseData["nodeid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Nodeid = types.Int64Value(intVal)
		} else {
			data.Nodeid = types.Int64Null()
		}
	} else {
		data.Nodeid = types.Int64Null()
	}

	// Set ID for the resource using the legacy SDK v2 composite "domain,nameserver"
	// format (a domain may have multiple name server records, so the identity is the
	// pair). This keeps imported SDK v2 state and existing tooling working unchanged.
	data.Id = types.StringValue(fmt.Sprintf("%s,%s", data.Domain.ValueString(), data.Nameserver.ValueString()))

	return data
}
