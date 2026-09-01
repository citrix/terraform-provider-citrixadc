package nshttpparam

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// NshttpparamDataSourceModel is the data-source-specific model, decoupled from
// NshttpparamResourceModel.
//
// nshttpparam is a global singleton (no lookup key). A data source is a pure read
// surface, so it exposes the read/write attributes (as Computed outputs) AND the
// read-only attributes the resource deliberately omits (builtin, feature), which
// are populated only from a GET.
type NshttpparamDataSourceModel struct {
	Id                        types.String `tfsdk:"id"`
	Conmultiplex              types.String `tfsdk:"conmultiplex"`
	Dropinvalreqs             types.String `tfsdk:"dropinvalreqs"`
	Http2serverside           types.String `tfsdk:"http2serverside"`
	Ignoreconnectcodingscheme types.String `tfsdk:"ignoreconnectcodingscheme"`
	Insnssrvrhdr              types.String `tfsdk:"insnssrvrhdr"`
	Logerrresp                types.String `tfsdk:"logerrresp"`
	Markconnreqinval          types.String `tfsdk:"markconnreqinval"`
	Markhttp09inval           types.String `tfsdk:"markhttp09inval"`
	Maxreusepool              types.Int64  `tfsdk:"maxreusepool"`
	Nssrvrhdr                 types.String `tfsdk:"nssrvrhdr"`

	// Read-only (GET-only) attributes from the NITRO read-only set. Never
	// settable; populated from GET.
	Builtin types.List   `tfsdk:"builtin"`
	Feature types.String `tfsdk:"feature"`
}

func NshttpparamDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"conmultiplex": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Reuse server connections for requests from more than one client connections.",
			},
			"dropinvalreqs": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Drop invalid HTTP requests or responses.",
			},
			"http2serverside": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable/Disable HTTP/2 on server side",
			},
			"ignoreconnectcodingscheme": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Ignore Coding scheme in CONNECT request.",
			},
			"insnssrvrhdr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable Citrix ADC server header insertion for Citrix ADC generated HTTP responses.",
			},
			"logerrresp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Server header value to be inserted.",
			},
			"markconnreqinval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Mark CONNECT requests as invalid.",
			},
			"markhttp09inval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Mark HTTP/0.9 requests as invalid.",
			},
			"maxreusepool": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum limit on the number of connections, from the Citrix ADC to a particular server that are kept in the reuse pool. This setting is helpful for optimal memory utilization and for reducing the idle connections to the server just after the peak time.",
			},
			"nssrvrhdr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The server header value to be inserted. If no explicit header is specified then NSBUILD.RELEASE is used as default server header.",
			},

			// Read-only (GET-only) attributes surfaced by the data source (these
			// are intentionally NOT modeled on the resource). All Computed.
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Flag to determine if the http param is built-in or not. Possible values: [ MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL ]",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
		},
	}
}

// nshttpparamDataSourceSetAttrFromGet projects a NITRO nshttpparam GET response
// onto the data-source model. nshttpparam is a global singleton, so the ID is
// static. The shared utils.MapGet* helpers fill each attribute (or leave it Null
// when the GET omits it).
func nshttpparamDataSourceSetAttrFromGet(ctx context.Context, data *NshttpparamDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In nshttpparamDataSourceSetAttrFromGet Function")

	data.Conmultiplex = utils.MapGetString(g, "conmultiplex")
	data.Dropinvalreqs = utils.MapGetString(g, "dropinvalreqs")
	data.Http2serverside = utils.MapGetString(g, "http2serverside")
	data.Ignoreconnectcodingscheme = utils.MapGetString(g, "ignoreconnectcodingscheme")
	data.Insnssrvrhdr = utils.MapGetString(g, "insnssrvrhdr")
	data.Logerrresp = utils.MapGetString(g, "logerrresp")
	data.Markconnreqinval = utils.MapGetString(g, "markconnreqinval")
	data.Markhttp09inval = utils.MapGetString(g, "markhttp09inval")
	data.Maxreusepool = utils.MapGetInt64(g, "maxreusepool")
	data.Nssrvrhdr = utils.MapGetString(g, "nssrvrhdr")

	// Read-only attributes.
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")

	// nshttpparam has no unique lookup key; use a static ID.
	data.Id = types.StringValue("nshttpparam-config")
}
