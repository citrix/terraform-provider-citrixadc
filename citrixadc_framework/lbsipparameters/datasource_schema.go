package lbsipparameters

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// LbsipparametersDataSourceModel is the data-source-specific model, decoupled
// from LbsipparametersResourceModel. A data source is a pure read surface, so it
// can expose the full GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits.
type LbsipparametersDataSourceModel struct {
	Id                  types.String `tfsdk:"id"`
	Addrportvip         types.String `tfsdk:"addrportvip"`
	Retrydur            types.Int64  `tfsdk:"retrydur"`
	Rnatdstport         types.Int64  `tfsdk:"rnatdstport"`
	Rnatsecuredstport   types.Int64  `tfsdk:"rnatsecuredstport"`
	Rnatsecuresrcport   types.Int64  `tfsdk:"rnatsecuresrcport"`
	Rnatsrcport         types.Int64  `tfsdk:"rnatsrcport"`
	Sip503ratethreshold types.Int64  `tfsdk:"sip503ratethreshold"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/lbsipparameters.json). Never settable; populated from GET.
	Builtin types.List   `tfsdk:"builtin"`
	Feature types.String `tfsdk:"feature"`
}

func LbsipparametersDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"addrportvip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Add the rport parameter to the VIA headers of SIP requests that virtual servers receive from clients or servers.",
			},
			"retrydur": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in seconds, for which a client must wait before initiating a connection after receiving a 503 Service Unavailable response from the SIP server. The time value is sent in the \"Retry-After\" header in the 503 response.",
			},
			"rnatdstport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number with which to match the destination port in server-initiated SIP traffic. The rport parameter is added, without a value, to SIP packets that have a matching destination port number, and CALL-ID based persistence is implemented for the responses received by the virtual server.",
			},
			"rnatsecuredstport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number with which to match the destination port in server-initiated SIP over SSL traffic. The rport parameter is added, without a value, to SIP packets that have a matching destination port number, and CALL-ID based persistence is implemented for the responses received by the virtual server.",
			},
			"rnatsecuresrcport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number with which to match the source port in server-initiated SIP over SSL traffic. The rport parameter is added, without a value, to SIP packets that have a matching source port number, and CALL-ID based persistence is implemented for the responses received by the virtual server.",
			},
			"rnatsrcport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number with which to match the source port in server-initiated SIP traffic. The rport parameter is added, without a value, to SIP packets that have a matching source port number, and CALL-ID based persistence is implemented for the responses received by the virtual server.",
			},
			"sip503ratethreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of 503 Service Unavailable responses to generate, once every 10 milliseconds, when a SIP virtual server becomes unavailable.",
			},

			// Read-only (GET-only) attributes surfaced by the data source (these
			// are intentionally NOT modeled on the resource). All Computed.
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Flag to determine if SIP param is built-in or not. Possible values = MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL.",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
		},
	}
}

// lbsipparametersDataSourceSetAttrFromGet projects a NITRO lbsipparameters GET
// response onto the data-source model. The shared utils.MapGet* helpers fill
// each attribute from the GET (or leave it Null when the GET omits it).
func lbsipparametersDataSourceSetAttrFromGet(ctx context.Context, data *LbsipparametersDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In lbsipparametersDataSourceSetAttrFromGet Function")

	// Read/write attributes as read-back outputs.
	data.Addrportvip = utils.MapGetString(g, "addrportvip")
	data.Retrydur = utils.MapGetInt64(g, "retrydur")
	data.Rnatdstport = utils.MapGetInt64(g, "rnatdstport")
	data.Rnatsecuredstport = utils.MapGetInt64(g, "rnatsecuredstport")
	data.Rnatsecuresrcport = utils.MapGetInt64(g, "rnatsecuresrcport")
	data.Rnatsrcport = utils.MapGetInt64(g, "rnatsrcport")
	data.Sip503ratethreshold = utils.MapGetInt64(g, "sip503ratethreshold")

	// Read-only attributes.
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")

	// Set ID. Singleton config -> static ID.
	data.Id = types.StringValue("lbsipparameters-config")
}
