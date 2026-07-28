package vrid6_channel_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Vrid6ChannelBindingDataSourceModel describes the datasource data model.
// In addition to the resource identity attributes it exposes the read-only
// GET-response fields (flags, vlan) which are not part of the write payload.
type Vrid6ChannelBindingDataSourceModel struct {
	Id     types.String `tfsdk:"id"`
	VridId types.Int64  `tfsdk:"vrid_id"`
	Ifnum  types.String `tfsdk:"ifnum"`
	Flags  types.Int64  `tfsdk:"flags"`
	Vlan   types.Int64  `tfsdk:"vlan"`
}

func Vrid6ChannelBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"vrid_id": schema.Int64Attribute{
				Required:    true,
				Description: "Integer value that uniquely identifies a VMAC6 address.",
			},
			"ifnum": schema.StringAttribute{
				Required:    true,
				Description: "Interfaces to bind to the VMAC6, specified in (slot/port) notation (for example, 1/2).Use spaces to separate multiple entries.",
			},
			"flags": schema.Int64Attribute{
				Computed:    true,
				Description: "Flags.",
			},
			"vlan": schema.Int64Attribute{
				Computed:    true,
				Description: "The VLAN in which this VRID resides.",
			},
		},
	}
}
