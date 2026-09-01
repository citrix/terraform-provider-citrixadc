package authenticationnegotiateaction

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AuthenticationnegotiateactionDataSourceModel is the data-source-specific model,
// decoupled from AuthenticationnegotiateactionResourceModel. A data source is a
// pure read surface (Read only; no plan/apply lifecycle), so it can expose the
// FULL GET projection: the read/write attributes (as Computed outputs) AND the
// read-only attributes the resource deliberately omits. Every non-key attribute
// is Computed; the Framework's per-attribute model <-> schema reflection requires
// this model to have exactly the attributes the data-source schema declares.
type AuthenticationnegotiateactionDataSourceModel struct {
	Id                         types.String `tfsdk:"id"`
	Defaultauthenticationgroup types.String `tfsdk:"defaultauthenticationgroup"`
	Domain                     types.String `tfsdk:"domain"`
	Domainuser                 types.String `tfsdk:"domainuser"`
	Domainuserpasswd           types.String `tfsdk:"domainuserpasswd"`
	DomainuserpasswdWo         types.String `tfsdk:"domainuserpasswd_wo"`
	DomainuserpasswdWoVersion  types.Int64  `tfsdk:"domainuserpasswd_wo_version"`
	Keytab                     types.String `tfsdk:"keytab"`
	Name                       types.String `tfsdk:"name"`
	Ntlmpath                   types.String `tfsdk:"ntlmpath"`
	Ou                         types.String `tfsdk:"ou"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/authenticationnegotiateaction.json). Never settable;
	// populated from GET.
	Kcdspn types.String `tfsdk:"kcdspn"`
}

func AuthenticationnegotiateactionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"defaultauthenticationgroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is the default group that is chosen when the authentication succeeds in addition to extracted groups.",
			},
			"domain": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Domain name of the service principal that represnts Citrix ADC.",
			},
			"domainuser": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "User name of the account that is mapped with Citrix ADC principal. This can be given along with domain and password when keytab file is not available. If username is given along with keytab file, then that keytab file will be searched for this user's credentials.",
			},
			"domainuserpasswd": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Password of the account that is mapped to the Citrix ADC principal.",
			},
			"domainuserpasswd_wo": schema.StringAttribute{
				Optional:    true,
				Description: "Password of the account that is mapped to the Citrix ADC principal.",
			},
			"domainuserpasswd_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Increment this version to signal a domainuserpasswd_wo update.",
			},
			"keytab": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The path to the keytab file that is used to decrypt kerberos tickets presented to Citrix ADC. If keytab is not available, domain/username/password can be specified in the negotiate action configuration",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the AD KDC server profile (negotiate action).\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after AD KDC server profile is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my authentication action\" or 'my authentication action').",
			},
			"ntlmpath": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The path to the site that is enabled for NTLM authentication, including FQDN of the server. This is used when clients fallback to NTLM.",
			},
			"ou": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Active Directory organizational units (OU) attribute.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"kcdspn": schema.StringAttribute{
				Computed:    true,
				Description: "Host SPN extracted from keytab file.",
			},
		},
	}
}

// authenticationnegotiateactionDataSourceSetAttrFromGet projects a NITRO
// authenticationnegotiateaction GET response onto the data-source model. Because
// a data source has no plan/apply reconciliation, attributes are simply filled
// from the GET (or left Null when the GET omits them). The shared utils.MapGet*
// helpers implement that projection.
func authenticationnegotiateactionDataSourceSetAttrFromGet(ctx context.Context, data *AuthenticationnegotiateactionDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In authenticationnegotiateactionDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Defaultauthenticationgroup = utils.MapGetString(g, "defaultauthenticationgroup")
	data.Domain = utils.MapGetString(g, "domain")
	data.Domainuser = utils.MapGetString(g, "domainuser")
	data.Keytab = utils.MapGetString(g, "keytab")
	data.Ntlmpath = utils.MapGetString(g, "ntlmpath")
	data.Ou = utils.MapGetString(g, "ou")

	// domainuserpasswd / domainuserpasswd_wo(+version) are write-only or
	// action-only inputs the GET never returns -> Null.
	data.Domainuserpasswd = types.StringNull()
	data.DomainuserpasswdWo = types.StringNull()
	data.DomainuserpasswdWoVersion = types.Int64Null()

	// Read-only attributes.
	data.Kcdspn = utils.MapGetString(g, "kcdspn")
}
