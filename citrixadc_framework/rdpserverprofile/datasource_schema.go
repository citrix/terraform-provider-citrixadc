package rdpserverprofile

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// RdpserverprofileDataSourceModel is the data-source-specific model, decoupled
// from RdpserverprofileResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits (builtin,
// feature). Every non-key attribute is Computed; the Framework's per-attribute
// model <-> schema reflection requires this model to have exactly the attributes
// the data-source schema declares, which is why it cannot reuse the resource
// model.
type RdpserverprofileDataSourceModel struct {
	Id             types.String `tfsdk:"id"`
	Name           types.String `tfsdk:"name"` // Required lookup key
	Psk            types.String `tfsdk:"psk"`
	PskWo          types.String `tfsdk:"psk_wo"`
	PskWoVersion   types.Int64  `tfsdk:"psk_wo_version"`
	Rdpip          types.String `tfsdk:"rdpip"`
	Rdpport        types.Int64  `tfsdk:"rdpport"`
	Rdpredirection types.String `tfsdk:"rdpredirection"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/rdpserverprofile.json). Never settable; populated from GET.
	Builtin types.List   `tfsdk:"builtin"`
	Feature types.String `tfsdk:"feature"`
}

func RdpserverprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the rdp server profile",
			},
			"psk": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Pre shared key value",
			},
			"psk_wo": schema.StringAttribute{
				Optional:    true,
				Description: "Pre shared key value",
			},
			"psk_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Increment this version to signal a psk_wo update.",
			},
			"rdpip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IPv4 or IPv6 address of RDP listener. This terminates client RDP connections.",
			},
			"rdpport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "TCP port on which the RDP connection is established.",
			},
			"rdpredirection": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable/Disable RDP redirection support. This needs to be enabled in presence of connection broker or session directory with IP cookie(msts cookie) based redirection support",
			},

			// Read-only (GET-only) attributes surfaced by the data source (these
			// are intentionally NOT modeled on the resource). All Computed.
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Indicates that a variable is a built-in (SYSTEM INTERNAL) type. Possible values = MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
		},
	}
}

// rdpserverprofileDataSourceSetAttrFromGet projects a NITRO rdpserverprofile GET
// response onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func rdpserverprofileDataSourceSetAttrFromGet(ctx context.Context, data *RdpserverprofileDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In rdpserverprofileDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Rdpip = utils.MapGetString(g, "rdpip")
	data.Rdpport = utils.MapGetInt64(g, "rdpport")
	data.Rdpredirection = utils.MapGetString(g, "rdpredirection")

	// psk / psk_wo / psk_wo_version are write-only secret / version-tracker inputs
	// the GET never returns -> Null.
	data.Psk = types.StringNull()
	data.PskWo = types.StringNull()
	data.PskWoVersion = types.Int64Null()

	// Read-only attributes.
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
}
