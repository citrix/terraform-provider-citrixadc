package dnscaarec

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// DnscaarecDataSourceModel is the data-source-specific model, decoupled from
// DnscaarecResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits.
// The Framework's per-attribute model <-> schema reflection requires this model
// to have exactly the attributes the data-source schema declares.
type DnscaarecDataSourceModel struct {
	Id          types.String `tfsdk:"id"`
	Domain      types.String `tfsdk:"domain"` // Required lookup key
	Ecssubnet   types.String `tfsdk:"ecssubnet"`
	Flag        types.String `tfsdk:"flag"`
	Recordid    types.Int64  `tfsdk:"recordid"` // Required lookup key
	Tag         types.String `tfsdk:"tag"`
	Ttl         types.Int64  `tfsdk:"ttl"`
	Valuestring types.String `tfsdk:"valuestring"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/dnscaarec.json). Never settable; populated from GET.
	Authtype types.String `tfsdk:"authtype"`
}

func DnscaarecDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"domain": schema.StringAttribute{
				Required:    true,
				Description: "Domain name of the CAA record.",
			},
			"ecssubnet": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Subnet for which the cached CAA record need to be removed.",
			},
			"flag": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Flag associated with the CAA record.",
			},
			"recordid": schema.Int64Attribute{
				Required:    true,
				Description: "Unique, internally generated record ID. View the details of the CAA record to obtain its record ID. Records can be removedby either specifying the domain name and record id OR by specifying domain name and all other CAA record attributes as was supplied during the add command.",
			},
			"tag": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "String that represents the identifier of the property represented by the CAA record. The RFC currently defines three available tags - issue, issuwild and iodef.",
			},
			"ttl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time to Live (TTL), in seconds, for the record. TTL is the time for which the record must be cached by DNS proxies. The specified TTL is applied to all the resource records that are of the same record type and belong to the specified domain name. For example, if you add an address record, with a TTL of 36000, to the domain name example.com, the TTLs of all the address records of example.com are changed to 36000. If the TTL is not specified, the Citrix ADC uses either the DNS zone's minimum TTL or, if the SOA record is not available on the appliance, the default value of 3600.",
			},
			"valuestring": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Value associated with the chosen property tag in the CAA resource record. Enclose the string in single or double quotation marks.",
			},

			// Read-only (GET-only) attributes surfaced by the data source (these
			// are intentionally NOT modeled on the resource). All Computed.
			"authtype": schema.StringAttribute{
				Computed:    true,
				Description: "Authentication type. Possible values = ALL, ADNS, PROXY.",
			},
		},
	}
}

// dnscaarecDataSourceSetAttrFromGet projects a NITRO dnscaarec GET response onto
// the data-source model. Because a data source has no plan/apply reconciliation,
// attributes are simply filled from the GET (or left Null when the GET omits
// them). The shared utils.MapGet* helpers implement that projection.
func dnscaarecDataSourceSetAttrFromGet(ctx context.Context, data *DnscaarecDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In dnscaarecDataSourceSetAttrFromGet Function")

	data.Domain = utils.MapGetString(g, "domain")
	data.Ecssubnet = utils.MapGetString(g, "ecssubnet")
	data.Flag = utils.MapGetString(g, "flag")
	data.Recordid = utils.MapGetInt64(g, "recordid")
	data.Tag = utils.MapGetString(g, "tag")
	data.Ttl = utils.MapGetInt64(g, "ttl")
	data.Valuestring = utils.MapGetString(g, "valuestring")

	// Read-only attributes.
	data.Authtype = utils.MapGetString(g, "authtype")

	// Composite key: domain,recordid (multiple CAA records may share a domain).
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("domain:%s", utils.UrlEncode(data.Domain.ValueString())))
	idParts = append(idParts, fmt.Sprintf("recordid:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Recordid.ValueInt64()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))
}
