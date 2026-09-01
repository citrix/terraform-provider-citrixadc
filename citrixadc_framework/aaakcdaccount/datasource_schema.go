package aaakcdaccount

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AaakcdaccountDataSourceModel is the data-source-specific model, decoupled from
// AaakcdaccountResourceModel. A data source is a pure read surface, so it can
// expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits.
type AaakcdaccountDataSourceModel struct {
	Id                   types.String `tfsdk:"id"`
	Cacert               types.String `tfsdk:"cacert"`
	Delegateduser        types.String `tfsdk:"delegateduser"`
	Enterpriserealm      types.String `tfsdk:"enterpriserealm"`
	Kcdaccount           types.String `tfsdk:"kcdaccount"` // Required lookup key
	Kcdpassword          types.String `tfsdk:"kcdpassword"`
	KcdpasswordWo        types.String `tfsdk:"kcdpassword_wo"`
	KcdpasswordWoVersion types.Int64  `tfsdk:"kcdpassword_wo_version"`
	Keytab               types.String `tfsdk:"keytab"`
	Realmstr             types.String `tfsdk:"realmstr"`
	Saltexpression       types.String `tfsdk:"saltexpression"`
	Servicespn           types.String `tfsdk:"servicespn"`
	Usercert             types.String `tfsdk:"usercert"`
	Userrealm            types.String `tfsdk:"userrealm"`

	// Read-only (GET-only) attributes from the NITRO read-only set
	// (zion73x_readonly/aaakcdaccount.json). Never settable; populated from GET.
	Principle types.String `tfsdk:"principle"`
	Kcdspn    types.String `tfsdk:"kcdspn"`
}

func AaakcdaccountDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"cacert": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "CA Cert for UserCert or when doing PKINIT backchannel.",
			},
			"delegateduser": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Username that can perform kerberos constrained delegation.",
			},
			"enterpriserealm": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enterprise Realm of the user. This should be given only in certain KDC deployments where KDC expects Enterprise username instead of Principal Name",
			},
			"kcdaccount": schema.StringAttribute{
				Required:    true,
				Description: "The name of the KCD account.",
			},
			"kcdpassword": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
				Description: "Password for Delegated User.",
			},
			"kcdpassword_wo": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "Password for Delegated User.",
			},
			"kcdpassword_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Increment this version to signal a kcdpassword_wo update.",
			},
			"keytab": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The path to the keytab file. If specified other parameters in this command need not be given",
			},
			"realmstr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Kerberos Realm.",
			},
			"saltexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Salt expression used by Kerberos impersonation. When configured, this expression will be used for key derivation with AES-128 or AES-256 encryption types. For RC4 encryption, the salt is not used. If the salt expression is not set, the default behavior is to derive the salt value from the Kerberos principal.",
			},
			"servicespn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Service SPN. When specified, this will be used to fetch kerberos tickets. If not specified, Citrix ADC will construct SPN using service fqdn",
			},
			"usercert": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "SSL Cert (including private key) for Delegated User.",
			},
			"userrealm": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Realm of the user",
			},

			// Read-only (GET-only) attributes surfaced by the data source. All Computed.
			"principle": schema.StringAttribute{
				Computed:    true,
				Description: "SPN extracted from keytab file.",
			},
			"kcdspn": schema.StringAttribute{
				Computed:    true,
				Description: "Host SPN extracted from keytab file.",
			},
		},
	}
}

// aaakcdaccountDataSourceSetAttrFromGet projects a NITRO aaakcdaccount GET
// response onto the data-source model. Attributes are simply filled from the
// GET (or left Null when the GET omits them) via the shared utils.MapGet*
// helpers.
func aaakcdaccountDataSourceSetAttrFromGet(ctx context.Context, data *AaakcdaccountDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In aaakcdaccountDataSourceSetAttrFromGet Function")

	if v, ok := g["kcdaccount"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Kcdaccount = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Cacert = utils.MapGetString(g, "cacert")
	data.Delegateduser = utils.MapGetString(g, "delegateduser")
	data.Enterpriserealm = utils.MapGetString(g, "enterpriserealm")
	data.Keytab = utils.MapGetString(g, "keytab")
	data.Realmstr = utils.MapGetString(g, "realmstr")
	data.Saltexpression = utils.MapGetString(g, "saltexpression")
	data.Servicespn = utils.MapGetString(g, "servicespn")
	data.Usercert = utils.MapGetString(g, "usercert")
	data.Userrealm = utils.MapGetString(g, "userrealm")

	// kcdpassword / kcdpassword_wo(+version) are write-only secrets the GET never
	// returns -> Null.
	data.Kcdpassword = types.StringNull()
	data.KcdpasswordWo = types.StringNull()
	data.KcdpasswordWoVersion = types.Int64Null()

	// Read-only attributes.
	data.Principle = utils.MapGetString(g, "principle")
	data.Kcdspn = utils.MapGetString(g, "kcdspn")
}
