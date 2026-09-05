package dnssoarec

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// DnssoarecDataSourceModel is the data-source-specific model, decoupled from
// DnssoarecResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only metadata attributes that the resource deliberately
// omits (authtype). Every non-key attribute is Computed; the Framework's
// per-attribute model <-> schema reflection requires this model to have exactly
// the attributes the data-source schema declares, which is why it cannot reuse
// the resource model.
type DnssoarecDataSourceModel struct {
	Id           types.String `tfsdk:"id"`
	Contact      types.String `tfsdk:"contact"`
	Domain       types.String `tfsdk:"domain"` // Required lookup key
	Ecssubnet    types.String `tfsdk:"ecssubnet"`
	Expire       types.Int64  `tfsdk:"expire"`
	Minimum      types.Int64  `tfsdk:"minimum"`
	Nodeid       types.Int64  `tfsdk:"nodeid"`
	Originserver types.String `tfsdk:"originserver"`
	Refresh      types.Int64  `tfsdk:"refresh"`
	Retry        types.Int64  `tfsdk:"retry"`
	Serial       types.Int64  `tfsdk:"serial"`
	Ttl          types.Int64  `tfsdk:"ttl"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/dnssoarec.json). Never settable; populated from GET.
	Authtype types.String `tfsdk:"authtype"`
}

func DnssoarecDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"contact": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Email address of the contact to whom domain issues can be addressed. In the email address, replace the @ sign with a period (.). For example, enter domainadmin.example.com instead of domainadmin@example.com.",
			},
			"domain": schema.StringAttribute{
				Required:    true,
				Description: "Domain name for which to add the SOA record.",
			},
			"ecssubnet": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Subnet for which the cached SOA record need to be removed.",
			},
			"expire": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in seconds, after which the zone data on a secondary name server can no longer be considered authoritative because all refresh and retry attempts made during the period have failed. After the expiry period, the secondary server stops serving the zone. Typically one week. Not used by the primary server.",
			},
			"minimum": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Default time to live (TTL) for all records in the zone. Can be overridden for individual records.",
			},
			"nodeid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique number that identifies the cluster node.",
			},
			"originserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Domain name of the name server that responds authoritatively for the domain.",
			},
			"refresh": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in seconds, for which a secondary server must wait between successive checks on the value of the serial number.",
			},
			"retry": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in seconds, between retries if a secondary server's attempt to contact the primary server for a zone refresh fails.",
			},
			"serial": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The secondary server uses this parameter to determine whether it requires a zone transfer from the primary server.",
			},
			"ttl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time to Live (TTL), in seconds, for the record. TTL is the time for which the record must be cached by DNS proxies. The specified TTL is applied to all the resource records that are of the same record type and belong to the specified domain name. For example, if you add an address record, with a TTL of 36000, to the domain name example.com, the TTLs of all the address records of example.com are changed to 36000. If the TTL is not specified, the Citrix ADC uses either the DNS zone's minimum TTL or, if the SOA record is not available on the appliance, the default value of 3600.",
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

// dnssoarecDataSourceSetAttrFromGet projects a NITRO dnssoarec GET response onto
// the data-source model. Because a data source has no plan/apply reconciliation,
// attributes are simply filled from the GET (or left Null when the GET omits
// them). The shared utils.MapGet* helpers implement that projection.
func dnssoarecDataSourceSetAttrFromGet(ctx context.Context, data *DnssoarecDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In dnssoarecDataSourceSetAttrFromGet Function")

	if v, ok := g["domain"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Domain = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Contact = utils.MapGetString(g, "contact")
	data.Ecssubnet = utils.MapGetString(g, "ecssubnet")
	data.Expire = utils.MapGetInt64(g, "expire")
	data.Minimum = utils.MapGetInt64(g, "minimum")
	data.Nodeid = utils.MapGetInt64(g, "nodeid")
	data.Originserver = utils.MapGetString(g, "originserver")
	data.Refresh = utils.MapGetInt64(g, "refresh")
	data.Retry = utils.MapGetInt64(g, "retry")
	data.Serial = utils.MapGetInt64(g, "serial")
	data.Ttl = utils.MapGetInt64(g, "ttl")

	// Read-only metadata.
	data.Authtype = utils.MapGetString(g, "authtype")
}
