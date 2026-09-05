package vxlan

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// VxlanDataSourceModel is the data-source-specific model, decoupled from
// VxlanResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits
// (td, partitionname). Every non-key attribute is Computed; the Framework's
// per-attribute model <-> schema reflection requires this model to have exactly
// the attributes the data-source schema declares, which is why it cannot reuse
// the resource model.
type VxlanDataSourceModel struct {
	Id      types.String `tfsdk:"id"`
	Vxlanid types.Int64  `tfsdk:"vxlanid"` // Required lookup key

	// Read/write attributes, surfaced here as Computed outputs.
	Dynamicrouting     types.String `tfsdk:"dynamicrouting"`
	Innervlantagging   types.String `tfsdk:"innervlantagging"`
	Ipv6dynamicrouting types.String `tfsdk:"ipv6dynamicrouting"`
	Port               types.Int64  `tfsdk:"port"`
	Protocol           types.String `tfsdk:"protocol"`
	Type               types.String `tfsdk:"type"`
	Vlan               types.Int64  `tfsdk:"vlan"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/vxlan.json). Never settable; populated from GET.
	Td            types.Int64  `tfsdk:"td"`
	Partitionname types.String `tfsdk:"partitionname"`
}

func VxlanDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"dynamicrouting": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable dynamic routing on this VXLAN.",
			},
			"vxlanid": schema.Int64Attribute{
				Required:    true,
				Description: "A positive integer, which is also called VXLAN Network Identifier (VNI), that uniquely identifies a VXLAN.",
			},
			"innervlantagging": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies whether Citrix ADC should generate VXLAN packets with inner VLAN tag.",
			},
			"ipv6dynamicrouting": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable all IPv6 dynamic routing protocols on this VXLAN. Note: For the ENABLED setting to work, you must configure IPv6 dynamic routing protocols from the VTYSH command line.",
			},
			"port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies UDP destination port for VXLAN packets.",
			},
			"protocol": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "VXLAN-GPE next protocol. RESERVED, IPv4, IPv6, ETHERNET, NSH",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "VXLAN encapsulation type. VXLAN, VXLANGPE",
			},
			"vlan": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "ID of VLANs whose traffic is allowed over this VXLAN. If you do not specify any VLAN IDs, the Citrix ADC allows traffic of all VLANs that are not part of any other VXLANs.",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"td": schema.Int64Attribute{
				Computed:    true,
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
			"partitionname": schema.StringAttribute{
				Computed:    true,
				Description: "The Partition to which this vxlan is bound.",
			},
		},
	}
}

// vxlanDataSourceSetAttrFromGet projects a NITRO vxlan GET response onto the
// data-source model. Because a data source has no plan/apply reconciliation,
// attributes are simply filled from the GET (or left Null when the GET omits
// them) — no unknown->null resolution or plan preservation is required. The
// shared utils.MapGet* helpers implement that projection.
func vxlanDataSourceSetAttrFromGet(ctx context.Context, data *VxlanDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In vxlanDataSourceSetAttrFromGet Function")

	if v, ok := g["id"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
	}
	data.Vxlanid = utils.MapGetInt64(g, "id")

	// Read/write attributes as read-back outputs.
	data.Dynamicrouting = utils.MapGetString(g, "dynamicrouting")
	data.Innervlantagging = utils.MapGetString(g, "innervlantagging")
	data.Ipv6dynamicrouting = utils.MapGetString(g, "ipv6dynamicrouting")
	data.Port = utils.MapGetInt64(g, "port")
	data.Protocol = utils.MapGetString(g, "protocol")
	data.Type = utils.MapGetString(g, "type")
	data.Vlan = utils.MapGetInt64(g, "vlan")

	// Read-only metadata.
	data.Td = utils.MapGetInt64(g, "td")
	data.Partitionname = utils.MapGetString(g, "partitionname")
}
