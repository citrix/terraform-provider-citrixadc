package sslcrl

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SslcrlDataSourceModel is the data-source-specific model, decoupled from
// SslcrlResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only CRL metadata attributes that the resource
// deliberately omits (issuer, version, lastupdate, nextupdate, ...). Every
// non-key attribute is Computed; the Framework's per-attribute model <-> schema
// reflection requires this model to have exactly the attributes the data-source
// schema declares, which is why it cannot reuse the resource model.
type SslcrlDataSourceModel struct {
	Id                types.String `tfsdk:"id"`
	Basedn            types.String `tfsdk:"basedn"`
	Binary            types.String `tfsdk:"binary"`
	Binddn            types.String `tfsdk:"binddn"`
	Cacert            types.String `tfsdk:"cacert"`
	Cacertfile        types.String `tfsdk:"cacertfile"`
	Cakeyfile         types.String `tfsdk:"cakeyfile"`
	Crlname           types.String `tfsdk:"crlname"` // Required lookup key
	Crlpath           types.String `tfsdk:"crlpath"`
	Day               types.Int64  `tfsdk:"day"`
	Gencrl            types.String `tfsdk:"gencrl"`
	Indexfile         types.String `tfsdk:"indexfile"`
	Inform            types.String `tfsdk:"inform"`
	Interval          types.String `tfsdk:"interval"`
	Method            types.String `tfsdk:"method"`
	Password          types.String `tfsdk:"password"`
	PasswordWo        types.String `tfsdk:"password_wo"`
	PasswordWoVersion types.Int64  `tfsdk:"password_wo_version"`
	Port              types.Int64  `tfsdk:"port"`
	Refresh           types.String `tfsdk:"refresh"`
	Revoke            types.String `tfsdk:"revoke"`
	Scope             types.String `tfsdk:"scope"`
	Server            types.String `tfsdk:"server"`
	Time              types.String `tfsdk:"time"`
	Url               types.String `tfsdk:"url"`

	// Read-only (GET-only) CRL metadata from the NITRO doc read-only set
	// (zion73x_readonly/sslcrl.json). Never settable; populated from GET.
	Flags            types.Int64  `tfsdk:"flags"`
	Lastupdatetime   types.Int64  `tfsdk:"lastupdatetime"`
	Version          types.Int64  `tfsdk:"version"`
	Signaturealgo    types.String `tfsdk:"signaturealgo"`
	Issuer           types.String `tfsdk:"issuer"`
	Lastupdate       types.String `tfsdk:"lastupdate"`
	Nextupdate       types.String `tfsdk:"nextupdate"`
	Daystoexpiration types.Int64  `tfsdk:"daystoexpiration"`
}

func SslcrlDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"basedn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Base distinguished name (DN), which is used in an LDAP search to search for a CRL. Citrix recommends searching for the Base DN instead of the Issuer Name from the CA certificate, because the Issuer Name field might not exactly match the LDAP directory structure's DN.",
			},
			"binary": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Set the LDAP-based CRL retrieval mode to binary.",
			},
			"binddn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Bind distinguished name (DN) to be used to access the CRL object in the LDAP repository if access to the LDAP repository is restricted or anonymous access is not allowed.",
			},
			"cacert": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "CA certificate that has issued the CRL. Required if CRL Auto Refresh is selected. Install the CA certificate on the appliance before adding the CRL.",
			},
			"cacertfile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of and, optionally, path to the CA certificate file.\n/nsconfig/ssl/ is the default path.",
			},
			"cakeyfile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of and, optionally, path to the CA key file. /nsconfig/ssl/ is the default path",
			},
			"crlname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the Certificate Revocation List (CRL). Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the CRL is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my crl\" or 'my crl').",
			},
			"crlpath": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Path to the CRL file. /var/netscaler/ssl/ is the default path.",
			},
			"day": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Day on which to refresh the CRL, or, if the Interval parameter is not set, the number of days after which to refresh the CRL. If Interval is set to MONTHLY, specify the date. If Interval is set to WEEKLY, specify the day of the week (for example, Sun=0 and Sat=6). This parameter is not applicable if the Interval is set to DAILY.",
			},
			"gencrl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of and, optionally, path to the CRL file to be generated. The list of certificates that have been revoked is obtained from the index file. /nsconfig/ssl/ is the default path.",
			},
			"indexfile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of and, optionally, path to the file containing the serial numbers of all the certificates that are revoked. Revoked certificates are appended to the file. /nsconfig/ssl/ is the default path",
			},
			"inform": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Input format of the CRL file. The two formats supported on the appliance are:\nPEM - Privacy Enhanced Mail.\nDER - Distinguished Encoding Rule.",
			},
			"interval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "CRL refresh interval. Use the NONE setting to unset this parameter.",
			},
			"method": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Method for CRL refresh. If LDAP is selected, specify the method, CA certificate, base DN, port, and LDAP server name. If HTTP is selected, specify the CA certificate, method, URL, and port. Cannot be changed after a CRL is added.",
			},
			"password": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
				Description: "Password to access the CRL in the LDAP repository if access to the LDAP repository is restricted or anonymous access is not allowed.",
			},
			"password_wo": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "Password to access the CRL in the LDAP repository if access to the LDAP repository is restricted or anonymous access is not allowed.",
			},
			"password_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Increment this version to signal a password_wo update.",
			},
			"port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port for the LDAP server.",
			},
			"refresh": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Set CRL auto refresh.",
			},
			"revoke": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of and, optionally, path to the certificate to be revoked. /nsconfig/ssl/ is the default path.",
			},
			"scope": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Extent of the search operation on the LDAP server. Available settings function as follows:\nOne - One level below Base DN.\nBase - Exactly the same level as Base DN.",
			},
			"server": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of the LDAP server from which to fetch the CRLs.",
			},
			"time": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in hours (1-24) and minutes (1-60), at which to refresh the CRL.",
			},
			"url": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL of the CRL distribution point.",
			},

			// Read-only (GET-only) CRL metadata surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"flags": schema.Int64Attribute{
				Computed:    true,
				Description: "CRL status flag.",
			},
			"lastupdatetime": schema.Int64Attribute{
				Computed:    true,
				Description: "Last CRL refresh time.",
			},
			"version": schema.Int64Attribute{
				Computed:    true,
				Description: "CRL version.",
			},
			"signaturealgo": schema.StringAttribute{
				Computed:    true,
				Description: "Signature algorithm.",
			},
			"issuer": schema.StringAttribute{
				Computed:    true,
				Description: "Issuer name.",
			},
			"lastupdate": schema.StringAttribute{
				Computed:    true,
				Description: "Last update time.",
			},
			"nextupdate": schema.StringAttribute{
				Computed:    true,
				Description: "Next update time.",
			},
			"daystoexpiration": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of days remaining for the CRL to expire.",
			},
		},
	}
}

// sslcrlDataSourceSetAttrFromGet projects a NITRO sslcrl GET response onto the
// data-source model. Because a data source has no plan/apply reconciliation,
// attributes are simply filled from the GET (or left Null when the GET omits
// them). The shared utils.MapGet* helpers implement that projection.
func sslcrlDataSourceSetAttrFromGet(ctx context.Context, data *SslcrlDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In sslcrlDataSourceSetAttrFromGet Function")

	if v, ok := g["crlname"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Crlname = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Basedn = utils.MapGetString(g, "basedn")
	data.Binary = utils.MapGetString(g, "binary")
	data.Binddn = utils.MapGetString(g, "binddn")
	data.Cacert = utils.MapGetString(g, "cacert")
	data.Cacertfile = utils.MapGetString(g, "cacertfile")
	data.Cakeyfile = utils.MapGetString(g, "cakeyfile")
	data.Crlpath = utils.MapGetString(g, "crlpath")
	data.Day = utils.MapGetInt64(g, "day")
	data.Gencrl = utils.MapGetString(g, "gencrl")
	data.Indexfile = utils.MapGetString(g, "indexfile")
	data.Inform = utils.MapGetString(g, "inform")
	data.Interval = utils.MapGetString(g, "interval")
	data.Method = utils.MapGetString(g, "method")
	data.Port = utils.MapGetInt64(g, "port")
	data.Refresh = utils.MapGetString(g, "refresh")
	data.Revoke = utils.MapGetString(g, "revoke")
	data.Scope = utils.MapGetString(g, "scope")
	data.Server = utils.MapGetString(g, "server")
	data.Time = utils.MapGetString(g, "time")
	data.Url = utils.MapGetString(g, "url")

	// password / password_wo(+version) are write-only secret inputs the GET
	// never returns -> Null.
	data.Password = types.StringNull()
	data.PasswordWo = types.StringNull()
	data.PasswordWoVersion = types.Int64Null()

	// Read-only CRL metadata.
	data.Flags = utils.MapGetInt64(g, "flags")
	data.Lastupdatetime = utils.MapGetInt64(g, "lastupdatetime")
	data.Version = utils.MapGetInt64(g, "version")
	data.Signaturealgo = utils.MapGetString(g, "signaturealgo")
	data.Issuer = utils.MapGetString(g, "issuer")
	data.Lastupdate = utils.MapGetString(g, "lastupdate")
	data.Nextupdate = utils.MapGetString(g, "nextupdate")
	data.Daystoexpiration = utils.MapGetInt64(g, "daystoexpiration")
}
