package dnsnsrec

import (
	"context"
	"fmt"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// DnsnsrecDataSourceModel is the data-source-specific model, decoupled from
// DnsnsrecResourceModel. A data source is a pure read surface, so it exposes the
// full GET projection: the read/write attributes (as Computed outputs) AND the
// read-only attributes the resource deliberately omits (authtype). Every
// non-key attribute is Computed.
type DnsnsrecDataSourceModel struct {
	Id         types.String `tfsdk:"id"`
	Domain     types.String `tfsdk:"domain"` // Required lookup key
	Ecssubnet  types.String `tfsdk:"ecssubnet"`
	Nameserver types.String `tfsdk:"nameserver"`
	Nodeid     types.Int64  `tfsdk:"nodeid"`
	Ttl        types.Int64  `tfsdk:"ttl"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/dnsnsrec.json). Never settable; populated from GET.
	Authtype types.String `tfsdk:"authtype"`
}

func DnsnsrecDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"domain": schema.StringAttribute{
				Required:    true,
				Description: "Domain name.",
			},
			"ecssubnet": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Subnet for which the cached name server record need to be removed.",
			},
			"nameserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Host name of the name server to add to the domain.",
			},
			"nodeid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique number that identifies the cluster node.",
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
				Description: "Record type. Possible values: [ ALL, ADNS, PROXY ]",
			},
		},
	}
}

// dnsnsrecDataSourceSetAttrFromGet projects a NITRO dnsnsrec GET response onto
// the data-source model. Attributes are filled from the GET (or left Null when
// the GET omits them) via the shared utils.MapGet* helpers.
func dnsnsrecDataSourceSetAttrFromGet(ctx context.Context, data *DnsnsrecDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In dnsnsrecDataSourceSetAttrFromGet Function")

	// Read/write attributes as read-back outputs.
	data.Domain = utils.MapGetString(g, "domain")
	data.Ecssubnet = utils.MapGetString(g, "ecssubnet")
	data.Nameserver = utils.MapGetString(g, "nameserver")
	data.Nodeid = utils.MapGetInt64(g, "nodeid")
	data.Ttl = utils.MapGetInt64(g, "ttl")

	// Read-only attributes.
	data.Authtype = utils.MapGetString(g, "authtype")

	// ID uses the legacy SDK v2 composite "domain,nameserver" format (a domain
	// may have multiple name-server records, so the identity is the pair).
	data.Id = types.StringValue(fmt.Sprintf("%s,%s", data.Domain.ValueString(), data.Nameserver.ValueString()))
}
