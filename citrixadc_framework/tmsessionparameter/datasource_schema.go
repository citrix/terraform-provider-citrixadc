package tmsessionparameter

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// TmsessionparameterDataSourceModel is the data-source-specific model, decoupled
// from TmsessionparameterResourceModel.
//
// tmsessionparameter is a singleton (global) configuration object, so the data
// source takes no lookup key and always reads the live values. A data source is a
// pure read surface (Read only; no plan/apply lifecycle), so it can expose the FULL
// GET projection: the configurable attributes (as Computed outputs) AND the
// read-only attributes the resource deliberately omits. Every attribute is Computed.
type TmsessionparameterDataSourceModel struct {
	Id                         types.String `tfsdk:"id"`
	Defaultauthorizationaction types.String `tfsdk:"defaultauthorizationaction"`
	Homepage                   types.String `tfsdk:"homepage"`
	Httponlycookie             types.String `tfsdk:"httponlycookie"`
	Kcdaccount                 types.String `tfsdk:"kcdaccount"`
	Persistentcookie           types.String `tfsdk:"persistentcookie"`
	Persistentcookievalidity   types.Int64  `tfsdk:"persistentcookievalidity"`
	Sesstimeout                types.Int64  `tfsdk:"sesstimeout"`
	Sso                        types.String `tfsdk:"sso"`
	Ssocredential              types.String `tfsdk:"ssocredential"`
	Ssodomain                  types.String `tfsdk:"ssodomain"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/tmsessionparameter.json). Never settable; populated from GET.
	Name                    types.String `tfsdk:"name"`
	Tmsessionpolicybindtype types.String `tfsdk:"tmsessionpolicybindtype"`
	Tmsessionpolicycount    types.Int64  `tfsdk:"tmsessionpolicycount"`
}

func TmsessionparameterDataSourceSchema() schema.Schema {
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
			"persistentcookie": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use persistent SSO cookies for the traffic session. A persistent cookie remains on the user device and is sent with each HTTP request. The cookie becomes stale if the session ends.",
			},
			"persistentcookievalidity": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer specifying the number of minutes for which the persistent cookie remains valid. Can be set only if the persistence cookie setting is enabled.",
			},
			"sesstimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Session timeout, in minutes. If there is no traffic during the timeout period, the user is disconnected and must reauthenticate to access the intranet resources.",
			},
			"sso": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log users on to all web applications automatically after they authenticate, or pass users to the web application logon page to authenticate for each application. Note that this configuration does not honor the following authentication types for security reason. BASIC, DIGEST, and NTLM (without Negotiate NTLM2 Key or Negotiate Sign Flag). Use TM TrafficAction to configure SSO for these authentication types.",
			},
			"ssocredential": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use primary or secondary authentication credentials for single sign-on.",
			},
			"ssodomain": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Domain to use for single sign-on.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"name": schema.StringAttribute{
				Computed:    true,
				Description: "Name returned by the appliance for the TM session parameter object.",
			},
			"tmsessionpolicybindtype": schema.StringAttribute{
				Computed:    true,
				Description: "Indicates current bind type (Classic Policy / Advanced Policy) for TM session policy across all bind entities.",
			},
			"tmsessionpolicycount": schema.Int64Attribute{
				Computed:    true,
				Description: "Count of TM session policies across all bind entities.",
			},
		},
	}
}

// tmsessionparameterDataSourceSetAttrFromGet projects a NITRO tmsessionparameter
// GET response onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when the
// GET omits them). The shared utils.MapGet* helpers implement that projection.
func tmsessionparameterDataSourceSetAttrFromGet(ctx context.Context, data *TmsessionparameterDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In tmsessionparameterDataSourceSetAttrFromGet Function")

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
	data.Name = utils.MapGetString(g, "name")
	data.Tmsessionpolicybindtype = utils.MapGetString(g, "tmsessionpolicybindtype")
	data.Tmsessionpolicycount = utils.MapGetInt64(g, "tmsessionpolicycount")

	// Singleton resource - static ID.
	data.Id = types.StringValue("tmsessionparameter-config")
}
