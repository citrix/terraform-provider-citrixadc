package routerdynamicrouting

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func RouterdynamicroutingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"commandstring": schema.StringAttribute{
				Required:    true,
				Description: "command to be executed",
			},
			"nodeid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique number that identifies the cluster node.",
			},
		},
	}
}
