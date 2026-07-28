package vrid6_trackinterface_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Vrid6TrackinterfaceBindingDataSourceModel describes the datasource data model.
// In addition to the resource identity attributes it exposes the read-only
// GET-response field (flags). trackinterface bindings have no vlan field.
type Vrid6TrackinterfaceBindingDataSourceModel struct {
	Id         types.String `tfsdk:"id"`
	VridId     types.Int64  `tfsdk:"vrid_id"`
	Trackifnum types.String `tfsdk:"trackifnum"`
	Flags      types.Int64  `tfsdk:"flags"`
}

func Vrid6TrackinterfaceBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"vrid_id": schema.Int64Attribute{
				Required:    true,
				Description: "Integer value that uniquely identifies a VMAC6 address.",
			},
			"trackifnum": schema.StringAttribute{
				Required:    true,
				Description: "Interfaces which need to be tracked for this vrID.",
			},
			"flags": schema.Int64Attribute{
				Computed:    true,
				Description: "Flags.",
			},
		},
	}
}
