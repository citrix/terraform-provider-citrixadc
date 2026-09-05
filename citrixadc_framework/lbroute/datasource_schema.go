package lbroute

import (
	"context"
	"fmt"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// LbrouteDataSourceModel is the data-source-specific model, decoupled from
// LbrouteResourceModel. A data source is a pure read surface, so it can expose
// the full GET projection: the read/write attributes (as Computed outputs) AND
// the read-only attributes that the resource deliberately omits.
type LbrouteDataSourceModel struct {
	Id          types.String `tfsdk:"id"`
	Gatewayname types.String `tfsdk:"gatewayname"`
	Netmask     types.String `tfsdk:"netmask"`
	Network     types.String `tfsdk:"network"`
	Td          types.Int64  `tfsdk:"td"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/lbroute.json). Never settable; populated from GET.
	Flags types.String `tfsdk:"flags"`
}

func LbrouteDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"gatewayname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The name of the route.",
			},
			"netmask": schema.StringAttribute{
				Required:    true,
				Description: "The netmask to which the route belongs.",
			},
			"network": schema.StringAttribute{
				Required:    true,
				Description: "The IP address of the network to which the route belongs.",
			},
			"td": schema.Int64Attribute{
				Required:    true,
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},

			// Read-only (GET-only) attributes surfaced by the data source (these
			// are intentionally NOT modeled on the resource). All Computed.
			"flags": schema.StringAttribute{
				Computed:    true,
				Description: "State of the configured gateway. Possible values = UP, DOWN, UNKNOWN, BUSY, OUT OF SERVICE, GOING OUT OF SERVICE, DOWN WHEN GOING OUT OF SERVICE, NS_EMPTY_STR, Unknown, DISABLED.",
			},
		},
	}
}

// lbrouteDataSourceSetAttrFromGet projects a NITRO lbroute GET response onto the
// data-source model. The shared utils.MapGet* helpers fill each attribute from
// the GET (or leave it Null when the GET omits it).
func lbrouteDataSourceSetAttrFromGet(ctx context.Context, data *LbrouteDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In lbrouteDataSourceSetAttrFromGet Function")

	// Read/write attributes as read-back outputs.
	data.Gatewayname = utils.MapGetString(g, "gatewayname")
	data.Netmask = utils.MapGetString(g, "netmask")
	data.Network = utils.MapGetString(g, "network")
	// td is a config-supplied key; NITRO omits it for the default traffic
	// domain (0), so preserve the configured value instead of nulling it.
	if tdv, tdok := g["td"]; tdok && tdv != nil {
		if iv, err := utils.ConvertToInt64(tdv); err == nil {
			data.Td = types.Int64Value(iv)
		}
	}

	// Read-only attributes.
	data.Flags = utils.MapGetString(g, "flags")

	// Set ID. SDK v2 used "network,netmask,gatewayname".
	data.Id = types.StringValue(fmt.Sprintf("%s,%s,%s", data.Network.ValueString(), data.Netmask.ValueString(), data.Gatewayname.ValueString()))
}
