package dnsmxrec

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// DnsmxrecDataSourceModel is the data-source-specific model, decoupled from
// DnsmxrecResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits.
// The Framework's per-attribute model <-> schema reflection requires this model
// to have exactly the attributes the data-source schema declares.
type DnsmxrecDataSourceModel struct {
	Id        types.String `tfsdk:"id"`
	Domain    types.String `tfsdk:"domain"` // Required lookup key
	Ecssubnet types.String `tfsdk:"ecssubnet"`
	Mx        types.String `tfsdk:"mx"`
	Nodeid    types.Int64  `tfsdk:"nodeid"`
	Pref      types.Int64  `tfsdk:"pref"`
	Ttl       types.Int64  `tfsdk:"ttl"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/dnsmxrec.json). Never settable; populated from GET.
	Authtype types.String `tfsdk:"authtype"`
}

func DnsmxrecDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"domain": schema.StringAttribute{
				Required:    true,
				Description: "Domain name for which to add the MX record.",
			},
			"ecssubnet": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Subnet for which the cached MX record need to be removed.",
			},
			"mx": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Host name of the mail exchange server.",
			},
			"nodeid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique number that identifies the cluster node.",
			},
			"pref": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Priority number to assign to the mail exchange server. A domain name can have multiple mail servers, with a priority number assigned to each server. The lower the priority number, the higher the mail server's priority. When other mail servers have to deliver mail to the specified domain, they begin with the mail server with the lowest priority number, and use other configured mail servers, in priority order, as backups.",
			},
			"ttl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time to Live (TTL), in seconds, for the record. TTL is the time for which the record must be cached by DNS proxies. The specified TTL is applied to all the resource records that are of the same record type and belong to the specified domain name. For example, if you add an address record, with a TTL of 36000, to the domain name example.com, the TTLs of all the address records of example.com are changed to 36000. If the TTL is not specified, the Citrix ADC uses either the DNS zone's minimum TTL or, if the SOA record is not available on the appliance, the default value of 3600.",
			},

			// Read-only (GET-only) attributes surfaced by the data source (these
			// are intentionally NOT modeled on the resource). All Computed.
			"authtype": schema.StringAttribute{
				Computed:    true,
				Description: "Record type. Possible values = ALL, ADNS, PROXY.",
			},
		},
	}
}

// dnsmxrecDataSourceSetAttrFromGet projects a NITRO dnsmxrec GET response onto
// the data-source model. Because a data source has no plan/apply reconciliation,
// attributes are simply filled from the GET (or left Null when the GET omits
// them). The shared utils.MapGet* helpers implement that projection.
func dnsmxrecDataSourceSetAttrFromGet(ctx context.Context, data *DnsmxrecDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In dnsmxrecDataSourceSetAttrFromGet Function")

	if v, ok := g["domain"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Domain = types.StringValue(utils.AnyToString(v))
	}

	data.Ecssubnet = utils.MapGetString(g, "ecssubnet")
	data.Mx = utils.MapGetString(g, "mx")
	data.Nodeid = utils.MapGetInt64(g, "nodeid")
	data.Pref = utils.MapGetInt64(g, "pref")
	data.Ttl = utils.MapGetInt64(g, "ttl")

	// Read-only attributes.
	data.Authtype = utils.MapGetString(g, "authtype")
}
