package sslcertkey

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SslCertKeyDataSourceModel is the data-source-specific model, decoupled from
// SslCertKeyResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only cert-metadata attributes that the resource
// deliberately omits (serial, issuer, subject, daystoexpiration, status, SANs,
// ...). Every non-key attribute is Computed; the Framework's per-attribute model
// <-> schema reflection requires this model to have exactly the attributes the
// data-source schema declares, which is why it cannot reuse the resource model.
type SslCertKeyDataSourceModel struct {
	Id      types.String `tfsdk:"id"`
	Certkey types.String `tfsdk:"certkey"` // Required lookup key

	// Read/write attributes, surfaced here as Computed outputs.
	Cert                        types.String `tfsdk:"cert"`
	Key                         types.String `tfsdk:"key"`
	Password                    types.Bool   `tfsdk:"password"`
	Fipskey                     types.String `tfsdk:"fipskey"`
	Hsmkey                      types.String `tfsdk:"hsmkey"`
	Inform                      types.String `tfsdk:"inform"`
	Expirymonitor               types.String `tfsdk:"expirymonitor"`
	NotificationPeriod          types.Int64  `tfsdk:"notificationperiod"`
	Bundle                      types.String `tfsdk:"bundle"`
	LinkCertKeyName             types.String `tfsdk:"linkcertkeyname"`
	NoDomainCheck               types.Bool   `tfsdk:"nodomaincheck"`
	OcspStaplingCache           types.Bool   `tfsdk:"ocspstaplingcache"`
	DeleteFromDevice            types.Bool   `tfsdk:"deletefromdevice"`
	DeleteCertKeyFilesOnRemoval types.String `tfsdk:"deletecertkeyfilesonremoval"`
	Passplain                   types.String `tfsdk:"passplain"`
	CertHash                    types.String `tfsdk:"cert_hash"`
	KeyHash                     types.String `tfsdk:"key_hash"`

	// Read-only (GET-only) cert metadata from the NITRO doc read-only set
	// (zion73x_readonly/sslcertkey.json). Never settable; populated from GET.
	Signaturealg        types.String `tfsdk:"signaturealg"`
	Certificatetype     types.List   `tfsdk:"certificatetype"`
	Serial              types.String `tfsdk:"serial"`
	Issuer              types.String `tfsdk:"issuer"`
	Clientcertnotbefore types.String `tfsdk:"clientcertnotbefore"`
	Clientcertnotafter  types.String `tfsdk:"clientcertnotafter"`
	Daystoexpiration    types.Int64  `tfsdk:"daystoexpiration"`
	Subject             types.String `tfsdk:"subject"`
	Publickey           types.String `tfsdk:"publickey"`
	Publickeysize       types.Int64  `tfsdk:"publickeysize"`
	Version             types.Int64  `tfsdk:"version"`
	Priority            types.Int64  `tfsdk:"priority"`
	Status              types.String `tfsdk:"status"`
	Passcrypt           types.String `tfsdk:"passcrypt"`
	Data                types.Int64  `tfsdk:"data"`
	Servicename         types.String `tfsdk:"servicename"`
	Sandns              types.String `tfsdk:"sandns"`
	Sanipadd            types.String `tfsdk:"sanipadd"`
	Ocspresponsestatus  types.String `tfsdk:"ocspresponsestatus"`
	Builtin             types.List   `tfsdk:"builtin"`
	Feature             types.String `tfsdk:"feature"`
	Certkeydigest       types.String `tfsdk:"certkeydigest"`
	Certificatesource   types.String `tfsdk:"certificatesource"`
	Certkeystatus       types.String `tfsdk:"certkeystatus"`
}

func SslCertKeyDataSourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Data source to read SSL certificate key pair configuration.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the SSL certificate key pair.",
			},
			"certkey": schema.StringAttribute{
				Required:    true,
				Description: "Name of the certificate and private-key pair to read.",
			},
			"cert": schema.StringAttribute{
				Computed:    true,
				Description: "Name of and path to the X509 certificate file.",
			},
			"key": schema.StringAttribute{
				Computed:    true,
				Description: "Name of and path to the private-key file.",
			},
			"password": schema.BoolAttribute{
				Computed:    true,
				Description: "Passphrase that was used to encrypt the private-key.",
			},
			"fipskey": schema.StringAttribute{
				Computed:    true,
				Description: "Name of the FIPS key in the Hardware Security Module (HSM).",
			},
			"hsmkey": schema.StringAttribute{
				Computed:    true,
				Description: "Name of the HSM key in the External Hardware Security Module (HSM).",
			},
			"inform": schema.StringAttribute{
				Computed:    true,
				Description: "Input format of the certificate and the private-key files (PEM, DER, or PFX).",
			},
			"expirymonitor": schema.StringAttribute{
				Computed:    true,
				Description: "Issue an alert when the certificate is about to expire.",
			},
			"notificationperiod": schema.Int64Attribute{
				Computed:    true,
				Description: "Time, in days, before certificate expiration at which to generate an alert.",
			},
			"bundle": schema.StringAttribute{
				Computed:    true,
				Description: "Parse the certificate chain as a single file.",
			},
			"linkcertkeyname": schema.StringAttribute{
				Computed:    true,
				Description: "Name of the Certificate Authority certificate-key pair linked to this certificate.",
			},
			"nodomaincheck": schema.BoolAttribute{
				Computed:    true,
				Description: "Override the check for matching domain names during certificate update.",
			},
			"ocspstaplingcache": schema.BoolAttribute{
				Computed:    true,
				Description: "Clear cached ocspStapling response.",
			},
			"deletefromdevice": schema.BoolAttribute{
				Computed:    true,
				Description: "Delete cert/key file from file system.",
			},
			"deletecertkeyfilesonremoval": schema.StringAttribute{
				Computed:    true,
				Description: "Delete certificate and key files when the certificate is removed.",
			},
			"passplain": schema.StringAttribute{
				Sensitive:   true,
				Computed:    true,
				Description: "Pass phrase used to encrypt the private-key. Required when adding an encrypted private-key in PEM format.",
			},
			"cert_hash": schema.StringAttribute{
				Optional:    true,
				Description: "SHA256 hash of the certificate.",
			},
			"key_hash": schema.StringAttribute{
				Optional:    true,
				Description: "SHA256 hash of the private key.",
			},

			// Read-only (GET-only) cert metadata surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"signaturealg": schema.StringAttribute{
				Computed:    true,
				Description: "Algorithm used to sign the certificate.",
			},
			"certificatetype": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Certificate type (e.g. ROOT_CERT, INTERM_CERT, SERVER_CERT, CLIENT_CERT).",
			},
			"serial": schema.StringAttribute{
				Computed:    true,
				Description: "Serial number of the certificate.",
			},
			"issuer": schema.StringAttribute{
				Computed:    true,
				Description: "Distinguished name of the certificate issuer.",
			},
			"clientcertnotbefore": schema.StringAttribute{
				Computed:    true,
				Description: "Date and time from which the certificate is valid.",
			},
			"clientcertnotafter": schema.StringAttribute{
				Computed:    true,
				Description: "Date and time after which the certificate expires.",
			},
			"daystoexpiration": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of days remaining before the certificate expires.",
			},
			"subject": schema.StringAttribute{
				Computed:    true,
				Description: "Distinguished name of the certificate subject.",
			},
			"publickey": schema.StringAttribute{
				Computed:    true,
				Description: "Public key algorithm of the certificate.",
			},
			"publickeysize": schema.Int64Attribute{
				Computed:    true,
				Description: "Public key size, in bits.",
			},
			"version": schema.Int64Attribute{
				Computed:    true,
				Description: "Version number of the certificate.",
			},
			"priority": schema.Int64Attribute{
				Computed:    true,
				Description: "Priority of the certificate.",
			},
			"status": schema.StringAttribute{
				Computed:    true,
				Description: "Validity status of the certificate.",
			},
			"passcrypt": schema.StringAttribute{
				Computed:    true,
				Sensitive:   true,
				Description: "Encrypted passphrase of the private-key as stored on the appliance.",
			},
			"data": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of references to the certificate-key pair.",
			},
			"servicename": schema.StringAttribute{
				Computed:    true,
				Description: "Service name to which the certificate is bound.",
			},
			"sandns": schema.StringAttribute{
				Computed:    true,
				Description: "Subject Alternative Name (DNS) entries of the certificate.",
			},
			"sanipadd": schema.StringAttribute{
				Computed:    true,
				Description: "Subject Alternative Name (IP address) entries of the certificate.",
			},
			"ocspresponsestatus": schema.StringAttribute{
				Computed:    true,
				Description: "OCSP response status for the certificate.",
			},
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Flag to determine whether the certificate-key pair is built-in.",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this configuration.",
			},
			"certkeydigest": schema.StringAttribute{
				Computed:    true,
				Description: "Digest (fingerprint) of the certificate.",
			},
			"certificatesource": schema.StringAttribute{
				Computed:    true,
				Description: "Source of the certificate.",
			},
			"certkeystatus": schema.StringAttribute{
				Computed:    true,
				Description: "Status of the certificate-key pair.",
			},
		},
	}
}

// sslcertkeyDataSourceSetAttrFromGet projects a NITRO sslcertkey GET response
// onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them) — no unknown->null resolution or plan preservation is
// required. The shared utils.MapGet* helpers implement that projection.
func sslcertkeyDataSourceSetAttrFromGet(ctx context.Context, data *SslCertKeyDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In sslcertkeyDataSourceSetAttrFromGet Function")

	if v, ok := g["certkey"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Certkey = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Cert = utils.MapGetString(g, "cert")
	data.Key = utils.MapGetString(g, "key")
	data.Fipskey = utils.MapGetString(g, "fipskey")
	data.Hsmkey = utils.MapGetString(g, "hsmkey")
	data.Inform = utils.MapGetString(g, "inform")
	data.Expirymonitor = utils.MapGetString(g, "expirymonitor")
	data.NotificationPeriod = utils.MapGetInt64(g, "notificationperiod")
	data.Bundle = utils.MapGetString(g, "bundle")
	data.LinkCertKeyName = utils.MapGetString(g, "linkcertkeyname")
	data.OcspStaplingCache = utils.MapGetBool(g, "ocspstaplingcache")
	data.DeleteFromDevice = utils.MapGetBool(g, "deletefromdevice")
	data.DeleteCertKeyFilesOnRemoval = utils.MapGetString(g, "deletecertkeyfilesonremoval")

	// passplain is a secret input the GET never returns -> Null.
	data.Password = types.BoolNull()
	data.NoDomainCheck = types.BoolNull()
	data.Passplain = types.StringNull()

	// Read-only cert metadata.
	data.Signaturealg = utils.MapGetString(g, "signaturealg")
	data.Certificatetype = utils.MapGetStringList(g, "certificatetype")
	data.Serial = utils.MapGetString(g, "serial")
	data.Issuer = utils.MapGetString(g, "issuer")
	data.Clientcertnotbefore = utils.MapGetString(g, "clientcertnotbefore")
	data.Clientcertnotafter = utils.MapGetString(g, "clientcertnotafter")
	data.Daystoexpiration = utils.MapGetInt64(g, "daystoexpiration")
	data.Subject = utils.MapGetString(g, "subject")
	data.Publickey = utils.MapGetString(g, "publickey")
	data.Publickeysize = utils.MapGetInt64(g, "publickeysize")
	data.Version = utils.MapGetInt64(g, "version")
	data.Priority = utils.MapGetInt64(g, "priority")
	data.Status = utils.MapGetString(g, "status")
	data.Passcrypt = utils.MapGetString(g, "passcrypt")
	data.Data = utils.MapGetInt64(g, "data")
	data.Servicename = utils.MapGetString(g, "servicename")
	data.Sandns = utils.MapGetString(g, "sandns")
	data.Sanipadd = utils.MapGetString(g, "sanipadd")
	data.Ocspresponsestatus = utils.MapGetString(g, "ocspresponsestatus")
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
	data.Certkeydigest = utils.MapGetString(g, "certkeydigest")
	data.Certificatesource = utils.MapGetString(g, "certificatesource")
	data.Certkeystatus = utils.MapGetString(g, "certkeystatus")
}
