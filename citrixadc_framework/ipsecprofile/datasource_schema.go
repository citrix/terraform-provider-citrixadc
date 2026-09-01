package ipsecprofile

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// IpsecprofileDataSourceModel is the data-source-specific model, decoupled from
// IpsecprofileResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits. Every
// non-key attribute is Computed; the Framework's per-attribute model <-> schema
// reflection requires this model to have exactly the attributes the data-source
// schema declares, which is why it cannot reuse the resource model.
type IpsecprofileDataSourceModel struct {
	Id                    types.String `tfsdk:"id"`
	Encalgo               types.List   `tfsdk:"encalgo"`
	Hashalgo              types.List   `tfsdk:"hashalgo"`
	Ikeretryinterval      types.Int64  `tfsdk:"ikeretryinterval"`
	Ikeversion            types.String `tfsdk:"ikeversion"`
	Lifetime              types.Int64  `tfsdk:"lifetime"`
	Livenesscheckinterval types.Int64  `tfsdk:"livenesscheckinterval"`
	Name                  types.String `tfsdk:"name"` // Required lookup key
	Peerpublickey         types.String `tfsdk:"peerpublickey"`
	Perfectforwardsecrecy types.String `tfsdk:"perfectforwardsecrecy"`
	Privatekey            types.String `tfsdk:"privatekey"`
	Psk                   types.String `tfsdk:"psk"`
	PskWo                 types.String `tfsdk:"psk_wo"`
	PskWoVersion          types.Int64  `tfsdk:"psk_wo_version"`
	Publickey             types.String `tfsdk:"publickey"`
	Replaywindowsize      types.Int64  `tfsdk:"replaywindowsize"`
	Retransmissiontime    types.Int64  `tfsdk:"retransmissiontime"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/ipsecprofile.json). Never settable; populated from GET.
	Responderonly types.String `tfsdk:"responderonly"`
	Builtin       types.List   `tfsdk:"builtin"`
	Feature       types.String `tfsdk:"feature"`
}

func IpsecprofileDataSourceSchema() schema.Schema {
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
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the ipsec profile",
			},
			"peerpublickey": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Peer public key file path",
			},
			"perfectforwardsecrecy": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable/Disable PFS.",
			},
			"privatekey": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Private key file path",
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
			"publickey": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Public key file path",
			},
			"replaywindowsize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "IPSec Replay window size for the data traffic",
			},
			"retransmissiontime": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The interval in seconds to retry sending the IKE messages to peer, three consecutive attempts are done with doubled interval after every failure.",
			},

			// Read-only (GET-only) metadata surfaced by the data source (these are
			// intentionally NOT modeled on the resource). All Computed; null when the
			// appliance omits them.
			"responderonly": schema.StringAttribute{
				Computed:    true,
				Description: "Responder Only config for IKED. Possible values: YES, NO.",
			},
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Indicates that a variable is a built-in (SYSTEM INTERNAL) type.",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
		},
	}
}

// ipsecprofileDataSourceSetAttrFromGet projects a NITRO ipsecprofile GET response
// onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that projection.
func ipsecprofileDataSourceSetAttrFromGet(ctx context.Context, data *IpsecprofileDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In ipsecprofileDataSourceSetAttrFromGet Function")

	// Named resource keyed on "name"; the ID is the plain name value.
	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Encalgo = utils.MapGetStringList(g, "encalgo")
	data.Hashalgo = utils.MapGetStringList(g, "hashalgo")
	data.Ikeretryinterval = utils.MapGetInt64(g, "ikeretryinterval")
	data.Ikeversion = utils.MapGetString(g, "ikeversion")
	data.Lifetime = utils.MapGetInt64(g, "lifetime")
	data.Livenesscheckinterval = utils.MapGetInt64(g, "livenesscheckinterval")
	data.Peerpublickey = utils.MapGetString(g, "peerpublickey")
	data.Perfectforwardsecrecy = utils.MapGetString(g, "perfectforwardsecrecy")
	data.Privatekey = utils.MapGetString(g, "privatekey")
	data.Publickey = utils.MapGetString(g, "publickey")
	data.Replaywindowsize = utils.MapGetInt64(g, "replaywindowsize")
	data.Retransmissiontime = utils.MapGetInt64(g, "retransmissiontime")

	// psk / psk_wo / psk_wo_version are write-only or action-only inputs the GET
	// never returns -> Null.
	data.Psk = types.StringNull()
	data.PskWo = types.StringNull()
	data.PskWoVersion = types.Int64Null()

	// Read-only metadata.
	data.Responderonly = utils.MapGetString(g, "responderonly")
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
}
