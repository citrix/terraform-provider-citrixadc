package botsettings

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// BotsettingsDataSourceModel is the data-source-specific model, decoupled from
// BotsettingsResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only metadata attributes that the resource deliberately
// omits (builtin, feature). Every non-key attribute is Computed; the Framework's
// per-attribute model <-> schema reflection requires this model to have exactly
// the attributes the data-source schema declares, which is why it cannot reuse
// the resource model.
type BotsettingsDataSourceModel struct {
	Id                         types.String `tfsdk:"id"`
	Defaultnonintrusiveprofile types.String `tfsdk:"defaultnonintrusiveprofile"`
	Defaultprofile             types.String `tfsdk:"defaultprofile"`
	Dfprequestlimit            types.Int64  `tfsdk:"dfprequestlimit"`
	Javascriptname             types.String `tfsdk:"javascriptname"`
	Proxypassword              types.String `tfsdk:"proxypassword"`
	Proxyport                  types.Int64  `tfsdk:"proxyport"`
	Proxyserver                types.String `tfsdk:"proxyserver"`
	Proxyusername              types.String `tfsdk:"proxyusername"`
	Sessioncookiename          types.String `tfsdk:"sessioncookiename"`
	Sessiontimeout             types.Int64  `tfsdk:"sessiontimeout"`
	Signatureautoupdate        types.String `tfsdk:"signatureautoupdate"`
	Signatureurl               types.String `tfsdk:"signatureurl"`
	Trapurlautogenerate        types.String `tfsdk:"trapurlautogenerate"`
	Trapurlinterval            types.Int64  `tfsdk:"trapurlinterval"`
	Trapurllength              types.Int64  `tfsdk:"trapurllength"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/botsettings.json). Never settable; populated from GET.
	Builtin types.List   `tfsdk:"builtin"`
	Feature types.String `tfsdk:"feature"`
}

func BotsettingsDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"defaultnonintrusiveprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Profile to use when the feature is not enabled but feature is licensed. NonIntrusive checks will be disabled and IPRep cronjob(24 Hours) will be removed if this is set to BOT_BYPASS.",
			},
			"defaultprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Profile to use when a connection does not match any policy. Default setting is \" \", which sends unmatched connections back to the Citrix ADC without attempting to filter them further.",
			},
			"dfprequestlimit": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of requests to allow without bot session cookie if device fingerprint is enabled",
			},
			"javascriptname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the JavaScript that the Bot Management feature  uses in response.\nMust begin with a letter or number, and can consist of from 1 to 31 letters, numbers, and the hyphen (-) and underscore (_) symbols.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my cookie name\" or 'my cookie name').",
			},
			"proxypassword": schema.StringAttribute{
				Sensitive:   true,
				Optional:    true,
				Computed:    true,
				Description: "Password with which user logs on.",
			},
			"proxyport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Proxy Server Port to get updated signatures from AWS.",
			},
			"proxyserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Proxy Server IP to get updated signatures from AWS.",
			},
			"proxyusername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Proxy Username",
			},
			"sessioncookiename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the SessionCookie that the Bot Management feature uses for tracking.\nMust begin with a letter or number, and can consist of from 1 to 31 letters, numbers, and the hyphen (-) and underscore (_) symbols.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my cookie name\" or 'my cookie name').",
			},
			"sessiontimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Timeout, in seconds, after which a user session is terminated.",
			},
			"signatureautoupdate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Flag used to enable/disable bot auto update signatures",
			},
			"signatureurl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL to download the bot signature mapping file from server",
			},
			"trapurlautogenerate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable/disable trap URL auto generation. When enabled, trap URL is updated within the configured interval.",
			},
			"trapurlinterval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time in seconds after which trap URL is updated.",
			},
			"trapurllength": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Length of the auto-generated trap URL.",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Flag to determine if bot engine setting is built-in or not. Possible values = MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL.",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
		},
	}
}

// botsettingsDataSourceSetAttrFromGet projects a NITRO botsettings GET response
// onto the data-source model. botsettings is a singleton, so the ID is a static
// identifier rather than a lookup key. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func botsettingsDataSourceSetAttrFromGet(ctx context.Context, data *BotsettingsDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In botsettingsDataSourceSetAttrFromGet Function")

	// botsettings is a singleton config - use a static ID.
	data.Id = types.StringValue("botsettings-config")

	// Read/write attributes as read-back outputs.
	data.Defaultnonintrusiveprofile = utils.MapGetString(g, "defaultnonintrusiveprofile")
	data.Defaultprofile = utils.MapGetString(g, "defaultprofile")
	data.Dfprequestlimit = utils.MapGetInt64(g, "dfprequestlimit")
	data.Javascriptname = utils.MapGetString(g, "javascriptname")
	data.Proxyport = utils.MapGetInt64(g, "proxyport")
	data.Proxyserver = utils.MapGetString(g, "proxyserver")
	data.Proxyusername = utils.MapGetString(g, "proxyusername")
	data.Sessioncookiename = utils.MapGetString(g, "sessioncookiename")
	data.Sessiontimeout = utils.MapGetInt64(g, "sessiontimeout")
	data.Signatureautoupdate = utils.MapGetString(g, "signatureautoupdate")
	data.Signatureurl = utils.MapGetString(g, "signatureurl")
	data.Trapurlautogenerate = utils.MapGetString(g, "trapurlautogenerate")
	data.Trapurlinterval = utils.MapGetInt64(g, "trapurlinterval")
	data.Trapurllength = utils.MapGetInt64(g, "trapurllength")

	// proxypassword is a secret input the GET never returns -> Null.
	data.Proxypassword = types.StringNull()

	// Read-only metadata.
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
}
