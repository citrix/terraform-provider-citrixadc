package appflowcollector

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AppflowcollectorDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"ipaddress": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IPv4 address of the collector.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the collector. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at\n(@), equals (=), and hyphen (-) characters.\n Only four collectors can be configured.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my appflow collector\" or 'my appflow collector').",
			},
			"netprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Netprofile to associate with the collector. The IP address defined in the profile is used as the source IP address for AppFlow traffic for this collector.  If you do not set this parameter, the Citrix ADC IP (NSIP) address is used as the source IP address.",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New name for the collector. Must begin with an ASCII alphabetic or underscore (_) character, and must\ncontain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at(@), equals (=), and hyphen (-) characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my appflow coll\" or 'my appflow coll').",
			},
			"port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port on which the collector listens.",
			},
			"transport": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of collector: either logstream or ipfix or rest.",
			},
		},
	}
}
