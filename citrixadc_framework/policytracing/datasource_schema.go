package policytracing

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

// PolicytracingDataSourceSchema backs the get(all) datasource. transactionid,
// detail and nodeid are optional show/get filters; capturesslhandshakepolicies,
// filterexpr and protocoltype are read-only output fields returned by GET.
func PolicytracingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"capturesslhandshakepolicies": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Set it to yes if need to capture the SSL handshake policies",
			},
			"detail": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Show detailed information of the captured records",
			},
			"filterexpr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Policy tracing filter expression. For example: http.req.url.startswith(\"/this\").",
			},
			"nodeid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique number that identifies the cluster node.",
			},
			"protocoltype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "protocol type for which policy records needs to be collected",
			},
			"transactionid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique ID to identify the current transaction",
			},
		},
	}
}
