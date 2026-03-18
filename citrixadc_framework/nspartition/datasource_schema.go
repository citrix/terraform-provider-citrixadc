package nspartition

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func NspartitionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"force": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Switches to new admin partition without prompt for saving configuration. Configuration will not be saved",
			},
			"maxbandwidth": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum bandwidth, in Kbps, that the partition can consume. A zero value indicates the bandwidth is unrestricted on the partition and it can consume up to the system limits.",
			},
			"maxconn": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of concurrent connections that can be open in the partition. A zero value indicates no limit on number of open connections.",
			},
			"maxmemlimit": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum memory, in megabytes, allocated to the partition.  A zero value indicates the memory is unlimited on the partition and it can consume up to the system limits.",
			},
			"minbandwidth": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum bandwidth, in Kbps, that the partition can consume. A zero value indicates the bandwidth is unrestricted on the partition and it can consume up to the system limits",
			},
			"partitionmac": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Special MAC address for the partition which is used for communication over shared vlans in this partition. If not specified, the MAC address is auto-generated.",
			},
			"partitionname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the Partition. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.",
			},
			"save": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Switches to new admin partition without prompt for saving configuration. Configuration will be saved",
			},
		},
	}
}
