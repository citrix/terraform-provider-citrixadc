package nd6

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Nd6DataSourceModel is the data-source-specific model, decoupled from
// Nd6ResourceModel. A data source is a pure read surface, so it exposes the
// read/write attributes (as Computed outputs) plus the read-only (GET-only)
// attributes the resource deliberately omits.
type Nd6DataSourceModel struct {
	Id       types.String `tfsdk:"id"`
	Ifnum    types.String `tfsdk:"ifnum"`
	Mac      types.String `tfsdk:"mac"`
	Neighbor types.String `tfsdk:"neighbor"`
	Nodeid   types.Int64  `tfsdk:"nodeid"`
	Td       types.Int64  `tfsdk:"td"`
	Vlan     types.Int64  `tfsdk:"vlan"`
	Vtep     types.String `tfsdk:"vtep"`
	Vxlan    types.Int64  `tfsdk:"vxlan"`

	// Read-only (GET-only) attributes from the NITRO read-only set
	// (zion73x_readonly/nd6.json). Never settable; populated from GET.
	State        types.String `tfsdk:"state"`
	Timeout      types.Int64  `tfsdk:"timeout"`
	Flags        types.Int64  `tfsdk:"flags"`
	Controlplane types.Bool   `tfsdk:"controlplane"`
	Channel      types.Int64  `tfsdk:"channel"`
}

func Nd6DataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"ifnum": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Interface through which the adjacent network device is available, specified in slot/port notation (for example, 1/3). Use spaces to separate multiple entries.",
			},
			"mac": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "MAC address of the adjacent network device.",
			},
			"neighbor": schema.StringAttribute{
				Required:    true,
				Description: "Link-local IPv6 address of the adjacent network device to add to the ND6 table.",
			},
			"nodeid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique number that identifies the cluster node.",
			},
			"td": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
			"vlan": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer value that uniquely identifies the VLAN on which the adjacent network device exists.",
			},
			"vtep": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of the VXLAN tunnel endpoint (VTEP) through which the IPv6 address of this ND6 entry is reachable.",
			},
			"vxlan": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "ID of the VXLAN on which the IPv6 address of this ND6 entry is reachable.",
			},

			// Read-only (GET-only) attributes surfaced by the data source.
			"state": schema.StringAttribute{
				Computed:    true,
				Description: "ND6 state (INCOMPLETE, REACHABLE, STALE, DELAY, PROBE).",
			},
			"timeout": schema.Int64Attribute{
				Computed:    true,
				Description: "Time elapsed.",
			},
			"flags": schema.Int64Attribute{
				Computed:    true,
				Description: "Flag for static/permanent entry.",
			},
			"controlplane": schema.BoolAttribute{
				Computed:    true,
				Description: "This nd6 entry is populated by a control plane protocol.",
			},
			"channel": schema.Int64Attribute{
				Computed:    true,
				Description: "The tunnel that is bound to a netbridge.",
			},
		},
	}
}

// nd6DataSourceSetAttrFromGet projects a NITRO nd6 GET response onto the
// data-source model using the shared utils.MapGet* helpers.
func nd6DataSourceSetAttrFromGet(ctx context.Context, data *Nd6DataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In nd6DataSourceSetAttrFromGet Function")

	if v, ok := g["neighbor"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
	}

	data.Ifnum = utils.MapGetString(g, "ifnum")
	data.Mac = utils.MapGetString(g, "mac")
	data.Neighbor = utils.MapGetString(g, "neighbor")
	data.Nodeid = utils.MapGetInt64(g, "nodeid")
	// td is a config-supplied key; NITRO omits it for the default traffic
	// domain (0), so preserve the configured value instead of nulling it.
	if tdv, tdok := g["td"]; tdok && tdv != nil {
		if iv, err := utils.ConvertToInt64(tdv); err == nil {
			data.Td = types.Int64Value(iv)
		}
	}
	data.Vlan = utils.MapGetInt64(g, "vlan")
	data.Vtep = utils.MapGetString(g, "vtep")
	data.Vxlan = utils.MapGetInt64(g, "vxlan")

	// Read-only (GET-only) attributes.
	data.State = utils.MapGetString(g, "state")
	data.Timeout = utils.MapGetInt64(g, "timeout")
	data.Flags = utils.MapGetInt64(g, "flags")
	data.Controlplane = utils.MapGetBool(g, "controlplane")
	data.Channel = utils.MapGetInt64(g, "channel")
}
