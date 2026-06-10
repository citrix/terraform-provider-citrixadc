package lsnsipalgcall

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func LsnsipalgcallDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"callid": schema.StringAttribute{
				Required:    true,
				Description: "Call ID for the SIP call.",
			},
			"nodeid": schema.Int64Attribute{
				Optional:    true,
				Description: "Unique number that identifies the cluster node.",
			},
		},
	}
}
