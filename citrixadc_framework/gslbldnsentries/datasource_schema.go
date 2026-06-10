package gslbldnsentries

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

// GslbldnsentriesDataSourceSchema is a read-only, get(all)-backed filter
// datasource. nodeid is an optional GET filter; the datasource returns the first
// LDNS entry matching the supplied filter. The rich read-only output fields are
// not in tfdata and are not modelled.
func GslbldnsentriesDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"nodeid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique number that identifies the cluster node.",
			},
		},
	}
}
