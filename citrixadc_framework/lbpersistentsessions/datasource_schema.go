package lbpersistentsessions

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

// LbpersistentsessionsDataSourceSchema is a read-only, get(all)-backed filter
// datasource. vserver and nodeid are optional filters (nodeid is a valid GET
// filter); the datasource returns the first session matching the supplied
// filters. The rich read-only output fields are not in tfdata and are not
// modelled.
func LbpersistentsessionsDataSourceSchema() schema.Schema {
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
			"persistenceparameter": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The persistence parameter whose persistence sessions are to be flushed.",
			},
			"vserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The name of the virtual server.",
			},
		},
	}
}
