package quicprofile

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// QuicprofileDataSourceModel is the data-source-specific model, decoupled from
// QuicprofileResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits (refcnt,
// builtin, feature). Every non-key attribute is Computed; the Framework's
// per-attribute model <-> schema reflection requires this model to have exactly
// the attributes the data-source schema declares, which is why it cannot reuse
// the resource model.
type QuicprofileDataSourceModel struct {
	Id                             types.String `tfsdk:"id"`
	Ackdelayexponent               types.Int64  `tfsdk:"ackdelayexponent"`
	Activeconnectionidlimit        types.Int64  `tfsdk:"activeconnectionidlimit"`
	Activeconnectionmigration      types.String `tfsdk:"activeconnectionmigration"`
	Congestionctrlalgorithm        types.String `tfsdk:"congestionctrlalgorithm"`
	Initialmaxdata                 types.Int64  `tfsdk:"initialmaxdata"`
	Initialmaxstreamdatabidilocal  types.Int64  `tfsdk:"initialmaxstreamdatabidilocal"`
	Initialmaxstreamdatabidiremote types.Int64  `tfsdk:"initialmaxstreamdatabidiremote"`
	Initialmaxstreamdatauni        types.Int64  `tfsdk:"initialmaxstreamdatauni"`
	Initialmaxstreamsbidi          types.Int64  `tfsdk:"initialmaxstreamsbidi"`
	Initialmaxstreamsuni           types.Int64  `tfsdk:"initialmaxstreamsuni"`
	Maxackdelay                    types.Int64  `tfsdk:"maxackdelay"`
	Maxidletimeout                 types.Int64  `tfsdk:"maxidletimeout"`
	Maxudpdatagramsperburst        types.Int64  `tfsdk:"maxudpdatagramsperburst"`
	Maxudppayloadsize              types.Int64  `tfsdk:"maxudppayloadsize"`
	Name                           types.String `tfsdk:"name"` // Required lookup key
	Newtokenvalidityperiod         types.Int64  `tfsdk:"newtokenvalidityperiod"`
	Retrytokenvalidityperiod       types.Int64  `tfsdk:"retrytokenvalidityperiod"`
	Statelessaddressvalidation     types.String `tfsdk:"statelessaddressvalidation"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/quicprofile.json). Never settable; populated from GET.
	Refcnt  types.Int64  `tfsdk:"refcnt"`
	Builtin types.List   `tfsdk:"builtin"`
	Feature types.String `tfsdk:"feature"`
}

func QuicprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"ackdelayexponent": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "An integer value advertised by the Citrix ADC to the remote QUIC endpoint, indicating an exponent that the remote QUIC endpoint should use, to decode the ACK Delay field in QUIC ACK frames sent by the Citrix ADC.",
			},
			"activeconnectionidlimit": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "An integer value advertised by the Citrix ADC to the remote QUIC endpoint, specifying the maximum number of QUIC connection IDs from the remote QUIC endpoint, that the Citrix ADC is willing to store.",
			},
			"activeconnectionmigration": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify whether the Citrix ADC should allow the remote QUIC endpoint to perform active QUIC connection migration.",
			},
			"congestionctrlalgorithm": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the congestion control algorithm to be used for QUIC connections. The default congestion control algorithm is CUBIC.",
			},
			"initialmaxdata": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "An integer value advertised by the Citrix ADC to the remote QUIC endpoint, specifying the initial value, in bytes, for the maximum amount of data that can be sent on a QUIC connection.",
			},
			"initialmaxstreamdatabidilocal": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "An integer value advertised by the Citrix ADC to the remote QUIC endpoint, specifying the initial flow control limit, in bytes, for bidirectional QUIC streams initiated by the Citrix ADC.",
			},
			"initialmaxstreamdatabidiremote": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "An integer value advertised by the Citrix ADC to the remote QUIC endpoint, specifying the initial flow control limit, in bytes, for bidirectional QUIC streams initiated by the remote QUIC endpoint.",
			},
			"initialmaxstreamdatauni": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "An integer value advertised by the Citrix ADC to the remote QUIC endpoint, specifying the initial flow control limit, in bytes, for unidirectional streams initiated by the remote QUIC endpoint.",
			},
			"initialmaxstreamsbidi": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "An integer value advertised by the Citrix ADC to the remote QUIC endpoint, specifying the initial maximum number of bidirectional streams the remote QUIC endpoint may initiate.",
			},
			"initialmaxstreamsuni": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "An integer value advertised by the Citrix ADC to the remote QUIC endpoint, specifying the initial maximum number of unidirectional streams the remote QUIC endpoint may initiate.",
			},
			"maxackdelay": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "An integer value advertised by the Citrix ADC to the remote QUIC endpoint, specifying the maximum amount of time, in milliseconds, by which the Citrix ADC will delay sending acknowledgments.",
			},
			"maxidletimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "An integer value advertised by the Citrix ADC to the remote QUIC endpoint, specifying the maximum idle timeout, in seconds, for a QUIC connection. A QUIC connection will be silently discarded by the Citrix ADC if it remains idle for longer than the minimum of the idle timeout values advertised by the Citrix ADC and the remote QUIC endpoint, and three times the current Probe Timeout (PTO).",
			},
			"maxudpdatagramsperburst": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "An integer value, specifying the maximum number of UDP datagrams that can be transmitted by the Citrix ADC in a single transmission burst on a QUIC connection.",
			},
			"maxudppayloadsize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "An integer value advertised by the Citrix ADC to the remote QUIC endpoint, specifying the size of the largest UDP datagram payload, in bytes, that the Citrix ADC is willing to receive on a QUIC connection.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the QUIC profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@),equals sign (=), and hyphen (-) characters. Cannot be changed after the profile is created.",
			},
			"newtokenvalidityperiod": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "An integer value, specifying the validity period, in seconds, of address validation tokens issued through QUIC NEW_TOKEN frames sent by the Citrix ADC.",
			},
			"retrytokenvalidityperiod": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "An integer value, specifying the validity period, in seconds, of address validation tokens issued through QUIC Retry packets sent by the Citrix ADC.",
			},
			"statelessaddressvalidation": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify whether the Citrix ADC should perform stateless address validation for QUIC clients, by sending tokens in QUIC Retry packets during QUIC connection establishment, and by sending tokens in QUIC NEW_TOKEN frames after QUIC connection establishment.",
			},

			// Read-only (GET-only) attributes surfaced by the data source (these
			// are intentionally NOT modeled on the resource). All Computed.
			"refcnt": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of entities using this profile.",
			},
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Flag to determine if the QUIC profile is built-in or not. Possible values = MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
		},
	}
}

// quicprofileDataSourceSetAttrFromGet projects a NITRO quicprofile GET response
// onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func quicprofileDataSourceSetAttrFromGet(ctx context.Context, data *QuicprofileDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In quicprofileDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Ackdelayexponent = utils.MapGetInt64(g, "ackdelayexponent")
	data.Activeconnectionidlimit = utils.MapGetInt64(g, "activeconnectionidlimit")
	data.Activeconnectionmigration = utils.MapGetString(g, "activeconnectionmigration")
	data.Congestionctrlalgorithm = utils.MapGetString(g, "congestionctrlalgorithm")
	data.Initialmaxdata = utils.MapGetInt64(g, "initialmaxdata")
	data.Initialmaxstreamdatabidilocal = utils.MapGetInt64(g, "initialmaxstreamdatabidilocal")
	data.Initialmaxstreamdatabidiremote = utils.MapGetInt64(g, "initialmaxstreamdatabidiremote")
	data.Initialmaxstreamdatauni = utils.MapGetInt64(g, "initialmaxstreamdatauni")
	data.Initialmaxstreamsbidi = utils.MapGetInt64(g, "initialmaxstreamsbidi")
	data.Initialmaxstreamsuni = utils.MapGetInt64(g, "initialmaxstreamsuni")
	data.Maxackdelay = utils.MapGetInt64(g, "maxackdelay")
	data.Maxidletimeout = utils.MapGetInt64(g, "maxidletimeout")
	data.Maxudpdatagramsperburst = utils.MapGetInt64(g, "maxudpdatagramsperburst")
	data.Maxudppayloadsize = utils.MapGetInt64(g, "maxudppayloadsize")
	data.Newtokenvalidityperiod = utils.MapGetInt64(g, "newtokenvalidityperiod")
	data.Retrytokenvalidityperiod = utils.MapGetInt64(g, "retrytokenvalidityperiod")
	data.Statelessaddressvalidation = utils.MapGetString(g, "statelessaddressvalidation")

	// Read-only attributes.
	data.Refcnt = utils.MapGetInt64(g, "refcnt")
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
}
