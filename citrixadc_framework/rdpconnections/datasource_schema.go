package rdpconnections

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// RdpconnectionsDataSourceModel is the datasource-only model. It intentionally
// does NOT reuse the resource model (RdpconnectionsResourceModel): the resource
// is an action-only kill with a minimal schema (id/username/all), while the
// datasource is backed by get(all) and exposes the read-only RDP connection
// telemetry. Per the NITRO doc get response, endpointport/targetport are
// Integer and peid is Double; they map to types.Int64 (populated via
// utils.ConvertToInt64, which tolerates the float64 JSON wire format). All other
// telemetry fields are strings. "all" is a kill-only selector, not a GET filter,
// so it is omitted from the datasource; username is the sole filter.
type RdpconnectionsDataSourceModel struct {
	Id       types.String `tfsdk:"id"`
	Username types.String `tfsdk:"username"`

	// Read-only RDP connection telemetry (Computed) from get(all).
	Endpointip   types.String `tfsdk:"endpointip"`
	Endpointport types.Int64  `tfsdk:"endpointport"`
	Targetip     types.String `tfsdk:"targetip"`
	Targetport   types.Int64  `tfsdk:"targetport"`
	Peid         types.Int64  `tfsdk:"peid"`
}

func RdpconnectionsDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			// Optional filter mirroring CLI "show rdp connections [-userName]".
			"username": schema.StringAttribute{
				Optional:    true,
				Description: "User name for which to display connections.",
			},

			// Read-only RDP connection telemetry from get(all).
			"endpointip": schema.StringAttribute{
				Computed:    true,
				Description: "IP address of the RDP connection endpoint.",
			},
			"endpointport": schema.Int64Attribute{
				Computed:    true,
				Description: "Port of the RDP connection endpoint (1-65535).",
			},
			"targetip": schema.StringAttribute{
				Computed:    true,
				Description: "IP address of the RDP connection target.",
			},
			"targetport": schema.Int64Attribute{
				Computed:    true,
				Description: "Port of the RDP connection target (1-65535).",
			},
			"peid": schema.Int64Attribute{
				Computed:    true,
				Description: "Packet engine (core) ID handling the RDP connection.",
			},
		},
	}
}
