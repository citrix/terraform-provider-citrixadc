package dnssrvrec

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// DnssrvrecDataSourceModel is the data-source-specific model, decoupled from
// DnssrvrecResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only metadata attributes that the resource deliberately
// omits (authtype). Every non-key attribute is Computed; the Framework's
// per-attribute model <-> schema reflection requires this model to have exactly
// the attributes the data-source schema declares, which is why it cannot reuse
// the resource model.
type DnssrvrecDataSourceModel struct {
	Id        types.String `tfsdk:"id"`
	Domain    types.String `tfsdk:"domain"` // Required lookup key
	Ecssubnet types.String `tfsdk:"ecssubnet"`
	Nodeid    types.Int64  `tfsdk:"nodeid"`
	Port      types.Int64  `tfsdk:"port"`
	Priority  types.Int64  `tfsdk:"priority"`
	Target    types.String `tfsdk:"target"` // Required lookup key
	Ttl       types.Int64  `tfsdk:"ttl"`
	Weight    types.Int64  `tfsdk:"weight"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/dnssrvrec.json). Never settable; populated from GET.
	Authtype types.String `tfsdk:"authtype"`
}

func DnssrvrecDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"domain": schema.StringAttribute{
				Required:    true,
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
				Optional:    true,
				Computed:    true,
				Description: "Port on which the target host listens for client requests.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer specifying the priority of the target host. The lower the number, the higher the priority. If multiple target hosts have the same priority, selection is based on the Weight parameter.",
			},
			"target": schema.StringAttribute{
				Required:    true,
				Description: "Target host for the specified service.",
			},
			"ttl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time to Live (TTL), in seconds, for the record. TTL is the time for which the record must be cached by DNS proxies. The specified TTL is applied to all the resource records that are of the same record type and belong to the specified domain name. For example, if you add an address record, with a TTL of 36000, to the domain name example.com, the TTLs of all the address records of example.com are changed to 36000. If the TTL is not specified, the Citrix ADC uses either the DNS zone's minimum TTL or, if the SOA record is not available on the appliance, the default value of 3600.",
			},
			"weight": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Weight for the target host. Aids host selection when two or more hosts have the same priority. A larger number indicates greater weight.",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"authtype": schema.StringAttribute{
				Computed:    true,
				Description: "Record type. Possible values = ALL, ADNS, PROXY.",
			},
		},
	}
}

// dnssrvrecDataSourceSetAttrFromGet projects a NITRO dnssrvrec GET response onto
// the data-source model. Because a data source has no plan/apply reconciliation,
// attributes are simply filled from the GET (or left Null when the GET omits
// them). The shared utils.MapGet* helpers implement that projection.
func dnssrvrecDataSourceSetAttrFromGet(ctx context.Context, data *DnssrvrecDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In dnssrvrecDataSourceSetAttrFromGet Function")

	// Read/write attributes as read-back outputs.
	data.Domain = utils.MapGetString(g, "domain")
	data.Target = utils.MapGetString(g, "target")
	data.Ecssubnet = utils.MapGetString(g, "ecssubnet")
	data.Nodeid = utils.MapGetInt64(g, "nodeid")
	data.Port = utils.MapGetInt64(g, "port")
	data.Priority = utils.MapGetInt64(g, "priority")
	data.Ttl = utils.MapGetInt64(g, "ttl")
	data.Weight = utils.MapGetInt64(g, "weight")

	// Composite ID matches the SDK v2 "domain,target" format.
	data.Id = types.StringValue(data.Domain.ValueString() + "," + data.Target.ValueString())

	// Read-only metadata.
	data.Authtype = utils.MapGetString(g, "authtype")
}
