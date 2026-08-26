package servicegroup

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func ServicegroupDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"aigwprofilename": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"appflowlog": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"autodelayedtrofs": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"autodisabledelay": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"autodisablegraceful": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"autoscale": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"bootstrap": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"cacheable": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"cachetype": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"cip": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"cipheader": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"cka": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"clttimeout": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"cmp": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"comment": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"customserverid": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"dbsttl": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"delay": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"downstateflush": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"dupweight": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"graceful": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"hashid": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"healthmonitor": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"httpprofilename": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"includemembers": schema.BoolAttribute{
				Optional: true,
				Computed: true,
			},
			"maxbandwidth": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"maxclient": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"maxreq": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"mcpprofilename": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"memberport": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"monconnectionclose": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"monitornamesvc": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"monthreshold": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"nameserver": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"netprofile": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"pathmonitor": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"pathmonitorindv": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"port": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"quicprofilename": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"riseapbrstatsmsgcode": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"rtspsessionidremap": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"serverid": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"servername": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"servicegroupname": schema.StringAttribute{
				Required: true,
			},
			"servicetype": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"sp": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"state": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"svrtimeout": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"tcpb": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"tcpprofilename": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"td": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"topicname": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"useproxyport": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"usip": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"wasmmodule": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"weight": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},

			// Convenience blocks (shared model with the resource).
			"lbvservers": schema.SetAttribute{
				Optional:    true,
				Computed:    true,
				ElementType: types.StringType,
			},
			"lbmonitor": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"servicegroupmembers": schema.SetAttribute{
				Optional:    true,
				Computed:    true,
				ElementType: types.StringType,
			},
			"servicegroupmembers_by_servername": schema.SetAttribute{
				Optional:    true,
				Computed:    true,
				ElementType: types.StringType,
			},
		},
	}
}
