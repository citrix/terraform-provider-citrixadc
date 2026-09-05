package systemsession

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SystemsessionDataSourceModel is a dedicated read-only model exposing the GET-only
// session fields. It is intentionally separate from SystemsessionKillResourceModel, which
// is kept minimal (id, sid, all) for the kill action.
type SystemsessionDataSourceModel struct {
	Id                    types.String `tfsdk:"id"`
	Sid                   types.Int64  `tfsdk:"sid"`
	Username              types.String `tfsdk:"username"`
	Logintime             types.String `tfsdk:"logintime"`
	Logintimelocal        types.String `tfsdk:"logintimelocal"`
	Lastactivitytime      types.String `tfsdk:"lastactivitytime"`
	Lastactivitytimelocal types.String `tfsdk:"lastactivitytimelocal"`
	Expirytime            types.String `tfsdk:"expirytime"`
	Numofconnections      types.String `tfsdk:"numofconnections"`
	Currentconn           types.String `tfsdk:"currentconn"`
	Clienttype            types.String `tfsdk:"clienttype"`
	Partitionname         types.String `tfsdk:"partitionname"`
	Clientipaddress       types.String `tfsdk:"clientipaddress"`
}

func SystemsessionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"sid": schema.Int64Attribute{
				Required:    true,
				Description: "ID of the system session about which to display information.",
			},
			// Read-only session fields returned by GET.
			"username": schema.StringAttribute{
				Computed:    true,
				Description: "Name of the user who is logged in.",
			},
			"logintime": schema.StringAttribute{
				Computed:    true,
				Description: "Time when the user logged in.",
			},
			"logintimelocal": schema.StringAttribute{
				Computed:    true,
				Description: "Time (local) when the user logged in.",
			},
			"lastactivitytime": schema.StringAttribute{
				Computed:    true,
				Description: "Time of last activity in the session.",
			},
			"lastactivitytimelocal": schema.StringAttribute{
				Computed:    true,
				Description: "Time (local) of last activity in the session.",
			},
			"expirytime": schema.StringAttribute{
				Computed:    true,
				Description: "Time when the session expires.",
			},
			"numofconnections": schema.StringAttribute{
				Computed:    true,
				Description: "Number of connections in the session.",
			},
			"currentconn": schema.StringAttribute{
				Computed:    true,
				Description: "Indicates the current connection.",
			},
			"clienttype": schema.StringAttribute{
				Computed:    true,
				Description: "Type of client used for the session.",
			},
			"partitionname": schema.StringAttribute{
				Computed:    true,
				Description: "Name of the partition for the session.",
			},
			"clientipaddress": schema.StringAttribute{
				Computed:    true,
				Description: "Client IP address for the session.",
			},
		},
	}
}

// systemsessionDataSourceSetAttrFromGet projects a NITRO systemsession GET response onto
// the data-source model. Every attribute is filled from the GET (or left Null when the GET
// omits it) via the shared utils.MapGet* helpers. The lookup key (sid) and id are set by
// the datasource Read.
func systemsessionDataSourceSetAttrFromGet(ctx context.Context, data *SystemsessionDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In systemsessionDataSourceSetAttrFromGet Function")

	data.Username = utils.MapGetString(g, "username")
	data.Logintime = utils.MapGetString(g, "logintime")
	data.Logintimelocal = utils.MapGetString(g, "logintimelocal")
	data.Lastactivitytime = utils.MapGetString(g, "lastactivitytime")
	data.Lastactivitytimelocal = utils.MapGetString(g, "lastactivitytimelocal")
	data.Expirytime = utils.MapGetString(g, "expirytime")
	data.Numofconnections = utils.MapGetString(g, "numofconnections")
	data.Currentconn = utils.MapGetString(g, "currentconn")
	data.Clienttype = utils.MapGetString(g, "clienttype")
	data.Partitionname = utils.MapGetString(g, "partitionname")
	data.Clientipaddress = utils.MapGetString(g, "clientipaddress")
}
