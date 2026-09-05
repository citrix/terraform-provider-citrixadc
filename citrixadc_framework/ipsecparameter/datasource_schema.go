package ipsecparameter

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// IpsecparameterDataSourceModel is the data-source-specific model, decoupled from
// IpsecparameterResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits. Every
// non-key attribute is Computed; the Framework's per-attribute model <-> schema
// reflection requires this model to have exactly the attributes the data-source
// schema declares, which is why it cannot reuse the resource model. ipsecparameter
// is a singleton (no lookup key); the ID is static.
type IpsecparameterDataSourceModel struct {
	Id                    types.String `tfsdk:"id"`
	Encalgo               types.List   `tfsdk:"encalgo"`
	Hashalgo              types.List   `tfsdk:"hashalgo"`
	Ikeretryinterval      types.Int64  `tfsdk:"ikeretryinterval"`
	Ikeversion            types.String `tfsdk:"ikeversion"`
	Lifetime              types.Int64  `tfsdk:"lifetime"`
	Livenesscheckinterval types.Int64  `tfsdk:"livenesscheckinterval"`
	Perfectforwardsecrecy types.String `tfsdk:"perfectforwardsecrecy"`
	Replaywindowsize      types.Int64  `tfsdk:"replaywindowsize"`
	Retransmissiontime    types.Int64  `tfsdk:"retransmissiontime"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/ipsecparameter.json). Never settable; populated from GET.
	Responderonly types.String `tfsdk:"responderonly"`
}

func IpsecparameterDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"encalgo": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Type of encryption algorithm (Note: Selection of AES enables AES128)",
			},
			"hashalgo": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Type of hashing algorithm",
			},
			"ikeretryinterval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "IKE retry interval for bringing up the connection",
			},
			"ikeversion": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IKE Protocol Version",
			},
			"lifetime": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Lifetime of IKE SA in seconds. Lifetime of IPSec SA will be (lifetime of IKE SA/8)",
			},
			"livenesscheckinterval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of seconds after which a notify payload is sent to check the liveliness of the peer. Additional retries are done as per retransmit interval setting. Zero value disables liveliness checks.",
			},
			"perfectforwardsecrecy": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable/Disable PFS.",
			},
			"replaywindowsize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "IPSec Replay window size for the data traffic",
			},
			"retransmissiontime": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The interval in seconds to retry sending the IKE messages to peer, three consecutive attempts are done with doubled interval after every failure,\nincreases for every retransmit till 6 retransmits.",
			},

			// Read-only (GET-only) metadata surfaced by the data source (this is
			// intentionally NOT modeled on the resource). Computed; null when the
			// appliance omits it.
			"responderonly": schema.StringAttribute{
				Computed:    true,
				Description: "Responder Only config for IKED.",
			},
		},
	}
}

// ipsecparameterDataSourceSetAttrFromGet projects a NITRO ipsecparameter GET
// response onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that projection.
func ipsecparameterDataSourceSetAttrFromGet(ctx context.Context, data *IpsecparameterDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In ipsecparameterDataSourceSetAttrFromGet Function")

	// Read/write attributes as read-back outputs.
	data.Encalgo = utils.MapGetStringList(g, "encalgo")
	data.Hashalgo = utils.MapGetStringList(g, "hashalgo")
	data.Ikeretryinterval = utils.MapGetInt64(g, "ikeretryinterval")
	data.Ikeversion = utils.MapGetString(g, "ikeversion")
	data.Lifetime = utils.MapGetInt64(g, "lifetime")
	data.Livenesscheckinterval = utils.MapGetInt64(g, "livenesscheckinterval")
	data.Perfectforwardsecrecy = utils.MapGetString(g, "perfectforwardsecrecy")
	data.Replaywindowsize = utils.MapGetInt64(g, "replaywindowsize")
	data.Retransmissiontime = utils.MapGetInt64(g, "retransmissiontime")

	// Read-only metadata.
	data.Responderonly = utils.MapGetString(g, "responderonly")

	// Singleton resource with no unique attributes - static ID.
	data.Id = types.StringValue("ipsecparameter-config")
}
