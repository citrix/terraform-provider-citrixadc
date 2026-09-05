package sslparameter

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SslparameterDataSourceModel is the data-source-specific model, decoupled from
// SslparameterResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits.
// Every non-key attribute is Computed; the Framework's per-attribute model <->
// schema reflection requires this model to have exactly the attributes the
// data-source schema declares, which is why it cannot reuse the resource model.
type SslparameterDataSourceModel struct {
	Id                       types.String `tfsdk:"id"`
	Crlmemorysizemb          types.Int64  `tfsdk:"crlmemorysizemb"`
	Cryptodevdisablelimit    types.Int64  `tfsdk:"cryptodevdisablelimit"`
	Defaultprofile           types.String `tfsdk:"defaultprofile"`
	Denysslreneg             types.String `tfsdk:"denysslreneg"`
	Dropreqwithnohostheader  types.String `tfsdk:"dropreqwithnohostheader"`
	Encrypttriggerpktcount   types.Int64  `tfsdk:"encrypttriggerpktcount"`
	Heterogeneoussslhw       types.String `tfsdk:"heterogeneoussslhw"`
	Hybridfipsmode           types.String `tfsdk:"hybridfipsmode"`
	Insertcertspace          types.String `tfsdk:"insertcertspace"`
	Insertionencoding        types.String `tfsdk:"insertionencoding"`
	Ndcppcompliancecertcheck types.String `tfsdk:"ndcppcompliancecertcheck"`
	Ocspcachesize            types.Int64  `tfsdk:"ocspcachesize"`
	Operationqueuelimit      types.Int64  `tfsdk:"operationqueuelimit"`
	Pushenctriggertimeout    types.Int64  `tfsdk:"pushenctriggertimeout"`
	Pushflag                 types.Int64  `tfsdk:"pushflag"`
	Quantumsize              types.String `tfsdk:"quantumsize"`
	Sendclosenotify          types.String `tfsdk:"sendclosenotify"`
	Sigdigesttype            types.List   `tfsdk:"sigdigesttype"`
	Snihttphostmatch         types.String `tfsdk:"snihttphostmatch"`
	Softwarecryptothreshold  types.Int64  `tfsdk:"softwarecryptothreshold"`
	Sslierrorcache           types.String `tfsdk:"sslierrorcache"`
	Sslimaxerrorcachemem     types.Int64  `tfsdk:"sslimaxerrorcachemem"`
	Ssltriggertimeout        types.Int64  `tfsdk:"ssltriggertimeout"`
	Strictcachecks           types.String `tfsdk:"strictcachecks"`
	Undefactioncontrol       types.String `tfsdk:"undefactioncontrol"`
	Undefactiondata          types.String `tfsdk:"undefactiondata"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/sslparameter.json). Never settable; populated from GET.
	Svctls1112disable types.String `tfsdk:"svctls1112disable"`
	Montls1112disable types.String `tfsdk:"montls1112disable"`
}

func SslparameterDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"crlmemorysizemb": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum memory size to use for certificate revocation lists (CRLs). This parameter reserves memory for a CRL but sets a limit to the maximum memory that the CRLs loaded on the appliance can consume.",
			},
			"cryptodevdisablelimit": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Limit to the number of disabled SSL chips after which the ADC restarts. A value of zero implies that the ADC does not automatically restart.",
			},
			"defaultprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Global parameter used to enable default profile feature.",
			},
			"denysslreneg": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Deny renegotiation in specified circumstances. Available settings function as follows:\n* NO - Allow SSL renegotiation.\n* FRONTEND_CLIENT - Deny secure and nonsecure SSL renegotiation initiated by the client.\n* FRONTEND_CLIENTSERVER - Deny secure and nonsecure SSL renegotiation initiated by the client or the Citrix ADC during policy-based client authentication.\n* ALL - Deny all secure and nonsecure SSL renegotiation.\n* NONSECURE - Deny nonsecure SSL renegotiation. Allows only clients that support RFC 5746.",
			},
			"dropreqwithnohostheader": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Host header check for SNI enabled sessions. If this check is enabled and the HTTP request does not contain the host header for SNI enabled sessions(i.e vserver or profile bound to vserver has SNI enabled and 'Client Hello' arrived with SNI extension), the request is dropped.",
			},
			"encrypttriggerpktcount": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of queued packets after which encryption is triggered. Use this setting for SSL transactions that send small packets from server to Citrix ADC.",
			},
			"heterogeneoussslhw": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "To support both cavium and coleto based platforms in cluster environment, this mode has to be enabled.",
			},
			"hybridfipsmode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "When this mode is enabled, system will use additional crypto hardware to accelerate symmetric crypto operations.",
			},
			"insertcertspace": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "To insert space between lines in the certificate header of request",
			},
			"insertionencoding": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Encoding method used to insert the subject or issuer's name in HTTP requests to servers.",
			},
			"ndcppcompliancecertcheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Determines whether or not additional checks are carried out during a TLS handshake when validating an X.509 certificate received from the peer.\nSettings apply as follows:\nYES - (1) During certificate verification, ignore the\n          Common Name field (inside the subject name) if\n          Subject Alternative Name X.509 extension is present\n          in the certificate for backend connection.\n      (2) Verify the Extended Key Usage X.509 extension\n          server/client leaf certificate received over the wire\n          is consistent with the peer's role.\n          (applicable for frontend and backend connections)\n      (3) Verify the Basic Constraint CA field set to TRUE\n          for non-leaf certificates. (applicable for frontend,\n          backend connections and CAs bound to the Citrix ADC.\nNO  - (1) Verify the Common Name field (inside the subject name)\n          irrespective of Subject Alternative Name X.509\n          extension.\n      (2) Ignore the Extended Key Usage X.509 extension\n          for server/client leaf certificate.\n      (3) Do not verify the Basic Constraint CA true flag\n          for non-leaf certificates.",
			},
			"ocspcachesize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Size, per packet engine, in megabytes, of the OCSP cache. A maximum of 10% of the packet engine memory can be assigned. Because the maximum allowed packet engine memory is 4GB, the maximum value that can be assigned to the OCSP cache is approximately 410 MB.",
			},
			"operationqueuelimit": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Limit in percentage of capacity of the crypto operations queue beyond which new SSL connections are not accepted until the queue is reduced.",
			},
			"pushenctriggertimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "PUSH encryption trigger timeout value. The timeout value is applied only if you set the Push Encryption Trigger parameter to Timer in the SSL virtual server settings.",
			},
			"pushflag": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Insert PUSH flag into decrypted, encrypted, or all records. If the PUSH flag is set to a value other than 0, the buffered records are forwarded on the basis of the value of the PUSH flag. Available settings function as follows:\n0 - Auto (PUSH flag is not set.)\n1 - Insert PUSH flag into every decrypted record.\n2 -Insert PUSH flag into every encrypted record.\n3 - Insert PUSH flag into every decrypted and encrypted record.",
			},
			"quantumsize": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Amount of data to collect before the data is pushed to the crypto hardware for encryption. For large downloads, a larger quantum size better utilizes the crypto resources.",
			},
			"sendclosenotify": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Send an SSL Close-Notify message to the client at the end of a transaction.",
			},
			"sigdigesttype": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Signature Digest Algorithms that are supported by appliance. Default value is \"ALL\" and it will enable the following algorithms depending on the platform.\nOn VPX: ECDSA-SHA1 ECDSA-SHA224 ECDSA-SHA256 ECDSA-SHA384 ECDSA-SHA512 RSA-SHA1 RSA-SHA224 RSA-SHA256 RSA-SHA384 RSA-SHA512 DSA-SHA1 DSA-SHA224 DSA-SHA256 DSA-SHA384 DSA-SHA512\nOn MPX with Nitrox-III and coleto cards: RSA-SHA1 RSA-SHA224 RSA-SHA256 RSA-SHA384 RSA-SHA512 ECDSA-SHA1 ECDSA-SHA224 ECDSA-SHA256 ECDSA-SHA384 ECDSA-SHA512\nOthers: RSA-SHA1 RSA-SHA224 RSA-SHA256 RSA-SHA384 RSA-SHA512.\nNote:ALL doesnot include RSA-MD5 for any platform.",
			},
			"snihttphostmatch": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Controls how the HTTP 'Host' header value is validated. These checks are performed only if the session is SNI enabled (i.e when vserver or profile bound to vserver has SNI enabled and 'Client Hello' arrived with SNI extension) and HTTP request contains 'Host' header.\nAvailable settings function as follows:\nCERT   - Request is forwarded if the 'Host' value is covered\n         by the certificate used to establish this SSL session.\n         Note: 'CERT' matching mode cannot be applied in\n         TLS 1.3 connections established by resuming from a\n         previous TLS 1.3 session. On these connections, 'STRICT'\n         matching mode will be used instead.\nSTRICT - Request is forwarded only if value of 'Host' header\n         in HTTP is identical to the 'Server name' value passed\n         in 'Client Hello' of the SSL connection.\nNO     - No validation is performed on the HTTP 'Host'\n         header value.",
			},
			"softwarecryptothreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Citrix ADC CPU utilization threshold (in percentage) beyond which crypto operations are not done in software.\nA value of zero implies that CPU is not utilized for doing crypto in software.",
			},
			"sslierrorcache": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable dynamically learning and caching the learned information to make the subsequent interception or bypass decision. When enabled, NS does the lookup of this cached data to do early bypass.",
			},
			"sslimaxerrorcachemem": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the maximum memory that can be used for caching the learned data. This memory is used as a LRU cache so that the old entries gets replaced with new entry once the set memory limit is fully utilised. A value of 0 decides the limit automatically.",
			},
			"ssltriggertimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in milliseconds, after which encryption is triggered for transactions that are not tracked on the Citrix ADC because their length is not known. There can be a delay of up to 10ms from the specified timeout value before the packet is pushed into the queue.",
			},
			"strictcachecks": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable strict CA certificate checks on the appliance.",
			},
			"undefactioncontrol": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the undefined built-in control action: CLIENTAUTH, NOCLIENTAUTH, NOOP, RESET, or DROP.",
			},
			"undefactiondata": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the undefined built-in data action: NOOP, RESET or DROP.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (intentionally NOT modeled on the resource). All Computed.
			"svctls1112disable": schema.StringAttribute{
				Computed:    true,
				Description: "Disable TLS 1.1 and 1.2 for dynamic and VPN created services. Possible values = YES, NO.",
			},
			"montls1112disable": schema.StringAttribute{
				Computed:    true,
				Description: "Disable TLS 1.1 and 1.2 for secure (https) monitors bound to SSL_BRIDGE services. Possible values = YES, NO.",
			},
		},
	}
}

// sslparameterDataSourceSetAttrFromGet projects a NITRO sslparameter GET response
// onto the data-source model. sslparameter is a singleton (no lookup key), so the
// id is a fixed synthetic value. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them) via the shared utils.MapGet* helpers.
func sslparameterDataSourceSetAttrFromGet(ctx context.Context, data *SslparameterDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In sslparameterDataSourceSetAttrFromGet Function")

	// Singleton: fixed id matching the resource getter.
	data.Id = types.StringValue("sslparameter-config")

	// Read/write attributes as read-back outputs.
	data.Crlmemorysizemb = utils.MapGetInt64(g, "crlmemorysizemb")
	data.Cryptodevdisablelimit = utils.MapGetInt64(g, "cryptodevdisablelimit")
	data.Defaultprofile = utils.MapGetString(g, "defaultprofile")
	data.Denysslreneg = utils.MapGetString(g, "denysslreneg")
	data.Dropreqwithnohostheader = utils.MapGetString(g, "dropreqwithnohostheader")
	data.Encrypttriggerpktcount = utils.MapGetInt64(g, "encrypttriggerpktcount")
	data.Heterogeneoussslhw = utils.MapGetString(g, "heterogeneoussslhw")
	data.Hybridfipsmode = utils.MapGetString(g, "hybridfipsmode")
	data.Insertcertspace = utils.MapGetString(g, "insertcertspace")
	data.Insertionencoding = utils.MapGetString(g, "insertionencoding")
	data.Ndcppcompliancecertcheck = utils.MapGetString(g, "ndcppcompliancecertcheck")
	data.Ocspcachesize = utils.MapGetInt64(g, "ocspcachesize")
	data.Operationqueuelimit = utils.MapGetInt64(g, "operationqueuelimit")
	data.Pushenctriggertimeout = utils.MapGetInt64(g, "pushenctriggertimeout")
	data.Pushflag = utils.MapGetInt64(g, "pushflag")
	data.Quantumsize = utils.MapGetString(g, "quantumsize")
	data.Sendclosenotify = utils.MapGetString(g, "sendclosenotify")
	data.Sigdigesttype = utils.MapGetStringList(g, "sigdigesttype")
	data.Snihttphostmatch = utils.MapGetString(g, "snihttphostmatch")
	data.Softwarecryptothreshold = utils.MapGetInt64(g, "softwarecryptothreshold")
	data.Sslierrorcache = utils.MapGetString(g, "sslierrorcache")
	data.Sslimaxerrorcachemem = utils.MapGetInt64(g, "sslimaxerrorcachemem")
	data.Ssltriggertimeout = utils.MapGetInt64(g, "ssltriggertimeout")
	data.Strictcachecks = utils.MapGetString(g, "strictcachecks")
	data.Undefactioncontrol = utils.MapGetString(g, "undefactioncontrol")
	data.Undefactiondata = utils.MapGetString(g, "undefactiondata")

	// Read-only attributes.
	data.Svctls1112disable = utils.MapGetString(g, "svctls1112disable")
	data.Montls1112disable = utils.MapGetString(g, "montls1112disable")
}
