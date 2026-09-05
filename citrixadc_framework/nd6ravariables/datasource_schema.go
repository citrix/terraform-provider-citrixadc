package nd6ravariables

import (
	"context"
	"fmt"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Nd6ravariablesDataSourceModel is the data-source-specific model, decoupled
// from Nd6ravariablesResourceModel. A data source is a pure read surface, so it
// exposes the read/write attributes (as Computed outputs) plus the read-only
// (GET-only) attributes the resource deliberately omits.
type Nd6ravariablesDataSourceModel struct {
	Id                       types.String `tfsdk:"id"`
	Ceaserouteradv           types.String `tfsdk:"ceaserouteradv"`
	Currhoplimit             types.Int64  `tfsdk:"currhoplimit"`
	Defaultlifetime          types.Int64  `tfsdk:"defaultlifetime"`
	Linkmtu                  types.Int64  `tfsdk:"linkmtu"`
	Managedaddrconfig        types.String `tfsdk:"managedaddrconfig"`
	Maxrtadvinterval         types.Int64  `tfsdk:"maxrtadvinterval"`
	Minrtadvinterval         types.Int64  `tfsdk:"minrtadvinterval"`
	Onlyunicastrtadvresponse types.String `tfsdk:"onlyunicastrtadvresponse"`
	Otheraddrconfig          types.String `tfsdk:"otheraddrconfig"`
	Reachabletime            types.Int64  `tfsdk:"reachabletime"`
	Retranstime              types.Int64  `tfsdk:"retranstime"`
	Sendrouteradv            types.String `tfsdk:"sendrouteradv"`
	Srclinklayeraddroption   types.String `tfsdk:"srclinklayeraddroption"`
	Vlan                     types.Int64  `tfsdk:"vlan"`

	// Read-only (GET-only) attributes from the NITRO read-only set
	// (zion73x_readonly/nd6ravariables.json). Never settable; populated from GET.
	Lastrtadvtime  types.Int64 `tfsdk:"lastrtadvtime"`
	Nextrtadvdelay types.Int64 `tfsdk:"nextrtadvdelay"`
}

func Nd6ravariablesDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"ceaserouteradv": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Cease router advertisements on this vlan.",
			},
			"currhoplimit": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Current Hop limit.",
			},
			"defaultlifetime": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Default life time, in seconds.",
			},
			"linkmtu": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The Link MTU.",
			},
			"managedaddrconfig": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Value to be placed in the Managed address configuration flag field.",
			},
			"maxrtadvinterval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum time allowed between unsolicited multicast RAs, in seconds.",
			},
			"minrtadvinterval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum time interval between RA messages, in seconds.",
			},
			"onlyunicastrtadvresponse": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Send only Unicast Router Advertisements in respond to Router Solicitations.",
			},
			"otheraddrconfig": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Value to be placed in the Other configuration flag field.",
			},
			"reachabletime": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Reachable time, in milliseconds.",
			},
			"retranstime": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Retransmission time, in milliseconds.",
			},
			"sendrouteradv": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "whether the router sends periodic RAs and responds to Router Solicitations.",
			},
			"srclinklayeraddroption": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Include source link layer address option in RA messages.",
			},
			"vlan": schema.Int64Attribute{
				Required:    true,
				Description: "The VLAN number.",
			},

			// Read-only (GET-only) attributes surfaced by the data source.
			"lastrtadvtime": schema.Int64Attribute{
				Computed:    true,
				Description: "Last RA sent timestamp.",
			},
			"nextrtadvdelay": schema.Int64Attribute{
				Computed:    true,
				Description: "Next RA delay.",
			},
		},
	}
}

// nd6ravariablesDataSourceSetAttrFromGet projects a NITRO nd6ravariables GET
// response onto the data-source model using the shared utils.MapGet* helpers.
func nd6ravariablesDataSourceSetAttrFromGet(ctx context.Context, data *Nd6ravariablesDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In nd6ravariablesDataSourceSetAttrFromGet Function")

	data.Ceaserouteradv = utils.MapGetString(g, "ceaserouteradv")
	data.Currhoplimit = utils.MapGetInt64(g, "currhoplimit")
	data.Defaultlifetime = utils.MapGetInt64(g, "defaultlifetime")
	data.Linkmtu = utils.MapGetInt64(g, "linkmtu")
	data.Managedaddrconfig = utils.MapGetString(g, "managedaddrconfig")
	data.Maxrtadvinterval = utils.MapGetInt64(g, "maxrtadvinterval")
	data.Minrtadvinterval = utils.MapGetInt64(g, "minrtadvinterval")
	data.Onlyunicastrtadvresponse = utils.MapGetString(g, "onlyunicastrtadvresponse")
	data.Otheraddrconfig = utils.MapGetString(g, "otheraddrconfig")
	data.Reachabletime = utils.MapGetInt64(g, "reachabletime")
	data.Retranstime = utils.MapGetInt64(g, "retranstime")
	data.Sendrouteradv = utils.MapGetString(g, "sendrouteradv")
	data.Srclinklayeraddroption = utils.MapGetString(g, "srclinklayeraddroption")
	// vlan is the required lookup key supplied via config; only overwrite it when
	// the GET actually echoes it, so the id below is always resolvable.
	if v, ok := g["vlan"]; ok && v != nil {
		data.Vlan = utils.MapGetInt64(g, "vlan")
	}

	// Read-only (GET-only) attributes.
	data.Lastrtadvtime = utils.MapGetInt64(g, "lastrtadvtime")
	data.Nextrtadvdelay = utils.MapGetInt64(g, "nextrtadvdelay")

	// Set ID from the single unique key (vlan).
	data.Id = types.StringValue(fmt.Sprintf("%d", data.Vlan.ValueInt64()))
}
