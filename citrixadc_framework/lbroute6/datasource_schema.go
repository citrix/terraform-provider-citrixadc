package lbroute6

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Lbroute6DataSourceModel is the data-source-specific model, decoupled from
// Lbroute6ResourceModel. A data source is a pure read surface, so it can expose
// the full GET projection: the read/write attributes (as Computed outputs) AND
// the read-only attributes that the resource deliberately omits.
type Lbroute6DataSourceModel struct {
	Id          types.String `tfsdk:"id"`
	Gatewayname types.String `tfsdk:"gatewayname"`
	Network     types.String `tfsdk:"network"`
	Td          types.Int64  `tfsdk:"td"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/lbroute6.json). Never settable; populated from GET.
	Flags types.String `tfsdk:"flags"`
}

func Lbroute6DataSourceSchema() schema.Schema {
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
			"network": schema.StringAttribute{
				Required:    true,
				Description: "The destination network.",
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

// lbroute6DataSourceSetAttrFromGet projects a NITRO lbroute6 GET response onto
// the data-source model. The shared utils.MapGet* helpers fill each attribute
// from the GET (or leave it Null when the GET omits it).
func lbroute6DataSourceSetAttrFromGet(ctx context.Context, data *Lbroute6DataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In lbroute6DataSourceSetAttrFromGet Function")

	// Read/write attributes as read-back outputs.
	data.Gatewayname = utils.MapGetString(g, "gatewayname")
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

	// Set ID. SDK v2 used the plain network value.
	data.Id = types.StringValue(data.Network.ValueString())
}
