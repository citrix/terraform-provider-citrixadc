package snmpengineid

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SnmpengineidDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"engineid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "A hexadecimal value of at least 10 characters, uniquely identifying the engineid",
			},
			"ownernode": schema.Int64Attribute{
				Required:    true,
				Description: "ID of the cluster node for which you are setting the engineid",
			},
		},
	}
}
