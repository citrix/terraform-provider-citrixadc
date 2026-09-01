package cmpparameter

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// CmpparameterDataSourceModel is the data-source-specific model, decoupled from
// CmpparameterResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits
// (builtin, feature). Every non-key attribute is Computed; the Framework's
// per-attribute model <-> schema reflection requires this model to have exactly
// the attributes the data-source schema declares, which is why it cannot reuse
// the resource model.
type CmpparameterDataSourceModel struct {
	Id                          types.String `tfsdk:"id"`
	Addvaryheader               types.String `tfsdk:"addvaryheader"`
	Cmpbypasspct                types.Int64  `tfsdk:"cmpbypasspct"`
	Cmplevel                    types.String `tfsdk:"cmplevel"`
	Cmponpush                   types.String `tfsdk:"cmponpush"`
	Externalcache               types.String `tfsdk:"externalcache"`
	Heurexpiry                  types.String `tfsdk:"heurexpiry"`
	Heurexpiryhistwt            types.Int64  `tfsdk:"heurexpiryhistwt"`
	Heurexpirythres             types.Int64  `tfsdk:"heurexpirythres"`
	Minressize                  types.Int64  `tfsdk:"minressize"`
	Policytype                  types.String `tfsdk:"policytype"`
	Quantumsize                 types.Int64  `tfsdk:"quantumsize"`
	Randomgzipfilename          types.String `tfsdk:"randomgzipfilename"`
	Randomgzipfilenamemaxlength types.Int64  `tfsdk:"randomgzipfilenamemaxlength"`
	Randomgzipfilenameminlength types.Int64  `tfsdk:"randomgzipfilenameminlength"`
	Servercmp                   types.String `tfsdk:"servercmp"`
	Varyheadervalue             types.String `tfsdk:"varyheadervalue"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/cmpparameter.json). Never settable; populated from GET.
	Builtin types.List   `tfsdk:"builtin"`
	Feature types.String `tfsdk:"feature"`
}

func CmpparameterDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"addvaryheader": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Control insertion of the Vary header in HTTP responses compressed by Citrix ADC. Intermediate caches store different versions of the response for different values of the headers present in the Vary response header.",
			},
			"cmpbypasspct": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Citrix ADC CPU threshold after which compression is not performed. Range: 0 - 100",
			},
			"cmplevel": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify a compression level. Available settings function as follows:\n * Optimal - Corresponds to a gzip GZIP level of 5-7.\n * Best speed - Corresponds to a gzip level of 1.\n * Best compression - Corresponds to a gzip level of 9.",
			},
			"cmponpush": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Citrix ADC does not wait for the quantum to be filled before starting to compress data. Upon receipt of a packet with a PUSH flag, the appliance immediately begins compression of the accumulated packets.",
			},
			"externalcache": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable insertion of  Cache-Control: private response directive to indicate response message is intended for a single user and must not be cached by a shared or proxy cache.",
			},
			"heurexpiry": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Heuristic basefile expiry.",
			},
			"heurexpiryhistwt": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "For heuristic basefile expiry, weightage to be given to historical delta compression ratio, specified as percentage.  For example, to give 25% weightage to historical ratio (and therefore 75% weightage to the ratio for current delta compression transaction), specify 25.",
			},
			"heurexpirythres": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Threshold compression ratio for heuristic basefile expiry, multiplied by 100. For example, to set the threshold ratio to 1.25, specify 125.",
			},
			"minressize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Smallest response size, in bytes, to be compressed.",
			},
			"policytype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of the policy. The only possible value is ADVANCED",
			},
			"quantumsize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum quantum of data to be filled before compression begins.",
			},
			"randomgzipfilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Control the addition of a random filename of random length in the GZIP header to apply the Heal-the-BREACH mitigation for the BREACH attack.",
			},
			"randomgzipfilenamemaxlength": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum length of the random filename to be added in the GZIP header to apply the Heal-the-BREACH mitigation for the BREACH attack.",
			},
			"randomgzipfilenameminlength": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum length of the random filename to be added in the GZIP header to apply the Heal-the-BREACH mitigation for the BREACH attack.",
			},
			"servercmp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow the server to send compressed data to the Citrix ADC. With the default setting, the Citrix ADC appliance handles all compression.",
			},
			"varyheadervalue": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The value of the HTTP Vary header for compressed responses. If this argument is not specified, a default value of \"Accept-Encoding\" will be used.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Flag to determine whether compression is default or not.",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
		},
	}
}

// cmpparameterDataSourceSetAttrFromGet projects a NITRO cmpparameter GET response
// onto the data-source model. cmpparameter is a singleton (unnamed) config
// resource, so the id is a static value. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func cmpparameterDataSourceSetAttrFromGet(ctx context.Context, data *CmpparameterDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In cmpparameterDataSourceSetAttrFromGet Function")

	// Read/write attributes as read-back outputs.
	data.Addvaryheader = utils.MapGetString(g, "addvaryheader")
	data.Cmpbypasspct = utils.MapGetInt64(g, "cmpbypasspct")
	data.Cmplevel = utils.MapGetString(g, "cmplevel")
	data.Cmponpush = utils.MapGetString(g, "cmponpush")
	data.Externalcache = utils.MapGetString(g, "externalcache")
	data.Heurexpiry = utils.MapGetString(g, "heurexpiry")
	data.Heurexpiryhistwt = utils.MapGetInt64(g, "heurexpiryhistwt")
	data.Heurexpirythres = utils.MapGetInt64(g, "heurexpirythres")
	data.Minressize = utils.MapGetInt64(g, "minressize")
	data.Policytype = utils.MapGetString(g, "policytype")
	data.Quantumsize = utils.MapGetInt64(g, "quantumsize")
	data.Randomgzipfilename = utils.MapGetString(g, "randomgzipfilename")
	data.Randomgzipfilenamemaxlength = utils.MapGetInt64(g, "randomgzipfilenamemaxlength")
	data.Randomgzipfilenameminlength = utils.MapGetInt64(g, "randomgzipfilenameminlength")
	data.Servercmp = utils.MapGetString(g, "servercmp")
	data.Varyheadervalue = utils.MapGetString(g, "varyheadervalue")

	// Read-only attributes.
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")

	// cmpparameter is a singleton (unnamed) config resource - static ID.
	data.Id = types.StringValue("cmpparameter-config")
}
