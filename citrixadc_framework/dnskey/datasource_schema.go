package dnskey

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// DnskeyDataSourceModel is the data-source-specific model, decoupled from
// DnskeyResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits
// (state, type, tag, timestamps, rollover fail rc). The Framework's
// per-attribute model <-> schema reflection requires this model to have exactly
// the attributes the data-source schema declares.
type DnskeyDataSourceModel struct {
	Id                 types.String `tfsdk:"id"`
	Keyname            types.String `tfsdk:"keyname"` // Required lookup key
	Algorithm          types.String `tfsdk:"algorithm"`
	Autorollover       types.String `tfsdk:"autorollover"`
	Expires            types.Int64  `tfsdk:"expires"`
	Filenameprefix     types.String `tfsdk:"filenameprefix"`
	Keysize            types.Int64  `tfsdk:"keysize"`
	Keytype            types.String `tfsdk:"keytype"`
	Notificationperiod types.Int64  `tfsdk:"notificationperiod"`
	Password           types.String `tfsdk:"password"`
	Privatekey         types.String `tfsdk:"privatekey"`
	Publickey          types.String `tfsdk:"publickey"`
	Revoke             types.Bool   `tfsdk:"revoke"`
	Rollovermethod     types.String `tfsdk:"rollovermethod"`
	Src                types.String `tfsdk:"src"`
	Ttl                types.Int64  `tfsdk:"ttl"`
	Units1             types.String `tfsdk:"units1"`
	Units2             types.String `tfsdk:"units2"`
	Zonename           types.String `tfsdk:"zonename"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/dnskey.json). Never settable; populated from GET.
	State             types.String `tfsdk:"state"`
	Type              types.String `tfsdk:"type"`
	Tag               types.Int64  `tfsdk:"tag"`
	Createtimestr     types.String `tfsdk:"createtimestr"`
	Activationtimestr types.String `tfsdk:"activationtimestr"`
	Expirytimestr     types.String `tfsdk:"expirytimestr"`
	Deletiontimestr   types.String `tfsdk:"deletiontimestr"`
	Rolloverfailrc    types.Int64  `tfsdk:"rolloverfailrc"`
}

func DnskeyDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"algorithm": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Algorithm to generate the key.",
			},
			"autorollover": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Flag to enable/disable key rollover automatically.\nNote:\n* Key name will be appended with _AR1 for successor key. For e.g. current key=k1, successor key=k1_AR1.\n* Key name can be truncated if current name length is more than 58 bytes to accomodate the suffix.",
			},
			"expires": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time period for which to consider the key valid, after the key is used to sign a zone.",
			},
			"filenameprefix": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Common prefix for the names of the generated public and private key files and the Delegation Signer (DS) resource record. During key generation, the .key, .private, and .ds suffixes are appended automatically to the file name prefix to produce the names of the public key, the private key, and the DS record, respectively.",
			},
			"keyname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the public-private key pair to publish in the zone.",
			},
			"keysize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Size of the key, in bits.",
			},
			"keytype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of key to create.",
			},
			"notificationperiod": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time at which to generate notification of key expiration, specified as number of days, hours, or minutes before expiry. Must be less than the expiry period. The notification is an SNMP trap sent to an SNMP manager. To enable the appliance to send the trap, enable the DNSKEY-EXPIRY SNMP alarm. \nIn case autorollover option is enabled, rollover for successor key will be intiated at this time. No notification trap will be sent.",
			},
			"password": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
				Description: "Passphrase for reading the encrypted public/private DNS keys",
			},
			"privatekey": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "File name of the private key.",
			},
			"publickey": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "File name of the public key.",
			},
			"revoke": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Revoke the key. Note: This operation is non-reversible.",
			},
			"rollovermethod": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Method used for automatic rollover.\n* Key type: ZSK, Method: PrePublication or DoubleSignature.\n* Key type: KSK, Method: DoubleRRSet.",
			},
			"src": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL (protocol, host, path, and file name) from where the DNS key file will be imported. NOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access. This is a mandatory argument",
			},
			"ttl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time to Live (TTL), in seconds, for the DNSKEY resource record created in the zone. TTL is the time for which the record must be cached by the DNS proxies. If the TTL is not specified, either the DNS zone's minimum TTL or the default value of 3600 is used.",
			},
			"units1": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Units for the expiry period.",
			},
			"units2": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Units for the notification period.",
			},
			"zonename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the zone for which to create a key.",
			},

			// Read-only (GET-only) attributes surfaced by the data source (these
			// are intentionally NOT modeled on the resource). All Computed.
			"state": schema.StringAttribute{
				Computed:    true,
				Description: "Current key state. Possible values = Created, Activated, Deactivated, Revoked.",
			},
			"type": schema.StringAttribute{
				Computed:    true,
				Description: "Key type. Possible values = KSK, KeySigningKey, ZSK, ZoneSigningKey.",
			},
			"tag": schema.Int64Attribute{
				Computed:    true,
				Description: "Key tag/ID.",
			},
			"createtimestr": schema.StringAttribute{
				Computed:    true,
				Description: "Key creation time.",
			},
			"activationtimestr": schema.StringAttribute{
				Computed:    true,
				Description: "Key activation time.",
			},
			"expirytimestr": schema.StringAttribute{
				Computed:    true,
				Description: "Key expiry time.",
			},
			"deletiontimestr": schema.StringAttribute{
				Computed:    true,
				Description: "Key deletion time if autorollover option is enabled.",
			},
			"rolloverfailrc": schema.Int64Attribute{
				Computed:    true,
				Description: "Reason code in case rollover event failed.",
			},
		},
	}
}

// dnskeyDataSourceSetAttrFromGet projects a NITRO dnskey GET response onto the
// data-source model. Because a data source has no plan/apply reconciliation,
// attributes are simply filled from the GET (or left Null when the GET omits
// them). The shared utils.MapGet* helpers implement that projection.
func dnskeyDataSourceSetAttrFromGet(ctx context.Context, data *DnskeyDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In dnskeyDataSourceSetAttrFromGet Function")

	if v, ok := g["keyname"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Keyname = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Algorithm = utils.MapGetString(g, "algorithm")
	data.Autorollover = utils.MapGetString(g, "autorollover")
	data.Expires = utils.MapGetInt64(g, "expires")
	data.Filenameprefix = utils.MapGetString(g, "filenameprefix")
	data.Keysize = utils.MapGetInt64(g, "keysize")
	data.Keytype = utils.MapGetString(g, "keytype")
	data.Notificationperiod = utils.MapGetInt64(g, "notificationperiod")
	data.Privatekey = utils.MapGetString(g, "privatekey")
	data.Publickey = utils.MapGetString(g, "publickey")
	data.Revoke = utils.MapGetBool(g, "revoke")
	data.Rollovermethod = utils.MapGetString(g, "rollovermethod")
	data.Src = utils.MapGetString(g, "src")
	data.Ttl = utils.MapGetInt64(g, "ttl")
	data.Units1 = utils.MapGetString(g, "units1")
	data.Units2 = utils.MapGetString(g, "units2")
	data.Zonename = utils.MapGetString(g, "zonename")

	// password is a secret input the GET never returns -> Null.
	data.Password = types.StringNull()

	// Read-only attributes.
	data.State = utils.MapGetString(g, "state")
	data.Type = utils.MapGetString(g, "type")
	data.Tag = utils.MapGetInt64(g, "tag")
	data.Createtimestr = utils.MapGetString(g, "createtimestr")
	data.Activationtimestr = utils.MapGetString(g, "activationtimestr")
	data.Expirytimestr = utils.MapGetString(g, "expirytimestr")
	data.Deletiontimestr = utils.MapGetString(g, "deletiontimestr")
	data.Rolloverfailrc = utils.MapGetInt64(g, "rolloverfailrc")
}
