package quicbridgeprofile

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// QuicbridgeprofileDataSourceModel is the data-source-specific model, decoupled
// from QuicbridgeprofileResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits (e.g.
// refcnt). Every non-key attribute is Computed; the Framework's per-attribute
// model <-> schema reflection requires this model to have exactly the attributes
// the data-source schema declares, which is why it cannot reuse the resource
// model.
type QuicbridgeprofileDataSourceModel struct {
	Id               types.String `tfsdk:"id"`
	Name             types.String `tfsdk:"name"` // Required lookup key
	Routingalgorithm types.String `tfsdk:"routingalgorithm"`
	Serveridlength   types.Int64  `tfsdk:"serveridlength"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/quicbridgeprofile.json). Never settable; populated from GET.
	Refcnt types.Int64 `tfsdk:"refcnt"`
}

func QuicbridgeprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the QUIC profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@),equals sign (=), and hyphen (-) characters. Cannot be changed after the profile is created.",
			},
			"routingalgorithm": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Routing algorithm to generate routable connection IDs.",
			},
			"serveridlength": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Length of serverid to encode/decode server information",
			},

			// Read-only (GET-only) attributes surfaced by the data source (these
			// are intentionally NOT modeled on the resource). All Computed.
			"refcnt": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of entities using this profile.",
			},
		},
	}
}

// quicbridgeprofileDataSourceSetAttrFromGet projects a NITRO quicbridgeprofile
// GET response onto the data-source model. Because a data source has no
// plan/apply reconciliation, attributes are simply filled from the GET (or left
// Null when the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func quicbridgeprofileDataSourceSetAttrFromGet(ctx context.Context, data *QuicbridgeprofileDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In quicbridgeprofileDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Routingalgorithm = utils.MapGetString(g, "routingalgorithm")
	data.Serveridlength = utils.MapGetInt64(g, "serveridlength")

	// Read-only attributes.
	data.Refcnt = utils.MapGetInt64(g, "refcnt")
}
