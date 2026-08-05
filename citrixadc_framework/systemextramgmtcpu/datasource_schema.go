package systemextramgmtcpu

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SystemextramgmtcpuDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"enabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Boolean value indicating the effective state of the extra management CPU.",
			},
			// The datasource shares SystemextramgmtcpuResourceModel with the resource,
			// so every model field must be declared here even though these are
			// resource-only write knobs (never returned by NITRO / meaningless for a
			// datasource read).
			"reboot": schema.BoolAttribute{
				Optional: true,
				Computed: true,
			},
			"reachable_timeout": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"reachable_poll_delay": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"reachable_poll_interval": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"reachable_poll_timeout": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
		},
	}
}
