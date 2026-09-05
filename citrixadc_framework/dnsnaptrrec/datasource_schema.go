package dnsnaptrrec

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// DnsnaptrrecDataSourceModel is the data-source-specific model, decoupled from
// DnsnaptrrecResourceModel. A data source is a pure read surface, so it exposes
// the full GET projection: the read/write attributes (as Computed outputs) AND
// the read-only attributes the resource deliberately omits (authtype,
// vservername). Every non-key attribute is Computed.
type DnsnaptrrecDataSourceModel struct {
	Id          types.String `tfsdk:"id"`
	Domain      types.String `tfsdk:"domain"` // Required lookup key
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

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/dnsnaptrrec.json). Never settable; populated from GET.
	Authtype    types.String `tfsdk:"authtype"`
	Vservername types.String `tfsdk:"vservername"`
}

func DnsnaptrrecDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"domain": schema.StringAttribute{
				Required:    true,
				Description: "Name of the domain for the NAPTR record.",
			},
			"ecssubnet": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Subnet for which the cached NAPTR record need to be removed.",
			},
			"flags": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "flags for this NAPTR.",
			},
			"nodeid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique number that identifies the cluster node.",
			},
			"order": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "An integer specifying the order in which the NAPTR records MUST be processed in order to accurately represent the ordered list of Rules. The ordering is from lowest to highest",
			},
			"preference": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "An integer specifying the preference of this NAPTR among NAPTR records having same order. lower the number, higher the preference.",
			},
			"recordid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique, internally generated record ID. View the details of the naptr record to obtain its record ID. Records can be removed by either specifying the domain name and record id OR by specifying\ndomain name and all other naptr record attributes as was supplied during the add command.",
			},
			"regexp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The regular expression, that specifies the substitution expression for this NAPTR",
			},
			"replacement": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The replacement domain name for this NAPTR.",
			},
			"services": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Service Parameters applicable to this delegation path.",
			},
			"ttl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time to Live (TTL), in seconds, for the record. TTL is the time for which the record must be cached by DNS proxies. The specified TTL is applied to all the resource records that are of the same record type and belong to the specified domain name. For example, if you add an address record, with a TTL of 36000, to the domain name example.com, the TTLs of all the address records of example.com are changed to 36000. If the TTL is not specified, the Citrix ADC uses either the DNS zone's minimum TTL or, if the SOA record is not available on the appliance, the default value of 3600.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"authtype": schema.StringAttribute{
				Computed:    true,
				Description: "Authentication type. Possible values: [ ALL, ADNS, PROXY ]",
			},
			"vservername": schema.StringAttribute{
				Computed:    true,
				Description: "Virtual server name.",
			},
		},
	}
}

// dnsnaptrrecDataSourceSetAttrFromGet projects a NITRO dnsnaptrrec GET response
// onto the data-source model. Attributes are filled from the GET (or left Null
// when the GET omits them) via the shared utils.MapGet* helpers.
func dnsnaptrrecDataSourceSetAttrFromGet(ctx context.Context, data *DnsnaptrrecDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In dnsnaptrrecDataSourceSetAttrFromGet Function")

	if v, ok := g["domain"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Domain = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Ecssubnet = utils.MapGetString(g, "ecssubnet")
	data.Flags = utils.MapGetString(g, "flags")
	data.Nodeid = utils.MapGetInt64(g, "nodeid")
	data.Order = utils.MapGetInt64(g, "order")
	data.Preference = utils.MapGetInt64(g, "preference")
	data.Recordid = utils.MapGetInt64(g, "recordid")
	data.Regexp = utils.MapGetString(g, "regexp")
	data.Replacement = utils.MapGetString(g, "replacement")
	data.Services = utils.MapGetString(g, "services")
	data.Ttl = utils.MapGetInt64(g, "ttl")

	// Read-only attributes.
	data.Authtype = utils.MapGetString(g, "authtype")
	data.Vservername = utils.MapGetString(g, "vservername")
}
