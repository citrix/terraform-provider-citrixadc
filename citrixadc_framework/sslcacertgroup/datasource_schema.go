package sslcacertgroup

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SslcacertgroupDataSourceModel is the data-source-specific model, decoupled from
// SslcacertgroupResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only metadata attributes that the resource deliberately
// omits. Every non-key attribute is Computed; the Framework's per-attribute
// model <-> schema reflection requires this model to have exactly the attributes
// the data-source schema declares, which is why it cannot reuse the resource
// model.
type SslcacertgroupDataSourceModel struct {
	Id              types.String `tfsdk:"id"`
	Cacertgroupname types.String `tfsdk:"cacertgroupname"` // Required lookup key

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/sslcacertgroup.json). Never settable; populated from GET.
	Cacertgroupreferences types.Int64  `tfsdk:"cacertgroupreferences"`
	Ocspcheck             types.String `tfsdk:"ocspcheck"`
	Crlcheck              types.String `tfsdk:"crlcheck"`
}

func SslcacertgroupDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"cacertgroupname": schema.StringAttribute{
				Required:    true,
				Description: "Name given to the CA certificate group. The name will be used to add the CA certificates to the group. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my file\" or 'my file').",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"cacertgroupreferences": schema.Int64Attribute{
				Computed:    true,
				Description: "Count for ssl actions referring to this ca certificate group.",
			},
			"ocspcheck": schema.StringAttribute{
				Computed:    true,
				Description: "The state of the OCSP check parameter. (Mandatory/Optional).",
			},
			"crlcheck": schema.StringAttribute{
				Computed:    true,
				Description: "The state of the CRL check parameter. (Mandatory/Optional).",
			},
		},
	}
}

// sslcacertgroupDataSourceSetAttrFromGet projects a NITRO sslcacertgroup GET
// response onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func sslcacertgroupDataSourceSetAttrFromGet(ctx context.Context, data *SslcacertgroupDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In sslcacertgroupDataSourceSetAttrFromGet Function")

	if v, ok := g["cacertgroupname"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Cacertgroupname = types.StringValue(utils.AnyToString(v))
	}

	// Read-only metadata.
	data.Cacertgroupreferences = utils.MapGetInt64(g, "cacertgroupreferences")
	data.Ocspcheck = utils.MapGetString(g, "ocspcheck")
	data.Crlcheck = utils.MapGetString(g, "crlcheck")
}
