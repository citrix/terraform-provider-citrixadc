package tmsessionaction

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// TmsessionactionDataSourceModel is the data-source-specific model, decoupled from
// TmsessionactionResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the configurable attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits. Every
// non-key attribute is Computed.
type TmsessionactionDataSourceModel struct {
	Id                         types.String `tfsdk:"id"`
	Defaultauthorizationaction types.String `tfsdk:"defaultauthorizationaction"`
	Homepage                   types.String `tfsdk:"homepage"`
	Httponlycookie             types.String `tfsdk:"httponlycookie"`
	Kcdaccount                 types.String `tfsdk:"kcdaccount"`
	Name                       types.String `tfsdk:"name"` // Required lookup key
	Persistentcookie           types.String `tfsdk:"persistentcookie"`
	Persistentcookievalidity   types.Int64  `tfsdk:"persistentcookievalidity"`
	Sesstimeout                types.Int64  `tfsdk:"sesstimeout"`
	Sso                        types.String `tfsdk:"sso"`
	Ssocredential              types.String `tfsdk:"ssocredential"`
	Ssodomain                  types.String `tfsdk:"ssodomain"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/tmsessionaction.json). Never settable; populated from GET.
	Builtin types.List   `tfsdk:"builtin"`
	Feature types.String `tfsdk:"feature"`
}

func TmsessionactionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"defaultauthorizationaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow or deny access to content for which there is no specific authorization policy.",
			},
			"homepage": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Web address of the home page that a user is displayed when authentication vserver is bookmarked and used to login.",
			},
			"httponlycookie": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow only an HTTP session cookie, in which case the cookie cannot be accessed by scripts.",
			},
			"kcdaccount": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Kerberos constrained delegation account name",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the session action. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after a session action is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my action\" or 'my action').",
			},
			"persistentcookie": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable persistent SSO cookies for the traffic management (TM) session. A persistent cookie remains on the user device and is sent with each HTTP request. The cookie becomes stale if the session ends. This setting is overwritten if a traffic action sets persistent cookie to OFF.\nNote: If persistent cookie is enabled, make sure you set the persistent cookie validity.",
			},
			"persistentcookievalidity": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer specifying the number of minutes for which the persistent cookie remains valid. Can be set only if the persistent cookie setting is enabled.",
			},
			"sesstimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Session timeout, in minutes. If there is no traffic during the timeout period, the user is disconnected and must reauthenticate to access intranet resources.",
			},
			"sso": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use single sign-on (SSO) to log users on to all web applications automatically after they authenticate, or pass users to the web application logon page to authenticate to each application individually. Note that this configuration does not honor the following authentication types for security reason. BASIC, DIGEST, and NTLM (without Negotiate NTLM2 Key or Negotiate Sign Flag). Use TM TrafficAction to configure SSO for these authentication types.",
			},
			"ssocredential": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use the primary or secondary authentication credentials for single sign-on (SSO).",
			},
			"ssodomain": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Domain to use for single sign-on (SSO).",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Indicates that a variable is a built-in (SYSTEM INTERNAL) type (MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL). A list of strings.",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this configuration.",
			},
		},
	}
}

// tmsessionactionDataSourceSetAttrFromGet projects a NITRO tmsessionaction GET
// response onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when the
// GET omits them). The shared utils.MapGet* helpers implement that projection.
func tmsessionactionDataSourceSetAttrFromGet(ctx context.Context, data *TmsessionactionDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In tmsessionactionDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Defaultauthorizationaction = utils.MapGetString(g, "defaultauthorizationaction")
	data.Homepage = utils.MapGetString(g, "homepage")
	data.Httponlycookie = utils.MapGetString(g, "httponlycookie")
	data.Kcdaccount = utils.MapGetString(g, "kcdaccount")
	data.Persistentcookie = utils.MapGetString(g, "persistentcookie")
	data.Persistentcookievalidity = utils.MapGetInt64(g, "persistentcookievalidity")
	data.Sesstimeout = utils.MapGetInt64(g, "sesstimeout")
	data.Sso = utils.MapGetString(g, "sso")
	data.Ssocredential = utils.MapGetString(g, "ssocredential")
	data.Ssodomain = utils.MapGetString(g, "ssodomain")

	// Read-only (GET-only) attributes.
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
}
