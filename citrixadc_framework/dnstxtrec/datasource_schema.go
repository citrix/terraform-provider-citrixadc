package dnstxtrec

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// DnstxtrecDataSourceModel is the data-source-specific model, decoupled from
// DnstxtrecResourceModel. A data source is a pure read surface, so it can expose
// the FULL GET projection: the read/write attributes (as Computed outputs) AND
// the read-only attributes the resource deliberately omits (authtype).
type DnstxtrecDataSourceModel struct {
	Id     types.String `tfsdk:"id"`
	Domain types.String `tfsdk:"domain"` // Required lookup key

	// Read/write attributes, surfaced here as Computed outputs.
	String    types.List   `tfsdk:"string"`
	Ecssubnet types.String `tfsdk:"ecssubnet"`
	Nodeid    types.Int64  `tfsdk:"nodeid"`
	Recordid  types.Int64  `tfsdk:"recordid"`
	Ttl       types.Int64  `tfsdk:"ttl"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/dnstxtrec.json). Never settable; populated from GET.
	Authtype types.String `tfsdk:"authtype"`
}

func DnstxtrecDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"string": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Information to store in the TXT resource record. Enclose the string in single or double quotation marks. A TXT resource record can contain up to six strings, each of which can contain up to 255 characters. If you want to add a string of more than 255 characters, evaluate whether splitting it into two or more smaller strings, subject to the six-string limit, works for you.",
			},
			"domain": schema.StringAttribute{
				Required:    true,
				Description: "Name of the domain for the TXT record.",
			},
			"ecssubnet": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Subnet for which the cached TXT record need to be removed.",
			},
			"nodeid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique number that identifies the cluster node.",
			},
			"recordid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique, internally generated record ID. View the details of the TXT record to obtain its record ID. Mutually exclusive with the string parameter.",
			},
			"ttl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time to Live (TTL), in seconds, for the record. TTL is the time for which the record must be cached by DNS proxies. The specified TTL is applied to all the resource records that are of the same record type and belong to the specified domain name. For example, if you add an address record, with a TTL of 36000, to the domain name example.com, the TTLs of all the address records of example.com are changed to 36000. If the TTL is not specified, the Citrix ADC uses either the DNS zone's minimum TTL or, if the SOA record is not available on the appliance, the default value of 3600.",
			},

			// Read-only (GET-only) metadata surfaced by the data source (these are
			// intentionally NOT modeled on the resource). All Computed.
			"authtype": schema.StringAttribute{
				Computed:    true,
				Description: "Authentication type. Possible values: [ ALL, ADNS, PROXY ].",
			},
		},
	}
}

// dnstxtrecDataSourceSetAttrFromGet projects a NITRO dnstxtrec GET response onto
// the data-source model. Attributes are simply filled from the GET (or left Null
// when the GET omits them) via the shared utils.MapGet* helpers.
func dnstxtrecDataSourceSetAttrFromGet(ctx context.Context, data *DnstxtrecDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In dnstxtrecDataSourceSetAttrFromGet Function")

	if v, ok := g["domain"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Domain = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.String = utils.MapGetStringList(g, "String")
	data.Ecssubnet = utils.MapGetString(g, "ecssubnet")
	data.Nodeid = utils.MapGetInt64(g, "nodeid")
	data.Recordid = utils.MapGetInt64(g, "recordid")
	data.Ttl = utils.MapGetInt64(g, "ttl")

	// Read-only metadata.
	data.Authtype = utils.MapGetString(g, "authtype")
}
