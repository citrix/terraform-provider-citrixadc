package lsnrtspalgsession

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func LsnrtspalgsessionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"nodeid": schema.Int64Attribute{
				Optional:    true,
				Description: "Unique number that identifies the cluster node.",
			},
			"sessionid": schema.StringAttribute{
				Required:    true,
				Description: "Session ID for the RTSP call.",
			},
		},
	}
}
