package lldpneighbors

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// LldpneighborsDataSourceModel is the datasource-only model. It intentionally
// does NOT reuse the resource model (LldpneighborsResourceModel): the resource
// is an action-only clear with a minimal schema (id/ifnum/nodeid), while the
// datasource exposes the full read-only LLDP neighbor telemetry returned by
// get(all). All telemetry fields are string on the vendored NITRO struct
// (vendor/.../resource/config/lldp/lldpneighbors.go), so they map to
// types.String here.
type LldpneighborsDataSourceModel struct {
	Id     types.String `tfsdk:"id"`
	Ifnum  types.String `tfsdk:"ifnum"`
	Nodeid types.Int64  `tfsdk:"nodeid"`

	// Read-only LLDP neighbor telemetry (Computed).
	Chassisidsubtype   types.String `tfsdk:"chassisidsubtype"`
	Chassisid          types.String `tfsdk:"chassisid"`
	Portidsubtype      types.String `tfsdk:"portidsubtype"`
	Portid             types.String `tfsdk:"portid"`
	Ttl                types.String `tfsdk:"ttl"`
	Portdescription    types.String `tfsdk:"portdescription"`
	Sys                types.String `tfsdk:"sys"`
	Sysdesc            types.String `tfsdk:"sysdesc"`
	Mgmtaddresssubtype types.String `tfsdk:"mgmtaddresssubtype"`
	Mgmtaddress        types.String `tfsdk:"mgmtaddress"`
	Iftype             types.String `tfsdk:"iftype"`
	Ifnumber           types.String `tfsdk:"ifnumber"`
	Vlan               types.String `tfsdk:"vlan"`
	Vlanid             types.String `tfsdk:"vlanid"`
	Portprotosupported types.String `tfsdk:"portprotosupported"`
	Portprotoenabled   types.String `tfsdk:"portprotoenabled"`
	Portprotoid        types.String `tfsdk:"portprotoid"`
	Portvlanid         types.String `tfsdk:"portvlanid"`
	Protocolid         types.String `tfsdk:"protocolid"`
	Linkaggrcapable    types.String `tfsdk:"linkaggrcapable"`
	Linkaggrenabled    types.String `tfsdk:"linkaggrenabled"`
	Linkaggrid         types.String `tfsdk:"linkaggrid"`
	Flag               types.String `tfsdk:"flag"`
	Syscapabilities    types.String `tfsdk:"syscapabilities"`
	Syscapenabled      types.String `tfsdk:"syscapenabled"`
	Autonegsupport     types.String `tfsdk:"autonegsupport"`
	Autonegenabled     types.String `tfsdk:"autonegenabled"`
	Autonegadvertised  types.String `tfsdk:"autonegadvertised"`
	Autonegmautype     types.String `tfsdk:"autonegmautype"`
	Mtu                types.String `tfsdk:"mtu"`
}

func LldpneighborsDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"ifnum": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Interface Name",
			},
			"nodeid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique number that identifies the cluster node.",
			},

			// Read-only LLDP neighbor telemetry from get(all).
			"chassisidsubtype": schema.StringAttribute{
				Computed:    true,
				Description: "Chassis ID subtype of the LLDP neighbor.",
			},
			"chassisid": schema.StringAttribute{
				Computed:    true,
				Description: "Chassis ID of the LLDP neighbor.",
			},
			"portidsubtype": schema.StringAttribute{
				Computed:    true,
				Description: "Port ID subtype of the LLDP neighbor.",
			},
			"portid": schema.StringAttribute{
				Computed:    true,
				Description: "Port ID of the LLDP neighbor.",
			},
			"ttl": schema.StringAttribute{
				Computed:    true,
				Description: "Time to live of the LLDP neighbor advertisement.",
			},
			"portdescription": schema.StringAttribute{
				Computed:    true,
				Description: "Port description of the LLDP neighbor.",
			},
			"sys": schema.StringAttribute{
				Computed:    true,
				Description: "System name of the LLDP neighbor.",
			},
			"sysdesc": schema.StringAttribute{
				Computed:    true,
				Description: "System description of the LLDP neighbor.",
			},
			"mgmtaddresssubtype": schema.StringAttribute{
				Computed:    true,
				Description: "Management address subtype of the LLDP neighbor.",
			},
			"mgmtaddress": schema.StringAttribute{
				Computed:    true,
				Description: "Management address of the LLDP neighbor.",
			},
			"iftype": schema.StringAttribute{
				Computed:    true,
				Description: "Interface type of the LLDP neighbor.",
			},
			"ifnumber": schema.StringAttribute{
				Computed:    true,
				Description: "Interface number of the LLDP neighbor.",
			},
			"vlan": schema.StringAttribute{
				Computed:    true,
				Description: "VLAN of the LLDP neighbor.",
			},
			"vlanid": schema.StringAttribute{
				Computed:    true,
				Description: "VLAN ID of the LLDP neighbor.",
			},
			"portprotosupported": schema.StringAttribute{
				Computed:    true,
				Description: "Port protocol VLANs supported by the LLDP neighbor.",
			},
			"portprotoenabled": schema.StringAttribute{
				Computed:    true,
				Description: "Port protocol VLANs enabled on the LLDP neighbor.",
			},
			"portprotoid": schema.StringAttribute{
				Computed:    true,
				Description: "Port protocol VLAN ID of the LLDP neighbor.",
			},
			"portvlanid": schema.StringAttribute{
				Computed:    true,
				Description: "Port VLAN ID of the LLDP neighbor.",
			},
			"protocolid": schema.StringAttribute{
				Computed:    true,
				Description: "Protocol ID of the LLDP neighbor.",
			},
			"linkaggrcapable": schema.StringAttribute{
				Computed:    true,
				Description: "Whether the LLDP neighbor is link-aggregation capable.",
			},
			"linkaggrenabled": schema.StringAttribute{
				Computed:    true,
				Description: "Whether link aggregation is enabled on the LLDP neighbor.",
			},
			"linkaggrid": schema.StringAttribute{
				Computed:    true,
				Description: "Link aggregation ID of the LLDP neighbor.",
			},
			"flag": schema.StringAttribute{
				Computed:    true,
				Description: "Flag of the LLDP neighbor entry.",
			},
			"syscapabilities": schema.StringAttribute{
				Computed:    true,
				Description: "System capabilities of the LLDP neighbor.",
			},
			"syscapenabled": schema.StringAttribute{
				Computed:    true,
				Description: "Enabled system capabilities of the LLDP neighbor.",
			},
			"autonegsupport": schema.StringAttribute{
				Computed:    true,
				Description: "Whether auto-negotiation is supported by the LLDP neighbor.",
			},
			"autonegenabled": schema.StringAttribute{
				Computed:    true,
				Description: "Whether auto-negotiation is enabled on the LLDP neighbor.",
			},
			"autonegadvertised": schema.StringAttribute{
				Computed:    true,
				Description: "Auto-negotiation capabilities advertised by the LLDP neighbor.",
			},
			"autonegmautype": schema.StringAttribute{
				Computed:    true,
				Description: "Auto-negotiation MAU type of the LLDP neighbor.",
			},
			"mtu": schema.StringAttribute{
				Computed:    true,
				Description: "MTU of the LLDP neighbor.",
			},
		},
	}
}
