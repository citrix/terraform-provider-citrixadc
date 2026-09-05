package sslfipskey

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SslfipskeyDataSourceModel is the data-source-specific model, decoupled from
// SslfipskeyResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits.
// Every non-key attribute is Computed; the Framework's per-attribute model <->
// schema reflection requires this model to have exactly the attributes the
// data-source schema declares, which is why it cannot reuse the resource model.
type SslfipskeyDataSourceModel struct {
	Id          types.String `tfsdk:"id"`
	Curve       types.String `tfsdk:"curve"`
	Exponent    types.String `tfsdk:"exponent"`
	Fipskeyname types.String `tfsdk:"fipskeyname"` // Required lookup key
	Inform      types.String `tfsdk:"inform"`
	Iv          types.String `tfsdk:"iv"`
	Key         types.String `tfsdk:"key"`
	Keytype     types.String `tfsdk:"keytype"`
	Modulus     types.Int64  `tfsdk:"modulus"`
	Wrapkeyname types.String `tfsdk:"wrapkeyname"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/sslfipskey.json). Never settable; populated from GET.
	Size types.Int64 `tfsdk:"size"`
}

func SslfipskeyDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"curve": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Only p_256 (prime256v1) and P_384 (secp384r1) are supported.",
			},
			"exponent": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Exponent value for the FIPS key to be created. Available values function as follows:\n 3=3 (hexadecimal)\nF4=10001 (hexadecimal)",
			},
			"fipskeyname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the FIPS key. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the FIPS key is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my fipskey\" or 'my fipskey').",
			},
			"inform": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Input format of the key file. Available formats are:\nSIM - Secure Information Management; select when importing a FIPS key. If the external FIPS key is encrypted, first decrypt it, and then import it.\nPEM - Privacy Enhanced Mail; select when importing a non-FIPS key.",
			},
			"iv": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Initialization Vector (IV) to use for importing the key. Required for importing a non-FIPS key.",
			},
			"key": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of and, optionally, path to the key file to be imported.\n /nsconfig/ssl/ is the default path.",
			},
			"keytype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Only RSA key and ECDSA Key are supported.",
			},
			"modulus": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Modulus, in multiples of 64, of the FIPS key to be created.",
			},
			"wrapkeyname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the wrap key to use for importing the key. Required for importing a non-FIPS key.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (intentionally NOT modeled on the resource). All Computed.
			"size": schema.Int64Attribute{
				Computed:    true,
				Description: "Size.",
			},
		},
	}
}

// sslfipskeyDataSourceSetAttrFromGet projects a NITRO sslfipskey GET response
// onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them) via the shared utils.MapGet* helpers.
func sslfipskeyDataSourceSetAttrFromGet(ctx context.Context, data *SslfipskeyDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In sslfipskeyDataSourceSetAttrFromGet Function")

	if v, ok := g["fipskeyname"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Fipskeyname = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Curve = utils.MapGetString(g, "curve")
	data.Exponent = utils.MapGetString(g, "exponent")
	data.Inform = utils.MapGetString(g, "inform")
	data.Keytype = utils.MapGetString(g, "keytype")
	data.Modulus = utils.MapGetInt64(g, "modulus")

	// iv / key / wrapkeyname are write-only import inputs the GET never
	// returns -> Null.
	data.Iv = types.StringNull()
	data.Key = types.StringNull()
	data.Wrapkeyname = types.StringNull()

	// Read-only attributes.
	data.Size = utils.MapGetInt64(g, "size")
}
