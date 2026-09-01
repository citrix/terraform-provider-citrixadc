package lbgroup

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// LbgroupDataSourceModel is the data-source-specific model, decoupled from
// LbgroupResourceModel. A data source is a pure read surface (Read only), so it
// exposes the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits (td).
type LbgroupDataSourceModel struct {
	Id                       types.String `tfsdk:"id"`
	Backuppersistencetimeout types.Int64  `tfsdk:"backuppersistencetimeout"`
	Cookiedomain             types.String `tfsdk:"cookiedomain"`
	Cookiename               types.String `tfsdk:"cookiename"`
	Mastervserver            types.String `tfsdk:"mastervserver"`
	Name                     types.String `tfsdk:"name"` // Required lookup key
	Newname                  types.String `tfsdk:"newname"`
	Persistencebackup        types.String `tfsdk:"persistencebackup"`
	Persistencetype          types.String `tfsdk:"persistencetype"`
	Persistmask              types.String `tfsdk:"persistmask"`
	Rule                     types.String `tfsdk:"rule"`
	Timeout                  types.Int64  `tfsdk:"timeout"`
	Usevserverpersistency    types.String `tfsdk:"usevserverpersistency"`
	V6persistmasklen         types.Int64  `tfsdk:"v6persistmasklen"`

	// Read-only (GET-only) metadata from the NITRO read-only set
	// (zion73x_readonly/lbgroup.json). Never settable; populated from GET.
	Td types.Int64 `tfsdk:"td"`
}

func LbgroupDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"backuppersistencetimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time period, in minutes, for which backup persistence is in effect.",
			},
			"cookiedomain": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Domain attribute for the HTTP cookie.",
			},
			"cookiename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use this parameter to specify the cookie name for COOKIE peristence type. It specifies the name of cookie with a maximum of 32 characters. If not specified, cookie name is internally generated.",
			},
			"mastervserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "When USE_VSERVER_PERSISTENCE is enabled, one can use this setting to designate a member vserver as master which is responsible to create the persistence sessions",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the load balancing virtual server group.",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New name for the load balancing virtual server group.",
			},
			"persistencebackup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of backup persistence for the group.",
			},
			"persistencetype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of persistence for the group. Available settings function as follows:\n* SOURCEIP - Create persistence sessions based on the client IP.\n* COOKIEINSERT - Create persistence sessions based on a cookie in client requests. The cookie is inserted by a Set-Cookie directive from the server, in its first response to a client.\n* RULE - Create persistence sessions based on a user defined rule.\n* NONE - Disable persistence for the group.",
			},
			"persistmask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Persistence mask to apply to source IPv4 addresses when creating source IP based persistence sessions.",
			},
			"rule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression, or name of a named expression, against which traffic is evaluated.\n\nThe following requirements apply only to the Citrix ADC CLI:\n* If the expression includes one or more spaces, enclose the entire expression in double quotation marks.\n* If the expression itself includes double quotation marks, escape the quotations by using the \\ character.\n* Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.",
			},
			"timeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time period for which a persistence session is in effect.",
			},
			"usevserverpersistency": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use this parameter to enable vserver level persistence on group members. This allows member vservers to have their own persistence, but need to be compatible with other members persistence rules. When this setting is enabled persistence sessions created by any of the members can be shared by other member vservers.",
			},
			"v6persistmasklen": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Persistence mask to apply to source IPv6 addresses when creating source IP based persistence sessions.",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (intentionally NOT modeled on the resource). All Computed.
			"td": schema.Int64Attribute{
				Computed:    true,
				Description: "Integer value that uniquely identifies the traffic domain in which the entity is configured. Defaults to 0 (default traffic domain).",
			},
		},
	}
}

// lbgroupDataSourceSetAttrFromGet projects a NITRO lbgroup GET response onto the
// data-source model. Attributes are simply filled from the GET (or left Null when
// the GET omits them) via the shared utils.MapGet* helpers.
func lbgroupDataSourceSetAttrFromGet(ctx context.Context, data *LbgroupDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In lbgroupDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	} else {
		data.Id = data.Name
	}

	// Read/write attributes as read-back outputs.
	data.Backuppersistencetimeout = utils.MapGetInt64(g, "backuppersistencetimeout")
	data.Cookiedomain = utils.MapGetString(g, "cookiedomain")
	data.Cookiename = utils.MapGetString(g, "cookiename")
	data.Mastervserver = utils.MapGetString(g, "mastervserver")
	data.Persistencebackup = utils.MapGetString(g, "persistencebackup")
	data.Persistencetype = utils.MapGetString(g, "persistencetype")
	data.Persistmask = utils.MapGetString(g, "persistmask")
	data.Rule = utils.MapGetString(g, "rule")
	data.Timeout = utils.MapGetInt64(g, "timeout")
	data.Usevserverpersistency = utils.MapGetString(g, "usevserverpersistency")
	data.V6persistmasklen = utils.MapGetInt64(g, "v6persistmasklen")

	// newname is rename-only and never returned by GET -> Null.
	data.Newname = types.StringNull()

	// Read-only metadata.
	data.Td = utils.MapGetInt64(g, "td")
}
