package channel_interface_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func ChannelInterfaceBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"channelid": schema.StringAttribute{
				Required:    true,
				Description: "ID of the LA channel or the cluster LA channel to which you want to bind interfaces. Specify an LA channel in LA/x notation, where x can range from 1 to 8 or a cluster LA channel in CLA/x notation or  Link redundant channel in LR/x notation , where x can range from 1 to 4.",
			},
			"ifnum": schema.ListAttribute{
				ElementType: types.StringType,
				Required:    true,
				Description: "Interfaces to be bound to the LA channel of a Citrix ADC or to the LA channel of a cluster configuration.\nFor an LA channel of a Citrix ADC, specify an interface in C/U notation (for example, 1/3).\nFor an LA channel of a cluster configuration, specify an interface in N/C/U notation (for example, 2/1/3).\nwhere C can take one of the following values:\n* 0 - Indicates a management interface.\n* 1 - Indicates a 1 Gbps port.\n* 10 - Indicates a 10 Gbps port.\nU is a unique integer for representing an interface in a particular port group.\nN is the ID of the node to which an interface belongs in a cluster configuration.\nUse spaces to separate multiple entries.",
			},
		},
	}
}
