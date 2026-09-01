package aaatacacsparams

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AaatacacsparamsDataSourceModel is the data-source-specific model, decoupled
// from AaatacacsparamsResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only metadata attributes that the resource deliberately
// omits (builtin, feature). Every non-key attribute is Computed; the Framework's
// per-attribute model <-> schema reflection requires this model to have exactly
// the attributes the data-source schema declares, which is why it cannot reuse
// the resource model.
type AaatacacsparamsDataSourceModel struct {
	Id                         types.String `tfsdk:"id"`
	Accounting                 types.String `tfsdk:"accounting"`
	Auditfailedcmds            types.String `tfsdk:"auditfailedcmds"`
	Authorization              types.String `tfsdk:"authorization"`
	Authtimeout                types.Int64  `tfsdk:"authtimeout"`
	Defaultauthenticationgroup types.String `tfsdk:"defaultauthenticationgroup"`
	Groupattrname              types.String `tfsdk:"groupattrname"`
	Serverip                   types.String `tfsdk:"serverip"`
	Serverport                 types.Int64  `tfsdk:"serverport"`
	Tacacssecret               types.String `tfsdk:"tacacssecret"`
	TacacssecretWo             types.String `tfsdk:"tacacssecret_wo"`
	TacacssecretWoVersion      types.Int64  `tfsdk:"tacacssecret_wo_version"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/aaatacacsparams.json). Never settable; populated from GET.
	Builtin types.List   `tfsdk:"builtin"`
	Feature types.String `tfsdk:"feature"`
}

func AaatacacsparamsDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"accounting": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Send accounting messages to the TACACS+ server.",
			},
			"auditfailedcmds": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The option for sending accounting messages to the TACACS+ server.",
			},
			"authorization": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use streaming authorization on the TACACS+ server.",
			},
			"authtimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of seconds that the Citrix ADC waits for a response from the TACACS+ server.",
			},
			"defaultauthenticationgroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is the default group that is chosen when the authentication succeeds in addition to extracted groups.",
			},
			"groupattrname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "TACACS+ group attribute name.Used for group extraction on the TACACS+ server.",
			},
			"serverip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of your TACACS+ server.",
			},
			"serverport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number on which the TACACS+ server listens for connections.",
			},
			"tacacssecret": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
				Description: "Key shared between the TACACS+ server and clients. Required for allowing the Citrix ADC to communicate with the TACACS+ server.",
			},
			"tacacssecret_wo": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "Key shared between the TACACS+ server and clients. Required for allowing the Citrix ADC to communicate with the TACACS+ server.",
			},
			"tacacssecret_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Increment this version to signal a tacacssecret_wo update.",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
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

// aaatacacsparamsDataSourceSetAttrFromGet projects a NITRO aaatacacsparams GET
// response onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func aaatacacsparamsDataSourceSetAttrFromGet(ctx context.Context, data *AaatacacsparamsDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In aaatacacsparamsDataSourceSetAttrFromGet Function")

	// aaatacacsparams is a singleton; use the same static ID as the resource.
	data.Id = types.StringValue("aaatacacsparams-config")

	// Read/write attributes as read-back outputs.
	data.Accounting = utils.MapGetString(g, "accounting")
	data.Auditfailedcmds = utils.MapGetString(g, "auditfailedcmds")
	data.Authorization = utils.MapGetString(g, "authorization")
	data.Authtimeout = utils.MapGetInt64(g, "authtimeout")
	data.Defaultauthenticationgroup = utils.MapGetString(g, "defaultauthenticationgroup")
	data.Groupattrname = utils.MapGetString(g, "groupattrname")
	data.Serverip = utils.MapGetString(g, "serverip")
	data.Serverport = utils.MapGetInt64(g, "serverport")

	// tacacssecret / tacacssecret_wo(+version) are write-only or secret inputs
	// the GET never returns -> Null.
	data.Tacacssecret = types.StringNull()
	data.TacacssecretWo = types.StringNull()
	data.TacacssecretWoVersion = types.Int64Null()

	// Read-only metadata.
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
}
