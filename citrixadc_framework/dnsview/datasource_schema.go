package dnsview

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// DnsviewDataSourceModel is the data-source-specific model, decoupled from
// DnsviewResourceModel. A data source is a pure read surface, so it can expose
// the FULL GET projection: the lookup key (as Computed output) AND the read-only
// attributes the resource deliberately omits (flags).
type DnsviewDataSourceModel struct {
	Id       types.String `tfsdk:"id"`
	Viewname types.String `tfsdk:"viewname"` // Required lookup key

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/dnsview.json). Never settable; populated from GET.
	Flags types.Int64 `tfsdk:"flags"`
}

func DnsviewDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"viewname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the DNS view.",
			},

			// Read-only (GET-only) metadata surfaced by the data source (these are
			// intentionally NOT modeled on the resource). All Computed.
			"flags": schema.Int64Attribute{
				Computed:    true,
				Description: "Flags controlling display.",
			},
		},
	}
}

// dnsviewDataSourceSetAttrFromGet projects a NITRO dnsview GET response onto the
// data-source model. Attributes are simply filled from the GET (or left Null when
// the GET omits them) via the shared utils.MapGet* helpers.
func dnsviewDataSourceSetAttrFromGet(ctx context.Context, data *DnsviewDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In dnsviewDataSourceSetAttrFromGet Function")

	if v, ok := g["viewname"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Viewname = types.StringValue(utils.AnyToString(v))
	}

	// Read-only metadata.
	data.Flags = utils.MapGetInt64(g, "flags")
}
