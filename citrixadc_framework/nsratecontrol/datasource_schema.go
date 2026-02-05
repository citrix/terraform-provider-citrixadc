package nsratecontrol

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func NsratecontrolDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"icmpthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of ICMP packets permitted per 10 milliseconds.",
			},
			"tcprstthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The number of TCP RST packets permitted per 10 milli second. zero means rate control is disabled and 0xffffffff means every thing is rate controlled",
			},
			"tcpthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of SYNs permitted per 10 milliseconds.",
			},
			"udpthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of UDP packets permitted per 10 milliseconds.",
			},
		},
	}
}
